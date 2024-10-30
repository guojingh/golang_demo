package sequence

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const sequenceKey = "sequence_key"

// 基于 redis 实现一个取号器

type Redis struct {
	// redis 连接
	redis redis.Redis
}

func NewRedis(redisConf redis.RedisConf) Sequence {
	rds := redis.MustNewRedis(redisConf)
	return &Redis{
		redis: *rds,
	}
}

func (r *Redis) Next() (resp uint64, err error) {
	// 使用redis实现发号器的思路
	// incr(原子操作)
	rest, err := r.redis.Incr(sequenceKey)
	if err != nil {
		logx.Errorw("redis.Incr(sequenceKey) failed", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}

	return uint64(rest), nil
}
