package goscala

type Pair[K comparable, V any] interface {
	Key() K
	Value() V
	Get() (K, V)
}


type pair[K comparable, V any] struct {
	key K
	val V
}

var _ Pair[int, int] = &pair[int, int]{}

func (p *pair[K, V]) Key() K { 
	return p.key
}

func (p *pair[K, V]) Value() V {
	return p.val
}

func (p *pair[K, V]) Get() (K, V) {
	return p.key, p.val
}

func P[K comparable, V any](k K, v V) Pair[K, V] {
	return &pair[K, V] { 
		key: k, 
		val: v,
	}
}