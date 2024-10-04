package kafka

import (
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

var (
	client  sarama.SyncProducer
	msgChan chan *sarama.ProducerMessage
)

// kafka 相关操作
// Init 是初始化全局的 Kafka Client
func Init(address []string, chanSize int64) (err error) {
	// 生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //ACK
	config.Producer.Partitioner = sarama.NewRandomPartitioner //分区
	config.Producer.Return.Successes = true                   //确认

	//连接 Kafka
	client, err = sarama.NewSyncProducer(address, config)
	if err != nil {
		logrus.Error("kafka:producer closed, err:", err)
		return
	}
	// 初始还 MsgChan
	msgChan = make(chan *sarama.ProducerMessage, chanSize)
	// 起一个后台的 goroutine 从 MsgChan 中读数据
	go sendMsg()
	return
}

// 从MsgChan 中读取消息msg,发送给 kafka
func sendMsg() {
	for {
		select {
		case msg := <-msgChan:
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				logrus.Warning("send msg failed, err:", err)
				return
			}
			logrus.Info("send msg to kafka success. pid:%v offset:%v", pid, offset)
		}
	}
}

// 定义一个函数向外暴露 msgChan
func MsgChan(msg *sarama.ProducerMessage) {
	msgChan <- msg
}
