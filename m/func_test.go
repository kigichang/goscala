package m_test

import (
	"testing"

	"github.com/kigichang/goscala/m"
	"github.com/stretchr/testify/assert"
)

func TestCurrying2(t *testing.T) {
	sum := func(a, b int) int {
		return a + b
	}

	assert.Equal(t, sum(1, 2), m.Currying2(sum)(1)(2))
}

func TestCurrying3(t *testing.T) {
	sum := func(a, b, c int) int {
		return a + b + c
	}
	assert.Equal(t, sum(1, 2, 3), m.Currying3(sum)(1)(2)(3))
}

func TestCurrying3To2(t *testing.T) {
	sum := func(a, b, c int) int {
		return a + b + c
	}
	assert.Equal(t, sum(1, 2, 3), m.Currying3To2(sum)(1)(2, 3))
}

func TestFuncCompose(t *testing.T) {
	negative := func(a int) int {
		return -a
	}

	square := func(a int) int {
		return a * a
	}

	assert.Equal(t, square(negative(3)), m.FuncCompose(square, negative)(3))
	assert.Equal(t, negative(square(3)), m.FuncCompose(negative, square)(3))
}

func TestFuncAndThen(t *testing.T) {
	negative := func(a int) int {
		return -a
	}

	square := func(a int) int {
		return a * a
	}

	assert.Equal(t, square(negative(3)), m.FuncAndThen(negative, square)(3))
	assert.Equal(t, negative(square(3)), m.FuncAndThen(square, negative)(3))
}
