package goscala

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeOption(t *testing.T) {
	v := 0
	opt := MakeOption[int](v)
	fmt.Println(opt)
	assert.Equal(t, false, opt.IsDefined())
	assert.Panics(t, func() { opt.Get() })

	v = 1
	opt = MakeOption[int](v)
	assert.Equal(t, true, opt.IsDefined())
	assert.Equal(t, v, opt.Get())
}

func TestSome(t *testing.T) {
	v := 0
	opt := Some[int](v)
	fmt.Println(opt)
	assert.Equal(t, true, opt.IsDefined())
	assert.Equal(t, v, opt.Get())

	v = 1
	opt = Some[int](v)
	assert.Equal(t, true, opt.IsDefined())
	assert.Equal(t, v, opt.Get())
}

func TestNone(t *testing.T) {
	opt := None[int]()
	fmt.Println(opt)
	assert.Equal(t, false, opt.IsDefined())
	assert.Panics(t, func() { opt.Get() })
}

func TestOptionUnless(t *testing.T) {
	v := 0
	cond := func() bool { return true }
	opt := OptionUnless[int](cond, v)

	fmt.Println(opt)
	assert.Equal(t, false, opt.IsDefined())
	assert.Panics(t, func() { opt.Get() })

	cond = func() bool { return false }
	opt = OptionUnless[int](cond, v)

	fmt.Println(opt)
	assert.Equal(t, true, opt.IsDefined())
	assert.Equal(t, v, opt.Get())
}

func TestOptionWhen(t *testing.T) {
	v := 0
	cond := func() bool { return true }
	opt := OptionWhen[int](cond, v)

	fmt.Println(opt)
	assert.Equal(t, true, opt.IsDefined())
	assert.Equal(t, v, opt.Get())

	cond = func() bool { return false }
	opt = OptionWhen[int](cond, v)

	fmt.Println(opt)
	assert.Equal(t, false, opt.IsDefined())
	assert.Panics(t, func() { opt.Get() })
}

func TestOptionCollect(t *testing.T) {
	p := func(v int) (s string, ok bool) {
		if ok = (v != 0); ok {
			s = strconv.Itoa(v)
		}
		return
	}

	o := MakeOption[int](0)
	ans := OptionCollect[int, string](o, p)
	assert.Equal(t, false, ans.IsDefined())

	o = MakeOption[int](100)
	ans = OptionCollect[int, string](o, p)
	assert.Equal(t, true, ans.IsDefined())
	assert.Equal(t, "100", ans.Get())

	o = Some(0)
	ans = OptionCollect[int, string](o, p)
	assert.Equal(t, false, ans.IsDefined())
}

func TestOptionEquals(t *testing.T) {
	o := Some(100)

	assert.True(t, o.Equals(o, Equal[int]))
	assert.True(t, o.Equals(Some(100), Equal[int]))
	assert.False(t, o.Equals(Some(101), Equal[int]))
	assert.False(t, o.Equals(None[int](), Equal[int]))

	o = None[int]()
	assert.False(t, o.Equals(Some(0), Equal[int]))
	assert.False(t, o.Equals(Some(100), Equal[int]))
	assert.True(t, o.Equals(None[int](), Equal[int]))
}

func TestOptionContains(t *testing.T) {
	o := Some(100)
	assert.True(t, o.Contains(100, Equal[int]))
	assert.False(t, o.Contains(101, Equal[int]))

	o = None[int]()
	assert.False(t, o.Contains(0, Equal[int]))
	assert.False(t, o.Contains(100, Equal[int]))
}

func TestOptionExists(t *testing.T) {

	p := func(v int) bool {
		return v > 0
	}

	o := Some(1)
	assert.True(t, o.Exists(p))

	o = Some(-1)
	assert.False(t, o.Exists(p))

	o = None[int]()
	assert.False(t, o.Exists(p))
}

func TestOptionFilter(t *testing.T) {
	s := Some[int](100)

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

	n := None[int]()

	s1 = n.Filter(f1)
	assert.Equal(t, false, s1.IsDefined())

	s2 = n.Filter(f2)
	assert.Equal(t, false, s2.IsDefined())
}

func TestOptionFilterNot(t *testing.T) {
	s := Some[int](100)

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

	n := None[int]()

	s1 = n.FilterNot(f1)
	assert.Equal(t, false, s1.IsDefined())

	s2 = n.FilterNot(f2)
	assert.Equal(t, false, s2.IsDefined())
}

func TestOptionFlatMap(t *testing.T) {
	s := Some[int](100)

	f := func(v int) Option[string] {
		return Some[string](fmt.Sprintf("%d", v))
	}

	s1 := OptionFlatMap[int, string](s)(f)
	assert.Equal(t, true, s1.IsDefined())
	assert.Equal(t, "100", s1.Get())

	n := None[int]()
	s1 = OptionFlatMap[int, string](n)(f)
	assert.Equal(t, false, s1.IsDefined())
}

func TestOptionFold(t *testing.T) {
	f := func(v int) string {
		return fmt.Sprintf("%d", v)
	}

	z := "zero"

	assert.Equal(t, "100", OptionFold(Some[int](100), z, f))
	assert.Equal(t, "zero", OptionFold(None[int](), z, f))
}

