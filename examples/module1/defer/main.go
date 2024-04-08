package main

import (
	"fmt"
	"time"
)

/*func test03() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	loopFunc()
	time.Sleep(time.Second)
}

func loopFunc() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {

		//闭包，匿名函数
		go func(i int) {
			lock.Lock()
			defer lock.Unlock()
			fmt.Println("loopFunc:", i)
		}(i)
	}
}
*/

func main() {
	loopFunc()
	time.Sleep(time.Second)
}

func loopFunc() {

	for i := 0; i < 3; i++ {
		go fmt.Println("loopFunc", i)
	}
}
