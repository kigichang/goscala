package goscala

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeft(t *testing.T) {
	v := 0
	e := Left[int, string](v)
	fmt.Println(e)
	assert.Equal(t, true, e.IsLeft())
	assert.Equal(t, false, e.IsRight())
	assert.Equal(t, v, e.Left())
	assert.Panics(t, func() { e.Right() })
}

func TestRight(t *testing.T) {
	v := "hello"
	e := Right[int, string](v)
	fmt.Println(e)
	assert.Equal(t, false, e.IsLeft())
	assert.Equal(t, true, e.IsRight())
	assert.Equal(t, v, e.Right())
	assert.Panics(t, func() { e.Left() })
}

func TestMakeEither(t *testing.T) {
	left := 1
	right := "hello"

	e := MakeEither[int, string](left, right)
	fmt.Println(e)
	assert.Equal(t, false, e.IsLeft())
	assert.Equal(t, true, e.IsRight())
	assert.Equal(t, right, e.Right())
	assert.Panics(t, func() { e.Left() })

	left = 1
	right = ""
	e = MakeEither[int, string](left, right)
	fmt.Println(e)
	assert.Equal(t, true, e.IsLeft())
	assert.Equal(t, false, e.IsRight())
	assert.Equal(t, left, e.Left())
	assert.Panics(t, func() { e.Right() })

	assert.Panics(t, func() { MakeEither[int, string](0, "") })
}

func TestMakeEitherWithBool(t *testing.T) {
	e := MakeEitherWithBool[int](100, true)

	assert.Equal(t, true, e.IsRight())
	assert.Equal(t, 100, e.Get())

	e = MakeEitherWithBool[int](0, false)
	assert.Equal(t, false, e.IsRight())
	assert.Equal(t, true, e.IsLeft())
	assert.Panics(t, func() { e.Get() })
	assert.Equal(t, false, e.Swap().Get())
}

func TestEitherExists(t *testing.T) {
	/*
	   Right(12).exists(_ > 10)   // true
	   Right(7).exists(_ > 10)    // false
	   Left(12).exists(_ => true) // false
	*/

	e := Right[int, int](12)
	p1 := func(v int) bool {
		return v > 10
	}

	p2 := func(v int) bool {
		return true
	}

	assert.Equal(t, true, e.Exists(p1))

	e = Right[int, int](7)
	assert.Equal(t, false, e.Exists(p1))

	e = Left[int, int](12)
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

	e := Right[int, int](12).FilterOrElse(p1, -1)
	assert.Equal(t, true, e.IsRight())
	assert.Equal(t, 12, e.Right())

	e = Right[int, int](7).FilterOrElse(p1, -1)
	assert.Equal(t, true, e.IsLeft())
	assert.Equal(t, -1, e.Left())

	e = Left[int, int](7).FilterOrElse(p2, -1)
	assert.Equal(t, true, e.IsLeft())
	assert.Equal(t, 7, e.Left())
}

func TestEitherFlatMap(t *testing.T) {
	f := func(v string) Either[int, int64] {
		a, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return Left[int, int64](0)
		}
		return Right[int, int64](a)
	}

	r := Right[int, string]("1000")

	e := EitherFlatMap(r, f)

	assert.Equal(t, true, e.IsRight())
	assert.Equal(t, int64(1000), e.Right())

	r = Right[int, string]("abc")
	e = EitherFlatMap(r, f)
	assert.Equal(t, false, e.IsRight())
	assert.Equal(t, true, e.IsLeft())
	assert.Equal(t, 0, e.Left())

	l := Left[int, string](100)
	e = EitherFlatMap(l, f)

	assert.Equal(t, true, e.IsLeft())
	assert.Equal(t, l.Left(), e.Left())
}

func TestEitherFold(t *testing.T) {
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

	r := Right[int, string]("1000")
	e := EitherFold(r, fa, fb)
	assert.Equal(t, int64(1000*10), e)

	r = Right[int, string]("abc")
	e = EitherFold(r, fa, fb)
	assert.Equal(t, int64(-1), e)

	l := Left[int, string](100)
	e = EitherFold(l, fa, fb)
	assert.Equal(t, int64(100+10), e)
}

func TestEitherForall(t *testing.T) {
	p1 := func(v int) bool {
		return v > 10
	}

	p2 := func(_ int) bool {
		return false
	}

	assert.Equal(t, true, Right[int, int](12).Forall(p1))
	assert.Equal(t, false, Right[int, int](7).Forall(p1))
	assert.Equal(t, true, Left[int, int](0).Forall(p2))
}

