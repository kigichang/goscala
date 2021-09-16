package goscala

// Func2[T1, T2, R any] represents function: T1, T2 => R
type Func2[T1, T2, R any] func(T1, T2) R

type EqualFunc[T1 any] Func2[T1, T1, bool]
type CompareFunc[T1 any] Func2[T1, T1, int]

func (f Func2[T1, T2, R]) String() string {
	return typstr(f)
}

// Func2Bool[T1, T2, R any] represents function: T1, T2 => (R, bool)
type Func2Bool[T1, T2, R any] func(T1, T2) (R, bool)

func (f Func2Bool[T1, T2, R]) String() string {
	return typstr(f)
}

// Func2Err[T1, R any] represents function: T1, T2 => (R, error)
type Func2Err[T1, T2, R any] func(T1, T2) (R, error)

func (f Func2Err[T1, T2, R]) String() string {
	return typstr(f)
}
