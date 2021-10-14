// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package slices

import (
	"constraints"
	"sort"

	gs "github.com/kigichang/goscala"
)

type Slice[T any] []T

func (s Slice[T]) Clone() Slice[T] {
	ret := make([]T, len(s))
	copy(ret, s)
	return ret
}

func (s Slice[T]) Len() int {
	return len(s)
}

func (s Slice[T]) Cap() int {
	return cap(s)
}

func (s Slice[T]) Slice() []T {
	return []T(s)
}

func (s Slice[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s Slice[T]) Forall(fn func(T) bool) bool {
	if s.IsEmpty() {
		return true
	}

	for i := range s {
		if !fn(s[i]) {
			return false
		}
	}

	return true
}

func (s Slice[T]) Foreach(fn func(T)) {
	for i := range s {
		fn(s[i])
	}
}

func (s Slice[T]) Head() gs.Option[T] {
	if !s.IsEmpty() {
		return gs.Some[T](s[0])
	}
	return gs.None[T]()
}

func (s Slice[T]) Last() gs.Option[T] {
	if !s.IsEmpty() {
		return gs.Some[T](s[s.Len()-1])
	}
	return gs.None[T]()
}

func (s Slice[T]) Tail() Slice[T] {
	if s.IsEmpty() {
		return Empty[T]()
	}

	return s[1:]
}

func (s Slice[T]) Equals(eq func(T, T) bool) func(Slice[T]) bool {
	return func(that Slice[T]) bool {
		if &s == &that {
			return true
		}

		if s.Len() != that.Len() {
			return false
		}

		for i := range s {
			if !eq(s[i], that[i]) {
				return false
			}
		}
		return true
	}

}

func (s Slice[T]) Contains(eq func(T, T) bool) func(T) bool {
	return func(elem T) bool {
		p := gs.Currying2(eq)(elem)
		return s.Exists(p)
	}
}

func (s Slice[T]) Exists(p func(T) bool) (ok bool) {
	return s.Find(p).IsDefined()
}

func (s Slice[T]) Filter(p func(T) bool) Slice[T] {
	ret := Empty[T]()

	for i := range s {
		if p(s[i]) {
			ret = append(ret, s[i])
		}
	}

	return ret
}

func (s Slice[T]) FilterNot(p func(T) bool) Slice[T] {
	return s.Filter(func(v T) bool { return !p(v) })
}

func (s Slice[T]) Find(p func(T) bool) gs.Option[T] {
	for i := range s {
		if p(s[i]) {
			return gs.Some[T](s[i])
		}
	}
	return gs.None[T]()
}

func (s Slice[T]) FindLast(p func(T) bool) gs.Option[T] {
	size := s.Len()
	for i := size - 1; i >= 0; i-- {
		if p(s[i]) {
			return gs.Some[T](s[i])
		}
	}
	return gs.None[T]()
}

func (s Slice[T]) Partition(p func(T) bool) (Slice[T], Slice[T]) {
	a, b := Empty[T](), Empty[T]()

	for i := range s {
		if p(s[i]) {
			b = append(b, s[i])
		} else {
			a = append(a, s[i])
		}
	}

	return a, b
}

func (s Slice[T]) Reverse() Slice[T] {
	size := s.Len()

	ret := make([]T, size)

	for i := range s {
		ret[size-1-i] = s[i]
	}

	return ret
}

func (s Slice[T]) indexHelper(n int) (idx int, size int) {
	size = s.Len()
	idx = n
	if n < 0 {
		idx = size + n
	}
	return
}

func (s Slice[T]) SplitAt(n int) (Slice[T], Slice[T]) {

	idx, size := s.indexHelper(n)

	if idx <= 0 {
		return Empty[T](), s
	}

	if idx >= size {
		return s, Empty[T]()
	}

	return s[0:idx], s[idx:]
}

func (s Slice[T]) Take(n int) Slice[T] {
	a, b := s.SplitAt(n)
	if n >= 0 {
		return a
	}

	return b
}

func (s Slice[T]) TakeWhile(p func(T) bool) Slice[T] {
	ret := Empty[T]()

	for i := range s {
		if !p(s[i]) {
			break
		}
		ret = append(ret, s[i])
	}

	return ret
}

func (s Slice[T]) Count(p func(T) bool) (ret int) {
	for i := range s {
		if p(s[i]) {
			ret += 1
		}
	}
	return
}

func (s Slice[T]) Drop(n int) Slice[T] {
	a, b := s.SplitAt(n)
	if n >= 0 {
		return b
	}
	return a
}

func (s Slice[T]) DropWhile(p func(T) bool) Slice[T] {
	for i := range s {
		if !p(s[i]) {
			return s[i:]
		}
	}
	return Empty[T]()
}

func (s Slice[T]) ReduceRight(fn func(T, T) T) gs.Option[T] {
	size := s.Len()
	if size <= 0 {
		return gs.None[T]()
	}

	if size == 1 {
		return gs.Some[T](s[0])
	}

	return gs.Some[T](FoldRight(s[:size-1], s[size-1], fn))
}

func (s Slice[T]) ReduceLeft(fn func(T, T) T) gs.Option[T] {
	size := s.Len()
	if size <= 0 {
		return gs.None[T]()
	}

	if size == 1 {
		return gs.Some[T](s[0])
	}

	return gs.Some[T](FoldLeft(s[1:], s[0], fn))
}

func (s Slice[T]) Reduce(fn func(T, T) T) gs.Option[T] {
	return s.ReduceLeft(fn)
}

func (s Slice[T]) IndexWhereFrom(p func(T) bool, from int) int {
	size := s.Len()
	from = gs.Max(0, from)
	for i := from; i < size; i++ {
		if p(s[i]) {
			return i
		}
	}

	return -1
}

func (s Slice[T]) IndexWhere(p func(T) bool) int {
	return s.IndexWhereFrom(p, 0)
}

func (s Slice[T]) IndexFrom(elem T, from int, fn func(T, T) bool) int {
	p := func(v T) bool {
		return fn(v, elem)
	}
	return s.IndexWhereFrom(p, from)
}

func (s Slice[T]) Index(elem T, fn func(T, T) bool) int {
	return s.IndexFrom(elem, 0, fn)
}

func (s Slice[T]) LastIndexWhereFrom(p func(T) bool, from int) int {
	size := s.Len()
	from = gs.Min(from, size-1)

	for i := from; i >= 0; i-- {
		if p(s[i]) {
			return i
		}
	}

	return -1
}

func (s Slice[T]) LastIndexWhere(p func(T) bool) int {
	return s.LastIndexWhereFrom(p, s.Len())
}

func (s Slice[T]) LastIndexFrom(elem T, from int, fn func(T, T) bool) int {
	p := func(v T) bool {
		return fn(v, elem)
	}
	return s.LastIndexWhereFrom(p, from)
}

func (s Slice[T]) LastIndex(elem T, fn func(T, T) bool) int {
	return s.LastIndexFrom(elem, s.Len(), fn)
}

func (s Slice[T]) Max(compare func(T, T) int) gs.Option[T] {
	size := s.Len()
	if size == 0 {
		return gs.None[T]()
	}

	if size == 1 {
		return gs.Some[T](s[0])
	}

	v := s[0]

	for i := 1; i < size; i++ {
		if compare(v, s[i]) < 0 {
			v = s[i]
		}
	}
	return gs.Some[T](v)
}

func (s Slice[T]) Min(compare func(T, T) int) gs.Option[T] {
	cmp := func(a, b T) int {
		return -compare(a, b)
	}

	return s.Max(cmp)
}

func (s Slice[T]) Sort(compare func(T, T) int) Slice[T] {
	sort.SliceStable(
		s,
		func(i, j int) bool {
			return compare(s[i], s[j]) < 0
		},
	)
	return s
}

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

	return Slice[T](make([]T, size, cap))
}

