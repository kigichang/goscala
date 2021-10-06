// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package try_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/kigichang/goscala"
	"github.com/kigichang/goscala/try"
	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {
	v := 0
	tr := try.Success[int](v)
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
	assert.Panics(t, func() { try.Failure[int](nil) })

	testErr := fmt.Errorf("test failure")
	tr := try.Failure[int](testErr)
	t.Log(tr)
	assert.Equal(t, false, tr.IsSuccess())
	assert.Equal(t, true, tr.IsFailure())
	assert.Equal(t, testErr, tr.Failed())
	assert.Panics(t, func() { tr.Get() })
}

func TestTryEquals(t *testing.T) {
	err := fmt.Errorf("%w", goscala.ErrUnsupported)

	tr := try.Failure[int](err)

	assert.True(t,
		tr.Equals(goscala.Equal[int])(
			try.Failure[int](goscala.ErrUnsupported)))

	assert.False(t,
		tr.Equals(goscala.Eq[int])(try.Failure[int](fmt.Errorf("Test"))))

	tr = try.Success(100)
	assert.True(t, tr.Equals(goscala.Eq[int])(try.Success(100)))
	assert.False(t, tr.Equals(goscala.Eq[int])(try.Success(101)))

	assert.False(t, try.Failure[int](goscala.ErrUnsupported).Equals(goscala.Eq[int])(try.Success(100)))
	assert.False(t, try.Success(100).Equals(goscala.Eq[int])(try.Failure[int](goscala.ErrUnsupported)))

}

func TestTryFilter(t *testing.T) {
	predict := func(v int) bool {
		return v > 0
	}

	tr := try.Success(1)
	tr2 := tr.Filter(predict)
	assert.True(t, tr.IsSuccess())
	assert.Equal(t, tr.Get(), tr2.Get())

	tr = try.Success(-1)
	tr2 = tr.Filter(predict)
	assert.True(t, tr2.IsFailure())
	assert.Equal(t, goscala.ErrUnsatisfied, tr2.Failed())

	err := fmt.Errorf("tr filter error")
	tr = try.Failure[int](err)
	tr2 = tr.Filter(predict)
	assert.True(t, tr2.IsFailure())
	assert.Equal(t, tr.Failed(), tr2.Failed())
	assert.Equal(t, err, tr2.Failed())
}

func TestTryForeach(t *testing.T) {
	sum := 1
	try.Success(1).Foreach(func(v int) {
		sum += 1
	})

	assert.Equal(t, 1+1, sum)

	sum = 1
	try.Failure[int](goscala.ErrUnsupported).Foreach(func(v int) {
		sum += 1
	})
	assert.Equal(t, 1, sum)
}

func TestTryGetOrElse(t *testing.T) {
	assert.Equal(t, 1, try.Success(1).GetOrElse(-1))
	assert.Equal(t, -1, try.Failure[int](goscala.ErrUnsatisfied).GetOrElse(-1))
}

func TestTryOrElse(t *testing.T) {
	ans := try.Success(-1)
	assert.Equal(t, 1, try.Success(1).OrElse(ans).Get())
	assert.Equal(t, -1, try.Failure[int](goscala.ErrUnsupported).OrElse(ans).Get())
}

func TestTryRecover(t *testing.T) {
	r := func(_ error) (int, bool) {
		return 0, true
	}

	assert.Equal(t, 1, try.Success(1).Recover(r).Get())
	assert.Equal(t, 0, try.Failure[int](goscala.ErrUnsupported).Recover(r).Get())

	r = func(err error) (a int, ok bool) {
		if ok = errors.Is(err, goscala.ErrUnsupported); ok {
			a = 0
		}
		return
	}

	assert.Equal(t, 1, try.Success(1).Recover(r).Get())
	assert.Equal(t, 0, try.Failure[int](goscala.ErrUnsupported).Recover(r).Get())

	err := fmt.Errorf("test try recover error")
	assert.Equal(t, err, try.Failure[int](err).Recover(r).Failed())

}

func TestTryRecoverWith(t *testing.T) {
	r := func(_ error) (goscala.Try[int], bool) {
		return try.Success(0), true
	}

	assert.Equal(t, 1, try.Success(1).RecoverWith(r).Get())
	assert.Equal(t, 0, try.Failure[int](goscala.ErrUnsupported).RecoverWith(r).Get())

	r = func(err error) (ret goscala.Try[int], ok bool) {
		if ok = errors.Is(err, goscala.ErrUnsupported); ok {
			ret = try.Success(0)
		}
		return
	}

	assert.Equal(t, 1, try.Success(1).RecoverWith(r).Get())
	assert.Equal(t, 0, try.Failure[int](goscala.ErrUnsupported).RecoverWith(r).Get())

	err := fmt.Errorf("test try recover with error")
	assert.Equal(t, err, try.Failure[int](err).RecoverWith(r).Failed())

}
