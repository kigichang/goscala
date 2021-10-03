package goscala

import (
	"context"
	"fmt"
)

type Future[T any] interface {
	fmt.Stringer
	Completed() bool
	Context() context.Context
	Value() Try[T]
	Pass() context.Context
	Foreach(func(T))
	OnComplete(func(Try[T]))
	Wait()
}