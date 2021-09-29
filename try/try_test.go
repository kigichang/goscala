package try_test

import (
	"fmt"
	"testing"

	"github.com/kigichang/goscala"
	"github.com/kigichang/goscala/try"
	"github.com/stretchr/testify/assert"
)

func TestErr(t *testing.T) {
	v := 1
	err := fmt.Errorf("try error")

	tr := try.Err[int](v, nil)

	t.Log(tr)
	assert.Equal(t, true, tr.IsSuccess())
	assert.Equal(t, false, tr.IsFailure())
	assert.Equal(t, v, tr.Get())
	assert.Equal(t, goscala.ErrUnsupported, tr.Failed())

	tr = try.Err[int](v, err)
	fmt.Println(tr)
	assert.Equal(t, false, tr.IsSuccess())
	assert.Equal(t, true, tr.IsFailure())
	assert.Equal(t, err, tr.Failed())
	assert.Panics(t, func() { tr.Get() })
}
