package goscala_test

import (
	"testing"

	"github.com/kigichang/goscala"
	"github.com/stretchr/testify/assert"
)

func TestLeft(t *testing.T) {
	v := 0
	e := goscala.Left[int, string](v)
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
	e := goscala.Right[int, string](v)
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

func TestEitherOption(t *testing.T) {
	/*
	   Right(12).toOption // Some(12)
	   Left(12).toOption  // None
	*/
	o := goscala.Right[int, int](12).Option()
	assert.Equal(t, true, o.IsDefined())
	assert.Equal(t, 12, o.Get())

	o = goscala.Left[string, int]("12").Option()
	assert.Equal(t, false, o.IsDefined())
	assert.Panics(t, func() { o.Get() })
}

func TestEitherExists(t *testing.T) {
	/*
	   Right(12).exists(_ > 10)   // true
	   Right(7).exists(_ > 10)    // false
	   Left(12).exists(_ => true) // false
	*/

	e := goscala.Right[int, int](12)
	p1 := func(v int) bool {
		return v > 10
	}

	p2 := func(v int) bool {
		return true
	}

	assert.Equal(t, true, e.Exists(p1))

	e = goscala.Right[int, int](7)
	assert.Equal(t, false, e.Exists(p1))

	e = goscala.Left[int, int](12)
	assert.Equal(t, false, e.Exists(p1))
	assert.Equal(t, false, e.Exists(p2))
}

func TestEitherFilterOrElse(t *testing.T) {
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

	e := goscala.Right[int, int](12).FilterOrElse(p1, -1)
	assert.Equal(t, true, e.IsRight())
	assert.Equal(t, 12, e.Right())

	e = goscala.Right[int, int](7).FilterOrElse(p1, -1)
	assert.Equal(t, true, e.IsLeft())
	assert.Equal(t, -1, e.Left())

	e = goscala.Left[int, int](7).FilterOrElse(p2, -1)
	assert.Equal(t, true, e.IsLeft())
	assert.Equal(t, 7, e.Left())
}

func TestEitherForall(t *testing.T) {
	p1 := func(v int) bool {
		return v > 10
	}

	p2 := func(_ int) bool {
		return false
	}

	assert.Equal(t, true, goscala.Right[int, int](12).Forall(p1))
	assert.Equal(t, false, goscala.Right[int, int](7).Forall(p1))
	assert.Equal(t, true, goscala.Left[int, int](0).Forall(p2))
}


func TestEitherForeach(t *testing.T) {
	r := goscala.Right[int, string]("right")
	str := "hello"
	r.Foreach(func(s string) {
		str += s
	})
	assert.Equal(t, "helloright", str)

	l := goscala.Left[int, string](100)
	str = "hello"
	l.Foreach(func(s string) {
		str += s
	})
	assert.Equal(t, "hello", str)
}

func TestEitherGetOrElse(t *testing.T) {
	/*
	   Right(12).getOrElse(17) // 12
	   Left(12).getOrElse(17)  // 17
	*/
	assert.Equal(t, 12, goscala.Right[int, int](12).GetOrElse(17))
	assert.Equal(t, 17, goscala.Left[int, int](12).GetOrElse(17))
}



func TestEitherSWap(t *testing.T) {
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
	left := goscala.Left[string, int]("left")
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