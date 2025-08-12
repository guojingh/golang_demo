package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// 创建一个新的 MCP 服务器实例
	s := server.NewMCPServer(
		"mcp demo",
		"1.0.0",
		server.WithToolCapabilities(false),
	)

	// 配置一个工具
	tool := mcp.NewTool("hello_world",
		mcp.WithDescription("Say hello to someone"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Name of the person go greet"),
		),
	)

	// 新增工具
	s.AddTool(tool, helloHandler)

	// 启动 MCP 服务器
	// if err := server.ServeStdio(s); err != nil {
	// 	fmt.Printf("Server error: %v\n", err)
	// }

	// 创建sse服务器
	sseServer := server.NewSSEServer(s)
	if err := sseServer.Start(":8888"); err != nil {
		fmt.Printf("SSE Server error: %v\n", err)
	}
}

// helloHandler 是处理 hello_world 工具请求的函数
func helloHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, ok := request.GetArguments()["name"].(string)
	if !ok {

		return nil, errors.New("name must be a string")

	}
	return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
}
