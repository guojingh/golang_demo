package main

import (
	"log"
	"net"
	"time"
)

// 写阻塞
func main() {
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	log.Println("dial ok")
	data := make([]byte, 65536)
	var total int
	for {
		// 因为客户端开始暂停消费，所以这里会发生阻塞
		n, err := conn.Write(data)
		if err != nil {
			total += n
			log.Printf("write %d bytes, error:%s\n", n, err)
			break
		}
		total += n
		log.Printf("write %d bytes this time, %d bytes in total\n", n, total)
	}
	log.Printf("write %d bytes in total\n", total)
	time.Sleep(time.Second * 10000)
}
