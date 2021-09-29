package goscala

func FuncCompose[T, A, R any](f func(T) R, g func(A) T) func(A) R {
	return func(v A) R {
		return f(g(v))
	}
}

func FuncUnitCompose[T, R any](f func(T) R, g func() T) func() R {
	return func() R {
		return f(g())
	}
}

func funcCCompose[T, A, C, R any](f func(T, C) R, g func(A) (T, C)) func(A) R {
	return func(a A) R {
		return f(g(a))
	}
}

func FuncBoolCompose[T, A, R any](f func(T, bool) R, g func(A) (T, bool)) func(A) R {
	return funcCCompose(f, g)
}

func FuncErrCompose[T, A, R any](f func(T, error) R, g func(A) (T, error)) func(A) R {
	return funcCCompose(f, g)
}

func FuncAndThen[T, R, U any](f func(T) R, g func(R) U) func(T) U {
	return func(v T) U {
		return g(f(v))
	}
}

func FuncUnitAndThen[R, U any](f func() R, g func(R) U) func() U {
	return func() U {
		return g(f())
	}
}

func funcCAndThen[T, R, C, U any](f func(T) (R, C), g func(R, C) U) func(T) U {
	return func(v T) U {
		return g(f(v))
	}
}

func FuncBoolAndThen[T, R, U any](f func(T) (R, bool), g func(R, bool) U) func(T) U {
	return funcCAndThen(f, g)
}

func FuncErrAndThen[T, R, U any](f func(T) (R, error), g func(R, error) U) func(T) U {
	return funcCAndThen(f, g)
}
