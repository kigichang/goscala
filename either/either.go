package either

import (
	"github.com/kigichang/goscala"
	"github.com/kigichang/goscala/monad"
)
func Bool[R any](v R, ok bool) goscala.Either[bool, R] {
	return monad.Fold[R, bool, goscala.Either[bool, R]](goscala.ValueBoolFunc(v, ok))(
		goscala.Left[bool, R],
		goscala.Right[bool, R],
	)
}

func Err[R any](v R, err error) goscala.Either[error, R] {
	return monad.Fold[R, error, goscala.Either[error, R]](goscala.ValueErrFunc(v, err))(
		goscala.Left[error, R],
		goscala.Right[error, R],
	)
}