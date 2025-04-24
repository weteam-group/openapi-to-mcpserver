package main

import (
	"github.com/higress-group/openapi-to-mcpserver/api/routes"
)

func main() {
	// 设置路由
	r := routes.SetupRouter()

	// 启动服务器
	r.Run(":8080")
} 