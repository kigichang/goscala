package goscala

import (
	"reflect"
)

func TypeStr(x interface{}) string {
	return reflect.TypeOf(x).String()
}

// IsZero returns true if v is an Zero value, or returns false.
func IsZero(v interface{}) bool {
	return reflect.ValueOf(v).IsZero()
}

func ZeroFunc[T any](v T) func() (T, bool) {
	return func() (T, bool) {
		return v, !IsZero(v)
	}
}

func Equal[T comparable](a, b T) bool {
	return a == b
}

func Compare[T Ordered](a, b T) int {
	if a == b {
		return 0
	}

	if a > b {
		return 1
	}

	return -1
}

func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Identity[T any](v T) T {
	return v
}

//func IdentityWithBool[T](v T, ok bool) (T, bool) {
//	return v, ok
//}
//
//func IdentityWithErr[T](v T, err error) (T, error) {
//	return v, err
//}