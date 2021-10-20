// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package seq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type traitIter[T any] struct {
	next func() bool
	get  func() T
}

func (t *traitIter[T]) Next() bool {
	return t.next()
}

func (t *traitIter[T]) Get() T {
	return t.get()
}

func TestSkip(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}

	idx := -1

	it := &traitIter[int]{
		next: func() bool {
			idx += 1
			return 0 <= idx && idx < len(s)
		},
		get: func() int {
			return s[idx]
		},
	}

	skip[int](it, 3).Next()
	assert.Equal(t, 3, it.Get())

	idx = len(s)
	it = &traitIter[int]{
		next: func() bool {
			idx -= 1
			return 0 <= idx && idx < len(s)
		},
		get: func() int {
			return s[idx]
		},
	}
	skip[int](it, 3).Next()
	assert.Equal(t, 2, it.Get())
}
