package goscala

import (
	"fmt"
)

var (
	ErrZeroValue    = fmt.Errorf("zero-value")
	ErrUnsupported  = fmt.Errorf("unsupported")
	ErrNil          = fmt.Errorf("nil")
	ErrNone         = fmt.Errorf("none")
	ErrNotSatisfied = fmt.Errorf("unsatisfied")
	ErrFalse        = fmt.Errorf("false")
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string | ~uintptr
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

//type VOrF[T any] interface {
//	type T | func() T
//}
//
//
//func GetValue[T VOrF[T]](x T) T {
//	if reflect.TypeOf(x).Kind() == reflect.Func {
//		values := reflect.ValueOf(x).Call(nil)
//		return values[0].Interface().(T)
//	}
//
//	return x
//}