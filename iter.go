package goscala

type Iter[T any] interface {
	Next() (int, T, bool)
}

type MapIter[K comparable, V any] interface {
	Next() (K, V, bool)
}

type iter[T any] struct {
	s *Slice[T]
	idx int
}

func (i *iter[T]) Next() (idx int, ret T, ok bool) {
	i.idx++

	ss := *(i.s)
	if i.idx < len(ss) {
		idx = i.idx
		ret = ss[idx]
		ok = true
	}
	return
}

func newIter[T any](s *Slice[T]) Iter[T] {
	return &iter[T] {
		s: s,
		idx: -1,
	}
}