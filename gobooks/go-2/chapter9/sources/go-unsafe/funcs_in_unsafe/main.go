package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// sizeof用于获取一个表达式值的大小
	type Foo struct {
		a int
		b string
		c [10]byte
		d float64
	}

	var i int = 5
	var a = [100]int{}
	var sl = a[:]
	var f Foo

	//Alignof用于获取一个表达式的内存地址对齐系数
	fmt.Println(unsafe.Sizeof(i))           // 8
	fmt.Println(unsafe.Sizeof(a))           // 800
	fmt.Println(unsafe.Sizeof(sl))          // 24 (注：返回的是切片描述符的大小)
	fmt.Println(unsafe.Sizeof(f))           // 48
	fmt.Println(unsafe.Sizeof(f.c))         // 10
	fmt.Println(unsafe.Sizeof((*int)(nil))) // 8

	// chapter9/sources/go-unsafe/funcs_in_unsafe.go
	fmt.Println(unsafe.Alignof(i))          // 8
	fmt.Println(unsafe.Alignof(f.a))        // 8
	fmt.Println(unsafe.Alignof(a))          // 8
	fmt.Println(unsafe.Alignof(sl))         // 8
	fmt.Println(unsafe.Alignof(f))          // 8
	fmt.Println(unsafe.Alignof(f.c))        // 1
	fmt.Println(unsafe.Alignof(struct{}{})) // 1 (注：空结构体的对齐系数为1)
	fmt.Println(unsafe.Alignof([0]int{}))   // 8 (注：长度为0的数组，其对齐系数依然与其元素
	// 类型的对齐系数相同)

	// Offsetof用于获取结构体中某字段的地址偏移量（相对于结构体变量的地址）​
	// Offsetof函数应用面较窄，仅用于求结构体中某字段的偏移值
	fmt.Println(unsafe.Offsetof(f.b)) //8
	fmt.Println(unsafe.Offsetof(f.d)) //40
}
