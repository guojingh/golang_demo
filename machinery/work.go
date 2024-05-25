package main

import (
	"fmt"
	"log"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
)

// SayHello 定义任务
func SayHello(args []string) (string, error) {
	for _, arg := range args {
		fmt.Println(arg)
	}
	return "ok", nil
}

func main() {
	// 将配置文件实例化
	cnf, err := config.NewFromYaml("./config.yaml", false)
	if err != nil {
		log.Println("config.NewFromYaml failed, err:", err)
		return
	}

	// 根据实例化的配置文件创建 server 实例
	server, err := machinery.NewServer(cnf)
	if err != nil {
		log.Println("machinery.NewServer failed, err:", err)
		return
	}

	// 为消费者程序注册任务
	err = server.RegisterTask("SayHello", SayHello)
	if err != nil {
		log.Println("server.RegisterTask failed, err:", err)
		return
	}

	// 创建 worker 实例并绑定任务队列名
	worker := server.NewWorker("sms", 1)
	// 运行 worker 监听逻辑，监听消息队列中的任务
	err = worker.Launch()
	if err != nil {
		log.Println("start worker error", err)
		return
	}
}
