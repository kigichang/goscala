package opt

import (
	"github.com/kigichang/goscala"
	"github.com/kigichang/goscala/monad"
)

func noneBool[T any](b bool) goscala.Option[T] {
	if !b {
		return goscala.None[T]()
	}
	panic("can not make None with true")
}

func noneErr[T any](err error) goscala.Option[T] {
	if err != nil {
		return goscala.None[T]()
	}
	panic("can not make None with nil error")
}

func MakeWithBool[T any](v T, ok bool) goscala.Option[T] {
	return monad.Fold[T, bool, goscala.Option[T]](goscala.IdentityBool(v, ok))(noneBool[T], goscala.Some[T])
}

func MakeWithErr[T any](v T, err error) goscala.Option[T] {
	return monad.Fold[T, error, goscala.Option[T]](goscala.IdentityErr(v, err))(noneErr[T], goscala.Some[T])
}
