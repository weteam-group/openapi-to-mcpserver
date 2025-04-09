package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/higress-group/openapi-to-mcpserver/pkg/converter"
	"github.com/higress-group/openapi-to-mcpserver/pkg/models"
	"github.com/higress-group/openapi-to-mcpserver/pkg/parser"
	"gopkg.in/yaml.v3"
)

func main() {
	// Define command-line flags
	inputFile := flag.String("input", "", "Path to the OpenAPI specification file (JSON or YAML)")
	outputFile := flag.String("output", "", "Path to the output MCP configuration file (YAML)")
	serverName := flag.String("server-name", "openapi-server", "Name of the MCP server")
	toolNamePrefix := flag.String("tool-prefix", "", "Prefix for tool names")
	format := flag.String("format", "yaml", "Output format (yaml or json)")

	// Parse command-line flags
	flag.Parse()

	// Validate required flags
	if *inputFile == "" {
		fmt.Println("Error: input file is required")
		flag.Usage()
		os.Exit(1)
	}

	if *outputFile == "" {
		fmt.Println("Error: output file is required")
		flag.Usage()
		os.Exit(1)
	}

	// Create a new parser
	p := parser.NewParser()

	// Parse the OpenAPI specification
	err := p.ParseFile(*inputFile)
	if err != nil {
		fmt.Printf("Error parsing OpenAPI specification: %v\n", err)
		os.Exit(1)
	}

	// Create a new converter
	c := converter.NewConverter(p, models.ConvertOptions{
		ServerName:     *serverName,
		ToolNamePrefix: *toolNamePrefix,
	})

	// Convert the OpenAPI specification to an MCP configuration
	config, err := c.Convert()
	if err != nil {
		fmt.Printf("Error converting OpenAPI specification: %v\n", err)
		os.Exit(1)
	}

	// Create the output directory if it doesn't exist
	outputDir := filepath.Dir(*outputFile)
	if outputDir != "" && outputDir != "." {
		err = os.MkdirAll(outputDir, 0755)
		if err != nil {
			fmt.Printf("Error creating output directory: %v\n", err)
			os.Exit(1)
		}
	}

	// Write the MCP configuration to the output file
	var data []byte
	if *format == "json" {
		data, err = json.MarshalIndent(config, "", "  ")
	} else {
		var buffer bytes.Buffer
		encoder := yaml.NewEncoder(&buffer)
		encoder.SetIndent(2)

		if err := encoder.Encode(config); err != nil {
			fmt.Printf("Error encoding YAML: %v\n", err)
			return
		}
		data = buffer.Bytes()
	}
	if err != nil {
		fmt.Printf("Error marshaling MCP configuration: %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile(*outputFile, data, 0644)
	if err != nil {
		fmt.Printf("Error writing MCP configuration: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully converted OpenAPI specification to MCP configuration: %s\n", *outputFile)
}
