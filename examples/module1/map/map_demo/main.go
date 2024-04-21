package main

import (
	"fmt"
)

// 空 map 和 nil map 的区别
func main() {

	//这是一个 nil map
	var m1 map[int]string
	fmt.Println(m1)

	//这是一个空 map
	m2 := make(map[int]string)
	fmt.Println(m2)

	//对 nil map 取值不会 panic 会返回一个空值
	fmt.Println(m1[1])
	//对 nil map 存值会 panic
	m1[1] = "hello"

	//对空 map 存取值都不会 panic
	fmt.Println(m2[1])
	m2[1] = "hello"
	if m1 == nil {
		fmt.Println("m1 是 nil")
	}

	if m2 == nil {
		fmt.Println("m2 是 nil")
	}
}
