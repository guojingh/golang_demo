package main

import "fmt"

// 比较
// 切片变量之间不能直接通过操作符进行等值比较，但可以与nil进行等值比较
func main() {
	var a = []byte{'a', 'b', 'c'}
	var b = []byte{'a', 'b', 'c'}

	// 切片变量之间不能直接通过操作符进行等值比较：invalid operation: a == b (slice can only be compared to nil)
	// if a == b {
	// 	fmt.Println("a=b")
	// } else {
	// 	fmt.Println("a!=b")
	// }

	if a != nil && b != nil {
		fmt.Println("a and b !=nil")
	}
}
