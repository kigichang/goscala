package iter

import gs "github.com/kigichang/goscala"

func FromSlice[T any](s ...T) gs.Iterator[T] {
	idx := -1
	ss := &s
	return &TraitIterator[T]{
		FnLen: func() int {
			return len(*ss)
		},
		FnCap: func() int {
			return cap(*ss)
		},
		FnNext: func() (ok bool) {
			idx++
			if ok = (idx < len(*ss)); !ok {
				idx = len(*ss)
			}
			return
		},
		FnGet: func() T {
			a := *ss
			return a[idx]
		},
	}
}

func Map[T, U any](a gs.Iterator[T], fn func(T) U) gs.Iterator[U] {
	return &TraitIterator[U]{
		FnLen:  a.Len,
		FnCap:  a.Cap,
		FnNext: a.Next,
		FnGet: func() U {
			return fn(a.Get())
		},
	}
}

func FlatMap[T, U any](a gs.Iterator[T], fn func(T) gs.Iterator[U]) gs.Iterator[U] {
	var cur gs.Iterator[U] = &TraitIterator[U]{
		FnLen: func() int {
			return 0
		},
		FnCap: func() int {
			return 0
		},
		FnNext: func() bool {
			return false
		},
		FnGet: func() (ret U) {
			return
		},
	}
	return &TraitIterator[U]{
		FnLen: a.Len,
		FnCap: a.Cap,
		FnNext: func() bool {
			if !cur.Next() {
				if !a.Next() {
					return false
				}
				cur = fn(a.Get())
				return cur.Next()
			}
			return true
		},
		FnGet: func() U {
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

func FoldLeft[T, U any](a gs.Iterator[T], z U, fn func(U, T) U) U {
	zz := z

	for a.Next() {
		zz = fn(zz, a.Get())
	}
	return zz
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
	return FromSlice[T](dst...)
}

func FoldRight[T, U any](a gs.Iterator[T], z U, fn func(T, U) U) U {
	iter := Reverse(a)
	return FoldLeft(iter, z, func(a U, b T) U {
		return fn(b, a)
	})
}

func ScanLeft[T, U any](a gs.Iterator[T], z U, fn func(U, T) U) gs.Iterator[U] {
	zz := z
	first := true
	return &TraitIterator[U]{
		FnLen: func() int {
			return a.Len() + 1
		},
		FnCap: func() int {
			return a.Cap() + 1
		},
		FnNext: func() (ok bool) {
			return first || a.Next()
		},
		FnGet: func() U {
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
		FnLen: func() int {
			// length unknown
			return 0
		},
		FnCap: a.Cap,
		FnNext: func() bool {
			for a.Next() {
				if p(a.Get()) {
					return true
				}
			}
			return false
		},
		FnGet: func() T {
			return a.Get()
		},
	}
}

func FilterNot[T any](a gs.Iterator[T], p func(T) bool) gs.Iterator[T] {
	return Filter(a, func(v T) bool {
		return !p(v)
	})
}

func Equals[T any](a gs.Iterator[T], eq func(T, T) bool) func(gs.Iterator[T]) bool {
	return func(that gs.Iterator[T]) bool {
		if a == that {
			return true
		}

		for a.Next() {
			if !that.Next() || !eq(a.Get(), that.Get()) {
				return false
			}
		}

		return true
	}
}

func Find[T any](a gs.Iterator[T], p func(T) bool) (ret T, ok bool) {
	for a.Next() {
		v := a.Get()
		if ok = p(v); ok {
			ret = v
			return
		}
	}
	return
}

func Exists[T any](a gs.Iterator[T], p func(T) bool) (ok bool) {
	_, ok = Find(a, p)
	return
}

func Contains[T any](a gs.Iterator[T], eq func(T, T) bool) func(T) bool {
	return func(that T) bool {
		return Exists(a, gs.Currying2(eq)(that))
	}
}
