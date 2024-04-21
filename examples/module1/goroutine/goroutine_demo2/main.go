package main

import (
	"fmt"
	"runtime"
	"sync"
)

//可以使用 go build -race 用竞争检测器标志来编译程序
//显示的结果就是存在并发危险的地方

var (
	counter int
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter()
	go incCounter()

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter() {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter

		runtime.Gosched()

		value++

		counter = value
	}
}
