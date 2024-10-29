package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	UserRPC zrpc.RpcClientConf // 连接其它微服务的RPC客户端

	Mysql struct { // Mysql ??
		DataSource string
	}

	CacheRedis cache.CacheConf
}
