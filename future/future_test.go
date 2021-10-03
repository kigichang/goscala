package future_test

import (
	"context"
	"testing"

	gs "github.com/kigichang/goscala"
	"github.com/kigichang/goscala/future"
	"github.com/stretchr/testify/assert"
)

func TestErr(t *testing.T) {
	f := future.Err(context.Background(), func() (int, error) { return 0, nil })
	f.Wait()
	t.Log(f)
	assert.True(t, f.Completed())
	assert.Equal(t, 0, f.Value().Get())
}

func TestFlatMapAndMap(t *testing.T) {
	f := future.Err(context.Background(), func() (int, error) { return 5, nil })
	g := future.Err(context.Background(), func() (int, error) { return 3, nil })

	h := future.FlatMap(context.Background(), f, func(a int) gs.Future[int] {
		return future.Map(context.Background(), g, func(b int) int {
			return a * b
		})
	})

	h.Wait()

	t.Log(h)
}
