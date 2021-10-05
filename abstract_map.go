package goscala

import (
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

func (m *abstractMap[K, V]) Keys() Slice[K] {
	it := m.Range()
	ret := make([]K, 0, it.Len())

	for it.Next() {
		k, _ := it.Get()
		ret = append(ret, k)
	}
	return ret
}

func (m *abstractMap[K, V]) Values() Slice[V] {
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

func (m *abstractMap[K, V]) Find(p func(K, V) bool) Option[Pair[K, V]] {
	if k, v, ok := pair.Find(m.Range(), p); ok {
		return Some[Pair[K, V]](P(k, v))
	}
	return None[Pair[K, V]]()
}

func (m *abstractMap[K, V]) Exists(p func(K, V) bool) bool {
	return pair.Exists(m.Range(), p)
}

func (m *abstractMap[K, V]) Filter(p func(K, V) bool) Slice[Pair[K, V]] {
	ret := []Pair[K, V]{}
	it := m.Range()

	for it.Next() {
		if k, v := it.Get(); p(k, v) {
			ret = append(ret, P(k, v))
		}
	}
	return ret
}

func (m *abstractMap[K, V]) FilterNot(p func(K, V) bool) Slice[Pair[K, V]] {
	return m.Filter(func(k K, v V) bool {
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

func (m *abstractMap[K, V]) Partition(p func(K, V) bool) (Slice[Pair[K, V]], Slice[Pair[K, V]]) {
	a, b := []Pair[K, V]{}, []Pair[K, V]{}
	it := m.Range()

	for it.Next() {
		if k, v := it.Get(); p(k, v) {
			b = append(b, P(k, v))
		} else {
			a = append(a, P(k, v))
		}
	}

	return a, b
}

func (m *abstractMap[K, V]) Slice() Slice[Pair[K, V]] {
	it := m.Range()
	ret := make([]Pair[K, V], 0, it.Len())
	for it.Next() {
		ret = append(ret, P(it.Get()))
	}
	return ret
}
