package main

import (
	"fmt"
	"sync"
)

// go routine 向外通信
// goroutine 的通信是通过 channel 实现的
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	c := make(chan int, 1)

	go func() {
		defer wg.Done()
		c <- 1
		return
	}()

	wg.Wait()
	fmt.Println(<-c)
}
