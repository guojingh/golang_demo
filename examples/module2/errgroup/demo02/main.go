package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

// 并发执行一组文件读取操作，并等待它们全部完成或任何一个读取操作返回错误。
func main() {

	g := new(errgroup.Group)
	urls := []string{
		"http://example.com",
		"http://example.net",
		"http://example.org1223456",
	}

	for _, url := range urls {
		url := url
		g.Go(func() error {
			resp, err := http.Get(url)
			if err != nil {
				return err
			}

			defer resp.Body.Close()
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("one of the requests returned an error:", err)
	}
}
