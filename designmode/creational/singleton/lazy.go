package singleton

import "sync"

/*
1.懒汉方式是开源项目使用最多的方式，但它的缺点就是“非并发安全”，在实际使用中需要加锁。
2.为了解决懒汉方式非并发安全的问题，需要对实例进行加锁
3.使用 once.Do 可以确保 lazy 实例全局只被创建一次，once.Do 函数可以确保当同时创还能多个对象时，只被一个动作执行
*/
type Lazy struct{}

var lazy *Lazy
var mu sync.Mutex
var once sync.Once

// GetLazy 非并发安全实现
func GetLazy() *Lazy {
	if lazy == nil {
		lazy = &Lazy{}
	}
	return lazy
}

// GetLazySafe 加锁实现
func GetLazySafe() *Lazy {
	if lazy == nil {
		mu.Lock()
		if lazy == nil {
			lazy = &Lazy{}
		}
		mu.Lock()
	}
	return lazy
}

// GetLazyOnce 更加优雅的一种实现
func GetLazyOnce() *Lazy {
	once.Do(func() {
		lazy = &Lazy{}
	})
	return lazy
}
