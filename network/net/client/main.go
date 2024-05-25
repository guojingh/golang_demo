package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// golang关于网络的都在 net 包里面
func main() {

	fmt.Println("客户端启动...")
	//net.Dial 建立网络连接，指定协议和网络地址
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("连接失败, ", err)
	}

	fmt.Printf("连接成功 conn=%v, 服务器=%v\n", conn, conn.RemoteAddr().String())

	//试图循环读取标准输入流里面的数据失败，后面再研究吧
	/*	for {
			reader := bufio.NewReader(os.Stdin)

			if _, err := reader.ReadString('e'); err != nil {
				break
			}

			str, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("接收标准输入失败")
			}

			n, err := conn.Write([]byte(str))
			if err != nil {
				fmt.Println("连接失败")
			}

			fmt.Printf("发送成功%d个数据", n)
		}
	*/
	//利用标准输入 os.Stdin 建立有缓存的 Reader
	reader := bufio.NewReader(os.Stdin)
	//读取数据以回车为结束标志
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("接收标准输入失败")
	}

	//将数据以字节数据的方式写入连接
	n, err := conn.Write([]byte(str))
	if err != nil {
		fmt.Println("连接失败")
	}

	fmt.Printf("发送成功%d个数据", n)

	fmt.Println("发送数据工作完成")
}
