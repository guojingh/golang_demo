package main

import (
	"bytes"
	"fmt"
)

// bytes包提供了 bytes.Equal函数比较两个字节切片是否相等
// 起底层实现也是通过字符串间的比较：string(a) == string(b)
func main() {
	fmt.Println(bytes.Equal([]byte{'a', 'b', 'c'}, []byte{'a', 'b', 'd'}))
	fmt.Println(bytes.Equal([]byte{'a', 'b', 'c'}, []byte{'a', 'b', 'c'}))
	fmt.Println(bytes.Equal([]byte{'a', 'b', 'c'}, []byte{'b', 'a', 'd'}))
	fmt.Println(bytes.Equal([]byte{}, nil))
}
