package goscala

func MakePair[K comparable, V any](k K, v V) Pair[K, V] {
	return P[K, V](k, v)
}

func P[K comparable, V any](k K, v V) Pair[K, V] {
	
	return &pair[K, V] {
		tup: &tuple2[K, V] {
		v1: k,
		v2: v,
	}}
}