func From[T any](a ...T) Slice[T] {
	if a == nil {
		return Empty[T]()
	}
	return a
}

func One[T any](elem T) Slice[T] {
	return []T{elem}
}

func Empty[T any]() Slice[T] {
	return []T{}
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

func CollectFirst[T, U any](s Slice[T], pf func(T) (U, bool)) gs.Option[U] {
	for i := range s {
		if v, ok := pf(s[i]); ok {
			return gs.Some[U](v)
		}
	}
	return gs.None[U]()
}

func FoldLeft[T, U any](s Slice[T], z U, fn func(a U, b T) U) U {
	zz := z
	for i := range s {
		zz = fn(zz, s[i])
	}
	return zz
}

func FoldRight[T, U any](s Slice[T], z U, fn func(T, U) U) U {
	zz := z
	size := len(s)
	for i := size - 1; i >= 0; i-- {
		zz = fn(s[i], zz)
	}
	return zz
}

func Fold[T, U any](s Slice[T], z U, fn func(a U, b T) U) U {
	return FoldLeft(s, z, fn)
}

func ScanLeft[T, U any](s Slice[T], z U, fn func(U, T) U) Slice[U] {
	return FoldLeft[T, Slice[U]](s, One[U](z), func(a Slice[U], b T) Slice[U] {
		return append(a, fn(a[len(a)-1], b))
	})
}

func ScanRight[T, U any](s Slice[T], z U, fn func(T, U) U) Slice[U] {
	result := FoldRight[T, Slice[U]](s, One[U](z), func(a T, b Slice[U]) Slice[U] {
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

func Scan[T, U any](s Slice[T], z U, fn func(U, T) U) Slice[U] {
	return ScanLeft(s, z, fn)
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
		ret[k] = m[k].Reduce(fn).Get()
	}
	return ret
}

func maxBy[T any, B constraints.Ordered](s Slice[T], fn func(T) B, cmp func(B, B) int) gs.Option[T] {
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

//func ToMap[K comparable, V any](s Slice[gs.Pair[K, V]]) gs.Map[K, V] {
//	ret := gs.MkMap[K, V](s.Len())
//
//	for i := range s {
//		ret.Add(s[i])
//	}
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
