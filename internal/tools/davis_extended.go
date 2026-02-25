package tools

import (
	"context"
	"fmt"

	"github.com/dynatrace/dynatrace-platform-mcp-server/internal/client"
	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

// RegisterDavisExtendedTools registers additional Davis AI tools not in the base set.
func RegisterDavisExtendedTools(s *mcpserver.MCPServer, h *Handlers, isEnabled func(string) bool) {
	// ==================== Davis Analyzer Extended ====================
	if isEnabled("dt_davis_analyzer_documentation") {
		s.AddTool(mcp.Tool{
			Name:        "dt_davis_analyzer_documentation",
			Description: `Get documentation for a Davis analyzer.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"name": map[string]interface{}{"type": "string", "description": "Analyzer name"},
				},
				Required: []string{"name"},
			},
		}, h.handleDavisAnalyzerDocumentation)
	}

	if isEnabled("dt_davis_analyzer_input_schema") {
		s.AddTool(mcp.Tool{
			Name:        "dt_davis_analyzer_input_schema",
			Description: `Get JSON schema for analyzer input parameters.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"name": map[string]interface{}{"type": "string", "description": "Analyzer name"},
				},
				Required: []string{"name"},
			},
		}, h.handleDavisAnalyzerInputSchema)
	}

	if isEnabled("dt_davis_analyzer_result_schema") {
		s.AddTool(mcp.Tool{
			Name:        "dt_davis_analyzer_result_schema",
			Description: `Get JSON schema for analyzer result.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"name": map[string]interface{}{"type": "string", "description": "Analyzer name"},
				},
				Required: []string{"name"},
			},
		}, h.handleDavisAnalyzerResultSchema)
	}

	if isEnabled("dt_davis_analyzer_validate") {
		s.AddTool(mcp.Tool{
			Name:        "dt_davis_analyzer_validate",
			Description: `Validate input parameters for an analyzer.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"name":  map[string]interface{}{"type": "string", "description": "Analyzer name"},
					"input": map[string]interface{}{"type": "object", "description": "Input parameters to validate"},
				},
				Required: []string{"name", "input"},
			},
		}, h.handleDavisAnalyzerValidate)
	}

	// ==================== CoPilot Extended - NL2DQL (HIGH VALUE!) ====================
	if isEnabled("dt_copilot_nl2dql") {
		s.AddTool(mcp.Tool{
			Name: "dt_copilot_nl2dql",
			Description: `Convert natural language to DQL query using Davis CoPilot.
	
	This is a high-value tool for autonomous query generation. Given a natural language
	description, it returns a valid DQL query.
	
	Example:
	  Input: "Show me all services with error rate above 5%"
	  Output: DQL query to fetch services with high error rates`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"text":    map[string]interface{}{"type": "string", "description": "Natural language description of the query"},
					"context": map[string]interface{}{"type": "object", "description": "Optional context (entity IDs, timeframe, etc.)"},
				},
				Required: []string{"text"},
			},
		}, h.handleCopilotNL2DQL)
	}

	if isEnabled("dt_copilot_dql2nl") {
		s.AddTool(mcp.Tool{
			Name:        "dt_copilot_dql2nl",
			Description: `Explain a DQL query in natural language using Davis CoPilot.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"query":   map[string]interface{}{"type": "string", "description": "DQL query to explain"},
					"context": map[string]interface{}{"type": "object", "description": "Optional context"},
				},
				Required: []string{"query"},
			},
		}, h.handleCopilotDQL2NL)
	}

	if isEnabled("dt_copilot_document_search") {
		s.AddTool(mcp.Tool{
			Name:        "dt_copilot_document_search",
			Description: `Search documents using Davis CoPilot.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"query":   map[string]interface{}{"type": "string", "description": "Search query"},
					"context": map[string]interface{}{"type": "object", "description": "Optional context"},
				},
				Required: []string{"query"},
			},
		}, h.handleCopilotDocumentSearch)
	}

	// ==================== CoPilot Feedback ====================
	if isEnabled("dt_copilot_conversation_feedback") {
		s.AddTool(mcp.Tool{
			Name:        "dt_copilot_conversation_feedback",
			Description: `Submit feedback for a CoPilot conversation.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"conversation_id": map[string]interface{}{"type": "string", "description": "Conversation ID"},
					"rating":          map[string]interface{}{"type": "string", "description": "Rating (positive, negative)"},
					"comment":         map[string]interface{}{"type": "string", "description": "Optional comment"},
				},
				Required: []string{"conversation_id", "rating"},
			},
		}, h.handleCopilotConversationFeedback)
	}

	if isEnabled("dt_copilot_nl2dql_feedback") {
		s.AddTool(mcp.Tool{
			Name:        "dt_copilot_nl2dql_feedback",
			Description: `Submit feedback for NL2DQL generation.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"request_id":      map[string]interface{}{"type": "string", "description": "Request ID"},
					"rating":          map[string]interface{}{"type": "string", "description": "Rating (positive, negative)"},
					"corrected_query": map[string]interface{}{"type": "string", "description": "Corrected DQL if rating is negative"},
					"comment":         map[string]interface{}{"type": "string", "description": "Optional comment"},
				},
				Required: []string{"request_id", "rating"},
			},
		}, h.handleCopilotNL2DQLFeedback)
	}

	if isEnabled("dt_copilot_dql2nl_feedback") {
		s.AddTool(mcp.Tool{
			Name:        "dt_copilot_dql2nl_feedback",
			Description: `Submit feedback for DQL2NL explanation.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"request_id": map[string]interface{}{"type": "string", "description": "Request ID"},
					"rating":     map[string]interface{}{"type": "string", "description": "Rating (positive, negative)"},
					"comment":    map[string]interface{}{"type": "string", "description": "Optional comment"},
				},
				Required: []string{"request_id", "rating"},
			},
		}, h.handleCopilotDQL2NLFeedback)
	}
}

// ==================== Handler Implementations ====================

func (h *Handlers) handleDavisAnalyzerDocumentation(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name := getStringParam(req.Params.Arguments, "name")
	if name == "" {
		return toolError(fmt.Errorf("name is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/davis/analyzers/v1/analyzers/"+name+"/documentation")
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	// Documentation might be text/markdown
	return toolResult(string(resp.Body)), nil
}

func (h *Handlers) handleDavisAnalyzerInputSchema(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name := getStringParam(req.Params.Arguments, "name")
	if name == "" {
		return toolError(fmt.Errorf("name is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/davis/analyzers/v1/analyzers/"+name+"/json-schema/input")
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

func (h *Handlers) handleDavisAnalyzerResultSchema(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name := getStringParam(req.Params.Arguments, "name")
	if name == "" {
		return toolError(fmt.Errorf("name is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/davis/analyzers/v1/analyzers/"+name+"/json-schema/result")
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

func (h *Handlers) handleDavisAnalyzerValidate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	name := getStringParam(args, "name")
	if name == "" {
		return toolError(fmt.Errorf("name is required")), nil
	}

	input, ok := args["input"].(map[string]interface{})
	if !ok {
		return toolError(fmt.Errorf("input is required")), nil
	}

	resp, err := h.Client.Post(ctx, "/platform/davis/analyzers/v1/analyzers/"+name+":validate", input)
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

func (h *Handlers) handleCopilotNL2DQL(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	text := getStringParam(args, "text")
	if text == "" {
		return toolError(fmt.Errorf("text is required")), nil
	}

	body := map[string]interface{}{
		"text": text,
	}
	if ctx, ok := args["context"].(map[string]interface{}); ok {
		body["context"] = ctx
	}

	resp, err := h.Client.Post(ctx, "/platform/davis/copilot/v1/skills/nl2dql:generate", body)
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

func (h *Handlers) handleCopilotDQL2NL(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

	resp, err := h.Client.Post(ctx, "/platform/davis/copilot/v1/skills/dql2nl:explain", body)
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

func (h *Handlers) handleCopilotDocumentSearch(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

	resp, err := h.Client.Post(ctx, "/platform/davis/copilot/v1/skills/document-search:execute", body)
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

func (h *Handlers) handleCopilotConversationFeedback(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	conversationID := getStringParam(args, "conversation_id")
	rating := getStringParam(args, "rating")
	if conversationID == "" || rating == "" {
		return toolError(fmt.Errorf("conversation_id and rating are required")), nil
	}

	body := map[string]interface{}{
		"conversationId": conversationID,
		"rating":         rating,
	}
	if comment := getStringParam(args, "comment"); comment != "" {
		body["comment"] = comment
	}

	resp, err := h.Client.Post(ctx, "/platform/davis/copilot/v1/skills/conversations:feedback", body)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult("Feedback submitted"), nil
}

func (h *Handlers) handleCopilotNL2DQLFeedback(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	requestID := getStringParam(args, "request_id")
	rating := getStringParam(args, "rating")
	if requestID == "" || rating == "" {
		return toolError(fmt.Errorf("request_id and rating are required")), nil
	}

	body := map[string]interface{}{
		"requestId": requestID,
		"rating":    rating,
	}
	if corrected := getStringParam(args, "corrected_query"); corrected != "" {
		body["correctedQuery"] = corrected
	}
	if comment := getStringParam(args, "comment"); comment != "" {
		body["comment"] = comment
	}

	resp, err := h.Client.Post(ctx, "/platform/davis/copilot/v1/skills/nl2dql:feedback", body)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult("NL2DQL feedback submitted"), nil
}

func (h *Handlers) handleCopilotDQL2NLFeedback(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	requestID := getStringParam(args, "request_id")
	rating := getStringParam(args, "rating")
	if requestID == "" || rating == "" {
		return toolError(fmt.Errorf("request_id and rating are required")), nil
	}

	body := map[string]interface{}{
		"requestId": requestID,
		"rating":    rating,
	}
	if comment := getStringParam(args, "comment"); comment != "" {
		body["comment"] = comment
	}

	resp, err := h.Client.Post(ctx, "/platform/davis/copilot/v1/skills/dql2nl:feedback", body)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult("DQL2NL feedback submitted"), nil
}
