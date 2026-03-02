package main

import (
	"log"
	"sync"
	"time"
)

// 通过 sync.Once 实现单例模式
type Foo struct{}

var once sync.Once
var instance *Foo

func GetInstance(id int) *Foo {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("goroutine-%d: caught a panic:%s", id, e)
		}
	}()
	log.Printf("goroutine-%d: enter GetInstance\n", id)
	once.Do(func() {
		instance = &Foo{}
		time.Sleep(3 * time.Second)
		log.Printf("goroutine-%d: the addr of instance is %p\n", id, instance)
		panic("panic in once.Do function")
	})

	return instance
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			inst := GetInstance(i)
			log.Printf("goroutine-%d: the addr of instance returned is %p\n", i, inst)
			wg.Done()
		}(i + 1)
	}

	time.Sleep(5 * time.Second)
	inst := GetInstance(0)
	log.Printf("goroutine-%d: the addr of instance returned is %p\n", 0, inst)

	wg.Wait()
	log.Println("all goroutine is exit")
}
