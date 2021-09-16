package goscala

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakePair(t *testing.T) {
	v1 := 1
	v2 := "a"

	pair := MakePair(v1, v2)
	a1, a2 := pair.Get()

	assert.Equal(t, v1, a1)
	assert.Equal(t, v2, a2)
	assert.Equal(t, v1, pair.Key())
	assert.Equal(t, v2, pair.Value())
	assert.Equal(t, v1, pair.V1())
	assert.Equal(t, v2, pair.V2())
}
