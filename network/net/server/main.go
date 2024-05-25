package main

import (
	"fmt"
	"io"
	"net"
)

// 用来启动协程，不同客户端的连接使用不同的 goroutine 进行处理
func process(conn net.Conn) {
	defer conn.Close()

	fmt.Println("hello")
	bytes := make([]byte, 1024)
	for {
		//读取客户端传过来的数据
		n, err := conn.Read(bytes)
		if err != nil && err == io.EOF {
			return
		}
		fmt.Println(string(bytes[0:n]))
	}
}

func main() {
	fmt.Println("服务器端启动...")
	//服务端启动调用 net.listen() 进行监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("服务器启动失败，", err)
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("服务器关闭失败：", err)
		}
	}(listener)

	for {
		//等待服务端建立连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("等待客户端连接失败：", err)
		} else {
			fmt.Printf("等待连接成功，conn=%v, 客户端=%v\n", conn, conn.RemoteAddr().String())
		}

		//启动协程处理客户端传过来的数据
		go process(conn)
	}
}
