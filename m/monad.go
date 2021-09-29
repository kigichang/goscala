package m

func Fold[T, C, U any](fetch func() (T, C), fa func(C) U, fb func(T) U) U {

	var v T
	var c C
	v, c = fetch()

	var x interface{} = c
	ok := false
	switch xv := x.(type) {
	case bool:
		ok = xv
	case error:
		ok = false
	default:
		ok = (xv == nil)
	}

	if ok {
		return fb(v)
	}
	return fa(c)

}

func FoldBool[T, U any](fetch func() (T, bool), z func() U, f func(T) U) U {
	if v, ok := fetch(); ok {
		return f(v)
	}
	return z()
}

func FoldErr[T, U any](fetch func() (T, error), z func(error) U, f func(T) U) U {
	return Fold[T, error, U](fetch, z, f)
}
