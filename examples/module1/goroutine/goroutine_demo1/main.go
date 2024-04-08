package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 使用两个协程打印素数
// wg用来等待协程完成
var wg sync.WaitGroup

// 测试goroutine的调度
func main() {
	//分配一个处理器给程序使用
	runtime.GOMAXPROCS(1)
	//计算加2表示要等待2个goroutine
	wg.Add(2)

	//创建2个goroutine
	fmt.Println("Create goroutine")
	go printPrime("A")
	go printPrime("B")

	//等待goroutine完成
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

func printPrime(prefix string) {
	defer wg.Done()
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
