// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package goscala

import (
	"fmt"
)

type Try[T any] interface {
	fmt.Stringer
	Fetcher[T]
	IsSuccess() bool
	Success() T
	Failed() error
	IsFailure() bool
	Equals(func(T, T) bool) func(Try[T]) bool
	Get() T
	Filter(func(T) bool) Try[T]
	Foreach(func(T))
	GetOrElse(z T) T
	OrElse(z Try[T]) Try[T]
	Recover(func(error) (T, bool)) Try[T]
	RecoverWith(func(error) (Try[T], bool)) Try[T]
}
