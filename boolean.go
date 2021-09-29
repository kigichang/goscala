package goscala

func True() bool {
	return true
}

func False() bool {
	return false
}

func Cond[T any](ok bool, t, f T) T {
	if ok {
		return t
	}
	return f
}

func Ternary[T any](cond func() bool, succ func() T, fail func() T) T {
	if cond() {
		return succ()
	}
	return fail()
}
