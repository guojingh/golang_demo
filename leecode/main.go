package main

import (
	"fmt"
	"strings"
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

// 102. 二叉树的层序遍历
// 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法解题
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	// 深度
	depth := 0

	var order func(root *TreeNode, depth int)
	order = func(root *TreeNode, depth int) {
		// 跳出条件
		if root == nil {
			return
		}

		// 为当前层次创建内部数组
		if depth == len(res) {
			res = append(res, []int{})
		}

		//将该层数组加入到res中
		res[depth] = append(res[depth], root.Val)
		// 继续向下遍历
		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}
	order(root, depth)
	return res
}

// 148.排序链表
// 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
func sortList(head *ListNode) *ListNode {

	return nil
}

// 108.将有序数组转换成二叉搜索树
// 给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 平衡 二叉搜索树。
// 高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	// 二分法找出中间节点
	mid := len(nums) / 2
	// 以mid为跟节点定义 TreeNode
	treeNode := &TreeNode{Val: nums[mid]}
	// treeNode 左边节点用 [:mid] 填充
	treeNode.Left = sortedArrayToBST(nums[:mid])
	// treeNode 右边节点用 [mid:] 填充
	treeNode.Right = sortedArrayToBST(nums[mid+1:])
	return treeNode
}

// 98. 验证二叉搜索树
// 有效 二叉搜索树定义如下：

// 节点的左子树只包含 严格小于 当前节点的数。
// 节点的右子树只包含 严格大于 当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
func isValidBST(root *TreeNode) bool {
	list := &[]int{}
	// 递归中序遍历获得数组
	rootList(root, list)
	fmt.Printf("list = %v\n", list)
	// 判断list数组是否递增（不可出现重复值）
	for i := 1; i < len(*list); i++ {
		if (*list)[i] <= (*list)[i-1] {
			return false
		}
	}
	return true
}

// 递归遍历数组
func rootList(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	rootList(root.Left, list)
	*list = append(*list, root.Val)
	rootList(root.Right, list)
}

// 230. 二叉搜索树中第 K 小的元素
// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 小的元素（k 从 1 开始计数）。
func kthSmallest(root *TreeNode, k int) int {
	// 中序遍历二叉搜索树为列表
	list := &[]int{}
	rootList(root, list)
	return (*list)[k-1]
}

// 199. 二叉树的右视图
// 给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
func rightSideView(root *TreeNode) []int {
	// 1.层序遍历获取一个二维数组
	res := [][]int{}
	// 深度
	depth := 0

	var order func(root *TreeNode, depth int)
	// 定义函数
	order = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		// 新建该层的二维数组元素
		if depth == len(res) {
			res = append(res, []int{})
		}

		// 将根节点加入数组中
		res[depth] = append(res[depth], root.Val)
		// 遍历其右节点
		order(root.Right, depth+1)
		// 遍历其左节点
		order(root.Left, depth+1)
	}
	order(root, depth)

	fmt.Printf("res = %v\n", res)
	var nodes []int
	// 2.获取二维数组中，每个子数组的第一个元素
	for _, nums := range res {
		nodes = append(nodes, nums[0])
	}

	return nodes
}

// 114. 二叉树展开为链表
// 给你二叉树的根结点 root ，请你将它展开为一个单链表：

// 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。
func flatten(root *TreeNode) {
	res := &[]*TreeNode{}
	preOrder(root, res)
	// 遍历将数组变成单链表
	for i := 0; i < len(*res)-1; i++ {
		(*res)[i].Left = nil
		(*res)[i].Right = (*res)[i+1]
	}
}

// 先序遍历成数组
func preOrder(root *TreeNode, res *[]*TreeNode) {
	if root == nil {
		return
	}
	*res = append(*res, root)
	preOrder(root.Left, res)
	preOrder(root.Right, res)
}

// 105. 从前序与中序遍历序列构造二叉树
// 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历，
// inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
// 假设改二叉树没有重复元素
func buildTree(preorder []int, inorder []int) *TreeNode {
	//
	if len(preorder) == 0 {
		return nil
	}

	treeMap := make(map[int]int)
	for i, num := range inorder {
		treeMap[num] = i
	}
	i := treeMap[preorder[0]]
	//preorder = [3,9,20,15,7]
	//inorder = [9,3,15,20,7]
	root := &TreeNode{Val: preorder[0]}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}

// 35. 搜索插入位置
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 请必须使用时间复杂度为 O(log n) 的算法。
func searchInsert(nums []int, target int) int {
	n := len(nums)
	l := 0
	r := n - 1
	for l <= r {
		mid := l + ((r - l) / 2)
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
	}

	return r + 1
}

// 70. 爬楼梯
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
func climbStairs(n int) int {

	return 0
}

// 46. 全排列
// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
func permute(nums []int) [][]int {

	return nil
}

// 1 2 , 0 1
// func main() {

// 	//preorder = [3,9,20,15,7]
// 	//inorder = [9,3,15,20,7]

// 	// node := &TreeNode{Val: 5}
// 	// node.Left = &TreeNode{Val: 1}
// 	// node.Right = &TreeNode{Val: 4}
// 	// node.Right.Left = &TreeNode{Val: 3}
// 	// node.Right.Right = &TreeNode{Val: 6}
// 	// fmt.Printf("node Tree = %v\n", rightSideView(node))
// 	// preorder := []int{3, 9, 20, 15, 7}
// 	// inorder := []int{9, 3, 15, 20, 7}
// 	// buildTree(preorder, inorder)

// 	nums := []int{1, 2, 3, 4, 5}

// 	result := searchInsert(nums, 2)
// 	fmt.Printf("nums = %d\n", result)
// 	fmt.Println(5 / 2)
// }

