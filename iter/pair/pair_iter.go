// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pair

type Iter[K comparable, V any] interface {
	Next() bool
	Get() (K, V)
}
