package slices

import (
	"sort"

	gs "github.com/kigichang/goscala"
	//"github.com/kigichang/goscala/iter"
)

type Slice[T any] []T

//var _ gs.Seq[int] = Slice[int]{}

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
	return s
}

func (s Slice[T]) Get(i int) T {
	return s[i]
}

func (s Slice[T]) Range() gs.Iterator[T] {
	return newSliceIterator(&s)
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

func (s Slice[T]) Head() (ret T, ok bool) {
	if ok = !s.IsEmpty(); ok {
		ret = s[0]
	}
	return
}

func (s Slice[T]) Last() (ret T, ok bool) {
	if ok = !s.IsEmpty(); ok {
		ret = s[s.Len()-1]
	}
	return
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
	_, ok = s.Find(p)
	return
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

func (s Slice[T]) Find(p func(T) bool) (ret T, ok bool) {
	for i := range s {
		if p(s[i]) {
			ret = s[i]
			ok = true
			return
		}
	}
	return
}

func (s Slice[T]) FindLast(p func(T) bool) (ret T, ok bool) {
	size := s.Len()
	for i := size - 1; i >= 0; i-- {
		if ok = p(s[i]); ok {
			ret = s[i]
			return
		}
	}
	return
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

func (s Slice[T]) ReduceRight(fn func(T, T) T) (ret T, ok bool) {
	size := s.Len()
	if size <= 0 {
		return
	}

	if ok = (size == 1); ok {
		ret = s[0]
		return
	}

	return FoldRight(s[:size-1], s[size-1], fn), true
}

func (s Slice[T]) ReduceLeft(fn func(T, T) T) (ret T, ok bool) {
	size := s.Len()
	if size <= 0 {
		return
	}

	if ok = (size == 1); ok {
		ret = s[0]
		return
	}

	return FoldLeft(s[1:], s[0], fn), true
}

func (s Slice[T]) Reduce(fn func(T, T) T) (T, bool) {
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

func (s Slice[T]) Max(compare func(T, T) int) (ret T, ok bool) {
	size := s.Len()
	if size == 0 {
		return
	}

	if ok = (size == 1); ok {
		ret = s[0]
		return
	}

	v := s[0]

	for i := 1; i < size; i++ {
		if compare(v, s[i]) < 0 {
			v = s[i]
		}
	}
	return v, true
}

func (s Slice[T]) Min(compare func(T, T) int) (T, bool) {
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
