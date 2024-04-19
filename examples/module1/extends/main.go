package main

import (
	"fmt"
)

// 在 golang 中实现继承---通过组合实现继承
type Animal struct {
	Name string
}

func (a *Animal) Eat() {
	fmt.Printf("%v is eating", a.Name)
	fmt.Println()
}

type Cat struct {
	Animal *Animal
}

func main() {

	cat := &Cat{
		Animal: &Animal{
			Name: "小明",
		},
	}

	cat.Animal.Eat() //小明 is eating

	l := new([10]int)
	fmt.Println(l[:])
}
