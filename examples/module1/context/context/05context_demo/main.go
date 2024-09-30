package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	//尽管 ctx 会过期，但在任何情况下调用它的cancel函数都是很好的实践
	//如果不这样做，可能使上下文及其父类存活的时间超过必要的时间
	defer cancel() // 调用 cancel 可以在 deadline 提前取消

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("===", ctx.Err())
	}
}
