package maps

import (
	"constraints"

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

func CollectMap[K1, K2 comparable, V1, V2 any](m gs.Map[K1, V1], pf func(K1, V1) (K2, V2, bool)) gs.Map[K2, V2] {
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

func FlatMap[K1, K2 comparable, V1, V2 any](m gs.Map[K1, V1], fn func(K1, V1) gs.Sliceable[gs.Pair[K2, V2]]) gs.Map[K2, V2] {
	return slices.ToMap(FlatMapSlice(m, fn))
}

func GroupMap[K1, K2 comparable, V1, V2 any](m gs.Map[K1, V1], groupBy func(K1, V1) K2, op func(K1, V1) V2) gs.Map[K2, gs.Slice[V2]] {
	ret := Make[K2, gs.Slice[V2]]()

	it := m.Range()

	for it.Next() {
		k, v := it.Get()
		k2 := groupBy(k, v)
		v2 := op(k, v)

		x, _ := ret.Get(k2)
		x = append(x, v2)
		ret.Put(k2, x)
	}

	return ret
}

func GroupBy[K, K1 comparable, V any](m gs.Map[K, V], groupBy func(K, V) K1) gs.Map[K1, gs.Map[K, V]] {
	op := func(k K, v V) gs.Pair[K, V] {
		return gs.P(k, v)
	}

	m1 := GroupMap(m, groupBy, op)

	ret := Make[K1, gs.Map[K, V]](m1.Len())

	it := m1.Range()

	for it.Next() {
		k, v := it.Get()
		ret.Put(k, slices.ToMap(v))
	}

	return ret
}

func GroupMapReduce[K1, K2 comparable, V1, V2 any](m gs.Map[K1, V1], groupBy func(K1, V1) K2, op func(K1, V1) V2, reduce func(V2, V2) V2) gs.Map[K2, V2] {
	m2 := GroupMap(m, groupBy, op)

	ret := Make[K2, V2](m2.Len())

	it := m2.Range()

	for it.Next() {
		k, v := it.Get()
		r, _ := v.Reduce(reduce)
		ret.Put(k, r)
	}

	return ret
}

func MapMap[K1, K2 comparable, V1, V2 any](m gs.Map[K1, V1], fn func(K1, V1) (K2, V2)) gs.Map[K2, V2] {
	ret := Make[K2, V2](m.Len())

	it := m.Range()

	for it.Next() {
		k2, v2 := fn(it.Get())
		ret.Put(k2, v2)
	}

	return ret
}

func MapSlice[K comparable, V, T any](m gs.Map[K, V], fn func(K, V) T) gs.Slice[T] {
	ret := slices.Make[T](m.Len())

	it := m.Range()

	for it.Next() {
		ret = append(ret, fn(it.Get()))
	}

	return ret
}

func MaxBy[K comparable, V any, B constraints.Ordered](m gs.Map[K, V], fn func(K, V) B) gs.Option[gs.Pair[K, V]] {
	fn1 := func(p gs.Pair[K, V]) B {
		return fn(p.Key(), p.Value())
	}

	return slices.MaxBy(m.Slice(), fn1)
}

func MinBy[K comparable, V any, B constraints.Ordered](m gs.Map[K, V], fn func(K, V) B) gs.Option[gs.Pair[K, V]] {
	fn1 := func(p gs.Pair[K, V]) B {
		return fn(p.Key(), p.Value())
	}
	return slices.MinBy(m.Slice(), fn1)
}

func PartitionMap[K comparable, V, A, B any](m gs.Map[K, V], fn func(K, V) gs.Either[A, B]) (gs.Slice[A], gs.Slice[B]) {
	return slices.PartitionMap(m.Slice(), func(p gs.Pair[K, V]) gs.Either[A, B] {
		return fn(p.Key(), p.Value())
	})
}
