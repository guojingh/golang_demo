package main

import (
	"fmt"
	"log_transfer/es"
	"log_transfer/kafka"
	"log_transfer/model"
	"strings"

	"github.com/go-ini/ini"
)

// log transfer
// 从 kafka 中消费日志数据，写入 es

func main() {
	//1.加载配置文件
	var cfg = new(model.Config)
	err := ini.MapTo(cfg, "./config/logtransfer.ini")
	if err != nil {
		fmt.Printf("load ini fail, err:%v\n", err)
		panic(err)
	}

	fmt.Println("load config success...")

	//2.连接 ES
	addr := strings.Split(cfg.ESConf.Address, ",")
	err = es.Init(addr, cfg.ESConf.Index, cfg.ESConf.MaxSize, cfg.ESConf.GoNum)
	if err != nil {
		fmt.Printf("init es fail, err:%v\n", err)
		panic(err)
	}
	fmt.Println("init es success...")

	//3.连接 kafka
	addrList := strings.Split(cfg.KafkaConf.Address, ",")
	err = kafka.Init(addrList, cfg.KafkaConf.Topic)
	if err != nil {
		fmt.Printf("init kafka fail, err:%v\n", err)
		panic(err)
	}
	fmt.Println("init kafka success...")

	// 在这儿停顿
	select {}
}
