package main

import (
	"fmt"
	"sync"
)

// 两个协程交替打印输出
func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	c1 := make(chan int, 1)
	c2 := make(chan int)

	c1 <- 1

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-c1
			fmt.Printf("func_01: %d\n", i)
			c2 <- 1
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-c2
			fmt.Printf("func_02: %d\n", i)
			c1 <- 1
		}
	}()

	wg.Wait()
}
