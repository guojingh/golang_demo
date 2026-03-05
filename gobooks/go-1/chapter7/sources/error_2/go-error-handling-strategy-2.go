package main

import (
	"errors"
	"fmt"
)

// errors.As() 函数的应用：用来把错误链中的某个具体类型提取出来，适合做类型的匹配，并拿到其字段/方法。
// errors.As(err, &target)
type MyError struct {
	e string
}

func (e *MyError) Error() string {
	return e.e
}

func main() {
	var err = &MyError{"my error type"}
	err1 := fmt.Errorf("wrap err1: %w", err)
	err2 := fmt.Errorf("wrap err2: %w", err1)
	var e *MyError
	if errors.As(err2, &e) {
		println("MyError is on the chain of err2")
		println(e == err)
		return
	}

	fmt.Println("MyError is not on the chain of err2")

}
