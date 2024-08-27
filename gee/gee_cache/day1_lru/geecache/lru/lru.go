package lru

import "container/list"

// 最近最少使用淘汰策略，对并发访问不安全
type Cache struct {
	maxBytes  int64                         // 最大内存容量
	nBytes    int64                         // 已存大小
	ll        *list.List                    // Go 标准库实现的双向链表
	cache     map[string]*list.Element      // 双向链表中对应节点的指针
	OnEvicted func(key string, value Value) // 某条记录被移除后的回调
}

type entry struct {
	key   string
	value Value
}

// Len() 计算Value占用内存大小
type Value interface {
	Len() int
}

// 创建 Cache
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// 添加缓存数据
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		// 将链表中的节点移动到队尾，这是双向链表作为队列，队首队尾是相对应的，这里约定front为队尾
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nBytes += int64(len(key)) + int64(value.Len())
	}

	// 内存超出最大时进行淘汰
	for c.maxBytes != 0 && c.maxBytes < c.nBytes {
		c.RemoveOldest()
	}
}

// 获取缓存元素
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// 移除最近最少使用的元素
func (c *Cache) RemoveOldest() {
	// 取到队首节点，从链表中删除
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		// 从字典中删除该节点的映射关系
		delete(c.cache, kv.key)
		c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			// 删除之后进行回调的函数
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// 计算缓存元素的内存占用
func (c *Cache) Len() int {
	return c.ll.Len()
}
