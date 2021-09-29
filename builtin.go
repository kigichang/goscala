package goscala

import (
	"constraints"
	"reflect"
)

type Fetcher[T any] interface {
	Fetch() (T, bool)
	FetchErr() (T, error)
}

func TypeStr(x interface{}) string {
	return reflect.TypeOf(x).String()
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Compare[T constraints.Ordered](a, b T) int {
	if a == b {
		return 0
	}

	if a > b {
		return 1
	}
	return -1
}

func Equal[T comparable](a, b T) bool {
	return Eq[T](a, b)
}

func Eq[T comparable](a, b T) bool {
	return a == b
}

func SliceEmpty[T any]() []T {
	return []T{}
}

func SliceOne[T any](v T) []T {
	return []T { v }
}