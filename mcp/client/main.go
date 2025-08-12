package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/mark3labs/mcp-go/mcp"
)

func main() {
	// 创建sse客户端
	sseTransport, err := transport.NewSSE("http://127.0.0.1:8888" + "/sse")
	if err != nil {
		log.Fatalf("failed to created SSE transport: %v", err)
	}

	ctx := context.Background()
	if err := sseTransport.Start(ctx); err != nil {
		log.Fatalf("failed to start SSE transport: %v", err)
	}

	// 创建一个新的客户端示例
	c := client.NewClient(sseTransport)
	defer c.Close()

	// 初始化连接 (客户端调用工具之前一定要进行初始化)
	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "MCP-Go Simple Client Example",
		Version: "1.0.0",
	}

	initRequest.Params.Capabilities = mcp.ClientCapabilities{}

	serverInfo, err := c.Initialize(ctx, initRequest)
	if err != nil {
		log.Fatalf("failed to initialize: %v", err)
	}

	// 显示服务器信息
	fmt.Printf("connected to server: %s(version %s)\n", serverInfo.ServerInfo.Name, serverInfo.ServerInfo.Version)
	fmt.Printf("Server capabilities: %+v\n", serverInfo.Capabilities)

	// Test Ping
	if err := c.Ping(ctx); err != nil {
		log.Fatalf("ping failed: %v", err)
	}

	// 列出服务器支持的工具
	if serverInfo.Capabilities.Tools != nil {
		fmt.Println("fetching avaliable tools...")
		toolsRequest := mcp.ListToolsRequest{}
		toolsResult, err := c.ListTools(ctx, toolsRequest)
		if err != nil {
			log.Printf("failed to list tools: %v", err)
		} else {
			fmt.Printf("server has %d tools available\n", len(toolsResult.Tools))
			for i, tool := range toolsResult.Tools {
				fmt.Printf(" %d. %s - %s\n", i+1, tool.Name, tool.Description)
			}
		}
	}

	// 测试调用工具
	callToolRequest := mcp.CallToolRequest{
		Params: struct {
			Name      string    `json:"name"`
			Arguments any       `json:"arguments,omitempty"`
			Meta      *mcp.Meta `json:"_meta,omitempty"`
		}{
			Name: "hello_world",
			Arguments: map[string]interface{}{
				"name": "World",
			},
		},
	}

	result, err := c.CallTool(ctx, callToolRequest)
	if err != nil {
		log.Fatalf("调用 hello_world 调用工具失败：%v", err)
	}

	// 打印返回值
	fmt.Printf("工具返回结果： %+v\n", result.Content)

}
