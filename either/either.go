package either

import (
	gs "github.com/kigichang/goscala"
)

func Bool[R any](v R, ok bool) gs.Either[bool, R] {
	return gs.Fold[R, bool, gs.Either[bool, R]](gs.VF2(v, ok))(
		gs.Left[bool, R],
		gs.Right[bool, R],
	)
}

func Err[R any](v R, err error) gs.Either[error, R] {
	return gs.Fold[R, error, gs.Either[error, R]](gs.VF2(v, err))(
		gs.Left[error, R],
		gs.Right[error, R],
	)
}

func Cond[L, R any](cond func() bool, lv L, rv R) gs.Either[L, R] {
	return gs.FoldBool[R, gs.Either[L, R]](func() (R, bool) {
		return rv, cond()
	})(
		gs.FuncUnitAndThen[L, gs.Either[L, R]](gs.VF(lv))(gs.Left[L, R]),
		gs.Right[L, R],
	)
}

func Fold[L, R, T any](e gs.Either[L, R]) func(func(L) T, func(R) T) T {
	return func(fa func(L) T, fb func(R) T) T {
		return gs.FoldBool[R, T](e.Fetch)(
			gs.FuncUnitAndThen[L, T](e.Left)(fa),
			fb,
		)
	}
}

func FlatMap[L, R, R1 any](e gs.Either[L, R]) func(func(R) gs.Either[L, R1]) gs.Either[L, R1] {
	return func(fn func(R) gs.Either[L, R1]) gs.Either[L, R1] {
		return gs.FoldBool[R, gs.Either[L, R1]](e.Fetch)(
			gs.FuncUnitAndThen[L, gs.Either[L, R1]](e.Left)(gs.Left[L, R1]),
			fn,
		)
	}
}

func Map[L, R, R1 any](e gs.Either[L, R]) func(func(R) R1) gs.Either[L, R1] {
	return func(fn func(R) R1) gs.Either[L, R1] {
		return gs.FoldBool[R, gs.Either[L, R1]](e.Fetch)(
			gs.FuncUnitAndThen[L, gs.Either[L, R1]](e.Left)(gs.Left[L, R1]),
			gs.FuncAndThen[R, R1, gs.Either[L, R1]](fn)(gs.Right[L, R1]),
		)
	}
}
