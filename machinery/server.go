package main

import (
	"log"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
)

func NewSumTaskSignature(val []string) *tasks.Signature {
	// 新建任务签名，执行任务名Name，指定传递给任务的参数Args
	signature := &tasks.Signature{
		Name: "SayHello",
		Args: []tasks.Arg{
			{
				Type:  "[]string",
				Value: val,
			},
		},
	}
	return signature
}

func main() {
	// 加载配置文件
	cnf, err := config.NewFromYaml("./config.yaml", false)
	if err != nil {
		log.Println("config failed", err)
		return
	}

	// 新建 server 实例
	server, err := machinery.NewServer(cnf)
	if err != nil {
		log.Println("start server failed", err)
		return
	}

	// 下发任务给消费者
	signature := NewSumTaskSignature([]string{"1", "2", "3", "4", "5"})
	asyncResult, err := server.SendTask(signature)
	if err != nil {
		log.Fatal(err)
	}

	// 每秒获取一次消息队列中的结果
	res, err := asyncResult.Get(1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("get res is %v\n", tasks.HumanReadableResults(res))
}
