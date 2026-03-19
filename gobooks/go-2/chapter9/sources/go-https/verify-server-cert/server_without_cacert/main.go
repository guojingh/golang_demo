package main

import (
	"fmt"
	"net/http"
)

// 对服务端公钥证书的校验
// 创建一个新的HTTPS web服务，该服务通过CA 新签发的证书
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world\n")
	})

	fmt.Println(http.ListenAndServeTLS("localhost:8081",
		"../server.crt",
		"../server.key", nil))
}
