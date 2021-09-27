package goscala

import "reflect"

func TypeStr(x interface{}) string {
	return reflect.TypeOf(x).String()
}

//func Fetch[T any](v T) func() T {
//	return func() T {
//		return v
//	}
//}
//
//func FetchBool[T any](v T, ok bool) func() (T, bool) {
//	return func() (T, bool) {
//		return v, ok
//	}
//}
//
//func FetchErr[T any](v T, err error) func() (T, error) {
//	return func() (T, error) {
//		return v, err
//	}
//}

//func ValueFunc[T any](v T) func() T {
//	return func() T {
//		return v
//	}
//}
//
//func ValueBoolFunc[T any](v T, ok bool) func() (T, bool) {
//	return func() (T, bool) {
//		return v, ok
//	}
//}
//
//func ValueErrFunc[T any](v T, err error) func() (T, error) {
//	return func() (T, error) {
//		return v, err
//	}
//}
//
//func Identity[T any](v T) T {
//	return Id(v)
//}
//
//func Id[T any](v T) T {
//	return v
//}
//
//func IdBool[T any](v T, ok bool) (T, bool) {
//	return v, ok
//}
//
//func IdentityBool[T any](v T, ok bool) (T, bool) {
//	return IdBool(v, ok)
//}
//
//func IdErr[T any](v T, err error) (T, error) {
//	return v, err
//}
//
//func IdentityErr[T any](v T, err error) (T, error) {
//	return IdErr(v, err)
//}

func Eq[T comparable](a, b T) bool {
	return a == b
}

func Equal[T comparable](a, b T) bool {
	return Eq(a, b)
}

func C[T any](flag bool, a, b T) T {
	if flag {
		return b
	}
	return a
}

//func Compare[T Ordered](a, b T) int {
//	if a == b {
//		return 0
//	}
//
//	if a > b {
//		return 1
//	}
//
//	return -1
//}
//
//func Max[T Ordered](a, b T) T {
//	if a > b {
//		return a
//	}
//	return b
//}
//
//func Min[T Ordered](a, b T) T {
//	if a < b {
//		return a
//	}
//	return b
//}

//func Ternary[T any](cond func() bool, succ func() T, fail func() T) T {
//	if cond() {
//		return succ()
//	}
//	return fail()
//}
//
//func True() bool {
//	return true
//}
//
//func False() bool {
//	return false
//}
