// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package opt

import (
	"fmt"

	gs "github.com/kigichang/goscala"
)

type option[T any] struct {
	defined bool
	v       T
}

var _ gs.Option[int] = &option[int]{}

func (opt *option[T]) String() string {
	if opt.defined {
		return fmt.Sprintf(`Some(%v)`, opt.v)
	}

	return fmt.Sprintf(`None[%s]`, gs.TypeStr(opt.v))
}

func (opt *option[T]) Fetch() (T, bool) {
	return opt.v, opt.defined
}

func (opt *option[T]) FetchErr() (T, error) {
	return opt.v, gs.Cond(opt.defined, nil, gs.ErrEmpty)
}

func (opt *option[T]) IsDefined() bool {
	return opt.defined
}

func (opt *option[T]) IsEmpty() bool {
	return !opt.defined
}

func (opt *option[T]) Get() T {
	if opt.defined {
		return opt.v
	}
	panic(fmt.Sprintf(`can not get value from %v`, opt))
}

func (opt *option[T]) GetOrElse(z T) T {
	return gs.Default(z)(opt.Fetch)
}

func (opt *option[T]) OrElse(z gs.Option[T]) gs.Option[T] {
	return gs.Cond(opt.defined, gs.Option[T](opt), z)
}

func (opt *option[T]) Contains(v T, eq func(T, T) bool) bool {
	return gs.Partial(
		gs.Currying2(eq)(v),
		gs.False,
	)(opt.Fetch)

}

func (opt *option[T]) Exists(p func(T) bool) bool {
	return opt.Filter(p).IsDefined()
}

func (opt *option[T]) Equals(eq func(T, T) bool) func(gs.Option[T]) bool {
	return func(that gs.Option[T]) bool {
		return gs.Partial(
			func(x T) bool {
				return that.IsDefined() && eq(that.Get(), x)
			},
			that.IsEmpty,
		)(opt.Fetch)
	}
}

func (opt *option[T]) Filter(p func(T) bool) gs.Option[T] {
	return gs.Partial(
		gs.Predict(Some[T], None[T])(p),
		None[T],
	)(opt.Fetch)
}

func (opt *option[T]) FilterNot(p func(T) bool) gs.Option[T] {
	return opt.Filter(func(v T) bool {
		return !p(v)
	})
}

func (opt *option[T]) Forall(p func(T) bool) bool {
	return gs.Partial(
		p,
		gs.True,
	)(opt.Fetch)
}

func (opt *option[T]) Foreach(f func(T)) {
	gs.Partial(gs.UnitWrap(f), gs.Unit)(opt.Fetch)
}

func Some[T any](v T) gs.Option[T] {
	return &option[T]{
		defined: true,
		v:       v,
	}
}

func None[T any]() gs.Option[T] {
	return &option[T]{
		defined: false,
	}
}

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
