package goscala

import (
	"fmt"

	"github.com/kigichang/goscala/monad"
)

type Option[T any] interface {
	fmt.Stringer
	Fetcher[T]
	IsDefined() bool
	IsEmpty() bool

	Contains(func(T, T) bool) func(T) bool
	Exists(func(T) bool) bool
	Equals(func(T, T) bool) func(Option[T]) bool
	Filter(func(T) bool) Option[T]
	FilterNot(func(T) bool) Option[T]
	Forall(p func(T) bool) bool
	Foreach(f func(T))
	Get() T
	GetOrElse(z T) T
	OrElse(Option[T]) Option[T]
	Slice() []T
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
	return monad.FoldBool[T, T](opt.Fetch)(
		ValueFunc[T](z),
		Id[T],
	)
}

func (opt *option[T]) OrElse(z Option[T]) Option[T] {
	return monad.FoldBool[T, Option[T]](opt.Fetch)(
		ValueFunc(z),
		Some[T],
	)
}

func (opt *option[T]) Contains(eq func(T, T) bool) func(T) bool {
	return func(v T) bool {
		return monad.FoldBool[T, bool](opt.Fetch)(False, func(x T) bool { return eq(v, x) })
	}
}

func (opt *option[T]) Exists(p func(T) bool) bool {
	//return monad.Fold[T, bool, bool](opt.Fetch)(Id[bool], p)
	return opt.Filter(p).IsDefined()
}

func (opt *option[T]) Equals(eq func(T, T) bool) func(Option[T]) bool {
	return func(that Option[T]) bool {
		return monad.FoldBool[T, bool](opt.Fetch)(
			that.IsEmpty,
			func(x T) bool {
				return that.IsDefined() && eq(that.Get(), x)
			},
		)
	}
}

func (opt *option[T]) Filter(p func(T) bool) Option[T] {
	return monad.FoldBool[T, Option[T]](opt.Fetch)(
		None[T],
		func(x T) Option[T] {
			if p(x) {
				return opt
			}
			return None[T]()
		},
	)
}

func (opt *option[T]) FilterNot(p func(T) bool) Option[T] {
	return opt.Filter(func(v T) bool {
		return !p(v)
	})
}

func (opt *option[T]) Forall(p func(T) bool) bool {
	return monad.FoldBool[T, bool](opt.Fetch)(
		True,
		p,
	)
}

func (opt *option[T]) Foreach(f func(T)) {
	monad.FoldBool[T, Unit](opt.Fetch)(
		UnitFunc,
		UnitWrap(f),
	)
}

func (opt *option[T]) Slice() []T {
	return monad.FoldBool[T, []T](opt.Fetch)(
		monad.EmptySlice[T],
		monad.ElemSlice[T],
	)
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
