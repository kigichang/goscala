package opt

import (
	gs "github.com/kigichang/goscala"
)

//func noneBool[T any](b bool) gs.Option[T] {
//	if !b {
//		return gs.None[T]()
//	}
//	panic("can not make None with true")
//}
//
func noneErr[T any](err error) gs.Option[T] {
	if err != nil {
		return gs.None[T]()
	}
	panic("can not make None with nil error")
}

func Bool[T any](v T, ok bool) gs.Option[T] {
	return gs.FoldBool[T, gs.Option[T]](gs.VF2(v, ok))(
		gs.None[T],
		gs.Some[T],
	)
}

func Err[T any](v T, err error) gs.Option[T] {
	return gs.FoldErr[T, gs.Option[T]](gs.VF2(v, err))(noneErr[T], gs.Some[T])
}

func When[T any](cond func() bool) func(T) gs.Option[T] {
	return func(z T) gs.Option[T] {
		return gs.Ternary(
			cond,
			gs.FuncUnitAndThen[T, gs.Option[T]](gs.VF(z))(gs.Some[T]),
			gs.None[T],
		)
	}
}

func Unless[T any](cond func() bool) func(T) gs.Option[T] {
	return func(z T) gs.Option[T] {
		return gs.Ternary(
			func() bool { return !cond() },
			gs.FuncUnitAndThen[T, gs.Option[T]](gs.VF(z))(gs.Some[T]),
			gs.None[T],
		)
	}
}

func Collect[T any, U any](opt gs.Option[T]) func(func(T) (U, bool)) gs.Option[U] {
	return func(fn func(T) (U, bool)) gs.Option[U] {
		return gs.FoldBool[T, gs.Option[U]](opt.Fetch)(
			gs.None[U],
			gs.FuncBoolAndThen[T, U, gs.Option[U]](fn)(Bool[U]),
		)
	}
}

func FlatMap[T, U any](opt gs.Option[T]) func(func(T) gs.Option[U]) gs.Option[U] {
	return func(fn func(T) gs.Option[U]) gs.Option[U] {
		return gs.FoldBool[T, gs.Option[U]](opt.Fetch)(
			gs.None[U],
			fn,
		)
	}
}

func Map[T, U any](opt gs.Option[T]) func(func(T) U) gs.Option[U] {
	return func(fn func(T) U) gs.Option[U] {
		return gs.FoldBool[T, gs.Option[U]](opt.Fetch)(
			gs.None[U],
			gs.FuncAndThen[T, U, gs.Option[U]](fn)(gs.Some[U]),
		)
	}
}

func Fold[T, U any](opt gs.Option[T]) func(U) func(func(T) U) U {
	return func(z U) func(func(T) U) U {
		return func(fn func(T) U) U {
			return gs.FoldBool[T, U](opt.Fetch)(
				gs.VF(z),
				fn,
			)
		}
	}
}

func Left[T, R any](opt gs.Option[T]) func(R) gs.Either[T, R] {
	return func(z R) gs.Either[T, R] {
		return gs.FoldBool[T, gs.Either[T, R]](opt.Fetch)(
			gs.VF(gs.Right[T, R](z)),
			gs.Left[T, R],
		)
	}
}

func Right[L, T any](opt gs.Option[T]) func(L) gs.Either[L, T] {
	return func(z L) gs.Either[L, T] {
		return gs.FoldBool[T, gs.Either[L, T]](opt.Fetch)(
			gs.VF(gs.Left[L, T](z)),
			gs.Right[L, T],
		)
	}
}
