package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 有缓存读取文件，大文件操作
// 需要打开文件 使用，os.Open()
// 读取文件 bufio.ReadString('\n') -- 读取到'\n'字符，也就是一行一行进行读取
// 最后退出循环读取 err == io.EOF, io.EOF 表示文件读取完毕
// 关闭文件 os.Close
func main() {

	//打开文件
	file, err := os.Open("D:/git/test.txt")
	if err != nil {
		fmt.Println("文件打开失败: ", err)
	}

	//关闭文件
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			fmt.Println("关闭文件失败：", err)
		}
	}(file)
	//读取文件
	//创建一个 Reader
	reader := bufio.NewReader(file)
	//读取操作
	for {
		str, err := reader.ReadString('\n')
		//io.EOF 表示已经读取到文件的结尾
		if err == io.EOF {
			break
		}

		//如果没读到文件结尾就正常输出文件内容即可。
		fmt.Println(str)
	}

	//文件读取成功
	fmt.Println("文件读取成功")
}
