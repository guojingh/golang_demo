package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/bytedance/sonic"
)

type QueryRequestInfo struct {
	Query string `json:"query"`
}

type RequestPayload struct {
	ProductName      string           `json:"product_name"`
	ServiceName      string           `json:"service_name"`
	HandlerName      string           `json:"handler_name"`
	Identifier       string           `json:"Identifier"`
	Version          string           `json:"version"`
	ExpIdList        []interface{}    `json:"exp_id_list"`
	ParamList        []string         `json:"param_list"`
	QueryRequestInfo QueryRequestInfo `json:"query_request_info"`
}

func main() {
	// 从命令行参数获取查询字符串，如果没有提供则使用默认值
	query := ""
	if len(os.Args) > 1 {
		query = os.Args[1]
	}

	// 构建请求payload
	payload := RequestPayload{
		ProductName: "AgentsEcologyAgentsEcologyNew",
		ServiceName: "VSRank",
		HandlerName: "rank_mcp_main",
		Identifier:  "1",
		Version:     "default",
		ExpIdList:   []interface{}{},
		ParamList: []string{
			"params={\"page_size\" : 3}",
			"observe_proto_json_message={\"enabled\":false}",
			"deliverType=",
			"browserType=",
			"source_type=",
			"resourceId=",
			"threshold=0.2",
			"is_debug=0",
			// "id=cSefBjMAvFDgXcJC3imWri,BYpzGTARSG5pjeZvNXpq3u",
			// "id=hG7AeTgvYgwCRGpmNqHKtr",
			// "id=d9JPSY6c6D9nHUMggHJea2",
			// "id=Sguwf4odWKf4ro7rWWzZZ2",
			// "id=kU7e6QXJL6RZ6scWgm4DwZ,jpjckXWJSaVyvCYHyBYvth",
			// "id=j5dNqWPwzyPKoL6DW9BrhF",
			"id=j5dNqWPwzyPKoL6DW9BrhF",
		},
		QueryRequestInfo: QueryRequestInfo{
			Query: query,
		},
	}

	// 将payload转换为JSON
	jsonData, err := sonic.Marshal(payload)
	if err != nil {
		fmt.Printf("错误：JSON序列化失败: %v\n", err)
		return
	}
	fmt.Println("payload", string(jsonData))

	// 创建HTTP请求
	// baseURL := "http://10.221.87.15:2391"
	// baseURL = "http://10.221.62.140:3302"
	baseURL := "http://10.229.133.144:2137"
	baseURL = "http://10.68.72.14:2318"
	baseURL = "http://10.41.146.82:2091"
	url := baseURL + "/baidu.vs.VsRankService/Search"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("错误：创建请求失败: %v\n", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("错误：发送请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("错误：读取响应失败: %v\n", err)
		return
	}

	// 打印响应状态和内容
	fmt.Printf("响应状态: %s\n", resp.Status)
	fmt.Printf("响应内容: %s\n", string(body))
}
