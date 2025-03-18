package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

// 并发执行一组耗时的计算任务，并限制最大同时执行的goroutine数

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)
	maxWorkers := 5
	for i := 0; i < maxWorkers; i++ {
		g.Go(func() error {
			result := compute()
			fmt.Println(result)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("one of the computations returned an error:", err)
	}
}

func compute() int {
	// 模拟计算时间消耗，以及返回数值
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	return rand.Intn(100)
}
