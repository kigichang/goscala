package maps

import (
	gs "github.com/kigichang/goscala"
	"github.com/kigichang/goscala/slices"
)

func Make[K comparable, V any](a ...int) gs.Map[K, V] {
	return gs.MkMap[K, V](a...)
}

func Empty[K comparable, V any]() gs.Map[K, V] {
	return Make[K, V]()
}

func From[K comparable, V any](pairs ...gs.Pair[K, V]) gs.Map[K, V] {
	m := Make[K, V](len(pairs))

	for i := range pairs {
		m.Add(pairs[i])
	}
	return m
}

func Collect[K comparable, V, T any](m gs.Map[K, V], pf func(K, V) (T, bool)) gs.Slice[T] {
	return slices.Collect(
		m.Slice(), 
		func(p gs.Pair[K, V]) (T, bool) {
			return pf(p.Key(), p.Value())
		},
	)
}

func CollectMap[K1, K2 comparable, V1, V2 any](m gs.Map[K1, V1], pf func(K1, V1)(K2, V2, bool)) gs.Map[K2, V2] {
	ret := Make[K2, V2]()

	iter := m.Range()

	for iter.Next() {
		if k2, v2, ok := pf(iter.Get()); ok {
			ret.Put(k2, v2)
		}
	}
	return ret
}

func CollectFirst[K comparable, V, T any](m gs.Map[K, V], pf func(K, V) (T, bool)) gs.Option[T] {
	iter := m.Range()

	for iter.Next() {
		if v, ok := pf(iter.Get()); ok {
			gs.Some[T](v)
		}
	}
	return gs.None[T]()
}

func FlatMapSlice[K comparable, V, T any](m gs.Map[K, V], fn func(K, V) gs.Sliceable[T]) gs.Slice[T] {
	ret := slices.Empty[T]()

	iter := m.Range()
	for iter.Next() {
		ret = append(ret, fn(iter.Get()).Slice()...)
	}
	return ret
}