package factory

import "fmt"

type Person2 interface {
	Greet2()
}

type person2 struct {
	name string
	age  int
}

func (p person2) Greet2() {
	fmt.Printf("Hi! My name is %s", p.name)
}

// NewPerson2 返回一个接口，而非结构体
func NewPerson2(name string, age int) Person2 {
	return person2{
		name: name,
		age:  age,
	}
}
