### 实现一个内存缓存系统
- 支持设定过期时间，精确到秒
- 支持设定最大内存，当内存超出时做合适处理
- 支持并发安全
- 按照一下接口要求实现
```go
//接口示例
type Cache interface {
	//size: 1KB 100KB 1MB 2MB 1GB
	SetMaxMemory(size string) bool
	//将 value 写入缓存
	Set(key string, val interface{}, expire time.Duration) bool
	//根据 key 获取 value
	Get(key string) (interface{}, bool)
	//删除 key 值
	Del(key string) bool
	//判断 key 是否存在
	Exists(key string) bool
	//清空所有key
	Flush() bool
	//获取缓存中key的数量
	Keys() int64
}

//使用示例
cache := NewMemCache()
cache.SetMaxMemory("100MB")
cache.Set("int", 1)
cache.Set("bool", false)
cache.Set("data", map[string]interface{}{"a":1})
cache.Get("int")
cache.Del("int")
cache.Flush()
cache.Keys()
```

### 分析
- 编写单元测试
- 支持设定过期时间，精确到秒
- 支持设定最大内存，当内存超出时做合适处理
- 支持并发安全
- 按照一下接口要求实现