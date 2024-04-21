package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/Jeffail/tunny"
)

/*
目的：实现 goroutine 的并发数量
1.通过 有缓存的 channel 实现
2.通过第三方库实现 "github.com/Jeffail/tunny"
*/

func main() {

	//ctrByChannel()
	ctrByTunny()
}

// 通过 有缓存的 channel 实现
func ctrByChannel() {
	//定义并发的最大数量
	count := 10
	//定义总共要执行的goroutine
	sum := 100
	//用来控制主程等待
	wg := sync.WaitGroup{}
	//count数量缓存的 channel
	c := make(chan struct{}, count)

	//开启 100 次循环
	for i := 0; i < sum; i++ {
		//wg+1
		wg.Add(1)
		//channel 中塞值
		c <- struct{}{}
		//开启协程
		go func(j int) {
			//wg-1
			defer wg.Done()
			//输出
			fmt.Println(j)
			//channel 中取值
			<-c
		}(i)
	}

	wg.Wait()
}

// 通过第三方库实现
func ctrByTunny() {
	//创建新的线程池，这个 func 是后面每个 pool.Process 都会调用的
	pool := tunny.NewFunc(10, func(i interface{}) interface{} {
		fmt.Println(i)
		time.Sleep(time.Second)
		return nil
	})

	//关闭资源
	defer pool.Close()
	for i := 0; i < 500; i++ {
		//从协程池里面获取资源
		go pool.Process(i)
	}
	time.Sleep(time.Second * 100)
}
