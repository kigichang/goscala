// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package goscala

func PartialV[T, U any](succ func(T) U, fail func() U) func(T, bool) U {
	return func(v T, ok bool) U {
		if ok {
			return succ(v)
		}
		return fail()
	}
}

func Partial[T, U any](succ func(T) U, fail func() U) func(func() (T, bool)) U {
	return func(pf func() (T, bool)) U {
		return PartialV(succ, fail)(pf())
	}
}

func PartialErrV[T, U any](succ func(T) U, fail func(error) U) func(T, error) U {
	return func(v T, err error) U {
		if err == nil {
			return succ(v)
		}
		return fail(err)
	}
}

func PartialErr[T, U any](succ func(T) U, fail func(error) U) func(func() (T, error)) U {
	return func(pf func() (T, error)) U {
		return PartialErrV(succ, fail)(pf())
	}
}

func GetOrElse[T any](v T, ok bool) func(T) T {
	return func(z T) T {
		return Cond(ok, v, z)
	}
}
