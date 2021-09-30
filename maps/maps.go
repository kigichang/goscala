package maps

import (
	gs "github.com/kigichang/goscala"
)

func Make[K comparable, V any](a ...int) gs.Map[K, V] {
	return gs.MkMap[K, V](a...)
}

func From[K comparable, V any](pairs ...gs.Pair[K, V]) gs.Map[K, V] {
	m := Make[K, V](len(pairs))

	for i := range pairs {
		m.Add(pairs[i])
	}
	return m
}