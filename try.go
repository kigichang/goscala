package goscala

import (
	"errors"
	"fmt"

	"github.com/kigichang/goscala/monad"
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

	Option() Option[T]
	Either() Either[error, T]
	Slice() []T
}

type try[T any] struct {
	v T
	err error
}

var _ Try[int] = &try[int]{}

func (t *try[T]) either() *either[error, T] {
	if t.err == nil {
		return right[error, T](t.v)
	}
	return left[error, T](t.err)
}

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
	if !t.IsFailure() {
		return t.err
	}
	panic(fmt.Errorf(`can not get failed value from %v`, t))
}

func (t *try[T]) Fetch() (T, bool) {
	return t.v, t.err == nil
}

func (t *try[T]) fetchErr() (T, error) {
	return t.v, t.err
}

func (t *try[T]) Equals(eq func(T, T) bool) func(Try[T]) bool {
	return func(that Try[T]) bool {
		return monad.FoldErr[T, bool](t.fetchErr)(
			func(err error) bool {
				return that.IsFailure() && errors.Is(err, that.Failed())
			},
			func(v T) bool {
				return that.IsSuccess() && eq(v, that.Success())
			},
		)
	}
}

func (t *try[T]) Get() T {
	return t.Success()
}

func (t *try[T]) Option() Option[T] {
	return monad.FoldBool[T, Option[T]](t.Fetch)(
		None[T],
		Some[T],
	)
}

func (t *try[T]) Either() Either[error, T] {
	return t.either()
}

func (t *try[T]) Slice() []T {
	return monad.FoldBool[T, []T](t.Fetch)(
		monad.EmptySlice[T],
		monad.ElemSlice[T],
	)
}

func success[T any](v T) *try[T] {
	return &try[T] {
		v: v,
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

	return &try[T] {
		err: err,
	}
}

func Failure[T any](err error) Try[T] {
	return failure[T](err)
}