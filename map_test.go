package goscala

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMapLen(t *testing.T) {
	m := MapFrom(P(1, "1"), P(2, "2"), P(3, "3"))
	assert.Equal(t, 3, m.Len())
}
