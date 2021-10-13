// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package either_test

import (
	"strconv"
	"testing"

	gs "github.com/kigichang/goscala"
	"github.com/kigichang/goscala/either"
	"github.com/stretchr/testify/assert"
)

func TestLeft(t *testing.T) {
	v := 0
	e := either.Left[int, string](v)
	t.Log(e)
	assert.Equal(t, true, e.IsLeft())
	assert.Equal(t, false, e.IsRight())
	assert.Equal(t, v, e.Left())
	assert.Panics(t, func() { e.Right() })
	assert.Panics(t, func() { e.Get() })
	_, ok := e.Fetch()
	assert.False(t, ok)
}

func TestRight(t *testing.T) {
	v := "hello"
	e := either.Right[int, string](v)
	t.Log(e)
	assert.Equal(t, false, e.IsLeft())
	assert.Equal(t, true, e.IsRight())
	assert.Equal(t, v, e.Right())
	assert.Equal(t, v, e.Get())
	assert.Panics(t, func() { e.Left() })

	x, ok := e.Fetch()
	assert.Equal(t, v, x)
	assert.True(t, ok)
}

func TestExists(t *testing.T) {
	/*
	   Right(12).exists(_ > 10)   // true
	   Right(7).exists(_ > 10)    // false
	   Left(12).exists(_ => true) // false
	*/

	e := either.Right[int, int](12)
	p1 := func(v int) bool {
		return v > 10
	}

	p2 := func(v int) bool {
		return true
	}

	assert.Equal(t, true, e.Exists(p1))

	e = either.Right[int, int](7)
	assert.Equal(t, false, e.Exists(p1))

	e = either.Left[int, int](12)
	assert.Equal(t, false, e.Exists(p1))
	assert.Equal(t, false, e.Exists(p2))
}

func TestFilterOrElse(t *testing.T) {
	/*
	   Right(12).filterOrElse(_ > 10, -1)   // Right(12)
	   Right(7).filterOrElse(_ > 10, -1)    // Left(-1)
	   Left(7).filterOrElse(_ => false, -1) // Left(7)
	*/
	p1 := func(v int) bool {
		return v > 10
	}

	p2 := func(_ int) bool {
		return false
	}

	e := either.Right[int, int](12).FilterOrElse(p1, -1)
	assert.Equal(t, true, e.IsRight())
	assert.Equal(t, 12, e.Right())

	e = either.Right[int, int](7).FilterOrElse(p1, -1)
	assert.Equal(t, true, e.IsLeft())
	assert.Equal(t, -1, e.Left())

	e = either.Left[int, int](7).FilterOrElse(p2, -1)
	assert.Equal(t, true, e.IsLeft())
	assert.Equal(t, 7, e.Left())
}

func TestForall(t *testing.T) {
	p1 := func(v int) bool {
		return v > 10
	}

	p2 := func(_ int) bool {
		return false
	}

	assert.Equal(t, true, either.Right[int, int](12).Forall(p1))
	assert.Equal(t, false, either.Right[int, int](7).Forall(p1))
	assert.Equal(t, true, either.Left[int, int](0).Forall(p2))
}

func TestForeach(t *testing.T) {
	r := either.Right[int, string]("right")
	str := "hello"
	r.Foreach(func(s string) {
		str += s
	})
	assert.Equal(t, "helloright", str)

	l := either.Left[int, string](100)
	str = "hello"
	l.Foreach(func(s string) {
		str += s
	})
	assert.Equal(t, "hello", str)
}

func TestGetOrElse(t *testing.T) {
	/*
	   Right(12).getOrElse(17) // 12
	   Left(12).getOrElse(17)  // 17
	*/
	assert.Equal(t, 12, either.Right[int, int](12).GetOrElse(17))
	assert.Equal(t, 17, either.Left[int, int](12).GetOrElse(17))
}

