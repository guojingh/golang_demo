package main

import (
	"bytes"
	"fmt"
	"strings"
)

//分割

func main() {

	//Fields相关函数
	//Fields 会把 '\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP) 识别为空白字符串
	fmt.Printf("%q\n", strings.Fields("go java python"))
	fmt.Printf("%q\n", strings.Fields("\tgo \f \u0085 \u00a0 java \n\rpython"))
	fmt.Printf("%q\n", strings.Fields(" \t \n\r  "))
	fmt.Printf("%q\n", bytes.Fields([]byte("go java python")))
	fmt.Printf("%q\n", bytes.Fields([]byte("\tgo \f \u0085 \u00a0 java \n\rpython")))
	fmt.Printf("%q\n", bytes.Fields([]byte(" \t \n\r  ")))
	fmt.Println("===============")

	//FieldsFunc函数
	//按自定义逻辑对原字符串进行分割
	splitFunc := func(r rune) bool {
		return r == rune('\n')
	}
	fmt.Printf("%q\n", strings.FieldsFunc("\tgo \f \u0085 \u00a0 java \n\n\rpython", splitFunc))
	fmt.Printf("%q\n", bytes.FieldsFunc([]byte("\tgo \f \u0085 \u00a0 java \n\n\rpython"), splitFunc))

	fmt.Println("==============")
	//Split相关函数
	//Split可以更为通用地对字符串或字节切片进行分割
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a,b,c", "b"))
	fmt.Printf("%q\n", strings.Split("Go社区欢迎你", ""))
	fmt.Printf("%q\n", strings.Split("abc", "de"))
	fmt.Printf("%q\n", strings.SplitN("a,b,c,d", ",", 2))
	fmt.Printf("%q\n", strings.SplitN("a,b,c,d", ",", 3))
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c,d", ","))
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c,d", ",", 2))
}
