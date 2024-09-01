package geecache

import (
	"geecache/lru"
	"sync"
)

type cache struct {
	mu         *sync.Mutex // 互斥锁
	lru        *lru.Cache  // 淘汰算法结点实现
	cacheBytes int64       //缓存区大小
}

func (c *cache) Add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// 如果 lru 为空，则进行创建
	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}
	// 添加元素
	c.lru.Add(key, value)
}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}

	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}
	return
}
