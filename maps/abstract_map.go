package maps

import (
	gs "github.com/kigichang/goscala"
	"github.com/kigichang/goscala/iter/pair"
)

type mapRanger[K comparable, V any] interface {
	Range() pair.Iter[K, V]
}

type abstractMap[K comparable, V any] struct {
	mapRanger[K, V]
}

func (m *abstractMap[K, V]) Len() int {
	return m.Range().Len()
}

func (m *abstractMap[K, V]) Keys() gs.Slice[K] {
	it := m.Range()
	ret := make([]K, 0, it.Len())

	for it.Next() {
		k, _ := it.Get()
		ret = append(ret, k)
	}
	return ret
}

func (m *abstractMap[K, V]) Values() gs.Slice[V] {
	it := m.Range()
	ret := make([]V, 0, it.Len())

	for it.Next() {
		_, v := it.Get()
		ret = append(ret, v)
	}
	return ret
}

func (m *abstractMap[K, V]) Contains(k K) bool {
	it := m.Range()
	for it.Next() {
		if k1, _ := it.Get(); k1 == k {
			return true
		}
	}
	return false
}

func (m *abstractMap[K, V]) Count(p func(K, V) bool) int {
	return pair.Count(m.Range(), p)
}

func (m *abstractMap[K, V]) Find(p func(K, V) bool) gs.Option[gs.Pair[K, V]] {
	if k, v, ok := pair.Find(m.Range(), p); ok {
		return gs.Some[gs.Pair[K, V]](gs.P(k, v))
	}
	return gs.None[gs.Pair[K, V]]()
}

func (m *abstractMap[K, V]) Exists(p func(K, V) bool) bool {
	return pair.Exists(m.Range(), p)
}

func (m *abstractMap[K, V]) filter(p func(K, V) bool) gs.Slice[gs.Pair[K, V]] {
	ret := []gs.Pair[K, V]{}
	it := m.Range()

	for it.Next() {
		if k, v := it.Get(); p(k, v) {
			ret = append(ret, gs.P(k, v))
		}
	}
	return ret
}

func (m *abstractMap[K, V]) filterNot(p func(K, V) bool) gs.Slice[gs.Pair[K, V]] {
	return m.filter(func(k K, v V) bool {
		return !p(k, v)
	})
}

func (m *abstractMap[K, V]) Forall(p func(K, V) bool) bool {
	it := m.Range()
	for it.Next() {
		if !p(it.Get()) {
			return false
		}
	}
	return true
}

func (m *abstractMap[K, V]) Foreach(fn func(K, V)) {
	it := m.Range()
	for it.Next() {
		fn(it.Get())
	}
}

func (m *abstractMap[K, V]) partition(p func(K, V) bool) (gs.Slice[gs.Pair[K, V]], gs.Slice[gs.Pair[K, V]]) {
	a, b := []gs.Pair[K, V]{}, []gs.Pair[K, V]{}
	it := m.Range()

	for it.Next() {
		if k, v := it.Get(); p(k, v) {
			b = append(b, gs.P(k, v))
		} else {
			a = append(a, gs.P(k, v))
		}
	}

	return a, b
}

func (m *abstractMap[K, V]) Slice() gs.Slice[gs.Pair[K, V]] {
	it := m.Range()
	ret := make([]gs.Pair[K, V], 0, it.Len())
	for it.Next() {
		ret = append(ret, gs.P(it.Get()))
	}
	return ret
}
