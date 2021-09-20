package goscala

// Func[R any] represents function: => R
type Func[R any] func() R
type Condition = Func[bool]
type SliceFunc[R any] func() Slice[R]

func (f Func[R]) String() string {
	return TypeStr(f)
}

// FuncBool[R any] represents function: => (R, bool)
type FuncBool[R any] func() (R, bool)

func FuncBoolMaker[R any](v R, ok bool) func() (R, bool) {
	return func() (R, bool) {
		return v, ok
	}
}

type FetchFunc[R any] func() (R, bool)

func (f FuncBool[R]) String() string {
	return TypeStr(f)
}

// FuncErr[R any] represents function: () => (R, error)
type FuncErr[R any] func() (R, error)
type FetchErr[R any] func() (R, error)

func (f FuncErr[R]) String() string {
	return TypeStr(f)
}