func TestEitherForeach(t *testing.T) {
	r := Right[int, string]("right")
	str := "hello"
	r.Foreach(func(s string) {
		str += s
	})
	assert.Equal(t, "helloright", str)

	l := Left[int, string](100)
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
	assert.Equal(t, 12, Right[int, int](12).GetOrElse(17))
	assert.Equal(t, 17, Left[int, int](12).GetOrElse(17))

}

func TestEitherMap(t *testing.T) {
	/*
	   Right(12).map(x => "flower") // Result: Right("flower")
	   Left(12).map(x => "flower")  // Result: Left(12)
	*/
	f := func(_ int) string {
		return "flower"
	}

	e := EitherMap(Right[int, int](12), f)
	assert.Equal(t, true, e.IsRight())
	assert.Equal(t, "flower", e.Right())

	e = EitherMap(Left[int, int](12), f)
	assert.Equal(t, true, e.IsLeft())
	assert.Equal(t, 12, e.Left())
}

func TestEigherOrElse(t *testing.T) {
	/*
	   Right(1) orElse Left(2) // Right(1)
	   Left(1) orElse Left(2)  // Left(2)
	   Left(1) orElse Left(2) orElse Right(3) // Right(3)
	*/

	f := Left[int, int](2)

	f1 := Right[int, int](3)

	e := Right[int, int](1).OrElse(f)
	assert.Equal(t, true, e.IsRight())
	assert.Equal(t, 1, e.Right())

	e = Left[int, int](1).OrElse(f)
	assert.Equal(t, true, e.IsLeft())
	assert.Equal(t, 2, e.Left())

	e = Left[int, int](1).OrElse(f).OrElse(f1)
	assert.Equal(t, true, e.IsRight())
	assert.Equal(t, false, e.IsLeft())
	assert.Equal(t, 3, e.Right())
	assert.Panics(t, func() { e.Left() })
}

func TestEitherGet(t *testing.T) {
	assert.Equal(t, 1, Right[int, int](1).Get())
	assert.Panics(t, func() { Left[int, int](100).Get() })
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
	left := Left[string, int]("left")
	right := left.Swap()
	assert.Equal(t, true, right.IsRight())
	assert.Equal(t, "left", right.Right())
	assert.Panics(t, func() { right.Left() })

	r1 := Right[string, int](2)
	r2 := Left[int, string](3).Swap()
	result := EitherFlatMap(r1, func(v1 int) Either[string, int] {
		return EitherMap(r2, func(v2 int) int {
			return v1 * v2
		})
	})

	assert.Equal(t, true, result.IsRight())
	assert.Equal(t, 2*3, result.Get())

}

func TestEitherOption(t *testing.T) {
	/*
	   Right(12).toOption // Some(12)
	   Left(12).toOption  // None
	*/
	o := Right[int, int](12).Option()
	assert.Equal(t, true, o.IsDefined())
	assert.Equal(t, 12, o.Get())

	o = Left[string, int]("12").Option()
	assert.Equal(t, false, o.IsDefined())
	assert.Panics(t, func() { o.Get() })
}

func TestEitherCond(t *testing.T) {
	lv := 100
	rv := "abc"

	e := EitherCond(func() bool { return true }, lv, rv)
	assert.True(t, e.IsRight())
	assert.False(t, e.IsLeft())
	assert.Equal(t, rv, e.Right())

	e = EitherCond(func() bool { return false }, lv, rv)
	assert.False(t, e.IsRight())
	assert.True(t, e.IsLeft())
	assert.Equal(t, lv, e.Left())

}

func TestEitherContains(t *testing.T) {
	/*
	   // Returns true because value of Right is "something" which equals "something".
	   Right("something") contains "something"

	   // Returns false because value of Right is "something" which does not equal "anything".
	   Right("something") contains "anything"

	   // Returns false because it's not a Right value.
	   Left("something") contains "something"
	*/

	assert.True(t, Right[string, string]("something").Contains("something", Equal[string]))
	assert.False(t, Right[string, string]("something").Contains("anything", Equal[string]))
	assert.False(t, Left[string, string]("something").Contains("something", Equal[string]))
}

func TestEitherEquals(t *testing.T) {

	assert.True(t, Right[int, string]("something").Equals(Right[int, string]("something"), Equal[int], Equal[string]))
	assert.False(t, Right[int, string]("something").Equals(Right[int, string]("anything"), Equal[int], Equal[string]))
	assert.True(t, Left[int, string](0).Equals(Left[int, string](0), Equal[int], Equal[string]))
	assert.False(t, Left[int, string](0).Equals(Left[int, string](100), Equal[int], Equal[string]))

	assert.False(t, Right[int, string]("something").Equals(Left[int, string](100), Equal[int], Equal[string]))
}
