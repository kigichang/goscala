package goscala

import "reflect"

func TypeStr(x interface{}) string {
	return reflect.TypeOf(x).String()
}
