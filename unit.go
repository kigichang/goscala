// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package goscala

type UnitRef = struct{}

var unit = UnitRef{}

func Unit() UnitRef {
	return unit
}

func UnitWrap[T any](fn func(T)) func(T) UnitRef {
	return func(v T) UnitRef {
		fn(v)
		return unit
	}
}
