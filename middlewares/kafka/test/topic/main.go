package main

import "github.com/cncamp/golang/middlewares/kafka/test"

func main() {

	//创建Topic
	//test.CreateTopicByConn("my-topic-02", test.Address)
	//查询所有Topic
	test.GetTopicList(test.Address)
}
