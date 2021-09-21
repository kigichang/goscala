package goscala

func makeTry[T any](v T, err error) (ret *try[T]) {

	e := either[error, T]{
		right: err == nil,
		rv:    v,
		lv:    err,
	}

	t := try[T](e)
	return &t
}

func Success[T any](v T) Try[T] {
	return makeTry(v, nil)
}

func Failure[T any](err error) Try[T] {
	if err == nil {
		err = ErrNil
	}

	var x T
	return makeTry(x, err)
}

func MakeTry[T any](v T, err error) Try[T] {
	if err != nil {
		return Failure[T](err)
	}
	return Success[T](v)
}

func MakeTryWithBool[T any](v T, ok bool) Try[T] {
	if ok {
		return Success[T](v)
	}
	return Failure[T](ErrUnsatisfied)
}

func TryCollect[T, U any](t Try[T], pf PartialFunc[T, U]) Try[U] {
	if t.IsSuccess() {
		return MakeTryWithBool(pf(t.Get()))
	}
	return Failure[U](t.Failed())
}

func TryFlatMap[T, U any](t Try[T], f Func1[T, Try[U]]) Try[U] {
	if t.IsSuccess() {
		return f(t.Get())
	}
	return Failure[U](t.Failed())
}

func TryFold[T, U any](t Try[T], succ Func1[T, U], fail Func1[error, U]) U {
	if t.IsSuccess() {
		return succ(t.Get())
	}
	return fail(t.Failed())
}

func TryMap[T, U any](t Try[T], f Func1[T, U]) Try[U] {
	if t.IsSuccess() {
		return Success[U](f(t.Get()))
	}

	return Failure[U](t.Failed())
}

func TryMapWithErr[T, U any](t Try[T], f Func1Err[T, U]) Try[U] {
	if t.IsSuccess() {
		return MakeTry(f(t.Get()))
	}

	return Failure[U](t.Failed())
}

func TryMapWithBool[T, U any](t Try[T], pf PartialFunc[T, U]) Try[U] {
	if t.IsSuccess() {
		return MakeTryWithBool(pf(t.Get()))
	}
	return Failure[U](t.Failed())
}

func TryTransform[T, U any](t Try[T], succ Func1[T, Try[U]], fail Func1[error, Try[U]]) Try[U] {
	if t.IsSuccess() {
		return succ(t.Get())
	}

	return fail(t.Failed())
}