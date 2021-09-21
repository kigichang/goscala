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

func FoldBool[T, U any](fetch func() (T, bool)) func(func(bool) U, func(T) U) U {
	return Fold[T, bool, U](fetch)
}

func FoldErr[T, U any](fetch func() (T, error)) func(func(error) U, func(T) U) U {
	return Fold[T, error, U](fetch)
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

func ScanLeft[T, U any](s []T) func(U) func(func(U, T) U) []U {
	return func(z U) func(func(U, T) U) []U {
		return func(fn func(U, T) U) []U {
			return FoldLeft[T, []U](s)([]U{z})(func (a []U, b T) []U {
				return append(a, fn(a[len(a)-1], b))
			})
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

func ScanRight[T, U any](s []T) func(U) func(func(T, U) U) []U {
	return func(z U) func(func(T, U) U) []U {
		return func (fn func(T, U) U) []U {
			result := FoldRight[T, []U](s)([]U{z})(func(a T, b []U) []U {
				return append(b, fn(a, b[len(b) - 1]))
			})

			size := len(result)
			half := size / 2

			for i := 0; i < half; i++ {
				tmp := result[i]
				result[i] = result[size - 1 - i] 
				result[size - 1 - i] = tmp
			}

			return result
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

func FuncCompose[T, A, R any](f func(T) R) func(func(A) T) func(A) R {
	return func(g func(A) T) func(A) R {
		return func(v A) R {
			return f(g(v))
		}
	}
}

func FuncAndThen[T, R, U any](f func(T) R) func(func(R) U) func(T) U {
	return func(g func(R) U) func(T) U {
		return func(v T) U {
			return g(f(v))
		}
	}
}

func FuncBoolAndThen[T, R, U any](f func(T) (R, bool)) func(func (R, bool) U) func (T) U {
	return func(g func(R, bool) U) func(T) U {
		return func(v T) U {
			return g(f(v))
		}
	}
}