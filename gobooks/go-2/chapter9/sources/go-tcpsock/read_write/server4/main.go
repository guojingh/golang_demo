package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// 读操作超时的情况
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		time.Sleep(5 * time.Second)
		var buf = make([]byte, 65536)
		log.Println("start to read from conn")
		c.SetReadDeadline(time.Now().Add(time.Microsecond * 5))
		n, err := c.Read(buf)
		if err != nil {
			log.Printf("conn read %d bytes, error:%s\n", n, err)
			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
				continue
			}
		}
		log.Printf("conn read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Println("listen err:", err)
		return
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("accept error:", err)
			break
		}
		fmt.Println("create a new goroutine headle conn")
		go handleConn(c)
	}
}
