package goscala

import (
	"fmt"
)

type Try[T any] interface {
	fmt.Stringer

	Equals(Try[T], EqualFunc[T]) bool
	
	IsSuccess() bool
	IsFailure() bool
	Get() T
	Failed() error
	Filter(Predict[T]) Try[T]
	Foreach(func(T))
	GetOrElse(z T) T
	OrElse(z Try[T]) Try[T]
	Recover(PartialFunc[error, T]) Try[T]
	RecoverWith(PartialFunc[error, Try[T]]) Try[T]

	Option() Option[T]
	Either() Either[error, T]
	Slice() Slice[T]
}

type try[T any] either[error, T]

func (t *try[T]) Either() Either[error, T] {
	return ((*either[error, T])(t))
}

func (t *try[T]) String() string {
	e := t.Either()
	if e.IsRight() {
		return fmt.Sprintf(`Success(%v)`, e.Right())
	}
	return fmt.Sprintf(`Failure(%v)`, e.Left())
}

func (t *try[T]) Equals(that Try[T], fn EqualFunc[T]) bool {
	if t == that {
		return true
	}

	if t.IsSuccess() == that.IsSuccess() {
		if t.IsSuccess() {
			return fn(t.Get(), that.Get())
		}
		return t.Failed() == that.Failed()
	}

	return false
}

func (t *try[T]) IsSuccess() bool {
	return t.Either().IsRight()
}

func (t *try[T]) IsFailure() bool {
	return t.Either().IsLeft()
}

func (t *try[T]) Get() T {
	e := t.Either()
	if e.IsRight() {
		return e.Get()
	}

	panic(fmt.Sprintf(`can not get value from %v`, t))
}

func (t *try[T]) Failed() error {
	e := t.Either()
	if e.IsLeft() {
		return e.Left()
	}
	return ErrUnsupported
}

func (t *try[T]) Filter(p Predict[T]) Try[T] {

	if !t.IsSuccess() || p(t.Get()) {
		return t
	}

	return Failure[T](ErrUnsatisfied)
}

func (t *try[T]) Foreach(f func(T)) {
	if t.IsSuccess() {
		f(t.Get())
	}
}

func (t *try[T]) GetOrElse(z T) T {
	if t.IsSuccess() {
		return t.Get()
	}
	return z
}

func (t *try[T]) OrElse(z Try[T]) Try[T] {
	if t.IsSuccess() {
		return t
	}

	return z
}

func (t *try[T]) Recover(pf PartialFunc[error, T]) Try[T] {

	if t.IsFailure() {
		if v, ok := pf(t.Failed()); ok {
			return Success[T](v)
		}
	}
	return t
}

func (t *try[T]) RecoverWith(pf PartialFunc[error, Try[T]]) Try[T] {
	if t.IsFailure() {
		if v, ok := pf(t.Failed()); ok {
			return v
		}
	}
	return t
}

func (t *try[T]) Option() Option[T] {
	if t.IsSuccess() {
		return Some[T](t.Get())
	}
	return None[T]()
}

func (t *try[T]) Slice() Slice[T] {
	if t.IsSuccess() {
		return SliceFrom(t.Get())
	}
	return SliceEmpty[T]()
}
