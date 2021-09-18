package goscala

func MapEmpty[K comparable, V any]() Map[K, V] {
	return map[K]V{}
}

func MakeMap[K comparable, V any](size int) Map[K, V] {
	return make(map[K]V, size)
}

func MapFrom[K comparable, V any](pairs ...Pair[K, V]) Map[K, V] {
	ret := MakeMap[K, V](len(pairs))

	for i := range pairs {
		ret[pairs[i].Key()] = pairs[i].Value()
	}
	return ret
}

func MapCollect[K comparable, V, T any](m Map[K, V], pf func(K, V) (T, bool)) Slice[T] {
	ret := SliceEmpty[T]()
	for key := range m {
		if v, ok := pf(key, m[key]); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func MapColectMap[K1, K2 comparable, V1, V2 any](m Map[K1, V1], pf func(K1, V1)(K2, V2, bool)) Map[K2, V2] {
	ret := MapEmpty[K2, V2]()

	for key := range m {
		if k2, v2, ok := pf(key, m[key]); ok {
			ret[k2] = v2
		}
	}

	return ret
}

func MapCollectFirst[K comparable, V, T any](m Map[K, V], pf func(K, V)(T, bool)) Option[T] {
	for key := range m {
		if v, ok := pf(key, m[key]); ok {
			return Some[T](v)
		}
	}
	return None[T]()
}

func MapFlatMapSlice[K comparable, V, T any](m Map[K, V], fn func(K, V) Sliceable[T]) Slice[T] {
	ret := SliceEmpty[T]()
	for key := range m {
		ret = append(ret, fn(key, m[key]).Slice()...)
	}
	return ret
}

func MapFlatMap[K1, K2 comparable, V1, V2 any](m Map[K1, V1], fn func(K1, V1) Sliceable[Pair[K2, V2]]) Map[K2, V2] {
	return SliceToMap(MapFlatMapSlice(m, fn))
}

func MapGroupMap[K1, K2 comparable, V1, V2 any](m Map[K1, V1], groupBy func(K1, V1) K2, op func(K1, V1) V2) Map[K2, Slice[V2]] {
	ret := MapEmpty[K2, Slice[V2]]()

	for key := range m {
		key2 := groupBy(key, m[key])
		val2 := op(key, m[key])
		ret[key2] = append(ret[key2], val2)
	}
	return ret
}

func MapGroupBy[K, K1 comparable, V any](m Map[K, V], groupBy func(K, V) K1) Map[K1, Map[K, V]] {
	op := func(k K, v V) Pair[K, V] {
		return P(k, v)
	}

	m1 := MapGroupMap(m, groupBy, op)

	ret := MakeMap[K1, Map[K, V]](m1.Len())

	for k1 := range m1 {
		ret[k1] = SliceToMap(m1[k1])
	}
	return ret
}

func MapGroupMapReduce[K1, K2 comparable, V1, V2 any](m Map[K1, V1], groupBy func(K1, V1) K2, op func(K1, V1) V2, reduce func(V2, V2) V2) Map[K2, V2] {
	m2 := MapGroupMap(m, groupBy, op)

	ret := MakeMap[K2, V2](m2.Len())

	for key := range m2 {
		ret[key]= m2[key].Reduce(reduce).Get()
	}
	return ret
}

func MapMap[K1, K2 comparable, V1, V2 any](m Map[K1, V1], fn func(K1, V1) (K2, V2)) Map[K2, V2] {
	ret := MakeMap[K2, V2](m.Len())

	for key := range m {
		k2, v2 := fn(key, m[key])
		ret[k2] = v2
	}

	return ret
}

func MapMapSlice[K comparable, V, T any](m Map[K, V], fn func(K, V) T) Slice[T] {
	ret := MakeSlice[T](m.Len())

	for key := range m {
		ret = append(ret, fn(key, m[key]))
	}

	return ret
}

func MapMax[K comparable, V any, B Ordered](m Map[K, V], fn func(K, V) B) Option[Pair[K, V]] {
	fn1 := func(p Pair[K, V]) B {
		return fn(p.Key(), p.Value())
	}

	return SliceMaxBy(m.Slice(), fn1)
}

func MapMin[K comparable, V any, B Ordered](m Map[K, V], fn func(K, V) B) Option[Pair[K, V]] {
	fn1 := func(p Pair[K, V]) B {
		return fn(p.Key(), p.Value())
	}
	return SliceMinBy(m.Slice(), fn1)
}

func MapPartitionMap[K comparable, V, A, B any](m Map[K, V], fn func(K, V) Either[A, B]) (Slice[A], Slice[B]) {
	return SlicePartitionMap(m.Slice(), func(p Pair[K, V]) Either[A, B] {
		return fn(p.Key(), p.Value())
	})
}