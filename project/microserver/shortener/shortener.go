package main

import (
	"flag"
	"fmt"

	"github.com/guojinghu/shortenter/internal/config"
	"github.com/guojinghu/shortenter/internal/handler"
	"github.com/guojinghu/shortenter/internal/svc"
	"github.com/guojinghu/shortenter/pkg/base62"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/shortener-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	fmt.Printf("load conf:%#v\n", c)

	// base62 模块初始化
	base62.MustInit(c.BaseString)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
