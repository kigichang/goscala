// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package goscala_test

import (
	"testing"

	gs "github.com/kigichang/goscala"
	"github.com/kigichang/goscala/opt"
	"github.com/kigichang/goscala/slices"
	"github.com/stretchr/testify/assert"
)

func TestSliceClone(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	dst := src.Clone()

	assert.False(t, &src == &dst)
	assert.True(t, src.Equals(gs.Eq[int])(dst))
}

func TestSliceForAll(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	p1 := func(v int) bool {
		return v >= 0
	}

	p2 := func(v int) bool {
		return v > 5
	}

	assert.True(t, slices.Empty[int]().Forall(p2))
	assert.True(t, src.Forall(p1))
	assert.False(t, src.Forall(p2))
}

func TestSliceForeach(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	sum := 0
	src.Foreach(func(v int) {
		sum += v
	})
	assert.Equal(t, 45, sum)
}

func TestSliceHead(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)
	o := opt.Bool[int](slices.Empty[int]().Head())
	assert.False(t, o.IsDefined())

	o = opt.Bool[int](src.Head())
	assert.Equal(t, 1, o.Get())
}

func TestSliceTail(t *testing.T) {
	src := slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	assert.True(t, slices.Empty[int]().Tail().IsEmpty())
	assert.Equal(t, src.Tail(), slices.From(3, 5, 7, 9, 2, 4, 6, 8))
}

func TestSliceEquals(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)
	eq := src.Equals(gs.Eq[int])

	assert.False(t, eq(slices.From(1, 2, 3)))
	assert.True(t, eq(src))
	assert.True(t, eq(slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)))
}

func TestSliceContains(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)
	contain := src.Contains(gs.Eq[int])
	assert.False(t, slices.Empty[int]().Contains(gs.Eq[int])(1))
	assert.False(t, contain(-1))
	assert.True(t, contain(5))
}

func TestSliceExists(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	p1 := func(v int) bool {
		return v > 0
	}
	p2 := func(v int) bool {
		return v < 0
	}

	assert.False(t, slices.Empty[int]().Exists(p1))
	assert.False(t, slices.Empty[int]().Exists(p2))

	assert.True(t, src.Exists(p1))
	assert.False(t, src.Exists(p2))

}

func TestSliceFilter(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	p := func(v int) bool {
		return (v & 0x01) == 0
	}

	s := src.Filter(p)

	assert.True(t, s.Equals(gs.Eq[int])(slices.From(2, 4, 6, 8)))
}

func TestSliceFilterNot(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	p := func(v int) bool {
		return (v & 0x01) == 0
	}

	s := src.FilterNot(p)

	assert.True(t, s.Equals(gs.Eq[int])(slices.From(1, 3, 5, 7, 9)))
}

func TestSliceFind(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	p1 := func(v int) bool {
		return v > 5
	}
	p2 := func(v int) bool {
		return v < 0
	}

	o := opt.Bool[int](src.Find(p1))
	assert.Equal(t, 7, o.Get())

	o = opt.Bool[int](src.Find(p2))
	assert.False(t, o.IsDefined())
}

func TestSliceFindLast(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	p1 := func(v int) bool {
		return v > 5
	}
	p2 := func(v int) bool {
		return v < 0
	}

	o := opt.Bool[int](src.FindLast(p1))
	assert.Equal(t, 8, o.Get())

	o = opt.Bool[int](src.FindLast(p2))
	assert.False(t, o.IsDefined())
}

func TestSlicePartition(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	p := func(v int) bool {
		return (v & 0x01) == 0
	}

	a, b := src.Partition(p)

	assert.True(t, a.Equals(gs.Eq[int])(slices.From(1, 3, 5, 7, 9)))
	assert.True(t, b.Equals(gs.Eq[int])(slices.From(2, 4, 6, 8)))
}

func TestSliceReverse(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	s := src.Reverse()
	assert.True(t, s.Equals(gs.Eq[int])(slices.From(8, 6, 4, 2, 9, 7, 5, 3, 1)))
}

func TestSplitAt(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	a, b := src.SplitAt(100)

	assert.True(t, src.Equals(gs.Eq[int])(a))
	assert.Equal(t, 0, b.Len())

	a, b = src.SplitAt(-100)

	assert.True(t, src.Equals(gs.Eq[int])(b))
	assert.Equal(t, 0, a.Len())

	a, b = src.SplitAt(5)

	assert.True(t, a.Equals(gs.Eq[int])(src[:5]))
	assert.True(t, b.Equals(gs.Eq[int])(src[5:]))

	a, b = src.SplitAt(-3)
	idx := src.Len() - 3
	assert.True(t, a.Equals(gs.Eq[int])(src[:idx]))
	assert.True(t, b.Equals(gs.Eq[int])(src[idx:]))
}

