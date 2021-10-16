// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package seq

import (
	"constraints"

	"github.com/kigichang/goscala"
)

type Iterator[T any] interface {
	Next() bool
	Get() T
}

func skip[T any](it Iterator[T], n int) {
	for i := 0; i < n && it.Next(); i++ {
	}
}

type Interface[T any] interface {
	Len() int
	Iterate() Iterator[T]
	BackIterate() Iterator[T]
	Append(T)
}

type Constraint[T any] interface {
	Interface[T]
}

func cloneFrom[E Constraint[T], T any](it Iterator[T], empty func() E) E {
	ret := empty()
	for it.Next() {
		ret.Append(it.Get())
	}
	return ret
}

func Clone[E Constraint[T], T any](s E, empty func() E) E {
	return cloneFrom(s.Iterate(), empty)
}

func Forall[E Constraint[T], T any](s E, fn func(T) bool) bool {

	for it := s.Iterate(); it.Next(); {
		if !fn(it.Get()) {
			return false
		}
	}
	return true
}

func Foreach[E Constraint[T], T any](s E, fn func(T)) {

	for it := s.Iterate(); it.Next(); {
		fn(it.Get())
	}
}

func Equals[E Constraint[T], T any](eq func(T, T) bool, a, b E) bool {
	if a.Len() != b.Len() {
		return false
	}

	ita := a.Iterate()
	itb := b.Iterate()

	for ita.Next() {
		if !itb.Next() || !eq(ita.Get(), itb.Get()) {
			return false
		}
	}
	return true

}

func find[T any](it Iterator[T], p func(T) bool) goscala.Option[T] {
	for it.Next() {
		v := it.Get()
		if p(v) {
			return goscala.Some[T](v)
		}
	}
	return goscala.None[T]()
}

func Find[E Constraint[T], T any](s E, p func(T) bool) goscala.Option[T] {
	return find(s.Iterate(), p)
}

func Contains[E Constraint[T], T any](s E, eq func(T, T) bool) func(T) bool {
	return func(elem T) bool {
		p := goscala.Currying2(eq)(elem)
		return Exists(s, p)
	}
}

func Exists[E Constraint[T], T any](s E, p func(T) bool) (ok bool) {
	return Find(s, p).IsDefined()
}

func Filter[E Constraint[T], T any](s E, empty func() E, p func(T) bool) E {
	ret := empty()

	for it := s.Iterate(); it.Next(); {
		v := it.Get()
		if p(v) {
			ret.Append(v)
		}
	}

	return ret
}

func FilterNot[E Constraint[T], T any](s E, empty func() E, p func(T) bool) E {
	return Filter(s, empty, func(v T) bool { return !p(v) })
}

func FindLast[E Constraint[T], T any](s E, p func(T) bool) goscala.Option[T] {
	return find(s.BackIterate(), p)
}

func Partition[E Constraint[T], T any](s E, empty func() E, p func(T) bool) (E, E) {
	a, b := empty(), empty()

	for it := s.Iterate(); it.Next(); {
		v := it.Get()
		if p(v) {
			b.Append(v)
		} else {
			a.Append(v)
		}
	}

	return a, b
}

func Reverse[E Constraint[T], T any](s E, empty func() E) E {
	ret := empty()

	for it := s.BackIterate(); it.Next(); {
		ret.Append(it.Get())
	}
	return ret
}

func indexHelper[E Constraint[T], T any](s E, n int) (idx int, size int) {
	size = s.Len()
	idx = n
	if n < 0 {
		idx = size + n
	}
	return
}

func SplitAt[E Constraint[T], T any](s E, empty func() E, n int) (E, E) {

	idx, size := indexHelper[E, T](s, n)

	if idx <= 0 {
		return empty(), Clone[E, T](s, empty)
	}

	if idx >= size {
		return Clone[E, T](s, empty), empty()
	}

	a, b := empty(), empty()

	for i, it := 0, s.Iterate(); it.Next(); i++ {
		if i < idx {
			a.Append(it.Get())
		} else {
			b.Append(it.Get())
		}
	}
	return a, b
}

func Take[E Constraint[T], T any](s E, empty func() E, n int) E {
	idx, _ := indexHelper[E, T](s, n)

	if n < 0 {
		it := s.Iterate()
		skip(it, idx)

		return cloneFrom(it, empty)
	}

	ret := empty()
	for i, it := 0, s.Iterate(); i < idx && it.Next(); i++ {
		ret.Append(it.Get())
	}
	return ret
}

func TakeWhile[E Constraint[T], T any](s E, empty func() E, p func(T) bool) E {
	ret := empty()

	for it := s.Iterate(); it.Next(); {
		v := it.Get()
		if !p(v) {
			break
		}
		ret.Append(v)
	}
	return ret
}

func Count[E Constraint[T], T any](s E, p func(T) bool) (ret int) {
	for it := s.Iterate(); it.Next(); {
		if p(it.Get()) {
			ret += 1
		}
	}
	return
}

