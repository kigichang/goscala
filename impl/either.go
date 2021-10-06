package impl

import (
	"fmt"

	gs "github.com/kigichang/goscala"
)

type Either[L, R any] struct {
	OK bool
	L  L
	R  R
}

var _ gs.Either[int, int] = &Either[int, int]{}

func (e *Either[L, R]) String() string {
	if e.OK {
		return fmt.Sprintf(`Right(%v)`, e.R)
	}
	return fmt.Sprintf(`Left(%v)`, e.L)
}

func (e *Either[L, R]) Fetch() (R, bool) {
	return e.R, e.OK
}

func (e *Either[L, R]) FetchErr() (R, error) {
	return e.R, gs.Cond(e.OK, nil, gs.ErrLeft)
}

func (e *Either[L, R]) fetchAll() (R, L) {
	return e.R, e.L
}

func (e *Either[L, R]) IsRight() bool {
	return e.OK
}

func (e *Either[L, R]) IsLeft() bool {
	return !e.OK
}

func (e *Either[L, R]) Left() L {
	if !e.OK {
		return e.L
	}
	panic(fmt.Errorf("can not get left value from %v", e))
}

func (e *Either[L, R]) Right() R {
	if e.OK {
		return e.R
	}
	panic(fmt.Errorf("can not get right value from %v", e))
}

func (e *Either[L, R]) Get() R {
	return e.Right()
}

//func (e *Either[L, R]) Try() gs.Try[R] {
//	if e.IsRight() {
//		return Success[R](e.OK())
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
	return &Either[L, R]{
		OK: false,
		L:  v,
	}
}

func Right[L, R any](v R) gs.Either[L, R] {
	return &Either[L, R]{
		OK: true,
		R:  v,
	}
}
