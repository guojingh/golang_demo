package main

import (
	"bufio"
	"fmt"
	"os"
)

//写文件
//1.打开文件，使用 os.OpenFile("文件路径","各种权限"，“”)
//2.defer 关闭文件
//3.bufio.NewWriter() 创建输入流
//4.writer.WriteString() 写入字符串
//5.writer.Flush() 数据刷新到底层的 io 接口

func main() {
	//1.打开文件
	file, err := os.OpenFile("D:/git/test2.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件失败，", err)
	}
	//关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("文件关闭失败")
		}
	}(file)

	//2.写入文件
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		_, err := writer.WriteString("Hello World\n")
		if err != nil {
			return
		}
	}

	//缓存区刷新
	if err := writer.Flush(); err != nil {
		fmt.Println("缓存去内容刷新失败")
	}
}
