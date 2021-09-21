package goscala

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeTuple2(t *testing.T) {
	v1 := 1
	v2 := "a"

	tup := MakeTuple2(v1, v2)
	a1, a2 := tup.Get()

	assert.Equal(t, v1, a1)
	assert.Equal(t, v2, a2)
	assert.Equal(t, v1, tup.V1())
	assert.Equal(t, v2, tup.V2())
}
