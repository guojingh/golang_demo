package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 使用 无缓存 channel 模拟网球比赛
// 这段代码好有意思 哈哈哈
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	//创建一个无缓存通道
	court := make(chan int)

	//计数加2,，表示要等待2个goroutine
	wg.Add(2)

	//启动2个选手
	go player("小明", court)
	go player("小李", court)

	//发球
	court <- 1
	//等待游戏结束
	wg.Wait()
}

func player(name string, court chan int) {
	//在函数退出时通过 Done 来通知main函数工作已完成
	defer wg.Done()

	for {
		//等待把球打过来
		ball, ok := <-court
		if !ok {
			//如果通道被关闭，我们就赢了
			fmt.Printf("player %s Won\n", name)
			return
		}

		//选随机数，然后用这个数判断我们是否丢球
		n := rand.Intn(100)
		if n%3 == 0 {
			fmt.Printf("player %s Missed\n", name)
			//关闭通道，我们输了
			close(court)
			return
		}

		//显示击球数，并将击球数加1
		fmt.Printf("player %s Hit %d\n", name, ball)
		ball++

		//将球打向对手
		court <- ball
	}
}
