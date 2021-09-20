package goscala

// Func3[T1, T2, R any] represents function: T1, T2, T3 => R
type Func3[T1, T2, T3, R any] func(T1, T2, T3) R
type Currying3[T1, T2, T3, R any] func(T1) func(T2) func(T3) R

func Curried3[T1, T2, T3, R any](fn Func3[T1, T2, T3, R]) Currying3[T1, T2, T3, R] {
	return func(a T1) func(T2) func(T3) R {
		return func(b T2) func(T3) R {
			return func(c T3) R {
				return fn(a, b, c)
			}
		}
	}

}

func (f Func3[T1, T2, T3, R]) String() string {
	return typstr(f)
}

func (f Func3[T1, T2, T3, R]) Curried() Currying3[T1, T2, T3, R] {
	return Curried3(f)
}

// Func3Bool[T1, T2, T3, R any] represents function: T1, T2, T3 => (R, bool)
type Func3Bool[T1, T2, T3, R any] func(T1, T2, T3) (R, bool)

func (f Func3Bool[T1, T2, T3, R]) String() string {
	return typstr(f)
}

// Func3Err[T1, R any] represents function: T1, T2, T3 => (R, error)
type Func3Err[T1, T2, T3, R any] func(T1, T2, T3) (R, error)

func (f Func3Err[T1, T2, T3, R]) String() string {
	return typstr(f)
}
