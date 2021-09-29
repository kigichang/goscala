package goscala

import (
	"github.com/kigichang/goscala/m"
)

func Fold[T, C, U any](fetch func() (T, C)) func(func(C) U, func(T) U) U {
	return m.Currying3To2(m.Fold[T, C, U])(fetch)
}

func FoldBool[T, U any](fetch func() (T, bool)) func(func() U, func(T) U) U {
	return m.Currying3To2(m.FoldBool[T, U])(fetch)
}

func FoldErr[T, U any](fetch func() (T, error)) func(func(error) U, func(T) U) U {
	return m.Currying3To2(m.FoldErr[T, U])(fetch)
}

func FoldLeft[T, U any](s []T) func(U) func(func(U, T) U) U {
	return m.Currying3(m.FoldLeft[T, U])(s)
}

func ScanLeft[T, U any](s []T) func(U) func(func(U, T) U) []U {
	return m.Currying3(m.ScanLeft[T, U])(s)
}

func FoldRight[T, U any](s []T) func(U) func(func(T, U) U) U {
	return m.Currying3(m.FoldRight[T, U])(s)
}

func ScanRight[T, U any](s []T) func(U) func(func(T, U) U) []U {
	return m.Currying3(m.ScanRight[T, U])(s)
}
