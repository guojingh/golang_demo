package main

import (
	"fmt"
	"os"
)

// os.File结构体类型实现了io.Reader io.Writer 接口，因此可通过Reader，Writer方法直接读写字节序列
func directWriteByteSliceToFile(path string, data []byte) (int, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file error:", err)
		return 0, err
	}

	defer func() {
		// 强制把内核缓冲区里的数据立刻写入到磁盘，不等系统调度，为了数据不丢失
		f.Sync()
		f.Close()
	}()
	return f.Write(data)
}

func directReadByteSliceFromFile(path string, data []byte) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("open file error:", err)
		return 0, err
	}

	defer f.Close()
	return f.Read(data)
}

func main() {
	file := "./foo.text"
	text := "hello, gopher"
	buf := make([]byte, 20)

	n, err := directWriteByteSliceToFile(file, []byte(text))
	if err != nil {
		fmt.Println("write file error:", err)
		return
	}
	fmt.Printf("write %d bytes to file .\n", n)

	n, err = directReadByteSliceFromFile(file, buf)
	if err != nil {
		fmt.Println("read file error:", err)
		return
	}

	//%q 是带转义的字符串格式，会把不可见字符转义显示出来
	fmt.Printf("read %d bytes from file: %q\n", n, buf)
}
