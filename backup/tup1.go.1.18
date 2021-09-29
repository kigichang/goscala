package goscala

import (
	"fmt"
)

type Tuple1[T any] interface {
	fmt.Stringer
	Get() T
	V1() T
}


type tuple1[T any] struct {
	v T
}

func (t *tuple1[T]) String() string {
	return fmt.Sprintf(`(%v)`, t.v)
}

func (t *tuple1[T]) Get() T {
	return t.v
}

func (t *tuple1[T]) V1() T {
	return t.Get()
}