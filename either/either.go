package either

import (
	gs "github.com/kigichang/goscala"
)

func Bool[R any](v R, ok bool) gs.Either[bool, R] {
	return gs.FoldV(
		gs.Right[bool, R],
		gs.Left[bool, R],
	)(v, ok)
}

func Err[R any](v R, err error) gs.Either[error, R] {
	return gs.FoldV(
		gs.Right[error, R],
		gs.Left[error, R],
	)(v, err)
}

func Cond[L, R any](cond func() bool, lv L, rv R) gs.Either[L, R] {
	return gs.PFV(
		gs.Right[L, R],
		gs.FuncUnitAndThen[L, gs.Either[L, R]](gs.VF(lv))(gs.Left[L, R]),
	)(rv, cond())
}

func _Fold[L, R, T any](e gs.Either[L, R]) func(func(L) T, func(R) T) T {
	return func(fa func(L) T, fb func(R) T) T {
		return gs.PFF(
			fb,
			gs.FuncUnitAndThen[L, T](e.Left)(fa),
		)(e.Fetch)
	}
}
func Fold[L, R, T any](left func(L) T, right func(R) T) func(gs.Either[L, R]) T {
	return func(e gs.Either[L, R]) T {
		return gs.PFF(
			right,
			gs.FuncUnitAndThen[L, T](e.Left)(left),
		)(e.Fetch)
	}
}

func FlatMap[L, R, R1 any](e gs.Either[L, R]) func(func(R) gs.Either[L, R1]) gs.Either[L, R1] {
	return func(fn func(R) gs.Either[L, R1]) gs.Either[L, R1] {
		return gs.PFF(
			fn,
			gs.FuncUnitAndThen[L, gs.Either[L, R1]](e.Left)(gs.Left[L, R1]),
		)(e.Fetch)
	}
}

func Map[L, R, R1 any](e gs.Either[L, R]) func(func(R) R1) gs.Either[L, R1] {
	return func(fn func(R) R1) gs.Either[L, R1] {
		return gs.PFF(
			gs.FuncAndThen[R, R1, gs.Either[L, R1]](fn)(gs.Right[L, R1]),
			gs.FuncUnitAndThen[L, gs.Either[L, R1]](e.Left)(gs.Left[L, R1]),
		)(e.Fetch)
	}
}
