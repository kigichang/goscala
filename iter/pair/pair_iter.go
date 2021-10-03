// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package pair

type Iter[K comparable, V any] interface {
	Next() bool
	Get() (K, V)
}

type abstractIter[K comparable, V any] struct {
	next func() bool
	get  func() (K, V)
}

func ToMap[K comparable, V any](a Iter[K, V]) map[K]V {
	ret := make(map[K]V)
	for a.Next() {
		k, v := a.Get()
		ret[k] = v
	}
	return ret
}

func Count[K comparable, V any](a Iter[K, V], p func(K, V) bool) (ret int) {
	for a.Next() {
		if p(a.Get()) {
			ret++
		}
	}
	return
}

func Find[K comparable, V any](a Iter[K, V], p func(K, V) bool) (retK K, retV V, ok bool) {
	for a.Next() {
		if ok = p(a.Get()); ok {
			retK, retV = a.Get()
			return
		}
	}
	return
}

func Exists[K comparable, V any](a Iter[K, V], p func(K, V) bool) (ok bool) {
	_, _, ok = Find(a, p)
	return
}
