package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func route1() http.Handler {
	r := gin.New()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "你好，我是一号服务")
	})

	return r
}

func route2() http.Handler {
	r := gin.New()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "你好，我是二号服务")
	})

	return r
}

var e errgroup.Group

func main() {
	server1 := &http.Server{
		Addr:         "127.0.0.1:9001",
		Handler:      route1(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	server2 := &http.Server{
		Addr:         "127.0.0.1:9002",
		Handler:      route2(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	e.Go(func() error {
		return server1.ListenAndServe()
	})

	e.Go(func() error {
		return server2.ListenAndServe()
	})

	if err := e.Wait(); err != nil {
		log.Fatal("系统异常退出！")
	}
}
