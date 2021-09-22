package monad

func EmptySlice[T any]() []T {
	return []T{}
}

func ElemSlice[T any](v T) []T {
	return []T{v}
}