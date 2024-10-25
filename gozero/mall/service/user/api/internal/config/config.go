package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	Mysql struct { // Mysql ??
		DataSource string
	}

	CacheRedis cache.CacheConf
}
