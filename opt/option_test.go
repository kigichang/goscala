// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package opt_test

import (
	"fmt"
	"strconv"
	"testing"

	gs "github.com/kigichang/goscala"
	"github.com/kigichang/goscala/opt"
	"github.com/stretchr/testify/assert"
)

func TestSome(t *testing.T) {
	v := 0
	o := opt.Some[int](v)
	t.Log(o)
	assert.Equal(t, true, o.IsDefined())
	assert.Equal(t, v, o.Get())

	x, ok := o.Fetch()
	assert.Equal(t, v, x)
	assert.True(t, ok)

	v = 1
	o = opt.Some[int](v)
	assert.Equal(t, true, o.IsDefined())
	assert.Equal(t, v, o.Get())

	x, ok = o.Fetch()
	assert.Equal(t, v, x)
	assert.True(t, ok)
}

func TestNone(t *testing.T) {
	o := opt.None[int]()
	t.Log(o)
	assert.Equal(t, false, o.IsDefined())
	assert.Panics(t, func() { o.Get() })
	x, ok := o.Fetch()
	assert.Equal(t, 0, x)
	assert.False(t, ok)

}

func TestMakeWithBool(t *testing.T) {
	o := opt.Bool(0, true)

	assert.True(t, o.IsDefined())
	assert.Equal(t, 0, o.Get())

	v, ok := o.Fetch()
	assert.Equal(t, 0, v)
	assert.True(t, ok)

	o = opt.Bool(100, false)
	assert.False(t, o.IsDefined())
	assert.Panics(t, func() { o.Get() })

	v, ok = o.Fetch()
	assert.Equal(t, 0, v)
	assert.False(t, ok)

}

func TestMakeWitErr(t *testing.T) {
	o := opt.Err(0, nil)

	assert.True(t, o.IsDefined())
	assert.Equal(t, 0, o.Get())

	v, ok := o.Fetch()
	assert.Equal(t, 0, v)
	assert.True(t, ok)

	o = opt.Err(100, fmt.Errorf("test"))
	assert.False(t, o.IsDefined())
	assert.Panics(t, func() { o.Get() })

	v, ok = o.Fetch()
	assert.Equal(t, 0, v)
	assert.False(t, ok)
}

func TestContains(t *testing.T) {
	o := opt.Some(100)
	assert.True(t, o.Contains(100, gs.Eq[int]))
	assert.False(t, o.Contains(101, gs.Eq[int]))

	o = opt.None[int]()
	assert.False(t, o.Contains(0, gs.Eq[int]))
	assert.False(t, o.Contains(100, gs.Eq[int]))
}

func TestExists(t *testing.T) {

	p := func(v int) bool {
		return v > 0
	}

	o := opt.Some(1)
	assert.True(t, o.Exists(p))

	o = opt.Some(-1)
	assert.False(t, o.Exists(p))

	o = opt.None[int]()
	assert.False(t, o.Exists(p))
}

func TestEquals(t *testing.T) {
	o := opt.Some(100)
	eq := o.Equals(gs.Eq[int])

	assert.True(t, eq(o))
	assert.True(t, eq(opt.Some(100)))
	assert.False(t, eq(opt.Some(101)))
	assert.False(t, eq(opt.None[int]()))

	o = opt.None[int]()
	eq = o.Equals(gs.Eq[int])
	assert.False(t, eq(opt.Some(0)))
	assert.False(t, eq(opt.Some(100)))
	assert.True(t, eq(opt.None[int]()))
}

func TestFilter(t *testing.T) {
	s := opt.Some[int](100)

	f1 := func(v int) bool {
		return v == 100
	}

	f2 := func(v int) bool {
		return v < 0
	}

	s1 := s.Filter(f1)
	assert.Equal(t, s.IsDefined(), s1.IsDefined())
	assert.Equal(t, s.Get(), s1.Get())

	s2 := s.Filter(f2)
	assert.Equal(t, false, s2.IsDefined())

	n := opt.None[int]()

	s1 = n.Filter(f1)
	assert.Equal(t, false, s1.IsDefined())

	s2 = n.Filter(f2)
	assert.Equal(t, false, s2.IsDefined())
}

func TestFilterNot(t *testing.T) {
	s := opt.Some[int](100)

	f1 := func(v int) bool {
		return v == 100
	}

	f2 := func(v int) bool {
		return v < 0
	}

	s1 := s.FilterNot(f1)
	assert.Equal(t, false, s1.IsDefined())

	s2 := s.FilterNot(f2)
	assert.Equal(t, s.IsDefined(), s2.IsDefined())
	assert.Equal(t, s.Get(), s2.Get())

	n := opt.None[int]()

	s1 = n.FilterNot(f1)
	assert.Equal(t, false, s1.IsDefined())

	s2 = n.FilterNot(f2)
	assert.Equal(t, false, s2.IsDefined())
}

