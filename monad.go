package goscala

func MonadFold[T, U any](z U, pf FuncBool[T], fn Func1[T, U]) U {
	v, ok := pf()
	if ok {
		return fn(v)
	}
	return z
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