package main

import (
	"log"
	"net"
	"time"
)

// 模拟一个延迟较大的糟糕的网络环境
func main() {
	log.Println("begin dial...")
	conn, err := net.DialTimeout("tcp", ":8888", time.Second*2)
	if err != nil {
		log.Println("dial error", err)
		return
	}
	defer conn.Close()
	log.Println("dial ok...")
}
