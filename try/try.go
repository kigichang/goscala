package try

import (
	"github.com/kigichang/goscala"
	gs "github.com/kigichang/goscala"
)

func Err[T any](v T, err error) gs.Try[T] {
	return gs.FoldErr[T, goscala.Try[T]](gs.VF2(v, err))(
		gs.Failure[T],
		gs.Success[T],
	)
}

func Bool[T any](v T, ok bool) gs.Try[T] {
	return gs.FoldBool[T, gs.Try[T]](gs.VF2(v, ok))(
		gs.FuncUnitAndThen[error, gs.Try[T]](gs.VF(gs.ErrUnsatisfied))(gs.Failure[T]),
		gs.Success[T],
	)
}
