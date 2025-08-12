package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var nums []int
var wg sync.WaitGroup

func main() {

	wg.Add(4)
	num1 := rand.Intn(100)
	fmt.Println(num1)
	num2 := rand.Intn(100)
	fmt.Println(num2)

	go func(num int) {
		defer wg.Done()
		nums = append(nums, num1)
	}(num1)

	go func(num int) {
		defer wg.Done()
		nums = append(nums, num1)
	}(num1)

	go func(num int) {
		defer wg.Done()
		nums = append(nums, num2)
	}(num2)

	go func(num int) {
		defer wg.Done()
		nums = append(nums, num2)
	}(num2)

	wg.Wait()
	fmt.Printf("%v\n", nums)
}

type People struct {
	Name string
	Age  int
}
