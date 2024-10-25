package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"trim_service/pb"

	apiconsul "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

const serviceName = "trim_service"

var (
	port       = flag.Int("port", 8975, "service port")
	consulAddr = flag.String("consul", "172.16.56.134:8500", "consul address")
)

// trim service
type server struct {
	pb.UnimplementedTrimServer
}

func (s *server) TrimSpace(_ context.Context, req *pb.TrimRequest) (*pb.TrimResponse, error) {
	ov := req.GetS()
	v := strings.ReplaceAll(ov, " ", "")
	fmt.Printf("ov:%#v v:%#v\n", ov, v)
	return &pb.TrimResponse{S: v}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterTrimServer(s, &server{})

	// 注册服务
	cc, err := NewConsulClient(*consulAddr)
	if err != nil {
		fmt.Printf("failed to NewConsulClient: %v\n", err)
		return
	}
	ipInfo, err := getOutboundIP()
	if err != nil {
		fmt.Printf("getOutboundIP failed, err:%v\n", err)
		return
	}

	if err := cc.RegisterServie(serviceName, ipInfo.String(), *port); err != nil {
		fmt.Printf("Register service failed, err:%v\n", err)
		return
	}

	go func() {
		if err := s.Serve(lis); err != nil {
			fmt.Printf("serve failed, err:%v\n", err)
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 提出时注销程序
	if err := cc.Deregister(fmt.Sprintf("%s-%s-%d", serviceName, ipInfo.String(), *port)); err != nil {
		fmt.Printf("Deregister service failed, err:%v\n", err)
		return
	}
}

type ConsulClient struct {
	client *apiconsul.Client
}

// NewConsulClient 新建 consulClient
func NewConsulClient(consulAddr string) (*ConsulClient, error) {
	cfg := apiconsul.DefaultConfig()
	cfg.Address = consulAddr
	client, err := apiconsul.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &ConsulClient{client}, nil
}

// RegisterServie 注册服务
func (c *ConsulClient) RegisterServie(serviceName, ip string, port int) error {
	srv := &apiconsul.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", serviceName, ip, port),
		Name:    serviceName,
		Tags:    []string{"qimi", "trim"},
		Address: ip,
		Port:    port,
	}

	return c.client.Agent().ServiceRegister(srv)
}

// Deregister 注销服务
func (c *ConsulClient) Deregister(serviceID string) error {
	return c.client.Agent().ServiceDeregister(serviceID)
}

// 获取本机出口IP
func getOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}
