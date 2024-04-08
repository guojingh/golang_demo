package main

import (
	"context"

	"github.com/cncamp/golang/middlewares/kafka/test"
)

func main() {

	//向kafka写数据
	test.WriteByConn(context.Background(), "tcp", test.Address, test.Topic, test.Partition)
}