func Drop[E Constraint[T], T any](s E, empty func() E, n int) E {
	idx, _ := indexHelper[E, T](s, n)

	if n < 0 {
		ret := empty()
		for i, it := 0, s.Iterate(); i < idx && it.Next(); i++ {
			ret.Append(it.Get())
		}
		return ret
	}

	it := s.Iterate()

	skip(it, idx)
	return cloneFrom(it, empty)
}

func DropWhile[E Constraint[T], T any](s E, empty func() E, p func(T) bool) E {
	it := s.Iterate()
	ret := empty()
	appended := false

	for it.Next() {
		v := it.Get()
		if appended {
			ret.Append(v)
		} else if !p(v) {
			ret.Append(v)
			appended = true
		}
	}

	return ret
}

func fold[T, U any](it Iterator[T], z U, fn func(U, T) U) U {
	zz := z
	for it.Next() {
		zz = fn(zz, it.Get())
	}
	return zz
}

func FoldLeft[E Constraint[T], T, U any](s E, z U, fn func(a U, b T) U) U {
	return fold(s.Iterate(), z, fn)
}

func FoldRight[E Constraint[T], T, U any](s E, z U, fn func(T, U) U) U {
	return fold(s.BackIterate(), z, func(a U, b T) U {
		return fn(b, a)
	})
}

func Fold[E Constraint[T], T, U any](s E, z U, fn func(a U, b T) U) U {
	return Fold(s, z, fn)
}

func ReduceLeft[E Constraint[T], T any](s E, fn func(T, T) T) goscala.Option[T] {
	it := s.Iterate()
	if !it.Next() {
		return goscala.None[T]()
	}
	z := it.Get()

	return goscala.Some[T](fold(it, z, fn))
}

func ReduceRight[E Constraint[T], T any](s E, fn func(T, T) T) goscala.Option[T] {
	it := s.BackIterate()
	if !it.Next() {
		return goscala.None[T]()
	}
	z := it.Get()

	return goscala.Some(fold(it, z, fn))
}

func Reduce[E Constraint[T], T any](s E, fn func(T, T) T) goscala.Option[T] {
	return ReduceLeft(s, fn)
}

func IndexWhereFrom[E Constraint[T], T any](s E, p func(T) bool, from int) int {
	from = goscala.Max(0, from)
	it := s.Iterate()
	skip(it, from)
	for i := from; it.Next(); i++ {
		if p(it.Get()) {
			return i
		}
	}
	return -1
}

func IndexWhere[E Constraint[T], T any](s E, p func(T) bool) int {
	return IndexWhereFrom(s, p, 0)
}

func IndexFrom[E Constraint[T], T any](s E, elem T, from int, eq func(T, T) bool) int {
	p := goscala.Currying2(eq)(elem)
	return IndexWhereFrom(s, p, from)
}

func Index[E Constraint[T], T any](s E, elem T, eq func(T, T) bool) int {
	return IndexFrom(s, elem, 0, eq)
}

func LastIndexWhereFrom[E Constraint[T], T any](s E, p func(T) bool, from int) int {
	bound := s.Len() - 1
	from = goscala.Min(from, bound)
	it := s.BackIterate()
	skip(it, bound-from)

	for i := from; i >= 0 && it.Next(); i-- {
		if p(it.Get()) {
			return i
		}
	}

	return -1
}

func LastIndexWhere[E Constraint[T], T any](s E, p func(T) bool) int {
	return LastIndexWhereFrom(s, p, s.Len()-1)
}

func LastIndexFrom[E Constraint[T], T any](s E, elem T, from int, eq func(T, T) bool) int {
	p := goscala.Currying2(eq)(elem)
	return LastIndexWhereFrom(s, p, from)
}

func LastIndex[E Constraint[T], T any](s E, elem T, fn func(T, T) bool) int {
	return LastIndexFrom(s, elem, s.Len()-1, fn)
}

func Max[E Constraint[T], T any](s E, compare func(T, T) int) goscala.Option[T] {
	it := s.Iterate()
	if !it.Next() {
		return goscala.None[T]()
	}

	v := it.Get()

	for it.Next() {
		u := it.Get()
		if compare(v, u) < 0 {
			v = u
		}
	}

	return goscala.Some[T](v)
}

func Min[E Constraint[T], T any](s E, compare func(T, T) int) goscala.Option[T] {
	cmp := func(a, b T) int {
		return -compare(a, b)
	}

	return Max(s, cmp)
}

func From[E Constraint[T], T any](empty func() E, a ...T) E {
	ret := empty()

	for i := range a {
		ret.Append(a[i])
	}

	return ret
}

func One[E Constraint[T], T any](empty func() E, elem T) E {
	ret := empty()
	ret.Append(elem)
	return ret
}

func Fill[E Constraint[T], T any](empty func() E, size int, v T) E {
	ret := empty()
	for i := 0; i < size; i++ {
		ret.Append(v)
	}
	return ret
}

func Range[E Constraint[T], T goscala.Numeric](empty func() E, start, end, step T) E {
	ret := empty()

	for i := start; i < end; i += step {
		ret.Append(i)
	}

	return ret
}

