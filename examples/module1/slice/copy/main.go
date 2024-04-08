package main

import (
	"fmt"
	"time"
)

func main() {
	var a []int
	b := make([]int, 1000000)
	c := make([]int, 1000000)

	for i := 0; i < 1000000; i++ {
		a = append(a, i)
	}

	copyTest(a, b)
	dengyuTest(a, c)

}

func copyTest(a []int, b []int) {
	begin := time.Now().UnixMicro()
	copy(b, a)
	end := time.Now().UnixMicro()

	fmt.Println(end - begin)
}

func dengyuTest(a []int, b []int) {
	begin := time.Now().UnixMicro()
	b = a
	end := time.Now().UnixMicro()

	fmt.Println(end - begin)
}
