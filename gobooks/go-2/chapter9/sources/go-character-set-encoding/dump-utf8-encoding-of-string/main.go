package main

import "fmt"

// 打印字符字面量底层的内存空间来进行验证
func main() {
	var s = "中"
	fmt.Printf("Unicode字符：%s => 其UTF-8内存编码表示为：", s)
	for _, v := range []byte(s) {
		fmt.Printf("0x%X ", v)
	}

	fmt.Printf("\n")
}
