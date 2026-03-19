package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// socket 关闭，socket中还有未读取的数据
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		time.Sleep(time.Second * 10)
		var buf = make([]byte, 10)
		log.Println("start to read from conn")
		n, err := c.Read(buf)
		if err != nil {
			log.Println("conn read error:", err)
			return
		}
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
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
			log.Println("accept err:", err)
			break
		}

		fmt.Println("create a new goroutine headle conn")
		go handleConn(c)
	}
}
