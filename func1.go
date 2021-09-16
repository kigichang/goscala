package goscala

// Func1[T1, R any] represents function: T1 => R
type Func1[T1, R any] func(T1) R

type EqualMethod[T1 any] func(T1) bool

type Predict[T1 any] Func1[T1, bool]

func (f Func1[T1, R]) String() string {
	return typstr(f)
}

// Func1Bool[T1, R any] represents function: T1 => (R, bool)
type Func1Bool[T1, R any] func(T1) (R, bool)
type PartialFunc[T, U any] Func1Bool[T, U]

//func MakePartialFunc[T, U any](p Predict[T], act Func1[T, U]) PartialFunc[T, U] {
//	return PartialFunc[T, U](func(v T) (ret U, ok bool) {
//		ok = p(v)
//		if ok {
//			ret = act(v)
//		}
//		return
//	})
//}

func (f Func1Bool[T1, R]) String() string {
	return typstr(f)
}

// Func1Err[T1, R any] represents function: T1 => (R, error)
type Func1Err[T1, R any] func(T1) (R, error)

func (f Func1Err[T1, R]) String() string {
	return typstr(f)
}
