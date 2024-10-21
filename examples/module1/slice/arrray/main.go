package main

import (
	"fmt"
	"time"
)

func main() {

	slice := []int{10, 20, 30, 40, 50}

	newSlice := slice[1:3:4]

	//运行发生 panic 切片不能操作超出其长度的位置
	//newSlice[3] = 35

	time.Now()

	newSlice = append(newSlice, 60)
	fmt.Println(slice, newSlice)

}