func TestSliceTake(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	a := src.Take(100)
	assert.True(t, src.Equals(gs.Eq[int])(a))

	a = src.Take(-100)
	assert.True(t, src.Equals(gs.Eq[int])(a))

	a = src.Take(5)
	assert.True(t, a.Equals(gs.Eq[int])(src[:5]))

	a = src.Take(-3)
	idx := src.Len() - 3
	assert.True(t, a.Equals(gs.Eq[int])(src[idx:]))
}
func TestSliceTakeWhile(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	p := func(v int) bool {
		return (v & 0x01) == 1
	}

	a := src.TakeWhile(p)
	assert.True(t, a.Equals(gs.Eq[int])(slices.From(1, 3, 5, 7, 9)))
}

func TestSliceCount(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	p := func(v int) bool {
		return (v & 0x01) == 1
	}
	assert.Equal(t, 5, src.Count(p))
}

func TestSliceDrop(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	b := src.Drop(100)
	assert.Equal(t, 0, b.Len())

	b = src.Drop(-100)
	assert.Equal(t, 0, b.Len())

	b = src.Drop(5)
	assert.True(t, b.Equals(gs.Eq[int])(src[5:]))

	b = src.Drop(-3)
	idx := src.Len() - 3
	assert.True(t, b.Equals(gs.Eq[int])(src[:idx]))
}

func TestSliceDropWhile(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	p := func(v int) bool {
		return (v & 0x01) == 1
	}

	a := src.DropWhile(p)
	assert.True(t, a.Equals(gs.Eq[int])(slices.From(2, 4, 6, 8)))
}

func TestSliceReduceLeft(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	fn1 := func(v1, v2 int) int {
		return v1 - v2
	}

	fn2 := func(v1, v2 int) int {
		return v1 + v2
	}

	o := opt.Bool[int](src.ReduceLeft(fn1))
	assert.Equal(t, -43, o.Get())

	o = opt.Bool[int](src.Reduce(fn1))
	assert.Equal(t, -43, o.Get())

	o = opt.Bool[int](src.ReduceLeft(fn2))
	assert.Equal(t, 45, o.Get())

	o = opt.Bool[int](src.Reduce(fn2))
	assert.Equal(t, 45, o.Get())
}

func TestSliceReduceRight(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	fn1 := func(v1, v2 int) int {
		return v1 - v2
	}

	fn2 := func(v1, v2 int) int {
		return v1 + v2
	}

	o := opt.Bool[int](src.ReduceRight(fn1))
	assert.Equal(t, 9, o.Get())

	o = opt.Bool[int](src.ReduceRight(fn2))
	assert.Equal(t, 45, o.Get())
}

func TestSliceIndexWhereFrom(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	p1 := func(v int) bool {
		return (v & 0x01) == 0
	}

	p2 := func(v int) bool {
		return (v & 0x01) == 1
	}

	assert.Equal(t, 5, src.IndexWhereFrom(p1, 0))
	assert.Equal(t, 5, src.IndexWhere(p1))
	assert.Equal(t, -1, src.IndexWhereFrom(p2, 6))
	assert.Equal(t, 0, src.IndexWhere(p2))
}

func TestSliceLastIndexWhereFrom(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	p1 := func(v int) bool {
		return (v & 0x01) == 0
	}

	p2 := func(v int) bool {
		return (v & 0x01) == 1
	}

	assert.Equal(t, -1, src.LastIndexWhereFrom(p1, 4))
	assert.Equal(t, 8, src.LastIndexWhere(p1))
	assert.Equal(t, 4, src.LastIndexWhereFrom(p2, 6))
	assert.Equal(t, 4, src.LastIndexWhere(p2))
}

func TestSliceMax(t *testing.T) {
	src := slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	assert.Equal(t, 9, opt.Bool[int](src.Max(gs.Compare[int])).Get())
}

func TestSliceMin(t *testing.T) {
	src := slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)
	assert.Equal(t, 1, opt.Bool[int](src.Min(gs.Compare[int])).Get())
}

func TestSliceSort(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)
	assert.True(t, src.Sort(gs.Compare[int]).Equals(gs.Eq[int])(slices.From(1, 2, 3, 4, 5, 6, 7, 8, 9)))
}
