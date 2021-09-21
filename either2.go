package goscala

import (
	"fmt"
)

type Either2[L, R any] interface {
	fmt.Stringer
	Fetcher[R]

	IsLeft() bool
	IsRight() bool

	Left() L
	Right() R
}