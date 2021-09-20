package monad

func Fold[T, C, U any](fetch func() (T, C)) func(func(C) U, func(T) U) U {
	return func(z func(C) U, f func(T) U) U {
		var v T
		var c C
		v, c = fetch()
		
		var x interface{} = c
		ok := false
		switch xv := x.(type) {
		case bool:
			ok = xv
		default:
			ok = (xv == nil)
		}


		if ok {
			return f(v)
		}
		return z(c)
	}
}

func FoldLeft[T, U any](s []T) func(U) func(func(U, T) U) U {
	return func(z U) func(func(U, T) U) U {
		return func (fn func(a U, b T) U) U {
			zz := z
			for i := range s {
				zz = fn(zz, s[i])
			}
			return zz
		}
	}
}

func FoldRight[T, U any](s []T) func(U) func(func(T, U) U) U {
	return func (z U) func(func(T, U) U) U {
		return func(fn func(T, U) U) U {
			zz := z
			size := len(s)
			for i := size -1; i >= 0; i-- {
				zz = fn(s[i], zz)
			}
			return zz
		}
	}
}

func Map[T, U any](s []T) func(func(T) U) []U {
	return func(fn func(T) U) []U {
		return FoldLeft[T, []U](s)([]U{})(func (z []U, a T) []U {
			z = append(z, fn(a))
			return z
		}) 
	}
}

func FlatMap[T, U any](s []T) func(func(T) []U) []U {
	return func(fn func(T) []U) []U {
		return FoldLeft[T, []U](s)([]U{})(func(z []U, a T) []U {
			z = append(z, fn(a)...)
			return z
		})
	}
}