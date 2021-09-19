package goscala

func Some[T any](v T) Option[T] {
	return &option[T]{
		defined: true,
		v:       v,
	}
}

func None[T any]() Option[T] {
	return &option[T]{
		defined: false,
	}
}

func MakeOption[T any](v T) Option[T] {
	return &option[T]{
		defined: !IsZero(v),
		v:       v,
	}
}

func MakeOptionWithBool[T any](v T, ok bool) Option[T] {
	if ok {
		return Some[T](v)
	}
	return None[T]()
}

func MakeOptionWithErr[T any](v T, err error) Option[T] {
	if err != nil {
		return None[T]()
	}
	return Some[T](v)
}

func OptionWhen[T any](cond Condition, v T) Option[T] {
	if cond() {
		return Some[T](v)
	}
	return None[T]()
}

func OptionFlatMap[T, U any](o Option[T], f Func1[T, Option[U]]) Option[U] {
	if o.IsDefined() {
		return f(o.Get())
	}
	return None[U]()
}

func OptionFold[T, U any](o Option[T], z U, f Func1[T, U]) U {
	if o.IsDefined() {
		return f(o.Get())
	}
	return z
}

func OptionMap[T, U any](o Option[T], f Func1[T, U]) Option[U] {
	if o.IsDefined() {
		return Some[U](f(o.Get()))
	}
	return None[U]()
}

func OptionMapWithErr[T, U any](o Option[T], f Func1Err[T, U]) Option[U] {
	if o.IsDefined() {
		return MakeOptionWithErr[U](f(o.Get()))
	}
	return None[U]()
}

func OptionMapWithBool[T, U any](o Option[T], f Func1Bool[T, U]) Option[U] {
	if o.IsDefined() {
		return MakeOptionWithBool[U](f(o.Get()))
	}
	return None[U]()
}

func OptionToLeft[T, R any](o Option[T], z R) Either[T, R] {
	if o.IsDefined() {
		return Left[T, R](o.Get())
	}
	return Right[T, R](z)
}

func OptionToRight[L, T any](o Option[T], z L) Either[L, T] {
	if o.IsDefined() {
		return Right[L, T](o.Get())
	}
	return Left[L, T](z)
}

func OptionToTry[T, U any](o Option[T], fn Func1[T, U]) Try[U] {
	if o.IsDefined() {
		return Success[U](fn(o.Get()))
	}
	return Failure[U](ErrNone)
}

func OptionZip[T, U any](o Option[T], that Option[U]) Option[Tuple2[T, U]] {
	if o.IsDefined() && that.IsDefined() {
		return Some[Tuple2[T, U]](MakeTuple2[T, U](o.Get(), that.Get()))

	}
	return None[Tuple2[T, U]]()
}

func OptionUnzip[T1, T2 any](o Option[Tuple2[T1, T2]]) (Option[T1], Option[T2]) {
	if o.IsDefined() {
		return Some[T1](o.Get().V1()), Some[T2](o.Get().V2())
	}
	return None[T1](), None[T2]()
}

func OptionUnzip3[T1, T2, T3 any](o Option[Tuple3[T1, T2, T3]]) (Option[T1], Option[T2], Option[T3]) {
	if o.IsDefined() {
		return Some[T1](o.Get().V1()), Some[T2](o.Get().V2()), Some[T3](o.Get().V3())
	}
	return None[T1](), None[T2](), None[T3]()
}

func OptionUnless[T any](cond Condition, v T) Option[T] {
	if !cond() {
		return Some[T](v)
	}
	return None[T]()
}

func OptionCollect[T, U any](o Option[T], pf PartialFunc[T, U]) Option[U] {
	if o.IsDefined() {
		return MakeOptionWithBool[U](pf(o.Get()))
	}
	return None[U]()
}