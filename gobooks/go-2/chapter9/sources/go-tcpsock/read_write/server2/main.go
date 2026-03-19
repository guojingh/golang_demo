package main

import (
	"fmt"
	"log"
	"net"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		//从连接上读取数据
		var buf = make([]byte, 10)
		log.Println("start to read from conn")
		n, err := c.Read(buf)
		if err != nil {
			log.Println("conn read err:", err)
			return
		}
		// buf[:n] 只转换实际读到的数据
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}

func main() {

	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Println("listen error:", err)
		return
	}
	defer l.Close()
	for {
		fmt.Println("等待接收socket链接...")
		c, err := l.Accept()
		if err != nil {
			log.Println("accept error:", err)
			return
		}
		go handleConn(c)
	}
}
