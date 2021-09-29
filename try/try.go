package try

import (
	m "github.com/kigichang/gomonad"
	"github.com/kigichang/goscala"
	s "github.com/kigichang/goscala"
)

func Err[T any](v T, err error) s.Try[T] {
	return m.FoldErr[T, goscala.Try[T]](m.VF2(v, err))(
		s.Failure[T],
		s.Success[T],
	)
}

func Bool[T any](v T, ok bool) goscala.Try[T] {
	return m.FoldBool[T, s.Try[T]](m.VF2(v, ok))(
		m.FuncUnitAndThen[error, s.Try[T]](m.VF(s.ErrUnsatisfied))(s.Failure[T]),
		s.Success[T],
	)
}
