package main

import (
	"fmt"
	"time"
)

// time包提供了多种创建timer的方式
func create_timer_by_afterfunc() {
	_ = time.AfterFunc(1*time.Second, func() {
		fmt.Println("timer created by afterfunc fired")
	})
}

func create_timer_by_newtimer() {
	timer := time.NewTimer(2 * time.Second)
	select {
	case <-timer.C:
		fmt.Println("timer created by newtimer fired!")
	}
}

func create_timer_by_after() {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("timer created by after fired!")
	}
}

func create_ticker_by_newticker() {
	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("ticker create by newticker fired!")
		}
	}

}

func main() {
	//create_timer_by_afterfunc()
	create_timer_by_newtimer()
	//create_timer_by_after()
	//create_ticker_by_newticker()
}
