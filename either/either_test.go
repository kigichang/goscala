// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package either_test

import (
	"strconv"
	"testing"

	"github.com/kigichang/goscala"
	"github.com/kigichang/goscala/either"
	"github.com/stretchr/testify/assert"
)

func TestCond(t *testing.T) {
	lv := 100
	rv := "abc"

	e := either.Cond(func() bool { return true }, lv, rv)
	assert.True(t, e.IsRight())
	assert.False(t, e.IsLeft())
	assert.Equal(t, rv, e.Right())

	e = either.Cond(func() bool { return false }, lv, rv)
	assert.False(t, e.IsRight())
	assert.True(t, e.IsLeft())
	assert.Equal(t, lv, e.Left())
}

func TestFlatMap(t *testing.T) {
	f := func(v string) goscala.Either[int, int64] {
		a, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return goscala.Left[int, int64](0)
		}
		return goscala.Right[int, int64](a)
	}

	r := goscala.Right[int, string]("1000")

	e := either.FlatMap(r, f)

	assert.Equal(t, true, e.IsRight())
	assert.Equal(t, int64(1000), e.Right())

	r = goscala.Right[int, string]("abc")
	e = either.FlatMap(r, f)
	assert.Equal(t, false, e.IsRight())
	assert.Equal(t, true, e.IsLeft())
	assert.Equal(t, 0, e.Left())

	l := goscala.Left[int, string](100)
	e = either.FlatMap(l, f)

	assert.Equal(t, true, e.IsLeft())
	assert.Equal(t, l.Left(), e.Left())
}

func TestFold(t *testing.T) {
	fa := func(v int) int64 {
		return int64(v + 10)
	}

	fb := func(s string) int64 {
		a, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return -1
		}
		return a * 10
	}

	r := goscala.Right[int, string]("1000")
	e := either.Fold(r, fa, fb)
	assert.Equal(t, int64(1000*10), e)

	r = goscala.Right[int, string]("abc")
	e = either.Fold(r, fa, fb)
	assert.Equal(t, int64(-1), e)

	l := goscala.Left[int, string](100)
	e = either.Fold(l, fa, fb)
	assert.Equal(t, int64(100+10), e)
}

func TestMap(t *testing.T) {
	/*
	   Right(12).map(x => "flower") // Result: Right("flower")
	   Left(12).map(x => "flower")  // Result: Left(12)
	*/
	f := func(_ int) string {
		return "flower"
	}

	e := either.Map(goscala.Right[int, int](12), f)
	assert.Equal(t, true, e.IsRight())
	assert.Equal(t, "flower", e.Right())

	e = either.Map(goscala.Left[int, int](12), f)
	assert.Equal(t, true, e.IsLeft())
	assert.Equal(t, 12, e.Left())
}
