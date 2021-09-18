package goscala

import (
	//"fmt"
	"sort"
)

func MakeSlice[T any](a int, b...int) Slice[T] {
	
	c := a
	if len(b) > 0 {
		c = b[0]
	}
	if c < a {
		c = a
	}

	return Slice[T](make([]T, a, c))
}

func SliceEmpty[T any]() Slice[T] {
	return []T{}
}

func SliceFrom[T any](a ...T) Slice[T] {
	return a
}

func SliceFill[T any](size int, v T) Slice[T] {
	ret := MakeSlice[T](size)
	for i := range ret {
		ret[i] = v
	}
	return ret
}

func SliceRange[T Number] (start, end, step T) Slice[T] {
	ret := SliceEmpty[T]()

	for i := start; i < end; i += step {
		ret = append(ret, i)
	}

	return ret
} 

func SliceTabulate[T any](size int, f Func1[int, T]) Slice[T] {
	if size <= 0 {
		return SliceEmpty[T]()
	}

	ret := MakeSlice[T](size)

	for i := 0; i < size; i++ {
		ret[i] = f(i)
	}

	return ret
}

func SliceCollect[T, U any](s Slice[T], pf PartialFunc[T, U]) Slice[U] {
	ret := SliceEmpty[U]()

	for i := range s {
		v, ok := pf(s[i])
		if ok {
			ret = append(ret, v)
		}
	}

	return ret
}

func SliceCollectFirst[T, U any](s Slice[T], pf PartialFunc[T, U]) Option[U] {
	for i := range s {
		v, ok := pf(s[i])
		if ok {
			return Some[U](v)
		}
	}
	return None[U]()
}

func SliceFlatMap[T, U any](s Slice[T], fn Func1[T, Sliceable[U]]) Slice[U] {
	ret := SliceEmpty[U]()
	for i := range s {
		v := fn(s[i]).Slice()
		ret = append(ret, v...)
	}
	return ret
}

func SliceFoldRight[T, U any](s Slice[T], z U, fn Func2[T, U, U]) U {
	size := s.Len()

	for i:= size - 1; i >= 0; i-- {
		z = fn(s[i], z)
	}

	return z
}

func SliceFoldLeft[T, U any](s Slice[T], z U, fn Func2[U, T, U]) U {
	for i := range s {
		z = fn(z, s[i])
	}
	return z
}

func SliceFold[T, U any](s Slice[T], z U, fn Func2[U, T, U]) U {
	return SliceFoldLeft(s, z, fn)
}

func SliceMap[T, U any](s Slice[T], fn Func1[T, U]) Slice[U] {
	ret := SliceEmpty[U]()

	for i := range s {
		ret = append(ret, fn(s[i]))
	}

	return ret
}

func SlicePartitionMap[T, A, B any](s Slice[T], fn Func1[T, Either[A, B]]) (Slice[A], Slice[B]) {
	a, b := SliceEmpty[A](), SliceEmpty[B]()

	for i := range s {
		v := fn(s[i])

		if v.IsRight() {
			b = append(b, v.Right())
		} else {
			a = append(a, v.Left())
		}
	}

	return a, b
}

func SliceScanLeft[T, U any](s Slice[T], z U, fn Func2[U, T, U]) Slice[U] {
	ret := MakeSlice[U](s.Len()+1)
	ret[0] = z

	for i := range s {
		z = fn(z, s[i])
		ret[i+1] = z
	}

	return ret
}

func SliceScanRight[T, U any](s Slice[T], z U, fn Func2[T, U, U]) Slice[U] {
	size := s.Len()
	ret := MakeSlice[U](size+1)
	ret[size] = z

	for i := range s {
		z = fn(s[size-1-i], z)
		ret[size-1-i] = z
	}

	return ret
}

func SliceScan[T, U any](s Slice[T], z U, fn Func2[U, T, U]) Slice[U] {
	return SliceScanLeft(s, z, fn)
}

func SliceGroupBy[T any, K comparable](s Slice[T], fn Func1[T, K]) Map[K, Slice[T]] {
	
	return SliceGroupMap(
		s,
		fn,
		func(v T) T {
			return v
		},
	)
}

func SliceGroupMap[T any, K comparable, V any](s Slice[T], key Func1[T, K], val Func1[T, V]) Map[K, Slice[V]] {
	ret := map[K]Slice[V]{}

	for i := range s {
		k := key(s[i])
		v := val(s[i])
		ret[k] = append(ret[k], v)
	}
	return ret
}

func SliceGroupMapReduce[T any, K comparable, V any](s Slice[T], key Func1[T, K], val Func1[T, V], fn Func2[V, V, V]) Map[K, V] {
	m := SliceGroupMap(s, key, val)
	ret := map[K]V{}
	for k := range m {
		ret[k] = m[k].Reduce(fn).Get()
	}
	return ret
}

func sliceMaxBy[T any, B Ordered](s Slice[T], fn Func1[T, B], cmp CompareFunc[B]) Option[T] {
	size := s.Len()
	if size == 0 {
		return None[T]()
	}

	if size == 1 {
		return Some[T](s[0])
	}

	v := s[0]
	x := fn(s[0])
	for i := 1; i < size; i++ {
		y := fn(s[i])
		if cmp(x, y) < 0 {
			x = y
			v = s[i]
		}
	}

	return Some[T](v)
}

func SliceMaxBy[T any, B Ordered](s Slice[T], fn Func1[T, B]) Option[T] {
	return sliceMaxBy(s, fn, Compare[B])
}

func SliceMinBy[T any, B Ordered](s Slice[T], fn Func1[T, B]) Option[T] {
	cmp := func(v1, v2 B) int {
		return -Compare(v1, v2)
	}
	return sliceMaxBy(s, fn, cmp)
}

func SliceSortBy[T any, B Ordered](s Slice[T], fn Func1[T, B]) Slice[T] {
	sort.SliceStable(s, func(i, j int) bool {
		return fn(s[i]) < fn(s[j])
	})
	return s
}

func SliceToMap[K comparable, V any](s Slice[Pair[K, V]]) Map[K, V] {
	ret := MakeMap[K, V](s.Len())

	for i := range s {
		ret[s[i].Key()] = s[i].Value()
	}

	return ret
}

func SliceToSet[K comparable](s Slice[K]) Set[K] {
	ret := MakeMap[K, bool](s.Len())
	for i := range s {
		ret[s[i]] = true
	}
	return Set[K](ret)
}