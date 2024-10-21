package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	service2 "github.com/cncamp/golang/grpc/rpcdemo/netrpc/server/service"
)

// server.go 表市注册服务
func main() {
	service := new(service2.ServiceA)
	rpc.Register(service)
	//基与 HTTP 协议
	//rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatal("listen err: ", err)
	}
	//http.Serve(l, nil)

	// 基于 tcp 的 RPC
	for {
		conn, _ := l.Accept()
		//rpc.ServeConn(conn)

		//使用 JSON 协议
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
