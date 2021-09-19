package intern

type a[T any] struct {
	v T
}

func (x *a[T]) Get() T {
	return x.v
}

func AA[T any](v T) *a[T] {
	return &a[T] {
		v: v,
	}
}