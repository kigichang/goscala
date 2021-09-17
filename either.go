package goscala

import (
	"fmt"
)

// Either represents Either of Scala. Either means it may be one of two possible types L and R.
// Either has two subtypes Left and Right. Usually, Right means positive like ture and success.
// Left meas negative like false and failure oppositely.
type Either[L, R any] interface {
	fmt.Stringer

	Contains(R, EqualFunc[R]) bool
	Equals(Either[L, R], EqualFunc[L], EqualFunc[R]) bool

	IsLeft() bool
	IsRight() bool
	Left() L
	Right() R
	Get() R
	Option() Option[R]
	Exists(p Predict[R]) bool
	FilterOrElse(p Predict[R], z L) Either[L, R]
	Forall(p Predict[R]) bool
	Foreach(f func(R))
	GetOrElse(z R) R
	OrElse(z Either[L, R]) Either[L, R]
	Swap() Either[R, L]
	//Try() Try[R]
	Slice() Slice[R]
}

// either implements Either.
type either[L, R any] struct {
	right bool
	lv    L
	rv    R
}

func (e *either[L, R]) String() string {
	if e.right {
		return fmt.Sprintf(`Right(%v)`, e.rv)
	}
	return fmt.Sprintf(`Left(%v)`, e.lv)
}

// Contains returns true if e contains given value.
func (e *either[L, R]) Contains(z R, fn EqualFunc[R]) bool {
	if e.IsRight() {
		return fn(e.Right(), z)
	}
	return false
}

// Equals returns true if e is same as that.
func (e *either[L, R]) Equals(that Either[L, R], lf EqualFunc[L], rf EqualFunc[R]) bool {
	if e == that {
		return true
	}

	if e.IsRight() == that.IsRight() {
		if e.IsRight() {
			return rf(e.Right(), that.Right())
		}
		return lf(e.Left(), that.Left())
	}

	return false
}

// IsLeft returns true if e is Left subtype.
func (e *either[L, R]) IsLeft() bool {
	return !e.right
}

// Left returns left value if e is Left, or panic.
func (e *either[L, R]) Left() L {
	if e.IsLeft() {
		return e.lv
	}
	panic(fmt.Sprintf(`can not get left value from %v`, e))
}

// IsRight returns true if e is Right subtype.
func (e *either[L, R]) IsRight() bool {
	return e.right
}

// Right returns right value if e is Right, or panic.
func (e *either[L, R]) Right() R {
	if e.IsRight() {
		return e.rv
	}
	panic(fmt.Sprintf(`can not get right value from %v`, e))
}

// Get returns right value if e is Right, or panic.
func (e *either[L, R]) Get() R {
	return e.Right()
}

// Option converts to Option with type R.
// Returns Some of right value if e is Right,
// or returns None.
func (e *either[L, R]) Option() Option[R] {
	if e.IsRight() {
		return Some[R](e.rv)
	}
	return None[R]()
}

// Exists returns false if e is Left, or result applying predicate p to Right.
func (e *either[L, R]) Exists(p Predict[R]) bool {
	if e.right {
		return p(e.rv)
	}
	return false
}

// FilterOrElse returns existing Right if e is right and the given predicate p holds the right value,
// or returns Left contains value z if e is right and the given predicate p does not hold the right value,
// or return existing Left if e is Left.
//
// Right(12).filterOrElse(_ > 10, -1)   // Right(12)
// Right(7).filterOrElse(_ > 10, -1)    // Left(-1)
// Left(7).filterOrElse(_ => false, -1) // Left(7)
func (e *either[L, R]) FilterOrElse(p Predict[R], z L) Either[L, R] {
	if !e.right || p(e.rv) {
		return e
	}

	return Left[L, R](z)
}

// Forall returns true if e is Left, or result applying given predicate p to right value of e.
func (e *either[L, R]) Forall(p Predict[R]) bool {
	if e.right {
		return p(e.rv)
	}
	return true
}

// Foreach executes given function f if e is Right.
func (e *either[L, R]) Foreach(f func(R)) {
	if e.right {
		f(e.rv)
	}
}

// GetOrElse returns right value if e is Right, or given argument z if e is Left.
func (e *either[L, R]) GetOrElse(z R) R {
	if e.right {
		return e.rv
	}

	return z
}

// GetOrElse returns Right if e is Right, or given argument z if e is Left.
func (e *either[L, R]) OrElse(z Either[L, R]) Either[L, R] {
	if !e.right {
		return z
	}
	return e
}

// Swap returns new Either[R, L] swapping Left and Right position.
func (e *either[L, R]) Swap() Either[R, L] {
	if e.right {
		return Left[R, L](e.rv)
	}

	return Right[R, L](e.lv)
}

// Slice converts to Slice with type R.
// Slice contains one value if e is Right, or empty if e is Left.
func (e *either[L, R]) Slice() Slice[R] {
	if e.right {
		return SliceFrom(e.rv)
	}
	return SliceEmpty[R]()
}
