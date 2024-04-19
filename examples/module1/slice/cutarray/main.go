package main

import "fmt"

//对数组切成切片的测试
func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("原数组", arr)    //原数组 [1 2 3 4 5 6 7 8 9 10]
	fmt.Println("对数组进行截取：===") //对数组进行截取：===
	//指定 max ，max 的大小不能超过要切（数组，切片）的容量

	//max:10 low:2 high:5
	s1 := arr[2:5:10]
	//len = high - low; cap = max - low
	fmt.Printf("数组截取之后的类型%T, 长度%d, 容量%d\n", s1, len(s1), cap(s1)) //数组截取之后的类型[]int, 长度3, 容量8

	fmt.Println("原切片")
	fmt.Println("对切片进行截取===")
	//low:1 high:7 未指定 max
	s2 := s1[1:7]
	//未指定 max 的情况 len = high - low; cap = high - low
	fmt.Printf("切片截取之后的类型%T, 长度%d, 容量%d\n", s2, len(s2), cap(s2)) //切片截取之后的类型[]int, 长度6, 容量7

	//利用数组创建切片，切片操作的是同一个底层数组
	s1[0] = 9999
	s1[1] = 8888

	fmt.Println(arr) // [1 2 9999 8888 5 6 7 8 9 10]
}
