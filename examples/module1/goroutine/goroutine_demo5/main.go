package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 通过有缓存channel模拟处理任务
const (
	numberGoroutines = 4  // 要使用的goroutine数量
	taskLoad         = 10 // 要处理的工作数量
)

var wg sync.WaitGroup

func init() {
	//初始化随机数种子
	rand.Seed(time.Now().Unix())
}

func main() {
	//创建一个有缓存通道来管理工作
	tasks := make(chan string, taskLoad)
	//启动 goroutine来处理工作
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	//增加一组要完成的工作
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}
	//所有工作完成后关闭通道
	close(tasks)
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()
	for {
		//等待分配工作
		task, ok := <-tasks
		if !ok {
			//意味着通道已空，并且已经关闭
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}
		//显示我么要开始工作了
		fmt.Printf("Worker: %d : Start %s\n", worker, task)
		//随机等待一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		//显示我们完成了工作
		fmt.Printf("Worker: %d : Completed %s \n", worker, task)
	}
}
