package goscala_test

import (
	"testing"

	"github.com/kigichang/goscala"
	"github.com/stretchr/testify/assert"
)

func TestSome(t *testing.T) {
	v := 0
	o := goscala.Some[int](v)
	t.Log(o)
	assert.Equal(t, true, o.IsDefined())
	assert.Equal(t, v, o.Get())

	x, ok := o.Fetch()
	assert.Equal(t, v, x)
	assert.True(t, ok)

	v = 1
	o = goscala.Some[int](v)
	assert.Equal(t, true, o.IsDefined())
	assert.Equal(t, v, o.Get())

	x, ok = o.Fetch()
	assert.Equal(t, v, x)
	assert.True(t, ok)
}

func TestNone(t *testing.T) {
	o := goscala.None[int]()
	t.Log(o)
	assert.Equal(t, false, o.IsDefined())
	assert.Panics(t, func() { o.Get() })
	x, ok := o.Fetch()
	assert.Equal(t, 0, x)
	assert.False(t, ok)
}
