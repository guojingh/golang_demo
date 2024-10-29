package main

import (
	"flag"
	"fmt"

	"api/internal/config"
	"api/internal/handler"
	"api/internal/logic"
	"api/internal/middleware"
	"api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	fmt.Printf("--> conf:%#v\n", c)

	// 初始化一个雪花算法节点
	if err := logic.InitSonyflakeNode(1); err != nil {
		fmt.Printf("init validator Trans failed,err:%v\n", err)
		return
	}

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 应用全局中间件
	server.Use(middleware.CopyResp)
	// 中间件中调用其他服务
	server.Use(middleware.MiddlewareWithAnother(true))

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
