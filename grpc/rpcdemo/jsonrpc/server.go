package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	service := new(ServiceA)

	// 注册 RPC 服务
	rpc.Register(service)
	len, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatal("listen err:", err)
	}

	for {
		conn, _ := len.Accept()
		//使用 json 协议
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
