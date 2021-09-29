package goscala

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCurrying2(t *testing.T) {
	sum := func(a, b int) int {
		return a + b
	}

	curried := Curried2(sum)

	//Curried2[int, int, int](sum)

	x := curried(1)

	assert.Equal(t, sum(1, 2), x(2))
}
