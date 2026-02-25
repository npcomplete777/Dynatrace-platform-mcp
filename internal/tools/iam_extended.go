package tools

import (
	"context"
	"fmt"

	"github.com/dynatrace/dynatrace-platform-mcp-server/internal/client"
	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

// RegisterIAMExtendedTools registers additional IAM tools.
func RegisterIAMExtendedTools(s *mcpserver.MCPServer, h *Handlers, isEnabled func(string) bool) {
	// ==================== User Management ====================
	if isEnabled("dt_iam_user_get") {
		s.AddTool(mcp.Tool{
			Name:        "dt_iam_user_get",
			Description: `Get a specific IAM user by UUID.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"level_type": map[string]interface{}{"type": "string", "description": "Level type (account, environment)"},
					"level_id":   map[string]interface{}{"type": "string", "description": "Level ID"},
					"uuid":       map[string]interface{}{"type": "string", "description": "User UUID"},
				},
				Required: []string{"level_type", "level_id", "uuid"},
			},
		}, h.handleIAMUserGet)
	}

	if isEnabled("dt_iam_user_create") {
		s.AddTool(mcp.Tool{
			Name:        "dt_iam_user_create",
			Description: `Create a new IAM user.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"level_type": map[string]interface{}{"type": "string", "description": "Level type"},
					"level_id":   map[string]interface{}{"type": "string", "description": "Level ID"},
					"email":      map[string]interface{}{"type": "string", "description": "User email"},
					"groups":     map[string]interface{}{"type": "array", "description": "Group UUIDs to assign"},
				},
				Required: []string{"level_type", "level_id", "email"},
			},
		}, h.handleIAMUserCreate)
	}

	if isEnabled("dt_iam_user_update") {
		s.AddTool(mcp.Tool{
			Name:        "dt_iam_user_update",
			Description: `Update an IAM user.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"level_type": map[string]interface{}{"type": "string", "description": "Level type"},
					"level_id":   map[string]interface{}{"type": "string", "description": "Level ID"},
					"uuid":       map[string]interface{}{"type": "string", "description": "User UUID"},
					"groups":     map[string]interface{}{"type": "array", "description": "Updated group UUIDs"},
				},
				Required: []string{"level_type", "level_id", "uuid"},
			},
		}, h.handleIAMUserUpdate)
	}

	if isEnabled("dt_iam_user_delete") {
		s.AddTool(mcp.Tool{
			Name:        "dt_iam_user_delete",
			Description: `Delete an IAM user.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"level_type": map[string]interface{}{"type": "string", "description": "Level type"},
					"level_id":   map[string]interface{}{"type": "string", "description": "Level ID"},
					"uuid":       map[string]interface{}{"type": "string", "description": "User UUID"},
				},
				Required: []string{"level_type", "level_id", "uuid"},
			},
		}, h.handleIAMUserDelete)
	}

	// ==================== Service Users ====================
	if isEnabled("dt_iam_service_users_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_iam_service_users_list",
			Description: `List IAM service users.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"level_type": map[string]interface{}{"type": "string", "description": "Level type"},
					"level_id":   map[string]interface{}{"type": "string", "description": "Level ID"},
				},
				Required: []string{"level_type", "level_id"},
			},
		}, h.handleIAMServiceUsersList)
	}

	// ==================== Group Management ====================
	if isEnabled("dt_iam_group_get") {
		s.AddTool(mcp.Tool{
			Name:        "dt_iam_group_get",
			Description: `Get a specific IAM group by UUID.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"level_type": map[string]interface{}{"type": "string", "description": "Level type"},
					"level_id":   map[string]interface{}{"type": "string", "description": "Level ID"},
					"uuid":       map[string]interface{}{"type": "string", "description": "Group UUID"},
				},
				Required: []string{"level_type", "level_id", "uuid"},
			},
		}, h.handleIAMGroupGet)
	}

	if isEnabled("dt_iam_group_create") {
		s.AddTool(mcp.Tool{
			Name:        "dt_iam_group_create",
			Description: `Create a new IAM group.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"level_type":  map[string]interface{}{"type": "string", "description": "Level type"},
					"level_id":    map[string]interface{}{"type": "string", "description": "Level ID"},
					"name":        map[string]interface{}{"type": "string", "description": "Group name"},
					"description": map[string]interface{}{"type": "string", "description": "Group description"},
					"permissions": map[string]interface{}{"type": "array", "description": "Permission definitions"},
				},
				Required: []string{"level_type", "level_id", "name"},
			},
		}, h.handleIAMGroupCreate)
	}

	if isEnabled("dt_iam_group_update") {
		s.AddTool(mcp.Tool{
			Name:        "dt_iam_group_update",
			Description: `Update an IAM group.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"level_type":  map[string]interface{}{"type": "string", "description": "Level type"},
					"level_id":    map[string]interface{}{"type": "string", "description": "Level ID"},
					"uuid":        map[string]interface{}{"type": "string", "description": "Group UUID"},
					"name":        map[string]interface{}{"type": "string", "description": "New group name"},
					"description": map[string]interface{}{"type": "string", "description": "New description"},
					"permissions": map[string]interface{}{"type": "array", "description": "Updated permissions"},
				},
				Required: []string{"level_type", "level_id", "uuid"},
			},
		}, h.handleIAMGroupUpdate)
	}

	if isEnabled("dt_iam_group_delete") {
		s.AddTool(mcp.Tool{
			Name:        "dt_iam_group_delete",
			Description: `Delete an IAM group.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"level_type": map[string]interface{}{"type": "string", "description": "Level type"},
					"level_id":   map[string]interface{}{"type": "string", "description": "Level ID"},
					"uuid":       map[string]interface{}{"type": "string", "description": "Group UUID"},
				},
				Required: []string{"level_type", "level_id", "uuid"},
			},
		}, h.handleIAMGroupDelete)
	}

	// ==================== Platform Management Extended ====================
	if isEnabled("dt_environment_license_settings") {
		s.AddTool(mcp.Tool{
			Name:        "dt_environment_license_settings",
			Description: `Get license settings for the environment.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"keys": map[string]interface{}{"type": "array", "description": "Optional filter keys", "items": map[string]interface{}{"type": "string"}},
				},
			},
		}, h.handleEnvironmentLicenseSettings)
	}

	// ==================== State Management Extended ====================
	if isEnabled("dt_all_user_app_states_delete") {
		s.AddTool(mcp.Tool{
			Name:        "dt_all_user_app_states_delete",
			Description: `Delete all user app states for all users for an app.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"app_id": map[string]interface{}{"type": "string", "description": "App ID"},
				},
				Required: []string{"app_id"},
			},
		}, h.handleAllUserAppStatesDelete)
	}
}

// ==================== Handler Implementations ====================

func (h *Handlers) handleIAMUserGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	levelType := getStringParam(args, "level_type")
	levelID := getStringParam(args, "level_id")
	uuid := getStringParam(args, "uuid")
	if levelType == "" || levelID == "" || uuid == "" {
		return toolError(fmt.Errorf("level_type, level_id, and uuid are required")), nil
	}

	path := fmt.Sprintf("/platform/iam/v1/organizational-levels/%s/%s/users/%s", levelType, levelID, uuid)
	resp, err := h.Client.Get(ctx, path)
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

func (h *Handlers) handleIAMUserCreate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	levelType := getStringParam(args, "level_type")
	levelID := getStringParam(args, "level_id")
	email := getStringParam(args, "email")
	if levelType == "" || levelID == "" || email == "" {
		return toolError(fmt.Errorf("level_type, level_id, and email are required")), nil
	}

	body := map[string]interface{}{
		"email": email,
	}
	if groups, ok := args["groups"].([]interface{}); ok {
		body["groups"] = groups
	}

	path := fmt.Sprintf("/platform/iam/v1/organizational-levels/%s/%s/users", levelType, levelID)
	resp, err := h.Client.Post(ctx, path, body)
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

func (h *Handlers) handleIAMUserUpdate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	levelType := getStringParam(args, "level_type")
	levelID := getStringParam(args, "level_id")
	uuid := getStringParam(args, "uuid")
	if levelType == "" || levelID == "" || uuid == "" {
		return toolError(fmt.Errorf("level_type, level_id, and uuid are required")), nil
	}

	body := map[string]interface{}{}
	if groups, ok := args["groups"].([]interface{}); ok {
		body["groups"] = groups
	}

	path := fmt.Sprintf("/platform/iam/v1/organizational-levels/%s/%s/users/%s", levelType, levelID, uuid)
	resp, err := h.Client.Put(ctx, path, body)
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

func (h *Handlers) handleIAMUserDelete(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	levelType := getStringParam(args, "level_type")
	levelID := getStringParam(args, "level_id")
	uuid := getStringParam(args, "uuid")
	if levelType == "" || levelID == "" || uuid == "" {
		return toolError(fmt.Errorf("level_type, level_id, and uuid are required")), nil
	}

	path := fmt.Sprintf("/platform/iam/v1/organizational-levels/%s/%s/users/%s", levelType, levelID, uuid)
	resp, err := h.Client.Delete(ctx, path)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(fmt.Sprintf("User %s deleted", uuid)), nil
}

func (h *Handlers) handleIAMServiceUsersList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	levelType := getStringParam(args, "level_type")
	levelID := getStringParam(args, "level_id")
	if levelType == "" || levelID == "" {
		return toolError(fmt.Errorf("level_type and level_id are required")), nil
	}

	path := fmt.Sprintf("/platform/iam/v1/organizational-levels/%s/%s/service-users", levelType, levelID)
	resp, err := h.Client.Get(ctx, path)
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

func (h *Handlers) handleIAMGroupGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	levelType := getStringParam(args, "level_type")
	levelID := getStringParam(args, "level_id")
	uuid := getStringParam(args, "uuid")
	if levelType == "" || levelID == "" || uuid == "" {
		return toolError(fmt.Errorf("level_type, level_id, and uuid are required")), nil
	}

	path := fmt.Sprintf("/platform/iam/v1/organizational-levels/%s/%s/groups/%s", levelType, levelID, uuid)
	resp, err := h.Client.Get(ctx, path)
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

func (h *Handlers) handleIAMGroupCreate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	levelType := getStringParam(args, "level_type")
	levelID := getStringParam(args, "level_id")
	name := getStringParam(args, "name")
	if levelType == "" || levelID == "" || name == "" {
		return toolError(fmt.Errorf("level_type, level_id, and name are required")), nil
	}

	body := map[string]interface{}{
		"name": name,
	}
	if desc := getStringParam(args, "description"); desc != "" {
		body["description"] = desc
	}
	if permissions, ok := args["permissions"].([]interface{}); ok {
		body["permissions"] = permissions
	}

	path := fmt.Sprintf("/platform/iam/v1/organizational-levels/%s/%s/groups", levelType, levelID)
	resp, err := h.Client.Post(ctx, path, body)
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

func (h *Handlers) handleIAMGroupUpdate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	levelType := getStringParam(args, "level_type")
	levelID := getStringParam(args, "level_id")
	uuid := getStringParam(args, "uuid")
	if levelType == "" || levelID == "" || uuid == "" {
		return toolError(fmt.Errorf("level_type, level_id, and uuid are required")), nil
	}

	body := map[string]interface{}{}
	if name := getStringParam(args, "name"); name != "" {
		body["name"] = name
	}
	if desc := getStringParam(args, "description"); desc != "" {
		body["description"] = desc
	}
	if permissions, ok := args["permissions"].([]interface{}); ok {
		body["permissions"] = permissions
	}

	path := fmt.Sprintf("/platform/iam/v1/organizational-levels/%s/%s/groups/%s", levelType, levelID, uuid)
	resp, err := h.Client.Put(ctx, path, body)
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

func (h *Handlers) handleIAMGroupDelete(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	levelType := getStringParam(args, "level_type")
	levelID := getStringParam(args, "level_id")
	uuid := getStringParam(args, "uuid")
	if levelType == "" || levelID == "" || uuid == "" {
		return toolError(fmt.Errorf("level_type, level_id, and uuid are required")), nil
	}

	path := fmt.Sprintf("/platform/iam/v1/organizational-levels/%s/%s/groups/%s", levelType, levelID, uuid)
	resp, err := h.Client.Delete(ctx, path)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(fmt.Sprintf("Group %s deleted", uuid)), nil
}

func (h *Handlers) handleEnvironmentLicenseSettings(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	params := map[string]string{}
	if keys, ok := args["keys"].([]interface{}); ok && len(keys) > 0 {
		for _, k := range keys {
			if ks, ok := k.(string); ok {
				params["keys"] = ks // Note: API may need comma-separated or multiple params
			}
		}
	}

	resp, err := h.Client.Get(ctx, "/platform/management/v1/environment/license/settings", client.WithQueryParams(params))
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

func (h *Handlers) handleAllUserAppStatesDelete(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	appID := getStringParam(req.Params.Arguments, "app_id")
	if appID == "" {
		return toolError(fmt.Errorf("app_id is required")), nil
	}

	resp, err := h.Client.Delete(ctx, "/platform/state-management/v1/"+appID+"/user-app-states")
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(fmt.Sprintf("All user app states deleted for %s", appID)), nil
}
