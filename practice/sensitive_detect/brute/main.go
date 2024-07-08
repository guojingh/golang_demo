package main

import (
	"fmt"
	"strings"
)

// 敏感词检测，暴力方法实现
func main() {
	sensitiveWords := []string{
		"傻逼",
		"傻叉",
		"垃圾",
		"妈的",
		"sb",
	}
	text := "什么垃圾打野，傻逼一样，叫你来开龙不来，sb"

	for _, word := range sensitiveWords {
		replaceChar := ""

		//先将 string 转换成 rune 类型，再求长度，更精确
		for i, wordLen := 0, len([]rune(word)); i < wordLen; i++ {
			//构造敏感词长度
			replaceChar += "*"
		}
		text = strings.Replace(text, word, replaceChar, -1)
	}

	fmt.Println("text =>", text)
}
