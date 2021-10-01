package goscala

import (
	"reflect"
	"github.com/kigichang/goscala/iter/pair"
)

type Map[K comparable, V any] interface {
	Len() int
	Keys() Slice[K]
	Values() Slice[V]
	Add(Pair[K, V])
	Put(K, V)
	Get(K) (V, bool)

	Contains(K) bool
	Count(func(K, V) bool) int

	Find(func(K, V) bool) Option[Pair[K, V]]
	Exists(func(K, V) bool) bool
	Filter(func(K, V) bool) Map[K, V]
	FilterNot(func(K, V) bool) Map[K, V]
	Forall(func(K, V) bool) bool
	Foreach(func(K, V))
	Partition(func(K, V) bool) (Map[K, V], Map[K, V])
	GetOrElse(K, V) V

	Slice() Slice[Pair[K, V]]
	Range() pair.Iter[K, V]
}

type _mapIter[K comparable, V any] struct {
	iter *reflect.MapIter
}

func (i *_mapIter[K, V]) Next() bool {
	return i.iter.Next()
}

func (i *_mapIter[K, V]) Get() (K, V) {
	k := i.iter.Key().Interface()
	v := i.iter.Value().Interface()

	return k.(K), v.(V)
}

type _map[K comparable, V any] map[K]V

func (m _map[K, V]) Len() int {
	return len(m)
}

func (m _map[K, V]) Keys() Slice[K] {
	ret := make([]K, len(m))
	idx := 0
	for k := range m {
		ret[idx] = k
		idx++
	}
	return ret
}

func (m _map[K, V]) Values() Slice[V] {
	ret := make([]V, len(m))
	idx := 0
	for k := range m {
		ret[idx] = m[k]
		idx++
	}
	return ret
}

func (m _map[K, V]) Add(p Pair[K, V]) {
	if p == nil {
		return
	}
	m.Put(p.Key(), p.Value())
}

func(m _map[K, V]) Put(k K, v V) {
	m[k] = v
}

func (m _map[K, V]) Get(k K) (v V, ok bool) {
	v, ok = m[k]
	return
}

func (m _map[K, V]) Contains(k K) (ok bool) {
	_, ok = m[k]
	return 
}

func (m _map[K, V]) Count(p func(K, V) bool) (ret int) {
	for k := range m {
		if p(k, m[k]) {
			ret++
		}
	}
	return
}

func (m _map[K, V]) Find(p func(K, V) bool) Option[Pair[K, V]] {
	for k := range m {
		if p(k, m[k]) {
			return Some[Pair[K, V]](P(k, m[k]))
		}
	}
	return None[Pair[K, V]]()
}

func (m _map[K, V]) Exists(p func(K, V) bool) bool {
	return m.Find(p).IsDefined()
}

func (m _map[K, V]) Filter(p func(K, V) bool) Map[K, V] {
	ret := MkMap[K, V]()
	for key := range m {
		if p(key, m[key]) {
			ret.Put(key, m[key])
		}
	}
	return ret
}

func (m _map[K, V]) FilterNot(p func(K, V) bool) Map[K, V] {
	return m.Filter(func(k K, v V) bool {
		return !p(k, v)
	})
}

func (m _map[K, V]) Forall(p func(K, V) bool) bool {
	for key := range m {
		if !p(key, m[key]) {
			return false
		}
	}

	return true
}

func (m _map[K, V]) Foreach(fn func(K, V)) {
	for key := range m {
		fn(key, m[key])
	}
}

func (m _map[K, V]) Partition(p func(K, V) bool) (Map[K, V], Map[K, V]) {
	a, b := MkMap[K, V](), MkMap[K, V]()

	for k := range m {
		if p(k, m[k]) {
			b.Put(k, m[k])
		} else {
			a.Put(k, m[k])
		}
	}
	return a, b
}

func (m _map[K, V]) GetOrElse(key K, z V) V {
	return GetOrElse(m.Get(key))(z)
}

func (m _map[K, V]) Slice() Slice[Pair[K, V]] {
	ret := make([]Pair[K, V], len(m))

	i := 0
	for k := range m {
		ret[i] = P(k, m[k])
	}
	return ret
}

func (m _map[K, V]) Range() pair.Iter[K, V] {
	return &_mapIter[K, V] {
		iter: reflect.ValueOf(m).MapRange(),
	}
}

func MkMap[K comparable, V any](a ...int) Map[K, V] {
	size := 0
	if len(a) > 0 {
		size = a[0]
	}
	
	return _map[K, V](make(map[K]V, size))
}