func Tabulate[E Constraint[T], T any](empty func() E, size int, f func(int) T) E {
	if size <= 0 {
		return empty()
	}

	ret := empty()

	for i := 0; i < size; i++ {
		ret.Append(f(i))
	}

	return ret
}

func Collect[E1 Constraint[T], E2 Constraint[U], T, U any](s E1, empty func() E2, pf func(T) (U, bool)) E2 {
	ret := empty()
	for it := s.Iterate(); it.Next(); {
		if v, ok := pf(it.Get()); ok {
			ret.Append(v)
		}
	}
	return ret
}

func CollectFirst[E Constraint[T], T, U any](s E, pf func(T) (U, bool)) goscala.Option[U] {
	for it := s.Iterate(); it.Next(); {
		if v, ok := pf(it.Get()); ok {
			return goscala.Some[U](v)
		}
	}

	return goscala.None[U]()
}

func ScanLeft[E1 Constraint[T], E2 Constraint[U], T, U any](s E1, empty func() E2, z U, fn func(U, T) U) E2 {
	last := z

	return FoldLeft(s, One(empty, z), func(a E2, b T) E2 {
		last = fn(last, b)
		a.Append(last)
		return a
	})
}

func ScanRight[E1 Constraint[T], E2 Constraint[U], T, U any](s E1, empty func() E2, z U, fn func(T, U) U) E2 {
	last := z
	result := FoldRight(s, One(empty, z), func(a T, b E2) E2 {
		last := fn(a, last)
		b.Append(last)
		return b
	})

	return cloneFrom(result.BackIterate(), empty)
}

func Scan[E1 Constraint[T], E2 Constraint[U], T, U any](s E1, empty func() E2, z U, fn func(U, T) U) E2 {
	return ScanLeft(s, empty, z, fn)
}

func Map[E1 Constraint[T], E2 Constraint[U], T, U any](s E1, empty func() E2, fn func(T) U) E2 {
	return FoldLeft(s, empty(), func(z E2, a T) E2 {
		z.Append(fn(a))
		return z
	})
}

func FlatMap[E1 Constraint[T], E2 Constraint[U], T, U any](s E1, empty func() E2, fn func(T) Interface[U]) E2 {
	return FoldLeft(s, empty(), func(z E2, a T) E2 {
		tmp := fn(a)
		for it := tmp.Iterate(); it.Next(); {
			z.Append(it.Get())
		}
		return z
	})
}

func PartitionMap[E1 Constraint[T], E2 Constraint[A], E3 Constraint[B], T, A, B any](s E1, emptya func() E2, emptyb func() E3, fn func(T) goscala.Either[A, B]) (E2, E3) {
	a, b := emptya(), emptyb()

	for it := s.Iterate(); it.Next(); {
		v := fn(it.Get())
		if v.IsRight() {
			b.Append(v.Right())
		} else {
			a.Append(v.Left())
		}
	}

	return a, b
}

func GroupMap[E1 Constraint[T], E2 Constraint[V], T any, K comparable, V any](s E1, empty func() E2, key func(T) K, val func(T) V) map[K]E2 {
	ret := map[K]E2{}

	for it := s.Iterate(); it.Next(); {
		tmp := it.Get()
		k := key(tmp)
		v := val(tmp)
		s, ok := ret[k]
		if !ok {
			s = empty()
		}
		s.Append(v)
		ret[k] = s
	}
	return ret
}

func GroupBy[E Constraint[T], T any, K comparable](s E, empty func() E, fn func(T) K) map[K]E {
	return GroupMap(s, empty, fn, goscala.Id[T])
}

func GroupMapReduce[E1 Constraint[T], E2 Constraint[V], T any, K comparable, V any](s E1, empty func() E2, key func(T) K, val func(T) V, fn func(V, V) V) map[K]V {
	m := GroupMap(s, empty, key, val)
	ret := map[K]V{}
	for k := range m {
		ret[k] = Reduce(m[k], fn).Get()
	}
	return ret
}

func maxBy[E Constraint[T], T any, B constraints.Ordered](s E, fn func(T) B, cmp func(B, B) int) goscala.Option[T] {
	it := s.Iterate()

	if !it.Next() {
		return goscala.None[T]()
	}

	v := it.Get()
	x := fn(v)

	for it.Next() {
		tmp := it.Get()
		y := fn(tmp)
		if cmp(x, y) < 0 {
			x = y
			v = tmp
		}
	}

	return goscala.Some[T](v)
}

func MaxBy[E Constraint[T], T any, B constraints.Ordered](s E, fn func(T) B) goscala.Option[T] {
	return maxBy(s, fn, goscala.Compare[B])
}

func MinBy[E Constraint[T], T any, B constraints.Ordered](s E, fn func(T) B) goscala.Option[T] {
	cmp := func(v1, v2 B) int {
		return -goscala.Compare(v1, v2)
	}
	return maxBy(s, fn, cmp)
}
