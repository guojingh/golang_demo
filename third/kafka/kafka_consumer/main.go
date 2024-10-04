package main

import (
	"fmt"
	"sync"

	"github.com/IBM/sarama"
)

// kafka 消费者
func main() {
	consumer, err := sarama.NewConsumer([]string{"172.16.56.129:9092", "172.16.56.130:9092", "172.16.56.131:9092"}, nil)
	if err != nil {
		fmt.Printf("Failed to start consumer: %s\n", err)
		return
	}

	partitionList, err := consumer.Partitions("web_log")
	if err != nil {
		fmt.Printf("Failed to get list of partitions: %s\n", err)
		return
	}
	fmt.Println(partitionList)
	var wg sync.WaitGroup
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(partitionConsumer sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d key:%s value:%s\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}
	wg.Wait()
}
