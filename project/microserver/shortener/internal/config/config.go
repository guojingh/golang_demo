package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	ShortUrlDB struct {
		DSN string
	}

	Sequence struct {
		DSN string
	}

	BaseString string // base62指定字符串

	Redis redis.RedisConf

	ShortUrlBlackList []string
	ShortDoamin       string

	CacheRedis cache.CacheConf
}
