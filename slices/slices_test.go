// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package slices_test

import (
	"strconv"
	"testing"

	gs "github.com/kigichang/goscala"
	"github.com/kigichang/goscala/slices"
	"github.com/stretchr/testify/assert"
)

func TestFill(t *testing.T) {
	s := slices.Fill(5, -1)
	seq := s.Equals(gs.Eq[int])
	assert.True(t, seq(slices.From(-1, -1, -1, -1, -1)))
}

func TestRange(t *testing.T) {
	assert.True(t, slices.Range(0, 10, 2).Equals(gs.Eq[int])(slices.From(0, 2, 4, 6, 8)))
}

func TestTabulate(t *testing.T) {
	fn := func(v int) string {
		return strconv.Itoa(v + 1)
	}

	assert.True(t, slices.Tabulate(5, fn).Equals(gs.Eq[string])(slices.From("1", "2", "3", "4", "5")))
}

func TestCollect(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	fn := func(v int) (s string, ok bool) {
		if ok = ((v & 0x01) == 0); ok {
			s = strconv.Itoa(v)
		}
		return
	}

	dst := slices.Collect(src, fn)
	assert.True(t, dst.Equals(gs.Eq[string])(slices.From("2", "4", "6", "8")))
}

func TestCollectFirst(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	fn := func(v int) (s string, ok bool) {
		if ok = ((v & 0x01) == 0); ok {
			s = strconv.Itoa(v)
		}
		return
	}

	dst := slices.CollectFirst(src, fn)
	assert.Equal(t, "2", dst.Get())

	fn = func(_ int) (string, bool) {
		return "", false
	}

	assert.False(t, slices.CollectFirst(src, fn).IsDefined())
}

func TestFlatMap(t *testing.T) {
	dst := slices.FlatMap(
		slices.From(1, 2, 3),
		func(v int) gs.Sliceable[int] {
			return slices.Map(
				slices.From(4, 5, 6),
				func(x int) int {
					return v * x
				})
		})

	ans := slices.From(4, 5, 6, 8, 10, 12, 12, 15, 18)
	assert.True(t, dst.Equals(gs.Eq[int])(ans))
}

func TestFoldLeft(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	fn1 := func(v1, v2 int) int {
		return v1 - v2
	}

	fn2 := func(v1, v2 int) int {
		return v1 + v2
	}

	assert.Equal(t, -45, slices.FoldLeft(src, 0, fn1))
	assert.Equal(t, -45, slices.Fold(src, 0, fn1))

	assert.Equal(t, 45, slices.FoldLeft(src, 0, fn2))
	assert.Equal(t, 45, slices.Fold(src, 0, fn2))
}

func TestFoldRight(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	fn1 := func(v1, v2 int) int {
		return v1 - v2
	}

	fn2 := func(v1, v2 int) int {
		return v1 + v2
	}

	assert.Equal(t, 9, slices.FoldRight(src, 0, fn1))
	assert.Equal(t, 45, slices.FoldRight(src, 0, fn2))
}

func TestPartitionMap(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	fn := func(v int) gs.Either[int, int] {
		if (v & 0x01) == 0 {
			return gs.Right[int, int](v)
		}
		return gs.Left[int, int](10 + v)
	}

	a, b := slices.PartitionMap(src, fn)
	assert.True(t, slices.From(11, 13, 15, 17, 19).Equals(gs.Eq[int])(a))
	assert.True(t, slices.From(2, 4, 6, 8).Equals(gs.Eq[int])(b))
}

func TestScanLeft(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	fn := func(v1, v2 int) int {
		return v1 + v2
	}

	dst := slices.ScanLeft(src, 100, fn)
	ans := slices.From(100, 101, 104, 109, 116, 125, 127, 131, 137, 145)
	assert.True(t, dst.Equals(gs.Eq[int])(ans))

	dst = slices.Scan(src, 100, fn)
	assert.True(t, dst.Equals(gs.Eq[int])(ans))
}

func TestScanRight(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	fn := func(v1, v2 int) int {
		return v1 + v2
	}
	dst := slices.ScanRight(src, 100, fn)
	ans := slices.From(145, 144, 141, 136, 129, 120, 118, 114, 108, 100)
	assert.True(t, dst.Equals(gs.Eq[int])(ans))
}

func TestGroupBy(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	fn := func(v int) bool {
		return (v & 0x01) == 0
	}
	m := slices.GroupBy(src, fn)

	assert.Equal(t, m[true], slices.From(2, 4, 6, 8))
	assert.Equal(t, m[false], slices.From(1, 3, 5, 7, 9))
}

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
	o := slices.Empty[int]().Head()
	assert.False(t, o.IsDefined())

	o = src.Head()
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

	o := src.Find(p1)
	assert.Equal(t, 7, o.Get())

	o = src.Find(p2)
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

	o := src.FindLast(p1)
	assert.Equal(t, 8, o.Get())

	o = src.FindLast(p2)
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

	o := src.ReduceLeft(fn1)
	assert.Equal(t, -43, o.Get())

	o = src.Reduce(fn1)
	assert.Equal(t, -43, o.Get())

	o = src.ReduceLeft(fn2)
	assert.Equal(t, 45, o.Get())

	o = src.Reduce(fn2)
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

	o := src.ReduceRight(fn1)
	assert.Equal(t, 9, o.Get())

	o = src.ReduceRight(fn2)
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

	assert.Equal(t, 9, src.Max(gs.Compare[int]).Get())
}

func TestSliceMin(t *testing.T) {
	src := slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)
	assert.Equal(t, 1, src.Min(gs.Compare[int]).Get())
}

func TestSliceSort(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)
	assert.True(t, src.Sort(gs.Compare[int]).Equals(gs.Eq[int])(slices.From(1, 2, 3, 4, 5, 6, 7, 8, 9)))
}
