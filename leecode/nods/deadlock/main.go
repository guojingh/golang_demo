package main

import "fmt"

// 写一个死锁并且分析
func main() {

	c1 := make(chan int)
	fmt.Println(<-c1)
}
