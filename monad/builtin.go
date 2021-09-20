package monad

func Equal[T comparable](a, b T) bool {
	return a == b
}

func Compare[T Ordered](a, b T) int {
	if a == b {
		return 0
	}

	if a > b {
		return 1
	}

	return -1
}

func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}