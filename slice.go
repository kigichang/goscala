package goscala

import (
	"github.com/kigichang/goscala/m"
)

func EmptySlice[T any]() []T {
	return []T{}
}

func ElemSlice[T any](v T) []T {
	return []T{v}
}

func Map[T, U any](s []T) func(func(T) U) []U {
	return m.Currying2(m.Map[T, U])(s)
}

func FlatMap[T, U any](s []T) func(func(T) []U) []U {
	return m.Currying2(m.FlatMap[T, U])(s)
}
