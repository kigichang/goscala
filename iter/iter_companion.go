package iter

import (
	gs "github.com/kigichang/goscala"
)

func Map[T, U any](a gs.Iterator[T], fn func(T) U) gs.Iterator[U] {
	return &TraitIterator[U]{
		len:  a.Len,
		cap:  a.Cap,
		next: a.Next,
		get: func() U {
			return fn(a.Get())
		},
	}
}

func FlatMap[T, U any](a gs.Iterator[T], fn func(T) gs.Iterator[U]) gs.Iterator[U] {
	var cur gs.Iterator[U] = &TraitIterator[U]{
		len: func() int {
			return 0
		},
		cap: func() int {
			return 0
		},
		next: func() bool {
			return false
		},
		get: func() (ret U) {
			return
		},
	}
	return &TraitIterator[U]{
		len: func() int {
			// length unknown
			return 0
		},
		cap: func() int {
			// capacity unknown
			return 0
		},
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

func Slice[T any](a gs.Iterator[T]) []T {
	ret := make([]T, 0, a.Cap())
	for a.Next() {
		ret = append(ret, a.Get())
	}
	return ret
}

func Reverse[T any](a gs.Iterator[T]) gs.Iterator[T] {
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

func FoldLeft[T, U any](a gs.Iterator[T], z U, fn func(U, T) U) U {
	zz := z

	for a.Next() {
		zz = fn(zz, a.Get())
	}
	return zz
}

func FoldRight[T, U any](a gs.Iterator[T], z U, fn func(T, U) U) U {
	it := Reverse(a)
	return FoldLeft(it, z, func(a U, b T) U {
		return fn(b, a)
	})
}

func ScanLeft[T, U any](a gs.Iterator[T], z U, fn func(U, T) U) gs.Iterator[U] {
	zz := z
	first := true
	return &TraitIterator[U]{
		len: func() int {
			return a.Len() + 1
		},
		cap: func() int {
			return a.Cap() + 1
		},
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

func ScanRight[T, U any](a gs.Iterator[T], z U, fn func(T, U) U) gs.Iterator[U] {
	it := Reverse(a)
	return Reverse(ScanLeft(it, z, func(a U, b T) U {
		return fn(b, a)
	}))
}

func Forall[T any](a gs.Iterator[T], fn func(T) bool) bool {
	for a.Next() {
		if !fn(a.Get()) {
			return false
		}
	}
	return true
}

func Foreach[T any](a gs.Iterator[T], fn func(T)) {
	for a.Next() {
		fn(a.Get())
	}
}

func Filter[T any](a gs.Iterator[T], p func(T) bool) gs.Iterator[T] {
	return &TraitIterator[T]{
		len: func() int {
			// length unknown
			return 0
		},
		cap: func() int {
			// capacity unknown
			return 0
		},
		next: func() bool {
			for a.Next() {
				if p(a.Get()) {
					return true
				}
			}
			return false
		},
		get: func() T {
			return a.Get()
		},
	}
}

func FilterNot[T any](a gs.Iterator[T], p func(T) bool) gs.Iterator[T] {
	return Filter(a, func(v T) bool {
		return !p(v)
	})
}
