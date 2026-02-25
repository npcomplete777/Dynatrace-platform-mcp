package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/npcomplete777/Dynatrace-platform-mcp/internal/client"
	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

// ==================== Automation Tools ====================

func RegisterAutomationTools(s *mcpserver.MCPServer, h *Handlers, isEnabled func(string) bool) {
	if isEnabled("dt_workflows_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_workflows_list",
			Description: "List automation workflows with optional filtering and pagination.",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"limit":  map[string]interface{}{"type": "integer", "description": "Results per page"},
					"offset": map[string]interface{}{"type": "integer", "description": "Pagination offset"},
					"search": map[string]interface{}{"type": "string", "description": "Search term"},
				},
			},
		}, h.handleWorkflowsList)
	}

	if isEnabled("dt_workflow_get") {
		s.AddTool(mcp.Tool{
			Name:        "dt_workflow_get",
			Description: "Get details of a specific workflow by ID.",
			InputSchema: mcp.ToolInputSchema{
				Type:       "object",
				Properties: map[string]interface{}{"id": map[string]interface{}{"type": "string", "description": "Workflow ID"}},
				Required:   []string{"id"},
			},
		}, h.handleWorkflowGet)
	}

	if isEnabled("dt_workflow_create") {
		s.AddTool(mcp.Tool{
			Name:        "dt_workflow_create",
			Description: "Create a new automation workflow.",
			InputSchema: mcp.ToolInputSchema{
				Type:       "object",
				Properties: map[string]interface{}{"workflow": map[string]interface{}{"type": "object", "description": "Workflow definition"}},
				Required:   []string{"workflow"},
			},
		}, h.handleWorkflowCreate)
	}

	if isEnabled("dt_workflow_update") {
		s.AddTool(mcp.Tool{
			Name:        "dt_workflow_update",
			Description: "Update an existing workflow.",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":       map[string]interface{}{"type": "string", "description": "Workflow ID"},
					"workflow": map[string]interface{}{"type": "object", "description": "Updated workflow definition"},
				},
				Required: []string{"id", "workflow"},
			},
		}, h.handleWorkflowUpdate)
	}

	if isEnabled("dt_workflow_delete") {
		s.AddTool(mcp.Tool{
			Name:        "dt_workflow_delete",
			Description: "Delete a workflow.",
			InputSchema: mcp.ToolInputSchema{
				Type:       "object",
				Properties: map[string]interface{}{"id": map[string]interface{}{"type": "string", "description": "Workflow ID"}},
				Required:   []string{"id"},
			},
		}, h.handleWorkflowDelete)
	}

	if isEnabled("dt_workflow_run") {
		s.AddTool(mcp.Tool{
			Name:        "dt_workflow_run",
			Description: "Trigger a workflow execution.",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":     map[string]interface{}{"type": "string", "description": "Workflow ID"},
					"params": map[string]interface{}{"type": "object", "description": "Input parameters"},
				},
				Required: []string{"id"},
			},
		}, h.handleWorkflowRun)
	}

	if isEnabled("dt_workflow_executions") {
		s.AddTool(mcp.Tool{
			Name:        "dt_workflow_executions",
			Description: "List executions of a workflow.",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"workflow_id": map[string]interface{}{"type": "string", "description": "Workflow ID"},
					"limit":       map[string]interface{}{"type": "integer", "description": "Results per page"},
				},
			},
		}, h.handleWorkflowExecutions)
	}

	if isEnabled("dt_action_executions") {
		s.AddTool(mcp.Tool{
			Name:        "dt_action_executions",
			Description: "List action executions.",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"limit": map[string]interface{}{"type": "integer", "description": "Results per page"},
				},
			},
		}, h.handleActionExecutions)
	}
}

func (h *Handlers) handleWorkflowsList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	params := make(map[string]string)
	if l := getIntParam(req.Params.Arguments, "limit", 0); l > 0 {
		params["limit"] = fmt.Sprintf("%d", l)
	}
	if o := getIntParam(req.Params.Arguments, "offset", 0); o > 0 {
		params["offset"] = fmt.Sprintf("%d", o)
	}
	if s := getStringParam(req.Params.Arguments, "search"); s != "" {
		params["search"] = s
	}
	resp, err := h.Client.Get(ctx, "/platform/automation/v1/workflows", client.WithQueryParams(params))
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

