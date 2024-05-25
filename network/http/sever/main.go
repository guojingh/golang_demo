package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	host = "127.0.0.1"
	port = ":9000"
)

type Hello struct {
	Hello string
}

func helloHandler(hp http.ResponseWriter, r *http.Request) {
	var param []byte

	hello := &Hello{
		Hello: "Hello Http",
	}

	json.Unmarshal(param, hello)
	fmt.Printf("%v", hello)

	response, _ := json.Marshal(hello)
	//给客户端返回
	_, err := hp.Write(response)
	if err != nil {
		log.Printf("hello 请求出错:%s", err)
	}
}

func main() {
	//注册一个处理函数
	http.HandleFunc("/hello", helloHandler)
	//http.ListenAndServe(port, nil) 启动服务，指定端口
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Printf("服务器异常:%s", err)
	}
	log.Println("服务器关闭")
}
