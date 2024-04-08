package main

import (
	"container/list"
	"fmt"
)

var i1 = 0

func main() {

	l1 := list.New()

	l1.PushBack(10)
	fmt.Println(l1.Front().Value)

	recursion := FibonacciWithRecursion(10)

	fmt.Println(recursion)
}

func FibonacciWithRecursion(i int) int {
	if i == 1 || i == 2 {
		return i - 1
	}

	i1++
	return FibonacciWithRecursion(i-1) + FibonacciWithRecursion(i-2)
}
