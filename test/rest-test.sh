curl --location --request POST 'http://127.0.0.1:8080/openapi-to-mcp' \
--header 'Content-Type: application/json' \
--data-raw '{
    "openapi_spec": "{\"openapi\":\"3.1.0\",\"info\":{\"title\":\"openapi-to-mcp\",\"description\":\"\",\"version\":\"1.0.0\"},\"tags\":[],\"paths\":{\"/health\":{\"get\":{\"summary\":\"健康检查\",\"deprecated\":false,\"description\":\"检查服务是否正常运行\",\"operationId\":\"checkHealth\",\"tags\":[],\"parameters\":[],\"responses\":{\"200\":{\"description\":\"服务正常运行\",\"content\":{\"application/json\":{\"schema\":{\"type\":\"object\",\"properties\":{\"status\":{\"type\":\"string\",\"examples\":[\"ok\"]}}}}},\"headers\":{}}},\"security\":[]}}},\"components\":{\"schemas\":{},\"securitySchemes\":{}},\"servers\":[{\"url\":\"http://127.0.0.1:8080\",\"description\":\"本机环境\"}]}",
    "options": {
        "server_name": "health",
        "tool_name_prefix": "mcp",
        "template": "",
        "validate": true
    },
    "format": "yaml"
}'