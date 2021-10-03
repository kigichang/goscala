// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package future

import (
	"context"
	"fmt"

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

func (f *_future[T]) Context() context.Context {
	return f.ctx
}

func (f *_future[T]) Value() gs.Try[T] {
	return f.val
}

func (f *_future[T]) Pass() context.Context {
	return withValue(f.ctx, &f.val, &f.completed)
}

func (f *_future[T]) OnComplete(fn func(gs.Try[T])) {
	go func(x *_future[T]) {
		ctx := x.Pass()
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

	go func(ap context.Context, x *_future[U]) {
		select {
		case <-ap.Done():
			v, completed := resulted[T](ap)
			if completed {
				x.val = try.Map(v, fn)
				x.completed = true
			}
			x.cancel()
		case <-x.ctx.Done():

		}
	}(a.Pass(), f)

	return f
}

func FlatMap[T, U any](ctx context.Context, a gs.Future[T], fn func(T) gs.Future[U]) gs.Future[U] {
	f := &_future[U]{}
	f.ctx, f.cancel = context.WithCancel(ctx)

	go func(ap context.Context, x *_future[U]) {
		select {
		case <-ap.Done():
			v, completed := resulted[T](ap)
			if completed {
				if v.IsSuccess() {
					g := fn(v.Success())
					go func(gp context.Context, y *_future[U]) {
						select {
						case <-gp.Done():
							v2, completed2 := resulted[U](gp)
							if completed2 {
								y.val = v2
								y.completed = true
							}
							y.cancel()
						case <-y.ctx.Done():
						}

					}(g.Pass(), f)
				} else {
					f.val = gs.Failure[U](v.Failed())
					f.completed = true
					f.cancel()
				}
			}
		case <-x.ctx.Done():

		}
	}(a.Pass(), f)

	return f
}
