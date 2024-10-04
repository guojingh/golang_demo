package kafka

import (
	"encoding/json"
	"fmt"
	"log_transfer/es"

	"github.com/IBM/sarama"
)

// 初始化 kafka 连接
// 从 kafka 里面取出日志数据

func Init(addr []string, topic string) (err error) {
	consumer, err := sarama.NewConsumer(addr, nil)
	if err != nil {
		fmt.Printf("Failed to start consumer: %s\n", err)
		return
	}

	// 拿到指定 topic 下面的所有分区列表
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Printf("Failed to get list of partitions: %s\n", err)
		return
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
			return err
		}
		//defer pc.AsyncClose()
		go func(partitionConsumer sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				//logDataChan <- msg // 为了将同步流程异步化，所以将取出的日志数据先放到channel中
				var m1 map[string]interface{}
				err := json.Unmarshal(msg.Value, &m1)
				if err != nil {
					fmt.Printf("unmarshal msg fail: %v\n", err)
					continue
				}
				es.PutLogData(m1)
			}
		}(pc)
	}
	return
}
