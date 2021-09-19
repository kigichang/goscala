package goscala

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceClone(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	dst := src.Clone()

	assert.False(t, &src == &dst)
	assert.True(t, src.Equals(dst, Equal[int]))
}

func TestSliceForAll(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	p1 := func(v int) bool {
		return v >= 0
	}

	p2 := func(v int) bool {
		return v > 5
	}

	assert.True(t, SliceEmpty[int]().Forall(p2))
	assert.True(t, src.Forall(p1))
	assert.False(t, src.Forall(p2))
}

func TestSliceForeach(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	sum := 0
	src.Foreach(func(v int) {
		sum += v
	})
	assert.Equal(t, 45, sum)
}

func TestSliceHead(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	assert.False(t, SliceEmpty[int]().Head().IsDefined())
	assert.Equal(t, 1, src.Head().Get())
}

func TestSliceTail(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	assert.False(t, SliceEmpty[int]().Last().IsDefined())
	assert.Equal(t, 8, src.Last().Get())
}

func TestSliceEquals(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	assert.False(t, src.Equals(SliceFrom(1, 2, 3), Equal[int]))
	assert.True(t, src.Equals(src, Equal[int]))
	assert.True(t, src.Equals(SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8), Equal[int]))
}

func TestSliceContains(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)
	assert.False(t, SliceEmpty[int]().Contains(1, Equal[int]))
	assert.False(t, src.Contains(-1, Equal[int]))
	assert.True(t, src.Contains(5, Equal[int]))
}

func TestSliceExists(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	p1 := func(v int) bool {
		return v > 0
	}
	p2 := func(v int) bool {
		return v < 0
	}

	assert.False(t, SliceEmpty[int]().Exists(p1))
	assert.False(t, SliceEmpty[int]().Exists(p2))

	assert.True(t, src.Exists(p1))
	assert.False(t, src.Exists(p2))

}

func TestSliceFilter(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	p := func(v int) bool {
		return (v & 0x01) == 0
	}

	s := src.Filter(p)

	assert.True(t, s.Equals(SliceFrom(2, 4, 6, 8), Equal[int]))
}

func TestSliceFilterNot(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	p := func(v int) bool {
		return (v & 0x01) == 0
	}

	s := src.FilterNot(p)

	assert.True(t, s.Equals(SliceFrom(1, 3, 5, 7, 9), Equal[int]))
}

func TestSliceFind(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	p1 := func(v int) bool {
		return v > 5
	}
	p2 := func(v int) bool {
		return v < 0
	}
	assert.Equal(t, 7, src.Find(p1).Get())
	assert.False(t, src.Find(p2).IsDefined())
}

func TestSliceFindLast(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	p1 := func(v int) bool {
		return v > 5
	}
	p2 := func(v int) bool {
		return v < 0
	}
	assert.Equal(t, 8, src.FindLast(p1).Get())
	assert.False(t, src.FindLast(p2).IsDefined())
}

func TestSlicePartition(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	p := func(v int) bool {
		return (v & 0x01) == 0
	}

	a, b := src.Partition(p)

	assert.True(t, a.Equals(SliceFrom(1, 3, 5, 7, 9), Equal[int]))
	assert.True(t, b.Equals(SliceFrom(2, 4, 6, 8), Equal[int]))
}

func TestSliceReverse(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	s := src.Reverse()
	assert.True(t, s.Equals(SliceFrom(8, 6, 4, 2, 9, 7, 5, 3, 1), Equal[int]))
}

func TestSplitAt(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	a, b := src.SplitAt(100)

	assert.True(t, src.Equals(a, Equal[int]))
	assert.Equal(t, 0, b.Len())

	a, b = src.SplitAt(-100)

	assert.True(t, src.Equals(b, Equal[int]))
	assert.Equal(t, 0, a.Len())

	a, b = src.SplitAt(5)

	assert.True(t, a.Equals(src[:5], Equal[int]))
	assert.True(t, b.Equals(src[5:], Equal[int]))

	a, b = src.SplitAt(-3)
	idx := src.Len() - 3
	assert.True(t, a.Equals(src[:idx], Equal[int]))
	assert.True(t, b.Equals(src[idx:], Equal[int]))
}

func TestSliceTake(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	a := src.Take(100)
	assert.True(t, src.Equals(a, Equal[int]))

	a = src.Take(-100)
	assert.True(t, src.Equals(a, Equal[int]))

	a = src.Take(5)
	assert.True(t, a.Equals(src[:5], Equal[int]))

	a = src.Take(-3)
	idx := src.Len() - 3
	assert.True(t, a.Equals(src[idx:], Equal[int]))
}

func TestSliceTakeWhile(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	p := func(v int) bool {
		return (v & 0x01) == 1
	}

	a := src.TakeWhile(p)
	assert.True(t, a.Equals(SliceFrom(1, 3, 5, 7, 9), Equal[int]))
}

func TestSliceCount(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	p := func(v int) bool {
		return (v & 0x01) == 1
	}
	assert.Equal(t, 5, src.Count(p))
}

func TestSliceDrop(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	b := src.Drop(100)
	assert.Equal(t, 0, b.Len())

	b = src.Drop(-100)
	assert.Equal(t, 0, b.Len())

	b = src.Drop(5)
	assert.True(t, b.Equals(src[5:], Equal[int]))

	b = src.Drop(-3)
	idx := src.Len() - 3
	assert.True(t, b.Equals(src[:idx], Equal[int]))
}

func TestSliceDropWhile(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	p := func(v int) bool {
		return (v & 0x01) == 1
	}

	a := src.DropWhile(p)
	assert.True(t, a.Equals(SliceFrom(2, 4, 6, 8), Equal[int]))
}

