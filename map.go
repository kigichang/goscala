package goscala

type Map[K comparable, V any] map[K]V

func (m Map[K, V]) Len() int {
	return len(m)
}

func (m Map[K, V]) Slice() Slice[Pair[K, V]] {
	// TODO:
	return nil
}