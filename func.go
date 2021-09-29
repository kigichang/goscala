package goscala

import "github.com/kigichang/goscala/m"

func FuncCompose[T, A, R any](f func(T) R) func(func(A) T) func(A) R {
	return m.Currying2(m.FuncCompose[T, A, R])(f)
}

func FuncAndThen[T, R, U any](f func(T) R) func(func(R) U) func(T) U {
	return m.Currying2(m.FuncAndThen[T, R, U])(f)
}

func FuncUnitAndThen[R, U any](f func() R) func(func(R) U) func() U {
	return m.Currying2(m.FuncUnitAndThen[R, U])(f)
}

func FuncBoolAndThen[T, R, U any](f func(T) (R, bool)) func(func(R, bool) U) func(T) U {
	return m.Currying2(m.FuncBoolAndThen[T, R, U])(f)
}
