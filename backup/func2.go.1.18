package goscala

// Func2[T1, T2, R any] represents function: T1, T2 => R
type Func2[T1, T2, R any] func(T1, T2) R
type Currying2[T1, T2, R any] func(T1) func(T2) R

func Curried2[T1, T2, R any](fn Func2[T1, T2, R]) Currying2[T1, T2, R] {
	return func(a T1) func(T2) R {
		return func(b T2) R {
			return fn(a, b)
		}
	}

}

func (f Func2[T1, T2, R]) String() string {
	return TypeStr(f)
}

type EqualFunc[T1 any] Func2[T1, T1, bool]
type CompareFunc[T1 any] Func2[T1, T1, int]

// Func2Bool[T1, T2, R any] represents function: T1, T2 => (R, bool)
type Func2Bool[T1, T2, R any] func(T1, T2) (R, bool)

func (f Func2Bool[T1, T2, R]) String() string {
	return TypeStr(f)
}

// Func2Err[T1, R any] represents function: T1, T2 => (R, error)
type Func2Err[T1, T2, R any] func(T1, T2) (R, error)

func (f Func2Err[T1, T2, R]) String() string {
	return TypeStr(f)
}
