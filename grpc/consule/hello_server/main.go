package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"

	"hello_server/pb"

	"google.golang.org/grpc/health"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// grpc server

const serviceName = "hello"

var port = 8976

type server struct {
	pb.UnimplementedGreeterServer

	// server 可以添加自定义字段
	addr string
}

// SayHello 是我们需要实现的方法
// 这个方法是我们对外提供的服务
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	reply := s.addr + "hello " + in.GetName()
	return &pb.HelloResponse{Reply: reply}, nil
}

func main() {
	flag.IntVar(&port, "port", 8976, "server port")
	flag.Parse()
	fmt.Printf("-->port:%d\n", port)
	// 启动服务
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("failed to listen, err:%v\n", err)
		return
	}

	s := grpc.NewServer() // 创建grpc服务

	// 获取本机的出口ip
	ipinfo, err := GetOutboundIP()
	if err != nil {
		fmt.Printf("GetOutboundIP failed,err:%v\n", err)
		return
	}
	fmt.Println(ipinfo.String())

	// 注册服务
	addr := fmt.Sprintf("%s:%d", ipinfo.String(), port)
	pb.RegisterGreeterServer(s, &server{addr: addr})
	// 给我们的gRPC服务增加了健康检查的处理逻辑
	healthpb.RegisterHealthServer(s, health.NewServer()) // consul 发来健康检查的RPC请求，这个负责返回OK

	// 连接至consul
	config := api.DefaultConfig()
	config.Address = "172.16.56.134:8500"
	cc, err := api.NewClient(config) // 127.0.0.1:8500
	if err != nil {
		fmt.Printf("api.NewClient failed, err:%v\n", err)
		return
	}

	// 将我们的gRPC服务注册到consul
	// 1.定义我们的服务
	// 配置健康检查策略，告诉consul如何进行健康检查
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", ipinfo.String(), port), // 外网地址
		Timeout:                        "5s",
		Interval:                       "5s", // 间隔
		DeregisterCriticalServiceAfter: "1m", // 1分钟后注销掉不健康的服务节点，一般不建议配置特别小（自动注销）
	}

	serviceID := fmt.Sprintf("%s-%s-%d", serviceName, ipinfo.String(), port) // 服务唯一ID
	srv := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceName,
		Tags:    []string{"qimi"},
		Address: ipinfo.String(),
		Port:    port,
		Check:   check,
	}
	// 2.注册服务到consul
	cc.Agent().ServiceRegister(srv)

	// 启动服务
	go func() {
		if err := s.Serve(l); err != nil {
			fmt.Printf("failed to serve,err:%v\n", err)
			return
		}
	}()

	// Ctrl + C 退出程序
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("wait quit signal...")
	<-quitCh // 没收到信号就阻塞

	// 程序退出的时候注销服务（手动注销）
	fmt.Println("service out...")
	err = cc.Agent().ServiceDeregister(serviceID)
	if err != nil {
		fmt.Printf("ServiceDeregister failed, err:%v\n", err)
	}
}

// GetOutboundIP 获取本机的出口IP
func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}

// 10.22.33.4:80/health   --> HTTP 200
// 127.0.0.1:8972/health  --> gRPC ok
