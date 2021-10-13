package goscala_test

import (
	"testing"

	gs "github.com/kigichang/goscala"
	"github.com/stretchr/testify/assert"
)

func TestTrue(t *testing.T) {
	assert.True(t, gs.True())
}

func TestFalse(t *testing.T) {
	assert.False(t, gs.False())
}

func TestCond(t *testing.T) {
	assert.Equal(t, 1, gs.Cond(true, 1, 0))
	assert.Equal(t, 0, gs.Cond(false, 1, 0))
}

func TestDefault(t *testing.T) {
	assert.Equal(t, 1, gs.Default(0)(func() (int, bool) {
		return 1, true
	}))

	assert.Equal(t, 0, gs.Default(0)(func() (int, bool) {
		return -1, false
	}))
}
