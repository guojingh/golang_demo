package main

import (
	"log"
	"net"
	"net/rpc"
)

func main() {
	service := new(ServiceA)
	//注册 RPC 服务
	rpc.Register(service)
	len, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	for {
		conn, _ := len.Accept()
		rpc.ServeConn(conn)
	}
}
