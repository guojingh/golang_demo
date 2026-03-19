package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// Go标准库没有直接提供简体中文编码与UTF-8编码之间的转换实现，
// 但Go标准库依赖的golang.org/x/text模块中提供了相关转换实现

func utf8ToGB18030(in []byte) ([]byte, error) {
	if !utf8.Valid(in) {
		return nil, errors.New("invalid utf-8 runes")
	}

	r := bytes.NewReader(in)
	t := transform.NewReader(r, simplifiedchinese.GB18030.NewEncoder())
	out, err := io.ReadAll(t)
	if err != nil {
		return nil, err
	}
	return out, err
}

func dumpToFile(in []byte, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(in)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var src = "中国人" // "\u4E2D\u56FD\u4EBA"
	var dst []byte

	for i, v := range src {
		fmt.Printf("Unicode字符：%s <=> 码点（rune）：%X <=> utf8编码内存表示：", string(v), v)
		//string类型做切片操作，操作的单位是字节，但是返回的结果是string
		s := src[i : i+3]
		for _, v := range []byte(s) {
			fmt.Printf("0x%X ", v)
		}

		t, err := utf8ToGB18030([]byte(s))
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}
		fmt.Printf("gb18030编码内存表示：")
		for _, v := range t {
			fmt.Printf("0x%X ", v)
		}
		fmt.Printf("\n")
		dst = append(dst, t...)
	}

	dumpToFile(dst, "gb18030.txt")
}
