package main

import (
	"bytes"
	"fmt"
	"strings"
)

// 修剪与变换
// Go标准库的bytes包和strings包提供了一系列Trim API可以辅助你完成对输入数据的修剪
// TrimSpace
// TrimSpace函数去除输入字符串/字节切片首部和尾部的空白字符，它对空白字符的定义与前面Fields函数采用的空白字符定义相同
func main() {
	// TrimSpace(string)
	fmt.Println(strings.TrimSpace("\t\n\f I love Go!!! \n\r"))
	// TrimSpace(byte)
	fmt.Printf("%q\n", bytes.TrimSpace([]byte("\t\n\f I love Go!!! \n\r")))

	//Trim函数允许我们自定义要修剪掉的字符集合
	//Trim、TrimLeft、TrimRight(string)
	fmt.Println(strings.Trim("\t\n fffI love go!!\n \rfff", "\t\n\r f"))
	fmt.Printf("%q\n", strings.TrimLeft("\t\n fffI love go!!!\n \rfff", "\t\n\r f"))
	fmt.Printf("%q\n", strings.TrimRight("\t\n fffI love go!!!\n \rfff", "\t\n\r f"))

	//Trim、TrimLeft、TrimRight([]byte)
	fmt.Printf("%q\n", bytes.Trim([]byte("\t\n fffI love go!!!\n \rfff"), "\t\n\r f"))
	fmt.Printf("%q\n", bytes.TrimLeft([]byte("\t\n fffI love go!!!\n \rfff"), "\t\n\r f"))
	fmt.Printf("%q\n", bytes.TrimRight([]byte("\t\n fffI love go!!!\n \rfff"), "\t\n\r f"))

	//TrimPrefix和TrimLeft的区别
	//TrimLeft的第二个参数应理解为一个字符的集合，而TrimPrefix的第二个参数应理解为一个整体的字符串。通过下面的例子我们就很容易理解
	//TrimSuffix和TrimRight亦是同理
	fmt.Printf("%q\n", strings.TrimLeft("prefix,prefix I love go!!", "prefix,"))
	fmt.Printf("%q\n", strings.TrimPrefix("prefix,prefix I love go!!", "prefix,"))

	// 变换
	// 大小写变换
	// strings包和bytes包提供了ToUpper和ToLower函数
	// ToUpper ToLower(string)
	fmt.Printf("%q\n", strings.ToUpper("i LoVe gOlaNg!!"))
	fmt.Printf("%q\n", strings.ToLower("i LoVe gOlaNg!!"))

	// ToUpper ToLower(bytes)
	fmt.Printf("%q\n", bytes.ToUpper([]byte("i LoVe gOlaNg!!")))
	fmt.Printf("%q\n", bytes.ToLower([]byte("i LoVe gOlaNg!!")))

	//Go标准库在strings和bytes包中提供了Map函数
	//该函数用于将原字符串/字节切片中的部分数据按照传入的映射规则变换为新数据
	trans := func(r rune) rune {
		switch {
		case r == 'p':
			return 'g'
		case r == 'y':
			return 'o'
		case r == 't':
			return 'l'
		case r == 'h':
			return 'a'
		case r == 'o':
			return 'n'
		case r == 'n':
			return 'g'
		}
		return r
	}

	fmt.Printf("%q\n", strings.Map(trans, "I like python!!!"))
	fmt.Printf("%q\n", bytes.Map(trans, []byte("I like python!!!")))
}
