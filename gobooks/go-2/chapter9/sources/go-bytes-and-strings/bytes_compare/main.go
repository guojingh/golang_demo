package main

import (
	"bytes"
	"fmt"
)

// bytes包和strings包均提供了Compare方法来对两个字符串/字节切片做排序比较
// 但是字符串通常通过操作符比较更为地道
// 比较时从左到右逐个字符对比这个数值（字节序），遇到第一个不同的位置就决出大小，后面的字符完全不看
func main() {
	var a = []byte{'a', 'b', 'c'}
	var b = []byte{'a', 'b', 'd'}
	var c = []byte{} //empty slice
	var d []byte     //nil slice

	fmt.Println(bytes.Compare(a, b)) //a<b
	fmt.Println(bytes.Compare(b, a)) //b>a
	fmt.Println(bytes.Compare(c, d))
	fmt.Println(bytes.Compare(c, nil))
	fmt.Println(bytes.Compare(d, nil))
	fmt.Println(bytes.Compare(nil, nil))
}
