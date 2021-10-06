package goscala_test

import (
	"testing"

	"github.com/kigichang/goscala/either"
)

func TestLeft(t *testing.T) {
	//impl.Left[int, int](0)
	either.Left[int, int](100)

	//assert.True(t, left.IsLeft())
}
