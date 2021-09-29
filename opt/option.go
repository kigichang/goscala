package opt

import (
	gs "github.com/kigichang/goscala"
)

func Bool[T any](v T, ok bool) gs.Option[T] {
	return gs.PFV(gs.Some[T], gs.None[T])(v, ok)
}

func Err[T any](v T, err error) gs.Option[T] {
	return gs.PFV(gs.Some[T], gs.None[T])(v, err == nil)
}

func When[T any](cond func() bool) func(T) gs.Option[T] {
	return func(z T) gs.Option[T] {
		return Bool(z, cond())
	}
}

func Unless[T any](cond func() bool) func(T) gs.Option[T] {
	return func(z T) gs.Option[T] {
		return Bool(z, !cond())
	}
}

func Collect[T any, U any](opt gs.Option[T]) func(func(T) (U, bool)) gs.Option[U] {
	return func(fn func(T) (U, bool)) gs.Option[U] {
		return gs.PFF(
			gs.FuncBoolAndThen[T, U, gs.Option[U]](fn)(gs.PFV(gs.Some[U], gs.None[U])),
			gs.None[U],
		)(opt.Fetch)
	}
}

func FlatMap[T, U any](opt gs.Option[T]) func(func(T) gs.Option[U]) gs.Option[U] {
	return func(fn func(T) gs.Option[U]) gs.Option[U] {
		return gs.PFF(fn, gs.None[U])(opt.Fetch)
	}
}

func Map[T, U any](opt gs.Option[T]) func(func(T) U) gs.Option[U] {
	return func(fn func(T) U) gs.Option[U] {
		return gs.PFF(
			gs.FuncAndThen[T, U, gs.Option[U]](fn)(gs.Some[U]),
			gs.None[U],
		)(opt.Fetch)
	}
}

func Fold[T, U any](opt gs.Option[T]) func(U) func(func(T) U) U {
	return func(z U) func(func(T) U) U {
		return func(fn func(T) U) U {
			return gs.PFF(fn, gs.VF(z))(opt.Fetch)
		}
	}
}

func Left[T, R any](opt gs.Option[T]) func(R) gs.Either[T, R] {
	return func(z R) gs.Either[T, R] {
		return gs.PFF(
			gs.Left[T, R],
			gs.FuncUnitAndThen[R, gs.Either[T, R]](gs.VF(z))(gs.Right[T, R]),
		)(opt.Fetch)
	}
}

func Right[L, T any](opt gs.Option[T]) func(L) gs.Either[L, T] {
	return func(z L) gs.Either[L, T] {
		return gs.PFF(
			gs.Right[L, T],
			gs.FuncUnitAndThen[L, gs.Either[L, T]](gs.VF(z))(gs.Left[L, T]),
		)(opt.Fetch)
	}
}
