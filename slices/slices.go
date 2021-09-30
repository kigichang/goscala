package slices

import (
	"constraints"
	"sort"
	gs "github.com/kigichang/goscala"
)

func toSlice[T any](s gs.Sliceable[T]) []T {
	return s.Slice()
}

func Make[T any](a ...int) gs.Slice[T] {

	size, cap := 0, 0

	if len(a) > 0 {
		size, cap = a[0], a[0]
	}

	if len(a) > 1 {
		cap = a[1]
	}

	if cap < size {
		cap = size
	}

	return gs.Slice[T](make([]T, size, cap))
}

func Empty[T any]() gs.Slice[T] {
	return gs.SliceEmpty[T]()
}

func One[T any](elem T) gs.Slice[T] {
	return gs.SliceOne[T](elem)
}

func From[T any](a ...T) gs.Slice[T] {
	if a == nil {
		return Empty[T]()
	}
	return a
}

func Fill[T any](size int, v T) gs.Slice[T] {
	ret := Make[T](size)
	for i := range ret {
		ret[i] = v
	}
	return ret
}

func Range[T gs.Numeric](start, end, step T) gs.Slice[T] {
	ret := Make[T](0)

	for i := start; i < end; i += step {
		ret = append(ret, i)
	}

	return ret
}

func Tabulate[T any](size int, f func(int) T) gs.Slice[T] {
	if size <= 0 {
		return Empty[T]()
	}

	ret := Make[T](size)

	for i := 0; i < size; i++ {
		ret[i] = f(i)
	}

	return ret
}

func Collect[T, U any](s gs.Slice[T], pf func(T) (U, bool)) gs.Slice[U] {
	ret := Empty[U]()
	for i := range s {
		if v, ok := pf(s[i]); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func CollectFirst[T, U any](s gs.Slice[T], pf func(T) (U, bool)) (ret U, ok bool) {
	for i := range s {
		if ret, ok = pf(s[i]); ok {
			return
		}
	}
	return
}

func FlatMap[T, U any](s gs.Slice[T], fn func(T) gs.Sliceable[U]) gs.Slice[U] {
	return gs.FlatMap(s, fn)
}

func FoldRight[T, U any](s gs.Slice[T], z U, fn func(T, U) U) U {
	return gs.FoldRight(s, z, fn)
}

func FoldLeft[T, U any](s gs.Slice[T], z U, fn func(U, T) U) U {
	return gs.FoldLeft(s, z, fn)
}

func Fold[T, U any](s gs.Slice[T], z U, fn func(U, T) U) U {
	return FoldLeft(s, z, fn)
}

func Map[T, U any](s gs.Slice[T], fn func(T) U) gs.Slice[U] {
	return gs.Map(s, fn)
}

func PartitionMap[T, A, B any](s gs.Slice[T], fn func(T) gs.Either[A, B]) (gs.Slice[A], gs.Slice[B]) {
	a, b := Empty[A](), Empty[B]()
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

func ScanLeft[T, U any](s gs.Slice[T], z U, fn func(U, T) U) gs.Slice[U] {
	return gs.ScanLeft(s, z, fn)
}

func Scan[T, U any](s gs.Slice[T], z U, fn func(U, T) U) gs.Slice[U] {
	return ScanLeft(s, z, fn)
}

func ScanRight[T, U any](s gs.Slice[T], z U, fn func(T, U) U) gs.Slice[U] {
	return gs.ScanRight(s, z, fn)
}

func GroupMap[T any, K comparable, V any](s gs.Slice[T], key func(T) K, val func(T) V) map[K]gs.Slice[V] {
	ret := map[K]gs.Slice[V]{}
	for i := range s {
		k := key(s[i])
		v := val(s[i])
		ret[k] = append(ret[k], v)
	}
	return ret
}

func GroupBy[T any, K comparable](s gs.Slice[T], fn func(T) K) map[K]gs.Slice[T] {
	return GroupMap(s,fn, gs.Id[T])
}

func GroupMapReduce[T any, K comparable, V any](s gs.Slice[T], key func(T) K, val func(T) V, fn func(V, V) V) map[K]V {
	m := GroupMap(s, key, val)
	ret := map[K]V{}
	for k := range m {
		ret[k], _ = m[k].Reduce(fn)
	}
	return ret
}

func maxBy[T any, B constraints.Ordered](s gs.Slice[T], fn func(T) B, cmp func(B, B) int) gs.Option[T] {
	size := s.Len()
	if size == 0 {
		return gs.None[T]()
	}

	if size == 1 {
		return gs.Some[T](s[0])
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

	return gs.Some[T](v)
}

func MaxBy[T any, B constraints.Ordered](s gs.Slice[T], fn func(T) B) gs.Option[T] {
	return maxBy(s, fn, gs.Compare[B])
}

func MinBy[T any, B constraints.Ordered](s gs.Slice[T], fn func(T) B) gs.Option[T] {
	cmp := func(v1, v2 B) int {
		return -gs.Compare(v1, v2)
	}
	return maxBy(s, fn, cmp)
}

func SortBy[T any, B constraints.Ordered](s gs.Slice[T], fn func(T) B) gs.Slice[T] {
	sort.SliceStable(s, func(i, j int) bool {
		return fn(s[i]) < fn(s[j])
	})
	return s
}
//
//func SliceToMap[K comparable, V any](s Slice[Pair[K, V]]) Map[K, V] {
//	ret := MakeMap[K, V](s.Len())
//
//	for i := range s {
//		ret[s[i].Key()] = s[i].Value()
//	}
//
//	return ret
//}
//
//func SliceToSet[K comparable](s Slice[K]) Set[K] {
//	ret := MakeMap[K, bool](s.Len())
//	for i := range s {
//		ret[s[i]] = true
//	}
//	return Set[K](ret)
//}
