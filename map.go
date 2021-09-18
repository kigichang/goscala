package goscala

type Map[K comparable, V any] map[K]V

type Set[K comparable] Map[K, bool]

func (m Map[K, V]) Len() int {
	return len(m)
}

func (m Map[K, V]) Slice() Slice[Pair[K, V]] {
	ret := MakeSlice[Pair[K, V]](m.Len())
	i := 0

	for key := range m {
		ret[i] = P(key, m[key])
		i++
	}

	return ret
}

func (m Map[K, V]) Contains(key K) (ok bool) {
	_, ok = m[key]
	return ok
}

func (m Map[K, V]) Count(p func(K, V) bool) (ret int) {
	for key := range m {
		if p(key, m[key]) {
			ret++
		}
	}
	return
}

func (m Map[K, V]) Find(p func(K, V) bool) Option[Pair[K, V]] {
	for key := range m {
		if p(key, m[key]) {
			return Some[Pair[K, V]](P(key, m[key]))
		}
	}
	return None[Pair[K, V]]()
}


func (m Map[K, V]) Exists(p func(K, V) bool) bool {
	return m.Find(p).IsDefined()
}

func (m Map[K, V]) Filter(p func(K, V) bool) Map[K, V] {
	ret := MapEmpty[K, V]()

	for key := range m {
		if p(key, m[key]) {
			ret[key] = m[key]
		}
	}

	return ret
}

func (m Map[K, V]) FilterNot(p func(K, V) bool) Map[K, V] {
	p2 := func(k K, v V) bool {
		return !p(k, v)
	}

	return m.Filter(p2)
}

func (m Map[K, V]) Forall(p func(K, V) bool) bool {
	for key := range m {
		if !p(key, m[key]) {
			return false
		}
	}

	return true
}

func (m Map[K, V]) Foreach(fn func(K, V)) {
	for key := range m {
		fn(key, m[key])
	}
}

func (m Map[K, V]) Get(key K) Option[V] {
	if v, ok := m[key]; ok {
		return Some[V](v)
	}
	return None[V]()
}

func (m Map[K, V]) GetOrElse(key K, z V) V {
	if v, ok := m[key]; ok {
		return v
	}
	return z
}

func (m Map[K, V]) Keys() Slice[K] {
	return MapMapSlice(m, func(k K, _ V) K {
		return k
	})
}

func (m Map[K, V]) KeySet() Set[K] {
	return SliceToSet(m.Keys())
}

func (m Map[K, V]) Partition(fn func(K, V) bool) (Map[K, V], Map[K, V]) {
	a, b := m.Slice().Partition(func(p Pair[K, V]) bool {
		return fn(p.Key(), p.Value())
	})

	return SliceToMap(a), SliceToMap(b)
}