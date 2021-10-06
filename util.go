// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package goscala

import "reflect"

func GetOrElse[T any](v T, ok bool) func(T) T {
	return func(z T) T {
		return Cond(ok, v, z)
	}
}

func TypeStr(x interface{}) string {
	return reflect.TypeOf(x).String()
}
