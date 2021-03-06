package maps

import (
	"github.com/kigichang/goscala/iter/pair"

	gs "github.com/kigichang/goscala"
	"reflect"
)

type generalIter[K comparable, V any] struct {
	len  func() int
	iter *reflect.MapIter
}

func (i *generalIter[K, V]) Len() int {
	return i.len()
}

func (i *generalIter[K, V]) Next() bool {
	return i.iter.Next()
}

func (i *generalIter[K, V]) Get() (K, V) {
	k := i.iter.Key().Interface()
	v := i.iter.Value().Interface()

	return k.(K), v.(V)
}

func newGeneralIter[K comparable, V any](m map[K]V) *generalIter[K, V] {
	return &generalIter[K, V]{
		len: func() int {
			return len(m)
		},
		iter: reflect.ValueOf(m).MapRange(),
	}
}

type _map[K comparable, V any] map[K]V

func (m _map[K, V]) Range() pair.Iter[K, V] {
	return newGeneralIter[K, V](m)
}

type generalMap[K comparable, V any] struct {
	*abstractMap[K, V]
}

func (g *generalMap[K, V]) Add(p gs.Pair[K, V]) {
	if p == nil {
		return
	}
	g.Put(p.Key(), p.Value())
}

func (g *generalMap[K, V]) Put(k K, v V) {
	m := g.abstractMap.mapRanger.(_map[K, V])
	m[k] = v
}

func (g *generalMap[K, V]) Get(k K) (ret V, ok bool) {
	m := g.abstractMap.mapRanger.(_map[K, V])
	ret, ok = m[k]
	return
}

func (g *generalMap[K, V]) Contains(k K) (ok bool) {
	_, ok = g.Get(k)
	return
}

func (g *generalMap[K, V]) Delete(k K) {
	m := g.abstractMap.mapRanger.(_map[K, V])
	delete(m, k)
}

func (g *generalMap[K, V]) GetOrElse(k K, z V) V {
	v, ok := g.Get(k)
	if ok {
		return v
	}
	return z
}

func (g *generalMap[K, V]) Partition(fn func(K, V) bool) (gs.Map[K, V], gs.Map[K, V]) {
	a, b := g.partition(fn)
	return formGeneralMap(a...), formGeneralMap(b...)
}

func (g *generalMap[K, V]) Filter(fn func(K, V) bool) gs.Map[K, V] {
	return formGeneralMap(g.filter(fn)...)
}

func (g *generalMap[K, V]) FilterNot(fn func(K, V) bool) gs.Map[K, V] {
	return formGeneralMap(g.filterNot(fn)...)
}

func newGeneralMap[K comparable, V any](a ...int) *generalMap[K, V] {

	size := 0
	if len(a) > 0 {
		size = a[0]
	}

	m := make(map[K]V, size)
	return &generalMap[K, V]{
		&abstractMap[K, V]{
			_map[K, V](m),
		},
	}
}

func formGeneralMap[K comparable, V any](s ...gs.Pair[K, V]) *generalMap[K, V] {
	m := newGeneralMap[K, V]()
	for i := range s {
		m.Add(s[i])
	}
	return m
}
