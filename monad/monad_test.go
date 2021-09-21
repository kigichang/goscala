package monad_test

import (
	"fmt"
	"testing"
	"strconv"
	"github.com/kigichang/goscala/monad"
	"github.com/stretchr/testify/assert"
)


func sliceEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestFoldBool(t *testing.T) {
	fetch1 := func() (int, bool) {
		return 100, true
	}

	fetch2 := func() (int, bool) {
		return 0, false
	}

	z := func() string {
		return "false"
	}

	result := monad.FoldBool[int, string](fetch1)(z, strconv.Itoa)
	assert.Equal(t, "100", result)
	
	result = monad.FoldBool[int, string](fetch2)(z, strconv.Itoa)
	assert.Equal(t, "false", result)
}

func TestFoldErr(t *testing.T) {
	fetch1 := func() (int, error) {
		return 100, nil
	}

	fetch2 := func() (int, error) {
		return 0, fmt.Errorf("0")
	}

	z := func(err error) string {
		return err.Error()
	}

	result := monad.FoldErr[int, string](fetch1)(z, strconv.Itoa)
	assert.Equal(t, "100", result)
	
	result = monad.FoldErr[int, string](fetch2)(z, strconv.Itoa)
	assert.Equal(t, "0", result)
}

func TestFoldLeft(t *testing.T) {
	s := []int{1,2,3,4,5,6,7,8,9}

	x := monad.FoldLeft[int, int](s)(0)
	result := x(func(a, b int) int {
		return a + b
	})

	assert.Equal(t, 45, result)

	result = x(func(a, b int) int {
		return a - b
	})

	assert.Equal(t, -45, result)

}

func TestFoldRight(t *testing.T) {
	s := []int{1, 3, 5, 7, 9, 2, 4, 6, 8}

	x := monad.FoldRight[int, int](s)(0)
	result := x(func(a, b int) int {
		return a + b
	})

	assert.Equal(t, 45, result)
	result = x(func(a, b int) int {
		return a - b
	})
	assert.Equal(t, 9, result)
}

func TestMap(t *testing.T) {
	s := []int{1,2,3}

	x := monad.Map[int, int](s)

	result := x(func(a int) int {
		return a + 1
	})

	assert.Equal(t, 2, result[0])
	assert.Equal(t, 3, result[1])
	assert.Equal(t, 4, result[2])
}

func TestFlatMap(t *testing.T) {
	a := []int{1,2,3}
	b := []int{4,5,6}

	xa := monad.FlatMap[int, int](a)
	xb := monad.Map[int, int](b)

	result := xa(func(x int) []int {
		return xb(func(y int) int {
			return x * y
		})
	})

	ans := []int {4, 5, 6, 8, 10, 12, 12, 15, 18}

	assert.Equal(t, len(result), len(ans))

	for i := range ans {
		assert.Equal(t, ans[i], result[i])
	}
}

func TestScanLeft(t *testing.T) {
	s := []int {1, 3, 5, 7, 9, 2, 4, 6, 8}

	result := monad.ScanLeft[int, int](s)(100)(func(v1, v2 int) int {
		return v1 + v2
	})
	ans := []int {100, 101, 104, 109, 116, 125, 127, 131, 137, 145}
	assert.True(t, sliceEqual(ans, result))
}

func TestScanRight(t *testing.T) {
	s := []int {1, 3, 5, 7, 9, 2, 4, 6, 8}
	result := monad.ScanRight[int, int](s)(100)(func(v1, v2 int) int {
		return v1 + v2
	})

	ans := []int{145, 144, 141, 136, 129, 120, 118, 114, 108, 100}
	assert.True(t, sliceEqual(ans, result))
}