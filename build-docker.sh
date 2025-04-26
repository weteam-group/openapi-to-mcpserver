#!/bin/bash

# 设置镜像名称和标签
IMAGE_NAME="openapi-to-mcpserver"
IMAGE_TAG=${1:-"latest"}  # 使用第一个参数作为标签，默认为"latest"

echo "开始构建 $IMAGE_NAME:$IMAGE_TAG 镜像..."

# 构建Docker镜像
docker build -t $IMAGE_NAME:$IMAGE_TAG .

# 检查构建结果
if [ $? -eq 0 ]; then
    echo "✅ 构建成功：$IMAGE_NAME:$IMAGE_TAG"
    
    # 显示构建的镜像
    echo "镜像详情："
    docker images | grep $IMAGE_NAME
else
    echo "❌ 构建失败"
    exit 1
fi 