package main

import (
	"fmt"
	"io"
	"os"
)

// Read方法会用读取到的内容填充传入的字节切片；
// 如果剩余数据长度小于切片长度，那么也只会使用这些剩余数据填充切片，并返回实际读取的数据长度；
// 只有当剩余长度为0时，才会返回io.EOF表示读到了文件末尾
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

func main() {
	file := "./foo.text"
	text := "hello, gopher"
	buf := make([]byte, 13)

	n, err := directWriteByteSliceToFile(file, []byte(text))
	if err != nil {
		fmt.Println("write file error:", err)
		return
	}
	fmt.Printf("write %d bytes to file.\n", n)

	f, err := os.Open(file)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	defer f.Close()
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("read meets EOF")
				return
			}
			fmt.Println("read file error:", err)
		}
		fmt.Printf("read %d bytes from file: %q\n", n, buf)
	}
}
