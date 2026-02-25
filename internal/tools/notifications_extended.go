package tools

import (
	"context"
	"fmt"

	"github.com/npcomplete777/Dynatrace-platform-mcp/internal/client"
	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

// RegisterNotificationsExtendedTools registers additional notification tools (v2 API).
func RegisterNotificationsExtendedTools(s *mcpserver.MCPServer, h *Handlers, isEnabled func(string) bool) {
	// ==================== Event Notifications ====================
	if isEnabled("dt_event_notifications_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_event_notifications_list",
			Description: `List event notifications.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"app_id":            map[string]interface{}{"type": "string", "description": "App ID filter"},
					"notification_type": map[string]interface{}{"type": "string", "description": "Notification type filter"},
					"resource_id":       map[string]interface{}{"type": "string", "description": "Resource ID filter"},
					"owner":             map[string]interface{}{"type": "string", "description": "Owner filter"},
					"limit":             map[string]interface{}{"type": "integer", "description": "Page size"},
					"offset":            map[string]interface{}{"type": "integer", "description": "Offset"},
				},
			},
		}, h.handleEventNotificationsList)
	}

	if isEnabled("dt_event_notification_create") {
		s.AddTool(mcp.Tool{
			Name:        "dt_event_notification_create",
			Description: `Create an event notification.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"notification_type": map[string]interface{}{"type": "string", "description": "Notification type"},
					"resource_id":       map[string]interface{}{"type": "string", "description": "Resource ID"},
					"trigger_config":    map[string]interface{}{"type": "object", "description": "Trigger configuration"},
				},
				Required: []string{"notification_type"},
			},
		}, h.handleEventNotificationCreate)
	}

	if isEnabled("dt_event_notification_get") {
		s.AddTool(mcp.Tool{
			Name:        "dt_event_notification_get",
			Description: `Get an event notification by ID.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{"type": "string", "description": "Notification ID"},
				},
				Required: []string{"id"},
			},
		}, h.handleEventNotificationGet)
	}

	if isEnabled("dt_event_notification_update") {
		s.AddTool(mcp.Tool{
			Name:        "dt_event_notification_update",
			Description: `Update an event notification.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":             map[string]interface{}{"type": "string", "description": "Notification ID"},
					"trigger_config": map[string]interface{}{"type": "object", "description": "Updated trigger configuration"},
					"enabled":        map[string]interface{}{"type": "boolean", "description": "Enable/disable notification"},
				},
				Required: []string{"id"},
			},
		}, h.handleEventNotificationUpdate)
	}

	if isEnabled("dt_event_notification_delete") {
		s.AddTool(mcp.Tool{
			Name:        "dt_event_notification_delete",
			Description: `Delete an event notification.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{"type": "string", "description": "Notification ID"},
				},
				Required: []string{"id"},
			},
		}, h.handleEventNotificationDelete)
	}

	// ==================== Resource Notifications ====================
	if isEnabled("dt_resource_notifications_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_resource_notifications_list",
			Description: `List resource notifications.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"app_id":            map[string]interface{}{"type": "string", "description": "App ID filter"},
					"notification_type": map[string]interface{}{"type": "string", "description": "Notification type filter"},
					"limit":             map[string]interface{}{"type": "integer", "description": "Page size"},
					"offset":            map[string]interface{}{"type": "integer", "description": "Offset"},
				},
			},
		}, h.handleResourceNotificationsList)
	}

	if isEnabled("dt_resource_notification_create") {
		s.AddTool(mcp.Tool{
			Name:        "dt_resource_notification_create",
			Description: `Create a resource notification.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"notification_type": map[string]interface{}{"type": "string", "description": "Notification type"},
					"resource_id":       map[string]interface{}{"type": "string", "description": "Resource ID"},
					"trigger_config":    map[string]interface{}{"type": "object", "description": "Trigger configuration"},
				},
				Required: []string{"notification_type", "resource_id"},
			},
		}, h.handleResourceNotificationCreate)
	}

	if isEnabled("dt_resource_notification_get_by_resource") {
		s.AddTool(mcp.Tool{
			Name:        "dt_resource_notification_get_by_resource",
			Description: `Get resource notification by notification type and resource ID.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"notification_type": map[string]interface{}{"type": "string", "description": "Notification type"},
					"resource_id":       map[string]interface{}{"type": "string", "description": "Resource ID"},
				},
				Required: []string{"notification_type", "resource_id"},
			},
		}, h.handleResourceNotificationGetByResource)
	}

	if isEnabled("dt_resource_notification_get") {
		s.AddTool(mcp.Tool{
			Name:        "dt_resource_notification_get",
			Description: `Get a resource notification by ID.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{"type": "string", "description": "Notification ID"},
				},
				Required: []string{"id"},
			},
		}, h.handleResourceNotificationGet)
	}

	if isEnabled("dt_resource_notification_update") {
		s.AddTool(mcp.Tool{
			Name:        "dt_resource_notification_update",
			Description: `Update a resource notification.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":             map[string]interface{}{"type": "string", "description": "Notification ID"},
					"trigger_config": map[string]interface{}{"type": "object", "description": "Updated trigger configuration"},
					"enabled":        map[string]interface{}{"type": "boolean", "description": "Enable/disable notification"},
				},
				Required: []string{"id"},
			},
		}, h.handleResourceNotificationUpdate)
	}

	if isEnabled("dt_resource_notification_delete") {
		s.AddTool(mcp.Tool{
			Name:        "dt_resource_notification_delete",
			Description: `Delete a resource notification.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{"type": "string", "description": "Notification ID"},
				},
				Required: []string{"id"},
			},
		}, h.handleResourceNotificationDelete)
	}

	// ==================== Self Notifications (v1) ====================
	if isEnabled("dt_self_notifications_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_self_notifications_list",
			Description: `List self notifications (v1).`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"limit":  map[string]interface{}{"type": "integer", "description": "Page size"},
					"offset": map[string]interface{}{"type": "integer", "description": "Offset"},
				},
			},
		}, h.handleSelfNotificationsList)
	}

	if isEnabled("dt_self_notification_create") {
		s.AddTool(mcp.Tool{
			Name:        "dt_self_notification_create",
			Description: `Create a self notification.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"notification_type": map[string]interface{}{"type": "string", "description": "Notification type"},
					"resource_id":       map[string]interface{}{"type": "string", "description": "Resource ID"},
					"trigger_config":    map[string]interface{}{"type": "object", "description": "Trigger configuration"},
				},
				Required: []string{"notification_type"},
			},
		}, h.handleSelfNotificationCreate)
	}

	if isEnabled("dt_self_notification_get") {
		s.AddTool(mcp.Tool{
			Name:        "dt_self_notification_get",
			Description: `Get a self notification by ID.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{"type": "string", "description": "Notification ID"},
				},
				Required: []string{"id"},
			},
		}, h.handleSelfNotificationGet)
	}

	if isEnabled("dt_self_notification_update") {
		s.AddTool(mcp.Tool{
			Name:        "dt_self_notification_update",
			Description: `Update a self notification.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":             map[string]interface{}{"type": "string", "description": "Notification ID"},
					"trigger_config": map[string]interface{}{"type": "object", "description": "Updated trigger configuration"},
					"enabled":        map[string]interface{}{"type": "boolean", "description": "Enable/disable notification"},
				},
				Required: []string{"id"},
			},
		}, h.handleSelfNotificationUpdate)
	}

	if isEnabled("dt_self_notification_delete") {
		s.AddTool(mcp.Tool{
			Name:        "dt_self_notification_delete",
			Description: `Delete a self notification.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{"type": "string", "description": "Notification ID"},
				},
				Required: []string{"id"},
			},
		}, h.handleSelfNotificationDelete)
	}
}

// ==================== Handler Implementations ====================

func (h *Handlers) handleEventNotificationsList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	params := map[string]string{}
	if appID := getStringParam(args, "app_id"); appID != "" {
		params["appId"] = appID
	}
	if notifType := getStringParam(args, "notification_type"); notifType != "" {
		params["notificationType"] = notifType
	}
	if resourceID := getStringParam(args, "resource_id"); resourceID != "" {
		params["resourceId"] = resourceID
	}
	if owner := getStringParam(args, "owner"); owner != "" {
		params["owner"] = owner
	}
	if limit := getIntParam(args, "limit", 0); limit > 0 {
		params["limit"] = fmt.Sprintf("%d", limit)
	}
	if offset := getIntParam(args, "offset", 0); offset > 0 {
		params["offset"] = fmt.Sprintf("%d", offset)
	}

	resp, err := h.Client.Get(ctx, "/platform/notification/v2/event-notifications", client.WithQueryParams(params))
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

func (h *Handlers) handleEventNotificationCreate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	notifType := getStringParam(args, "notification_type")
	if notifType == "" {
		return toolError(fmt.Errorf("notification_type is required")), nil
	}

	body := map[string]interface{}{
		"notificationType": notifType,
	}
	if resourceID := getStringParam(args, "resource_id"); resourceID != "" {
		body["resourceId"] = resourceID
	}
	if triggerConfig, ok := args["trigger_config"].(map[string]interface{}); ok {
		body["triggerConfig"] = triggerConfig
	}

	resp, err := h.Client.Post(ctx, "/platform/notification/v2/event-notifications", body)
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

func (h *Handlers) handleEventNotificationGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/notification/v2/event-notifications/"+id)
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

func (h *Handlers) handleEventNotificationUpdate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	body := map[string]interface{}{}
	if triggerConfig, ok := args["trigger_config"].(map[string]interface{}); ok {
		body["triggerConfig"] = triggerConfig
	}
	if enabled, ok := args["enabled"].(bool); ok {
		body["enabled"] = enabled
	}

	resp, err := h.Client.Put(ctx, "/platform/notification/v2/event-notifications/"+id, body)
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

func (h *Handlers) handleEventNotificationDelete(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Delete(ctx, "/platform/notification/v2/event-notifications/"+id)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(fmt.Sprintf("Event notification %s deleted", id)), nil
}

func (h *Handlers) handleResourceNotificationsList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	params := map[string]string{}
	if appID := getStringParam(args, "app_id"); appID != "" {
		params["appId"] = appID
	}
	if notifType := getStringParam(args, "notification_type"); notifType != "" {
		params["notificationType"] = notifType
	}
	if limit := getIntParam(args, "limit", 0); limit > 0 {
		params["limit"] = fmt.Sprintf("%d", limit)
	}
	if offset := getIntParam(args, "offset", 0); offset > 0 {
		params["offset"] = fmt.Sprintf("%d", offset)
	}

	resp, err := h.Client.Get(ctx, "/platform/notification/v2/resource-notifications", client.WithQueryParams(params))
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

func (h *Handlers) handleResourceNotificationCreate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	notifType := getStringParam(args, "notification_type")
	resourceID := getStringParam(args, "resource_id")
	if notifType == "" || resourceID == "" {
		return toolError(fmt.Errorf("notification_type and resource_id are required")), nil
	}

	body := map[string]interface{}{
		"notificationType": notifType,
		"resourceId":       resourceID,
	}
	if triggerConfig, ok := args["trigger_config"].(map[string]interface{}); ok {
		body["triggerConfig"] = triggerConfig
	}

	resp, err := h.Client.Post(ctx, "/platform/notification/v2/resource-notifications", body)
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

func (h *Handlers) handleResourceNotificationGetByResource(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	notifType := getStringParam(args, "notification_type")
	resourceID := getStringParam(args, "resource_id")
	if notifType == "" || resourceID == "" {
		return toolError(fmt.Errorf("notification_type and resource_id are required")), nil
	}

	resp, err := h.Client.Get(ctx, fmt.Sprintf("/platform/notification/v2/resource-notifications/%s/%s", notifType, resourceID))
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

func (h *Handlers) handleResourceNotificationGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/notification/v2/resource-notifications/"+id)
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

func (h *Handlers) handleResourceNotificationUpdate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	body := map[string]interface{}{}
	if triggerConfig, ok := args["trigger_config"].(map[string]interface{}); ok {
		body["triggerConfig"] = triggerConfig
	}
	if enabled, ok := args["enabled"].(bool); ok {
		body["enabled"] = enabled
	}

	resp, err := h.Client.Put(ctx, "/platform/notification/v2/resource-notifications/"+id, body)
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

func (h *Handlers) handleResourceNotificationDelete(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Delete(ctx, "/platform/notification/v2/resource-notifications/"+id)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(fmt.Sprintf("Resource notification %s deleted", id)), nil
}

func (h *Handlers) handleSelfNotificationsList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	params := map[string]string{}
	if limit := getIntParam(args, "limit", 0); limit > 0 {
		params["limit"] = fmt.Sprintf("%d", limit)
	}
	if offset := getIntParam(args, "offset", 0); offset > 0 {
		params["offset"] = fmt.Sprintf("%d", offset)
	}

	resp, err := h.Client.Get(ctx, "/platform/notification/v1/self-notifications", client.WithQueryParams(params))
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

func (h *Handlers) handleSelfNotificationCreate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	notifType := getStringParam(args, "notification_type")
	if notifType == "" {
		return toolError(fmt.Errorf("notification_type is required")), nil
	}

	body := map[string]interface{}{
		"notificationType": notifType,
	}
	if resourceID := getStringParam(args, "resource_id"); resourceID != "" {
		body["resourceId"] = resourceID
	}
	if triggerConfig, ok := args["trigger_config"].(map[string]interface{}); ok {
		body["triggerConfig"] = triggerConfig
	}

	resp, err := h.Client.Post(ctx, "/platform/notification/v1/self-notifications", body)
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

func (h *Handlers) handleSelfNotificationGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/notification/v1/self-notifications/"+id)
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

func (h *Handlers) handleSelfNotificationUpdate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	body := map[string]interface{}{}
	if triggerConfig, ok := args["trigger_config"].(map[string]interface{}); ok {
		body["triggerConfig"] = triggerConfig
	}
	if enabled, ok := args["enabled"].(bool); ok {
		body["enabled"] = enabled
	}

	resp, err := h.Client.Put(ctx, "/platform/notification/v1/self-notifications/"+id, body)
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

func (h *Handlers) handleSelfNotificationDelete(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Delete(ctx, "/platform/notification/v1/self-notifications/"+id)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(fmt.Sprintf("Self notification %s deleted", id)), nil
}
