package goscala

import "sync"

type LazyRef[T any] struct {
	once sync.Once
	fn   func() T
	v    T
}

func (z *LazyRef[T]) Get() T {
	z.once.Do(func() {
		z.v = z.fn()
	})
	return z.v
}

func Lazy[T any](fn func() T) *LazyRef[T] {
	return &LazyRef[T]{
		fn: fn,
	}
}

type Lazy2Ref[T1, T2 any] struct {
	once sync.Once
	fn   func() (T1, T2)
	v1   T1
	v2   T2
}

func (z *Lazy2Ref[T1, T2]) Get() (T1, T2) {
	z.once.Do(func() {
		z.v1, z.v2 = z.fn()
	})
	return z.v1, z.v2
}

func Lazy2[T1, T2 any](fn func() (T1, T2)) *Lazy2Ref[T1, T2] {
	return &Lazy2Ref[T1, T2]{
		fn: fn,
	}
}
