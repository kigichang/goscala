package goscala

import (
	"fmt"
)

type Option[T any] interface {
	fmt.Stringer
	Fetcher[T]
	IsDefined() bool
	IsEmpty() bool

	Contains(T, func(T, T) bool) bool
	Exists(func(T) bool) bool
	Equals(func(T, T) bool) func(Option[T]) bool
	Filter(func(T) bool) Option[T]
	FilterNot(func(T) bool) Option[T]
	Forall(p func(T) bool) bool
	Foreach(f func(T))
	Get() T
	GetOrElse(z T) T
	OrElse(Option[T]) Option[T]
	Slice() Slice[T]
}

type option[T any] struct {
	defined bool
	v       T
}

var _ Option[int] = &option[int]{}

func (opt *option[T]) String() string {
	if opt.defined {
		return fmt.Sprintf(`Some(%v)`, opt.v)
	}

	return fmt.Sprintf(`None[%s]`, TypeStr(opt.v))
}

func (opt *option[T]) Fetch() (T, bool) {
	return opt.v, opt.defined
}

func (opt *option[T]) FetchErr() (T, error) {
	return opt.v, Cond(opt.defined, nil, ErrEmpty)
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
	return Default(z)(opt.Fetch)
}

func (opt *option[T]) OrElse(z Option[T]) Option[T] {
	return Cond(opt.defined, Option[T](opt), z)
}

func (opt *option[T]) Contains(v T, eq func(T, T) bool) bool {
	return Partial(
		Currying2(eq)(v),
		False,
	)(opt.Fetch)

}

func (opt *option[T]) Exists(p func(T) bool) bool {
	//return Fold[T, bool, bool](opt.Fetch)(Id[bool], p)
	return opt.Filter(p).IsDefined()
}

func (opt *option[T]) Equals(eq func(T, T) bool) func(Option[T]) bool {
	return func(that Option[T]) bool {
		return Partial(
			func(x T) bool {
				return that.IsDefined() && eq(that.Get(), x)
			},
			that.IsEmpty,
		)(opt.Fetch)
	}
}

func (opt *option[T]) Filter(p func(T) bool) Option[T] {
	return Partial(
		Predict(Some[T], None[T])(p),
		None[T],
	)(opt.Fetch)
}

func (opt *option[T]) FilterNot(p func(T) bool) Option[T] {
	return opt.Filter(func(v T) bool {
		return !p(v)
	})
}

func (opt *option[T]) Forall(p func(T) bool) bool {
	return Partial(
		p,
		True,
	)(opt.Fetch)
}

func (opt *option[T]) Foreach(f func(T)) {
	Partial(UnitWrap(f), Unit)(opt.Fetch)
}

func (opt *option[T]) Slice() Slice[T] {
	return Partial(
		SliceOne[T],
		SliceEmpty[T],
	)(opt.Fetch)
}

func Some[T any](v T) Option[T] {
	return &option[T]{
		defined: true,
		v:       v,
	}
}

func None[T any]() Option[T] {
	return &option[T]{
		defined: false,
	}
}
