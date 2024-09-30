package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 使用channel的方式实现

func worker(ctx context.Context) {
	defer wg.Done()
	go worker2(ctx)
	for {
		fmt.Println("worker1")
		time.Sleep(time.Second)
		// 如何接收外部命令实现退出
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

func worker2(ctx context.Context) {
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()  // 调用 cancel() 告诉 goroutine 退出
	wg.Wait() //等待...
	time.Sleep(time.Second * 5)
	fmt.Println("over")
}
