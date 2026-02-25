// Package tools provides MCP tool handlers for Dynatrace Platform APIs.
package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/npcomplete777/Dynatrace-platform-mcp/internal/client"
	"github.com/mark3labs/mcp-go/mcp"
)

// Handlers holds the Dynatrace API client and provides tool handler methods.
type Handlers struct {
	Client *client.Client
}

// Helper functions for parameter extraction

func getStringParam(args map[string]interface{}, key string) string {
	if v, ok := args[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func getIntParam(args map[string]interface{}, key string, defaultVal int) int {
	if v, ok := args[key]; ok {
		switch n := v.(type) {
		case float64:
			return int(n)
		case int:
			return n
		}
	}
	return defaultVal
}

func getBoolParam(args map[string]interface{}, key string) bool {
	if v, ok := args[key]; ok {
		if b, ok := v.(bool); ok {
			return b
		}
	}
	return false
}

func getMapParam(args map[string]interface{}, key string) map[string]interface{} {
	if v, ok := args[key]; ok {
		if m, ok := v.(map[string]interface{}); ok {
			return m
		}
	}
	return nil
}

func getSliceParam(args map[string]interface{}, key string) []interface{} {
	if v, ok := args[key]; ok {
		if s, ok := v.([]interface{}); ok {
			return s
		}
	}
	return nil
}

func getStringSliceParam(args map[string]interface{}, key string) []string {
	if v, ok := args[key]; ok {
		switch s := v.(type) {
		case []interface{}:
			result := make([]string, 0, len(s))
			for _, item := range s {
				if str, ok := item.(string); ok {
					result = append(result, str)
				}
			}
			return result
		case []string:
			return s
		}
	}
	return nil
}

// Helper functions for results

func toolError(err error) *mcp.CallToolResult {
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: fmt.Sprintf("Error: %v", err),
			},
		},
		IsError: true,
	}
}

func textResult(text string) *mcp.CallToolResult {
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: text,
			},
		},
	}
}

func jsonResult(data interface{}) *mcp.CallToolResult {
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return toolError(fmt.Errorf("marshal result: %w", err))
	}
	return textResult(string(jsonBytes))
}

// toolResult formats data as a tool result - accepts []byte or string
func toolResult(data interface{}) *mcp.CallToolResult {
	switch v := data.(type) {
	case []byte:
		var parsed interface{}
		if err := json.Unmarshal(v, &parsed); err != nil {
			return textResult(string(v))
		}
		return jsonResult(parsed)
	case string:
		return textResult(v)
	default:
		return jsonResult(data)
	}
}

// rawResult formats a Response as a tool result
func rawResult(resp *client.Response) *mcp.CallToolResult {
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp)))
	}
	var parsed interface{}
	if err := json.Unmarshal(resp.Body, &parsed); err != nil {
		return textResult(string(resp.Body))
	}
	return jsonResult(parsed)
}

// Context helper
func toolCtx() context.Context {
	return context.Background()
}
