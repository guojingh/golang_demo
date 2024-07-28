package factory

import "fmt"

// Person3 工厂方法模式
type Person3 struct {
	name string
	age  int
}

func NewPersonFactory(age int) func(name string) Person3 {
	return func(name string) Person3 {
		return Person3{
			name: name,
			age:  age,
		}
	}
}

func main() {
	//创建具有默认年龄的工厂
	newBaby := NewPersonFactory(1)
	baby := newBaby("john")
	fmt.Printf("john is %v", baby)

	newTeenager := NewPersonFactory(16)
	teen := newTeenager("jill")
	fmt.Printf("teen is %v", teen)
}
