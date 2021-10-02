// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package iter_test

import (
	"testing"

	"github.com/kigichang/goscala/iter"
	"github.com/stretchr/testify/assert"
)

func TestGenIter(t *testing.T) {

	it := iter.Gen(1, 2, 3, 4)

	ss := iter.Slice(it)

	assert.Equal(t, ss, []int{1, 2, 3, 4})
}

func TestFoldLeft(t *testing.T) {
	src := []int{1, 3, 5, 7, 9, 2, 4, 6, 8}

	it := iter.Gen(src...)
	fn1 := func(v1, v2 int) int {
		return v1 - v2
	}
	assert.Equal(t, -45, iter.FoldLeft(it, 0, fn1))

	it = iter.Gen(src...)
	fn2 := func(v1, v2 int) int {
		return v1 + v2
	}
	assert.Equal(t, 45, iter.FoldLeft(it, 0, fn2))
}

func TestFoldRight(t *testing.T) {
	src := []int{1, 3, 5, 7, 9, 2, 4, 6, 8}

	it := iter.Gen(src...)
	fn1 := func(v1, v2 int) int {
		return v1 - v2
	}
	assert.Equal(t, 9, iter.FoldRight(it, 0, fn1))

	it = iter.Gen(src...)
	fn2 := func(v1, v2 int) int {
		return v1 + v2
	}
	assert.Equal(t, 45, iter.FoldRight(it, 0, fn2))
}

func TestScanLeft(t *testing.T) {
	src := []int{1, 3, 5, 7, 9, 2, 4, 6, 8}

	fn := func(v1, v2 int) int {
		return v1 + v2
	}

	dst := iter.Slice(iter.ScanLeft(iter.Gen(src...), 100, fn))
	ans := []int{100, 101, 104, 109, 116, 125, 127, 131, 137, 145}
	assert.Equal(t, ans, dst)
}

func TestScanRight(t *testing.T) {
	src := []int{1, 3, 5, 7, 9, 2, 4, 6, 8}
	fn := func(v1, v2 int) int {
		return v1 + v2
	}

	dst := iter.Slice(iter.ScanRight(iter.Gen(src...), 100, fn))
	ans := []int{145, 144, 141, 136, 129, 120, 118, 114, 108, 100}
	assert.Equal(t, ans, dst)
}

func TestFlatMapAndMap(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}

	dst := iter.FlatMap(
		iter.Gen(s1...),
		func(x int) iter.Iter[int] {
			return iter.Map(iter.Gen(s2...), func(y int) int {
				return x * y
			})
		},
	)

	ans := []int{4, 5, 6, 8, 10, 12, 12, 15, 18}
	assert.Equal(t, ans, iter.Slice(dst))
}
