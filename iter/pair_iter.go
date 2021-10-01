package iter

type PairIter[K comparable, V any] interface {
	Next() bool
	Get() (K, V)
}