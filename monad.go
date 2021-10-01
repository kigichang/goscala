// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package goscala

func Fold[T, C, U any](succ func(T) U, fail func(C) U) func(func() (T, C)) U {
	return func(fetch func() (T, C)) U {
		return FoldV(succ, fail)(fetch())
	}
}

func FoldV[T, C, U any](succ func(T) U, fail func(C) U) func(T, C) U {
	return func(v T, c C) U {
		var x interface{} = c
		ok := false
		switch xv := x.(type) {
		case bool:
			ok = xv
		case error:
			ok = (xv == nil)
		default:
			ok = (xv == nil)
		}

		if ok {
			return succ(v)
		}
		return fail(c)
	}
}

