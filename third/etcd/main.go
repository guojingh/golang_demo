package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

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

	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	str := `[{"path":"d:/logs/s4.log","topic":"s4_log"},{"path":"e:/logs/web.log","topic":"web_log"},{"path":"c:/logs/nazha.log","topic":"nazha"}]`
	//str := `[{"path":"d:/logs/s4.log","topic":"s4_log"},{"path":"e:/logs/web.log","topic":"web_log"},{"path":"c:/logs/nazha.log","topic":"nazha"},{"path":"c:/logs/nazha2.log","topic":"nazha2"}]`
	//str := `[{"path":"d:/logs/s4.log","topic":"s4_log"},{"path":"e:/logs/web.log","topic":"web_log"}]`
	_, err = cli.Put(ctx, "collect_log_192.168.3.101_conf", str)
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v", err)
		return
	}
	cancel()

	//get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	gr, err := cli.Get(ctx, "collect_log_conf")
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v", err)
		return
	}

	for _, ev := range gr.Kvs {
		fmt.Printf("key:%s, value:%s\n", ev.Key, ev.Value)
	}
	cancel()
}
