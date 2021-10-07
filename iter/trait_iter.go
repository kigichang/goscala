package iter

import (
	gs "github.com/kigichang/goscala"
)

type TraitIterator[T any] struct {
	FnLen  func() int
	FnCap  func() int
	FnNext func() bool
	FnGet  func() T
}

var _ gs.Iterator[int] = &TraitIterator[int]{}

func (i *TraitIterator[T]) Len() int {
	return i.FnLen()
}

func (i *TraitIterator[T]) Cap() int {
	return i.FnCap()
}

func (i *TraitIterator[T]) Next() bool {
	return i.FnNext()
}

func (i *TraitIterator[T]) Get() T {
	return i.FnGet()
}
