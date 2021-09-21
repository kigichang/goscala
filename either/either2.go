package either

import (
	"fmt"
	"github.com/kigichang/goscala"
)

type either[L, R any] struct {
	right bool
	lv L
	rv R
}

var _ goscala.Either2[int, string] = &either[int, string]{}

func Left[L, R any](v L) goscala.Either2[L, R] {
	return &either[L, R] {
		right: false,
		lv: v,
	}
}

func Right[L, R any](v R) goscala.Either2[L, R] {
	return &either[L, R] {
		right: true,
		rv: v,
	}
}


func (e *either[L, R]) String() string {
	if e.right {
		return fmt.Sprintf(`Right(%v)`, e.rv)
	}
	return fmt.Sprintf(`Left(%v)`, e.lv)
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

func (e *either[L, R]) Fetch() (R, bool) {
	return e.rv, e.right
}