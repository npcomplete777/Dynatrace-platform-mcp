package tools

import (
	"context"
	"fmt"

	"github.com/npcomplete777/Dynatrace-platform-mcp/internal/client"
	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

// RegisterOpenPipelineExtendedTools registers additional OpenPipeline tools.
func RegisterOpenPipelineExtendedTools(s *mcpserver.MCPServer, h *Handlers, isEnabled func(string) bool) {
	// ==================== Configuration Management ====================
	if isEnabled("dt_openpipeline_configuration_update") {
		s.AddTool(mcp.Tool{
			Name:        "dt_openpipeline_configuration_update",
			Description: `Update an OpenPipeline configuration.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":            map[string]interface{}{"type": "string", "description": "Configuration ID"},
					"version":       map[string]interface{}{"type": "string", "description": "Current version"},
					"configuration": map[string]interface{}{"type": "object", "description": "Updated configuration"},
				},
				Required: []string{"id", "version", "configuration"},
			},
		}, h.handleOpenPipelineConfigUpdate)
	}

	// ==================== DQL Processor ====================
	if isEnabled("dt_openpipeline_dql_autocomplete") {
		s.AddTool(mcp.Tool{
			Name:        "dt_openpipeline_dql_autocomplete",
			Description: `Get autocomplete suggestions for OpenPipeline DQL processor.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"query":           map[string]interface{}{"type": "string", "description": "Partial DQL query"},
					"cursor_position": map[string]interface{}{"type": "integer", "description": "Cursor position"},
					"context":         map[string]interface{}{"type": "object", "description": "Pipeline context"},
				},
				Required: []string{"query", "cursor_position"},
			},
		}, h.handleOpenPipelineDQLAutocomplete)
	}

	if isEnabled("dt_openpipeline_dql_verify") {
		s.AddTool(mcp.Tool{
			Name:        "dt_openpipeline_dql_verify",
			Description: `Verify OpenPipeline DQL processor syntax.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"query":   map[string]interface{}{"type": "string", "description": "DQL query to verify"},
					"context": map[string]interface{}{"type": "object", "description": "Pipeline context"},
				},
				Required: []string{"query"},
			},
		}, h.handleOpenPipelineDQLVerify)
	}

	// ==================== Matcher ====================
	if isEnabled("dt_openpipeline_matcher_autocomplete") {
		s.AddTool(mcp.Tool{
			Name:        "dt_openpipeline_matcher_autocomplete",
			Description: `Get autocomplete suggestions for OpenPipeline matcher.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"query":           map[string]interface{}{"type": "string", "description": "Partial matcher query"},
					"cursor_position": map[string]interface{}{"type": "integer", "description": "Cursor position"},
					"context":         map[string]interface{}{"type": "object", "description": "Pipeline context"},
				},
				Required: []string{"query", "cursor_position"},
			},
		}, h.handleOpenPipelineMatcherAutocomplete)
	}

	if isEnabled("dt_openpipeline_matcher_verify") {
		s.AddTool(mcp.Tool{
			Name:        "dt_openpipeline_matcher_verify",
			Description: `Verify OpenPipeline matcher syntax.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"query":   map[string]interface{}{"type": "string", "description": "Matcher query to verify"},
					"context": map[string]interface{}{"type": "object", "description": "Pipeline context"},
				},
				Required: []string{"query"},
			},
		}, h.handleOpenPipelineMatcherVerify)
	}

	if isEnabled("dt_openpipeline_lql_to_dql") {
		s.AddTool(mcp.Tool{
			Name:        "dt_openpipeline_lql_to_dql",
			Description: `Convert LQL (Log Query Language) to DQL.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"lql": map[string]interface{}{"type": "string", "description": "LQL query to convert"},
				},
				Required: []string{"lql"},
			},
		}, h.handleOpenPipelineLQLToDQL)
	}

	// ==================== Processor Preview ====================
	if isEnabled("dt_openpipeline_processor_preview") {
		s.AddTool(mcp.Tool{
			Name:        "dt_openpipeline_processor_preview",
			Description: `Preview processor output with sample data.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"processor":   map[string]interface{}{"type": "object", "description": "Processor configuration"},
					"sample_data": map[string]interface{}{"type": "array", "description": "Sample records to process"},
				},
				Required: []string{"processor", "sample_data"},
			},
		}, h.handleOpenPipelineProcessorPreview)
	}

	// ==================== Technologies ====================
	if isEnabled("dt_openpipeline_technologies_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_openpipeline_technologies_list",
			Description: `List OpenPipeline technologies.`,
			InputSchema: mcp.ToolInputSchema{
				Type:       "object",
				Properties: map[string]interface{}{},
			},
		}, h.handleOpenPipelineTechnologiesList)
	}

	if isEnabled("dt_openpipeline_technology_processors_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_openpipeline_technology_processors_list",
			Description: `List processors available for a technology.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"technology_id": map[string]interface{}{"type": "string", "description": "Technology ID"},
				},
				Required: []string{"technology_id"},
			},
		}, h.handleOpenPipelineTechnologyProcessorsList)
	}
}

