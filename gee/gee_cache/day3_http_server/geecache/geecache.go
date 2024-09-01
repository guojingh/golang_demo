package geecache

import (
	"fmt"
	"log"
	"sync"
)

// 区分不同的缓存组，即设置命名空间
type Group struct {
	name      string
	getter    Getter
	mainCache cache // 缓存颗粒
}

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

var (
	// 读写锁，因为最上层都的情况比较对
	mu sync.RWMutex
	// 存放全局的 groups
	groups = make(map[string]*Group)
)

func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	// getter 不存在直接抛出异常
	if getter == nil {
		panic("getter is nil")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cacheBytes},
	}
	groups[name] = g
	return g
}

func GetGroup(name string) *Group {
	mu.RLocker()
	g := groups[name]
	mu.RUnlock()
	return g
}

func (g *Group) Get(key string) (ByteView, error) {
	// 没传key 直接返回异常
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}

	// 从缓存中获取对应key的value
	if v, ok := g.mainCache.get(key); ok {
		log.Println("[GeeCache] hit")
		return v, nil
	}

	// 缓存中没有，则从数据源进行加载
	return g.load(key)
}

func (g *Group) load(key string) (value ByteView, err error) {
	// 从本地数据源进行加载
	return g.getLocally(key)
}

func (g *Group) getLocally(key string) (ByteView, error) {
	// 调用 group 从数据源查找数据的方法
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}

	value := ByteView{
		// 对数据源查到的数据拷贝一份进行赋值
		b: cloneBytes(bytes),
	}
	g.populateCache(key, value)
	return value, nil
}

func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.Add(key, value)
}
