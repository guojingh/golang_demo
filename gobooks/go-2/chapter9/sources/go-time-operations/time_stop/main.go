package main

import (
	"log"
	"sync"
	"time"
)

// Timer提供了Stop方法来将尚未触发的定时器从P中的最小堆中移除，使之失效，这样可以减小最小堆管理和垃圾回收的压力

func consume(c <-chan bool, timer *time.Timer) bool {
	if !timer.Stop() {
		select {
		case <-timer.C:
		default:
		}
		//<-timer.C
	}
	timer.Reset(5 * time.Second)
	select {
	case b := <-c:
		if b == false {
			log.Printf("recv false, continue")
			return true
		}
		log.Printf("recv true, return")
		return false
	case <-timer.C:
		log.Printf("timer expired")
		return true
	}
}

func main() {
	c := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)

	//生产者
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 7)
			c <- false
		}
		time.Sleep(time.Second * 7)
		c <- true
		wg.Done()
	}()

	// 消费者
	go func() {
		timer := time.NewTimer(time.Second * 5)
		for {
			if b := consume(c, timer); !b {
				wg.Done()
				return
			}
		}
	}()

	wg.Wait()

}
