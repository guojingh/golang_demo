package main

import (
	"fmt"
	"io/ioutil"
)

// 写出文件 小文件操作
// io/ioutil下的 ReadFile 方法 --- 不带缓存，一次性
func main() {

	/*	file, err := os.Open("D:/git/test.txt")
		if err != nil {
			fmt.Println("打开文件失败：", err)
		}*/

	//读取文件 --- 这个 ioutil.ReadFile 会在底层调用 os.Open() 打开文件和 file.Close() 关闭文件
	content, err := ioutil.ReadFile("D:/git/test.txt")
	if err != nil {
		fmt.Println("读取文件错误：", err)
	}
	//输出
	fmt.Printf("%v", string(content))

	/*	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件失败：", err)
		}
	}(file)*/
}
