package opt

import (
	"github.com/kigichang/goscala"
	"github.com/kigichang/goscala/monad"
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

func MakeWithBool[T any](v T, ok bool) goscala.Option[T] {
	return monad.FoldBool[T, goscala.Option[T]](goscala.ValueBoolFunc(v, ok))(
		goscala.None[T], 
		goscala.Some[T],
	)
}

func MakeWithErr[T any](v T, err error) goscala.Option[T] {
	return monad.FoldErr[T, goscala.Option[T]](goscala.ValueErrFunc(v, err))(noneErr[T], goscala.Some[T])
}

func When[T any](cond func() bool) func(T) goscala.Option[T] {
	return func(z T) goscala.Option[T] {
		return goscala.Ternary(
			cond, 
			monad.FuncUnitAndThen[T, goscala.Option[T]](goscala.ValueFunc(z))(goscala.Some[T]),
			goscala.None[T],
		)
	}
}

func Unless[T any](cond func() bool) func(T) goscala.Option[T] {
	return func(z T) goscala.Option[T] {
		return goscala.Ternary(
			func() bool { return !cond() }, 
			monad.FuncUnitAndThen[T, goscala.Option[T]](goscala.ValueFunc(z))(goscala.Some[T]),
			goscala.None[T],
		)
	}
}

func Collect[T any, U any](opt goscala.Option[T]) func(func(T) (U, bool)) goscala.Option[U] {
	return func(fn func(T) (U, bool)) goscala.Option[U] {
		return monad.FoldBool[T, goscala.Option[U]](opt.Fetch)(
			goscala.None[U],
			monad.FuncBoolAndThen[T, U, goscala.Option[U]](fn)(MakeWithBool[U]),
		)
	}
}

func FlatMap[T, U any](opt goscala.Option[T]) func(func(T) goscala.Option[U]) goscala.Option[U] {
	return func(fn func(T) goscala.Option[U]) goscala.Option[U] {
		return monad.FoldBool[T, goscala.Option[U]](opt.Fetch)(
			goscala.None[U],
			fn,
		)
	}
}

func Map[T, U any](opt goscala.Option[T]) func(func(T) U) goscala.Option[U] {
	return func(fn func(T) U) goscala.Option[U] {
		return monad.FoldBool[T, goscala.Option[U]](opt.Fetch)(
			goscala.None[U],
			monad.FuncAndThen[T, U, goscala.Option[U]](fn)(goscala.Some[U]),
		)
	}
}

func Fold[T, U any](opt goscala.Option[T]) func(U) func(func(T) U) U {
	return func(z U) func(func(T) U) U {
		return func (fn func(T) U) U {
			return monad.FoldBool[T, U](opt.Fetch)(
				goscala.ValueFunc(z),
				fn,
			)
		}
	}
}

func Left[T, R any](opt goscala.Option[T]) func(R) goscala.Either[T, R] {
	return func(z R) goscala.Either[T, R] {
		return monad.FoldBool[T, goscala.Either[T, R]](opt.Fetch)(
			goscala.ValueFunc(goscala.Right[T, R](z)),
			goscala.Left[T, R],
		)
	}
}

func Right[L, T any](opt goscala.Option[T]) func(L) goscala.Either[L, T] {
	return func(z L) goscala.Either[L, T] {
		return monad.FoldBool[T, goscala.Either[L, T]](opt.Fetch)(
			goscala.ValueFunc(goscala.Left[L, T](z)),
			goscala.Right[L, T],
		)
	}
}

