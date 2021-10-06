// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package goscala

import (
	"constraints"
	"fmt"
)

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

var (
	ErrUnsupported = fmt.Errorf("unsupported")
	ErrUnsatisfied = fmt.Errorf("unsatisfied")
	ErrEmpty       = fmt.Errorf("emtpy")
	ErrLeft        = fmt.Errorf("left")
)

type Fetcher[T any] interface {
	Fetch() (T, bool)
	FetchErr() (T, error)
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Compare[T constraints.Ordered](a, b T) int {
	if a == b {
		return 0
	}

	if a > b {
		return 1
	}
	return -1
}

func Equal[T comparable](a, b T) bool {
	return Eq[T](a, b)
}

func Eq[T comparable](a, b T) bool {
	return a == b
}
