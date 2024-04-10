package main

import (
	"log"
	"net/rpc"
)

func main() {
	//建立 http 连接
	_, err := rpc.Dial("tcp", "127.0.0.1:9091")
	if err != nil {
		log.Fatal("dialing", err)
	}
}