package goscala_test

import (
	"testing"

	"github.com/kigichang/goscala"
	"github.com/stretchr/testify/assert"
)

func TestLeft(t *testing.T) {
	v := 0
	e := goscala.Left[int, string](v)
	t.Log(e)
	assert.Equal(t, true, e.IsLeft())
	assert.Equal(t, false, e.IsRight())
	assert.Equal(t, v, e.Left())
	assert.Panics(t, func() { e.Right() })
	assert.Panics(t, func() { e.Get() })
	_, ok := e.Fetch()
	assert.False(t, ok)
}

func TestRight(t *testing.T) {
	v := "hello"
	e := goscala.Right[int, string](v)
	t.Log(e)
	assert.Equal(t, false, e.IsLeft())
	assert.Equal(t, true, e.IsRight())
	assert.Equal(t, v, e.Right())
	assert.Equal(t, v, e.Get())
	assert.Panics(t, func() { e.Left() })

	x, ok := e.Fetch()
	assert.Equal(t, v, x)
	assert.True(t, ok)
}