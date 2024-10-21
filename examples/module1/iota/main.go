package main

import "fmt"

// 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表的实现
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func reverseString(s *string) {
	str := []rune(*s)
	n := len(str)
	for i := 0; i < n/2; i++ {
		str[i], str[n-i-1] = str[n-i-1], str[i]
	}
	*s = string(str)
}

func main() {

	// 翻转字符串
	str := "hello,worl"
	fmt.Println(10 / 2)

	reverseString(&str)
	fmt.Println(str)
}
