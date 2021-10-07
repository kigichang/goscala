package slices

import gs "github.com/kigichang/goscala"

type sliceIterator[T any] struct {
	src *Slice[T]
	cur int
}

var _ gs.Iterator[int] = &sliceIterator[int]{}

func (i *sliceIterator[T]) Len() int {
	return len(*i.src)
}

func (i *sliceIterator[T]) Cap() int {
	return cap(*i.src)
}

func (i *sliceIterator[T]) Next() (ok bool) {
	i.cur++
	if ok = (i.cur < len(*i.src)); !ok {
		i.cur = len(*i.src)
	}
	return
}

func (i *sliceIterator[T]) Get() T {
	a := *i.src
	return a[i.cur]
}

func newSliceIterator[T any](src *Slice[T]) *sliceIterator[T] {
	return &sliceIterator[T]{
		src: src,
		cur: -1,
	}
}
