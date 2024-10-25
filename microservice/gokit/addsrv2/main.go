package main

import (
	"addsrv/pb"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
	"golang.org/x/sync/errgroup"

	//httptransport "github.com/go-kit/kit/transport/http"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials/insecure"
)

// go-kit addService demo3

var (
	httpAddr  = flag.Int("http-addr", 8080, "HTTP端口")
	gGRPCAddr = flag.Int("grpc-addr", 8972, "gRPC端口")
	trimAddr  = flag.String("trim-addr", "127.0.0.1:8975", "trim-service地址")
)

func main() {
	flag.Parse()
	// 前置资源初始化

	srv := NewService()

	// 初始化带 logger 的service
	logger := log.NewJSONLogger(os.Stdout)
	srv = NewLogMiddleware(logger, srv)

	// metrics
	// instrumentation
	fieldKeys := []string{"method", "error"}

	resquestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "addsrv",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)

	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "addsrv",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "addsrv",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method",
	}, []string{}) // not fields here

	// metrics 中间件
	srv = instrumentingMiddleware{
		requestCount:   resquestCount,
		requestLatency: requestLatency,
		countResult:    countResult,
		next:           srv,
	}

	// trim 相关
	// 1. inti grpc client
	// conn, err := grpc.Dial(*trimAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	fmt.Printf("grpc.Dial %s failed, err:%v\n", *trimAddr, err)
	// 	return
	// }

	// defer conn.Close()

	// trimEndpoint := makeTrimEndpoint(conn)

	// 从consul获取 trim service
	trimEndpoint, err := getTrimServiceFromConsul("172.16.56.134:8500", logger, "trim_service", nil)
	if err != nil {
		fmt.Printf("getTrimServiceFromConsul failed, err:%v\n", err)
		return
	}

	srv = NewServiceWithTrim(trimEndpoint, srv)

	var g errgroup.Group

	// HTTP
	g.Go(func() error {
		httpListener, err := net.Listen("tcp", fmt.Sprintf(":%d", *httpAddr))
		if err != nil {
			fmt.Printf("net.Listen %d failed, err:%v\n", *httpAddr, err)
			return err
		}

		defer httpListener.Close()
		// 初始化 go-kit logger
		logger := log.NewLogfmtLogger(os.Stderr)
		httpHandler := NewHTTPServer(srv, logger)

		// http
		//http.Handle("/metric", promhttp.Handler())

		// gin
		httpHandler.(*gin.Engine).GET("/metric", gin.WrapH(promhttp.Handler()))

		return http.Serve(httpListener, httpHandler)

	})

	// gRPC
	g.Go(func() error {
		// gRPC 服务
		grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", *gGRPCAddr))
		if err != nil {
			fmt.Printf("net.Listen %d failed, err:%v\n", *gGRPCAddr, err)
			return err
		}
		defer grpcListener.Close()

		s := grpc.NewServer()
		pb.RegisterAddServer(s, NewGRPCServer(srv))
		return s.Serve(grpcListener)
	})

	// wait
	if err := g.Wait(); err != nil {
		fmt.Printf("server exit with error:%v\n", err)
	}
}
