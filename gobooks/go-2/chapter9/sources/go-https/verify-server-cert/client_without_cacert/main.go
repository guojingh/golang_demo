package main

import (
	"fmt"
	"io"
	"net/http"
)

// 通过普通的客户端直接访问
func main() {
	resp, err := http.Get("https://localhost:8081")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
