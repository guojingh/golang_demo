package main

import (
	"bytes"
	"fmt"
	"strings"
)

// bytes包和strings包提供了EqualFold函数用于进行不区分大小写的Unicode字符的等值比较
// 字节切片在比较时，切片内的字节序列将被解释成字符的UTF-8编码表示后再进行比较
func main() {
	fmt.Println(strings.EqualFold("GoLang", "golang"))
	fmt.Println(bytes.Equal([]byte("GoLang"), []byte("Golang")))
	fmt.Println(bytes.EqualFold([]byte("GoLang"), []byte("golang")))
}
