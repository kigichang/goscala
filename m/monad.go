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

func FoldLeft[T, U any](s []T, z U, fn func(a U, b T) U) U {
	zz := z
	for i := range s {
		zz = fn(zz, s[i])
	}
	return zz
}

func ScanLeft[T, U any](s []T, z U, fn func(U, T) U) []U {
	return FoldLeft[T, []U](s, []U{z}, func(a []U, b T) []U {
		return append(a, fn(a[len(a)-1], b))
	})
}

func FoldRight[T, U any](s []T, z U, fn func(T, U) U) U {
	zz := z
	size := len(s)
	for i := size - 1; i >= 0; i-- {
		zz = fn(s[i], zz)
	}
	return zz
}

func ScanRight[T, U any](s []T, z U, fn func(T, U) U) []U {
	result := FoldRight[T, []U](s, []U{z}, func(a T, b []U) []U {
		return append(b, fn(a, b[len(b)-1]))
	})

	size := len(result)
	half := size / 2

	for i := 0; i < half; i++ {
		tmp := result[i]
		result[i] = result[size-1-i]
		result[size-1-i] = tmp
	}

	return result
}

func Map[T, U any](s []T, fn func(T) U) []U {
	return FoldLeft[T, []U](s, []U{}, func(z []U, a T) []U {
		z = append(z, fn(a))
		return z
	})
}

func FlatMap[T, U any](s []T, fn func(T) []U) []U {
	return FoldLeft[T, []U](s, []U{}, func(z []U, a T) []U {
		z = append(z, fn(a)...)
		return z
	})

}
