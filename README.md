# OpenAPI to MCP Server

A tool to convert OpenAPI specifications to MCP (Model Context Protocol) server configurations.

## Installation

```bash
go install github.com/higress-group/openapi-to-mcpserver/cmd/openapi-to-mcp@latest
```

## Usage

```bash
openapi-to-mcp --input path/to/openapi.json --output path/to/mcp-config.yaml
```

### Options

- `--input`: Path to the OpenAPI specification file (JSON or YAML) (required)
- `--output`: Path to the output MCP configuration file (YAML) (required)
- `--server-name`: Name of the MCP server (default: "openapi-server")
- `--tool-prefix`: Prefix for tool names (default: "")
- `--format`: Output format (yaml or json) (default: "yaml")
- `--validate`: Validate the OpenAPI specification (default: false)

## Example

```bash
openapi-to-mcp --input petstore.json --output petstore-mcp.yaml --server-name petstore
```

### Converting OpenAPI to Higress REST-to-MCP Configuration

This tool can be used to convert an OpenAPI specification to a Higress REST-to-MCP configuration. Here's a complete example:

1. Start with an OpenAPI specification (petstore.json):

```json
{
  "openapi": "3.0.0",
  "info": {
    "version": "1.0.0",
    "title": "Swagger Petstore",
    "description": "A sample API that uses a petstore as an example to demonstrate features in the OpenAPI 3.0 specification"
  },
  "servers": [
    {
      "url": "http://petstore.swagger.io/v1"
    }
  ],
  "paths": {
    "/pets": {
      "get": {
        "summary": "List all pets",
        "operationId": "listPets",
        "parameters": [
          {
            "name": "limit",
            "in": "query",
            "description": "How many items to return at one time (max 100)",
            "required": false,
            "schema": {
              "type": "integer",
              "format": "int32"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "A paged array of pets",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "pets": {
                      "type": "array",
                      "items": {
                        "type": "object",
                        "properties": {
                          "id": {
                            "type": "integer",
                            "description": "Unique identifier for the pet"
                          },
                          "name": {
                            "type": "string",
                            "description": "Name of the pet"
                          },
                          "tag": {
                            "type": "string",
                            "description": "Tag of the pet"
                          }
                        }
                      }
                    },
                    "nextPage": {
                      "type": "string",
                      "description": "URL to get the next page of pets"
                    }
                  }
                }
              }
            }
          }
        }
      },
      "post": {
        "summary": "Create a pet",
        "operationId": "createPets",
        "requestBody": {
          "description": "Pet to add to the store",
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "required": ["name"],
                "properties": {
                  "name": {
                    "type": "string",
                    "description": "Name of the pet"
                  },
                  "tag": {
                    "type": "string",
                    "description": "Tag of the pet"
                  }
                }
              }
            }
          }
        },
        "responses": {
          "201": {
            "description": "Null response"
          }
        }
      }
    },
    "/pets/{petId}": {
      "get": {
        "summary": "Info for a specific pet",
        "operationId": "showPetById",
        "parameters": [
          {
            "name": "petId",
            "in": "path",
            "required": true,
            "description": "The id of the pet to retrieve",
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Expected response to a valid request",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "id": {
                      "type": "integer",
                      "description": "Unique identifier for the pet"
                    },
                    "name": {
                      "type": "string",
                      "description": "Name of the pet"
                    },
                    "tag": {
                      "type": "string",
                      "description": "Tag of the pet"
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
```

2. Convert it to a Higress REST-to-MCP configuration:

```bash
openapi-to-mcp --input petstore.json --output petstore-mcp.yaml --server-name petstore
```

3. The resulting petstore-mcp.yaml file:

```yaml
server:
  name: petstore
tools:
  - name: showPetById
    description: Info for a specific pet
    args:
      - name: petId
        description: The id of the pet to retrieve
        type: string
        required: true
        position: path
    requestTemplate:
      url: /pets/{petId}
      method: GET
    responseTemplate:
      prependBody: |
        # API Response Information

        Below is the response from an API call. To help you understand the data, I've provided:

        1. A detailed description of all fields in the response structure
        2. The complete API response

        ## Response Structure

        > Content-Type: application/json

        - **id**: Unique identifier for the pet (Type: integer)
        - **name**: Name of the pet (Type: string)
        - **tag**: Tag of the pet (Type: string)

        ## Original Response

  - name: createPets
    description: Create a pet
    args:
      - name: name
        description: Name of the pet
        type: string
        required: true
        position: body
      - name: tag
        description: Tag of the pet
        type: string
        position: body
    requestTemplate:
      url: /pets
      method: POST
      headers:
        - key: Content-Type
          value: application/json
    responseTemplate: {}

  - name: listPets
    description: List all pets
    args:
      - name: limit
        description: How many items to return at one time (max 100)
        type: integer
        position: query
    requestTemplate:
      url: /pets
      method: GET
    responseTemplate:
      prependBody: |
        # API Response Information

        Below is the response from an API call. To help you understand the data, I've provided:

        1. A detailed description of all fields in the response structure
        2. The complete API response

        ## Response Structure

        > Content-Type: application/json

        - **pets**:  (Type: array)
          - **pets[].id**: Unique identifier for the pet (Type: integer)
          - **pets[].name**: Name of the pet (Type: string)
          - **pets[].tag**: Tag of the pet (Type: string)
        - **nextPage**: URL to get the next page of pets (Type: string)

        ## Original Response
```

4. This configuration can be used with Higress by adding it to your Higress gateway configuration.

Note how the tool automatically sets the `position` field for each parameter based on its location in the OpenAPI specification:
- The `petId` parameter is set to `position: path` because it's defined as `in: path` in the OpenAPI spec
- The `limit` parameter is set to `position: query` because it's defined as `in: query` in the OpenAPI spec
- The request body properties (`name` and `tag`) are set to `position: body`

The MCP server will automatically handle these parameters in the correct location when making API requests.

For more information about using this configuration with Higress REST-to-MCP, please refer to the [Higress REST-to-MCP documentation](https://higress.cn/en/ai/mcp-quick-start/#configuring-rest-api-mcp-server).

## Features

- Converts OpenAPI paths to MCP tools
- Supports both JSON and YAML OpenAPI specifications
- Generates MCP configuration with server and tool definitions
- Preserves parameter descriptions and types
- Automatically sets parameter positions based on OpenAPI parameter locations
- Handles path, query, header, cookie, and body parameters
- Generates response templates with field descriptions and improved formatting for LLM understanding
- Optional validation of OpenAPI specifications (disabled by default)
