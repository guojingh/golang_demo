package main

import (
	"log"
	"net"
)

// Go语言使用 net.Dial 或 net.DialTimeout 函数发起连接建立请求
func main() {
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	log.Println("dial ok")
}
