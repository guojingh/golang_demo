package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	pool := x509.NewCertPool()
	caCertPath := "../ca.crt"

	caCrt, err := os.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}

	// 将PEM格式的证书追加到证书池中
	ok := pool.AppendCertsFromPEM(caCrt)
	if !ok {
		fmt.Println("❌ CA 证书解析失败，请检查 ca.crt 格式是否为 PEM")
		return
	}
	fmt.Println("✅ CA 证书加载成功")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: pool},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:8081")
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Get err:", err)
		return
	}
	fmt.Println(string(body))
}
