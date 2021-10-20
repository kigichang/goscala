// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package seq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type slices[T any] []T

func (s slices[T]) Append(v T) Interface[T] {
	return append(s, v)
}

func (s slices[T]) Len() int {
	return len(s)
}

func (s slices[T]) Iterate() Iterator[T] {
	idx := -1

	it := &traitIter[T]{
		next: func() bool {
			idx += 1
			return 0 <= idx && idx < s.Len()
		},
		get: func() T {
			return s[idx]
		},
	}
	return it
}

func (s slices[T]) BackIterate() Iterator[T] {
	idx := s.Len()
	it := &traitIter[T]{
		next: func() bool {
			idx -= 1
			return 0 <= idx && idx < s.Len()
		},
		get: func() T {
			return s[idx]
		},
	}
	return it
}

func slicesEmpty[T any]() slices[T] {
	return slices[T]([]T{})

}

func TestClone(t *testing.T) {
	s := slices[int]([]int{0, 1, 2, 3, 4, 5})

	assert.Equal(t, s, Clone[slices[int], int](s, slicesEmpty[int]))

	it := skip(s.Iterate(), 3)
	assert.Equal(t, s[3:], CopyFrom(it, slicesEmpty[int]))
}
