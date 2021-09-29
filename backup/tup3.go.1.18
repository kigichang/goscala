package goscala

import (
	"fmt"
)

type Tuple3[T1, T2, T3 any] interface {
	fmt.Stringer
	Get() (T1, T2, T3)
	V1() T1
	V2() T2
	V3() T3
}

type tuple3[T1, T2, T3 any] struct {
	v1 T1
	v2 T2
	v3 T3
}

func (t *tuple3[T1, T2, T3]) String() string {
	return fmt.Sprintf(`(%v,%v,%v)`, t.v1, t.v2, t.v3)
}

func (t *tuple3[T1, T2, T3]) Get() (T1, T2, T3) {
	return t.v1, t.v2, t.v3
}

func (t *tuple3[T1, T2, T3]) V1() T1 {
	return t.v1
}

func (t *tuple3[T1, T2, T3]) V2() T2{
	return t.v2
}

func (t *tuple3[T1, T2, T3]) V3() T3{
	return t.v3
}