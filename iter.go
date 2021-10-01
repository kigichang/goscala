package goscala

type Iter[T any] interface {
	Next() bool
	Get() T
}

type PairIter[K comparable, V any] interface {
	Next() bool
	Get() (K, V)
}