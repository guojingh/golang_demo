package main

import (
	"add_client/pb"
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var a = flag.Int64("a", 0, "通过-a输入第一个参数")
var b = flag.Int64("b", 0, "通过-b输入第二个参数")

func main() {
	flag.Parse()
	fmt.Println(*a, *b)

	// 连接 server 端
	conn, err := grpc.Dial("127.0.0.1:8973", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}

	defer conn.Close()
	// 创建客户端
	c := pb.NewAdderClient(conn)
	// 调用 RPC 方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 发起普通 RPC 调用
	// 带元数据
	md := metadata.Pairs(
		"token", "app-test-guojinghu",
	)

	ctx = metadata.NewOutgoingContext(ctx, md)

	// 声明两变量
	var header, trailer metadata.MD
	resp, err := c.Add(ctx, &pb.AddRequest{
		A: *a,
		B: *b,
	}, grpc.Header(&header), grpc.Trailer(&trailer))

	if err != nil {
		fmt.Printf("call add failed: %v", err)
		return
	}
	// 拿到响应数据之前可以获取 header
	fmt.Printf("header%v\n", header)

	// 拿到了 RPC 响应
	log.Printf("%d + %d = %d\n", *a, *b, resp.GetC())

	// 拿到了响应数据后可以获取 trailer
	fmt.Printf("trailer%v\n", trailer)
}
