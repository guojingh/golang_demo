package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"
)

// 一些安全要求十分严格的场景，也可以对客户的对公钥证书进行验证
func main() {
	pool := x509.NewCertPool()
	caCertPath := "../ca.crt"

	caCrt, err := os.ReadFile(caCertPath)
	if err != nil {
		log.Print("read crt file err:", err)
		return
	}
	// 添加ca证书验证客户端证书
	pool.AppendCertsFromPEM(caCrt)

	s := &http.Server{
		Addr: "localhost:8081",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello World!\n")
		}),
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	fmt.Println(s.ListenAndServeTLS("../server.crt", "../server.key"))
}
