package tools

import (
	"context"
	"fmt"

	"github.com/dynatrace/dynatrace-platform-mcp-server/internal/client"
	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

// RegisterAutomationExtendedTools registers additional automation tools not in the base set.
func RegisterAutomationExtendedTools(s *mcpserver.MCPServer, h *Handlers, isEnabled func(string) bool) {
	// ==================== Action Execution Logs ====================
	if isEnabled("dt_action_execution_log") {
		s.AddTool(mcp.Tool{
			Name:        "dt_action_execution_log",
			Description: `Get log output of a specific action execution.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":           map[string]interface{}{"type": "string", "description": "Action execution ID"},
					"admin_access": map[string]interface{}{"type": "boolean", "description": "Admin access"},
				},
				Required: []string{"id"},
			},
		}, h.handleActionExecutionLog)
	}

	// ==================== Execution Logs ====================
	if isEnabled("dt_execution_log") {
		s.AddTool(mcp.Tool{
			Name:        "dt_execution_log",
			Description: `Get log output of a workflow execution.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":           map[string]interface{}{"type": "string", "description": "Execution ID"},
					"admin_access": map[string]interface{}{"type": "boolean", "description": "Admin access"},
				},
				Required: []string{"id"},
			},
		}, h.handleExecutionLog)
	}

	if isEnabled("dt_execution_all_event_logs") {
		s.AddTool(mcp.Tool{
			Name:        "dt_execution_all_event_logs",
			Description: `Get all event logs for an execution.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":           map[string]interface{}{"type": "string", "description": "Execution ID"},
					"admin_access": map[string]interface{}{"type": "boolean", "description": "Admin access"},
				},
				Required: []string{"id"},
			},
		}, h.handleExecutionAllEventLogs)
	}

	if isEnabled("dt_execution_actions_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_execution_actions_list",
			Description: `List actions for an execution.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":           map[string]interface{}{"type": "string", "description": "Execution ID"},
					"admin_access": map[string]interface{}{"type": "boolean", "description": "Admin access"},
				},
				Required: []string{"id"},
			},
		}, h.handleExecutionActionsList)
	}

	// ==================== Execution Tasks ====================
	if isEnabled("dt_execution_tasks_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_execution_tasks_list",
			Description: `List tasks for an execution.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"execution_id": map[string]interface{}{"type": "string", "description": "Execution ID"},
					"admin_access": map[string]interface{}{"type": "boolean", "description": "Admin access"},
				},
				Required: []string{"execution_id"},
			},
		}, h.handleExecutionTasksList)
	}

	if isEnabled("dt_execution_task_get") {
		s.AddTool(mcp.Tool{
			Name:        "dt_execution_task_get",
			Description: `Get details of a specific task in an execution.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"execution_id": map[string]interface{}{"type": "string", "description": "Execution ID"},
					"task_id":      map[string]interface{}{"type": "string", "description": "Task ID"},
					"admin_access": map[string]interface{}{"type": "boolean", "description": "Admin access"},
				},
				Required: []string{"execution_id", "task_id"},
			},
		}, h.handleExecutionTaskGet)
	}

	if isEnabled("dt_execution_task_log") {
		s.AddTool(mcp.Tool{
			Name:        "dt_execution_task_log",
			Description: `Get log output of a specific task.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"execution_id": map[string]interface{}{"type": "string", "description": "Execution ID"},
					"task_id":      map[string]interface{}{"type": "string", "description": "Task ID"},
					"admin_access": map[string]interface{}{"type": "boolean", "description": "Admin access"},
				},
				Required: []string{"execution_id", "task_id"},
			},
		}, h.handleExecutionTaskLog)
	}

	if isEnabled("dt_execution_task_result") {
		s.AddTool(mcp.Tool{
			Name:        "dt_execution_task_result",
			Description: `Get result of a specific task.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"execution_id": map[string]interface{}{"type": "string", "description": "Execution ID"},
					"task_id":      map[string]interface{}{"type": "string", "description": "Task ID"},
					"admin_access": map[string]interface{}{"type": "boolean", "description": "Admin access"},
				},
				Required: []string{"execution_id", "task_id"},
			},
		}, h.handleExecutionTaskResult)
	}

	if isEnabled("dt_execution_task_input") {
		s.AddTool(mcp.Tool{
			Name:        "dt_execution_task_input",
			Description: `Get input of a specific task.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"execution_id": map[string]interface{}{"type": "string", "description": "Execution ID"},
					"task_id":      map[string]interface{}{"type": "string", "description": "Task ID"},
					"admin_access": map[string]interface{}{"type": "boolean", "description": "Admin access"},
				},
				Required: []string{"execution_id", "task_id"},
			},
		}, h.handleExecutionTaskInput)
	}

	if isEnabled("dt_execution_task_cancel") {
		s.AddTool(mcp.Tool{
			Name:        "dt_execution_task_cancel",
			Description: `Cancel a specific task in an execution.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"execution_id": map[string]interface{}{"type": "string", "description": "Execution ID"},
					"task_id":      map[string]interface{}{"type": "string", "description": "Task ID"},
					"admin_access": map[string]interface{}{"type": "boolean", "description": "Admin access"},
				},
				Required: []string{"execution_id", "task_id"},
			},
		}, h.handleExecutionTaskCancel)
	}

	if isEnabled("dt_execution_transitions_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_execution_transitions_list",
			Description: `List transitions for an execution.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"execution_id": map[string]interface{}{"type": "string", "description": "Execution ID"},
					"admin_access": map[string]interface{}{"type": "boolean", "description": "Admin access"},
				},
				Required: []string{"execution_id"},
			},
		}, h.handleExecutionTransitionsList)
	}

	// ==================== Workflow Extended ====================
	if isEnabled("dt_workflow_duplicate") {
		s.AddTool(mcp.Tool{
			Name:        "dt_workflow_duplicate",
			Description: `Duplicate a workflow.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":   map[string]interface{}{"type": "string", "description": "Workflow ID to duplicate"},
					"name": map[string]interface{}{"type": "string", "description": "Name for the duplicate"},
				},
				Required: []string{"id"},
			},
		}, h.handleWorkflowDuplicate)
	}

	if isEnabled("dt_workflow_export") {
		s.AddTool(mcp.Tool{
			Name:        "dt_workflow_export",
			Description: `Export a workflow.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{"type": "string", "description": "Workflow ID"},
				},
				Required: []string{"id"},
			},
		}, h.handleWorkflowExport)
	}

	if isEnabled("dt_workflow_history_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_workflow_history_list",
			Description: `List workflow version history.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{"type": "string", "description": "Workflow ID"},
				},
				Required: []string{"id"},
			},
		}, h.handleWorkflowHistoryList)
	}

	if isEnabled("dt_workflow_history_get") {
		s.AddTool(mcp.Tool{
			Name:        "dt_workflow_history_get",
			Description: `Get a specific version from workflow history.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":      map[string]interface{}{"type": "string", "description": "Workflow ID"},
					"version": map[string]interface{}{"type": "integer", "description": "Version number"},
				},
				Required: []string{"id", "version"},
			},
		}, h.handleWorkflowHistoryGet)
	}

	if isEnabled("dt_workflow_history_restore") {
		s.AddTool(mcp.Tool{
			Name:        "dt_workflow_history_restore",
			Description: `Restore a workflow to a previous version.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":      map[string]interface{}{"type": "string", "description": "Workflow ID"},
					"version": map[string]interface{}{"type": "integer", "description": "Version to restore"},
				},
				Required: []string{"id", "version"},
			},
		}, h.handleWorkflowHistoryRestore)
	}

	if isEnabled("dt_workflow_tasks_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_workflow_tasks_list",
			Description: `List tasks defined in a workflow.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{"type": "string", "description": "Workflow ID"},
				},
				Required: []string{"id"},
			},
		}, h.handleWorkflowTasksList)
	}

	if isEnabled("dt_workflow_reset_throttles") {
		s.AddTool(mcp.Tool{
			Name:        "dt_workflow_reset_throttles",
			Description: `Reset throttles for a workflow.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{"type": "string", "description": "Workflow ID"},
				},
				Required: []string{"id"},
			},
		}, h.handleWorkflowResetThrottles)
	}

	// ==================== Business Calendars CRUD ====================
	if isEnabled("dt_business_calendar_get") {
		s.AddTool(mcp.Tool{
			Name:        "dt_business_calendar_get",
			Description: `Get a business calendar by ID.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{"type": "string", "description": "Calendar ID"},
				},
				Required: []string{"id"},
			},
		}, h.handleBusinessCalendarGet)
	}

	if isEnabled("dt_business_calendar_create") {
		s.AddTool(mcp.Tool{
			Name:        "dt_business_calendar_create",
			Description: `Create a new business calendar.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"title":            map[string]interface{}{"type": "string", "description": "Calendar title"},
					"description":      map[string]interface{}{"type": "string", "description": "Description"},
					"timezone":         map[string]interface{}{"type": "string", "description": "Timezone"},
					"week_start":       map[string]interface{}{"type": "integer", "description": "Week start day (0=Monday)"},
					"week_days":        map[string]interface{}{"type": "array", "description": "Working days config"},
					"holidays":         map[string]interface{}{"type": "array", "description": "Holiday definitions"},
					"holiday_calendar": map[string]interface{}{"type": "string", "description": "Holiday calendar key"},
				},
				Required: []string{"title"},
			},
		}, h.handleBusinessCalendarCreate)
	}

	if isEnabled("dt_business_calendar_update") {
		s.AddTool(mcp.Tool{
			Name:        "dt_business_calendar_update",
			Description: `Update a business calendar.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":               map[string]interface{}{"type": "string", "description": "Calendar ID"},
					"title":            map[string]interface{}{"type": "string", "description": "Calendar title"},
					"description":      map[string]interface{}{"type": "string", "description": "Description"},
					"timezone":         map[string]interface{}{"type": "string", "description": "Timezone"},
					"week_start":       map[string]interface{}{"type": "integer", "description": "Week start day"},
					"week_days":        map[string]interface{}{"type": "array", "description": "Working days config"},
					"holidays":         map[string]interface{}{"type": "array", "description": "Holiday definitions"},
					"holiday_calendar": map[string]interface{}{"type": "string", "description": "Holiday calendar key"},
				},
				Required: []string{"id"},
			},
		}, h.handleBusinessCalendarUpdate)
	}

	if isEnabled("dt_business_calendar_delete") {
		s.AddTool(mcp.Tool{
			Name:        "dt_business_calendar_delete",
			Description: `Delete a business calendar.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{"type": "string", "description": "Calendar ID"},
				},
				Required: []string{"id"},
			},
		}, h.handleBusinessCalendarDelete)
	}

	if isEnabled("dt_business_calendar_duplicate") {
		s.AddTool(mcp.Tool{
			Name:        "dt_business_calendar_duplicate",
			Description: `Duplicate a business calendar.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":   map[string]interface{}{"type": "string", "description": "Calendar ID to duplicate"},
					"name": map[string]interface{}{"type": "string", "description": "Name for duplicate"},
				},
				Required: []string{"id"},
			},
		}, h.handleBusinessCalendarDuplicate)
	}

	if isEnabled("dt_business_calendar_history_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_business_calendar_history_list",
			Description: `List business calendar version history.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{"type": "string", "description": "Calendar ID"},
				},
				Required: []string{"id"},
			},
		}, h.handleBusinessCalendarHistoryList)
	}

	// ==================== Scheduling Rules CRUD ====================
	if isEnabled("dt_scheduling_rule_get") {
		s.AddTool(mcp.Tool{
			Name:        "dt_scheduling_rule_get",
			Description: `Get a scheduling rule by ID.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{"type": "string", "description": "Rule ID"},
				},
				Required: []string{"id"},
			},
		}, h.handleSchedulingRuleGet)
	}

	if isEnabled("dt_scheduling_rule_create") {
		s.AddTool(mcp.Tool{
			Name:        "dt_scheduling_rule_create",
			Description: `Create a new scheduling rule.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"title":             map[string]interface{}{"type": "string", "description": "Rule title"},
					"description":       map[string]interface{}{"type": "string", "description": "Description"},
					"rule_type":         map[string]interface{}{"type": "string", "description": "Rule type (cron, rrule, etc.)"},
					"rule":              map[string]interface{}{"type": "string", "description": "Rule expression"},
					"timezone":          map[string]interface{}{"type": "string", "description": "Timezone"},
					"business_calendar": map[string]interface{}{"type": "string", "description": "Business calendar ID"},
				},
				Required: []string{"title", "rule_type", "rule"},
			},
		}, h.handleSchedulingRuleCreate)
	}

	if isEnabled("dt_scheduling_rule_update") {
		s.AddTool(mcp.Tool{
			Name:        "dt_scheduling_rule_update",
			Description: `Update a scheduling rule.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":                map[string]interface{}{"type": "string", "description": "Rule ID"},
					"title":             map[string]interface{}{"type": "string", "description": "Rule title"},
					"description":       map[string]interface{}{"type": "string", "description": "Description"},
					"rule_type":         map[string]interface{}{"type": "string", "description": "Rule type"},
					"rule":              map[string]interface{}{"type": "string", "description": "Rule expression"},
					"timezone":          map[string]interface{}{"type": "string", "description": "Timezone"},
					"business_calendar": map[string]interface{}{"type": "string", "description": "Business calendar ID"},
				},
				Required: []string{"id"},
			},
		}, h.handleSchedulingRuleUpdate)
	}

	if isEnabled("dt_scheduling_rule_delete") {
		s.AddTool(mcp.Tool{
			Name:        "dt_scheduling_rule_delete",
			Description: `Delete a scheduling rule.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{"type": "string", "description": "Rule ID"},
				},
				Required: []string{"id"},
			},
		}, h.handleSchedulingRuleDelete)
	}

	if isEnabled("dt_scheduling_rule_duplicate") {
		s.AddTool(mcp.Tool{
			Name:        "dt_scheduling_rule_duplicate",
			Description: `Duplicate a scheduling rule.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":   map[string]interface{}{"type": "string", "description": "Rule ID to duplicate"},
					"name": map[string]interface{}{"type": "string", "description": "Name for duplicate"},
				},
				Required: []string{"id"},
			},
		}, h.handleSchedulingRuleDuplicate)
	}

	if isEnabled("dt_scheduling_rule_preview") {
		s.AddTool(mcp.Tool{
			Name:        "dt_scheduling_rule_preview",
			Description: `Preview scheduled times for a rule.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"rule_type":         map[string]interface{}{"type": "string", "description": "Rule type"},
					"rule":              map[string]interface{}{"type": "string", "description": "Rule expression"},
					"timezone":          map[string]interface{}{"type": "string", "description": "Timezone"},
					"business_calendar": map[string]interface{}{"type": "string", "description": "Business calendar ID"},
					"from":              map[string]interface{}{"type": "string", "description": "Start time (ISO)"},
					"to":                map[string]interface{}{"type": "string", "description": "End time (ISO)"},
				},
				Required: []string{"rule_type", "rule"},
			},
		}, h.handleSchedulingRulePreview)
	}

	// ==================== Schedules ====================
	if isEnabled("dt_holiday_calendars_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_holiday_calendars_list",
			Description: `List available holiday calendars.`,
			InputSchema: mcp.ToolInputSchema{
				Type:       "object",
				Properties: map[string]interface{}{},
			},
		}, h.handleHolidayCalendarsList)
	}

	if isEnabled("dt_holiday_calendar_get") {
		s.AddTool(mcp.Tool{
			Name:        "dt_holiday_calendar_get",
			Description: `Get a specific holiday calendar.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"key": map[string]interface{}{"type": "string", "description": "Holiday calendar key"},
				},
				Required: []string{"key"},
			},
		}, h.handleHolidayCalendarGet)
	}

	if isEnabled("dt_timezones_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_timezones_list",
			Description: `List available timezones.`,
			InputSchema: mcp.ToolInputSchema{
				Type:       "object",
				Properties: map[string]interface{}{},
			},
		}, h.handleTimezonesList)
	}

	if isEnabled("dt_schedule_preview") {
		s.AddTool(mcp.Tool{
			Name:        "dt_schedule_preview",
			Description: `Preview a schedule.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"schedule": map[string]interface{}{"type": "object", "description": "Schedule configuration"},
					"from":     map[string]interface{}{"type": "string", "description": "Start time"},
					"to":       map[string]interface{}{"type": "string", "description": "End time"},
				},
				Required: []string{"schedule"},
			},
		}, h.handleSchedulePreview)
	}

	// ==================== Event Triggers ====================
	if isEnabled("dt_event_trigger_filter_preview") {
		s.AddTool(mcp.Tool{
			Name:        "dt_event_trigger_filter_preview",
			Description: `Preview event trigger filter.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"filter_query": map[string]interface{}{"type": "string", "description": "DQL filter query"},
					"event_type":   map[string]interface{}{"type": "string", "description": "Event type"},
				},
				Required: []string{"filter_query", "event_type"},
			},
		}, h.handleEventTriggerFilterPreview)
	}

	// ==================== Automation Settings ====================
	if isEnabled("dt_automation_settings_get") {
		s.AddTool(mcp.Tool{
			Name:        "dt_automation_settings_get",
			Description: `Get automation settings.`,
			InputSchema: mcp.ToolInputSchema{
				Type:       "object",
				Properties: map[string]interface{}{},
			},
		}, h.handleAutomationSettingsGet)
	}

	if isEnabled("dt_automation_service_users_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_automation_service_users_list",
			Description: `List automation service users.`,
			InputSchema: mcp.ToolInputSchema{
				Type:       "object",
				Properties: map[string]interface{}{},
			},
		}, h.handleAutomationServiceUsersList)
	}

	if isEnabled("dt_automation_user_settings_get") {
		s.AddTool(mcp.Tool{
			Name:        "dt_automation_user_settings_get",
			Description: `Get current user's automation settings.`,
			InputSchema: mcp.ToolInputSchema{
				Type:       "object",
				Properties: map[string]interface{}{},
			},
		}, h.handleAutomationUserSettingsGet)
	}

	if isEnabled("dt_automation_user_permissions_get") {
		s.AddTool(mcp.Tool{
			Name:        "dt_automation_user_permissions_get",
			Description: `Get current user's automation permissions.`,
			InputSchema: mcp.ToolInputSchema{
				Type:       "object",
				Properties: map[string]interface{}{},
			},
		}, h.handleAutomationUserPermissionsGet)
	}

	if isEnabled("dt_automation_version_get") {
		s.AddTool(mcp.Tool{
			Name:        "dt_automation_version_get",
			Description: `Get automation engine version.`,
			InputSchema: mcp.ToolInputSchema{
				Type:       "object",
				Properties: map[string]interface{}{},
			},
		}, h.handleAutomationVersionGet)
	}
}

