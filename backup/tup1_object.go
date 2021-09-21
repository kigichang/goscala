package goscala

func MakeTuple1[T any](v T) Tuple1[T] {
	return &tuple1[T]{
		v: v,
	}
}