package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// strings bytes 包提供了快速创建满足 io.Reader 接口的方案
// 利用这两个包的NewReader函数并传入我们的数据域即可创建一个满足io.Reader接口的实例
func main() {
	var buf bytes.Buffer
	var s = "I Love go"

	_, err := io.Copy(&buf, strings.NewReader(s))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", buf.String())

	buf.Reset()
	var b = []byte("I Love go")
	_, err = io.Copy(&buf, bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", buf.String())
}
