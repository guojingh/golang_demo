package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"hello_client/pb"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// hello_client
const (
	defaultName = "ahu"
)

var (
	addr = flag.String("addr", "127.0.0.1:8972", "the address to connect to")
	name = flag.String("name", defaultName, "通过-name告诉server你是谁")
)

// 测试 自定义错误的 detail 信息
func main() {
	//解析命令行参数
	flag.Parse()

	//连接 Server
	conn, err := grpc.Dial(*addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		//客户端注册 一元拦截器
		grpc.WithUnaryInterceptor(unaryInterceptor),
		//客户端注册 流式拦截器
		grpc.WithStreamInterceptor(streamInterceptor))
	if err != nil {
		log.Fatalf("grpc.Dial failed,err:%v", err)
		return
	}

	defer conn.Close()
	//创建客户端
	c := pb.NewGreeterClient(conn)
	//调用 RPC 方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		s := status.Convert(err)        //将err 转换成status
		for _, d := range s.Details() { //获取details
			switch into := d.(type) {
			case *errdetails.QuotaFailure:
				fmt.Printf("Qupta failure: %s\n", into)
			default:
				fmt.Printf("Unexpected type: %s\n", into)
			}
		}
		fmt.Printf("c.SayHello failed, err:%v\n", err)
		return
	}
	//拿到 RPC的响应
	log.Printf("resp:%v\n", resp.GetReply())
}

/* func main() {
	flag.Parse()

	//连接到server端，此处禁用安全传输
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err.Error())
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	//执行RPC 调用，并打印收到的相应数据
	ctx, cancal := context.WithTimeout(context.Background(), time.Second)
	defer cancal()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("coult not greet: %s", err.Error())
	}

	log.Printf("Greet: %s", r.GetReply())

	//runLotsOfReplies(c)
	//runLotsOfGreeting(c)
	//runBidHello(c)
	//unaryCallWithMetadata(c, "ahu")
	bidirectionalWithMetadata(c, "ahu")
} */

func runLotsOfReplies(c pb.GreeterClient) {
	//server端流式 RPC
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.LotsOfReplies(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("c.LotsOfReplies failed, err: %v", err)
	}

	for {
		//接收数据端返回的流式数据，当收到io EOF或错误时退出
		res, err := stream.Recv()
		if err != nil {
			break
		}

		if err != nil {
			log.Fatalf("c.LotsOfReplies failed, err: %v", err)
		}
		log.Printf("got reply: %q\n", res.GetReply())
	}
}