func (h *Handlers) handleWorkflowGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}
	resp, err := h.Client.Get(ctx, "/platform/automation/v1/workflows/"+id)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

func (h *Handlers) handleWorkflowCreate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	workflow := getMapParam(req.Params.Arguments, "workflow")
	if workflow == nil {
		return toolError(fmt.Errorf("workflow is required")), nil
	}
	resp, err := h.Client.Post(ctx, "/platform/automation/v1/workflows", workflow)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

func (h *Handlers) handleWorkflowUpdate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}
	workflow := getMapParam(req.Params.Arguments, "workflow")
	if workflow == nil {
		return toolError(fmt.Errorf("workflow is required")), nil
	}
	resp, err := h.Client.Put(ctx, "/platform/automation/v1/workflows/"+id, workflow)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

func (h *Handlers) handleWorkflowDelete(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}
	resp, err := h.Client.Delete(ctx, "/platform/automation/v1/workflows/"+id)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

func (h *Handlers) handleWorkflowRun(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}
	params := getMapParam(req.Params.Arguments, "params")
	body := map[string]interface{}{}
	if params != nil {
		body["params"] = params
	}
	resp, err := h.Client.Post(ctx, "/platform/automation/v1/workflows/"+id+"/run", body)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

func (h *Handlers) handleWorkflowExecutions(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	params := make(map[string]string)
	if wid := getStringParam(req.Params.Arguments, "workflow_id"); wid != "" {
		params["workflow"] = wid
	}
	if l := getIntParam(req.Params.Arguments, "limit", 0); l > 0 {
		params["limit"] = fmt.Sprintf("%d", l)
	}
	resp, err := h.Client.Get(ctx, "/platform/automation/v1/executions", client.WithQueryParams(params))
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

func (h *Handlers) handleActionExecutions(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	params := make(map[string]string)
	if l := getIntParam(req.Params.Arguments, "limit", 0); l > 0 {
		params["limit"] = fmt.Sprintf("%d", l)
	}
	resp, err := h.Client.Get(ctx, "/platform/automation/v1/action-executions", client.WithQueryParams(params))
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

// ==================== Document Tools ====================

func RegisterDocumentTools(s *mcpserver.MCPServer, h *Handlers, isEnabled func(string) bool) {
	if isEnabled("dt_documents_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_documents_list",
			Description: "List Grail documents with optional filtering.",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"filter": map[string]interface{}{"type": "string", "description": "Filter expression"},
					"limit":  map[string]interface{}{"type": "integer", "description": "Max results"},
				},
			},
		}, h.handleDocumentsList)
	}

	if isEnabled("dt_document_get") {
		s.AddTool(mcp.Tool{
			Name:        "dt_document_get",
			Description: "Get a document by ID.",
			InputSchema: mcp.ToolInputSchema{
				Type:       "object",
				Properties: map[string]interface{}{"id": map[string]interface{}{"type": "string", "description": "Document ID"}},
				Required:   []string{"id"},
			},
		}, h.handleDocumentGet)
	}

	if isEnabled("dt_document_create") {
		s.AddTool(mcp.Tool{
			Name:        "dt_document_create",
			Description: "Create a new Grail document (dashboard, notebook, etc).",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"name":    map[string]interface{}{"type": "string", "description": "Document name"},
					"type":    map[string]interface{}{"type": "string", "description": "Document type (dashboard, notebook)"},
					"content": map[string]interface{}{"type": "object", "description": "Document content"},
				},
				Required: []string{"name", "type", "content"},
			},
		}, h.handleDocumentCreate)
	}

	if isEnabled("dt_document_update") {
		s.AddTool(mcp.Tool{
			Name:        "dt_document_update",
			Description: "Update an existing document.",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":      map[string]interface{}{"type": "string", "description": "Document ID"},
					"content": map[string]interface{}{"type": "object", "description": "Updated content"},
				},
				Required: []string{"id", "content"},
			},
		}, h.handleDocumentUpdate)
	}

	if isEnabled("dt_document_delete") {
		s.AddTool(mcp.Tool{
			Name:        "dt_document_delete",
			Description: "Delete a document.",
			InputSchema: mcp.ToolInputSchema{
				Type:       "object",
				Properties: map[string]interface{}{"id": map[string]interface{}{"type": "string", "description": "Document ID"}},
				Required:   []string{"id"},
			},
		}, h.handleDocumentDelete)
	}
}

