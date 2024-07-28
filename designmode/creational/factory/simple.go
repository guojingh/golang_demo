package factory

import "fmt"

// Person1 简单工厂模式
type Person1 struct {
	Name string
	Age  int
}

func (p Person1) Greet1() {
	fmt.Printf("Hi! My name is %s", p.Name)
}

func NewPerson1(name string, age int) *Person1 {
	return &Person1{
		Name: name,
		Age:  age,
	}
}
