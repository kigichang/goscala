package goscala

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {
	v := 0
	try := Success[int](v)
	fmt.Println(try)
	assert.Equal(t, true, try.IsSuccess())
	assert.Equal(t, false, try.IsFailure())
	assert.Equal(t, v, try.Get())
	assert.Equal(t, ErrUnsupported, try.Failed())
}

func TestFailure(t *testing.T) {
	try := Failure[int](nil)
	fmt.Println(try)
	assert.Equal(t, false, try.IsSuccess())
	assert.Equal(t, true, try.IsFailure())
	assert.Equal(t, ErrNil, try.Failed())
	assert.Panics(t, func() { try.Get() })

	testErr := fmt.Errorf("test failure")
	try = Failure[int](testErr)
	fmt.Println(try)
	assert.Equal(t, false, try.IsSuccess())
	assert.Equal(t, true, try.IsFailure())
	assert.Equal(t, testErr, try.Failed())
	assert.Panics(t, func() { try.Get() })
}

func TestMakeTry(t *testing.T) {
	v := 1
	err := fmt.Errorf("try error")

	try := MakeTry[int](v, nil)

	fmt.Println(try)
	assert.Equal(t, true, try.IsSuccess())
	assert.Equal(t, false, try.IsFailure())
	assert.Equal(t, v, try.Get())
	assert.Equal(t, ErrUnsupported, try.Failed())

	try = MakeTry[int](v, err)
	fmt.Println(try)
	assert.Equal(t, false, try.IsSuccess())
	assert.Equal(t, true, try.IsFailure())
	assert.Equal(t, err, try.Failed())
	assert.Panics(t, func() { try.Get() })
}

func TestTryCollect(t *testing.T) {

	pf := func(v int) (s string, ok bool) {
		if ok = (v > 0); ok {
			s = strconv.Itoa(v)
		}
		return
	}
	err := fmt.Errorf("try collect error")

	try := Success(1)
	try2 := TryCollect(try, pf)
	assert.True(t, try2.IsSuccess())
	assert.Equal(t, "1", try2.Get())

	try = Failure[int](err)
	try2 = TryCollect(try, pf)
	assert.True(t, try2.IsFailure())
	assert.Equal(t, err, try2.Failed())

	try = Success(-1)
	try2 = TryCollect(try, pf)
	assert.True(t, try2.IsFailure())
	assert.Equal(t, ErrUnsatisfied, try2.Failed())
}

func TestTryFlatMap(t *testing.T) {
	f := func(v int) Try[string] {
		return Success(fmt.Sprintf("%d", v))
	}

	v := 100
	try := Success(v)
	try2 := TryFlatMap(try, f)
	assert.True(t, try2.IsSuccess())
	assert.Equal(t, fmt.Sprintf(`%d`, v), try2.Get())

	err := fmt.Errorf("try flatmap error")
	try = Failure[int](err)
	try2 = TryFlatMap(try, f)
	assert.True(t, try2.IsFailure())
	assert.Equal(t, err, try2.Failed())
}

func TestTryFold(t *testing.T) {
	fail := func(err error) string {
		return err.Error()
	}

	succ := strconv.Itoa

	v := 100
	errStr := `test fold error`
	err := fmt.Errorf(errStr)

	try := Success(v)
	ans := TryFold(try, succ, fail)
	assert.Equal(t, fmt.Sprintf(`%v`, v), ans)

	try = Failure[int](err)
	ans = TryFold(try, succ, fail)
	assert.Equal(t, ans, errStr)
}

func TestTryMap(t *testing.T) {
	try := TryMap(Success(100), strconv.Itoa)
	assert.True(t, try.IsSuccess())
	assert.Equal(t, "100", try.Get())

	try = TryMap(Failure[int](ErrNil), strconv.Itoa)
	assert.True(t, try.IsFailure())
	assert.Equal(t, ErrNil, try.Failed())
}

func TestTryMapWithErr(t *testing.T) {
	try := TryMapWithErr(Success("100"), strconv.Atoi)
	assert.True(t, try.IsSuccess())
	assert.Equal(t, 100, try.Get())

	try = TryMapWithErr(Failure[string](ErrNil), strconv.Atoi)
	assert.True(t, try.IsFailure())
	assert.Equal(t, ErrNil, try.Failed())

	try = TryMapWithErr(Success("abc"), strconv.Atoi)
	assert.True(t, try.IsFailure())
}

