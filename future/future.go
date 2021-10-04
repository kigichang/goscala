// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package future

import (
	"context"
	"fmt"
	"time"

	gs "github.com/kigichang/goscala"
	"github.com/kigichang/goscala/try"
)

type _future[T any] struct {
	ctx       context.Context
	cancel    context.CancelFunc
	completed bool
	val       gs.Try[T]
}

var _ gs.Future[int] = &_future[int]{}

func (f *_future[T]) String() string {
	if f.completed {
		return fmt.Sprintf(`Future(%v)`, f.val)
	}

	return `Future(?)`
}

func (f *_future[T]) Completed() bool {
	return f.completed
}

func (f *_future[T]) PassValue() context.Context {
	return withValue(f.ctx, &f.val, &f.completed)
}

func (f *_future[T]) OnComplete(fn func(gs.Try[T])) {
	go func(x *_future[T]) {
		ctx := x.PassValue()
		<-ctx.Done()
		v, compleleted := resulted[T](ctx)
		if compleleted {
			fn(v)
		}
	}(f)
}

func (f *_future[T]) Foreach(fn func(T)) {
	f.OnComplete(func(v gs.Try[T]) {
		v.Foreach(fn)
	})
}

func (f *_future[T]) Wait() {
	if f.completed {
		return
	}
	<-f.ctx.Done()
}

func (f *_future[T]) Result(atMost time.Duration) (ret T, err error) {
	wait, cancel := context.WithTimeout(context.Background(), atMost)
	defer cancel()

	select {
	case <-f.ctx.Done():
		if f.completed {
			ret, err = f.val.FetchErr()
			return
		}
		err = f.ctx.Err()
	case <-wait.Done():
		err = wait.Err()
	}
	return
}

func (f *_future[T]) Filter(ctx context.Context, p func(T) bool) gs.Future[T] {
	return TransformWith[T, T](ctx, f, func(a gs.Try[T]) gs.Future[T] {
		if a.IsSuccess() {
			if p(a.Success()) {
				return Err[T](func() (ret T, _ error) {
					ret = a.Success()
					return
				})
			}

			return Err[T](func() (_ T, err error) {
				err = gs.ErrUnsatisfied
				return
			})
		}

		return Err[T](func() (_ T, err error) {
			err = a.Failed()
			return
		})
	})
}

func future[T any]() *_future[T] {
	f := &_future[T]{}
	f.ctx, f.cancel = context.WithCancel(context.Background())
	return f
}

func Make[T any](fn func() T) gs.Future[T] {
	f := future[T]()
	go func(x *_future[T]) {
		defer func() {
			if r := recover(); r != nil {
				switch rv := r.(type) {
				case error:
					x.val = gs.Failure[T](rv)
				default:
					x.val = gs.Failure[T](fmt.Errorf(`%v`, rv))
				}
			}

			x.completed = true
			x.cancel()
		}()
		x.val = gs.Success[T](fn())
	}(f)
	return f
}

func Err[T any](fn func() (T, error)) gs.Future[T] {
	f := future[T]()

	go func(x *_future[T]) {
		x.val = try.Err(fn())
		x.completed = true
		x.cancel()
	}(f)
	return f
}

func Map[T, U any](ctx context.Context, a gs.Future[T], fn func(T) U) gs.Future[U] {
	return Transform(ctx, a, func(x gs.Try[T]) gs.Try[U] {
		return try.Map(x, fn)
	})
}

func MapErr[T, U any](ctx context.Context, a gs.Future[T], fn func(T) (U, error)) gs.Future[U] {
	return Transform(ctx, a, func(x gs.Try[T]) gs.Try[U] {
		return try.MapErr(x, fn)
	})
}

func MapBool[T, U any](ctx context.Context, a gs.Future[T], fn func(T) (U, bool)) gs.Future[U] {
	return Transform(ctx, a, func(x gs.Try[T]) gs.Try[U] {
		return try.MapBool(x, fn)
	})
}

func FlatMap[T, U any](ctx context.Context, a gs.Future[T], fn func(T) gs.Future[U]) gs.Future[U] {
	return TransformWith(ctx, a, func(x gs.Try[T]) gs.Future[U] {
		if x.IsSuccess() {
			return fn(x.Success())
		}

		return Err[U](func() (ret U, err error) {
			err = x.Failed()
			return
		})
	})
}

func Transform[T, U any](ctx context.Context, a gs.Future[T], fn func(gs.Try[T]) gs.Try[U]) gs.Future[U] {
	f := future[U]()

	go func(apv context.Context, x *_future[U]) {
		select {
		case <-apv.Done():
			v, completed := resulted[T](apv)
			if completed {
				f.val = fn(v)
				f.completed = true
			}
		case <-ctx.Done():
		}
		x.cancel()
	}(a.PassValue(), f)

	return f
}

func TransformWith[T, U any](ctx context.Context, a gs.Future[T], fn func(gs.Try[T]) gs.Future[U]) gs.Future[U] {
	f := future[U]()

	go func(apv context.Context, x *_future[U]) {
		select {
		case <-apv.Done():
			v, completed := resulted[T](apv)
			if completed {
				b := fn(v)
				go func(bpv context.Context, y *_future[U]) {
					select {
					case <-bpv.Done():
						v2, completed2 := resulted[U](bpv)
						if completed2 {
							y.val = v2
							y.completed = true
						}
					case <-ctx.Done():
						// maybe cancelled.
					}
					y.cancel()
				}(b.PassValue(), x)
			} else {
				x.cancel()
			}
		case <-ctx.Done():
			// maybe cancelled.
			x.cancel()
		}
	}(a.PassValue(), f)

	return f
}