// ==================== Handler Implementations ====================

func (h *Handlers) handleActionExecutionLog(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	params := map[string]string{}
	if getBoolParam(args, "admin_access") {
		params["adminAccess"] = "true"
	}

	resp, err := h.Client.Get(ctx, "/platform/automation/v1/action-executions/"+id+"/log", client.WithQueryParams(params))
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(string(resp.Body)), nil
}

func (h *Handlers) handleExecutionLog(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	params := map[string]string{}
	if getBoolParam(args, "admin_access") {
		params["adminAccess"] = "true"
	}

	resp, err := h.Client.Get(ctx, "/platform/automation/v1/executions/"+id+"/log", client.WithQueryParams(params))
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(string(resp.Body)), nil
}

func (h *Handlers) handleExecutionAllEventLogs(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	params := map[string]string{}
	if getBoolParam(args, "admin_access") {
		params["adminAccess"] = "true"
	}

	resp, err := h.Client.Get(ctx, "/platform/automation/v1/executions/"+id+"/all-event-logs", client.WithQueryParams(params))
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

func (h *Handlers) handleExecutionActionsList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	params := map[string]string{}
	if getBoolParam(args, "admin_access") {
		params["adminAccess"] = "true"
	}

	resp, err := h.Client.Get(ctx, "/platform/automation/v1/executions/"+id+"/actions", client.WithQueryParams(params))
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

func (h *Handlers) handleExecutionTasksList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	execID := getStringParam(args, "execution_id")
	if execID == "" {
		return toolError(fmt.Errorf("execution_id is required")), nil
	}

	params := map[string]string{}
	if getBoolParam(args, "admin_access") {
		params["adminAccess"] = "true"
	}

	resp, err := h.Client.Get(ctx, "/platform/automation/v1/executions/"+execID+"/tasks", client.WithQueryParams(params))
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

func (h *Handlers) handleExecutionTaskGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	execID := getStringParam(args, "execution_id")
	taskID := getStringParam(args, "task_id")
	if execID == "" || taskID == "" {
		return toolError(fmt.Errorf("execution_id and task_id are required")), nil
	}

	params := map[string]string{}
	if getBoolParam(args, "admin_access") {
		params["adminAccess"] = "true"
	}

	resp, err := h.Client.Get(ctx, fmt.Sprintf("/platform/automation/v1/executions/%s/tasks/%s", execID, taskID), client.WithQueryParams(params))
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

func (h *Handlers) handleExecutionTaskLog(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	execID := getStringParam(args, "execution_id")
	taskID := getStringParam(args, "task_id")
	if execID == "" || taskID == "" {
		return toolError(fmt.Errorf("execution_id and task_id are required")), nil
	}

	params := map[string]string{}
	if getBoolParam(args, "admin_access") {
		params["adminAccess"] = "true"
	}

	resp, err := h.Client.Get(ctx, fmt.Sprintf("/platform/automation/v1/executions/%s/tasks/%s/log", execID, taskID), client.WithQueryParams(params))
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(string(resp.Body)), nil
}

func (h *Handlers) handleExecutionTaskResult(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	execID := getStringParam(args, "execution_id")
	taskID := getStringParam(args, "task_id")
	if execID == "" || taskID == "" {
		return toolError(fmt.Errorf("execution_id and task_id are required")), nil
	}

	params := map[string]string{}
	if getBoolParam(args, "admin_access") {
		params["adminAccess"] = "true"
	}

	resp, err := h.Client.Get(ctx, fmt.Sprintf("/platform/automation/v1/executions/%s/tasks/%s/result", execID, taskID), client.WithQueryParams(params))
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

func (h *Handlers) handleExecutionTaskInput(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	execID := getStringParam(args, "execution_id")
	taskID := getStringParam(args, "task_id")
	if execID == "" || taskID == "" {
		return toolError(fmt.Errorf("execution_id and task_id are required")), nil
	}

	params := map[string]string{}
	if getBoolParam(args, "admin_access") {
		params["adminAccess"] = "true"
	}

	resp, err := h.Client.Get(ctx, fmt.Sprintf("/platform/automation/v1/executions/%s/tasks/%s/input", execID, taskID), client.WithQueryParams(params))
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

func (h *Handlers) handleExecutionTaskCancel(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	execID := getStringParam(args, "execution_id")
	taskID := getStringParam(args, "task_id")
	if execID == "" || taskID == "" {
		return toolError(fmt.Errorf("execution_id and task_id are required")), nil
	}

	params := map[string]string{}
	if getBoolParam(args, "admin_access") {
		params["adminAccess"] = "true"
	}

	resp, err := h.Client.Post(ctx, fmt.Sprintf("/platform/automation/v1/executions/%s/tasks/%s/cancel", execID, taskID), nil, client.WithQueryParams(params))
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(fmt.Sprintf("Task %s cancelled", taskID)), nil
}

func (h *Handlers) handleExecutionTransitionsList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	execID := getStringParam(args, "execution_id")
	if execID == "" {
		return toolError(fmt.Errorf("execution_id is required")), nil
	}

	params := map[string]string{}
	if getBoolParam(args, "admin_access") {
		params["adminAccess"] = "true"
	}

	resp, err := h.Client.Get(ctx, "/platform/automation/v1/executions/"+execID+"/transitions", client.WithQueryParams(params))
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

func (h *Handlers) handleWorkflowDuplicate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	body := map[string]interface{}{}
	if name := getStringParam(args, "name"); name != "" {
		body["name"] = name
	}

	resp, err := h.Client.Post(ctx, "/platform/automation/v1/workflows/"+id+"/duplicate", body)
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

func (h *Handlers) handleWorkflowExport(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/automation/v1/workflows/"+id+"/export")
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

func (h *Handlers) handleWorkflowHistoryList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/automation/v1/workflows/"+id+"/history")
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

func (h *Handlers) handleWorkflowHistoryGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	version := getIntParam(args, "version", 0)
	if id == "" || version == 0 {
		return toolError(fmt.Errorf("id and version are required")), nil
	}

	resp, err := h.Client.Get(ctx, fmt.Sprintf("/platform/automation/v1/workflows/%s/history/%d", id, version))
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

func (h *Handlers) handleWorkflowHistoryRestore(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	version := getIntParam(args, "version", 0)
	if id == "" || version == 0 {
		return toolError(fmt.Errorf("id and version are required")), nil
	}

	resp, err := h.Client.Post(ctx, fmt.Sprintf("/platform/automation/v1/workflows/%s/history/%d/restore", id, version), nil)
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

func (h *Handlers) handleWorkflowTasksList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/automation/v1/workflows/"+id+"/tasks")
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

func (h *Handlers) handleWorkflowResetThrottles(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Post(ctx, "/platform/automation/v1/workflows/"+id+"/reset-throttles", nil)
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

func (h *Handlers) handleBusinessCalendarGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/automation/v1/business-calendars/"+id)
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

func (h *Handlers) handleBusinessCalendarCreate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	title := getStringParam(args, "title")
	if title == "" {
		return toolError(fmt.Errorf("title is required")), nil
	}

	body := map[string]interface{}{
		"title": title,
	}
	if desc := getStringParam(args, "description"); desc != "" {
		body["description"] = desc
	}
	if tz := getStringParam(args, "timezone"); tz != "" {
		body["timezone"] = tz
	}
	if ws := getIntParam(args, "week_start", -1); ws >= 0 {
		body["weekStart"] = ws
	}
	if wd, ok := args["week_days"].([]interface{}); ok {
		body["weekDays"] = wd
	}
	if h, ok := args["holidays"].([]interface{}); ok {
		body["holidays"] = h
	}
	if hc := getStringParam(args, "holiday_calendar"); hc != "" {
		body["holidayCalendar"] = hc
	}

	resp, err := h.Client.Post(ctx, "/platform/automation/v1/business-calendars", body)
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

func (h *Handlers) handleBusinessCalendarUpdate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	body := map[string]interface{}{}
	if title := getStringParam(args, "title"); title != "" {
		body["title"] = title
	}
	if desc := getStringParam(args, "description"); desc != "" {
		body["description"] = desc
	}
	if tz := getStringParam(args, "timezone"); tz != "" {
		body["timezone"] = tz
	}
	if ws := getIntParam(args, "week_start", -1); ws >= 0 {
		body["weekStart"] = ws
	}
	if wd, ok := args["week_days"].([]interface{}); ok {
		body["weekDays"] = wd
	}
	if hol, ok := args["holidays"].([]interface{}); ok {
		body["holidays"] = hol
	}
	if hc := getStringParam(args, "holiday_calendar"); hc != "" {
		body["holidayCalendar"] = hc
	}

	resp, err := h.Client.Put(ctx, "/platform/automation/v1/business-calendars/"+id, body)
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

func (h *Handlers) handleBusinessCalendarDelete(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Delete(ctx, "/platform/automation/v1/business-calendars/"+id)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(fmt.Sprintf("Business calendar %s deleted", id)), nil
}

func (h *Handlers) handleBusinessCalendarDuplicate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	body := map[string]interface{}{}
	if name := getStringParam(args, "name"); name != "" {
		body["name"] = name
	}

	resp, err := h.Client.Post(ctx, "/platform/automation/v1/business-calendars/"+id+"/duplicate", body)
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

func (h *Handlers) handleBusinessCalendarHistoryList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/automation/v1/business-calendars/"+id+"/history")
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

func (h *Handlers) handleSchedulingRuleGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/automation/v1/scheduling-rules/"+id)
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

func (h *Handlers) handleSchedulingRuleCreate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	title := getStringParam(args, "title")
	ruleType := getStringParam(args, "rule_type")
	rule := getStringParam(args, "rule")
	if title == "" || ruleType == "" || rule == "" {
		return toolError(fmt.Errorf("title, rule_type, and rule are required")), nil
	}

	body := map[string]interface{}{
		"title":    title,
		"ruleType": ruleType,
		"rule":     rule,
	}
	if desc := getStringParam(args, "description"); desc != "" {
		body["description"] = desc
	}
	if tz := getStringParam(args, "timezone"); tz != "" {
		body["timezone"] = tz
	}
	if bc := getStringParam(args, "business_calendar"); bc != "" {
		body["businessCalendar"] = bc
	}

	resp, err := h.Client.Post(ctx, "/platform/automation/v1/scheduling-rules", body)
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

func (h *Handlers) handleSchedulingRuleUpdate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	body := map[string]interface{}{}
	if title := getStringParam(args, "title"); title != "" {
		body["title"] = title
	}
	if desc := getStringParam(args, "description"); desc != "" {
		body["description"] = desc
	}
	if rt := getStringParam(args, "rule_type"); rt != "" {
		body["ruleType"] = rt
	}
	if r := getStringParam(args, "rule"); r != "" {
		body["rule"] = r
	}
	if tz := getStringParam(args, "timezone"); tz != "" {
		body["timezone"] = tz
	}
	if bc := getStringParam(args, "business_calendar"); bc != "" {
		body["businessCalendar"] = bc
	}

	resp, err := h.Client.Put(ctx, "/platform/automation/v1/scheduling-rules/"+id, body)
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

func (h *Handlers) handleSchedulingRuleDelete(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Delete(ctx, "/platform/automation/v1/scheduling-rules/"+id)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(fmt.Sprintf("Scheduling rule %s deleted", id)), nil
}

func (h *Handlers) handleSchedulingRuleDuplicate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	body := map[string]interface{}{}
	if name := getStringParam(args, "name"); name != "" {
		body["name"] = name
	}

	resp, err := h.Client.Post(ctx, "/platform/automation/v1/scheduling-rules/"+id+"/duplicate", body)
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

func (h *Handlers) handleSchedulingRulePreview(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	ruleType := getStringParam(args, "rule_type")
	rule := getStringParam(args, "rule")
	if ruleType == "" || rule == "" {
		return toolError(fmt.Errorf("rule_type and rule are required")), nil
	}

	body := map[string]interface{}{
		"ruleType": ruleType,
		"rule":     rule,
	}
	if tz := getStringParam(args, "timezone"); tz != "" {
		body["timezone"] = tz
	}
	if bc := getStringParam(args, "business_calendar"); bc != "" {
		body["businessCalendar"] = bc
	}
	if from := getStringParam(args, "from"); from != "" {
		body["from"] = from
	}
	if to := getStringParam(args, "to"); to != "" {
		body["to"] = to
	}

	resp, err := h.Client.Post(ctx, "/platform/automation/v1/scheduling-rules/preview", body)
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

func (h *Handlers) handleHolidayCalendarsList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	resp, err := h.Client.Get(ctx, "/platform/automation/v1/schedules/holiday-calendars")
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

func (h *Handlers) handleHolidayCalendarGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	key := getStringParam(req.Params.Arguments, "key")
	if key == "" {
		return toolError(fmt.Errorf("key is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/automation/v1/schedules/holiday-calendars/"+key)
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

func (h *Handlers) handleTimezonesList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	resp, err := h.Client.Get(ctx, "/platform/automation/v1/schedules/timezones")
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

func (h *Handlers) handleSchedulePreview(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	schedule, ok := args["schedule"].(map[string]interface{})
	if !ok {
		return toolError(fmt.Errorf("schedule is required")), nil
	}

	body := map[string]interface{}{
		"schedule": schedule,
	}
	if from := getStringParam(args, "from"); from != "" {
		body["from"] = from
	}
	if to := getStringParam(args, "to"); to != "" {
		body["to"] = to
	}

	resp, err := h.Client.Post(ctx, "/platform/automation/v1/schedules/preview", body)
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

func (h *Handlers) handleEventTriggerFilterPreview(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	filterQuery := getStringParam(args, "filter_query")
	eventType := getStringParam(args, "event_type")
	if filterQuery == "" || eventType == "" {
		return toolError(fmt.Errorf("filter_query and event_type are required")), nil
	}

	body := map[string]interface{}{
		"filterQuery": filterQuery,
		"eventType":   eventType,
	}

	resp, err := h.Client.Post(ctx, "/platform/automation/v1/event-triggers/filter-preview", body)
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

func (h *Handlers) handleAutomationSettingsGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	resp, err := h.Client.Get(ctx, "/platform/automation/v1/settings")
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

func (h *Handlers) handleAutomationServiceUsersList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	resp, err := h.Client.Get(ctx, "/platform/automation/v1/settings/service-users")
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

func (h *Handlers) handleAutomationUserSettingsGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	resp, err := h.Client.Get(ctx, "/platform/automation/v1/settings/user")
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

func (h *Handlers) handleAutomationUserPermissionsGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	resp, err := h.Client.Get(ctx, "/platform/automation/v1/settings/user-permissions")
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

func (h *Handlers) handleAutomationVersionGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	resp, err := h.Client.Get(ctx, "/platform/automation/v1/version")
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
