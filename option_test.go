package goscala_test

import (
	"fmt"
	"testing"

	gs "github.com/kigichang/goscala"
	"github.com/kigichang/goscala/opt"
	"github.com/stretchr/testify/assert"
)

func TestSome(t *testing.T) {
	v := 0
	o := gs.Some[int](v)
	t.Log(o)
	assert.Equal(t, true, o.IsDefined())
	assert.Equal(t, v, o.Get())

	x, ok := o.Fetch()
	assert.Equal(t, v, x)
	assert.True(t, ok)

	v = 1
	o = gs.Some[int](v)
	assert.Equal(t, true, o.IsDefined())
	assert.Equal(t, v, o.Get())

	x, ok = o.Fetch()
	assert.Equal(t, v, x)
	assert.True(t, ok)
}

func TestNone(t *testing.T) {
	o := gs.None[int]()
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

func TestOptionContains(t *testing.T) {
	o := gs.Some(100)
	assert.True(t, o.Contains(100, gs.Eq[int]))
	assert.False(t, o.Contains(101, gs.Eq[int]))

	o = gs.None[int]()
	assert.False(t, o.Contains(0, gs.Eq[int]))
	assert.False(t, o.Contains(100, gs.Eq[int]))
}

func TestOptionExists(t *testing.T) {

	p := func(v int) bool {
		return v > 0
	}

	o := gs.Some(1)
	assert.True(t, o.Exists(p))

	o = gs.Some(-1)
	assert.False(t, o.Exists(p))

	o = gs.None[int]()
	assert.False(t, o.Exists(p))
}

func TestOptionEquals(t *testing.T) {
	o := gs.Some(100)
	eq := o.Equals(gs.Eq[int])

	assert.True(t, eq(o))
	assert.True(t, eq(gs.Some(100)))
	assert.False(t, eq(gs.Some(101)))
	assert.False(t, eq(gs.None[int]()))

	o = gs.None[int]()
	eq = o.Equals(gs.Eq[int])
	assert.False(t, eq(gs.Some(0)))
	assert.False(t, eq(gs.Some(100)))
	assert.True(t, eq(gs.None[int]()))
}

func TestOptionFilter(t *testing.T) {
	s := gs.Some[int](100)

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

	n := gs.None[int]()

	s1 = n.Filter(f1)
	assert.Equal(t, false, s1.IsDefined())

	s2 = n.Filter(f2)
	assert.Equal(t, false, s2.IsDefined())
}

func TestOptionFilterNot(t *testing.T) {
	s := gs.Some[int](100)

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

	n := gs.None[int]()

	s1 = n.FilterNot(f1)
	assert.Equal(t, false, s1.IsDefined())

	s2 = n.FilterNot(f2)
	assert.Equal(t, false, s2.IsDefined())
}

func TestOptionForall(t *testing.T) {
	s := gs.Some[int](100)

	f1 := func(v int) bool {
		return v == 100
	}

	f2 := func(v int) bool {
		return v < 0
	}

	assert.Equal(t, true, s.Forall(f1))
	assert.Equal(t, false, s.Forall(f2))

	n := gs.None[int]()
	assert.Equal(t, true, n.Forall(f1))
	assert.Equal(t, true, n.Forall(f2))
}

func TestOptionForeach(t *testing.T) {
	sum := 123
	s := gs.Some[int](100)
	f := func(v int) {
		sum += v
	}
	s.Foreach(f)
	assert.Equal(t, 123+100, sum)

	sum = 123
	n := gs.None[int]()
	n.Foreach(f)
	assert.Equal(t, 123, sum)
}

func TestOptionGetOrElse(t *testing.T) {
	s := gs.Some[int](100)
	assert.Equal(t, 100, s.GetOrElse(-1))

	n := gs.None[int]()
	assert.Equal(t, -1, n.GetOrElse(-1))
}

func TestOptionOrElse(t *testing.T) {
	s := gs.Some[int](100)
	f := gs.Some[int](1)

	assert.Equal(t, s.IsDefined(), s.OrElse(f).IsDefined())
	assert.Equal(t, s.Get(), s.OrElse(f).Get())

	n := gs.None[int]()
	assert.Equal(t, true, n.OrElse(f).IsDefined())
	assert.Equal(t, 1, n.OrElse(f).Get())
}
