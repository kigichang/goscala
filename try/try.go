package try

import (
	gs "github.com/kigichang/goscala"
)

func Err[T any](v T, err error) gs.Try[T] {
	return gs.PFErrV(
		gs.Success[T],
		gs.Failure[T],
	)(v, err)
}

func Bool[T any](v T, ok bool) gs.Try[T] {
	return Err(v, gs.Cond(ok, nil, gs.ErrUnsatisfied))
}
