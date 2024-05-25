package main

import (
	"log"
	"net/http"

	"github.com/cncamp/golang/gin/Demo2/router"
	"golang.org/x/sync/errgroup"
)

const (
	Addr = "127.0.0.1:8088"
)

var e errgroup.Group

func main() {

	//0.gin 引擎
	r := router.SetUpRouter()
	//1.穿建web服务
	server := &http.Server{
		Addr:    Addr,
		Handler: r,
	}

	e.Go(func() error {
		return server.ListenAndServe()
	})

	err := e.Wait()
	if err != nil {
		log.Fatal("系统关闭")
	}
}
