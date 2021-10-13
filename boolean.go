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

func Cond[T any](ok bool, t, f T) T {
	if ok {
		return t
	}
	return f
}

func Default[T any](z T) func(func() (T, bool)) T {
	return func(fn func() (T, bool)) T {
		v, ok := fn()
		return Cond(ok, v, z)
	}
}
