# OpenAPI to MCP Server

一个将 OpenAPI 规范的文档转换为 MCP（Model Context Protocol）服务器配置的 Web 服务。

## 安装

```bash
git clone https://github.com/higress-group/openapi-to-mcpserver.git
cd openapi-to-mcpserver
go mod tidy
```

## 运行

```bash
go run api/main.go
```

服务将在 `:8080` 端口启动。

### 健康检查

```
GET /health
```

响应：
```json
{
  "status": "ok"
}
```

### OpenAPI 转换

```
POST /openapi-to-mcp
```

请求体：
```json
{
  "openapi_spec": "OpenAPI 规范内容（YAML 或 JSON 格式）",
  "options": {
    "server_name": "服务器名称（默认：openapi-server）",
    "tool_name_prefix": "工具名前缀（默认：空字符串）",
    "server_config": {},  // 可选，服务器配置
    "template": "模板文件路径（默认：空字符串）",
    "validate": "是否验证 OpenAPI 规范（默认：false）"
  },
  "format": "yaml"  // 或 "json"
}
```

### 服务器配置（可选）

`server_config` 是一个可选的配置项，用于自定义服务器的行为。如果未提供，将使用默认配置。

可用的配置项：

```json
{
  "host": "服务器主机地址（默认：localhost）",
  "port": "服务器端口号（默认：8080）",
  "base_path": "API 基础路径（默认：/）",
  "schemes": ["http", "https"],  // 支持的协议
  "timeout": "请求超时时间（单位：秒，默认：30）",
  "max_connections": "最大连接数（默认：100）",
  "cors": {
    "enabled": "是否启用 CORS（默认：false）",
    "allowed_origins": ["*"],  // 允许的源
    "allowed_methods": ["GET", "POST", "PUT", "DELETE"],  // 允许的方法
    "allowed_headers": ["Content-Type", "Authorization"]  // 允许的请求头
  },
  "auth": {
    "type": "认证类型（none/basic/jwt，默认：none）",
    "config": {}  // 认证配置，根据认证类型不同而不同
  },
  "logging": {
    "level": "日志级别（debug/info/warn/error，默认：info）",
    "format": "日志格式（text/json，默认：text）"
  }
}
```

响应：
- 成功：返回转换后的 MCP 配置（YAML 或 JSON 格式）
- 失败：返回错误信息

## 示例

使用 curl 调用 API：

```bash
curl -X POST http://localhost:8080/openapi-to-mcp \
  -H "Content-Type: application/json" \
  -d '{
    "openapi_spec": "您的 OpenAPI 规范内容",
    "options": {
      "server_name": "mcp-server-0",
      "tool_name_prefix": "mcp-",
      "validate": true
    },
    "format": "yaml"
  }'
```

## 开发

### 项目结构

```
.
├── api
│   ├── handlers      # HTTP 请求处理器
│   ├── routes        # 路由配置
│   └── main.go       # 服务入口
├── internal
│   ├── converter     # OpenAPI 到 MCP 的转换逻辑
│   ├── models        # 数据模型定义
│   └── parser        # OpenAPI 解析器
└── test              # 测试用例和示例文件
```

### 构建

```bash
go build -o openapi-to-mcp-server api/main.go
```
