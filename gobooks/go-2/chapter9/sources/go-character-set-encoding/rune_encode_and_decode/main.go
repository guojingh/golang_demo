package main

import (
	"fmt"
	"unicode/utf8"
)

// 在Go中，每个rune对应一个Unicode字符的码点，而Unicode字符在内存中的编码表示则放在[​]byte类型中

// rune --> []byte
func encodeRune() {
	var r rune = 0x4E2D // 0x4E2D为Unicode字符中“中”的码点
	buf := make([]byte, 3)
	n := utf8.EncodeRune(buf, r)

	fmt.Printf("the byte slice after encoding rune 0x4E2D is ")
	fmt.Printf("[ ")
	for i := 0; i < n; i++ {
		fmt.Printf("0x%X ", buf[i])
	}
	fmt.Printf("]\n")
	fmt.Printf("the unicode charactor is %s\n", string(buf))
}

// []byte --> rune
func dncodeRune() {
	var buf = []byte{0xE4, 0xB8, 0xAD}
	r, _ := utf8.DecodeRune(buf)
	fmt.Printf("the rune after decoding [0xE4, 0xB8, 0xAD] is 0x%X\n", r)
}

func main() {
	encodeRune()
	dncodeRune()
}
