package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	Auth struct { // jwt 鉴权配置
		AccessSecret string // jwt 秘钥
		AccessExpire int64  // 有效期，单位：秒
	}

	Mysql struct { // Mysql ??
		DataSource string
	}

	CacheRedis cache.CacheConf
}
