package goscala

func VF[T any](v T) func() T {
	return func() T {
		return v
	}
}

func VPF[T any](v T) func() (T, bool) {
	return func() (T, bool) {
		return v, true
	}
}

func ValueFunc[T any](v T) func() T {
	return VF(v)
}

func ValuePartialFunc[T any](v T) func() (T, bool) {
	return VPF(v)
}

func VF2[T, U any](v T, u U) func() (T, U) {
	return func() (T, U) {
		return v, u
	}
}

func ValueFunc2[T, U any](v T, u U) func() (T, U) {
	return VF2(v, u)
}

func Id[T any](v T) T {
	return v
}

func Identity[T any](v T) T {
	return Id(v)
}

func Id2[T, U any](v T, u U) (T, U) {
	return v, u
}

func Identity2[T, U any](v T, u U) (T, U) {
	return v, u
}
