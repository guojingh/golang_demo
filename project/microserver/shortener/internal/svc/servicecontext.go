package svc

import (
	"github.com/guojinghu/shortenter/internal/config"
	"github.com/guojinghu/shortenter/model"
	"github.com/guojinghu/shortenter/sequence"
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	ShortUrlModel model.ShortUrlMapModel // short_url_map

	// Sequence *sequence.MySQL
	Sequence sequence.Sequence

	ShortUrlBlackList map[string]struct{}

	// bloom filter
	Filter *bloom.Filter
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.ShortUrlDB.DSN)
	m := make(map[string]struct{}, len(c.ShortUrlBlackList))

	// 初始化布隆过滤器
	// 初始化 reidsBitSet
	store := redis.New(c.CacheRedis[0].Host, func(r *redis.Redis) {
		r.Type = redis.NodeType
	})

	// 声明一个bitSet , key="test_key"名且bits是1024位
	filter := bloom.New(store, "bloom_filter", 20*(1<<20))
	// 加载已有的短链接数据
	return &ServiceContext{
		Config:        c,
		ShortUrlModel: model.NewShortUrlMapModel(conn, c.CacheRedis),
		Sequence:      sequence.NewMySQL(c.Sequence.DSN), // sequence
		//Sequence: sequence.NewRedis(c.Redis),
		ShortUrlBlackList: m,
		Filter:            filter,
	}
}

// loadDataToBloomFilter 加载已有的数据到布隆过滤器
func LoadDataToBloomFilter() {

}
