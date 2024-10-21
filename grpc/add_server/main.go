package main

import (
	"add_server/pb"
	"context"
	"fmt"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedAdderServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	// 利用 defer 在发送完响应数据后，发送 trailer
	defer func() {
		trailer := metadata.Pairs(
			"timestamp", strconv.Itoa(int(time.Now().Unix())),
		)
		grpc.SetTrailer(ctx, trailer)
	}()
	// 在执行业务逻辑之前 check metadata 中是否包含 token
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok { // 没有元数据拒接
		return nil, status.Error(codes.Unauthenticated, "无效请求")
	}

	vl := md.Get("token")
	if len(vl) < 1 || vl[0] != "app-test-guojinghu" {
		return nil, status.Error(codes.Unauthenticated, "无效token")
	}

	/*	if vl, ok := md["token"]; ok {
		if len(vl) > 0 && vl[0] == "app-test-guojinghu" {

		}
	}*/
	reply := req.GetA() + req.GetB()
	// 发送数据前发送 header
	header := metadata.New(map[string]string{
		"location": "Beijing",
	})
	grpc.SendHeader(ctx, header)
	fmt.Printf("a=%d, b=%d, c=%d", req.GetA(), req.GetB(), reply)
	return &pb.AddResponse{C: reply}, nil
}

func main() {
	// 启动服务
	l, err := net.Listen("tcp", ":8973")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	// 创建 GRPC 服务
	s := grpc.NewServer()
	// 注册服务
	pb.RegisterAdderServer(s, &server{})
	// 启动服务
	err = s.Serve(l)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}

}
