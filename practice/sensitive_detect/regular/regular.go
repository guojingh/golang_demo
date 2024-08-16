package main

import (
	"fmt"
	"regexp"
	"strings"
)

// 正则匹配实现敏感词检测
// 简单实现
func regDemo01() {
	sensitiveWords := []string{
		"傻逼",
		"傻叉",
		"垃圾",
		"妈的",
		"sb",
	}

	text := "什么垃圾打野，傻逼一样，叫你来开龙不来，sb"
	//构造正则匹配字符
	regStr := strings.Join(sensitiveWords, "|")
	fmt.Println("regStr -> ", regStr)
	wordReg := regexp.MustCompile(regStr)
	text = wordReg.ReplaceAllString(text, "*")
	fmt.Println("text -> ", text)
}

// 优化后的 正则匹配实现敏感词检测
func regDemo02(sensitiveWords []string, matchContents []string) {
	//收集匹配到的敏感词
	banWords := make([]string, 0)

	//构造正则匹配字符
	regStr := strings.Join(sensitiveWords, "|")
	//根据正则匹配字符编译一个包含字母和数字的正则表达式
	wordReg := regexp.MustCompile(regStr)
	fmt.Println("regStr -> ", regStr)

	for _, text := range matchContents {
		textBytes := wordReg.ReplaceAllFunc([]byte(text), func(bytes []byte) []byte {
			banWords = append(banWords, string(bytes))
			textRunes := []rune(string(bytes))
			replaceBytes := make([]byte, 0)
			for i, runeLen := 0, len(textRunes); i < runeLen; i++ {
				replaceBytes = append(replaceBytes, byte('*'))
			}
			return replaceBytes
		})
		fmt.Println("srcText     -> ", text)
		fmt.Println("replaceText -> ", string(textBytes))
		fmt.Println("sensitiveWords -> ", banWords)
	}
}
