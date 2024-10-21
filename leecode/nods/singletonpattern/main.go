package main

import (
	"sync"
	"time"
)

type singleton struct {
	current time.Time
}

// 单例模式 (懒汉式，调用的时候去创建)

var instance *singleton
var lock *sync.Mutex

func GetInstance() *singleton {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = &singleton{current: time.Now()}
	}

	return instance
}

var instance2 = &singleton{
	current: time.Now(),
}

// 单例模式（饿汉式）
func GetInstance2() *singleton {
	return instance
}

func main() {

}
