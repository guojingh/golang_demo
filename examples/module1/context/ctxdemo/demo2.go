package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//协程取消

var wait = sync.WaitGroup{}

func main() {

	t1 := time.Now()

	ctx, cancel := context.WithCancel(context.Background())

	wait.Add(1)
	go func() {
		ip, err := GetIp(ctx)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(ip)
	}()

	go func() {
		time.Sleep(2 * time.Second)
		//取消协程
		cancel()
	}()
	wait.Wait()
	fmt.Println("执行完成", time.Since(t1))
}

func GetIp(ctx context.Context) (ip string, err error) {
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("协程取消", ctx.Err())
			err = ctx.Err()
			wait.Done()
			return
		}
	}()
	time.Sleep(4 * time.Second)
	ip = "192.168.222.134"
	wait.Done()
	return
}
