// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package iter

type Iter[T any] interface {
	Next() bool
	Get() T
}

type abstractIter[T any] struct {
	next func() bool
	get  func() T
}

func (a *abstractIter[T]) Next() bool {
	return a.next()
}

func (a *abstractIter[T]) Get() T {
	return a.get()
}

func newAbstractIter[T, U any](i Iter[T], fn func(T) U) Iter[U] {
	return &abstractIter[U]{
		next: i.Next,
		get: func() U {
			return fn(i.Get())
		},
	}
}

func GenIter[T any](s ...T) Iter[T] {
	idx := -1
	ss := &s
	return &abstractIter[T]{
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

func Map[T, U any](a Iter[T], fn func(T) U) Iter[U] {
	return newAbstractIter[T, U](a, fn)
}

func Slice[T any](a Iter[T]) []T {
	ret := []T{}
	for a.Next() {
		ret = append(ret, a.Get())
	}
	return ret
}

func FoldLeft[T, U any](a Iter[T], z U, fn func(U, T) U) U {
	zz := z

	for a.Next() {
		zz = fn(zz, a.Get())
	}
	return zz
}
