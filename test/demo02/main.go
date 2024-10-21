package main

import (
	"fmt"
	"sync"
)

func main() {
	// 三个 goroutine，两个 goroutine 计算（不打印），一个加到1万，一个加到2万
	// 第三个等他们算完，把和打印出来

	wg := new(sync.WaitGroup)
	sChan1 := make(chan int)
	sChan2 := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		sum1 := 0
		for i := 0; i < 10000; i++ {
			sum1 += i
		}
		sChan1 <- sum1
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		sum2 := 0
		for i := 0; i < 20000; i++ {
			sum2 += i
		}
		sChan2 <- sum2
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		sum1 := <-sChan1
		sum2 := <-sChan2
		sum := sum1 + sum2
		fmt.Printf("sum1 + sum2 = %d\n", sum)

	}()

	wg.Wait()
}
