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

func Gen[T any](s ...T) Iter[T] {
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

func FlatMap[T, U any](a Iter[T], fn func(T) Iter[U]) Iter[U] {
	var cur Iter[U] = &abstractIter[U]{
		next: func() bool {
			return false
		},
		get: func() (ret U) {
			return
		},
	}
	return &abstractIter[U]{
		next: func() bool {
			if !cur.Next() {
				if !a.Next() {
					return false
				}
				cur = fn(a.Get())
				return cur.Next()
			}
			return true
		},
		get: func() U {
			return cur.Get()
		},
	}
}

func Slice[T any](a Iter[T]) []T {
	ret := []T{}
	for a.Next() {
		ret = append(ret, a.Get())
	}
	return ret
}

func Reverse[T any](a Iter[T]) Iter[T] {
	dst := Slice(a)
	size := len(dst)
	half := size / 2
	end := size - 1
	for i := 0; i < half; i++ {
		tmp := dst[i]
		dst[i] = dst[end-i]
		dst[end-i] = tmp
	}
	return Gen(dst...)
}

func FoldLeft[T, U any](a Iter[T], z U, fn func(U, T) U) U {
	zz := z

	for a.Next() {
		zz = fn(zz, a.Get())
	}
	return zz
}

func FoldRight[T, U any](a Iter[T], z U, fn func(T, U) U) U {
	iter := Reverse(a)
	return FoldLeft(iter, z, func(a U, b T) U {
		return fn(b, a)
	})
}

func ScanLeft[T, U any](a Iter[T], z U, fn func(U, T) U) Iter[U] {
	zz := z
	first := true
	return &abstractIter[U]{
		next: func() (ok bool) {
			return first || a.Next()
		},
		get: func() U {
			if first {
				first = false
				return zz
			}
			zz = fn(zz, a.Get())
			return zz
		},
	}
}

func ScanRight[T, U any](a Iter[T], z U, fn func(T, U) U) Iter[U] {
	it := Reverse(a)
	return Reverse(ScanLeft(it, z, func(a U, b T) U {
		return fn(b, a)
	}))
}
