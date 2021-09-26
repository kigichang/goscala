package goscala_test

import (
	"fmt"
	"testing"

	"github.com/kigichang/goscala"
	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {
	v := 0
	try := goscala.Success[int](v)
	t.Log(try)
	assert.Equal(t, true, try.IsSuccess())
	assert.Equal(t, false, try.IsFailure())
	assert.Equal(t, v, try.Get())
	assert.Equal(t, goscala.ErrUnsupported, try.Failed())
}

func TestFailure(t *testing.T) {
	assert.Panics(t, func() { goscala.Failure[int](nil) })

	testErr := fmt.Errorf("test failure")
	try := goscala.Failure[int](testErr)
	t.Log(try)
	assert.Equal(t, false, try.IsSuccess())
	assert.Equal(t, true, try.IsFailure())
	assert.Equal(t, testErr, try.Failed())
	assert.Panics(t, func() { try.Get() })
}
