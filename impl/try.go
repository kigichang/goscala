// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package impl

import (
	"fmt"

	gs "github.com/kigichang/goscala"
)

type try[T any] struct {
	v   T
	err error
}

var _ gs.Try[int] = &try[int]{}

func (t *try[T]) String() string {
	if t.IsSuccess() {
		return fmt.Sprintf(`Success(%v)`, t.v)
	}
	return fmt.Sprintf(`Failure(%v)`, t.err)
}

func (t *try[T]) IsSuccess() bool {
	return t.err == nil
}

func (t *try[T]) Success() T {
	if t.IsSuccess() {
		return t.v
	}
	panic(fmt.Errorf(`can not get success value from %v`, t))
}

func (t *try[T]) Get() T {
	return t.Success()
}

func (t *try[T]) IsFailure() bool {
	return t.err != nil
}

func (t *try[T]) Failed() error {
	if t.IsFailure() {
		return t.err
	}
	return gs.ErrUnsupported
}

//func (t *try[T]) Either() gs.Either[error, T] {
//	if t.IsSuccess() {
//		return Right[error, T](t.Success())
//	}
//	return Left[error, T](t.Failed())
//}

func Success[T any](v T) *try[T] {
	return &try[T]{
		v: v,
	}
}

func Failure[T any](err error) *try[T] {
	return &try[T]{
		err: err,
	}
}