func (h *Handlers) handleDocumentsList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	params := make(map[string]string)
	if f := getStringParam(req.Params.Arguments, "filter"); f != "" {
		params["filter"] = f
	}
	if l := getIntParam(req.Params.Arguments, "limit", 0); l > 0 {
		params["page-size"] = fmt.Sprintf("%d", l)
	}
	resp, err := h.Client.Get(ctx, "/platform/document/v1/documents", client.WithQueryParams(params))
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

func (h *Handlers) handleDocumentGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}
	resp, err := h.Client.Get(ctx, "/platform/document/v1/documents/"+id)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

func (h *Handlers) handleDocumentCreate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name := getStringParam(req.Params.Arguments, "name")
	docType := getStringParam(req.Params.Arguments, "type")
	content := getMapParam(req.Params.Arguments, "content")
	if name == "" || docType == "" || content == nil {
		return toolError(fmt.Errorf("name, type, and content are required")), nil
	}
	_ = map[string]interface{}{
		"name":    name,
		"type":    docType,
		"content": content,
	}
	contentJSON, _ := json.Marshal(content)
	fields := map[string]string{"name": name, "type": docType, "content": string(contentJSON)}
	resp, err := h.Client.PostMultipart(ctx, "/platform/document/v1/documents", fields)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

func (h *Handlers) handleDocumentUpdate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	content := getMapParam(req.Params.Arguments, "content")
	if id == "" || content == nil {
		return toolError(fmt.Errorf("id and content are required")), nil
	}
	resp, err := h.Client.Put(ctx, "/platform/document/v1/documents/"+id, content)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

func (h *Handlers) handleDocumentDelete(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}
	resp, err := h.Client.Delete(ctx, "/platform/document/v1/documents/"+id)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

// ==================== Davis Tools ====================

func RegisterDavisTools(s *mcpserver.MCPServer, h *Handlers, isEnabled func(string) bool) {
	if isEnabled("dt_davis_analyzers_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_davis_analyzers_list",
			Description: "List available Davis analyzers.",
			InputSchema: mcp.ToolInputSchema{Type: "object", Properties: map[string]interface{}{}},
		}, h.handleDavisAnalyzersList)
	}

	if isEnabled("dt_davis_analyze") {
		s.AddTool(mcp.Tool{
			Name:        "dt_davis_analyze",
			Description: "Run a Davis analyzer.",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"analyzer_name": map[string]interface{}{"type": "string", "description": "Analyzer name"},
					"input":         map[string]interface{}{"type": "object", "description": "Analyzer input"},
				},
				Required: []string{"analyzer_name", "input"},
			},
		}, h.handleDavisAnalyze)
	}
}

func (h *Handlers) handleDavisAnalyzersList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	resp, err := h.Client.Get(ctx, "/platform/davis/v1/analyze")
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

func (h *Handlers) handleDavisAnalyze(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name := getStringParam(req.Params.Arguments, "analyzer_name")
	input := getMapParam(req.Params.Arguments, "input")
	if name == "" || input == nil {
		return toolError(fmt.Errorf("analyzer_name and input are required")), nil
	}
	resp, err := h.Client.Post(ctx, "/platform/davis/v1/analyze/"+name, input)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

// ==================== OpenPipeline Tools ====================

func RegisterOpenPipelineTools(s *mcpserver.MCPServer, h *Handlers, isEnabled func(string) bool) {
	if isEnabled("dt_openpipeline_configs_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_openpipeline_configs_list",
			Description: "List OpenPipeline configurations.",
			InputSchema: mcp.ToolInputSchema{Type: "object", Properties: map[string]interface{}{}},
		}, h.handleOpenPipelineConfigsList)
	}

	if isEnabled("dt_openpipeline_config_get") {
		s.AddTool(mcp.Tool{
			Name:        "dt_openpipeline_config_get",
			Description: "Get an OpenPipeline configuration.",
			InputSchema: mcp.ToolInputSchema{
				Type:       "object",
				Properties: map[string]interface{}{"id": map[string]interface{}{"type": "string", "description": "Config ID"}},
				Required:   []string{"id"},
			},
		}, h.handleOpenPipelineConfigGet)
	}
}

