package goscala

import (
	"fmt"

	"github.com/kigichang/gomonad"
)

type Either[L, R any] interface {
	fmt.Stringer
	Fetcher[R]
	IsRight() bool
	IsLeft() bool
	Get() R
	Left() L
	Right() R

	Exists(func(R) bool) bool
	FilterOrElse(func(R) bool, L) Either[L, R]
	Forall(func(R) bool) bool
	Foreach(func(R))
	GetOrElse(R) R
	OrElse(Either[L, R]) Either[L, R]
	Swap() Either[R, L]
	Option() Option[R]
	Slice() []R
	// Try() Try[R]
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

func (e *either[L, R]) Option() Option[R] {
	return gomonad.FoldBool[R, Option[R]](e.Fetch)(
		None[R],
		Some[R],
	)
}

func (e *either[L, R]) Exists(p func(R) bool) bool {
	return gomonad.FoldBool[R, bool](e.Fetch)(
		False,
		p,
	)
}

func (e *either[L, R]) FilterOrElse(p func(R) bool, z L) Either[L, R] {
	return gomonad.FoldBool[R, Either[L, R]](e.Fetch)(
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
	return gomonad.FoldBool[R, bool](e.Fetch)(
		True,
		p,
	)
}

func (e *either[L, R]) Foreach(fn func(R)) {
	gomonad.FoldBool[R, Unit](e.Fetch)(
		UnitFunc,
		UnitWrap(fn),
	)
}

func (e *either[L, R]) GetOrElse(z R) R {
	return gomonad.FoldBool[R, R](e.Fetch)(
		ValueFunc(z),
		Id[R],
	)
}

func (e *either[L, R]) OrElse(z Either[L, R]) Either[L, R] {
	return gomonad.FoldBool[R, Either[L, R]](e.Fetch)(
		ValueFunc(z),
		Right[L, R],
	)
}

func (e *either[L, R]) Swap() Either[R, L] {
	if e.right {
		return Left[R, L](e.rv)
	}
	return Right[R, L](e.lv)
}

func (e *either[L, R]) Slice() []R {
	return gomonad.FoldBool[R, []R](e.Fetch)(
		gomonad.EmptySlice[R],
		gomonad.ElemSlice[R],
	)
}

//func (e *either[L, R]) try() *try[R] {
//	if e.right {
//		return success[R](e.rv)
//	}
//
//	var x interface{} = e.lv
//	switch v := x.(type) {
//	case error:
//		return failure[R](v)
//	default:
//		return failure[R](fmt.Errorf(`%v`, v))
//	}
//
//}
//
//func (e *either[L, R]) Try() Try[R] {
//	return e.try()
//}

func Left[L, R any](v L) Either[L, R] {
	return left[L, R](v)
}

func left[L, R any](v L) *either[L, R] {
	return &either[L, R]{
		right: false,
		lv:    v,
	}
}

func Right[L, R any](v R) Either[L, R] {
	return right[L, R](v)
}

func right[L, R any](v R) *either[L, R] {
	return &either[L, R]{
		right: true,
		rv:    v,
	}
}
