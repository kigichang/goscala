// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package goscala

import "fmt"

type Seq[T any] interface {
	fmt.Stringer
	Iterable[T]
	Len() int
	Get(int) T
}
