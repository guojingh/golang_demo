package main

import (
	"fmt"
)

// 73. 矩阵置零
func setZeroes(matrix [][]int) {
	lineHasZero := false
	rowHasZero := false
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			lineHasZero = true
			break
		}
	}
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			rowHasZero = true
			break
		}
	}
	// 标记行列
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	// 置零
	for i := 1; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			for j := 1; j < len(matrix); j++ {
				matrix[j][i] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < len(matrix[i]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	if rowHasZero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
	if lineHasZero {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
	fmt.Printf("matrix=%v, lineHasZero=%t, rowHasZero=%t\n", matrix, lineHasZero, rowHasZero)
}

// 54. 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 {
		return res
	}
	// 定义上下左右边界
	left, right, up, down := 0, len(matrix[0])-1, 0, len(matrix)-1
	// 定义横纵坐标
	var x, y int
	// 基础条件
	for left <= right && up <= down {
		// 向右
		for y = left; y <= right && avoid(left, right, up, down); y++ {
			res = append(res, matrix[x][y])
		}
		// 超出边界回退
		y--
		// 边界收缩
		up++
		// 向下
		for x = up; x <= down && avoid(left, right, up, down); x++ {
			res = append(res, matrix[x][y])
		}
		x--
		right--
		// 向左
		for y = right; y >= left && avoid(left, right, up, down); y-- {
			res = append(res, matrix[x][y])
		}
		y++
		down--
		// 向上
		for x = down; x >= up && avoid(left, right, up, down); x-- {
			res = append(res, matrix[x][y])
		}
		x++
		left++
	}
	return res
}

// 碰壁条件
func avoid(left, right, up, down int) bool {
	return up <= down && left <= right
}

// 48. 旋转图像
// 给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
// 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：matrix = [[7,4,1],[8,5,2],[9,6,3]]

/*
	      1 2 3.                                        7 4 1
		  4 5 6.										8 5 2
		  7 8 9											9 6 3
*/
func rotate(matrix [][]int) {
	for i := 0; i < len(matrix)/2; i++ {
		start := i
		end := len(matrix) - i - 1
		for j := 0; j < end-start; j++ {
			matrix[start][start+j], matrix[end-j][start], matrix[end][end-j], matrix[start+j][end] = matrix[end-j][start], matrix[end][end-j], matrix[start+j][end], matrix[start][start+j]
		}
	}
}

// 2. 两数相加
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

// 2 ---- 4 ---- 3
// 5 ---- 6 ---- 4
// 7 ---- 0 ---- 8

// [9,9,9,9,9,9,9]

// [9,9,9,9]

// 8 9 9 9 0 0 0 1

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 进位标志
	count := 0
	// 同位数之和
	sum := 0
	// 结果链表
	head := ListNode{}
	// 当前指针
	cur := &head
	for l1 != nil && l2 != nil {
		sum = l1.Val + l2.Val + count
		if sum >= 10 {
			count = 1
		} else {
			count = 0
		}
		cur.Next = &ListNode{Val: sum % 10}
		// 指针后移
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
		// 和置零
		sum = 0
	}

	// l1链表未遍历完
	for l1 != nil {
		sum = l1.Val + count
		if sum >= 10 {
			count = 1
		} else {
			count = 0
		}
		cur.Next = &ListNode{Val: sum % 10}
		// 指针后移
		cur = cur.Next
		l1 = l1.Next
		// 和置零
		sum = 0
	}

	// l2链表未遍历完
	for l2 != nil {
		sum = l2.Val + count
		if sum >= 10 {
			count = 1
		} else {
			count = 0
		}
		cur.Next = &ListNode{Val: sum % 10}
		// 指针后移
		cur = cur.Next
		l2 = l2.Next
		// 和置零
		sum = 0
	}

	if count != 0 {
		cur.Next = &ListNode{Val: count}
	}

	return head.Next
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	// 进位标志
	count := 0
	// 同位数之和
	sum := 0
	// 结果链表
	head := ListNode{}
	// 当前指针
	cur := &head
	var a, b int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			a = l1.Val
		} else {
			a = 0
		}
		if l2 != nil {
			b = l2.Val
		} else {
			b = 0
		}
		sum = a + b + count
		count = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		// 指针后移
		cur = cur.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		// 和置零
		sum = 0
	}

	return head.Next
}

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 定义一个指针
	var cur *ListNode
	cur = head
	// 计算链表长度
	len := 0
	for cur != nil {
		len++
		cur = cur.Next
	}
	// 定位到待删除节点的前一个节点
	dummy := &ListNode{Next: head}
	cur = dummy
	for i := 0; i < len-n; i++ {
		cur = cur.Next
	}
	// 删除节点
	cur.Next = cur.Next.Next
	return dummy.Next
}

// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
// 输入：head = [1,2,3,4]
// 输出：       [2,1,4,3]
func swapPairs(head *ListNode) *ListNode {
	// 处理极限条件
	if head == nil || head.Next == nil {
		return head
	}
	// 定义双指针
	var slow, fast, dummy, tmp *ListNode
	dummy = &ListNode{Next: head}
	slow = dummy
	fast = head
	for fast.Next != nil {
		// 将双指针后面的节点记下来
		tmp = fast.Next.Next
		slow.Next = fast.Next
		fast.Next.Next = fast
		fast.Next = tmp
		// 双指针后移
		fast = fast.Next
		slow = slow.Next.Next
	}
	return dummy.Next
}

/*
给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。

构造这个链表的 深拷贝。 深拷贝应该正好由 n 个 全新 节点组成，其中每个新节点的值都设为其对应的原节点的值。
新节点的 next 指针和 random 指针也都应指向复制链表中的新节点，并使原链表和复制链表中的这些指针能够表示相同的链表状态。
复制链表中的指针都不应指向原链表中的节点 。

例如，如果原链表中有 X 和 Y 两个节点，其中 X.random --> Y 。那么在复制链表中对应的两个节点 x 和 y ，同样有 x.random --> y 。

返回复制链表的头节点。

用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：

val：一个表示 Node.val 的整数。
random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。
你的代码 只 接受原链表的头节点 head 作为传入参数。
*/

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	// 利用 哈希表解决 想不到啊
	if head == nil {
		return head
	}

	// 创建一个hash用来存放 p1 ---> p2
	mapNode := make(map[Node]*Node, 0)
	// 记录原链表头节点
	p1 := head
	// 设置新链表头节点前的虚拟节点
	temp := &Node{Val: -1}
	p2 := temp
	// 将旧链表放到hash的key中，value 设置新链表对应节点
	for p1 != nil {
		node := &Node{Val: p1.Val}
		mapNode[*p1] = node
		p1 = p1.Next
	}
	// 指针归位进行第二次循环遍历
	p1 = head
	for p1 != nil {
		// 拿到 p1 --> p2
		temp := mapNode[*p1]
		// 设置 p2节点Random指针指向的节点
		temp.Random = mapNode[*p1.Random]
		// 将p2链表串联起来
		p2.Next = temp
		// 指针后移
		p1 = p1.Next
		p2 = p2.Next
	}
	return temp.Next
}

func main() {

	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}

	res := swapPairs(head)
	fmt.Printf("res=%v\n", res)
}