func TestSliceReduceLeft(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	fn1 := func(v1, v2 int) int {
		return v1 - v2
	}

	fn2 := func(v1, v2 int) int {
		return v1 + v2
	}

	assert.Equal(t, -43, src.ReduceLeft(fn1).Get())
	assert.Equal(t, -43, src.Reduce(fn1).Get())

	assert.Equal(t, 45, src.ReduceLeft(fn2).Get())
	assert.Equal(t, 45, src.Reduce(fn2).Get())
}

func TestSliceReduceRight(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	fn1 := func(v1, v2 int) int {
		return v1 - v2
	}

	fn2 := func(v1, v2 int) int {
		return v1 + v2
	}

	assert.Equal(t, 9, src.ReduceRight(fn1).Get())
	assert.Equal(t, 45, src.ReduceRight(fn2).Get())
}

func TestSliceIndexWhereFrom(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

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
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

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
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)
	assert.Equal(t, 9, src.Max(Compare[int]).Get())
}

func TestSliceMin(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)
	assert.Equal(t, 1, src.Min(Compare[int]).Get())
}

func TestSliceSort(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)
	assert.True(t, src.Sort(Compare[int]).Equals(SliceFrom(1, 2, 3, 4, 5, 6, 7, 8, 9), Equal[int]))
}

func TestSliceFill(t *testing.T) {
	assert.True(t, SliceFill(5, -1).Equals(SliceFrom(-1, -1, -1, -1, -1), Equal[int]))
}

func TestSliceRange(t *testing.T) {
	assert.True(t, SliceRange(0, 10, 2).Equals(SliceFrom(0, 2, 4, 6, 8), Equal[int]))
}

func TestSliceTabulate(t *testing.T) {
	fn := func(v int) string {
		return strconv.Itoa(v + 1)
	}

	assert.True(t, SliceTabulate(5, fn).Equals(SliceFrom("1", "2", "3", "4", "5"), Equal[string]))
}

func TestSliceCollect(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	fn := func(v int) (s string, ok bool) {
		if ok = ((v & 0x01) == 0); ok {
			s = strconv.Itoa(v)
		}
		return
	}

	dst := SliceCollect(src, fn)
	assert.True(t, dst.Equals(SliceFrom("2", "4", "6", "8"), Equal[string]))
}

func TestSliceCollectFirst(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	fn := func(v int) (s string, ok bool) {
		if ok = ((v & 0x01) == 0); ok {
			s = strconv.Itoa(v)
		}
		return
	}

	dst := SliceCollectFirst(src, fn)
	assert.Equal(t, "2", dst.Get())

	fn = func(_ int) (string, bool) {
		return "", false
	}

	assert.False(t, SliceCollectFirst(src, fn).IsDefined())
}

func TestSliceScanLeft(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	fn := func(v1, v2 int) int {
		return v1 + v2
	}

	dst := SliceScanLeft(src, 100, fn)
	ans := SliceFrom(100, 101, 104, 109, 116, 125, 127, 131, 137, 145)
	assert.True(t, dst.Equals(ans, Equal[int]))

	dst = SliceScan(src, 100, fn)
	assert.True(t, dst.Equals(ans, Equal[int]))
}

func TestSliceScanRight(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	fn := func(v1, v2 int) int {
		return v1 + v2
	}
	dst := SliceScanRight(src, 100, fn)
	ans := SliceFrom(145, 144, 141, 136, 129, 120, 118, 114, 108, 100)
	assert.True(t, dst.Equals(ans, Equal[int]))
}

func TestSliceFoldLeft(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	fn1 := func(v1, v2 int) int {
		return v1 - v2
	}

	fn2 := func(v1, v2 int) int {
		return v1 + v2
	}

	assert.Equal(t, -45, SliceFoldLeft(src, 0, fn1))
	assert.Equal(t, -45, SliceFold(src, 0, fn1))

	assert.Equal(t, 45, SliceFoldLeft(src, 0, fn2))
	assert.Equal(t, 45, SliceFoldLeft(src, 0, fn2))
}

func TestSliceFoldRight(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	fn1 := func(v1, v2 int) int {
		return v1 - v2
	}

	fn2 := func(v1, v2 int) int {
		return v1 + v2
	}

	assert.Equal(t, 9, SliceFoldRight(src, 0, fn1))
	assert.Equal(t, 45, SliceFoldRight(src, 0, fn2))
}

func TestSliceFlatMap(t *testing.T) {
	dst := SliceFlatMap(SliceFrom(1, 2, 3), func(v int) Sliceable[int] {
		return SliceMap(SliceFrom(4, 5, 6), func(x int) int {
			return v * x
		})
	})

	ans := SliceFrom(4, 5, 6, 8, 10, 12, 12, 15, 18)
	assert.True(t, dst.Equals(ans, Equal[int]))
}

func TestSlicePartitionMap(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	fn := func(v int) Either[int, int] {
		if (v & 0x01) == 0 {
			return Right[int, int](v)
		}
		return Left[int, int](10 + v)
	}

	a, b := SlicePartitionMap(src, fn)
	assert.True(t, SliceFrom(11, 13, 15, 17, 19).Equals(a, Equal[int]))
	assert.True(t, SliceFrom(2, 4, 6, 8).Equals(b, Equal[int]))
}

func TestSliceGroupBy(t *testing.T) {
	var src = SliceFrom(1, 3, 5, 7, 9, 2, 4, 6, 8)

	fn := func(v int) bool {
		return (v & 0x01) == 0
	}
	m := SliceGroupBy(src, fn)
	t.Log(m)
}