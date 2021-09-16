package goscala

import (
	"sort"
)

type Sliceable[T any] interface {
	Slice() Slice[T]
}

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

func (s Slice[T]) Slice() Slice[T] {
	return s
}

func (s Slice[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s Slice[T]) Forall(fn Predict[T]) bool {
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

func (s Slice[T]) Head() Option[T] {
	if s.IsEmpty() {
		return None[T]()
	}
	return Some[T](s[0])
}

func (s Slice[T]) Last() Option[T] {
	if s.IsEmpty() {
		return None[T]()
	}
	return Some[T](s[s.Len() - 1])
}

func (s Slice[T]) Tail() Slice[T] {
	if s.IsEmpty() {
		return SliceEmpty[T]()
	}

	return s[1:]
}

func (s Slice[T]) Equals(that Slice[T], fn EqualFunc[T]) bool {
	if &s == &that {
		return true
	}

	if s.Len() != that.Len() {
		return false
	}

	for i := range s {
		if !fn(s[i], that[i]) {
			return false
		}
	}

	return true
}

func (s Slice[T]) Contains(elem T, fn EqualFunc[T]) bool {
	p := func(v T) bool {
		return fn(v, elem)
	}
	return s.Exists(p)
}

func (s Slice[T]) Exists(p Predict[T]) bool {
	return s.Find(p).IsDefined()
	//for i := range s {
	//	if p(s[i]) {
	//		return true
	//	}
	//}
	//
	//return false
}

func (s Slice[T]) Filter(p Predict[T]) Slice[T] {
	ret := SliceEmpty[T]()

	for i := range s {
		if p(s[i]) {
			ret = append(ret, s[i])
		}
	}

	return ret
}

func (s Slice[T]) FilterNot(p Predict[T]) Slice[T] {
	return s.Filter(func(v T) bool { return !p(v) } )
}

func (s Slice[T]) Find(p Predict[T]) Option[T] {
	for i := range s {
		if p(s[i]) {
			return Some[T](s[i])
		}
	}
	
	return None[T]()
}

func (s Slice[T]) FindLast(p Predict[T]) Option[T] {
	size := s.Len()
	for i := size - 1; i >= 0; i-- {
		if p(s[i]) {
			return Some[T](s[i])
		}
	}
	return None[T]()
}

func (s Slice[T]) Partition(p Predict[T]) (Slice[T], Slice[T]) {
	a, b := SliceEmpty[T](), SliceEmpty[T]()

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

	ret := MakeSlice[T](size)

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
		return SliceEmpty[T](), s
	}

	if idx >= size {
		return s, SliceEmpty[T]()
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

func (s Slice[T]) TakeWhile(p Predict[T]) Slice[T] {
	ret := SliceEmpty[T]()

	for i := range s {
		if !p(s[i]) {
			break
		}
		ret = append(ret, s[i])
	}

	return ret
}

func (s Slice[T]) Count(p Predict[T]) (ret int) {
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

func (s Slice[T]) DropWhile(p Predict[T]) Slice[T] {
	for i := range s {
		if !p(s[i]) {
			return s[i:]
		}
	}
	return SliceEmpty[T]()
}

func (s Slice[T]) ReduceRight(fn Func2[T, T, T]) Option[T] {
	size := s.Len()
	if size <= 0 {
		return None[T]()
	}

	if size == 1 {
		return Some[T](s[0])
	}

	return Some[T](SliceFoldRight(s[:size-1], s[size-1], fn))
}

func (s Slice[T]) ReduceLeft(fn Func2[T, T, T]) Option[T] {
	size := s.Len()
	if size <= 0 {
		return None[T]()
	}

	if size == 1 {
		return Some[T](s[0])
	}

	return Some[T](SliceFoldLeft(s[1:], s[0], fn))
}

func (s Slice[T]) Reduce(fn Func2[T, T, T]) Option[T] {
	return s.ReduceLeft(fn)
}

func (s Slice[T]) IndexWhereFrom(p Predict[T], from int) int {
	size := s.Len()
	from = Max(0, from)
	for i := from; i < size; i++ {
		if p(s[i]) {
			return i
		}
	}

	return -1
}

func (s Slice[T]) IndexWhere(p Predict[T]) int {
	return s.IndexWhereFrom(p, 0)
}


func (s Slice[T]) IndexFrom(elem T, from int, fn EqualFunc[T]) int {
	p := func(v T) bool {
		return fn(v, elem)
	}
	return s.IndexWhereFrom(p, from)
}

func (s Slice[T]) Index(elem T, fn EqualFunc[T]) int {
	return s.IndexFrom(elem, 0, fn)
}

func (s Slice[T]) LastIndexWhereFrom(p Predict[T], from int) int {
	size := s.Len()
	from = Min(from, size-1)

	for i := from; i >= 0; i-- {
		if p(s[i]) {
			return i
		}
	}

	return -1
}

func (s Slice[T]) LastIndexWhere(p Predict[T]) int {
	return s.LastIndexWhereFrom(p, s.Len())
}


func (s Slice[T]) LastIndexFrom(elem T, from int, fn EqualFunc[T]) int {
	p := func(v T) bool {
		return fn(v, elem)
	}
	return s.LastIndexWhereFrom(p, from)
}

func (s Slice[T]) LastIndex(elem T, fn EqualFunc[T]) int {
	return s.LastIndexFrom(elem, s.Len(), fn)
}

func (s Slice[T]) Max(compare CompareFunc[T]) Option[T] {
	size := s.Len()
	if size == 0 {
		return None[T]()
	}

	if size == 1 {
		return Some[T](s[0])
	}

	v := s[0]

	for i := 1; i < size; i++ {
		if compare(v, s[i]) < 0 {
			v = s[i]
		}
	}
	return Some[T](v)
}

func (s Slice[T]) Min(compare CompareFunc[T]) Option[T] {
	cmp := func(a, b T) int {
		return -compare(a, b)
	}

	return s.Max(cmp)
}

func (s Slice[T]) Sort(compare CompareFunc[T]) Slice[T] {
	sort.SliceStable(
		s, 
		func(i, j int) bool {
			return compare(s[i], s[j]) < 0
		},
	)
	return s
}