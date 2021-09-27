package goscala_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/kigichang/gomonad"
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

func TestTryOption(t *testing.T) {
	tr := goscala.Success(0)
	assert.Equal(t, goscala.Some(0).Get(), tr.Option().Get())

	tr = goscala.Failure[int](goscala.ErrUnsupported)
	assert.False(t, tr.Option().IsDefined())
}

func TestTrySlice(t *testing.T) {
	tr := goscala.Success(0)
	assert.Equal(t, 0, tr.Slice()[0])

	tr = goscala.Failure[int](goscala.ErrUnsupported)
	assert.Equal(t, 0, len(tr.Slice()))
}

func TestTryEquals(t *testing.T) {
	err := fmt.Errorf("%w", goscala.ErrUnsupported)

	tr := goscala.Failure[int](err)

	assert.True(t,
		tr.Equals(gomonad.Equal[int])(
			goscala.Failure[int](goscala.ErrUnsupported)))

	assert.False(t,
		tr.Equals(gomonad.Eq[int])(goscala.Failure[int](fmt.Errorf("Test"))))

	tr = goscala.Success(100)
	assert.True(t, tr.Equals(gomonad.Eq[int])(goscala.Success(100)))
	assert.False(t, tr.Equals(gomonad.Eq[int])(goscala.Success(101)))

	assert.False(t, goscala.Failure[int](goscala.ErrUnsupported).Equals(gomonad.Eq[int])(goscala.Success(100)))
	assert.False(t, goscala.Success(100).Equals(gomonad.Eq[int])(goscala.Failure[int](goscala.ErrUnsupported)))

}

func TesttrFilter(t *testing.T) {
	predict := func(v int) bool {
		return v > 0
	}

	tr := goscala.Success(1)
	tr2 := tr.Filter(predict)
	assert.True(t, tr.IsSuccess())
	assert.Equal(t, tr.Get(), tr2.Get())

	tr = goscala.Success(-1)
	tr2 = tr.Filter(predict)
	assert.True(t, tr2.IsFailure())
	assert.Equal(t, goscala.ErrUnsatisfied, tr2.Failed())

	err := fmt.Errorf("tr filter error")
	tr = goscala.Failure[int](err)
	tr2 = tr.Filter(predict)
	assert.True(t, tr2.IsFailure())
	assert.Equal(t, tr.Failed(), tr2.Failed())
	assert.Equal(t, err, tr2.Failed())
}

func TestTryForeach(t *testing.T) {
	sum := 1
	goscala.Success(1).Foreach(func(v int) {
		sum += 1
	})

	assert.Equal(t, 1+1, sum)

	sum = 1
	goscala.Failure[int](goscala.ErrUnsupported).Foreach(func(v int) {
		sum += 1
	})
	assert.Equal(t, 1, sum)
}

func TestTryGetOrElse(t *testing.T) {
	assert.Equal(t, 1, goscala.Success(1).GetOrElse(-1))
	assert.Equal(t, -1, goscala.Failure[int](goscala.ErrUnsatisfied).GetOrElse(-1))
}

func TestTryOrElse(t *testing.T) {
	ans := goscala.Success(-1)
	assert.Equal(t, 1, goscala.Success(1).OrElse(ans).Get())
	assert.Equal(t, -1, goscala.Failure[int](goscala.ErrUnsupported).OrElse(ans).Get())
}

func TestTryRecover(t *testing.T) {
	r := func(_ error) (int, bool) {
		return 0, true
	}

	assert.Equal(t, 1, goscala.Success(1).Recover(r).Get())
	assert.Equal(t, 0, goscala.Failure[int](goscala.ErrUnsupported).Recover(r).Get())

	r = func(err error) (a int, ok bool) {
		if ok = errors.Is(err, goscala.ErrUnsupported); ok {
			a = 0
		}
		return
	}

	assert.Equal(t, 1, goscala.Success(1).Recover(r).Get())
	assert.Equal(t, 0, goscala.Failure[int](goscala.ErrUnsupported).Recover(r).Get())

	err := fmt.Errorf("test try recover error")
	assert.Equal(t, err, goscala.Failure[int](err).Recover(r).Failed())

}

func TestTryRecoverWith(t *testing.T) {
	r := func(_ error) (goscala.Try[int], bool) {
		return goscala.Success(0), true
	}

	assert.Equal(t, 1, goscala.Success(1).RecoverWith(r).Get())
	assert.Equal(t, 0, goscala.Failure[int](goscala.ErrUnsupported).RecoverWith(r).Get())

	r = func(err error) (ret goscala.Try[int], ok bool) {
		if ok = errors.Is(err, goscala.ErrUnsupported); ok {
			ret = goscala.Success(0)
		}
		return
	}

	assert.Equal(t, 1, goscala.Success(1).RecoverWith(r).Get())
	assert.Equal(t, 0, goscala.Failure[int](goscala.ErrUnsupported).RecoverWith(r).Get())

	err := fmt.Errorf("test try recover with error")
	assert.Equal(t, err, goscala.Failure[int](err).RecoverWith(r).Failed())

}
