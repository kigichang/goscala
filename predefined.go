package goscala

import "fmt"

type Fetcher[T any] interface {
	Fetch() (T, bool)
}

type Unit struct{}

func UnitFunc() Unit {
	return Unit{}
}

func UnitWrap[T any](f func(T)) func(T) Unit {
	return func(v T) Unit {
		f(v)
		return Unit{}
	}
}

var (
	ErrUnsupported = fmt.Errorf("unsupported")
)
