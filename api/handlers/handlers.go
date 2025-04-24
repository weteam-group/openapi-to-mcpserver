package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/higress-group/openapi-to-mcpserver/internal/converter"
	"github.com/higress-group/openapi-to-mcpserver/internal/models"
	"github.com/higress-group/openapi-to-mcpserver/internal/parser"
)

type ConvertRequest struct {
	OpenAPISpec string `json:"openapi_spec" binding:"required"`
	Options     struct {
		ServerName     string                 `json:"server_name"`
		ToolNamePrefix string                 `json:"tool_name_prefix"`
		ServerConfig   map[string]interface{} `json:"server_config"`
		Template       string                 `json:"template"`
		Validate       bool                   `json:"validate"`
	} `json:"options"`
	Format string `json:"format" binding:"required,oneof=yaml json"`
}

// HealthCheck 处理健康检查请求
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// ConvertOpenAPI 处理 OpenAPI 转换请求
func ConvertOpenAPI(c *gin.Context) {
	var req ConvertRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 创建解析器
	p := parser.NewParser()
	p.SetValidation(req.Options.Validate)

	// 解析 OpenAPI 规范
	if err := p.ParseContent([]byte(req.OpenAPISpec)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "解析 OpenAPI 规范失败: " + err.Error()})
		return
	}

	// 创建转换器
	convertOptions := models.ConvertOptions{
		ServerName:     req.Options.ServerName,
		ToolNamePrefix: req.Options.ToolNamePrefix,
		ServerConfig:   req.Options.ServerConfig,
		TemplatePath:   req.Options.Template,
	}
	conv := converter.NewConverter(p, convertOptions)

	// 执行转换
	config, err := conv.Convert()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "转换失败: " + err.Error()})
		return
	}

	// 根据请求的格式返回结果
	if req.Format == "json" {
		c.JSON(http.StatusOK, config)
	} else {
		c.YAML(http.StatusOK, config)
	}
} 