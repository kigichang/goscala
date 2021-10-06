// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package goscala

import "fmt"

type Try[T any] interface {
	fmt.Stringer
	IsSuccess() bool
	IsFailure() bool

	Success() T
	Failed() error

	Get() T

	//Either() Either[error, T]
}
