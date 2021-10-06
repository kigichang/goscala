// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package goscala

type Pair[K comparable, V any] interface {
	Key() K
	Value() V
	Get() (K, V)
}
