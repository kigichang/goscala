package slices

import (
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

//
//func SliceGroupBy[T any, K comparable](s Slice[T], fn Func1[T, K]) Map[K, Slice[T]] {
//
//	return SliceGroupMap(
//		s,
//		fn,
//		func(v T) T {
//			return v
//		},
//	)
//}
//
//func SliceGroupMap[T any, K comparable, V any](s Slice[T], key Func1[T, K], val Func1[T, V]) Map[K, Slice[V]] {
//	ret := map[K]Slice[V]{}
//
//	for i := range s {
//		k := key(s[i])
//		v := val(s[i])
//		ret[k] = append(ret[k], v)
//	}
//	return ret
//}
//
//func SliceGroupMapReduce[T any, K comparable, V any](s Slice[T], key Func1[T, K], val Func1[T, V], fn Func2[V, V, V]) Map[K, V] {
//	m := SliceGroupMap(s, key, val)
//	ret := map[K]V{}
//	for k := range m {
//		ret[k] = m[k].Reduce(fn).Get()
//	}
//	return ret
//}
//
//func sliceMaxBy[T any, B Ordered](s Slice[T], fn Func1[T, B], cmp CompareFunc[B]) Option[T] {
//	size := s.Len()
//	if size == 0 {
//		return None[T]()
//	}
//
//	if size == 1 {
//		return Some[T](s[0])
//	}
//
//	v := s[0]
//	x := fn(s[0])
//	for i := 1; i < size; i++ {
//		y := fn(s[i])
//		if cmp(x, y) < 0 {
//			x = y
//			v = s[i]
//		}
//	}
//
//	return Some[T](v)
//}
//
//func SliceMaxBy[T any, B Ordered](s Slice[T], fn Func1[T, B]) Option[T] {
//	return sliceMaxBy(s, fn, Compare[B])
//}
//
//func SliceMinBy[T any, B Ordered](s Slice[T], fn Func1[T, B]) Option[T] {
//	cmp := func(v1, v2 B) int {
//		return -Compare(v1, v2)
//	}
//	return sliceMaxBy(s, fn, cmp)
//}
//
//func SliceSortBy[T any, B Ordered](s Slice[T], fn Func1[T, B]) Slice[T] {
//	sort.SliceStable(s, func(i, j int) bool {
//		return fn(s[i]) < fn(s[j])
//	})
//	return s
//}
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
