package goscala

import (
	"fmt"
)

type Tuple4[T1, T2, T3, T4 any] interface {
	fmt.Stringer
	Get() (T1, T2, T3, T4)
	V1() T1
	V2() T2
	V3() T3
	V4() T4
}

type tuple4[T1, T2, T3, T4 any] struct {
	v1 T1
	v2 T2
	v3 T3
	v4 T4
}

func (t *tuple4[T1, T2, T3, T4]) String() string {
	return fmt.Sprintf(`(%v,%v,%v,%v)`, t.v1, t.v2, t.v3, t.v4)
}

func (t *tuple4[T1, T2, T3, T4]) Get() (T1, T2, T3, T4) {
	return t.v1, t.v2, t.v3, t.v4
}

func (t *tuple4[T1, T2, T3, T4]) V1() T1 {
	return t.v1
}

func (t *tuple4[T1, T2, T3, T4]) V2() T2 {
	return t.v2
}

func (t *tuple4[T1, T2, T3, T4]) V3() T3 {
	return t.v3
}

func (t *tuple4[T1, T2, T3, T4]) V4() T4 {
	return t.v4
}
