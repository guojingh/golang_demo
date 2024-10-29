package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf

	Mysql struct { // Mysql数据库配置
		DataSource string
	}

	// redis
	CacheRedis cache.CacheConf

	// 引入Consul
	Consul consul.Conf
}
