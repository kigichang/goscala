package opt_test

import (
	"fmt"
	"testing"
	"strconv"
	"github.com/kigichang/goscala"
	"github.com/kigichang/goscala/monad"
	"github.com/kigichang/goscala/opt"
	"github.com/stretchr/testify/assert"
)

func TestMakeWithBool(t *testing.T) {
	o := opt.MakeWithBool(0, true)

	assert.True(t, o.IsDefined())
	assert.Equal(t, 0, o.Get())

	v, ok := o.Fetch()
	assert.Equal(t, 0, v)
	assert.True(t, ok)

	o = opt.MakeWithBool(100, false)
	assert.False(t, o.IsDefined())
	assert.Panics(t, func() { o.Get() })

	v, ok = o.Fetch()
	assert.Equal(t, 0, v)
	assert.False(t, ok)

}

func TestMakeWitErr(t *testing.T) {
	o := opt.MakeWithErr(0, nil)

	assert.True(t, o.IsDefined())
	assert.Equal(t, 0, o.Get())

	v, ok := o.Fetch()
	assert.Equal(t, 0, v)
	assert.True(t, ok)

	o = opt.MakeWithErr(100, fmt.Errorf("test"))
	assert.False(t, o.IsDefined())
	assert.Panics(t, func() { o.Get() })

	v, ok = o.Fetch()
	assert.Equal(t, 0, v)
	assert.False(t, ok)
}

func TestOptionMap(t *testing.T) {
	s := goscala.Some[int](100)

	s1 := opt.Map[int, string](s)(strconv.Itoa)
	assert.True(t, s1.IsDefined())
	assert.Equal(t, "100", s1.Get())

	n := goscala.None[int]()
	s1 = opt.Map[int, string](n)(strconv.Itoa)
	assert.False(t, s1.IsDefined())
}


func TestOptionFlatMap(t *testing.T) {
	s := goscala.Some[int](100)

	f := monad.FuncAndThen[int, string, goscala.Option[string]](strconv.Itoa)(goscala.Some[string])
	
	s1 := opt.FlatMap[int, string](s)(f)
	assert.Equal(t, true, s1.IsDefined())
	assert.Equal(t, "100", s1.Get())

	n := goscala.None[int]()
	s1 = opt.FlatMap[int, string](n)(f)
	assert.Equal(t, false, s1.IsDefined())
}

func TestOptionFold(t *testing.T) {
	z := "zero"

	assert.Equal(t, "100", opt.Fold[int, string](goscala.Some[int](100))(z)(strconv.Itoa))
	assert.Equal(t, "zero", opt.Fold[int, string](goscala.None[int]())(z)(strconv.Itoa))
}

func TestOptionCollect(t *testing.T) {
	p := func(v int) (s string, ok bool) {
		if ok = (v != 0); ok {
			s = strconv.Itoa(v)
		}
		return
	}

	o := opt.MakeWithBool(0, false)
	ans := opt.Collect[int, string](o)(p)
	assert.Equal(t, false, ans.IsDefined())

	o = opt.MakeWithBool(100, true)
	ans = opt.Collect[int, string](o)(p)
	assert.Equal(t, true, ans.IsDefined())
	assert.Equal(t, "100", ans.Get())

	o = goscala.Some(0)
	ans = opt.Collect[int, string](o)(p)
	assert.Equal(t, false, ans.IsDefined())
}

func TestOptionToLeft(t *testing.T) {
	v1 := 1
	v2 := "abc"
	o := goscala.Some[int](v1)

	e := opt.Left[int, string](o)(v2)
	assert.True(t, e.IsLeft())
	assert.Equal(t, v1, e.Left())

	o = goscala.None[int]()
	e = opt.Left[int, string](o)(v2)
	assert.True(t, e.IsRight())
	assert.Equal(t, v2, e.Right())
}

func TestOptionRight(t *testing.T) {
	v1 := 1
	v2 := "abc"
	o := goscala.Some[int](v1)

	e := opt.Right[string, int](o)(v2)
	assert.True(t, e.IsRight())
	assert.Equal(t, v1, e.Right())

	o = goscala.None[int]()
	e = opt.Right[string, int](o)(v2)
	assert.True(t, e.IsLeft())
	assert.Equal(t, v2, e.Left())
}