package goscala

import (
	"reflect"
	"testing"
)

func TestIter(t *testing.T) {
	m := map[string]int{"a": 1}

	iter := reflect.ValueOf(m).MapRange()

	for iter.Next() {
		t.Log(iter.Key(), iter.Value())
	}
}
