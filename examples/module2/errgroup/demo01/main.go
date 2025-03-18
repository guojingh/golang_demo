package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

// errgroup 使用方法 超时返回
// errgroup包提供了一种方便的方式来跟踪和处理多个goroutine中的错误。
// 它可以让你启动多个goroutine，并等待它们全部完成，或者在任何一个goroutine返回错误时立即取消所有其他goroutine。
func main() {

	// 创建一个带有取消信号的context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 创建带有 context 的 err group
	g, ctx := errgroup.WithContext(ctx)

	// 添加并发任务
	g.Go(func() error {
		select {
		case <-ctx.Done():
			fmt.Println("任务1被取消")
			return ctx.Err()
		case <-time.After(2 * time.Second):
			fmt.Println("任务1完成")
			return nil
		}
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
			fmt.Println("任务2被取消")
			return ctx.Err()
		case <-time.After(3 * time.Second):
			return nil
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Println("任务执行出错")
	} else {
		fmt.Println("所以任务执行完成")
	}
}
