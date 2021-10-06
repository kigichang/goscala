// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package either

import (
	"fmt"

	gs "github.com/kigichang/goscala"
)

type either[L, R any] struct {
	right bool
	lv    L
	rv    R
}

var _ gs.Either[int, string] = &either[int, string]{}

func (e *either[L, R]) String() string {
	if e.right {
		return fmt.Sprintf(`Right(%v)`, e.rv)
	}
	return fmt.Sprintf(`Left(%v)`, e.lv)
}

func (e *either[L, R]) Fetch() (R, bool) {
	return e.rv, e.right
}

func (e *either[L, R]) FetchErr() (R, error) {
	return e.rv, gs.Cond(e.right, nil, gs.ErrLeft)
}

func (e *either[L, R]) fetchAll() (R, L) {
	return e.rv, e.lv
}

func (e *either[L, R]) IsRight() bool {
	return e.right
}

func (e *either[L, R]) IsLeft() bool {
	return !e.right
}

func (e *either[L, R]) Left() L {
	if !e.right {
		return e.lv
	}
	panic(fmt.Errorf("can not get left value from %v", e))
}

func (e *either[L, R]) Right() R {
	if e.right {
		return e.rv
	}
	panic(fmt.Errorf("can not get right value from %v", e))
}

func (e *either[L, R]) Get() R {
	return e.Right()
}

func (e *either[L, R]) Exists(p func(R) bool) bool {
	return gs.Partial(p, gs.False)(e.Fetch)
}

func (e *either[L, R]) FilterOrElse(p func(R) bool, z L) gs.Either[L, R] {
	return gs.Partial(
		gs.Predict(Right[L, R], gs.VF(Left[L, R](z)))(p),
		gs.VF(gs.Either[L, R](e)),
	)(e.Fetch)

}

func (e *either[L, R]) Forall(p func(R) bool) bool {
	return gs.Partial(p, gs.True)(e.Fetch)
}

func (e *either[L, R]) Foreach(fn func(R)) {
	gs.Partial(gs.UnitWrap(fn), gs.Unit)(e.Fetch)
}

func (e *either[L, R]) GetOrElse(z R) R {
	return gs.Default(z)(e.Fetch)
}

func (e *either[L, R]) OrElse(z gs.Either[L, R]) gs.Either[L, R] {
	return gs.Cond(e.right, gs.Either[L, R](e), z)
}

func (e *either[L, R]) Swap() gs.Either[R, L] {
	if e.right {
		return Left[R, L](e.rv)
	}
	return Right[R, L](e.lv)
}

func Left[L, R any](v L) gs.Either[L, R] {
	return left[L, R](v)
}

func left[L, R any](v L) *either[L, R] {
	return &either[L, R]{
		right: false,
		lv:    v,
	}
}

func Right[L, R any](v R) gs.Either[L, R] {
	return right[L, R](v)
}

func right[L, R any](v R) *either[L, R] {
	return &either[L, R]{
		right: true,
		rv:    v,
	}
}
