package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/cncamp/golang/grpc/rpcdemo/netrpc/server/service"
)

func main() {
	//建立 http 连接
	//client, err := rpc.DialHTTP("tcp", "127.0.0.1:9091")
	// 基于 TCP 实现 PRC 调用
	//client, err := rpc.Dial("tcp", "127.0.0.1:9091")
	//基于 JSON 协议
	conn, err := net.Dial("tcp", "127.0.0.1:9091")

	//使用JSON协议
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	if err != nil {
		log.Fatal("dialing", err)
	}

	// 同步调用 client.Call
	args := &service.Args{X: 10, Y: 20}
	var reply int
	err = client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		log.Fatal("ServiceA.Add error:", err)
	}

	fmt.Printf("ServiceA.Add: %d+%d=%d\n", args.X, args.Y, reply)

	//异步调用 client.Go
	var reply2 int
	divCall := client.Go("ServiceA.Add", args, &reply2, nil)
	replayCall := <-divCall.Done
	fmt.Println(replayCall.Error)
	fmt.Println(reply2)
}
