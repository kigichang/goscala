package goscala

import (
	"errors"
	"fmt"
)

type Try[T any] interface {
	fmt.Stringer
	Fetcher[T]
	IsSuccess() bool
	Success() T
	Failed() error
	IsFailure() bool
	Equals(func(T, T) bool) func(Try[T]) bool
	Get() T
	Filter(func(T) bool) Try[T]
	Foreach(func(T))
	GetOrElse(z T) T
	OrElse(z Try[T]) Try[T]
	Recover(func(error) (T, bool)) Try[T]
	RecoverWith(func(error) (Try[T], bool)) Try[T]

	Option() Option[T]
	//Either() Either[error, T]
	Slice() []T
}

type try[T any] struct {
	v   T
	err error
}

var _ Try[int] = &try[int]{}

//func (t *try[T]) either() *either[error, T] {
//	if t.err == nil {
//		return right[error, T](t.v)
//	}
//	return left[error, T](t.err)
//}

func (t *try[T]) String() string {
	if t.IsSuccess() {
		return fmt.Sprintf(`Success(%v)`, t.v)
	}
	return fmt.Sprintf(`Failure(%v)`, t.err)
}

func (t *try[T]) IsSuccess() bool {
	return t.err == nil
}

func (t *try[T]) Success() T {
	if t.IsSuccess() {
		return t.v
	}
	panic(fmt.Errorf(`can not get success value from %v`, t))
}

func (t *try[T]) IsFailure() bool {
	return t.err != nil
}

func (t *try[T]) Failed() error {
	if t.IsFailure() {
		return t.err
	}
	return ErrUnsupported
}

func (t *try[T]) Fetch() (T, bool) {
	return t.v, t.err == nil
}

func (t *try[T]) FetchErr() (T, error) {
	return t.v, t.err
}

func (t *try[T]) Equals(eq func(T, T) bool) func(Try[T]) bool {
	return func(that Try[T]) bool {
		return PFErrF(
			func(v T) bool {
				return that.IsSuccess() && eq(v, that.Success())
			},
			func(err error) bool {
				return that.IsFailure() && errors.Is(err, that.Failed())
			},
		)(t.FetchErr)
	}
}

func (t *try[T]) Get() T {
	return t.Success()
}

func (t *try[T]) Filter(p func(T) bool) Try[T] {
	return PFErrF(
		Predict(
			Success[T],
			VF(Failure[T](ErrUnsatisfied)),
		)(p),
		Failure[T],
	)(t.FetchErr)
}

func (t *try[T]) Foreach(fn func(T)) {
	PFF(
		UnitWrap(fn),
		Unit,
	)(t.Fetch)
}

func (t *try[T]) GetOrElse(z T) T {
	return Default(z)(t.Fetch)
}

func (t *try[T]) OrElse(z Try[T]) Try[T] {
	return Cond(t.IsSuccess(), Try[T](t), z)
}

func (t *try[T]) Recover(pf func(error) (T, bool)) Try[T] {
	return PFErrF(
		Success[T],
		PredictTransform(Success[T], Failure[T])(pf),
	)(t.FetchErr)
}

func (t *try[T]) RecoverWith(pf func(error) (Try[T], bool)) Try[T] {
	return PFErrF(
		Success[T],
		PredictTransform(Id[Try[T]], Failure[T])(pf),
	)(t.FetchErr)
}

func (t *try[T]) Option() Option[T] {
	return PFF(
		Some[T],
		None[T],
	)(t.Fetch)
}

//func (t *try[T]) Either() Either[error, T] {
//	return t.either()
//}

func (t *try[T]) Slice() []T {
	return PFF(
		ElemSlice[T],
		EmptySlice[T],
	)(t.Fetch)
}

func success[T any](v T) *try[T] {
	return &try[T]{
		v:   v,
		err: nil,
	}
}

func Success[T any](v T) Try[T] {
	return success[T](v)
}

func failure[T any](err error) *try[T] {
	if err == nil {
		panic(fmt.Errorf("can not fail with nil error"))
	}

	return &try[T]{
		err: err,
	}
}

func Failure[T any](err error) Try[T] {
	return failure[T](err)
}
