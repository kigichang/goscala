// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package goscala

import "fmt"

type Option[T any] interface {
	fmt.Stringer
	Fetcher[T]

	IsDefined() bool
	IsEmpty() bool

	Get() T
	GetOrElse(T) T
	//OrElse(Option[T]) Option[T]
	//
	//Equals(func(T, T) bool) func(Option[T]) bool
	//Contains(func(T, T) bool) func(T) bool
	//Exists(func(T) bool) bool
	//
	//Filter(func(T) bool) Option[T]
	//FilterNot(func(T) bool) Option[T]
	//
	//Forall(p func(T) bool) bool
	//Foreach(f func(T))
}
