package grp

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a := Add(3, 8)
	if a != 11 {
		t.Errorf("test")
	}
}
