package goscala

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeTuple3(t *testing.T) {
	v1 := 1
	v2 := "a"
	v3 := 123.4

	tup := MakeTuple3(v1, v2, v3)
	a1, a2, a3 := tup.Get()

	assert.Equal(t, v1, a1)
	assert.Equal(t, v2, a2)
	assert.Equal(t, v3, a3)
	assert.Equal(t, v1, tup.V1())
	assert.Equal(t, v2, tup.V2())
	assert.Equal(t, v3, tup.V3())
}
