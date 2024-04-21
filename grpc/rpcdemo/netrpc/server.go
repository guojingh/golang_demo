package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// server.go 表市注册服务
func main1() {
	service := new(ServiceA)
	rpc.Register(service)
	//基与 HTTP 协议
	rpc.HandleHTTP()
	len, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatal("listen err: ", err)
	}
	http.Serve(len, nil)
}
