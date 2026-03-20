package main

import (
	"fmt"
	"time"
)

// 获取当前时间
func main() {
	/*
		这三个字段组成的结构体要同事表示两种时间-挂钟时间，单调时间
		type Time struct {
			wall uint64
			ext  int64
			loc *Location
		}

	*/
	t := time.Now() //返回一个time类型，用作对即时时间对抽象

	fmt.Println(t) //输出当前时间
}
