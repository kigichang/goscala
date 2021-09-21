package goscala

import (
	"fmt"

	"github.com/kigichang/goscala/monad"
)

type Either[L, R any] interface {
	fmt.Stringer
	Fetcher[R]
	IsRight() bool
	IsLeft() bool
	Get() R
	Left() L
	Right() R
	Option() Option[R]
	Exists(func(R) bool) bool
	FilterOrElse(func(R) bool, L) Either[L, R]
	Forall(func(R) bool) bool
	Foreach(func(R))
}

type either[L, R any] struct {
	right bool
	lv    L
	rv    R
}

var _ Either[int, string] = &either[int, string]{}

func (e *either[L, R]) String() string {
	if e.right {
		return fmt.Sprintf(`Right(%v)`, e.rv)
	}
	return fmt.Sprintf(`Left(%v)`, e.lv)
}

func (e *either[L, R]) Fetch() (R, bool) {
	return e.rv, e.right
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

func (e *either[L, R]) Option() Option[R] {
	return monad.FoldBool[R, Option[R]](e.Fetch)(
		None[R],
		Some[R],
	)
}

func (e *either[L, R]) Exists(p func(R) bool) bool {
	return monad.FoldBool[R, bool](e.Fetch)(
		False,
		p,
	)
}

func (e *either[L, R]) FilterOrElse(p func(R) bool, z L) Either[L, R] {
	return monad.FoldBool[R, Either[L, R]](e.Fetch)(
		ValueFunc(Either[L, R](e)),
		func(r R) Either[L, R] {
			if p(r) {
				return e
			}
			return Left[L, R](z)
		},
	)
}

func (e *either[L, R]) Forall(p func(R) bool) bool {
	return monad.FoldBool[R, bool](e.Fetch)(
		True,
		p,
	)
}

func (e *either[L, R]) Foreach(fn func(R)) {
	monad.FoldBool[R, Unit](e.Fetch)(
		UnitFunc,
		UnitWrap(fn),
	)
}

func Left[L, R any](v L) Either[L, R] {
	return &either[L, R]{
		right: false,
		lv:    v,
	}
}

func Right[L, R any](v R) Either[L, R] {
	return &either[L, R]{
		right: true,
		rv:    v,
	}
}