func (h *Handlers) handleOpenPipelineConfigsList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	resp, err := h.Client.Get(ctx, "/platform/openpipeline/v1/configurations")
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

func (h *Handlers) handleOpenPipelineConfigGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}
	resp, err := h.Client.Get(ctx, "/platform/openpipeline/v1/configurations/"+id)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

// ==================== Notification Tools ====================

func RegisterNotificationTools(s *mcpserver.MCPServer, h *Handlers, isEnabled func(string) bool) {
	if isEnabled("dt_notifications_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_notifications_list",
			Description: "List notification configurations.",
			InputSchema: mcp.ToolInputSchema{Type: "object", Properties: map[string]interface{}{}},
		}, h.handleNotificationsList)
	}
}

func (h *Handlers) handleNotificationsList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	resp, err := h.Client.Get(ctx, "/platform/notification/v1/notifications")
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

// ==================== Storage Tools ====================

func RegisterStorageTools(s *mcpserver.MCPServer, h *Handlers, isEnabled func(string) bool) {
	if isEnabled("dt_buckets_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_buckets_list",
			Description: "List Grail buckets.",
			InputSchema: mcp.ToolInputSchema{Type: "object", Properties: map[string]interface{}{}},
		}, h.handleBucketsList)
	}

	if isEnabled("dt_bucket_get") {
		s.AddTool(mcp.Tool{
			Name:        "dt_bucket_get",
			Description: "Get bucket details.",
			InputSchema: mcp.ToolInputSchema{
				Type:       "object",
				Properties: map[string]interface{}{"name": map[string]interface{}{"type": "string", "description": "Bucket name"}},
				Required:   []string{"name"},
			},
		}, h.handleBucketGet)
	}
}

func (h *Handlers) handleBucketsList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	resp, err := h.Client.Get(ctx, "/platform/storage/management/v1/bucket-definitions")
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

func (h *Handlers) handleBucketGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name := getStringParam(req.Params.Arguments, "name")
	if name == "" {
		return toolError(fmt.Errorf("name is required")), nil
	}
	resp, err := h.Client.Get(ctx, "/platform/storage/management/v1/bucket-definitions/"+name)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

// ==================== Vulnerability Tools ====================

func RegisterVulnerabilityTools(s *mcpserver.MCPServer, h *Handlers, isEnabled func(string) bool) {
	if isEnabled("dt_vulnerabilities_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_vulnerabilities_list",
			Description: "List security vulnerabilities.",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"limit": map[string]interface{}{"type": "integer", "description": "Max results"},
				},
			},
		}, h.handleVulnerabilitiesList)
	}
}

func (h *Handlers) handleVulnerabilitiesList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	params := make(map[string]string)
	if l := getIntParam(req.Params.Arguments, "limit", 0); l > 0 {
		params["pageSize"] = fmt.Sprintf("%d", l)
	}
	resp, err := h.Client.Get(ctx, "/platform/app-security/v1/security-problems", client.WithQueryParams(params))
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

// ==================== Hub Tools ====================

func RegisterHubTools(s *mcpserver.MCPServer, h *Handlers, isEnabled func(string) bool) {
	if isEnabled("dt_hub_items_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_hub_items_list",
			Description: "List Hub items (extensions, apps).",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"type": map[string]interface{}{"type": "string", "description": "Item type filter"},
				},
			},
		}, h.handleHubItemsList)
	}
}

func (h *Handlers) handleHubItemsList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	params := make(map[string]string)
	if t := getStringParam(req.Params.Arguments, "type"); t != "" {
		params["itemType"] = t
	}
	resp, err := h.Client.Get(ctx, "/platform/hub/v1/items", client.WithQueryParams(params))
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

// ==================== IAM Tools ====================

