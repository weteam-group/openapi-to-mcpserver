package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/higress-group/openapi-to-mcpserver/api/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 添加 CORS 中间件
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 健康检查接口
	r.GET("/health", handlers.HealthCheck)

	// OpenAPI 转换接口
	r.POST("/openapi-to-mcp", handlers.ConvertOpenAPI)

	return r
} 