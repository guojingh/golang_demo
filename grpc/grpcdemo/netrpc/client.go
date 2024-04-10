package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	//建立 http 连接
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9091")
	if err != nil {
		log.Fatal("dialing", err)
	}

	// 同步调用
	args := &Args{10, 20}
	var reply int
	err = client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		log.Fatal("ServiceA.Add error:", err)
	}

	fmt.Printf("ServiceA.Add: %d+%d=%d\n", args.X, args.Y, reply)

	//异步调用
	var reply2 int
	divCall := client.Go("ServiceA.Add", args, &reply2, nil)
	replayCall := <-divCall.Done
	fmt.Println(replayCall.Error)
	fmt.Println(reply2)
}
