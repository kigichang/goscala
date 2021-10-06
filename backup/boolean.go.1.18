// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package goscala

func True() bool {
	return true
}

func False() bool {
	return false
}

func Default[T any](z T) func(func() (T, bool)) T {
	return func(fn func() (T, bool)) T {
		if v, ok := fn(); ok {
			return v
		}
		return z
	}
}

func Cond[T any](ok bool, t, f T) T {
	if ok {
		return t
	}
	return f
}

func Ternary[T any](cond func() bool, succ func() T, fail func() T) T {
	if cond() {
		return succ()
	}
	return fail()
}
