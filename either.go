// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package goscala

import (
	"fmt"
)

type Either[L, R any] interface {
	fmt.Stringer
	Fetcher[R]
	IsRight() bool
	IsLeft() bool
	Get() R
	Left() L
	Right() R

	Exists(func(R) bool) bool
	FilterOrElse(func(R) bool, L) Either[L, R]
	Forall(func(R) bool) bool
	Foreach(func(R))
	GetOrElse(R) R
	OrElse(Either[L, R]) Either[L, R]
	Swap() Either[R, L]
}
