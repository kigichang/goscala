package goscala


func MonadFold[T, U any](fetch FetchFunc[T]) Func2[U, Func1[T, U], U] {
	return func(z U, fn Func1[T, U]) U {
		v, ok := fetch()
		if ok {
			return fn(v)
		}
		return z
	}
}

func MonadMaker[T, U any](v T) Func2[U, Func1[T, U], U] {
	return MonadFold[T, U](ZeroFunc[T](v))
}

func MonadFoldLeft[T, U any](s SliceFunc[T]) Func2[U, Func2[U, T, U], U] {
	return func(z U, fn Func2[U, T, U]) U {
		ss := s()
		for i := range ss {
			z = fn(z, ss[i])
		}
		return z
	}
}

//func MonadFoldLeft[T, U any](s Slice[T], z U, fn Func1[U, T, U]) U {
//	for i := range s {
//		z = fn(z, s[i])
//	}
//	return z
//}

func MonadFlatMap[R, X any](cond Condition, get Func[R], op Func1[R, X], flat Func1[X, X], fa Func[X]) X {
	if !cond() {
		return fa()
	}

	return flat(op(get()))	
}

func MonadMap[R, T, X any](cond Condition, get Func[R], op Func1[R, T], fa Func[X], fb Func1[T, X]) X {
	if cond() {
		return fb(op(get()))
	}
	return fa()
}