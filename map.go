package goscala

type Map[K comparable, V any] interface {
	Keys() Slice[K]
	Values() Slice[V]
}
