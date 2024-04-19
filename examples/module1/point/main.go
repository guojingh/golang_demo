package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Admin struct {
	Name string
	Age  int
}

// 首先说明 go 中指针类型是不能进行强转换的
// uintptr和 unsafe.Pointer的区别
// unsafe.Pointer 可以进行指针类型的转换
// uintptr可以进行指针类型的运算，那么 unsafe.Pointer 就是可以作为指针运算的中介
func main() {
	admin := Admin{
		Name: "seekload",
		Age:  18,
	}
	ptr := &admin
	//这里由于 name 是第一个字段，因此不需要进行偏移量的计算就能直接操作
	//所以获取name的地址，不需要进行偏移量的运算，直接获取第一个就是
	name := (*string)(unsafe.Pointer(ptr)) // 1

	*name = "四哥"

	s1 := make([]int, 3)
	s2 := []int{1, 2, 3}
	fmt.Println(reflect.DeepEqual(s1, s2))
	copy(s1, s2)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(*ptr)

	// 这里由于 Age 是第二个字段，因此需要加上它的一个偏移量，这里注意 unsafe.Offsetof 返回的是一个uintptr
	// uintptr(unsafe.Pointer(ptr)) 这个是一个对象的起始指针
	//那么就可以对指针进行加减操作，然后将其结果转换为 unsafe.Pointer，那么就可以强制指针类型转换了，最后再给其赋值。
	age := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + unsafe.Offsetof(ptr.Age))) // 2
	*age = 35

	fmt.Println(*ptr)
}
