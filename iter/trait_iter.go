package iter

import (
	gs "github.com/kigichang/goscala"
)

type TraitIterator[T any] struct {
	len  func() int
	cap  func() int
	next func() bool
	get  func() T
}

var _ gs.Iterator[int] = &TraitIterator[int]{}

func (a *TraitIterator[T]) Len() int {
	return a.len()
}

func (a *TraitIterator[T]) Cap() int {
	return a.cap()
}

func (a *TraitIterator[T]) Next() bool {
	return a.next()
}

func (a *TraitIterator[T]) Get() T {
	return a.get()
}

func Gen[T any](s ...T) gs.Iterator[T] {
	idx := -1
	ss := &s
	return &TraitIterator[T]{
		len: func() int {
			return len(*ss)
		},
		cap: func() int {
			return cap(*ss)
		},
		next: func() (ok bool) {
			idx++
			if ok = (idx < len(*ss)); !ok {
				idx = len(*ss)
			}
			return
		},
		get: func() T {
			a := *ss
			return a[idx]
		},
	}
}
