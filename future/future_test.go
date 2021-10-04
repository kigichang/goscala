package future_test

import (
	"context"
	"testing"
	"time"

	gs "github.com/kigichang/goscala"
	"github.com/kigichang/goscala/future"
	"github.com/stretchr/testify/assert"
)

func TestErr(t *testing.T) {
	f := future.Err(func() (int, error) { return 0, nil })
	f.Wait()
	t.Log(f)
	assert.True(t, f.Completed())

	v, err := f.Result(time.Second)
	assert.Equal(t, 0, v)
	assert.Nil(t, err)
}

func TestFlatMapAndMap(t *testing.T) {
	f := future.Err(func() (int, error) { return 5, nil })
	g := future.Err(func() (int, error) { return 3, nil })

	h := future.FlatMap(context.Background(), f, func(a int) gs.Future[int] {
		return future.Map(context.Background(), g, func(b int) int {
			return a * b
		})
	})

	h.Wait()

	t.Log(h)
}

func TestFilter(t *testing.T) {
	/*
		val f = Future { 5 }
		val g = f filter { _ % 2 == 1 }
		val h = f filter { _ % 2 == 0 }
		g foreach println // Eventually prints 5
		Await.result(h, Duration.Zero) // throw a NoSuchElementException
	*/

	f := future.Err(func() (int, error) {
		return 5, nil
	})

	p1 := func(a int) bool {
		return (a%2 == 1)
	}

	p2 := func(a int) bool {
		return (a % 2) == 0
	}

	g := f.Filter(context.Background(), p1)
	v, err := g.Result(time.Second)
	assert.Equal(t, 5, v)
	assert.Nil(t, err)

	h := f.Filter(context.Background(), p2)
	v, err = h.Result(time.Second)
	assert.Equal(t, 0, v)
	assert.Equal(t, gs.ErrUnsatisfied, err)

}