func TestOptionForall(t *testing.T) {
	s := Some[int](100)

	f1 := func(v int) bool {
		return v == 100
	}

	f2 := func(v int) bool {
		return v < 0
	}

	assert.Equal(t, true, s.Forall(f1))
	assert.Equal(t, false, s.Forall(f2))

	n := None[int]()
	assert.Equal(t, true, n.Forall(f1))
	assert.Equal(t, true, n.Forall(f2))
}

func TestOptionForeach(t *testing.T) {
	sum := 123
	s := Some[int](100)
	f := func(v int) {
		sum += v
	}
	s.Foreach(f)
	assert.Equal(t, 123+100, sum)

	sum = 123
	n := None[int]()
	n.Foreach(f)
	assert.Equal(t, 123, sum)
}

func TestOptionGetOrElse(t *testing.T) {
	s := Some[int](100)
	assert.Equal(t, 100, s.GetOrElse(-1))

	n := None[int]()
	assert.Equal(t, -1, n.GetOrElse(-1))
}

func TestOptionMap(t *testing.T) {
	s := Some[int](100)
	f := func(v int) string {
		return fmt.Sprintf("%d", v)
	}

	s1 := OptionMap[int, string](s)(f)
	assert.Equal(t, true, s1.IsDefined())
	assert.Equal(t, "100", s1.Get())

	n := None[int]()
	s1 = OptionMap[int, string](n)(f)
	assert.Equal(t, false, s1.IsDefined())
}

func TestOptionOrElse(t *testing.T) {
	s := Some[int](100)
	f := Some[int](1)

	assert.Equal(t, s.IsDefined(), s.OrElse(f).IsDefined())
	assert.Equal(t, s.Get(), s.OrElse(f).Get())

	n := None[int]()
	assert.Equal(t, true, n.OrElse(f).IsDefined())
	assert.Equal(t, 1, n.OrElse(f).Get())
}

func TestOptionZip(t *testing.T) {
	s1 := Some[int](100)
	s2 := Some[string]("abc")

	s := OptionZip(s1, s2)

	assert.Equal(t, true, s.IsDefined())
	assert.Equal(t, 100, s.Get().V1())
	assert.Equal(t, "abc", s.Get().V2())

	s2 = None[string]()

	s = OptionZip(s1, s2)
	assert.False(t, s.IsDefined())
}

func TestOptionUnzip(t *testing.T) {
	v1 := 1
	v2 := 2
	o := Some[Tuple2[int, int]](MakeTuple2(v1, v2))

	o1, o2 := OptionUnzip(o)
	assert.True(t, o1.IsDefined())
	assert.Equal(t, v1, o1.Get())
	assert.True(t, o2.IsDefined())
	assert.Equal(t, v2, o2.Get())

	o = None[Tuple2[int, int]]()
	o1, o2 = OptionUnzip(o)
	assert.False(t, o1.IsDefined())
	assert.False(t, o2.IsDefined())
}

func TestOptionUnzip3(t *testing.T) {
	v1 := 1
	v2 := 2
	v3 := 3
	o := Some[Tuple3[int, int, int]](MakeTuple3(v1, v2, v3))

	o1, o2, o3 := OptionUnzip3(o)
	assert.True(t, o1.IsDefined())
	assert.Equal(t, v1, o1.Get())
	assert.True(t, o2.IsDefined())
	assert.Equal(t, v2, o2.Get())
	assert.True(t, o3.IsDefined())
	assert.Equal(t, v3, o3.Get())

	o = None[Tuple3[int, int, int]]()
	o1, o2, o3 = OptionUnzip3(o)
	assert.False(t, o1.IsDefined())
	assert.False(t, o2.IsDefined())
	assert.False(t, o3.IsDefined())
}

//func TestOptionToLeft(t *testing.T) {
//	v1 := 1
//	v2 := "abc"
//	o := Some[int](v1)
//
//	e := OptionToLeft[int, string](o, v2)
//	assert.True(t, e.IsLeft())
//	assert.Equal(t, v1, e.Left())
//
//	o = None[int]()
//	e = OptionToLeft[int, string](o, v2)
//	assert.True(t, e.IsRight())
//	assert.Equal(t, v2, e.Right())
//}
//
//func TestOptionRight(t *testing.T) {
//	v1 := 1
//	v2 := "abc"
//	o := Some[int](v1)
//
//	e := OptionToRight[string, int](o, v2)
//	assert.True(t, e.IsRight())
//	assert.Equal(t, v1, e.Right())
//
//	o = None[int]()
//	e = OptionToRight[string, int](o, v2)
//	assert.True(t, e.IsLeft())
//	assert.Equal(t, v2, e.Left())
//}
//
//func TestOptionTry(t *testing.T) {
//	assert.Equal(t, 1, Some(1).Try().Get())
//	assert.Equal(t, ErrEmpty, None[int]().Try().Failed())
//}

func TestOptionMapWithErr(t *testing.T) {
	assert.Equal(t, 1, OptionMapWithErr(Some("1"), strconv.Atoi).Get())
	assert.False(t, OptionMapWithErr(Some("abc"), strconv.Atoi).IsDefined())
	assert.False(t, OptionMapWithErr(None[string](), strconv.Atoi).IsDefined())
}

func TestOptionWithBool(t *testing.T) {
	f := func(s string) (int, bool) {
		v, err := strconv.Atoi(s)
		return v, err == nil
	}

	assert.Equal(t, 1, OptionMapWithBool(Some("1"), f).Get())
	assert.False(t, OptionMapWithBool(Some("abc"), f).IsDefined())
	assert.False(t, OptionMapWithBool(None[string](), f).IsDefined())
}
