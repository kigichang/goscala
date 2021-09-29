package goscala

func Currying2[A, B, C any](f func(A, B) C) func(A) func(B) C {
	return func(a A) func(B) C {
		return func(b B) C {
			return f(a, b)
		}
	}
}

func Currying3[A, B, C, D any](f func(A, B, C) D) func(A) func(B) func(C) D {
	return func(a A) func(B) func(C) D {
		return func(b B) func(C) D {
			return func(c C) D {
				return f(a, b, c)
			}
		}
	}
}

func Currying3To2[A, B, C, D any](f func(A, B, C) D) func(A) func(B, C) D {
	return func(a A) func(B, C) D {
		return func(b B, c C) D {
			return f(a, b, c)
		}
	}
}
