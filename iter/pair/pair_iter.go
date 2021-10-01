package pair

type Iter[K comparable, V any] interface {
	Next() bool
	Get() (K, V)
}