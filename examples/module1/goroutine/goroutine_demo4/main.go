package main

import (
	"fmt"
	"sync"
	"time"
)

//使用无缓存channel模拟跑步接力，这个也是挺有意思的啦

// wg用来等待程序结束
var wg sync.WaitGroup

func main() {
	//创建一个无缓存的通道
	baton := make(chan int)

	//为最后一位跑步者计数加1
	wg.Add(1)

	//第一位跑步者持有接力棒
	go runnerFun(baton)

	//开始比赛
	baton <- 1
	//等待比赛结束
	wg.Wait()
}

func runnerFun(baton chan int) {
	var newRunner int
	//等待接力棒
	runner := <-baton
	//开始绕着跑到跑步
	fmt.Printf("Runner %d Running With Baton\n", runner)
	//创建下一位跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go runnerFun(baton)
	}
	//围绕跑道跑
	time.Sleep(100 * time.Millisecond)
	//比赛结束了吗
	if runner == 4 {
		fmt.Printf("Runner %d finished, Race Over\n", runner)
		wg.Done()
		return
	}
	//将接力棒交给下一位跑者
	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)
	baton <- newRunner
}
