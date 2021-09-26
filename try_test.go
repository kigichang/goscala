package goscala_test

import (
	"fmt"
	"testing"

	"github.com/kigichang/goscala"
	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {
	v := 0
	tr := goscala.Success[int](v)
	t.Log(tr)
	assert.Equal(t, true, tr.IsSuccess())
	assert.Equal(t, false, tr.IsFailure())
	assert.Equal(t, v, tr.Get())
	assert.Equal(t, goscala.ErrUnsupported, tr.Failed())
	assert.Equal(t, v, tr.Success())
	v2, ok := tr.Fetch()
	assert.Equal(t, v, v2)
	assert.True(t, ok)
}

func TestFailure(t *testing.T) {
	assert.Panics(t, func() { goscala.Failure[int](nil) })

	testErr := fmt.Errorf("test failure")
	tr := goscala.Failure[int](testErr)
	t.Log(tr)
	assert.Equal(t, false, tr.IsSuccess())
	assert.Equal(t, true, tr.IsFailure())
	assert.Equal(t, testErr, tr.Failed())
	assert.Panics(t, func() { tr.Get() })
}
