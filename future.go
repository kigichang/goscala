package goscala

import (
	"context"
	"fmt"
	"time"
)

type Future[T any] interface {
	fmt.Stringer
	Completed() bool
	PassValue() context.Context
	OnComplete(func(Try[T]))
	Foreach(func(T))
	Wait()
	Result(time.Duration) (T, error)
	Filter(context.Context, func(T) bool) Future[T]
}
