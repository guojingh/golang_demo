package main

import (
	"fmt"
	"net/url"
	"path"
	"sync"
	"time"
)

type a struct {
	b *b
}

type b struct {
	value  map[string]interface{}
	locker sync.RWMutex
}

func main() {

	// 输入的时间字符串
	input := "2024-07-15 00:00"

	// 定义时间布局
	layout := "2006-01-02 15:04"

	// 解析时间字符串
	parsedTime, err := time.Parse(layout, input)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}

	// 输出解析后的时间
	fmt.Println("Parsed time:", parsedTime)

	t, _ := GetBasePath("http://www.baidu.com/cncmp/liwienz?name=guojinghu")
	fmt.Println(t)
}

func GetBasePath(targetUrl string) (string, error) {
	myUrl, err := url.Parse(targetUrl)
	if err != nil {
		return "", err
	}

	return path.Base(myUrl.Path), nil
}
