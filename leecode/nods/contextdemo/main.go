package main

import "fmt"

func divide(a, b int) int {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err = ", err)
		}
	}()

	if b == 0 {
		panic("divided by zero")
	} else {
		return a / b
	}

}

func main() {
	_ = divide(2, 0)
}
