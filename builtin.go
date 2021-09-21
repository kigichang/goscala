package goscala

import "reflect"

func TypeStr(x interface{}) string {
	return reflect.TypeOf(x).String()
}

func Identity[T any](v T) func() T {
	return func() T {
		return v
	}
}

func IdentityBool[T any](v T, ok bool) func() (T, bool) {
	return func() (T, bool) {
		return v, ok
	}
}

func IdentityErr[T any](v T, err error) func() (T, error) {
	return func() (T, error) {
		return v, err
	}
}
