package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 如果多次调用Notify拦截某信号，但每次调用使用的channel不同，那么当应用进程收到异步信号时，
// Go运行时会给每个channel发送一份异步信号副本：
func main() {
	c1 := make(chan os.Signal, 1)
	c2 := make(chan os.Signal, 2)

	signal.Notify(c1, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(c2, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		s := <-c1
		fmt.Println("c1：收到异步信号", s)
	}()

	s := <-c2
	fmt.Println("c2：收到异步信号", s)
	time.Sleep(5 * time.Second)
}
