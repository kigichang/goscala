package future

import (
	"context"

	gs "github.com/kigichang/goscala"
)

type key int

const (
	keyResult key = 1
)

type _result[T any] struct {
	val       *gs.Try[T]
	completed *bool
}

func (r *_result[T]) Value() gs.Try[T] {
	return *(r.val)
}

func (r *_result[T]) Completed() bool {
	return *r.completed
}

func withValue[T any](parent context.Context, v *gs.Try[T], completed *bool) context.Context {
	return context.WithValue(
		parent,
		keyResult,
		&_result[T]{
			val:       v,
			completed: completed,
		})
}

func resulted[T any](ctx context.Context) (gs.Try[T], bool) {
	r := ctx.Value(keyResult).(*_result[T])
	if r == nil {
		return nil, false
	}

	return *r.val, *r.completed
}
