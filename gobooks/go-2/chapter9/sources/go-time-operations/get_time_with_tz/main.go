package main

import (
	"fmt"
	"time"
)

// 将本地时间转化为特定时区的即时时间
func main() {
	t := time.Now()
	fmt.Println(t)

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("load time location failed:", err)
		return
	}
	t = t.In(loc)
	fmt.Println(t)
}
