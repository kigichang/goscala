package try

import (
	gs "github.com/kigichang/goscala"
)

func Err[T any](v T, err error) gs.Try[T] {
	return gs.PartialErrV(
		gs.Success[T],
		gs.Failure[T],
	)(v, err)
}

func Bool[T any](v T, ok bool) gs.Try[T] {
	return Err(v, gs.Cond(ok, nil, gs.ErrUnsatisfied))
}

func Collect[T, U any](t gs.Try[T], pf func(T) (U, bool)) gs.Try[U] {
	return gs.PartialErr(
		gs.FuncBoolAndThen[T, U, gs.Try[U]](pf)(Bool[U]),
		gs.Failure[U],
	)(t.FetchErr)
}

func FlatMap[T, U any](t gs.Try[T], fn func(T) gs.Try[U]) gs.Try[U] {
	return gs.PartialErr(
		fn,
		gs.Failure[U],
	)(t.FetchErr)
}

func Map[T, U any](t gs.Try[T], fn func(T) U) gs.Try[U] {
	return gs.PartialErr(
		gs.FuncAndThen[T, U, gs.Try[U]](fn)(gs.Success[U]),
		gs.Failure[U],
	)(t.FetchErr)
}

func MapErr[T, U any](t gs.Try[T], fn func(T) (U, error)) gs.Try[U] {
	return gs.PartialErr(
		gs.FuncErrAndThen[T, U, gs.Try[U]](fn)(Err[U]),
		gs.Failure[U],
	)(t.FetchErr)
}

func MapBool[T, U any](t gs.Try[T], fn func(T) (U, bool)) gs.Try[U] {
	return gs.PartialErr(
		gs.FuncBoolAndThen[T, U, gs.Try[U]](fn)(Bool[U]),
		gs.Failure[U],
	)(t.FetchErr)
}

func Fold[T, U any](t gs.Try[T], succ func(T) U, fail func(error) U) U {
	return gs.PartialErr(
		succ,
		fail,
	)(t.FetchErr)
}

func Transform[T, U any](t gs.Try[T], succ func(T) gs.Try[U], fail func(error) gs.Try[U]) gs.Try[U] {
	return gs.PartialErr(
		succ,
		fail,
	)(t.FetchErr)
}
