// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package iter_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/kigichang/goscala/iter"
	"github.com/stretchr/testify/assert"
)

func TestGenIter(t *testing.T) {

	ans := []int{1, 2, 3, 4}
	it := iter.Gen(ans...)
	assert.Equal(t, len(ans), it.Len())
	assert.Equal(t, cap(ans), it.Cap())
	assert.Equal(t, iter.Slice(it), ans)
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

func TestForAll(t *testing.T) {
	src := []int{1, 3, 5, 7, 9, 2, 4, 6, 8}

	p1 := func(v int) bool {
		return v >= 0
	}

	p2 := func(v int) bool {
		return v > 5
	}

	assert.True(t, iter.Forall(iter.Gen([]int{}...), p2))
	assert.True(t, iter.Forall(iter.Gen(src...), p1))
	assert.False(t, iter.Forall(iter.Gen(src...), p2))
}

func TestForeach(t *testing.T) {
	src := []int{1, 3, 5, 7, 9, 2, 4, 6, 8}

	sum := 0
	iter.Foreach(iter.Gen(src...), func(v int) {
		sum += v
	})
	assert.Equal(t, 45, sum)
}

func TestFilter(t *testing.T) {
	src := []int{1, 3, 5, 7, 9, 2, 4, 6, 8}

	p := func(v int) bool {
		return (v & 0x01) == 0
	}

	s := iter.Slice(iter.Filter(iter.Gen(src...), p))

	assert.Equal(t, []int{2, 4, 6, 8}, s)
}

func TestFilterNot(t *testing.T) {
	src := []int{1, 3, 5, 7, 9, 2, 4, 6, 8}

	p := func(v int) bool {
		return (v & 0x01) == 0
	}

	s := iter.Slice(iter.FilterNot(iter.Gen(src...), p))

	assert.Equal(t, []int{1, 3, 5, 7, 9}, s)
}

type ctxKey int

const x ctxKey = 1

func TestContext(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	str := "abc"
	r1 := context.WithValue(ctx, x, &str)

	w := &sync.WaitGroup{}

	go func() {
		w.Add(1)
		defer w.Done()
		select {
		case <-r1.Done():
			v := r1.Value(x).(*string)
			t.Log("r1:", *v)
		}
	}()

	time.Sleep(5 * time.Second)
	str = "def"
	cancel()

	r2 := context.WithValue(ctx, x, &str)
	go func() {
		w.Add(1)
		defer w.Done()
		select {
		case <-r2.Done():
			v := r2.Value(x).(*string)
			t.Log("r2:", *v)
		}
	}()

	w.Wait()

}
