package main

import (
	"log"
	"net"
	"time"
)

// 写阻塞
func handleConn(conn net.Conn) {
	defer conn.Close()
	time.Sleep(time.Second * 10)
	for {
		time.Sleep(5 * time.Second)
		var buf = make([]byte, 60000)
		log.Println("start to read from conn")
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("conn read %d bytes, error: %s", n, err)
			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
				continue
			}
			break
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

		go handleConn(c)
	}
}
