package grp

type Numeric interface {
	int | int8
}

type A [T any] struct {
	v T
}

func (a *A[T]) Map[U any](f func(T) U) {
	
}

func Add[T Numeric](a, b T) T {
	return a + b
}