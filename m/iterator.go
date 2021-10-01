package m

type Iterable[T any] interface {
	HasNext() bool
	Next() T
	Add(T)
	AddIter(Iterable[T])
}

type Iterator[T any] struct {
	hasNext func() bool
	next    func() T
	add     func(T)
	addIter func(Iterable[T])
}

func (i *Iterator[T]) HasNext() bool {
	return i.hasNext()
}

func (i *Iterator[T]) Next() T {
	return i.next()
}

func (i *Iterator[T]) Add(v T) {
	i.add(v)
}

func (i *Iterator[T]) AddIter(that Iterable[T]) {
	i.addIter(that)
}

func FoldLeft[T, U any](s Iterable[T], z U, fn func(a U, b T) U) U {
	zz := z
	for s.HasNext() {
		zz = fn(zz, s.Next())
	}
	return zz
}

type A[T any] struct{}

func (a *A[T]) Hello() {
	println("abc")
}

type B[T any] struct {
	*A[T]
}

func ABC[T any](b *B[T]) {
	b.Hello()
}
