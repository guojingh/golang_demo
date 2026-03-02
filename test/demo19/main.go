package main

import (
	"fmt"
	"time"
)

// nil channel 的妙用
// 程序依次输出5和7这两个数字后退出
func main() {
	c1, c2 := make(chan int), make(chan int)
	go func() {
		time.Sleep(time.Second * 5)
		c1 <- 5
		// 关闭c1，无缓存channel关闭后会一直从中获取0值
		close(c1)
	}()

	go func() {
		time.Sleep(time.Second * 7)
		c2 <- 7
		close(c2)
	}()

	for {
		select {
		case x, ok := <-c1:
			if ok {
				fmt.Println(x)
			} else {
				c1 = nil
			}
		case x, ok := <-c2:
			if ok {
				fmt.Println(x)
			} else {
				c2 = nil
			}
		}
		if c1 == nil && c2 == nil {
			break
		}
	}

	fmt.Println("program end")
}
