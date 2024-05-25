package cache

import (
	"sync"
	"time"

	"github.com/cncamp/golang/project/cachepj/util"
)

type MemCache struct {
	//最大内存
	maxMemCache int64
	//最大内存字符串表示
	maxMemCacheStr string
	//当前使用内存大小
	currentCacheSize int64
	//键值对存储
	values map[string]*memCacheValue
	//锁
	locker sync.RWMutex
	//定期清楚过期数据
	cleanExpireItemInterval time.Duration
}

type memCacheValue struct {
	//value值
	val interface{}
	//过期时间
	expireTime time.Time
	//过期时间
	expire time.Duration
	//size
	size int64
}

func NewMemCache() *MemCache {
	m := &MemCache{
		values: make(map[string]*memCacheValue),
	}
	go m.cleanExpireItem()
	return m
}

// SetMaxMemory size: 1KB 100KB 1MB 2MB 1GB
func (m *MemCache) SetMaxMemory(size string) bool {
	maxMemCache, maxMemCacheStr := util.ParseSize(size)
	m.maxMemCache = maxMemCache
	m.maxMemCacheStr = maxMemCacheStr
	return false
}

// Set 将 value 写入缓存
func (m *MemCache) Set(key string, val interface{}, expire time.Duration) bool {
	m.locker.Lock()
	defer m.locker.Unlock()
	v := &memCacheValue{
		val:        val,
		expireTime: time.Now().Add(expire),
		size:       util.GetValSize(val),
	}

	m.values[key] = v
	return true
}

// Get 根据 key 获取 value
func (m *MemCache) Get(key string) (interface{}, bool) {
	m.locker.RLock()
	defer m.locker.RUnlock()
	value := m.values[key]
	if value != nil {
		return value.val, true
	}
	return nil, false
}

// Del 删除 key 值
func (m *MemCache) Del(key string) bool {
	return false
}

// Exists 判断 key 是否存在
func (m *MemCache) Exists(key string) bool {
	return false
}

// Flush 清空所有key
func (m *MemCache) Flush() bool {
	return true
}

// Keys 获取缓存中key的数量
func (m *MemCache) Keys() int64 {
	return int64(100)
}

// CleanCacheData 定期清楚缓存数据
func (m *MemCache) cleanExpireItem() {
	//创建定时器
	ticker := time.NewTicker(m.cleanExpireItemInterval)
	//监听 ticker
	for {
		select {
		case <-ticker.C:
			for key, value := range m.values {
				//如果过期时间不为0 并且过期时间已过
				if value.expire != 0 && time.Now().After(value.expireTime) {
					//清楚过期数据
					m.locker.Lock()
					m.Del(key)
					m.locker.Unlock()
				}
			}
		}
	}
}
