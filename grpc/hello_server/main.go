package main

import (
	"context"
	"fmt"
	"hello_server/pb"
	"io"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// hello server
type server struct {
	pb.UnimplementedGreeterServer
	mu    sync.Mutex     // count的并发锁
	count map[string]int //记录每个name 的请求访问次数
}

// SayHello 使我们需要实现的方法
// 这个方法是我们对外提供的服务
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count[in.Name]++ //记录用户的请求次数
	//超过一次就返回错误
	if s.count[in.Name] > 1 {
		st := status.New(codes.ResourceExhausted, "Request limit exceed.")
		ds, err := st.WithDetails(
			&errdetails.QuotaFailure{
				Violations: []*errdetails.QuotaFailure_Violation{{
					Subject:     fmt.Sprintf("name:%s", in.Name),
					Description: "限制每个 name 调用一次",
				}},
			},
		)
		if err != nil {
			return nil, st.Err()
		}
		return nil, ds.Err()
	}
	//正常返回响应
	reply := "hello " + in.GetName()
	return &pb.HelloResponse{Reply: reply}, nil
}

/* func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	s.UnarySayHello(ctx, in)
	return &pb.HelloResponse{Reply: "Hello" + in.Name}, nil
} */

func main() {
	//监听本地 8972 端口
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials(),
		//服务器注册 一元拦截器
		grpc.UnaryInterceptor(unaryInterceptor),
		//服务器注册 流式拦截器
		grpc.StreamInterceptor(streamInterceptor),
	)) //创建grpc 服务器
	//pb.RegisterGreeterServer(s, &server{}) //在grpc 服务端注册服务

	//注册服务，注意初始化 count
	pb.RegisterGreeterServer(s, &server{count: make(map[string]int)})
	//启动服务
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve,err: %v", err)
		return
	}

}

// LotsOfReplies 返或使用多种语言打招呼
func (s *server) LotsOfReplies(in *pb.HelloRequest, stream pb.Greeter_LotsOfRepliesServer) error {
	words := []string{
		"你好",
		"hello",
		"こんにちは",
		"안녕하세요",
	}

	for _, word := range words {
		data := &pb.HelloResponse{
			Reply: word + in.GetName(),
		}

		//使用 Send 方法返回多个数据
		if err := stream.Send(data); err != nil {
			return err
		}
	}
	return nil
}

// LotsOfGreetings 接收流式数据
func (s *server) LotsOfGreetings(stream pb.Greeter_LotsOfGreetingsServer) error {
	reply := "你好："
	for {
		//接收客户端发送来的流式数据
		res, err := stream.Recv()
		if err == io.EOF {
			//最终统一回复
			return stream.SendAndClose(&pb.HelloResponse{
				Reply: reply,
			})
		}

		if err != nil {
			return err
		}

		reply += res.GetName()
	}
}

// BidHello 双向流式打招呼
func (s *server) BidHello(stream pb.Greeter_BidHelloServer) error {
	//双向流带 metadata接收
	return s.BidirectionalStreamingSayHello(stream)

	//双向流普通接收
	/* 	for {
		//接收流式数据
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		//对接收的数据做些处理
		reply := magic(in.GetName())

		//返回流式响应
		if err := stream.Send(&pb.HelloResponse{Reply: reply}); err != nil {
			return err
		}
	} */
}

// UnarySayHello rpc 普通调用设置 header 和 trailer
func (s *server) UnarySayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	//通过 defer 设置 tralier
	defer func() {
		trailer := metadata.Pairs("timestamp", strconv.Itoa(int(time.Now().Unix())))
		grpc.SetTrailer(ctx, trailer)
	}()

	//从客户端请求上下文中读取metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "UnarySayHello: failed to get metadata")
	}
	if t, ok := md["token"]; ok {
		fmt.Printf("token form metadata:\n")
		if len(t) < 1 || t[0] != "app-test-ahu" {
			return nil, status.Error(codes.Unauthenticated, "认证失败")
		}
	}

	//创建和发送header
	header := metadata.New(map[string]string{"location": "chengdu"})
	grpc.SendHeader(ctx, header)

	fmt.Printf("request received: %v, say hello ...\n", in)
	return &pb.HelloResponse{Reply: in.Name}, nil

}

// BidirectionalStreamingSayHello 流式 RPC 调用客户端 metadata操作
func (s *server) BidirectionalStreamingSayHello(stream pb.Greeter_BidHelloServer) error {
	//在defer 中创建trailer 记录函数的返回时间
	defer func() {
		trailer := metadata.Pairs("timestamp", strconv.Itoa(int(time.Now().Unix())))
		stream.SetTrailer(trailer)
	}()

	//从 client 读取 metadata
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return status.Errorf(codes.DataLoss, "BidirectionalStreamingSayHello: failed to get metadata")
	}

	if t, ok := md["token"]; ok {
		fmt.Printf("token from metadata:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	}

	//创建和发送 header
	header := metadata.New(map[string]string{"location": "X2Q"})
	stream.SendHeader(header)

	//读取请求数据发送响应数据
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		fmt.Printf("request received %v, sending reply\n", in)
		if err := stream.Send(&pb.HelloResponse{Reply: in.Name}); err != nil {
			return err
		}
	}
}

// magic 对接收的数据进行处理
func magic(s string) string {
	s = strings.ReplaceAll(s, "吗", "")
	fmt.Println(s)
	s = strings.ReplaceAll(s, "吗", "")
	s = strings.ReplaceAll(s, "你", "我")
	s = strings.ReplaceAll(s, "?", "!")
	s = strings.ReplaceAll(s, "?", "!")
	return s
}

// unaryInterceptor 服务器一元拦截器
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// authentication (token verification)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
	}
	if !valid(md["authorization"]) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}
	m, err := handler(ctx, req)
	if err != nil {
		fmt.Printf("RPC failed with error %v\n", err)
	}
	return m, err
}

// streamInterceptor 服务器流拦截器
func streamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// authentication (token verification)
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return status.Errorf(codes.InvalidArgument, "missing metadata")
	}
	if !valid(md["authorization"]) {
		return status.Errorf(codes.Unauthenticated, "invalid token")
	}

	err := handler(srv, newWrappedStream(ss))
	if err != nil {
		fmt.Printf("RPC failed with error %v\n"m err)
	}
	return err
}

type wrappedStream struct {
	grpc.ServerStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	logger("Receive a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	logger("Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.SendMsg(m)
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}