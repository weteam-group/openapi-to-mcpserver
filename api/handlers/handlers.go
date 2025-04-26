package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/higress-group/openapi-to-mcpserver/internal/converter"
	"github.com/higress-group/openapi-to-mcpserver/internal/models"
	"github.com/higress-group/openapi-to-mcpserver/internal/parser"
)

type ConvertRequest struct {
	OpenAPISpec string `json:"openapi_spec" binding:"required"`
	Options     struct {
		ServerName       string                 `json:"server_name"`
		ToolNamePrefix   string                 `json:"tool_name_prefix"`
		ServerConfig     map[string]interface{} `json:"server_config"`
		ResponseTemplate string                 `json:"response_template"`
		Validate         bool                   `json:"validate"`
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
		// 处理绑定错误，提供更友好的错误提示
		if strings.Contains(err.Error(), "Key: 'ConvertRequest.OpenAPISpec'") &&
			strings.Contains(err.Error(), "Error:Field validation for 'OpenAPISpec' failed on the 'required' tag") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "必须提供 openapi_spec 参数，且不能为空"})
		} else if strings.Contains(err.Error(), "Key: 'ConvertRequest.Format'") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "format 参数必须为 'yaml' 或 'json'"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "请求格式错误: " + err.Error()})
		}
		return
	}

	// 专门校验 openapi_spec 是否为空
	if req.OpenAPISpec == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "必须提供 openapi_spec 参数，且不能为空"})
		return
	}

	// 创建解析器
	p := parser.NewParser()
	p.SetValidation(req.Options.Validate)

	// 解析 OpenAPI 规范
	if err := p.ParseContent([]byte(req.OpenAPISpec)); err != nil {
		errMsg := "解析 OpenAPI 规范失败"
		if strings.Contains(err.Error(), "unmarshal") {
			errMsg = "OpenAPI规范格式错误，请确保提供的是有效的YAML或JSON格式"
		} else if strings.Contains(err.Error(), "validation") {
			errMsg = "OpenAPI规范验证失败，请确保符合 OpenAPI-3.0 标准"
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg + ": " + err.Error()})
		return
	}

	// 创建转换器
	convertOptions := models.ConvertOptions{
		ServerName:       req.Options.ServerName,
		ToolNamePrefix:   req.Options.ToolNamePrefix,
		ServerConfig:     req.Options.ServerConfig,
		ResponseTemplate: req.Options.ResponseTemplate,
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
