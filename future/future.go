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

//func (f *_future[T]) Context() context.Context {
//	return f.ctx
//}

//func (f *_future[T]) Value() gs.Try[T] {
//	return f.val
//}

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

func Err[T any](ctx context.Context, fn func() (T, error)) gs.Future[T] {
	f := &_future[T]{}
	f.ctx, f.cancel = context.WithCancel(ctx)

	go func(x *_future[T]) {
		x.val = try.Err(fn())
		x.completed = true
		x.cancel()
	}(f)
	return f
}

func Map[T, U any](ctx context.Context, a gs.Future[T], fn func(T) U) gs.Future[U] {
	f := &_future[U]{}
	f.ctx, f.cancel = context.WithCancel(ctx)

	go func(apv context.Context, x *_future[U]) {
		select {
		case <-apv.Done():
			v, completed := resulted[T](apv)
			if completed {
				x.val = try.Map(v, fn)
				x.completed = true
			}
			x.cancel()
		case <-x.ctx.Done():
			// maybe cancelled.
		}
	}(a.PassValue(), f)

	return f
}

func FlatMap[T, U any](ctx context.Context, a gs.Future[T], fn func(T) gs.Future[U]) gs.Future[U] {
	f := &_future[U]{}
	f.ctx, f.cancel = context.WithCancel(ctx)

	go func(apv context.Context, x *_future[U]) {
		select {
		case <-apv.Done():
			v, completed := resulted[T](apv)
			if completed {
				if v.IsSuccess() {
					g := fn(v.Success())
					go func(gpv context.Context, y *_future[U]) {
						select {
						case <-gpv.Done():
							v2, completed2 := resulted[U](gpv)
							if completed2 {
								y.val = v2
								y.completed = true
							}
							y.cancel()
						case <-y.ctx.Done():
							// maybe concelled
						}

					}(g.PassValue(), f)
				} else {
					x.val = gs.Failure[U](v.Failed())
					x.completed = true
					x.cancel()
				}
			} else {
				x.cancel()
			}
		case <-x.ctx.Done():
			// maybe cancelled
		}
	}(a.PassValue(), f)

	return f
}
