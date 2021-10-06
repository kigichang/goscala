package opt

import (
	"fmt"

	gs "github.com/kigichang/goscala"
)

type option[T any] struct {
	defined bool
	v       T
}

var _ gs.Option[int] = &option[int]{}

func (opt *option[T]) String() string {
	if opt.defined {
		return fmt.Sprintf(`Some(%v)`, opt.v)
	}

	return fmt.Sprintf(`None[%s]`, gs.TypeStr(opt.v))
}

func (opt *option[T]) Fetch() (T, bool) {
	return opt.v, opt.defined
}

func (opt *option[T]) FetchErr() (T, error) {
	return opt.v, gs.Cond(opt.defined, nil, gs.ErrEmpty)
}

func (opt *option[T]) IsDefined() bool {
	return opt.defined
}

func (opt *option[T]) IsEmpty() bool {
	return !opt.defined
}

func (opt *option[T]) Get() T {
	if opt.defined {
		return opt.v
	}
	panic(fmt.Sprintf(`can not get value from %v`, opt))
}

func (opt *option[T]) GetOrElse(z T) T {
	return gs.Default(z)(opt.Fetch)
}

func (opt *option[T]) OrElse(z gs.Option[T]) gs.Option[T] {
	return gs.Cond(opt.defined, gs.Option[T](opt), z)
}

func (opt *option[T]) Contains(eq func(T, T) bool) func(T) bool {
	return func(v T) bool {
		return gs.Partial(
			gs.Currying2(eq)(v),
			gs.False,
		)(opt.Fetch)
	}
}

func (opt *option[T]) Exists(p func(T) bool) bool {
	//return Fold[T, bool, bool](opt.Fetch)(Id[bool], p)
	return opt.Filter(p).IsDefined()
}

func (opt *option[T]) Equals(eq func(T, T) bool) func(gs.Option[T]) bool {
	return func(that gs.Option[T]) bool {
		return gs.Partial(
			func(x T) bool {
				return that.IsDefined() && eq(that.Get(), x)
			},
			that.IsEmpty,
		)(opt.Fetch)
	}
}

func (opt *option[T]) Filter(p func(T) bool) gs.Option[T] {
	return gs.Partial(
		gs.Predict(Some[T], None[T])(p),
		None[T],
	)(opt.Fetch)
}

func (opt *option[T]) FilterNot(p func(T) bool) gs.Option[T] {
	return opt.Filter(func(v T) bool {
		return !p(v)
	})
}

func (opt *option[T]) Forall(p func(T) bool) bool {
	return gs.Partial(
		p,
		gs.True,
	)(opt.Fetch)
}

func (opt *option[T]) Foreach(f func(T)) {
	gs.Partial(gs.UnitWrap(f), gs.Unit)(opt.Fetch)
}

func Some[T any](v T) gs.Option[T] {
	return &option[T]{
		defined: true,
		v:       v,
	}
}

func None[T any]() gs.Option[T] {
	return &option[T]{
		defined: false,
	}
}
