package gee

import "net/http"

// HandlerFunc defines the request handler userd by gee
type HandlerFunc func(*Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router *router
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// POST defines the method to add GET reques
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// GET defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	//      1
	//   2     3
	//   4   5   6

	//  2 4 1 5 3 6
	var nodeList []int
	// 将二叉树进行中序遍历
	// 1.使用递归实现
	helper(root, &nodeList)
	return nodeList
}

func helper(node *TreeNode, nodeList *[]int) {
	if node == nil {
		return
	}

	helper(node.Left, nodeList)
	*nodeList = append(*nodeList, node.Val)
	helper(node.Right, nodeList)
}
