package main

import (
	ring "container/ring"
	"fmt"
)

// 统计最近10次平均值
func main() {

	const SIZE = 10

	//创建 ring用到container/ring
	ringArray := ring.New(SIZE)
	for i := 0; i < 10; i++ {
		ringArray.Value = i
		ringArray = ringArray.Next()
	}

	sum := 0
	ringArray.Do(func(i interface{}) {
		sum += i.(int)
	})

	fmt.Printf("avg is %v\n", float64(sum)/float64(SIZE))
}
