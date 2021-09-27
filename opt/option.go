package opt

import (
	"github.com/kigichang/gomonad"
	"github.com/kigichang/goscala"
)

//func noneBool[T any](b bool) goscala.Option[T] {
//	if !b {
//		return goscala.None[T]()
//	}
//	panic("can not make None with true")
//}
//
func noneErr[T any](err error) goscala.Option[T] {
	if err != nil {
		return goscala.None[T]()
	}
	panic("can not make None with nil error")
}

func Bool[T any](v T, ok bool) goscala.Option[T] {
	return gomonad.FoldBool[T, goscala.Option[T]](gomonad.VF2(v, ok))(
		goscala.None[T],
		goscala.Some[T],
	)
}

func Err[T any](v T, err error) goscala.Option[T] {
	return gomonad.FoldErr[T, goscala.Option[T]](gomonad.VF2(v, err))(noneErr[T], goscala.Some[T])
}

func When[T any](cond func() bool) func(T) goscala.Option[T] {
	return func(z T) goscala.Option[T] {
		return gomonad.Ternary(
			cond,
			gomonad.FuncUnitAndThen[T, goscala.Option[T]](gomonad.VF(z))(goscala.Some[T]),
			goscala.None[T],
		)
	}
}

func Unless[T any](cond func() bool) func(T) goscala.Option[T] {
	return func(z T) goscala.Option[T] {
		return gomonad.Ternary(
			func() bool { return !cond() },
			gomonad.FuncUnitAndThen[T, goscala.Option[T]](gomonad.VF(z))(goscala.Some[T]),
			goscala.None[T],
		)
	}
}

func Collect[T any, U any](opt goscala.Option[T]) func(func(T) (U, bool)) goscala.Option[U] {
	return func(fn func(T) (U, bool)) goscala.Option[U] {
		return gomonad.FoldBool[T, goscala.Option[U]](opt.Fetch)(
			goscala.None[U],
			gomonad.FuncBoolAndThen[T, U, goscala.Option[U]](fn)(Bool[U]),
		)
	}
}

func FlatMap[T, U any](opt goscala.Option[T]) func(func(T) goscala.Option[U]) goscala.Option[U] {
	return func(fn func(T) goscala.Option[U]) goscala.Option[U] {
		return gomonad.FoldBool[T, goscala.Option[U]](opt.Fetch)(
			goscala.None[U],
			fn,
		)
	}
}

func Map[T, U any](opt goscala.Option[T]) func(func(T) U) goscala.Option[U] {
	return func(fn func(T) U) goscala.Option[U] {
		return gomonad.FoldBool[T, goscala.Option[U]](opt.Fetch)(
			goscala.None[U],
			gomonad.FuncAndThen[T, U, goscala.Option[U]](fn)(goscala.Some[U]),
		)
	}
}

func Fold[T, U any](opt goscala.Option[T]) func(U) func(func(T) U) U {
	return func(z U) func(func(T) U) U {
		return func(fn func(T) U) U {
			return gomonad.FoldBool[T, U](opt.Fetch)(
				gomonad.VF(z),
				fn,
			)
		}
	}
}

func Left[T, R any](opt goscala.Option[T]) func(R) goscala.Either[T, R] {
	return func(z R) goscala.Either[T, R] {
		return gomonad.FoldBool[T, goscala.Either[T, R]](opt.Fetch)(
			gomonad.VF(goscala.Right[T, R](z)),
			goscala.Left[T, R],
		)
	}
}

func Right[L, T any](opt goscala.Option[T]) func(L) goscala.Either[L, T] {
	return func(z L) goscala.Either[L, T] {
		return gomonad.FoldBool[T, goscala.Either[L, T]](opt.Fetch)(
			gomonad.VF(goscala.Left[L, T](z)),
			goscala.Right[L, T],
		)
	}
}
