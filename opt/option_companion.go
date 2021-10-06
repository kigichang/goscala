// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package opt

import (
	gs "github.com/kigichang/goscala"
)

func Bool[T any](v T, ok bool) gs.Option[T] {
	return gs.PartialV(Some[T], None[T])(v, ok)
}

func Err[T any](v T, err error) gs.Option[T] {
	return gs.PartialV(Some[T], None[T])(v, err == nil)
}

func When[T any](cond func() bool, z T) gs.Option[T] {
	return Bool(z, cond())
}

func Unless[T any](cond func() bool, z T) gs.Option[T] {
	return Bool(z, !cond())
}

func Collect[T any, U any](opt gs.Option[T], fn func(T) (U, bool)) gs.Option[U] {
	return gs.Partial(
		gs.FuncBoolAndThen(fn, gs.PartialV(Some[U], None[U])),
		None[U],
	)(opt.Fetch)
}

func FlatMap[T, U any](opt gs.Option[T], fn func(T) gs.Option[U]) gs.Option[U] {
	return gs.Partial(fn, None[U])(opt.Fetch)
}

func Map[T, U any](opt gs.Option[T], fn func(T) U) gs.Option[U] {
	return gs.Partial(
		gs.FuncAndThen(fn, Some[U]),
		None[U],
	)(opt.Fetch)
}

func MapBool[T, U any](opt gs.Option[T], fn func(T) (U, bool)) gs.Option[U] {
	return gs.Partial(
		gs.FuncBoolAndThen(fn, Bool[U]),
		None[U],
	)(opt.Fetch)
}

func MapErr[T, U any](opt gs.Option[T], fn func(T) (U, error)) gs.Option[U] {
	return gs.Partial(
		gs.FuncErrAndThen(fn, Err[U]),
		None[U],
	)(opt.Fetch)
}

func Fold[T, U any](opt gs.Option[T]) func(U) func(func(T) U) U {
	return func(z U) func(func(T) U) U {
		return func(fn func(T) U) U {
			return gs.Partial(fn, gs.VF(z))(opt.Fetch)
		}
	}
}
