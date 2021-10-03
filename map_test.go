// Copyright © 2021 Kigi Chang <kigi.chang@gmail.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package goscala_test

import (
	"testing"

	gs "github.com/kigichang/goscala"
	"github.com/kigichang/goscala/maps"
	"github.com/kigichang/goscala/slices"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	m := gs.MkMap[int, string]()

	m.Add(gs.P(1, "1"))
	m.Add(nil)
	v, ok := m.Get(1)

	assert.Equal(t, "1", v)
	assert.True(t, ok)

	assert.Equal(t, m.Keys(), slices.From(1))
	assert.Equal(t, m.Values(), slices.From("1"))

	v, ok = m.Get(2)

	assert.Equal(t, "", v)
	assert.False(t, ok)

	assert.Equal(t, "abc", m.GetOrElse(2, "abc"))
	assert.Equal(t, "1", m.GetOrElse(1, "abc"))
}

func TestMapContains(t *testing.T) {
	m := maps.From(gs.P(1, "1"), gs.P(2, "2"))

	assert.True(t, m.Contains(1))
	assert.False(t, m.Contains(3))
}

func TestMapCount(t *testing.T) {
	m := maps.From(gs.P(1, "1"), gs.P(2, "2"), gs.P(-1, "-1"))

	count := m.Count(func(k int, _ string) bool {
		return k > 0
	})

	assert.Equal(t, 2, count)
}

func TestMapFindAndExists(t *testing.T) {
	m := maps.From(gs.P(1, "1"), gs.P(2, "2"))

	assert.True(t, m.Find(func(k int, _ string) bool {
		return k > 0
	}).IsDefined())

	assert.True(t, m.Find(func(k int, _ string) bool {
		return k > 10
	}).IsEmpty())

	assert.True(t, m.Exists(func(k int, _ string) bool {
		return k > 0
	}))

	assert.False(t, m.Exists(func(k int, _ string) bool {
		return k > 10
	}))
}

func TestMapFilter(t *testing.T) {
	m := maps.From(gs.P(1, "1"), gs.P(2, "2"), gs.P(-1, "-1"))

	m2 := m.Filter(func(k int, _ string) bool {
		return k < 0
	})

	assert.Equal(t, m2.Keys(), slices.From(-1))
	assert.Equal(t, m2.Values(), slices.From("-1"))
}

func TestMapFilterNot(t *testing.T) {
	m := maps.From(gs.P(1, "1"), gs.P(2, "2"), gs.P(-1, "-1"))

	m2 := m.FilterNot(func(k int, _ string) bool {
		return k > 0
	})

	assert.Equal(t, m2.Keys(), slices.From(-1))
	assert.Equal(t, m2.Values(), slices.From("-1"))
}

func TestMapForall(t *testing.T) {
	f1 := func(k int, _ string) bool {
		return k > 0
	}

	m := maps.From(gs.P(1, "1"), gs.P(2, "2"))
	assert.True(t, m.Forall(f1))

	m.Add(gs.P(-1, "-1"))
	assert.False(t, m.Forall(f1))

	assert.True(t, gs.MkMap[int, string]().Forall(f1))
}

func TestMapForeach(t *testing.T) {
	sum := 0
	m := maps.From(gs.P(1, "1"), gs.P(2, "2"))
	m.Foreach(func(k int, _ string) {
		sum += k
	})

	assert.Equal(t, 3, sum)
}

func TestMapPartition(t *testing.T) {
	m := maps.From(gs.P(1, "1"), gs.P(-1, "-1"))

	a, b := m.Partition(func(k int, _ string) bool {
		return k >= 0
	})

	assert.Equal(t, a.Keys(), slices.From(-1))
	assert.Equal(t, a.Values(), slices.From("-1"))

	assert.Equal(t, b.Keys(), slices.From(1))
	assert.Equal(t, b.Values(), slices.From("1"))
}