// 客户端调用 LotsOfGreetings 方法，向服务端发送流式数据
func runLotsOfGreeting(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//客户端流式RPC
	stream, err := c.LotsOfGreetings(ctx)
	if err != nil {
		log.Fatalf("c.runLotsOfGreeting failed, err: %v", err)
	}

	names := []string{"小明", "小李", "小赵"}
	for _, name := range names {
		//发送流式数据
		err := stream.Send(&pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("c.runLotsOfGreeting stream.Send(%v) failed, err: %v", name, err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("c.runLotsOfGreeting failed: %v", err)
	}

	log.Printf("got reply: %v", res.GetReply())
}

// runBidHello 客户端服务端双向流读取数据
func runBidHello(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	//双向流模式
	stream, err := c.BidHello(ctx)
	if err != nil {
		log.Fatalf("c.BidHello failed, err: %v", err)
	}
	waitc := make(chan struct{})

	//开启协程 用来接收服务器响应
	go func() {
		for {
			//接收服务器返回的相应
			in, err := stream.Recv()
			if err == io.EOF {
				//read done
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("c.BidHello stream.Recv() failed, err: %v", err)
			}
			fmt.Printf("AI: %s\n", in.GetReply())
		}
	}()

	//主协程 --- 从标准输入获取用户输入
	reader := bufio.NewReader(os.Stdin) //从标准输入生成读对象
	for {
		cmd, _ := reader.ReadString('\n') //读到换行
		cmd = strings.TrimSpace(cmd)
		if len(cmd) == 0 {
			continue
		}
		if strings.ToUpper(cmd) == "QUIT" {
			break
		}
		//将数据发送到服务器
		if err := stream.Send(&pb.HelloRequest{Name: cmd}); err != nil {
			log.Fatalf("c.BidHello stream.Send(%v) failed: %v", cmd, err)
		}
	}
	stream.CloseSend()
	<-waitc
}

// unaryCallWithMetadata 普通 RPC 调用客户端 metadata 操作
func unaryCallWithMetadata(c pb.GreeterClient, name string) {
	fmt.Println("--- UnarySayHello client ---")
	//创建 metadata
	md := metadata.Pairs(
		"token", "app-test-ahu",
		"request_id", "1234567",
	)
	//基于 metadata 创建 context
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	//RPC调用
	var header, trailer metadata.MD
	r, err := c.SayHello(
		ctx,
		&pb.HelloRequest{Name: name},
		grpc.Header(&header),   //接收服务端发送来的header
		grpc.Trailer(&trailer), //接收服务器发送来的trailer
	)
	if err != nil {
		log.Fatalf("failed to call SayHello: %v", err)
		return
	}
	//从header中获取location
	if t, ok := header["location"]; ok {
		fmt.Printf("location from header:\n")
		for i, e := range t {
			fmt.Printf("%d. %s\n", i, e)
		}
	} else {
		log.Printf("location expected but doesn't exist in header")
		return
	}
	//获取相应结果
	fmt.Printf("got response: %s\n", r.Reply)
	//从trailer 中获取 timestamp
	if t, ok := trailer["timestamp"]; ok {
		fmt.Printf("timestamp from trailer:\n")
		for i, e := range t {
			fmt.Printf("%d. %s\n", i, e)
		}
	} else {
		log.Printf("timestamp expected but doesn't exist in trailer")
	}
}

// 流式 RPC 调用客户端 metadata 操作
func bidirectionalWithMetadata(c pb.GreeterClient, name string) {
	//创建 metadata 和 context
	md := metadata.Pairs("token", "app-test-ahu")
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	//使用带有 metedata 的context 执行 RPC 调用
	stream, err := c.BidHello(ctx)
	if err != nil {
		log.Fatalf("failed to call BidiHello: %v\n", err)
	}

	go func() {
		// 当 header 到达时读取 header
		header, err := stream.Header()
		if err != nil {
			log.Fatalf("failed to get header from stream: %v", err)
		}

		//从返回响应的header 读取数据
		if l, ok := header["location"]; ok {
			fmt.Printf("location from header:\n")
			for i, e := range l {
				fmt.Printf(" %d. %s\n", i, e)
			}
		} else {
			log.Println("location expected but doesn't exist in header")
			return
		}

		//发送索引请求的数据到server
		for i := 0; i < 5; i++ {
			if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
				log.Fatalf("failed to send streaming: %v\n", err)
			}
		}
		stream.CloseSend()
	}()

	//读取所有响应
	var rpcStatus error
	fmt.Printf("got response:\n")
	for {
		r, err := stream.Recv()
		if err != nil {
			rpcStatus = err
			break
		}
		fmt.Printf(" - %s\n", r.Reply)
	}
	if rpcStatus != io.EOF {
		log.Printf("failed to finish server streaming: %v", rpcStatus)
		return
	}

	//当RPC 结束时读取trailer
	trailer := stream.Trailer()
	//从返回响应的trailer中读取 metadate
	if t, ok := trailer["timestamp"]; ok {
		fmt.Printf("timestamp from trailer:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Printf("timestmp expected but doesn't exise in trailer")
	}

}

// unaryInterceptor 客户端一元拦截器
func unaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	var credsConfigured bool
	for _, o := range opts {
		_, ok := o.(grpc.PerRPCCredsCallOption)
		if ok {
			credsConfigured = true
			break
		}
	}
	if !credsConfigured {
		opts = append(opts, grpc.PerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: "some-secret-token",
		})))
	}
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	end := time.Now()
	fmt.Printf("RPC: %s, start time: %s, end time: %s, err: %v\n", method, start.Format("Basic"), end.Format(time.RFC3339), err)
	return err
}

// 自定义一个拦截器
type wrappedStream struct {
	grpc.ClientStream
}

// wrappedStream 重写 grpc.ClientStream 接口的 RecvMsg 和 SendMsg 方法
func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("Receive a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.SendMsg(m)
}

func newWrappedStream(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedStream{s}
}

func streamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	var credsConfigured bool
	for _, o := range opts {
		_, ok := o.(*grpc.PerRPCCredsCallOption)
		if ok {
			credsConfigured = true
			break
		}
	}
	if !credsConfigured {
		opts = append(opts, grpc.PerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: "some-secret-token",
		})))
	}
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return newWrappedStream(s), nil
}
