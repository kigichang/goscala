// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package either

import (
	gs "github.com/kigichang/goscala"
	"github.com/kigichang/goscala/impl"
)

// Bool returns Right of v if ok is true, or Left of false.
func Bool[R any](v R, ok bool) gs.Either[bool, R] {
	return gs.FoldV(
		impl.Right[bool, R],
		impl.Left[bool, R],
	)(v, ok)
}

// Err returns Right of v if err is nil, or Left of err.
func Err[R any](v R, err error) gs.Either[error, R] {
	return gs.FoldV(
		impl.Right[error, R],
		impl.Left[error, R],
	)(v, err)
}

// Cond returns Right of b if given cond is satisfied, or Left of a.
func Cond[L, R any](cond func() bool, a L, b R) gs.Either[L, R] {
	return gs.PartialV(
		impl.Right[L, R],
		gs.FuncUnitAndThen(gs.VF(a), impl.Left[L, R]),
	)(b, cond())
}

func Fold[L, R, T any](e gs.Either[L, R], left func(L) T, right func(R) T) T {
	return gs.Partial(
		right,
		gs.FuncUnitAndThen(e.Left, left),
	)(e.Fetch)
}

func FlatMap[L, R, R1 any](e gs.Either[L, R], fn func(R) gs.Either[L, R1]) gs.Either[L, R1] {
	return gs.Partial(
		fn,
		gs.FuncUnitAndThen(e.Left, impl.Left[L, R1]),
	)(e.Fetch)

}

func Map[L, R, R1 any](e gs.Either[L, R], fn func(R) R1) gs.Either[L, R1] {
	return gs.Partial(
		gs.FuncAndThen(fn, impl.Right[L, R1]),
		gs.FuncUnitAndThen(e.Left, impl.Left[L, R1]),
	)(e.Fetch)
}
