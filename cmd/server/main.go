// Package main implements a comprehensive MCP server for Dynatrace Platform APIs.
package main

import (
	"fmt"
	"os"

	"github.com/dynatrace/dynatrace-platform-mcp-server/internal/client"
	"github.com/dynatrace/dynatrace-platform-mcp-server/internal/config"
	"github.com/dynatrace/dynatrace-platform-mcp-server/internal/server"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Configuration error: %v\n", err)
		os.Exit(1)
	}

	if err := cfg.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "Configuration validation error: %v\n", err)
		os.Exit(1)
	}

	httpClient := client.New(cfg)

	s := mcpserver.NewMCPServer(
		"Dynatrace Platform MCP Server",
		"2.0.0",
		mcpserver.WithToolCapabilities(true),
	)

	server.RegisterTools(s, httpClient, cfg)

	if err := mcpserver.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}
