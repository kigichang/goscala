// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package goscala

import "fmt"

var (
	ErrUnsupported = fmt.Errorf("unsupported")
	ErrUnsatisfied = fmt.Errorf("unsatisfied")
	ErrEmpty = fmt.Errorf("emtpy")
	ErrLeft = fmt.Errorf("left")
)