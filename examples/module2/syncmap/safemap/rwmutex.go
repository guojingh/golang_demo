package safemap

import (
	"fmt"
	"sync"
)

// 以 map[int]int 为例，借助 rwmutex 实现并发安全
type RWMap struct {
	sync.RWMutex
	m map[int]int
}

// 新建一个RWMap
func NewRWMap(n int) *RWMap {
	return &RWMap{
		m: make(map[int]int, n),
	}
}

// 从map中读取一个值
func (m *RWMap) Get(k int) (int, bool) {
	m.RLock()
	defer m.RUnlock()
	//在锁的保护下，从map中获取
	v, existed := m.m[k]
	return v, existed
}

// 设置一个键值对
func (m *RWMap) Set(k int, v int) {
	m.Lock()
	defer m.Unlock()
	m.m[k] = v
}

// 删除一个键
func (m *RWMap) Delete(k int) {
	m.Lock()
	defer m.Unlock()
	delete(m.m, k)
}

// 获取map的长度
func (m *RWMap) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.m)
}

// 遍历期间一直有读锁
func (m *RWMap) Each(f func(k, v int) bool) {
	m.RLock()
	defer m.RUnlock()

	for k, v := range m.m {
		if !f(k, v) {
			fmt.Printf("%v=%v\n", k, v)
			return
		}
	}
}
