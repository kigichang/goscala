package goscala

import (
	"fmt"
)

type Pair[K comparable, V any] interface {
	Tuple2[K, V]
	Tuple() Tuple2[K, V]
	Key() K
	Value() V
}

var _ Pair[int, int] = &pair[int, int]{}

type pair[K comparable, V any] struct {
	tup *tuple2[K, V]
}

func (p *pair[K, V]) Tuple() Tuple2[K, V] {
	return p.tup
}

func (p *pair[K, V]) String() string {
	return fmt.Sprintf(`(%v:%v)`, p.tup.v1, p.tup.v2)
}

func (p *pair[K, V]) Key() K {
	return p.tup.v1
}

func (p *pair[K, V]) V1() K {
	return p.Key()
}

func (p *pair[K, V]) Value() V {
	return p.tup.v2
}

func (p *pair[K, V]) V2() V {
	return p.Value()
}

func (p *pair[K, V]) Get() (K, V) {
	return p.tup.Get()
}
