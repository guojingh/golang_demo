package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// watch
func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"172.16.56.129:2379", "172.16.56.130:2379", "172.16.56.134:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v", err)
		return
	}
	defer cli.Close()

	// watch
	watchCh := cli.Watch(context.Background(), "s4")
	for wresp := range watchCh {
		for _, evt := range wresp.Events {
			fmt.Printf("type:%s key:%s value:%s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
		}
	}
}
