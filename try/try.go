// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package try

import (
	gs "github.com/kigichang/goscala"
	"github.com/kigichang/goscala/impl"
)

func Success[T any](v T) gs.Try[T] {
	return impl.Success[T](v)
}

func Failure[T any](err error) gs.Try[T] {
	return impl.Failure[T](err)
}
