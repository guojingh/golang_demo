package main

import (
	"bookstore/pb"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// BookStore
func main() {

	//连接数据库
	db, err := NewDB()
	if err != nil {
		fmt.Printf("NewDB err: %v\n", err)
		return
	}

	// 创建 server
	srv := server{
		bs: &bookstore{db: db},
	}

	// 启动 rpc 服务
	l, err := net.Listen("tcp", ":8091")
	if err != nil {
		fmt.Printf("failed to listen, err:%v\n", err)
		return
	}

	s := grpc.NewServer()
	//注册服务
	pb.RegisterBookstoreServer(s, &srv)

	// 同一个端口分别处理gRPC和HTTP
	// 1.创建gRPC-Gateway mux
	gwmux := runtime.NewServeMux()
	dops := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err = pb.RegisterBookstoreHandlerFromEndpoint(context.Background(), gwmux, "127.0.0.1:8091", dops); err != nil {
		log.Fatalf("failed to register grpc gateway, err:%v\n", err)
		return
	}

	// 2.新建HTTP mux
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	// 3.定义HTTP的 Server
	gwServer := &http.Server{
		Addr:    "127.0.0.0:8091",
		Handler: grpcHandlerFunc(s, mux),
	}

	// 4.启动
	fmt.Println("serving on 127.0.0.1:8091...")
	gwServer.Serve(l)
}

// grpcHandlerFunc 将gRPC请求和HTTP请求分别调用不同的handler处理
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}
