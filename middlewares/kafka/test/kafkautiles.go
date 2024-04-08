package test

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	Topic     = "my-topic"
	Partition = 0
	Address   = "192.168.222.134:9092"
)

/*
Conn 类型是 kafka-go 包的核心。它代表与 Kafka broker之间的连接。基于它实现了一套与Kafka交互的低级别 API。
*/

// WriteByConn Conn 生产发送消息
func WriteByConn(ctx context.Context, network string, address string, topic string, partition int) {

	//连接到kafka集群的leader节点
	conn, err := kafka.DialLeader(ctx, network, address, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	//设置发送消息超时时间
	if err = conn.SetWriteDeadline(time.Now().Add(10 * time.Second)); err != nil {
		return
	}

	//发送消息
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write message:", err)
	}

	//关闭连接
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close write:", err)
	}
}

// ReadByConn Conn 消费消息
func ReadByConn(ctx context.Context, network string, address string, topic string, partition int) {
	//连接到指定的 topic和 partition，连接到指定的leader节点
	conn, err := kafka.DialLeader(ctx, network, address, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	//设置发送消息超时时间
	if err = conn.SetReadDeadline(time.Now().Add(20 * time.Second)); err != nil {
		return
	}
	//读取一批消息，得到的batch是一系列消息的迭代器
	batch := conn.ReadBatch(10e3, 1e6) //10e3表示10乘以10的3次方，即10000；1e6表示1乘以10的6次方，即1000000

	//遍历读取消息
	//b := make([]byte, 10e3)
	for {
		//如果传入的Buffer过小，就会返回 io.ErrShortBuffer错误
		//n, err := batch.Read(b)
		//也可以使用这种方法获取信息
		msg, err := batch.ReadMessage()
		if err != nil {
			break
		}
		//fmt.Println(string(b[:n]))
		fmt.Println(string(msg.Value))
	}

	//关闭batch
	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	//关闭连接
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close Read:", err)
	}
}

// CreateTopicByConn 通过 Conn 创建Topic
func CreateTopicByConn(topic string, address string) {
	//连接到任意kafka节点
	conn, err := kafka.Dial("tcp", address)
	if err != nil {
		panic(err.Error())
	}

	//闭包关闭连接
	defer func(conn *kafka.Conn) {
		err = conn.Close()
		if err != nil {
			log.Fatal("CreateTopicByConn close conn failed:", err)
		}
	}(conn)

	//获取当前控制节点信息
	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}

	var controllerConn *kafka.Conn
	//连接到leader节点
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}

	//闭包关闭连接
	defer func(controllerConn *kafka.Conn) {
		err = controllerConn.Close()
		if err != nil {
			log.Fatal("CreateTopicByConn close controllerConn failed:", err)
		}
	}(controllerConn)

	topicConfigs := []kafka.TopicConfig{
		{Topic: topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}
	//创建topic
	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
}

// GetTopicList 获取topic列表
func GetTopicList(address string) {
	conn, err := kafka.Dial("tcp", address)
	if err != nil {
		panic(err.Error())
	}

	defer func(conn *kafka.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("GetTopicList close conn failed:", err)
		}
	}(conn)

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}
	//遍历所有分区获取 Topic
	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k := range m {
		fmt.Println(k)
	}
}
