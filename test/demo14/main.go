package main

import (
	"time"
)

// channel 一对一通知信号
type signal struct{}

// worker 函数
func worker() {
	println("worker is working...")
	time.Sleep(1 * time.Second)
}

// <-chan signal 返回值是一个只读通道(箭头在左边)，箭头在右边为只写通道
func spawn(f func()) <-chan signal {
	c := make(chan signal)
	go func() {
		println("worker start to working...")
		f()
		c <- signal{}
	}()
	return c
}

func main() {
	println("start a worker...")
	c := spawn(worker)
	<-c
	println("worker work done!...")
}
