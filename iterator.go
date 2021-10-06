// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package goscala

type Iterator[T any] interface {
	Len() int
	Cap() int
	Next() bool
	Get() T
}

type Iterable[T any] interface {
	Range() Iterator[T]
}

type PairIterator[K comparable, V any] interface {
	Len() int
	Next() bool
	Get() (K, V)
}

type PairIterable[K comparable, V any] interface {
	Range() PairIterator[K, V]
}
