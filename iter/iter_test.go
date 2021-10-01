package iter_test

import (
	"testing"

	"github.com/kigichang/goscala/iter"
	"github.com/stretchr/testify/assert"
)

func TestGenIter(t *testing.T) {

	it := iter.GenIter(1, 2, 3, 4)

	ss := iter.Slice(it)

	assert.Equal(t, ss, []int{1, 2, 3, 4})
}

func TestFoldLeft(t *testing.T) {
	src := []int{1, 3, 5, 7, 9, 2, 4, 6, 8}

	it := iter.GenIter(src...)
	fn1 := func(v1, v2 int) int {
		return v1 - v2
	}
	assert.Equal(t, -45, iter.FoldLeft(it, 0, fn1))

	it = iter.GenIter(src...)
	fn2 := func(v1, v2 int) int {
		return v1 + v2
	}
	assert.Equal(t, 45, iter.FoldLeft(it, 0, fn2))
}
