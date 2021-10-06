// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package goscala

import "fmt"

type Sequence[T any] interface {
	fmt.Stringer
	Get(int) T

	IsEmpty() bool
	IsNotEmpty() bool

	Len() int
	Cap() int
	Append(...T) Sequence[T]
	Head() Option[T]
	Last() Option[T]

	//Forall(func(T) bool) bool
	//Foreach(func(T))
	//

	//Tail() Sequence[T]
	//
	//Equals(func(T, T) bool) func(Sequence[T]) bool
	//Contains(func(T, T) bool) func(T) bool
	//Exists(func(T) bool) bool

}
