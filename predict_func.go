// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package goscala

func Predict[T, U any](succ func(T) U, fail func() U) func(func(T) bool) func(T) U {
	return func(p func(T) bool) func(T) U {
		return func(v T) U {
			if p(v) {
				return succ(v)
			}
			return fail()
		}
	}
}

func PredictTransform[T, U, R any](succ func(U) R, fail func(T) R) func(func(T) (U, bool)) func(T) R {
	return func(pf func(T) (U, bool)) func(T) R {
		return func(v T) R {
			if u, ok := pf(v); ok {
				return succ(u)
			}
			return fail(v)
		}
	}
}
