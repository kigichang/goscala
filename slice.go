// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package goscala

type Sliceable[T any] interface {
	Slice() []T
}

func sliceOne[T any](elem T) []T {
	return []T{elem}
}

func sliceEmpty[T any]() []T {
	return []T{}
}
