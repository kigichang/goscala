package goscala

type Fetcher[T any] interface {
	Fetch() (T, bool)
	FetchErr() (T, error)
}