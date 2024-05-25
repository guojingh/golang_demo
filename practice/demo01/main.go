package main

import (
	"sync"
)

type a struct {
	b *b
}

type b struct {
	value  map[string]interface{}
	locker sync.RWMutex
}

func main() {

	a := &a{
		b: &b{
			value: map[string]interface{}{"a": "b"},
		},
	}
	a.b.locker.RLocker()
	a.b.locker.RUnlock()

}