func TestSWap(t *testing.T) {
	/*
	   val left: Either[String, Int]  = Left("left")
	   val right: Either[Int, String] = left.swap // Result: Right("left")
	   val right = Right(2)
	   val left  = Left(3)
	   for {
	     r1 <- right
	     r2 <- left.swap
	   } yield r1 * r2 // Right(6)
	*/
	left := either.Left[string, int]("left")
	right := left.Swap()
	assert.Equal(t, true, right.IsRight())
	assert.Equal(t, "left", right.Right())
	assert.Panics(t, func() { right.Left() })

	//r1 := Right[string, int](2)
	//r2 := Left[int, string](3).Swap()
	//result := EitherFlatMap(r1, func(v1 int) Either[string, int] {
	//	return EitherMap(r2, func(v2 int) int {
	//		return v1 * v2
	//	})
	//})
	//
	//assert.Equal(t, true, result.IsRight())
	//assert.Equal(t, 2*3, result.Get())

}

func TestBool(t *testing.T) {
	e := either.Bool(1, true)

	assert.True(t, e.IsRight())
	assert.Equal(t, 1, e.Right())

	e = either.Bool(1, false)
	assert.True(t, e.IsLeft())
	assert.Equal(t, false, e.Left())
}

func TestErr(t *testing.T) {
	e := either.Err(1, nil)
	assert.True(t, e.IsRight())
	assert.Equal(t, 1, e.Right())

	e = either.Err(1, gs.ErrEmpty)
	assert.True(t, e.IsLeft())
	assert.Equal(t, gs.ErrEmpty, e.Left())
}

func TestCond(t *testing.T) {
	lv := 100
	rv := "abc"

	e := either.Cond(func() bool { return true }, lv, rv)
	assert.True(t, e.IsRight())
	assert.False(t, e.IsLeft())
	assert.Equal(t, rv, e.Right())

	e = either.Cond(func() bool { return false }, lv, rv)
	assert.False(t, e.IsRight())
	assert.True(t, e.IsLeft())
	assert.Equal(t, lv, e.Left())
}

func TestFlatMap(t *testing.T) {
	f := func(v string) gs.Either[int, int64] {
		a, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return either.Left[int, int64](0)
		}
		return either.Right[int, int64](a)
	}

	r := either.Right[int, string]("1000")

	e := either.FlatMap(r, f)

	assert.Equal(t, true, e.IsRight())
	assert.Equal(t, int64(1000), e.Right())

	r = either.Right[int, string]("abc")
	e = either.FlatMap(r, f)
	assert.Equal(t, false, e.IsRight())
	assert.Equal(t, true, e.IsLeft())
	assert.Equal(t, 0, e.Left())

	l := either.Left[int, string](100)
	e = either.FlatMap(l, f)

	assert.Equal(t, true, e.IsLeft())
	assert.Equal(t, l.Left(), e.Left())
}

func TestFold(t *testing.T) {
	fa := func(v int) int64 {
		return int64(v + 10)
	}

	fb := func(s string) int64 {
		a, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return -1
		}
		return a * 10
	}

	r := either.Right[int, string]("1000")
	e := either.Fold(r, fa, fb)
	assert.Equal(t, int64(1000*10), e)

	r = either.Right[int, string]("abc")
	e = either.Fold(r, fa, fb)
	assert.Equal(t, int64(-1), e)

	l := either.Left[int, string](100)
	e = either.Fold(l, fa, fb)
	assert.Equal(t, int64(100+10), e)
}

func TestMap(t *testing.T) {
	/*
	   Right(12).map(x => "flower") // Result: Right("flower")
	   Left(12).map(x => "flower")  // Result: Left(12)
	*/
	f := func(_ int) string {
		return "flower"
	}

	e := either.Map(either.Right[int, int](12), f)
	assert.Equal(t, true, e.IsRight())
	assert.Equal(t, "flower", e.Right())

	e = either.Map(either.Left[int, int](12), f)
	assert.Equal(t, true, e.IsLeft())
	assert.Equal(t, 12, e.Left())
}
