package goscala

import (
	"errors"
	"fmt"

	"github.com/kigichang/gomonad"
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

func (t *try[T]) fetchErr() (T, error) {
	return t.v, t.err
}

func (t *try[T]) Equals(eq func(T, T) bool) func(Try[T]) bool {
	return func(that Try[T]) bool {
		return gomonad.FoldErr[T, bool](t.fetchErr)(
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

func (t *try[T]) Filter(p func(T) bool) Try[T] {
	return gomonad.FoldErr[T, Try[T]](t.fetchErr)(
		Failure[T],
		gomonad.FuncAndThen[T, bool, Try[T]](p)(func(ok bool) Try[T] {
			if ok {
				return Success[T](t.v)
			}
			return Failure[T](ErrUnsatisfied)
		}),
	)
}

func (t *try[T]) Foreach(fn func(T)) {
	gomonad.FoldBool[T, struct{}](t.Fetch)(
		gomonad.Unit,
		gomonad.UnitWrap(fn),
	)
}

func (t *try[T]) GetOrElse(z T) T {
	return gomonad.FoldBool[T, T](t.Fetch)(
		gomonad.VF(z),
		gomonad.Id[T],
	)
}

func (t *try[T]) OrElse(z Try[T]) Try[T] {
	return gomonad.FoldBool[T, Try[T]](t.Fetch)(
		gomonad.VF(z),
		Success[T],
	)
}

func (t *try[T]) Recover(pf func(error) (T, bool)) Try[T] {
	return gomonad.FoldErr[T, Try[T]](t.fetchErr)(
		func(err error) Try[T] {
			if v, ok := pf(err); ok {
				return Success[T](v)
			}
			return Failure[T](err)
		},
		Success[T],
	)
}

func (t *try[T]) RecoverWith(pf func(error) (Try[T], bool)) Try[T] {
	return gomonad.FoldErr[T, Try[T]](t.fetchErr)(
		func(err error) Try[T] {
			if v, ok := pf(err); ok {
				return v
			}
			return Failure[T](err)
		},
		Success[T],
	)
}

func (t *try[T]) Option() Option[T] {
	return gomonad.FoldBool[T, Option[T]](t.Fetch)(
		None[T],
		Some[T],
	)
}

//func (t *try[T]) Either() Either[error, T] {
//	return t.either()
//}

func (t *try[T]) Slice() []T {
	return gomonad.FoldBool[T, []T](t.Fetch)(
		gomonad.EmptySlice[T],
		gomonad.ElemSlice[T],
	)
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
