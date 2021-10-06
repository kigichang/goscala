// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package goscala

import "fmt"

type Map[K comparable, V any] interface {
	fmt.Stringer
	PairIterable[K, V]
	Get(K) (V, bool)
	Put(K, V)
	Delete(K)
}
