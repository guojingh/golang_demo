package main

import "fmt"

const (
	a1 = iota
	b1
	c1
	d1
	e1
)

const (
	_ = iota
	a2
	b2
	c2
	d2
	e2
)

const (
	a3 = iota
	b3
	c3
	_
	d3
	e3
)

const (
	a4 = 1 << iota
	b4
	c4
	d4
	e4
)

func main() {

	fmt.Printf("a1=%d, b1=%d, c1=%d, d1=%d, e1=%d\n", a1, b1, c1, d1, e1)
	fmt.Printf("a2=%d, b2=%d, c2=%d, d2=%d, e2=%d\n", a2, b2, c2, d2, e2)
	fmt.Printf("a3=%d, b3=%d, c3=%d, d3=%d, e3=%d\n", a3, b3, c3, d3, e3)
	fmt.Printf("a4=%d, b4=%d, c4=%d, d4=%d, e4=%d\n", a4, b4, c4, d4, e4)
}
