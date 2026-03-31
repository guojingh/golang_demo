package main

import (
	"compress/gzip"
	"fmt"
	"os"
)

// 通过包裹类型实现数据压缩/解压缩操作
// Go标准库中的compress/gzip包提供了这样的包裹函数与包裹类型
func main() {
	file := "./hello_gopher.gz"
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	defer f.Close()

	zw := gzip.NewWriter(f)
	defer zw.Close()
	_, err = zw.Write([]byte("hello, gopher! I love golang!"))
	if err != nil {
		fmt.Println("write compressed data error:", err)
		return
	}
	fmt.Println("write compressed data ok")
}
