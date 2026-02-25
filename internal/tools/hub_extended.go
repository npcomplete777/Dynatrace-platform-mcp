package tools

import (
	"context"
	"fmt"

	"github.com/npcomplete777/Dynatrace-platform-mcp/internal/client"
	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

// RegisterHubExtendedTools registers additional Hub/catalog tools not in the base set.
func RegisterHubExtendedTools(s *mcpserver.MCPServer, h *Handlers, isEnabled func(string) bool) {
	// ==================== App Releases ====================
	if isEnabled("dt_hub_app_releases_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_hub_app_releases_list",
			Description: `List releases for an app.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{"type": "string", "description": "App ID"},
				},
				Required: []string{"id"},
			},
		}, h.handleHubAppReleasesList)
	}

	// ==================== Extensions ====================
	if isEnabled("dt_hub_extension_get") {
		s.AddTool(mcp.Tool{
			Name:        "dt_hub_extension_get",
			Description: `Get extension details from Hub.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{"type": "string", "description": "Extension ID"},
				},
				Required: []string{"id"},
			},
		}, h.handleHubExtensionGet)
	}

	if isEnabled("dt_hub_extension_releases_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_hub_extension_releases_list",
			Description: `List releases for an extension.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{"type": "string", "description": "Extension ID"},
				},
				Required: []string{"id"},
			},
		}, h.handleHubExtensionReleasesList)
	}

	// ==================== Technologies ====================
	if isEnabled("dt_hub_technology_get") {
		s.AddTool(mcp.Tool{
			Name:        "dt_hub_technology_get",
			Description: `Get technology details from Hub.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{"type": "string", "description": "Technology ID"},
				},
				Required: []string{"id"},
			},
		}, h.handleHubTechnologyGet)
	}

	// ==================== Categories ====================
	if isEnabled("dt_hub_categories_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_hub_categories_list",
			Description: `List Hub categories.`,
			InputSchema: mcp.ToolInputSchema{
				Type:       "object",
				Properties: map[string]interface{}{},
			},
		}, h.handleHubCategoriesList)
	}
}

// ==================== Handler Implementations ====================

func (h *Handlers) handleHubAppReleasesList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/hub/v1/catalog/apps/"+id+"/releases")
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

func (h *Handlers) handleHubExtensionGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/hub/v1/catalog/extensions/"+id)
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

func (h *Handlers) handleHubExtensionReleasesList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/hub/v1/catalog/extensions/"+id+"/releases")
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

func (h *Handlers) handleHubTechnologyGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/hub/v1/catalog/technologies/"+id)
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

func (h *Handlers) handleHubCategoriesList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	resp, err := h.Client.Get(ctx, "/platform/hub/v1/catalog/categories")
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
