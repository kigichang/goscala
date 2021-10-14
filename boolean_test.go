// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package goscala_test

import (
	"testing"

	"github.com/kigichang/goscala"
	"github.com/stretchr/testify/assert"
)

func TestTrue(t *testing.T) {
	assert.True(t, goscala.True())
}

func TestFalse(t *testing.T) {
	assert.False(t, goscala.False())
}

func TestCond(t *testing.T) {
	assert.Equal(t, 1, goscala.Cond(true, 1, -1))
	assert.Equal(t, -1, goscala.Cond(false, 1, -1))
}

func TestDefault(t *testing.T) {
	assert.Equal(t, 100, goscala.Default(1)(func() (int, bool) {
		return 100, true
	}))

	assert.Equal(t, 1, goscala.Default(1)(func() (int, bool) {
		return 100, false
	}))
}

func TestTernary(t *testing.T) {
	assert.Equal(t, 1, goscala.Ternary(
		goscala.True,
		func() int {
			return 1
		},
		func() int {
			return -1
		},
	))

	assert.Equal(t, -1, goscala.Ternary(
		goscala.False,
		func() int {
			return 1
		},
		func() int {
			return -1
		},
	))
}
