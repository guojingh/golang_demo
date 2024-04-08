package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

/*
Reader是由 kafka-go 包提供的另一个概念，对于从单个主题-分区（topic-partition）消费消息这种典型场景，使用它能够简化代码。
Reader 还实现了自动重连和偏移量管理，并支持使用 Context 支持异步取消和超时的 API。
注意： 当进程退出时，必须在 Reader 上调用 Close() 。Kafka服务器需要一个优雅的断开连接来阻止它继续尝试向已连接的客户端发送消息。
如果进程使用 SIGINT (shell 中的 Ctrl-C)或 SIGTERM (如 docker stop 或 kubernetes start)终止，那么下面给出的示例不会调用 Close()。
当同一topic上有新Reader连接时，可能导致延迟(例如，新进程启动或新容器运行)。在这种场景下应使用signal.Notify处理程序在进程关闭时关闭Reader。
*/

func ReadByReader() {
	//1.创建 Reader
	r := kafka.NewReader(kafka.ReaderConfig{
		//2.使用哪Reader建立连接指定 Address Topic Partition
		Brokers:   []string{"192.168.222.134:9092", "192.168.222.134:9093", "192.168.222.134:9094"},
		Topic:     "my-topic",
		Partition: 0,
		MaxBytes:  10e6, //10MB
	})
	err := r.SetOffset(0)
	if err != nil {
		return
	} //设置 offset

	//3.读取数据
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}
	//4.关闭连接
	if err = r.Close(); err != nil {
		log.Fatal("ReadByReader close reader failed:", err)
	}
}

// ReadByGroupID 使用消费者组读取消息
func ReadByGroupID() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"192.168.222.134:9092", "192.168.222.134:9093", "192.168.222.134:9094"},
		GroupID:        "consumer-group-id",
		Topic:          "my-topic",
		MaxBytes:       10e6, //10MB
		CommitInterval: time.Second,
	})

	ctx := context.Background()
	//接收消息
	for {
		//m, err := r.ReadMessage(ctx)  //在非显示提交下直接获取
		m, err := r.FetchMessage(ctx) //在显示提交下使用这种获取
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		//显示提交
		if err := r.CommitMessages(ctx, m); err != nil {
			log.Fatal("failed to commit message:", err)
		}
	}
	//4.关闭连接
	if err := r.Close(); err != nil {
		log.Fatal("ReadByReader close reader failed:", err)
	}
}
