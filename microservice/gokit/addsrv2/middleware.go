package main

import (
	"context"
	"errors"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"golang.org/x/time/rate"
)

// 日志中间件
func loggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "calling endpoint")
			start := time.Now()
			defer logger.Log("msg", "called endpoint", "cast", time.Since(start))
			return next(ctx, request)
		}
	}
}

var ErrRateLimit = errors.New("request rate limit")

// 限流中间件
// "golang.org/x/time/rate"
func rateMiddleware(limit *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			// 限流逻辑
			if limit.Allow() {
				return next(ctx, request)
			} else {
				return nil, errors.New("request rate limit")
			}
		}
	}
}
