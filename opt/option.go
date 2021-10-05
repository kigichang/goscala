// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package opt

import (
	gs "github.com/kigichang/goscala"
	"github.com/kigichang/goscala/impl"
)

func Bool[T any](v T, ok bool) gs.Option[T] {
	return gs.PartialV(gs.Some[T], gs.None[T])(v, ok)
}

func Err[T any](v T, err error) gs.Option[T] {
	return gs.PartialV(gs.Some[T], gs.None[T])(v, err == nil)
}

func When[T any](cond func() bool, z T) gs.Option[T] {
	return Bool(z, cond())
}

func Unless[T any](cond func() bool, z T) gs.Option[T] {
	return Bool(z, !cond())
}

func Collect[T any, U any](opt gs.Option[T], fn func(T) (U, bool)) gs.Option[U] {
	return gs.Partial(
		gs.FuncBoolAndThen(fn, gs.PartialV(gs.Some[U], gs.None[U])),
		gs.None[U],
	)(opt.Fetch)
}

func FlatMap[T, U any](opt gs.Option[T], fn func(T) gs.Option[U]) gs.Option[U] {
	return gs.Partial(fn, gs.None[U])(opt.Fetch)
}

func Map[T, U any](opt gs.Option[T], fn func(T) U) gs.Option[U] {
	return gs.Partial(
		gs.FuncAndThen(fn, gs.Some[U]),
		gs.None[U],
	)(opt.Fetch)
}

func MapBool[T, U any](opt gs.Option[T], fn func(T) (U, bool)) gs.Option[U] {
	return gs.Partial(
		gs.FuncBoolAndThen(fn, Bool[U]),
		gs.None[U],
	)(opt.Fetch)
}

func MapErr[T, U any](opt gs.Option[T], fn func(T) (U, error)) gs.Option[U] {
	return gs.Partial(
		gs.FuncErrAndThen(fn, Err[U]),
		gs.None[U],
	)(opt.Fetch)
}

func Fold[T, U any](opt gs.Option[T]) func(U) func(func(T) U) U {
	return func(z U) func(func(T) U) U {
		return func(fn func(T) U) U {
			return gs.Partial(fn, gs.VF(z))(opt.Fetch)
		}
	}
}

func Left[T, R any](opt gs.Option[T], z R) gs.Either[T, R] {
	return gs.Partial(
		impl.Left[T, R],
		gs.FuncUnitAndThen(gs.VF(z), impl.Right[T, R]),
	)(opt.Fetch)
}

func Right[L, T any](opt gs.Option[T], z L) gs.Either[L, T] {
	return gs.Partial(
		impl.Right[L, T],
		gs.FuncUnitAndThen(gs.VF(z), impl.Left[L, T]),
	)(opt.Fetch)
}
