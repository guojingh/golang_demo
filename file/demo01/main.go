package main

import (
	"fmt"
	"log"
	"os"
)

// 文件打开，关闭相关操作
// os包下的open函数和os包下的file结构体的close方法
func main() {
	//打开文件
	file, err := os.Open("D:/git/test.txt")

	if err != nil {
		fmt.Println("打开文件失败: ", err)
	}

	log.Println(file)

	//一系列操作

	//关闭文件
	/*	if err := file.Close(); err != nil {
		fmt.Println("关闭文件失败", err)
	}*/

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件失败", err)
		}
	}(file)
}
