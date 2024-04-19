package main

import (
	"fmt"
)

type alice string

type people struct {
	name string
	age  int
}

type man people

func main() {

	s := "hello"

	var a alice
	a = "world"

	fmt.Printf("s=%T;a=%T\n", s, a)

	b := string(a)
	fmt.Printf("b=%s;type=%T\n", b, b)

	p := man{
		name: "小明",
		age:  10,
	}

	m := people(p)
	fmt.Printf("p=%v;m=%v\n", p, m)
	fmt.Printf("p.type=%T;m.Type=%T\n", p, m)

}