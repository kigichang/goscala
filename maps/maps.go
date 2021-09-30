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
	s := slices.Collect(
		m.Slice(),
		func(p gs.Pair[K1, V1]) (gs.Pair[K2, V2], bool) {
			if k, v, ok := pf(p.Key(), p.Value()); ok {
				return gs.P(k, v), true
			}
			return nil, false
		},
	)

	return From[K2, V2](s...)
}