func TestTryMapWithBool(t *testing.T) {
	m := map[int]string{
		1: "1",
	}

	f := func(v int) (ret string, ok bool) {
		ret, ok = m[v]
		return
	}

	try := TryMapWithBool(Success(1), f)
	assert.True(t, try.IsSuccess())
	assert.Equal(t, "1", try.Get())

	try = TryMapWithBool(Failure[int](ErrNil), f)
	assert.True(t, try.IsFailure())
	assert.Equal(t, ErrNil, try.Failed())

	try = TryMapWithBool(Success(2), f)
	assert.True(t, try.IsFailure())
	assert.Equal(t, ErrUnsatisfied, try.Failed())
}

func TestTryTransform(t *testing.T) {
	succ := func(v string) Try[int] {
		return MakeTry(strconv.Atoi(v))
	}

	fail := func(err error) Try[int] {
		return Failure[int](err)
	}

	v := 123
	try := Success(fmt.Sprintf(`%d`, v))
	ans := TryTransform(try, succ, fail)
	assert.True(t, ans.IsSuccess())
	assert.Equal(t, v, ans.Get())

	try = Failure[string](ErrNil)
	ans = TryTransform(try, succ, fail)
	assert.True(t, ans.IsFailure())
	assert.Equal(t, ErrNil, ans.Failed())
}

func TestTryFilter(t *testing.T) {
	predict := func(v int) bool {
		return v > 0
	}

	try := Success(1)
	try2 := try.Filter(predict)
	assert.True(t, try.IsSuccess())
	assert.Equal(t, try.Get(), try2.Get())

	try = Success(-1)
	try2 = try.Filter(predict)
	assert.True(t, try2.IsFailure())
	assert.Equal(t, ErrUnsatisfied, try2.Failed())

	err := fmt.Errorf("try filter error")
	try = Failure[int](err)
	try2 = try.Filter(predict)
	assert.True(t, try2.IsFailure())
	assert.Equal(t, try.Failed(), try2.Failed())
	assert.Equal(t, err, try2.Failed())
}

func TestTryForeach(t *testing.T) {
	sum := 1
	Success(1).Foreach(func(v int) {
		sum += 1
	})

	assert.Equal(t, 1+1, sum)

	sum = 1
	Failure[int](ErrNil).Foreach(func(v int) {
		sum += 1
	})
	assert.Equal(t, 1, sum)
}

func TestTryGetOrElse(t *testing.T) {
	assert.Equal(t, 1, Success(1).GetOrElse(-1))
	assert.Equal(t, -1, Failure[int](ErrNil).GetOrElse(-1))
}

func TestTryOrElse(t *testing.T) {
	ans := Success(-1)
	assert.Equal(t, 1, Success(1).OrElse(ans).Get())
	assert.Equal(t, -1, Failure[int](ErrNil).OrElse(ans).Get())
}

func TestTryRecover(t *testing.T) {
	r := func(_ error) (int, bool) {
		return 0, true
	}

	assert.Equal(t, 1, Success(1).Recover(r).Get())
	assert.Equal(t, 0, Failure[int](ErrNil).Recover(r).Get())

	r = func(err error) (a int, ok bool) {
		if ok = (err == ErrNil); ok {
			a = 0
		}
		return
	}

	assert.Equal(t, 1, Success(1).Recover(r).Get())
	assert.Equal(t, 0, Failure[int](ErrNil).Recover(r).Get())

	err := fmt.Errorf("test try recover error")
	assert.Equal(t, err, Failure[int](err).Recover(r).Failed())

}

func TestTryRecoverWith(t *testing.T) {
	r := func(_ error) (Try[int], bool) {
		return Success(0), true
	}

	assert.Equal(t, 1, Success(1).RecoverWith(r).Get())
	assert.Equal(t, 0, Failure[int](ErrNil).RecoverWith(r).Get())

	r = func(err error) (ret Try[int], ok bool) {
		if ok = (err == ErrNil); ok {
			ret = Success(0)
		}
		return
	}

	assert.Equal(t, 1, Success(1).RecoverWith(r).Get())
	assert.Equal(t, 0, Failure[int](ErrNil).RecoverWith(r).Get())

	err := fmt.Errorf("test try recover with error")
	assert.Equal(t, err, Failure[int](err).RecoverWith(r).Failed())

}

func TestTryOption(t *testing.T) {
	assert.Equal(t, 1, Success(1).Option().Get())
	assert.False(t, Failure[int](ErrNil).Option().IsDefined())
}