func TestForall(t *testing.T) {
	s := opt.Some[int](100)

	f1 := func(v int) bool {
		return v == 100
	}

	f2 := func(v int) bool {
		return v < 0
	}

	assert.Equal(t, true, s.Forall(f1))
	assert.Equal(t, false, s.Forall(f2))

	n := opt.None[int]()
	assert.Equal(t, true, n.Forall(f1))
	assert.Equal(t, true, n.Forall(f2))
}

func TestForeach(t *testing.T) {
	sum := 123
	s := opt.Some[int](100)
	f := func(v int) {
		sum += v
	}
	s.Foreach(f)
	assert.Equal(t, 123+100, sum)

	sum = 123
	n := opt.None[int]()
	n.Foreach(f)
	assert.Equal(t, 123, sum)
}

func TestGetOrElse(t *testing.T) {
	s := opt.Some[int](100)
	assert.Equal(t, 100, s.GetOrElse(-1))

	n := opt.None[int]()
	assert.Equal(t, -1, n.GetOrElse(-1))
}

func TestOrElse(t *testing.T) {
	s := opt.Some[int](100)
	f := opt.Some[int](1)

	assert.Equal(t, s.IsDefined(), s.OrElse(f).IsDefined())
	assert.Equal(t, s.Get(), s.OrElse(f).Get())

	n := opt.None[int]()
	assert.Equal(t, true, n.OrElse(f).IsDefined())
	assert.Equal(t, 1, n.OrElse(f).Get())
}

func TestBool(t *testing.T) {
	o := opt.Bool(0, true)

	assert.True(t, o.IsDefined())
	assert.Equal(t, 0, o.Get())

	v, ok := o.Fetch()
	assert.Equal(t, 0, v)
	assert.True(t, ok)

	o = opt.Bool(100, false)
	assert.False(t, o.IsDefined())
	assert.Panics(t, func() { o.Get() })

	v, ok = o.Fetch()
	assert.Equal(t, 0, v)
	assert.False(t, ok)

}

func TestErr(t *testing.T) {
	o := opt.Err(0, nil)

	assert.True(t, o.IsDefined())
	assert.Equal(t, 0, o.Get())

	v, ok := o.Fetch()
	assert.Equal(t, 0, v)
	assert.True(t, ok)

	o = opt.Err(100, fmt.Errorf("test"))
	assert.False(t, o.IsDefined())
	assert.Panics(t, func() { o.Get() })

	v, ok = o.Fetch()
	assert.Equal(t, 0, v)
	assert.False(t, ok)
}

func TestMap(t *testing.T) {
	s := opt.Some[int](100)

	s1 := opt.Map(s, strconv.Itoa)
	assert.True(t, s1.IsDefined())
	assert.Equal(t, "100", s1.Get())

	n := opt.None[int]()
	s1 = opt.Map(n, strconv.Itoa)
	assert.False(t, s1.IsDefined())
}

func TestFlatMap(t *testing.T) {
	s := opt.Some[int](100)

	f := gs.FuncAndThen(strconv.Itoa, opt.Some[string])

	s1 := opt.FlatMap(s, f)
	assert.Equal(t, true, s1.IsDefined())
	assert.Equal(t, "100", s1.Get())

	n := opt.None[int]()
	s1 = opt.FlatMap(n, f)
	assert.Equal(t, false, s1.IsDefined())
}

func TestFold(t *testing.T) {
	z := "zero"

	assert.Equal(t, "100", opt.Fold[int, string](opt.Some[int](100))(z)(strconv.Itoa))
	assert.Equal(t, "zero", opt.Fold[int, string](opt.None[int]())(z)(strconv.Itoa))
}

func TestCollect(t *testing.T) {
	p := func(v int) (s string, ok bool) {
		if ok = (v != 0); ok {
			s = strconv.Itoa(v)
		}
		return
	}

	o := opt.Bool(0, false)
	ans := opt.Collect(o, p)
	assert.Equal(t, false, ans.IsDefined())

	o = opt.Bool(100, true)
	ans = opt.Collect(o, p)
	assert.Equal(t, true, ans.IsDefined())
	assert.Equal(t, "100", ans.Get())

	o = opt.Some(0)
	ans = opt.Collect(o, p)
	assert.Equal(t, false, ans.IsDefined())
}

func TestWhen(t *testing.T) {
	o := opt.When(gs.True, 0)
	assert.True(t, o.IsDefined())
	assert.Equal(t, 0, o.Get())

	o = opt.When(gs.False, 100)
	assert.False(t, o.IsDefined())
}

func TestUnless(t *testing.T) {
	o := opt.Unless(gs.True, 100)
	assert.False(t, o.IsDefined())

	o = opt.Unless(gs.False, 0)
	assert.True(t, o.IsDefined())
	assert.Equal(t, 0, o.Get())
}
