package goscala

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeTuple4(t *testing.T) {
	v1 := 1
	v2 := "a"
	v3 := 123.4
	v4 := [...]int{1, 2, 3}

	tup := MakeTuple4(v1, v2, v3, v4)
	a1, a2, a3, a4 := tup.Get()

	assert.Equal(t, v1, a1)
	assert.Equal(t, v2, a2)
	assert.Equal(t, v3, a3)
	assert.Equal(t, v4, a4)
	assert.Equal(t, v1, tup.V1())
	assert.Equal(t, v2, tup.V2())
	assert.Equal(t, v3, tup.V3())
	assert.Equal(t, v4, tup.V4())
}