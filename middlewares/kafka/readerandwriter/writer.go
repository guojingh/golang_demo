package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

/*
向Kafka发送消息，除了使用基于Conn的低级API，kafka-go包还提供了更高级别的 Writer 类型。大多数情况下使用Writer即可满足条件，它支持以下特性。

1.对错误进行自动重试和重新连接。
2.在可用分区之间可配置的消息分布。
3.向Kafka同步或异步写入消息。
4.使用Context的异步取消。
5.关闭时清除挂起的消息以支持正常关闭。
6.在发布消息之前自动创建不存在的topic。
*/

func WriteByWriter() {
	//1.创建writer
	w := &kafka.Writer{
		Addr:                   kafka.TCP("192.168.222.134:9092", "192.168.222.134:9093", "192.168.222.134:9094"),
		Topic:                  "my-topic-test",
		Balancer:               &kafka.LeastBytes{}, //指定分区的 Balancer模式为最小字节分布
		RequiredAcks:           kafka.RequireAll,    //ack模式
		Async:                  true,                //异步
		AllowAutoTopicCreation: true,                //允许自动创建Topic
		Logger:                 kafka.LoggerFunc(zap.NewExample().Sugar().Infof),
	}

	message := []kafka.Message{
		{
			Key:   []byte("key-A"),
			Value: []byte("Hello World!"),
		},
		{
			Key:   []byte("key-B"),
			Value: []byte("1111"),
		},
		{
			Key:   []byte("key-C"),
			Value: []byte("2222"),
		},
	}

	var err error
	const retries = 3
	//重试3次
	for i := 0; i < retries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		//2.发送消息
		err = w.WriteMessages(ctx, message...)
		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(time.Millisecond * 250)
			continue
		}

		if err != nil {
			log.Fatalf("unexpected error %v", err)
		}
		break
	}

	//3.关闭连接
	if err = w.Close(); err != nil {
		log.Fatal("WriteByWriter close failed:", err)
	}
}

// 自定义一个Logger
func logf(msg string, a ...interface{}) {
	fmt.Printf(msg, a...)
	fmt.Println()
}

// WriteByWriterTopics 写入多个Topic
func WriteByWriterTopics() {
	//1.创建 Writer
	w := &kafka.Writer{
		Addr: kafka.TCP("192.168.222.134:9092", "192.168.222.134:9093", "192.168.222.134:9094"),
		//这里如果不指定 Topic ，那么每个消息都要指定 Topic
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true, //允许自动创建Topic
		Logger:                 kafka.LoggerFunc(logf),
	}

	err := w.WriteMessages(context.Background(),
		// 注意: 每条消息都需要指定一个 Topic, 否则就会报错
		kafka.Message{
			Topic: "topic-A",
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Topic: "topic-B",
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Topic: "topic-C",
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
