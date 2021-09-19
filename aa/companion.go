package aa

import (
	"github.com/kigichang/goscala"
	"github.com/kigichang/goscala/intern"
)

func AA[T any](v T) goscala.A[T] {
	return intern.AA(v)
}