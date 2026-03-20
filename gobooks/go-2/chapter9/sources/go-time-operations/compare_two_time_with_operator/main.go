package main

import (
	"fmt"
	"time"
)

// 比较两个不同时区的时间
func main() {
	t := time.Now()
	fmt.Println(t)

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("load location time error:", err)
		return
	}
	t1 := t.In(loc)
	// 时间单位最好不要用==；而应该使用 time.Time 提供的 Equal 方法代替
	fmt.Println(t == t1)
}
