FROM golang:1.21-alpine AS builder

WORKDIR /app

# 复制项目文件
COPY . .

# 安装依赖并构建
RUN go mod download
RUN go build -o openapi-to-mcp-server api/main.go

# 使用轻量级的alpine镜像作为运行环境
FROM alpine:latest

WORKDIR /app

# 从构建阶段复制构建好的二进制文件
COPY --from=builder /app/openapi-to-mcp-server .
# 复制配置文件目录
COPY --from=builder /app/conf ./conf

# 暴露服务端口
EXPOSE 8080

# 运行应用
CMD ["./openapi-to-mcp-server"] 