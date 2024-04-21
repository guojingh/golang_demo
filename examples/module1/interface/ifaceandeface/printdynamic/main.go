package main

import (
	"fmt"
	"unsafe"
)

// 输出动态类型好动态值的方法
// 自定义一个 iface 类型
type iface struct {
	itab, data uintptr
}

func main() {
	var a interface{} = nil

	//定义一个类型为 *int 值为 nil 的接口
	var b interface{} = (*int)(nil)

	x := 5
	var c interface{} = (*int)(&x)

	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))
	ic := *(*iface)(unsafe.Pointer(&c))

	fmt.Println(ia, ib, ic)
	//先将 ic.data-uintptr 转换成 unsafe.Pointer 再将其转换成 int 指针最后取指得到值
	fmt.Println(*(*int)(unsafe.Pointer(ic.data)))
}
