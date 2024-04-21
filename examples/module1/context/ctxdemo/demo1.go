package main

import (
	"context"
	"fmt"
)

type UserInfo struct {
	name string
}

// 使用 context 进行传送数据 struct
func main1() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", UserInfo{
		name: "郭靖虎",
	})
	GetUser(ctx)
}

func GetUser(ctx context.Context) {
	// 可以使用断言获取数据
	fmt.Println(ctx.Value("name").(UserInfo).name)
}
