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
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/status"
)

var name = flag.String("name", "七米", "通过-name告诉server你是谁")

type wrappedStream struct {
	grpc.ClientStream
}

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

// grpc 客户端
// 调用 server 端的 SayHello 方法
func main() {
	flag.Parse() // 解析命令行参数

	// 连接 server 端
	// 加载证书
	creds, err := credentials.NewClientTLSFromFile("certs/server.crt", "guojinghu.com")
	if err != nil {
		fmt.Printf("credentials.NewClientTLSFromFile err: %v\n", err)
		return
	}
	conn, err := grpc.Dial("127.0.0.1:8972",
		//grpc.WithTransportCredentials(insecure.NewCredentials()),  // 不安全的连接
		grpc.WithTransportCredentials(creds), // 带有证书的连接
		grpc.WithUnaryInterceptor(unaryInterceptor),
		grpc.WithStreamInterceptor(streamInterceptor),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer conn.Close()
	// 创建客户端
	c := pb.NewGreeterClient(conn) //使用生成的Go代码
	// 调用 RPC 方法
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	// 普通的 RPC 调用
	resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		// 收到带 detail 的 error
		s := status.Convert(err)
		for _, d := range s.Details() {
			switch info := d.(type) {
			case *errdetails.QuotaFailure:
				fmt.Printf("Quota failure, message: %v\n", info)
			default:
				fmt.Printf("unexcepted type: %v\n", info)
			}
		}
		log.Printf("c.SayHello failed, err:%v\n", err)
		return
	}

	// 拿到了RPC 响应
	log.Printf("resp:%v\n", resp.GetReply())

	// 调用服务端流式的RPC
	//callLotsOfReplies(ctx, c)
	// 客户端流式 RPC
	//callLotsOfGreetings(ctx, c)
	// 双向流式 RPC
	//runBidiHello(ctx, c)
}

// 服务端流式 RPC
func callLotsOfReplies(ctx context.Context, c pb.GreeterClient) {
	stream, err := c.LotsOfReplies(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Println(err)
		return
	}

	// 依次从流式响应中读取返回的响应数据
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("stream.Recv err: %v", err)
			return
		}
		log.Printf("recv: %v\n", res.GetReply())
	}
}

// 客户端流式 RPC
func callLotsOfGreetings(ctx context.Context, c pb.GreeterClient) {
	// 客户端要流式的发送消息
	stream, err := c.LotsOfGreetings(ctx)
	if err != nil {
		log.Printf("c.LotsOfGreeting err: %v", err)
		return
	}

	names := []string{"张三", "李四", "王二麻子"}
	for _, n := range names {
		stream.Send(&pb.HelloRequest{Name: n})
	}
	// 流式发送结束之后要关闭流
	//stream.CloseSend()
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("stream.CloseAndRecv err: %v", err)
		return
	}

	log.Printf("res: %v\n", res.GetReply())
}

// 双向流模式
func runBidiHello(ctx context.Context, c pb.GreeterClient) {
	steam, err := c.BidHello(ctx)
	if err != nil {
		log.Fatalf("c.BidHello err: %v", err)
	}

	waitc := make(chan struct{})
	go func() {
		for {
			in, err := steam.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("steam.Recv err: %v", err)
			}
			fmt.Printf("AI:%s\n", in.GetReply())
		}
	}()

	// 从标准输入获取用户输入
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	for {
		cmd, _ := reader.ReadString('\n') // 读到换行
		cmd = strings.TrimSpace(cmd)      // 删除尾部空格
		if len(cmd) == 0 {
			continue
		}
		if strings.ToUpper(cmd) == "QUIT" {
			break
		}

		if err = steam.Send(&pb.HelloRequest{Name: cmd}); err != nil {
			log.Fatalf("c.BidHello stream.Send(%v) failed:%v", cmd, err)
		}
	}
	steam.CloseSend()
	<-waitc
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

// streamInterceptor 客户端流式拦截器
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
