package main

import (
	"context"

	"github.com/cncamp/golang/middlewares/kafka/test"
)

func main() {

	//向kafka读数据
	test.ReadByConn(context.Background(), "tcp", test.Address, test.Topic, test.Partition)
}
