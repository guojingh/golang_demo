package string

import (
	"fmt"
	"strings"
	"testing"
)

var s1 []string = []string{
	"Rob Pike",
	"Robert Griesemer",
	"Ken Thompson",
}

func concatStringByOperator(s1 []string) string {
	var s string

	for _, v := range s1 {
		s += v
	}

	return s
}

func concatStringBySprintf(s1 []string) string {
	var s string

	for _, v := range s1 {
		s = fmt.Sprintf("%s%s", s, v)
	}

	return s
}

// 将 s1 中的字符串安装 sep 的分隔符连接起来
func concatStringByJoin(s1 []string) string {
	return strings.Join(s1, "")
}

func concatStringByStringBulider(s1 []string) string {
	var b strings.Builder
	for _, v := range s1 {
		b.WriteString(v)
	}

	return b.String()

}

func main() {

}

// 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func TestString(t *testing.T) {
	s := "hello，golang语言"
	fmt.Println(reverseString(s))
	fmt.Println(reverseString(reverseString(s)))
	// output: 言语gnalog，olleh
	// output: hello，golang语言
}
