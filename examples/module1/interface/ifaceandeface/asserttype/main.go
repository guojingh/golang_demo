package main

import "fmt"

//断言和类型转换实例
func main() {

	converType()
	assertType()
}

//关于类型转换，转换前要相互兼容才行
func converType() {
	var i int = 9

	var f float64
	f = float64(i)
	fmt.Printf("%T, %v\n", f, f)

	f = 10.08
	a := int(f)
	fmt.Printf("%T, %v", a, a)

	//s := []int(i)  //cannot convert i (variable of type int) to type []int
}

type Student struct {
	Name string
	Age  int
}

func assertType() {
	var i interface{} = new(Student)
	//s := i.(Student) //interface conversion: interface {} is *main.Student, not main.Student
	//可以采用安全断言解决
	if s, ok := i.(Student); ok {
		fmt.Println(s)
	}
}
