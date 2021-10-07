package slices

import (
	"constraints"
	"sort"

	gs "github.com/kigichang/goscala"
	"github.com/kigichang/goscala/opt"
)

func Make[T any](a ...int) Slice[T] {

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

	return make([]T, size, cap)
}

func Empty[T any]() Slice[T] {
	return []T{}
}

func One[T any](elem T) Slice[T] {
	return []T{elem}
}

func From[T any](a ...T) Slice[T] {
	if a == nil {
		return Empty[T]()
	}
	return a
}

func Fill[T any](size int, v T) Slice[T] {
	ret := Make[T](size)
	for i := range ret {
		ret[i] = v
	}
	return ret
}

func Range[T gs.Numeric](start, end, step T) Slice[T] {
	ret := Make[T](0)

	for i := start; i < end; i += step {
		ret = append(ret, i)
	}

	return ret
}

func Tabulate[T any](size int, f func(int) T) Slice[T] {
	if size <= 0 {
		return Empty[T]()
	}

	ret := Make[T](size)

	for i := 0; i < size; i++ {
		ret[i] = f(i)
	}

	return ret
}

func Collect[T, U any](s Slice[T], pf func(T) (U, bool)) Slice[U] {
	ret := Empty[U]()
	for i := range s {
		if v, ok := pf(s[i]); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func CollectFirst[T, U any](s Slice[T], pf func(T) (U, bool)) (ret U, ok bool) {
	for i := range s {
		if ret, ok = pf(s[i]); ok {
			return
		}
	}
	return
}

func FlatMap[T, U any](s Slice[T], fn func(T) gs.Sliceable[U]) Slice[U] {
	return FoldLeft[T, Slice[U]](
		s,
		Empty[U](),
		func(z Slice[U], a T) Slice[U] {
			z = append(z, fn(a).Slice()...)
			return z
		},
	)
}

func FoldRight[T, U any](s Slice[T], z U, fn func(T, U) U) U {
	zz := z
	size := len(s)
	for i := size - 1; i >= 0; i-- {
		zz = fn(s[i], zz)
	}
	return zz
}

func FoldLeft[T, U any](s Slice[T], z U, fn func(U, T) U) U {
	zz := z
	for i := range s {
		zz = fn(zz, s[i])
	}
	return zz
}

func Fold[T, U any](s Slice[T], z U, fn func(U, T) U) U {
	return FoldLeft(s, z, fn)
}

func Map[T, U any](s Slice[T], fn func(T) U) Slice[U] {
	return FoldLeft[T, Slice[U]](
		s,
		Empty[U](),
		func(z Slice[U], a T) Slice[U] {
			z = append(z, fn(a))
			return z
		},
	)
}

func PartitionMap[T, A, B any](s Slice[T], fn func(T) gs.Either[A, B]) (Slice[A], Slice[B]) {
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

func ScanLeft[T, U any](s Slice[T], z U, fn func(U, T) U) Slice[U] {
	return FoldLeft[T, []U](s, []U{z}, func(a []U, b T) []U {
		return append(a, fn(a[len(a)-1], b))
	})
}

func Scan[T, U any](s Slice[T], z U, fn func(U, T) U) Slice[U] {
	return ScanLeft(s, z, fn)
}

func ScanRight[T, U any](s Slice[T], z U, fn func(T, U) U) Slice[U] {
	result := FoldRight[T, []U](s, []U{z}, func(a T, b []U) []U {
		return append(b, fn(a, b[len(b)-1]))
	})

	size := len(result)
	half := size / 2

	for i := 0; i < half; i++ {
		tmp := result[i]
		result[i] = result[size-1-i]
		result[size-1-i] = tmp
	}

	return result
}

func GroupMap[T any, K comparable, V any](s Slice[T], key func(T) K, val func(T) V) map[K]Slice[V] {
	ret := map[K]Slice[V]{}
	for i := range s {
		k := key(s[i])
		v := val(s[i])
		ret[k] = append(ret[k], v)
	}
	return ret
}

func GroupBy[T any, K comparable](s Slice[T], fn func(T) K) map[K]Slice[T] {
	return GroupMap(s, fn, gs.Id[T])
}

func GroupMapReduce[T any, K comparable, V any](s Slice[T], key func(T) K, val func(T) V, fn func(V, V) V) map[K]V {
	m := GroupMap(s, key, val)
	ret := map[K]V{}
	for k := range m {
		ret[k], _ = m[k].Reduce(fn)
	}
	return ret
}

func maxBy[T any, B constraints.Ordered](s Slice[T], fn func(T) B, cmp func(B, B) int) gs.Option[T] {
	size := s.Len()
	if size == 0 {
		return opt.None[T]()
	}

	if size == 1 {
		return opt.Some[T](s[0])
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

	return opt.Some[T](v)
}

func MaxBy[T any, B constraints.Ordered](s Slice[T], fn func(T) B) gs.Option[T] {
	return maxBy(s, fn, gs.Compare[B])
}

func MinBy[T any, B constraints.Ordered](s Slice[T], fn func(T) B) gs.Option[T] {
	cmp := func(v1, v2 B) int {
		return -gs.Compare(v1, v2)
	}
	return maxBy(s, fn, cmp)
}

func SortBy[T any, B constraints.Ordered](s Slice[T], fn func(T) B) Slice[T] {
	sort.SliceStable(s, func(i, j int) bool {
		return fn(s[i]) < fn(s[j])
	})
	return s
}
