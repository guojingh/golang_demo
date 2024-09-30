package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

type TraceCode string

func worker(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code")
	}
	log.Printf("%s worker func...", traceCode)
LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*50)
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12512312234")
	log.Printf("%s main 函数", "12512312234")
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over")
}
