package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func main() {

	httpClient := http.Client{
		Timeout: time.Second * 3,
	}

	for {
		go func() {
			_, err := httpClient.Get("https://www.xxx.com/")
			if err != nil {
				fmt.Printf("http.Get err: %v\n", err)
			}
			// do something...
		}()

		time.Sleep(time.Second * 1)
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}

	func serve(addr string, handler http.Handler, stop <-chan struct{}) error {}
}
