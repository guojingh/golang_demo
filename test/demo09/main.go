package main

import (
	"fmt"
	"sync"
	"time"
)

// 通知并等待多个goroutine退出
func worker(j int) {
	time.Sleep(time.Second * (time.Duration(j)))
}

func spawnGroup(n int) chan struct{} {
	quit := make(chan struct{})
	job := make(chan int)

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			name := fmt.Sprintf("worker-%d:", i)
			for {
				j, ok := <-job
				if !ok {
					println(name, "done")
					return
				}
				worker(j)
			}
		}(i)
	}

	go func() {
		<-quit
		close(job)
		wg.Wait()
		quit <- struct{}{}
	}()

	return quit
}

func main() {
	quit := spawnGroup(5)
	println("spawn a group of workers")

	time.Sleep(5 * time.Second)
	println("notify the worker to exit...")

	quit <- struct{}{}
	timer := time.NewTimer(time.Second * 2)
	defer timer.Stop()
	select {
	case <-quit:
		println("group workers done")
	case <-timer.C:
		println("wait worker exit timeout")

	}
}
