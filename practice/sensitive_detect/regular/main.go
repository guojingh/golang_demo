package main

import "fmt"

func main() {
	sensitiveWords := []string{
		"傻逼",
		"傻叉",
		"垃圾",
		"妈的",
		"sb",
	}
	matchContents := []string{
		"什么垃圾打野，傻逼一样，叫你来开龙不来，sb",
	}
	regDemo01()
	fmt.Println("=======================================")
	regDemo02(sensitiveWords, matchContents)
}
