package monad

func Fold[T, C, U any](fetch func() (T, C)) func(func(C) U, func(T) U) U {
	return Currying3To2(MFold[T, C, U])(fetch)
}

func FoldBool[T, U any](fetch func() (T, bool)) func(func() U, func(T) U) U {
	return Currying3To2(MFoldBool[T, U])(fetch)
}

func FoldErr[T, U any](fetch func() (T, error)) func(func(error) U, func(T) U) U {
	return Currying3To2(MFoldErr[T, U])(fetch)
}

func FoldLeft[T, U any](s []T) func(U) func(func(U, T) U) U {
	return Currying3(MFoldLeft[T, U])(s)
}

func ScanLeft[T, U any](s []T) func(U) func(func(U, T) U) []U {
	return Currying3(MScanLeft[T, U])(s)
}

func FoldRight[T, U any](s []T) func(U) func(func(T, U) U) U {
	return Currying3(MFoldRight[T, U])(s)
}

func ScanRight[T, U any](s []T) func(U) func(func(T, U) U) []U {
	return Currying3(MScanRight[T, U])(s)
}

func Map[T, U any](s []T) func(func(T) U) []U {
	return Currying2(MMap[T, U])(s)
}

func FlatMap[T, U any](s []T) func(func(T) []U) []U {
	return Currying2(MFlatMap[T, U])(s)
}

func FuncCompose[T, A, R any](f func(T) R) func(func(A) T) func(A) R {
	return Currying2(MFuncCompose[T, A, R])(f)
}

func FuncAndThen[T, R, U any](f func(T) R) func(func(R) U) func(T) U {
	return Currying2(MFuncAndThen[T, R, U])(f)
}

func FuncUnitAndThen[R, U any](f func() R) func(func(R) U) func() U {
	return Currying2(MFuncUnitAndThen[R, U])(f)
}

func FuncBoolAndThen[T, R, U any](f func(T) (R, bool)) func(func (R, bool) U) func (T) U {
	return Currying2(MFuncBoolAndThen[T, R, U])(f)
}