package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
)

// 提供客户端双向证书的客户端，一般用于 http 双向认证
func main() {
	pool := x509.NewCertPool()
	caCertPath := "../ca.crt"

	caCrt, err := os.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)
	cliCrt, err := tls.LoadX509KeyPair("../client.crt", "../client.key")
	if err != nil {
		fmt.Println("LoadX509KeyPair err:", err)
		return
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:8081")
	if err != nil {
		fmt.Println("Get err:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
