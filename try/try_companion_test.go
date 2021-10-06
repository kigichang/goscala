// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package try_test

import (
	"fmt"
	"strconv"
	"testing"

	gs "github.com/kigichang/goscala"
	"github.com/kigichang/goscala/try"
	"github.com/stretchr/testify/assert"
)

func TestErr(t *testing.T) {
	v := 1
	err := fmt.Errorf("tr error")
	tr := try.Err[int](v, nil)

	t.Log(tr)
	assert.Equal(t, true, tr.IsSuccess())
	assert.Equal(t, false, tr.IsFailure())
	assert.Equal(t, v, tr.Get())
	assert.Equal(t, gs.ErrUnsupported, tr.Failed())

	tr = try.Err[int](v, err)
	t.Log(tr)
	assert.Equal(t, false, tr.IsSuccess())
	assert.Equal(t, true, tr.IsFailure())
	assert.Equal(t, err, tr.Failed())
	assert.Panics(t, func() { tr.Get() })
}

func TestCollect(t *testing.T) {

	pf := func(v int) (s string, ok bool) {
		if ok = (v > 0); ok {
			s = strconv.Itoa(v)
		}
		return
	}
	err := fmt.Errorf("tr collect error")

	tr := try.Success(1)
	tr2 := try.Collect(tr, pf)
	assert.True(t, tr2.IsSuccess())
	assert.Equal(t, "1", tr2.Get())

	tr = try.Failure[int](err)
	tr2 = try.Collect(tr, pf)
	assert.True(t, tr2.IsFailure())
	assert.Equal(t, err, tr2.Failed())

	tr = try.Success(-1)
	tr2 = try.Collect(tr, pf)
	assert.True(t, tr2.IsFailure())
	assert.Equal(t, gs.ErrUnsatisfied, tr2.Failed())
}

func TestFlatMap(t *testing.T) {
	f := func(v int) gs.Try[string] {
		return try.Success(fmt.Sprintf("%d", v))
	}

	v := 100
	tr := try.Success(v)
	tr2 := try.FlatMap(tr, f)
	assert.True(t, tr2.IsSuccess())
	assert.Equal(t, fmt.Sprintf(`%d`, v), tr2.Get())

	err := fmt.Errorf("tr flatmap error")
	tr = try.Failure[int](err)
	tr2 = try.FlatMap(tr, f)
	assert.True(t, tr2.IsFailure())
	assert.Equal(t, err, tr2.Failed())
}

func TestFold(t *testing.T) {
	fail := func(err error) string {
		return err.Error()
	}

	succ := strconv.Itoa

	v := 100
	errStr := `test fold error`
	err := fmt.Errorf(errStr)

	tr := try.Success(v)
	ans := try.Fold(tr, succ, fail)
	assert.Equal(t, fmt.Sprintf(`%v`, v), ans)

	tr = try.Failure[int](err)
	ans = try.Fold(tr, succ, fail)
	assert.Equal(t, ans, errStr)
}

func TestMap(t *testing.T) {
	tr := try.Map(try.Success(100), strconv.Itoa)
	assert.True(t, tr.IsSuccess())
	assert.Equal(t, "100", tr.Get())

	tr = try.Map(try.Failure[int](gs.ErrEmpty), strconv.Itoa)
	assert.True(t, tr.IsFailure())
	assert.Equal(t, gs.ErrEmpty, tr.Failed())
}

func TestMapErr(t *testing.T) {
	tr := try.MapErr(try.Success("100"), strconv.Atoi)
	assert.True(t, tr.IsSuccess())
	assert.Equal(t, 100, tr.Get())

	tr = try.MapErr(try.Failure[string](gs.ErrEmpty), strconv.Atoi)
	assert.True(t, tr.IsFailure())
	assert.Equal(t, gs.ErrEmpty, tr.Failed())

	tr = try.MapErr(try.Success("abc"), strconv.Atoi)
	assert.True(t, tr.IsFailure())
}

func TestMapBool(t *testing.T) {
	m := map[int]string{
		1: "1",
	}

	f := func(v int) (ret string, ok bool) {
		ret, ok = m[v]
		return
	}

	tr := try.MapBool(try.Success(1), f)
	assert.True(t, tr.IsSuccess())
	assert.Equal(t, "1", tr.Get())

	tr = try.MapBool(try.Failure[int](gs.ErrEmpty), f)
	assert.True(t, tr.IsFailure())
	assert.Equal(t, gs.ErrEmpty, tr.Failed())

	tr = try.MapBool(try.Success(2), f)
	assert.True(t, tr.IsFailure())
	assert.Equal(t, gs.ErrUnsatisfied, tr.Failed())
}

func TestTryTransform(t *testing.T) {
	succ := func(v string) gs.Try[int] {
		return try.Err(strconv.Atoi(v))
	}

	fail := try.Failure[int]

	v := 123
	tr := try.Success(fmt.Sprintf(`%d`, v))
	ans := try.Transform(tr, succ, fail)
	assert.True(t, ans.IsSuccess())
	assert.Equal(t, v, ans.Get())

	tr = try.Failure[string](gs.ErrEmpty)
	ans = try.Transform(tr, succ, fail)
	assert.True(t, ans.IsFailure())
	assert.Equal(t, gs.ErrEmpty, ans.Failed())
}
