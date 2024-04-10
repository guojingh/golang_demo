package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg2 = sync.WaitGroup{}

func main() {

	t1 := time.Now()

	// 添加截至时间 ，第二个参数是时间即 2 秒钟之后取消
	// cancel 是一个 func() 用来取消
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	wg2.Add(1)
	go func() {
		ip, err := GetIp2(ctx, &wg2)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(ip)
	}()

	//马上取消 GetIp1() 的协程
	//cancel()

	/*	go func() {
		time.Sleep(2 * time.Second)
		//取消协程
		cancel()
	}()*/
	wg2.Wait()
	fmt.Println("执行完成", time.Since(t1))
}

func GetIp2(ctx context.Context, wg *sync.WaitGroup) (ip string, err error) {
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