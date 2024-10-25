package main

import (
	"addsrv/pb"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"golang.org/x/time/rate"
)

// transport

// 网络传输相关的，包括协议（HTTP，gPRC，thrift...）等
// 3. transport

// decode
// 请求来了之后根据协议（HTTP、HTTP2）和编码（JSON、pb、thrift）去解析数据
func decodeSumRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request SumReqest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func decodeConcatRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request ConcatReqest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

// 编码
// 把响应数据返回 按协议和编码返回
// w：代表响应的网络句柄
// response：业务层返回的响应数据
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func NewHTTPServer(svc AddService, logger log.Logger) http.Handler {
	// HTTP JSON 服务
	sum := makeSumEndpoint(svc)
	// go-kit/log
	//log.With(logger, "method", "sum")  派生子logger的效果
	sum = loggingMiddleware(log.With(logger, "method", "sum"))(sum)
	// 使用限流中间件
	sum = rateMiddleware(rate.NewLimiter(1, 1))(sum)

	sumHandler := httptransport.NewServer(
		sum, // 日志中间件包一层的sum endpoint
		decodeSumRequest,
		encodeResponse,
	)

	concatHandler := httptransport.NewServer(
		makeConcatEndpoint(svc),
		decodeConcatRequest,
		encodeResponse,
	)

	// github.com/gorilla/mux
	// r := mux.NewRouter()

	// r.http.Handle("/sum", sumHandler).Methods("POST")
	// r.http.Handle("/concat", concatHandler).Methods("POST")

	// use gin
	r := gin.Default()
	r.POST("/sum", gin.WrapH(sumHandler))
	r.POST("/concat", gin.WrapH(concatHandler))

	return r
}

// gRPC

type grpcServer struct {
	pb.UnimplementedAddServer

	sum    grpctransport.Handler
	concat grpctransport.Handler
}

func NewGRPCServer(svc AddService) pb.AddServer {
	return &grpcServer{
		sum: grpctransport.NewServer(
			makeSumEndpoint(svc), // endpoint
			decodeGRPCSumRequest,
			encodeGRPCSumResponse,
		),

		concat: grpctransport.NewServer(
			makeConcatEndpoint(svc),
			decodeGRPCConcatRequest,
			encodeGRPCConcatResponse,
		),
	}
}

func (s grpcServer) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	_, resp, err := s.sum.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.SumResponse), nil
}

func (s grpcServer) Concat(ctx context.Context, req *pb.ConcatRequest) (*pb.ConcatResponse, error) {
	_, resp, err := s.concat.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.ConcatResponse), nil
}

// gRPC请求与响应
// decodeGRPCSumRequest 将Sum方法的gRPC请求参数转为内部的SumRequest
func decodeGRPCSumRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.SumRequest)
	return SumReqest{A: int(req.A), B: int(req.B)}, nil
}

// decodeGRPCConcatRequest 将Concat方法的gRPC请求参数转为内部的ConcatRequest
func decodeGRPCConcatRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ConcatRequest)
	return ConcatReqest{A: req.A, B: req.B}, nil
}

// encodeGRPCSumResponse 封装Sum的gRPC响应
func encodeGRPCSumResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(SumResponse)
	return &pb.SumResponse{V: int64(resp.V), Err: resp.Err}, nil
}

// encodeGRPCConcatResponse 封装Concat的gRPC响应
func encodeGRPCConcatResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(ConcatResponse)
	return &pb.ConcatResponse{V: resp.V, Err: resp.Err}, nil
}

// trim相关

// encodeTrimRequest 将内部结构体转为protobuf格式
// 对发起gRPC请求
func encodeTrimRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(trimRequest)
	return &pb.TrimRequest{S: req.s}, nil
}

// decodeTrimResponse 将收到的gRPC响应转为内部的响应结构体
func decodeTrimResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.TrimResponse)
	return &pb.TrimResponse{S: resp.S}, nil
}
