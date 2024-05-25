package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 将 D:\git\test2.txt内容写到 D:\git\test.txt 文件里面
func main() {
	//1.定义源文件
	text2, err := os.OpenFile("D:/git/test2.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("打开文件失败 test2.txt，", err)
	}

	text, err := os.OpenFile("D:/git/test.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件失败test.txt，", err)
	}
	//2.定义目标文件
	//3.关闭文件资源
	defer func(text2 *os.File) {
		err := text2.Close()
		if err != nil {
			fmt.Println("关闭文件 test2 失败")
		}
	}(text2)

	defer func(text *os.File) {
		err := text.Close()
		if err != nil {
			fmt.Println("关闭文件 test 失败")
		}
	}(text)

	//4.复制文件内容
	reader := bufio.NewReader(text2)
	writer := bufio.NewWriter(text)

	for {
		str, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("读取文件失败，", err)
		}

		if err == io.EOF {
			break
		}

		_, err = writer.WriteString(str)
		if err != nil {
			fmt.Println("写入文件失败，", err)
		}
	}

	if err := writer.Flush(); err != nil {
		fmt.Println("写入缓存刷新到 io 接口失败")
	}

	fmt.Println("test2.txt ---> test.txt 文件复制成功！！！")
}
