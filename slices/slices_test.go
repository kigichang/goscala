package slices_test

import (
	"strconv"
	"testing"

	gs "github.com/kigichang/goscala"
	"github.com/kigichang/goscala/opt"
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

	dst := slices.Collect[int, string](src)(fn)
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

	dst := opt.Bool[string](slices.CollectFirst[int, string](src)(fn))
	assert.Equal(t, "2", dst.Get())

	fn = func(_ int) (string, bool) {
		return "", false
	}

	assert.False(t, opt.Bool[string](slices.CollectFirst[int, string](src)(fn)).IsDefined())
}

func TestFlatMap(t *testing.T) {
	dst := slices.FlatMap[int, int](slices.From(1, 2, 3))(func(v int) gs.Sliceable[int] {
		return slices.Map[int, int](slices.From(4, 5, 6))(func(x int) int {
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

	assert.Equal(t, -45, slices.FoldLeft[int, int](src)(0)(fn1))
	assert.Equal(t, -45, slices.Fold[int, int](src)(0)(fn1))

	assert.Equal(t, 45, slices.FoldLeft[int, int](src)(0)(fn2))
	assert.Equal(t, 45, slices.Fold[int, int](src)(0)(fn2))
}

func TestFoldRight(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	fn1 := func(v1, v2 int) int {
		return v1 - v2
	}

	fn2 := func(v1, v2 int) int {
		return v1 + v2
	}

	assert.Equal(t, 9, slices.FoldRight[int, int](src)(0)(fn1))
	assert.Equal(t, 45, slices.FoldRight[int, int](src)(0)(fn2))
}

func TestPartitionMap(t *testing.T) {
	var src = slices.From(1, 3, 5, 7, 9, 2, 4, 6, 8)

	fn := func(v int) gs.Either[int, int] {
		if (v & 0x01) == 0 {
			return gs.Right[int, int](v)
		}
		return gs.Left[int, int](10 + v)
	}

	a, b := slices.PartitionMap[int, int, int](src)(fn)
	assert.True(t, slices.From(11, 13, 15, 17, 19).Equals(gs.Eq[int])(a))
	assert.True(t, slices.From(2, 4, 6, 8).Equals(gs.Eq[int])(b))
}
