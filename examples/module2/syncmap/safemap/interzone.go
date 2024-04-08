package safemap

import (
	"sync"
)

/*
1.通过 RWMutex 实现的并发安全 map 性能上不太行，因此人们又创造除了 通过分区的方式实现并发 map
2.下面即为具体实现 demo
3.go社区经常使用的 分区加锁算法 concurrent map git：
*/

var SHARD_COUNT = 32

// 分成 SHARD_COUNT 个分区的 map
type ConcurrentMap []*ConcurrentMapShared

// 通过 RWMutex 保护线程的分片，包含一个 map
type ConcurrentMapShared struct {
	items map[string]interface{}
	sync.RWMutex
}

// 创建并发 map
func New() ConcurrentMap {
	m := make(ConcurrentMap, SHARD_COUNT)
	for i := 0; i < SHARD_COUNT; i++ {
		m[i] = &ConcurrentMapShared{items: make(map[string]interface{})}
	}
	return m
}

// 根据 key 计算分区索引
func (m ConcurrentMap) GetShared(key string) *ConcurrentMapShared {
	return m[uint(fnv32(key))%uint(SHARD_COUNT)]
}

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	keyLength := len(key)
	for i := 0; i < keyLength; i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

// 向 map 中添加 k-v
func (m ConcurrentMap) Set(key string, value interface{}) {
	//根据 key 计算出对应的分片
	shard := m.GetShared(key)
	//对这个分区加锁，执行业务操作
	shard.Lock()
	//从这个分区读取 key 值
	shard.items[key] = value
	shard.Unlock()
}

// 向 map 中获取 v
func (m ConcurrentMap) Get(key string) (interface{}, bool) {
	// 根据 key 计算出对应的分片
	shard := m.GetShared(key)
	//读操作加读锁
	shard.RLock()
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}
