package safemap

import (
	"fmt"
	"testing"
)

func TestEach(t *testing.T) {
	m := NewRWMap(10)
	m.Set(0, 0)
	m.Set(1, 1)
	m.Set(2, 2)
	m.Set(3, 3)
	m.Set(4, 4)

	r, b := m.Get(2)
	if b != false {
		fmt.Printf("m[%v] = %v\n", 2, r)
	}

	m.Each(eachTest)
}

func eachTest(k, v int) bool {
	return false
}
