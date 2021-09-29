package goscala

func PFV[T, U any](succ func(T) U, fail func() U) func(T, bool) U {
	return func(v T, ok bool) U {
		if ok {
			return succ(v)
		}
		return fail()
	}
}

func PFF[T, U any](succ func(T) U, fail func() U) func(func() (T, bool)) U {
	return func(pf func() (T, bool)) U {
		return PFV(succ, fail)(pf())
	}
}

func PFErrV[T, U any](succ func(T) U, fail func(error) U) func(T, error) U {
	return func(v T, err error) U {
		if err == nil {
			return succ(v)
		}
		return fail(err)
	}
}

func PFErrF[T, U any](succ func(T) U, fail func(error) U) func(func() (T, error)) U {
	return func(pf func() (T, error)) U {
		return PFErrV(succ, fail)(pf())
	}
}
