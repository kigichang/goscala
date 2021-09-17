package goscala

import (
	"fmt"
)

// Left returns left subtype of Either contains left value with type L.
func Left[L, R any](left L) Either[L, R] {
	return &either[L, R] {
		right: false,
		lv: left,
	}
}

// Right returns right subtype of Either contains right value with type R.
func Right[L, R any](right R) Either[L, R] {
	return &either[L, R] {
		right: true,
		rv: right,
	}
}

// MakeEither returns Right if right is not an zero value, 
// or returns Left if left is not a zero value.
// It is panic when both left and right are zero values.
func MakeEither[L, R any](left L, right R) Either[L, R] {
	if !IsZero(right) {
		return Right[L, R](right)
	}

	if !IsZero(left) {
		return Left[L, R](left)
	}

	panic("all arguments are zero-value")
}

// MakeEitherWithBool returns Right contains v if ok is true, 
// or return Left contains false.
func MakeEitherWithBool[T any](v T, ok bool) Either[bool, T] {
	if ok {
		return Right[bool, T](v)
	}

	return Left[bool, T](ok)
}

// MakeEitherWithErr returns Right contains v if err is nil,
// or returns Left contains err.
func MakeEitherWithErr[T any](v T, err error) Either[error, T] {
	if err == nil {
		return Right[error, T](v)
	}
	return Left[error, T](err)
}

// EitherCond returns Right contains rv if cond is satisfied,
// or returns Left contains lv.
func EitherCond[L, R any](cond Condition, lv L, rv R) Either[L, R] {
	if cond() {
		return Right[L, R](rv)
	}
	return Left[L, R](lv)
}

// EitherFlatMap binds function f if e is Right.
func EitherFlatMap[L, R, R1 any](e Either[L, R], f Func1[R, Either[L, R1]]) Either[L, R1] {
	if e.IsRight() {
		return f(e.Right())
	}

	return Left[L, R1](e.Left())
}

// EitherFold applies fb if e is Right, or applies fa.
func EitherFold[L, R, T any](e Either[L, R], fa Func1[L, T], fb Func1[R, T]) T {
	if e.IsRight() {
		return fb(e.Right())
	}
	return fa(e.Left())
}

// EitherMap applies function f if e is Right.
func EitherMap[L, R, R1 any](e Either[L, R], f Func1[R, R1]) Either[L, R1] {
	if e.IsRight() {
		return Right[L, R1](f(e.Right()))
	}
	return Left[L, R1](e.Left())
}

// EitherToTry applies function fn if e is Right, and returns Try contains the result.
func EitherToTry[L, R, R1 any](e Either[L, R], fn Func1[R, R1]) Try[R1] {
	if e.IsRight() {
		return Success[R1](fn(e.Right()))
	}

	var x interface{} = e.Left()
	switch v := x.(type) {
	case error:
		return Failure[R1](v)
	default:
		return Failure[R1](fmt.Errorf(`%v`, e.Left()))
	}
}