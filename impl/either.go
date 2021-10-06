package impl

import (
	"fmt"

	gs "github.com/kigichang/goscala"
)

type either[L, R any] struct {
	right bool
	lv    L
	rv    R
}

var _ gs.Either[int, int] = &either[int, int]{}

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

//func (e *either[L, R]) Try() gs.Try[R] {
//	if e.IsRight() {
//		return Success[R](e.Right())
//	}
//
//	var x interface{} = e.Left()
//	switch v := x.(type) {
//	case error:
//		return Failure[R](v)
//	default:
//		return Failure[R](fmt.Errorf("%v", v))
//	}
//}

func Left[L, R any](v L) gs.Either[L, R] {
	return &either[L, R]{
		right: false,
		lv:    v,
	}
}

func Right[L, R any](v R) gs.Either[L, R] {
	return &either[L, R]{
		right: true,
		rv:    v,
	}
}
