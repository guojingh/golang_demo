package main

import (
	"addsrv/pb"
	"context"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

// endpoint

// 一个 endponit 表示对外提供的一个方法
// 1.3.请求和响应
type SumReqest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SumResponse struct {
	V   int    `json:"v"`
	Err string `json:"err"`
}

type ConcatReqest struct {
	A string `json:"a"`
	B string `json:"b"`
}

type ConcatResponse struct {
	V   string `json:"v"`
	Err string `json:"err"`
}

// trim
type trimRequest struct {
	s string
}

type trimResponse struct {
	s string
}

// 2. endpoint
// 借助适配器将方法 -> endpoint
func makeSumEndpoint(srv AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SumReqest)
		v, err := srv.Sum(ctx, req.A, req.B) //方法调用
		if err != nil {
			return SumResponse{V: v, Err: err.Error()}, nil
		}
		return SumResponse{V: v}, nil
	}
}

func makeConcatEndpoint(srv AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ConcatReqest)
		v, err := srv.Concat(ctx, req.A, req.B) //方法调用
		if err != nil {
			return ConcatResponse{V: v, Err: err.Error()}, nil
		}
		return ConcatResponse{V: v}, nil
	}
}

// makeTrimEndpoint 客户端 endpoint
// 不是直接提服务，而是请求其他服务
func makeTrimEndpoint(conn *grpc.ClientConn) endpoint.Endpoint {
	return grpctransport.NewClient(
		conn,
		"pb.Trim",          // 服务名
		"TrimSpace",        // 方法名
		encodeTrimRequest,  // 编码
		decodeTrimResponse, // 解码
		pb.TrimResponse{},  // 接收结构
	).Endpoint()
}
