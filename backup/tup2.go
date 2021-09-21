package goscala

import (
	"fmt"
)

type Tuple2[T1, T2 any] interface {
	fmt.Stringer
	Get() (T1, T2)
	V1() T1
	V2() T2
}

type tuple2[T1, T2 any] struct {
	v1 T1
	v2 T2
}

func (t *tuple2[T1, T2]) String() string {
	return fmt.Sprintf(`(%v,%v)`, t.v1, t.v2)
}

func (t *tuple2[T1, T2]) Get() (T1, T2) {
	return t.v1, t.v2
}

func (t *tuple2[T1, T2]) V1() T1 {
	return t.v1
}

func (t *tuple2[T1, T2]) V2() T2 {
	return t.v2
}