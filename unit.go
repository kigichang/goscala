package goscala

type UnitRef = struct{}

var unit = UnitRef{}

func Unit() UnitRef {
	return unit
}

func UnitWrap[T any](fn func(T)) func(T) UnitRef {
	return func(v T) UnitRef {
		fn(v)
		return unit
	}
}
