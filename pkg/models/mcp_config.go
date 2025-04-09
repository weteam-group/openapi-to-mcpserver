package models

// MCPConfig represents the top-level MCP server configuration
type MCPConfig struct {
	Server ServerConfig `yaml:"server"`
	Tools  []Tool       `yaml:"tools,omitempty"`
}

// ServerConfig represents the MCP server configuration
type ServerConfig struct {
	Name       string                 `yaml:"name"`
	Config     map[string]interface{} `yaml:"config,omitempty"`
	AllowTools []string               `yaml:"allowTools,omitempty"`
}

// Tool represents an MCP tool configuration
type Tool struct {
	Name             string           `yaml:"name"`
	Description      string           `yaml:"description"`
	Args             []Arg            `yaml:"args"`
	RequestTemplate  RequestTemplate  `yaml:"requestTemplate"`
	ResponseTemplate ResponseTemplate `yaml:"responseTemplate"`
}

// Arg represents an MCP tool argument
type Arg struct {
	Name        string                 `yaml:"name"`
	Description string                 `yaml:"description"`
	Type        string                 `yaml:"type,omitempty"`
	Required    bool                   `yaml:"required,omitempty"`
	Default     interface{}            `yaml:"default,omitempty"`
	Enum        []interface{}          `yaml:"enum,omitempty"`
	Items       map[string]interface{} `yaml:"items,omitempty"`
	Properties  map[string]interface{} `yaml:"properties,omitempty"`
	Position    string                 `yaml:"position,omitempty"`
}

// RequestTemplate represents the MCP request template
type RequestTemplate struct {
	URL            string   `yaml:"url"`
	Method         string   `yaml:"method"`
	Headers        []Header `yaml:"headers,omitempty"`
	Body           string   `yaml:"body,omitempty"`
	ArgsToJsonBody bool     `yaml:"argsToJsonBody,omitempty"`
	ArgsToUrlParam bool     `yaml:"argsToUrlParam,omitempty"`
	ArgsToFormBody bool     `yaml:"argsToFormBody,omitempty"`
}

// Header represents an HTTP header
type Header struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

// ResponseTemplate represents the MCP response template
type ResponseTemplate struct {
	Body        string `yaml:"body,omitempty"`
	PrependBody string `yaml:"prependBody,omitempty"`
	AppendBody  string `yaml:"appendBody,omitempty"`
}

// ConvertOptions represents options for the conversion process
type ConvertOptions struct {
	ServerName     string
	ServerConfig   map[string]interface{}
	ToolNamePrefix string
}
