package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/encoding/unicode/utf32"
	"golang.org/x/text/transform"
)

// 把GB18030编码数据转换成utf-16和utf-32
func catFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}

func gb18030ToUTF16BE(in []byte) ([]byte, error) {
	r := bytes.NewReader(in)

	s := transform.NewReader(r, simplifiedchinese.GB18030.NewDecoder())
	d := transform.NewReader(s, unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM).NewEncoder())

	out, err := io.ReadAll(d)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func gb18030ToUTF32BE(in []byte) ([]byte, error) {
	r := bytes.NewReader(in)

	s := transform.NewReader(r, simplifiedchinese.GB18030.NewDecoder())
	d := transform.NewReader(s, utf32.UTF32(utf32.BigEndian, utf32.IgnoreBOM).NewEncoder())

	out, err := io.ReadAll(d)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func main() {
	src, err := catFile("../convert_utf8_to_gb18030/gb18030.txt")
	if err != nil {
		log.Println("cat file error:", err)
		return
	}

	//从gb18030到utf-16be
	dst, err := gb18030ToUTF16BE(src)
	if err != nil {
		fmt.Println("convert error:", err)
		return
	}
	fmt.Printf("UTF-16BE(no BOM)编码：")
	for _, v := range dst {
		fmt.Printf("0x%X ", v)
	}
	fmt.Printf("\n")

	//从gb18030到utf-32be
	dst1, err := gb18030ToUTF32BE(src)
	if err != nil {
		fmt.Println("convert error:", err)
		return
	}
	fmt.Printf("UTF-32BE(no BOM)编码：")
	for _, v := range dst1 {
		fmt.Printf("0x%X ", v)
	}
	fmt.Printf("\n")
}
