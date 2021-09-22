package either

import (
	"github.com/kigichang/goscala"
	"github.com/kigichang/goscala/monad"
)

func Bool[R any](v R, ok bool) goscala.Either[bool, R] {
	return monad.Fold[R, bool, goscala.Either[bool, R]](goscala.ValueBoolFunc(v, ok))(
		goscala.Left[bool, R],
		goscala.Right[bool, R],
	)
}

func Err[R any](v R, err error) goscala.Either[error, R] {
	return monad.Fold[R, error, goscala.Either[error, R]](goscala.ValueErrFunc(v, err))(
		goscala.Left[error, R],
		goscala.Right[error, R],
	)
}

func Cond[L, R any](cond func() bool, lv L, rv R) goscala.Either[L, R] {
	return monad.FoldBool[R, goscala.Either[L, R]](func() (R, bool) {
		return rv, cond()
	})(
		monad.FuncUnitAndThen[L, goscala.Either[L, R]](goscala.ValueFunc(lv))(goscala.Left[L, R]),
		goscala.Right[L, R],
	)
}

func Fold[L, R, T any](e goscala.Either[L, R]) func(func(L) T, func(R) T) T {
	return func(fa func(L) T, fb func(R) T) T {
		return monad.FoldBool[R, T](e.Fetch)(
			monad.FuncUnitAndThen[L, T](e.Left)(fa),
			fb,
		)
	}
}

func FlatMap[L, R, R1 any](e goscala.Either[L, R]) func(func(R) goscala.Either[L, R1]) goscala.Either[L, R1] {
	return func(fn func(R) goscala.Either[L, R1]) goscala.Either[L, R1] {
		return monad.FoldBool[R, goscala.Either[L, R1]](e.Fetch)(
			monad.FuncUnitAndThen[L, goscala.Either[L, R1]](e.Left)(goscala.Left[L, R1]),
			fn,
		)
	}
}

func Map[L, R, R1 any](e goscala.Either[L, R]) func(func(R) R1) goscala.Either[L, R1] {
	return func(fn func(R) R1) goscala.Either[L, R1] {
		return monad.FoldBool[R, goscala.Either[L, R1]](e.Fetch)(
			monad.FuncUnitAndThen[L, goscala.Either[L, R1]](e.Left)(goscala.Left[L, R1]),
			monad.FuncAndThen[R, R1, goscala.Either[L, R1]](fn)(goscala.Right[L, R1]),
		)
	}
}
