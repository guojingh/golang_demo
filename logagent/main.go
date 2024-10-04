package main

import (
	"fmt"
	"strings"

	"github.com/go-ini/ini"
	"github.com/guojinghu/logagent/common"
	"github.com/guojinghu/logagent/etcd"
	"github.com/guojinghu/logagent/kafka"
	"github.com/guojinghu/logagent/tailfile"
	"github.com/sirupsen/logrus"
)

//日志收集客户端
//类似的开源项目还有 filebeat
//收集指定目录下的日志文件，发送到 kafka 中

//技能包：
// 往Kafka发数据
// 使用tail读日志文件

// 整个logagent的配置结构体
type Config struct {
	KafkaConfig   `ini:"kafka"`
	CollectConfig `ini:"collect"`
	EtcdConfig    `ini:"etcd"`
}

type KafkaConfig struct {
	Address  string `ini:"address"`
	Topic    string `ini:"topic"`
	ChanSize int64  `ini:"chan_size"`
}

type CollectConfig struct {
	LogFilePath string `ini:"logfile_path"`
}

type EtcdConfig struct {
	Address    string `ini:"address"`
	CollectKey string `ini:"collect_key"`
}

func run() {
	for {
		select {}
	}
}

func main() {
	//-1: 获取本机 ip 为后续去 etcd 获取配置文件打下基础
	ip, err := common.GetOutBoundIP()
	if err != nil {
		logrus.Errorf("get ip failed, err:%v", err)
		return
	}
	var configObj = new(Config)
	// 0.读配置文件
	/*	cfg, err := ini.Load("./conf/config.ini")
		if err != nil {
			logrus.Error("load config failed,err:%v", err)
			return
		}
		kafkaAddr := cfg.Section("kafka").Key("address").String()
		fmt.Println(kafkaAddr)*/
	err = ini.MapTo(configObj, "./conf/config.ini")
	if err != nil {
		logrus.Error("load config failed,err:%v", err)
		return
	}
	fmt.Printf("%#v\n", configObj)

	// 1.初始化（做好准备工作）
	err = kafka.Init(strings.Split(configObj.KafkaConfig.Address, ","), configObj.KafkaConfig.ChanSize)
	if err != nil {
		logrus.Errorf("init kafka failed, err:%v", err)
		return
	}
	logrus.Info("init kafka success!")

	// 初始化 etcd 连接
	err = etcd.Init(strings.Split(configObj.EtcdConfig.Address, ","))
	if err != nil {
		logrus.Errorf("init etcd failed, err:%v", err)
		return
	}
	logrus.Info("init etcd success!")
	// 从 etcd 中拉取要收集的日志的配置项
	collectKey := fmt.Sprintf(configObj.EtcdConfig.CollectKey, ip)
	allConf, err := etcd.GetConf(collectKey)
	if err != nil {
		logrus.Errorf("get etcd config failed, err:%v", err)
		return
	}
	fmt.Println(allConf)

	// 派一个小弟去监控 etcd 中 configObj.EtcdConfig.CollectKey 对应值的变化
	go etcd.WatchConf(collectKey)
	// 2.根据配置中的日志路径使用 tailfile 去收集日志，初始化 tailfile
	// 把从etcd中获取的配置项传到 Init
	err = tailfile.Init(allConf)
	if err != nil {
		logrus.Errorf("init tailfile failed, err:%v", err)
		return
	}
	logrus.Info("init tailfile success!")
	run()
}
