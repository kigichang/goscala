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

type Fetcher[T any] interface {
	Fetch() (T, bool)
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
