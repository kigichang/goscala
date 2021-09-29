package goscala

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	assert.True(t, Equal(100, 100))
	assert.False(t, Equal("abc", "ABC"))
}

func TestCompare(t *testing.T) {
	assert.Equal(t, 0, Compare(1, 1))
	assert.Equal(t, 1, Compare(2, 1))
	assert.Equal(t, -1, Compare(2, 3))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 1, Max(-1, 1))
	assert.Equal(t, 1, Max(1, -1))
	assert.Equal(t, -1, Min(-1, 1))
	assert.Equal(t, -1, Min(1, -1))
}
