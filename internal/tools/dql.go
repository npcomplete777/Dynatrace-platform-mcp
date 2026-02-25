package tools

import (
	"context"
	"fmt"

	"github.com/dynatrace/dynatrace-platform-mcp-server/internal/client"
	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

// RegisterDQLTools registers DQL query tools.
func RegisterDQLTools(s *mcpserver.MCPServer, h *Handlers, isEnabled func(string) bool) {
	if isEnabled("dt_dql_query") {
		s.AddTool(mcp.Tool{
			Name: "dt_dql_query",
			Description: `Execute a DQL (Dynatrace Query Language) query against Grail.
	
	This is the primary tool for querying Dynatrace data. DQL supports:
	- Logs: fetch logs | filter ...
	- Metrics: timeseries avg(dt.host.cpu.usage)
	- Events: fetch events | filter ...
	- Entities: fetch dt.entity.service | fields ...
	- Bizevents: fetch bizevents | filter ...
	- Spans: fetch spans | filter ...
	
	Args:
	  - query (string, required): The DQL query to execute
	  - max_result_records (int, optional): Maximum records to return (default: 1000)
	  - request_timeout_ms (int, optional): Timeout in milliseconds (default: 60000)
	
	Returns:
	  Query results with records and metadata.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"query":              map[string]interface{}{"type": "string", "description": "DQL query to execute"},
					"max_result_records": map[string]interface{}{"type": "integer", "description": "Max records (default 1000)"},
					"request_timeout_ms": map[string]interface{}{"type": "integer", "description": "Timeout in ms (default 60000)"},
				},
				Required: []string{"query"},
			},
		}, h.handleDQLQuery)
	}

	if isEnabled("dt_dql_autocomplete") {
		s.AddTool(mcp.Tool{
			Name: "dt_dql_autocomplete",
			Description: `Get DQL query autocomplete suggestions.
	
	Args:
	  - query (string, required): Partial DQL query
	  - cursor_position (int, optional): Cursor position for context-aware suggestions
	
	Returns:
	  Autocomplete suggestions for the query.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"query":           map[string]interface{}{"type": "string", "description": "Partial DQL query"},
					"cursor_position": map[string]interface{}{"type": "integer", "description": "Cursor position"},
				},
				Required: []string{"query"},
			},
		}, h.handleDQLAutocomplete)
	}

	if isEnabled("dt_dql_parse") {
		s.AddTool(mcp.Tool{
			Name: "dt_dql_parse",
			Description: `Parse and validate a DQL query without executing it.
	
	Args:
	  - query (string, required): DQL query to parse
	
	Returns:
	  Parse result with AST or validation errors.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"query": map[string]interface{}{"type": "string", "description": "DQL query to parse"},
				},
				Required: []string{"query"},
			},
		}, h.handleDQLParse)
	}
}

// DQL Handlers

func (h *Handlers) handleDQLQuery(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments

	query := getStringParam(args, "query")
	if query == "" {
		return toolError(fmt.Errorf("query parameter is required")), nil
	}

	body := map[string]interface{}{
		"query": query,
	}

	if maxRecords := getIntParam(args, "max_result_records", 0); maxRecords > 0 {
		body["maxResultRecords"] = maxRecords
	}
	if timeout := getIntParam(args, "request_timeout_ms", 0); timeout > 0 {
		body["requestTimeoutMilliseconds"] = timeout
	}

	resp, err := h.Client.Post(ctx, "/platform/storage/query/v1/query:execute", body)
	if err != nil {
		return toolError(err), nil
	}

	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}

	return toolResult(resp.Body), nil
}

func (h *Handlers) handleDQLAutocomplete(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments

	query := getStringParam(args, "query")
	if query == "" {
		return toolError(fmt.Errorf("query parameter is required")), nil
	}

	body := map[string]interface{}{
		"query": query,
	}

	if cursorPos := getIntParam(args, "cursor_position", -1); cursorPos >= 0 {
		body["cursorPosition"] = cursorPos
	}

	resp, err := h.Client.Post(ctx, "/platform/storage/query/v1/query:autocomplete", body)
	if err != nil {
		return toolError(err), nil
	}

	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}

	return toolResult(resp.Body), nil
}

func (h *Handlers) handleDQLParse(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments

	query := getStringParam(args, "query")
	if query == "" {
		return toolError(fmt.Errorf("query parameter is required")), nil
	}

	body := map[string]interface{}{
		"query": query,
	}

	resp, err := h.Client.Post(ctx, "/platform/storage/query/v1/query:parse", body)
	if err != nil {
		return toolError(err), nil
	}

	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}

	return toolResult(resp.Body), nil
}
