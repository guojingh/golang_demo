package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//截止时间

var wg = sync.WaitGroup{}

func main3() {

	t1 := time.Now()

	// 添加截至时间 ，第二个参数是时间即 2 秒钟之后取消
	// cancel 是一个 func() 用来取消
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))

	wg.Add(1)
	go func() {
		ip, err := GetIp1(ctx, &wg)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(ip)
	}()

	//马上取消 GetIp1() 的协程
	cancel()

	/*	go func() {
		time.Sleep(2 * time.Second)
		//取消协程
		cancel()
	}()*/
	wg.Wait()
	fmt.Println("执行完成", time.Since(t1))
}

func GetIp1(ctx context.Context, wg *sync.WaitGroup) (ip string, err error) {
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("协程取消", ctx.Err())
			err = ctx.Err()
			wg.Done()
			return
		}
	}()
	time.Sleep(4 * time.Second)
	ip = "192.168.222.134"
	wg.Done()
	return
}