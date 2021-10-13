package goscala

import (
	"context"
	"fmt"
	"time"
)

type Future[T any] interface {
	fmt.Stringer
	Completed() bool
	//Context() context.Context
	//Value() Try[T]
	PassValue() context.Context
	OnComplete(func(Try[T]))
	Foreach(func(T))
	Wait()
	Result(time.Duration) (T, error)
	Filter(context.Context, func(T) bool) Future[T]
}
