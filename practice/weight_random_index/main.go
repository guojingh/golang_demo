package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	  功能：按照指定的一组权重随机返回数组索引
	  参数：weights []float32 权重切片
	  返回：加权随机索引index，index是 0 ~ len(weights)-1 之间的一个整数
	  示例如下：
	  按权重[0.1, 0.2, 0.3, 0.4]随机调用1000次该方法，返回0,1,2,3的次数将接近于1:2:3:4
	  	var weights = []float32{0.1, 0.2, 0.3, 0.4}
	  	var result [4]int
	  	rand.Seed(time.Now().Unix())
	  	for i := 0; i < 1000; i++ {
	  		result[WeightedRandomIndex(weights)]++
	  	}
		fmt.Printf("%v\n", result)
	  输出：
	    [112 174 304 410]
*/
func WeightedRandomIndex(weights []float32) int {
	if len(weights) == 1 {
		return 0
	}
	var sum float32 = 0.0
	for _, w := range weights {
		sum += w
	}
	r := rand.Float32() * sum
	var t float32 = 0.0
	for i, w := range weights {
		t += w
		if t > r {
			return i
		}
	}
	return len(weights) - 1
}

func main() {
	var weights = []float32{0.1, 0.2, 0.3, 0.4, 0.2, 0.2, 0.2}
	var result [7]int
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000; i++ {
		result[WeightedRandomIndex(weights)]++
	}
	fmt.Printf("%v\n", result)
}
