package goscala

import (
	"reflect"
)

func typstr(x interface{}) string {
	return reflect.TypeOf(x).String()
}

// IsZero returns true if v is an Zero value, or returns false.
func IsZero(v interface{}) bool {
	return reflect.ValueOf(v).IsZero()
}
