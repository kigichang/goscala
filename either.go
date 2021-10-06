// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package goscala

import "fmt"

type Either[L, R any] interface {
	fmt.Stringer
	Fetcher[R]

	IsLeft() bool
	IsRight() bool

	Left() L
	Right() R
	Get() R

	//Try() Try[R]
}