// ==================== Handler Implementations ====================

func (h *Handlers) handleOpenPipelineConfigUpdate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	version := getStringParam(args, "version")
	if id == "" || version == "" {
		return toolError(fmt.Errorf("id and version are required")), nil
	}

	configuration, ok := args["configuration"].(map[string]interface{})
	if !ok {
		return toolError(fmt.Errorf("configuration is required")), nil
	}

	params := map[string]string{"optimistic-locking-version": version}
	resp, err := h.Client.Put(ctx, "/platform/openpipeline/v1/configurations/"+id, configuration, client.WithQueryParams(params))
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	var result interface{}
	resp.JSON(&result)
	return jsonResult(result), nil
}

func (h *Handlers) handleOpenPipelineDQLAutocomplete(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	query := getStringParam(args, "query")
	cursorPos := getIntParam(args, "cursor_position", 0)
	if query == "" {
		return toolError(fmt.Errorf("query is required")), nil
	}

	body := map[string]interface{}{
		"query":          query,
		"cursorPosition": cursorPos,
	}
	if ctx, ok := args["context"].(map[string]interface{}); ok {
		body["context"] = ctx
	}

	resp, err := h.Client.Post(ctx, "/platform/openpipeline/v1/dqlProcessor/autocomplete", body)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	var result interface{}
	resp.JSON(&result)
	return jsonResult(result), nil
}

func (h *Handlers) handleOpenPipelineDQLVerify(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	query := getStringParam(args, "query")
	if query == "" {
		return toolError(fmt.Errorf("query is required")), nil
	}

	body := map[string]interface{}{
		"query": query,
	}
	if ctx, ok := args["context"].(map[string]interface{}); ok {
		body["context"] = ctx
	}

	resp, err := h.Client.Post(ctx, "/platform/openpipeline/v1/dqlProcessor/verify", body)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	var result interface{}
	resp.JSON(&result)
	return jsonResult(result), nil
}

func (h *Handlers) handleOpenPipelineMatcherAutocomplete(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	query := getStringParam(args, "query")
	cursorPos := getIntParam(args, "cursor_position", 0)
	if query == "" {
		return toolError(fmt.Errorf("query is required")), nil
	}

	body := map[string]interface{}{
		"query":          query,
		"cursorPosition": cursorPos,
	}
	if ctx, ok := args["context"].(map[string]interface{}); ok {
		body["context"] = ctx
	}

	resp, err := h.Client.Post(ctx, "/platform/openpipeline/v1/matcher/autocomplete", body)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	var result interface{}
	resp.JSON(&result)
	return jsonResult(result), nil
}

func (h *Handlers) handleOpenPipelineMatcherVerify(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	query := getStringParam(args, "query")
	if query == "" {
		return toolError(fmt.Errorf("query is required")), nil
	}

	body := map[string]interface{}{
		"query": query,
	}
	if ctx, ok := args["context"].(map[string]interface{}); ok {
		body["context"] = ctx
	}

	resp, err := h.Client.Post(ctx, "/platform/openpipeline/v1/matcher/verify", body)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	var result interface{}
	resp.JSON(&result)
	return jsonResult(result), nil
}

func (h *Handlers) handleOpenPipelineLQLToDQL(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	lql := getStringParam(req.Params.Arguments, "lql")
	if lql == "" {
		return toolError(fmt.Errorf("lql is required")), nil
	}

	body := map[string]interface{}{
		"lql": lql,
	}

	resp, err := h.Client.Post(ctx, "/platform/openpipeline/v1/matcher/lqlToDql", body)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	var result interface{}
	resp.JSON(&result)
	return jsonResult(result), nil
}

func (h *Handlers) handleOpenPipelineProcessorPreview(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	processor, ok := args["processor"].(map[string]interface{})
	if !ok {
		return toolError(fmt.Errorf("processor is required")), nil
	}

	sampleData, ok := args["sample_data"].([]interface{})
	if !ok {
		return toolError(fmt.Errorf("sample_data is required")), nil
	}

	body := map[string]interface{}{
		"processor":  processor,
		"sampleData": sampleData,
	}

	resp, err := h.Client.Post(ctx, "/platform/openpipeline/v1/preview/processor", body)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	var result interface{}
	resp.JSON(&result)
	return jsonResult(result), nil
}

func (h *Handlers) handleOpenPipelineTechnologiesList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	resp, err := h.Client.Get(ctx, "/platform/openpipeline/v1/technologies")
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	var result interface{}
	resp.JSON(&result)
	return jsonResult(result), nil
}

func (h *Handlers) handleOpenPipelineTechnologyProcessorsList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	techID := getStringParam(req.Params.Arguments, "technology_id")
	if techID == "" {
		return toolError(fmt.Errorf("technology_id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/openpipeline/v1/technologies/"+techID+"/processors")
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	var result interface{}
	resp.JSON(&result)
	return jsonResult(result), nil
}
