package main

// 使用 gorilla/websocket实现WebSocket的双向通信
import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 允许所以来源WebSocket连接
		return true
	},
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")

	for {
		//读取客户端消息
		messagType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		//处理消息
		fmt.Printf("Received message: %s\n", p)
		//发送消息给客户端
		err = conn.WriteMessage(messagType, p)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnection)

	//	启动 websocket 服务器
	fmt.Println("WebSocket Server listening on: 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
