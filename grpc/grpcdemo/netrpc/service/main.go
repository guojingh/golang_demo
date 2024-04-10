package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	service := new(ServiceA)
	rpc.Register(service)
	rpc.HandleHTTP() //基与 HTTP 协议
	l, e := net.Listen("tcp", ":9091")
	if e != nil {
		log.Fatal("listen err: ", e)
	}
	http.Serve(l, nil)
}