// chapter4/sources/method_set_6.go

func foo(b ...byte) {
	fmt.Println(string(b))
}

func concat(sep string, args ...interface{}) string {
	var result string
	for i, v := range args {
		if i != 0 {
			result += sep
		}
		switch v := v.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			result += fmt.Sprintf("%d", v)
		case string:
			result += fmt.Sprintf("%s", v)
		case []int:
			ints := v
			for i, v := range ints {
				if i != 0 {
					result += sep
				}
				result += fmt.Sprintf("%d", v)
			}
		case []string:
			strs := v
			result += strings.Join(strs, sep)
		default:
			fmt.Printf("the argument type [%T] is not supported\n", v)
			return ""
		}
	}
	return result
}

// 报名函数
type record struct {
	name    string
	gender  string
	age     uint16
	city    string
	country string
}

// 通过 ...模拟go语言可变长参数的传值
func enroll(args ...interface{}) (*record, error) {
	if len(args) > 5 || len(args) < 3 {
		return nil, fmt.Errorf("the number of arguments passed is wrong")
	}

	r := &record{
		city:    "Beijing", // 默认值
		country: "china",   // 默认值
	}

	for i, v := range args {
		switch i {
		case 0:
			name, ok := v.(string)
			if !ok {
				return nil, fmt.Errorf("name is not passed as string")
			}
			r.name = name
		case 1:
			gender, ok := v.(string)
			if !ok {
				return nil, fmt.Errorf("gender is not passed as string")
			}
			r.gender = gender
		case 2:
			age, ok := v.(int)
			if !ok {
				return nil, fmt.Errorf("age is not passed as int")
			}
			r.age = uint16(age)
		case 3:
			city, ok := v.(string)
			if !ok {
				return nil, fmt.Errorf("city is not passed as string")
			}
			r.city = city
		case 4:
			country, ok := v.(string)
			if !ok {
				return nil, fmt.Errorf("country is not passed as string")
			}
			r.country = country
		}
	}
	return r, nil
}

// go 选项模式 -- 精装房案例
type FinishedHouse struct {
	style                  int // 0:Chinese 1:American 2:European
	centralAirConditioning bool
	floorMaterial          string // "ground-tile" 或 "wood"
	wallMaterial           string // "latex" 或 "paper" 或 "diatom-mud"
}

// 通过参数暴露配置项
func NewFinishedHouse(style int, centralAirConditioning bool, floorMaterial, wallMaterial string) *FinishedHouse {
	h := &FinishedHouse{
		style:                  style,
		centralAirConditioning: centralAirConditioning,
		floorMaterial:          floorMaterial,
		wallMaterial:           wallMaterial,
	}

	return h
}

// 使用结构体封装配置选项
type Options struct {
	Style                  int // 0:Chinese 1:American 2:European
	CentralAirConditioning bool
	FloorMaterial          string // "ground-tile" 或 "wood"
	WallMaterial           string // "latex" 或 "paper" 或 "diatom-mud"
}

func NewFinishedHouse2(opts *Options) *FinishedHouse {
	// 如果opts为空，则使用默认配置
	var style int = 0
	var centralAirConditioning bool = true
	var floorMaterial string = "wood"
	var wallMaterial string = "paper"

	if opts != nil {
		style = opts.Style
		centralAirConditioning = opts.CentralAirConditioning
		floorMaterial = opts.FloorMaterial
		wallMaterial = opts.WallMaterial
	}

	h := &FinishedHouse{
		style:                  style,
		centralAirConditioning: centralAirConditioning,
		floorMaterial:          floorMaterial,
		wallMaterial:           wallMaterial,
	}

	return h
}

// 使用功能选项函数
type Option func(*FinishedHouse)

func NewFinishedHouse3(option ...Option) *FinishedHouse {
	h := &FinishedHouse{
		style:                  0,
		centralAirConditioning: true,
		floorMaterial:          "wood",
		wallMaterial:           "paper",
	}

	for _, option := range option {
		option(h)
	}

	return h
}

func withStyle(s int) Option {
	return func(h *FinishedHouse) {
		h.style = s
	}
}

func withCentralAirConditioning(c bool) Option {
	return func(h *FinishedHouse) {
		h.centralAirConditioning = c
	}
}

func withFloorMaterial(f string) Option {
	return func(h *FinishedHouse) {
		h.floorMaterial = f
	}
}

func withWallMaterial(w string) Option {
	return func(h *FinishedHouse) {
		h.wallMaterial = w
	}
}

func main() {
	// println(concat("_", 1, 2))
	// println(concat("_", "hello", "gopher"))

	// r, _ := enroll("小明", "男", 10)
	// fmt.Printf("%+v\n", *r)
	// r2, _ := enroll("小红", "女", 20, "Shanghai", "China111")
	// fmt.Printf("%+v\n", *r2)

	// fmt.Printf("%+v\n", NewFinishedHouse(0, true, "wood", "paper"))

	fmt.Printf("%+v\n", NewFinishedHouse2(nil)) //使用默认配置
	fmt.Printf("%+v\n", NewFinishedHouse2(&Options{
		Style:                  1,
		CentralAirConditioning: false,
		FloorMaterial:          "ground-tile",
		WallMaterial:           "latex",
	}))

	fmt.Println("-------")

	fmt.Printf("%+v\n", NewFinishedHouse3()) // 使用默认选项
	fmt.Printf("%+v\n", NewFinishedHouse3(
		withStyle(1),
		withFloorMaterial("ground-tile"),
		withCentralAirConditioning(false),
	))

	var i interface{}
	var e error
	println(i)
	println(e)
	println("i = nil:", i == nil)
	println("e = nil:", e == nil)
	println("i = e:", i == e)
}
