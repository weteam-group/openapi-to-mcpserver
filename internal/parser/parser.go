package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"gopkg.in/yaml.v3"
)

// Parser represents an OpenAPI parser
type Parser struct {
	document *openapi3.T
	validate bool
}

// NewParser creates a new OpenAPI parser
func NewParser() *Parser {
	return &Parser{
		validate: false,
	}
}

// SetValidation sets whether to validate the OpenAPI specification
func (p *Parser) SetValidation(validate bool) {
	p.validate = validate
}

// ParseFile parses an OpenAPI specification file
func (p *Parser) ParseFile(path string) error {
	// Read the file
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	return p.ParseContent(data)
}

// ParseContent parses OpenAPI specification from content
func (p *Parser) ParseContent(content []byte) error {
	var doc openapi3.T

	// 根据内容格式选择解析方式
	if isJSON(content) {
		// JSON 格式
		if err := json.Unmarshal(content, &doc); err != nil {
			return fmt.Errorf("failed to parse OpenAPI specification: %w", err)
		}
	} else {
		// YAML 格式
		if err := yaml.Unmarshal(content, &doc); err != nil {
			return fmt.Errorf("failed to parse OpenAPI specification: %w", err)
		}
	}

	// Validate the document if requested
	if p.validate {
		if err := doc.Validate(nil); err != nil {
			return fmt.Errorf("OpenAPI specification validation failed: %w", err)
		}
	}

	p.document = &doc
	return nil
}

// GetDocument returns the parsed OpenAPI document
func (p *Parser) GetDocument() *openapi3.T {
	return p.document
}

// GetPaths returns the paths from the OpenAPI document
func (p *Parser) GetPaths() map[string]*openapi3.PathItem {
	if p.document == nil {
		return nil
	}
	return p.document.Paths
}

// GetServers returns all servers in the OpenAPI document
func (p *Parser) GetServers() []*openapi3.Server {
	if p.document == nil {
		return nil
	}
	return p.document.Servers
}

// GetInfo returns the info section of the OpenAPI document
func (p *Parser) GetInfo() *openapi3.Info {
	if p.document == nil {
		return nil
	}
	return p.document.Info
}

// isJSON checks if the data is in JSON format
func isJSON(data []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(data, &js) == nil
}

// GetOperationID returns the operation ID for a given path and method
func (p *Parser) GetOperationID(path, method string, operation *openapi3.Operation) string {
	if operation.OperationID != "" {
		return operation.OperationID
	}

	// Generate a default operation ID based on the path and method
	return fmt.Sprintf("%s_%s", method, path)
}
