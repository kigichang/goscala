package goscala

import "fmt"

type Option[T any] interface {
	fmt.Stringer
	Fetcher[T]
	IsDefined() bool
	Get() T
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

func (opt *option[T]) Get() T {
	if opt.defined {
		return opt.v
	}
	panic(fmt.Sprintf(`can not get value from %v`, opt))
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
