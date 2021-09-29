package try_test

import (
	"fmt"
<<<<<<< HEAD
	"testing"

	"github.com/kigichang/goscala"
=======
	"strconv"
	"testing"

	gs "github.com/kigichang/goscala"
>>>>>>> 0ec6c5065beced5aa1c8726cf96ee1da6ef6d566
	"github.com/kigichang/goscala/try"
	"github.com/stretchr/testify/assert"
)

func TestErr(t *testing.T) {
	v := 1
<<<<<<< HEAD
	err := fmt.Errorf("try error")
=======
	err := fmt.Errorf("tr error")
>>>>>>> 0ec6c5065beced5aa1c8726cf96ee1da6ef6d566

	tr := try.Err[int](v, nil)

	t.Log(tr)
	assert.Equal(t, true, tr.IsSuccess())
	assert.Equal(t, false, tr.IsFailure())
	assert.Equal(t, v, tr.Get())
<<<<<<< HEAD
	assert.Equal(t, goscala.ErrUnsupported, tr.Failed())

	tr = try.Err[int](v, err)
	fmt.Println(tr)
=======
	assert.Equal(t, gs.ErrUnsupported, tr.Failed())

	tr = try.Err[int](v, err)
	t.Log(tr)
>>>>>>> 0ec6c5065beced5aa1c8726cf96ee1da6ef6d566
	assert.Equal(t, false, tr.IsSuccess())
	assert.Equal(t, true, tr.IsFailure())
	assert.Equal(t, err, tr.Failed())
	assert.Panics(t, func() { tr.Get() })
}
<<<<<<< HEAD
=======

func TestCollect(t *testing.T) {

	pf := func(v int) (s string, ok bool) {
		if ok = (v > 0); ok {
			s = strconv.Itoa(v)
		}
		return
	}
	err := fmt.Errorf("tr collect error")

	tr := gs.Success(1)
	tr2 := try.Collect[int, string](tr)(pf)
	assert.True(t, tr2.IsSuccess())
	assert.Equal(t, "1", tr2.Get())

	tr = gs.Failure[int](err)
	tr2 = try.Collect[int, string](tr)(pf)
	assert.True(t, tr2.IsFailure())
	assert.Equal(t, err, tr2.Failed())

	tr = gs.Success(-1)
	tr2 = try.Collect[int, string](tr)(pf)
	assert.True(t, tr2.IsFailure())
	assert.Equal(t, gs.ErrUnsatisfied, tr2.Failed())
}

func TestFlatMap(t *testing.T) {
	f := func(v int) gs.Try[string] {
		return gs.Success(fmt.Sprintf("%d", v))
	}

	v := 100
	tr := gs.Success(v)
	tr2 := try.FlatMap[int, string](tr)(f)
	assert.True(t, tr2.IsSuccess())
	assert.Equal(t, fmt.Sprintf(`%d`, v), tr2.Get())

	err := fmt.Errorf("tr flatmap error")
	tr = gs.Failure[int](err)
	tr2 = try.FlatMap[int, string](tr)(f)
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

	tr := gs.Success(v)
	ans := try.Fold[int, string](tr)(succ, fail)
	assert.Equal(t, fmt.Sprintf(`%v`, v), ans)

	tr = gs.Failure[int](err)
	ans = try.Fold[int, string](tr)(succ, fail)
	assert.Equal(t, ans, errStr)
}

func TestMap(t *testing.T) {
	tr := try.Map[int, string](gs.Success(100))(strconv.Itoa)
	assert.True(t, tr.IsSuccess())
	assert.Equal(t, "100", tr.Get())

	tr = try.Map[int, string](gs.Failure[int](gs.ErrEmpty))(strconv.Itoa)
	assert.True(t, tr.IsFailure())
	assert.Equal(t, gs.ErrEmpty, tr.Failed())
}

func TestMapErr(t *testing.T) {
	tr := try.MapErr[string, int](gs.Success("100"))(strconv.Atoi)
	assert.True(t, tr.IsSuccess())
	assert.Equal(t, 100, tr.Get())

	tr = try.MapErr[string, int](gs.Failure[string](gs.ErrEmpty))(strconv.Atoi)
	assert.True(t, tr.IsFailure())
	assert.Equal(t, gs.ErrEmpty, tr.Failed())

	tr = try.MapErr[string, int](gs.Success("abc"))(strconv.Atoi)
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

	tr := try.MapBool[int, string](gs.Success(1))(f)
	assert.True(t, tr.IsSuccess())
	assert.Equal(t, "1", tr.Get())

	tr = try.MapBool[int, string](gs.Failure[int](gs.ErrEmpty))(f)
	assert.True(t, tr.IsFailure())
	assert.Equal(t, gs.ErrEmpty, tr.Failed())

	tr = try.MapBool[int, string](gs.Success(2))(f)
	assert.True(t, tr.IsFailure())
	assert.Equal(t, gs.ErrUnsatisfied, tr.Failed())
}

func TestTryTransform(t *testing.T) {
	succ := func(v string) gs.Try[int] {
		return try.Err(strconv.Atoi(v))
	}

	fail := gs.Failure[int]

	v := 123
	tr := gs.Success(fmt.Sprintf(`%d`, v))
	ans := try.Transform[string, int](tr)(succ, fail)
	assert.True(t, ans.IsSuccess())
	assert.Equal(t, v, ans.Get())

	tr = gs.Failure[string](gs.ErrEmpty)
	ans = try.Transform[string, int](tr)(succ, fail)
	assert.True(t, ans.IsFailure())
	assert.Equal(t, gs.ErrEmpty, ans.Failed())
}
>>>>>>> 0ec6c5065beced5aa1c8726cf96ee1da6ef6d566
