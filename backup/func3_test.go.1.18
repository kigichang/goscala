package goscala

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCurrying3(t *testing.T) {
	var sum Func3[int, int, int, int] = func(a, b, c int) int {
		return a + b + c
	}

	curried := sum.Curried()

	x := curried(1)

	assert.Equal(t, sum(1, 2, 3), x(2)(3))
}
