# OpenAPI to MCP Server

一个将 OpenAPI 规范的文档转换为 MCP（Model Context Protocol）服务器配置的 Web 服务。

## 获取源码

```bash
git clone https://github.com/higress-group/openapi-to-mcpserver.git
cd openapi-to-mcpserver
```

## 配置文件

项目使用以下配置文件：

- `conf/response_template.md`: 默认的API响应描述模板。可以自定义此文件来修改生成的响应结构说明。

## 安装与运行

### 使用 Docker 运行（推荐）

**下载源码**
```bash
git clone https://github.com/higress-group/openapi-to-mcpserver.git
cd openapi-to-mcpserver
```

**Docker 镜像构建**
```bash
# 使用脚本构建
./build-docker.sh

# 或者手动构建
docker build -t openapi-to-mcpserver:latest .
```

**Docker Compose 运行**

```bash
docker-compose up -d
```

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

### OpenAPI 转换 MCP Yaml

```
POST /openapi-to-mcp
```

请求体：
```json
{
  "openapi_spec": "OpenAPI 3.0 规范内容（YAML 或 JSON 格式，必填）",
  "options": {
    "server_name": "服务器名称（默认：openapi-server）",
    "tool_name_prefix": "工具名前缀（默认：空字符串）",
    "server_config": {},  // 可选，服务器配置
    "response_template": "Markdown格式的响应描述模板（默认：空字符串）",
    "validate": "是否验证 OpenAPI 规范（默认：false）"
  },
  "format": "yaml"  // 或 "json"，必填
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
    "openapi_spec": "您的 OpenAPI 3.0 规范的内容",
    "options": {
      "server_name": "mcp-server-0",
      "tool_name_prefix": "mcp-",
      "validate": true,
      "response_template": "# 自定义API响应\n\n以下是API调用的返回结果解析："
    },
    "format": "yaml"
  }'
```

## 错误处理

API 可能返回以下错误：

| 状态码 | 错误原因 | 描述 |
|-------|---------|------|
| 400 | 请求体格式错误 | 请求体不是有效的 JSON 格式 |
| 400 | `openapi_spec` 缺失或为空 | 必须提供有效的 OpenAPI 规范内容 |
| 400 | `format` 参数错误 | 格式必须为 'yaml' 或 'json' |
| 400 | 解析 OpenAPI 规范失败 | 提供的 OpenAPI 内容不是有效的 YAML 或 JSON 格式 |
| 400 | OpenAPI 规范验证失败 | 当 `validate: true` 时，规范内容不符合 OpenAPI-3.0 标准 |
| 500 | 转换失败 | 服务器内部错误，转换过程中出现问题 |

## 常见问题

**Q: 为什么收到 "OpenAPI 规范格式错误" 的提示？**

A: 请确保提供的 `openapi_spec` 内容是有效的 YAML 或 JSON 格式。检查是否存在语法错误，如缺少闭合引号、缩进不正确等。

**Q: 为什么收到 "OpenAPI 规范验证失败" 的提示？**

A: 当设置 `validate: true` 时，系统会验证您的规范是否符合 OpenAPI 3.0 标准。请确保您的 OpenAPI 文档完全符合规范，包括必要的字段和正确的结构。

**Q: 是否支持 OpenAPI 3.1 或 Swagger 2.0？**

A: 目前系统主要支持 OpenAPI 3.0 规范。对于 Swagger 2.0 或其他版本的文档，建议先使用转换工具将其转换为 OpenAPI 3.0 格式。

**Q: 如何处理复杂的 OpenAPI 规范？**

A: 对于复杂的 OpenAPI 规范，建议：
1. 首先设置 `validate: false` 进行基本转换
2. 如遇问题，检查规范格式是否符合标准
3. 对于大型规范，可以分段处理，或者使用工具预处理简化结构

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
├── conf
│   └── response_template.md  # 默认响应模板
├── test              # 测试用例和示例文件
├── Dockerfile        # Docker 镜像构建文件
├── build-docker.sh   # Docker 镜像构建脚本
└── docker-compose.yaml  # Docker Compose 部署配置
```