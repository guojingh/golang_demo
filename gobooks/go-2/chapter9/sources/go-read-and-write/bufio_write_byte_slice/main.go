package main

import (
	"bufio"
	"fmt"
	"os"
)

// 通过包裹类型实现带缓冲I/O
func main() {
	file := "./bufio.txt"
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	defer func() {
		f.Sync()
		f.Close()
	}()
	data := []byte("I love golang!\n")
	//通过包裹函数创建带缓冲I/O的类型
	bio := bufio.NewWriterSize(f, 32) //初始缓存区大小为32字节
	//将15字节写入bio缓冲区，缓冲区缓冲15个字节，bufio.txt内容仍然为空
	bio.Write(data)
	//将15字节写入bio缓冲区，缓冲区缓冲30个字节，bufio.txt内容仍然为空
	bio.Write(data)
	//将15字节写入bio缓冲区后，bufio将32字节写入bufio.txt
	//bio缓冲区仍然有（15*3-32）字节
	bio.Write(data)

	bio.Flush()
}
