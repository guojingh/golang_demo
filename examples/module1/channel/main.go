package main

import (
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {

	semaphore
	c1 := make(chan int, 1)

	c1 <- 1
	go func(chan int) {
		fmt.Println("写协程开始...", time.Now())
		//time.Sleep(time.Second * 2)
		c1 <- 10
		fmt.Println("写协程结束...", time.Now())
	}(c1)

	go func(chan int) {
		fmt.Println("读协程开始...", time.Now())
		time.Sleep(time.Second * 3)
		fmt.Println(<-c1)
		fmt.Println("读协程结束...", time.Now())
	}(c1)

	time.Sleep(time.Second * 5)

}