func RegisterIAMTools(s *mcpserver.MCPServer, h *Handlers, isEnabled func(string) bool) {
	if isEnabled("dt_iam_groups_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_iam_groups_list",
			Description: "List IAM groups.",
			InputSchema: mcp.ToolInputSchema{Type: "object", Properties: map[string]interface{}{}},
		}, h.handleIAMGroupsList)
	}

	if isEnabled("dt_iam_policies_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_iam_policies_list",
			Description: "List IAM policies.",
			InputSchema: mcp.ToolInputSchema{Type: "object", Properties: map[string]interface{}{}},
		}, h.handleIAMPoliciesList)
	}
}

func (h *Handlers) handleIAMGroupsList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	accountURN := h.Client.GetConfig().GetAccountURN()
	if accountURN == "" {
		return toolError(fmt.Errorf("DT_ACCOUNT_URN environment variable not set")), nil
	}
	resp, err := h.Client.Get(ctx, "/iam/v1/accounts/"+accountURN+"/groups")
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

func (h *Handlers) handleIAMPoliciesList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	resp, err := h.Client.Get(ctx, "/iam/v1/policies")
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

// ==================== Platform Tools ====================

func RegisterPlatformTools(s *mcpserver.MCPServer, h *Handlers, isEnabled func(string) bool) {
	if isEnabled("dt_environment_info") {
		s.AddTool(mcp.Tool{
			Name:        "dt_environment_info",
			Description: "Get environment information.",
			InputSchema: mcp.ToolInputSchema{Type: "object", Properties: map[string]interface{}{}},
		}, h.handleEnvironmentInfo)
	}
}

func (h *Handlers) handleEnvironmentInfo(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	resp, err := h.Client.Get(ctx, "/platform/classic/environment-api/v2/entities")
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

// ==================== State Tools ====================

func RegisterStateTools(s *mcpserver.MCPServer, h *Handlers, isEnabled func(string) bool) {
	if isEnabled("dt_app_state_get") {
		s.AddTool(mcp.Tool{
			Name:        "dt_app_state_get",
			Description: "Get app state.",
			InputSchema: mcp.ToolInputSchema{
				Type:       "object",
				Properties: map[string]interface{}{"app_id": map[string]interface{}{"type": "string", "description": "App ID"}},
				Required:   []string{"app_id"},
			},
		}, h.handleAppStateGet)
	}
}

func (h *Handlers) handleAppStateGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	appID := getStringParam(req.Params.Arguments, "app_id")
	if appID == "" {
		return toolError(fmt.Errorf("app_id is required")), nil
	}
	resp, err := h.Client.Get(ctx, "/platform/state/v1/app-states/"+appID)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

// ==================== AppEngine Tools ====================

func RegisterAppEngineTools(s *mcpserver.MCPServer, h *Handlers, isEnabled func(string) bool) {
	if isEnabled("dt_apps_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_apps_list",
			Description: "List installed apps.",
			InputSchema: mcp.ToolInputSchema{Type: "object", Properties: map[string]interface{}{}},
		}, h.handleAppsList)
	}
}

func (h *Handlers) handleAppsList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	resp, err := h.Client.Get(ctx, "/platform/app-engine/registry/v1/apps")
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

// ==================== Email Tools ====================

func RegisterEmailTools(s *mcpserver.MCPServer, h *Handlers, isEnabled func(string) bool) {
	if isEnabled("dt_email_send") {
		s.AddTool(mcp.Tool{
			Name:        "dt_email_send",
			Description: "Send an email notification.",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"to":      map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": "string"}, "description": "Recipients"},
					"subject": map[string]interface{}{"type": "string", "description": "Email subject"},
					"body":    map[string]interface{}{"type": "string", "description": "Email body"},
				},
				Required: []string{"to", "subject", "body"},
			},
		}, h.handleEmailSend)
	}
}

func (h *Handlers) handleEmailSend(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	to := getStringSliceParam(req.Params.Arguments, "to")
	subject := getStringParam(req.Params.Arguments, "subject")
	body := getStringParam(req.Params.Arguments, "body")
	if len(to) == 0 || subject == "" || body == "" {
		return toolError(fmt.Errorf("to, subject, and body are required")), nil
	}
	emailBody := map[string]interface{}{
		"to":      to,
		"subject": subject,
		"body":    body,
	}
	resp, err := h.Client.Post(ctx, "/platform/email/v1/emails", emailBody)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}
