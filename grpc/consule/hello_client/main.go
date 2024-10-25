package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	//"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"hello_client/pb"

	_ "github.com/mbobakov/grpc-consul-resolver"
)

// grpc 客户端
// 调用server端的 SayHello 方法

var name = flag.String("name", "七米", "通过-name告诉server你是谁")

func main() {
	flag.Parse() // 解析命令行参数

	conn, err := grpc.Dial(
		"consul://172.16.56.134:8500/hello?healthy=true", // grpc 中使用consul名称解析器
		// 指定负载均衡策略，这里使用的是gRPC自带的round_robin策略
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("grpc.Dial failed,err:%v", err)
		return
	}
	defer conn.Close()
	// 创建客户端
	c := pb.NewGreeterClient(conn) // 使用生成的Go代码
	// 调用RPC方法
	for i := 0; i < 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
		if err != nil {
			fmt.Printf("c.SayHello failed, err:%v\n", err)
			return
		}

		// 拿到了RPC响应
		fmt.Printf("resp:%v\n", resp.GetReply())
	}
}
