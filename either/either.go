// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package either

import (
	gs "github.com/kigichang/goscala"
	"github.com/kigichang/goscala/impl"
)

func Left[L, R any](v L) gs.Either[L, R] {
	return &impl.Either[L, R]{
		OK: false,
		L:  v,
	}
}

func Right[L, R any](v R) gs.Either[L, R] {
	return &impl.Either[L, R]{
		OK: true,
		R:  v,
	}
}
