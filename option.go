package goscala

import (
	"fmt"
)

type Option[T any] interface {
	fmt.Stringer

	Contains(T, EqualFunc[T]) bool
	Exists(p Predict[T]) bool
	Equals(Option[T], EqualFunc[T]) bool
	Filter(p Predict[T]) Option[T]
	FilterNot(p Predict[T]) Option[T]
	Forall(p Predict[T]) bool
	Foreach(f func(T))
	Get() T
	GetOrElse(z T) T
	IsDefined() bool
	OrElse(z Option[T]) Option[T]
	Slice() Slice[T]
	//Try() Try[T]
}

type option[T any] struct {
	defined bool
	v       T
}

func (o *option[T]) String() string {
	if o.defined {
		return fmt.Sprintf(`Some(%v)`, o.v)
	}
	return fmt.Sprintf(`None(%s)`, typstr(o.v))
}

func (o *option[T]) Contains(elem T, fn EqualFunc[T]) bool {
	return o.defined && fn(o.v, elem)
}

func (o *option[T]) Exists(p Predict[T]) bool {
	return o.defined && p(o.v)
}

func (o *option[T]) Equals(that Option[T], fn EqualFunc[T]) bool {
	if o == that {
		return true
	}

	if o.IsDefined() == that.IsDefined() {
		return !o.IsDefined() || fn(o.Get(), that.Get())
	}

	return false
}

func (o *option[T]) Filter(p Predict[T]) Option[T] {
	if !o.defined || p(o.v) {
		return o
	}

	return None[T]()
}

func (o *option[T]) FilterNot(p Predict[T]) Option[T] {
	if !o.defined || !p(o.v) {
		return o
	}

	return None[T]()
}

func (o *option[T]) Forall(p Predict[T]) bool {
	if o.defined {
		return p(o.v)
	}
	return true
}

func (o *option[T]) Foreach(f func(T)) {
	if o.defined {
		f(o.v)
	}
}

func (o *option[T]) Get() T {
	if o.defined {
		return o.v
	}
	panic(fmt.Sprintf(`can not get value from %v`, o))
}

func (o *option[T]) GetOrElse(z T) T {
	if o.defined {
		return o.v
	}
	return z
}

func (o *option[T]) IsDefined() bool {
	return o.defined
}

func (o *option[T]) OrElse(z Option[T]) Option[T] {
	if o.defined {
		return o
	}
	return z
}

func (o *option[T]) Slice() Slice[T] {
	if o.defined {
		return SliceFrom(o.v)
	}
	return SliceEmpty[T]()
}
