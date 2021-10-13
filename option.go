// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package goscala

import (
	"fmt"
)

type Option[T any] interface {
	fmt.Stringer
	Fetcher[T]
	IsDefined() bool
	IsEmpty() bool

	Contains(T, func(T, T) bool) bool
	Exists(func(T) bool) bool
	Equals(func(T, T) bool) func(Option[T]) bool
	Filter(func(T) bool) Option[T]
	FilterNot(func(T) bool) Option[T]
	Forall(p func(T) bool) bool
	Foreach(f func(T))
	Get() T
	GetOrElse(z T) T
	OrElse(Option[T]) Option[T]
}
