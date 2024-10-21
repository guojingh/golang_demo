package main

import (
	"context"
	"fmt"
	"hello_server/pb"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type names struct {
	lock     sync.Mutex
	nameList map[string]struct{}
}

var nameList *names

// grep server

type server struct {
	pb.UnimplementedGreeterServer
}
type wrappedStream struct {
	grpc.ServerStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("Receive a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.SendMsg(m)
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

// SayHello 是我们需要实现的业务方法
// 这个方法是我们对外提供的服务
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	nameList.lock.Lock()
	defer nameList.lock.Unlock()
	for k, _ := range nameList.nameList {
		if k == in.GetName() {
			st := status.New(codes.ResourceExhausted, "request limit")
			// 添加错误详情信息，需要接收返回的status
			ds, err := st.WithDetails(
				&errdetails.QuotaFailure{
					Violations: []*errdetails.QuotaFailure_Violation{
						{
							Subject:     fmt.Sprintf("name:%s", in.Name),
							Description: "每个name只能调用一次 SayHello",
						},
					},
				},
			)
			if err != nil { // WithDetails 执行失败
				return nil, st.Err()
			}
			return nil, ds.Err()
		}
	}
	nameList.nameList[in.GetName()] = struct{}{}
	reply := "Hello" + in.GetName()
	return &pb.HelloResponse{Reply: reply}, nil
}

// LotsOfReplies 返回使用多种语言打招呼
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

		// 使用 Send 方法返回多个数据
		if err := stream.Send(data); err != nil {
			return err
		}
	}
	return nil
}

// LotsOfGreetings 接收流式数据
func (s *server) LotsOfGreetings(stream pb.Greeter_LotsOfGreetingsServer) error {
	reply := "你好： "
	for {
		// 接收客户端发来的流式数据
		res, err := stream.Recv()
		if err == io.EOF {
			// 最终同一回复(固定写法)
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

func (s *server) BidHello(stream pb.Greeter_BidHelloServer) error {
	for {
		// 接收流式请求
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		// 对收到的数据进行处理
		reply := magic(in.GetName())

		// 返回流式响应
		if err = stream.Send(&pb.HelloResponse{Reply: reply}); err != nil {
			return err
		}
	}
}

// magic 一段价值连城的“人工智能”代码
func magic(s string) string {
	s = strings.ReplaceAll(s, "吗", "")
	s = strings.ReplaceAll(s, "吧", "")
	s = strings.ReplaceAll(s, "你", "我")
	s = strings.ReplaceAll(s, "？", "!")
	s = strings.ReplaceAll(s, "?", "!")
	return s
}

func main() {
	nameList = &names{
		nameList: make(map[string]struct{}),
	}
	// 启动服务
	l, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen, err:%v\n", err)
		return
	}

	// 加载证书信息
	/*	creds, err := credentials.NewServerTLSFromFile("certs/server.crt", "certs/server.key")
		if err != nil {
			fmt.Printf("Failed to generate credentials %v\n", err)
			return
		}*/

	s := grpc.NewServer(
		//grpc.Creds(creds),
		grpc.Creds(insecure.NewCredentials()),
		//grpc.UnaryInterceptor(unaryInterceptor),
		grpc.StreamInterceptor(streamInterceptor),
	) // 创建 grpc 服务
	//注册服务
	pb.RegisterGreeterServer(s, &server{})
	// 启动服务
	/*	err = s.Serve(l)
		if err != nil {
			fmt.Printf("failed to serve, err:%v\n", err)
			return
		}
	*/
	// 开启 goroutine
	go func() {
		log.Fatalln(s.Serve(l))
	}()

	time.Sleep(time.Second * 5)
	// 下面是 grpc gateway 新加的内容
	// 创建一个连接到我们刚刚启动的 gRPC 服务器的客户端连接
	// gRPC-Gateway 就是通过它来代理请求 （将HTTP请求转换为PRC请求）
	conn, err := grpc.DialContext(
		context.Background(),
		"127.0.0.1:8972",
		grpc.WithBlock(), // 阻塞直到连接成功
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// 注册Greeter
	err = pb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	// 定义 HTTP server 配置
	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://127.0.0.1:8090")
	log.Fatalln(gwServer.ListenAndServe()) // 启动 HTTP 服务

}

// valid 校验认证信息.
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// 执行token认证的逻辑
	// 这里是为了演示方便简单判断token是否与"some-secret-token"相等
	return token == "some-secret-token"
}

// unaryInterceptor 服务端一元拦截器
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

// streamInterceptor 服务端流拦截器
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
		fmt.Printf("RPC failed with error %v\n", err)
	}
	return err
}
