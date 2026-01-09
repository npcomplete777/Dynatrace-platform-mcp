package tools

import (
	"context"
	"fmt"

	"github.com/dynatrace/dynatrace-platform-mcp-server/internal/client"
	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

// RegisterStorageExtendedTools registers additional Grail storage tools not in the base set.
func RegisterStorageExtendedTools(s *mcpserver.MCPServer, h *Handlers) {
	// ==================== Bucket Extended ====================
	s.AddTool(mcp.Tool{
		Name:        "dt_bucket_update",
		Description: `Update a storage bucket.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"name":           map[string]interface{}{"type": "string", "description": "Bucket name"},
				"version":        map[string]interface{}{"type": "integer", "description": "Current version for optimistic locking"},
				"retention_days": map[string]interface{}{"type": "integer", "description": "New retention in days"},
				"display_name":   map[string]interface{}{"type": "string", "description": "Display name"},
			},
			Required: []string{"name", "version"},
		},
	}, h.handleBucketUpdate)

	s.AddTool(mcp.Tool{
		Name:        "dt_bucket_truncate",
		Description: `Truncate all data in a bucket.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"name": map[string]interface{}{"type": "string", "description": "Bucket name"},
			},
			Required: []string{"name"},
		},
	}, h.handleBucketTruncate)

	// ==================== Filter Segments ====================
	s.AddTool(mcp.Tool{
		Name:        "dt_filter_segments_list",
		Description: `List filter segments.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"filter":    map[string]interface{}{"type": "string", "description": "Filter query"},
				"page_size": map[string]interface{}{"type": "integer", "description": "Page size"},
			},
		},
	}, h.handleFilterSegmentsList)

	s.AddTool(mcp.Tool{
		Name:        "dt_filter_segment_get",
		Description: `Get a filter segment by UID.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"uid":            map[string]interface{}{"type": "string", "description": "Filter segment UID"},
				"add_fields":     map[string]interface{}{"type": "array", "description": "Additional fields to include"},
				"include_filter": map[string]interface{}{"type": "boolean", "description": "Include filter details"},
			},
			Required: []string{"uid"},
		},
	}, h.handleFilterSegmentGet)

	s.AddTool(mcp.Tool{
		Name:        "dt_filter_segment_create",
		Description: `Create a new filter segment.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"name":        map[string]interface{}{"type": "string", "description": "Segment name"},
				"description": map[string]interface{}{"type": "string", "description": "Description"},
				"is_public":   map[string]interface{}{"type": "boolean", "description": "Public visibility"},
				"variables":   map[string]interface{}{"type": "object", "description": "Variable definitions"},
				"includes":    map[string]interface{}{"type": "array", "description": "Include conditions"},
				"excludes":    map[string]interface{}{"type": "array", "description": "Exclude conditions"},
			},
			Required: []string{"name"},
		},
	}, h.handleFilterSegmentCreate)

	s.AddTool(mcp.Tool{
		Name:        "dt_filter_segment_update",
		Description: `Update a filter segment.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"uid":         map[string]interface{}{"type": "string", "description": "Filter segment UID"},
				"version":     map[string]interface{}{"type": "string", "description": "Current version"},
				"name":        map[string]interface{}{"type": "string", "description": "New name"},
				"description": map[string]interface{}{"type": "string", "description": "New description"},
				"is_public":   map[string]interface{}{"type": "boolean", "description": "Public visibility"},
				"variables":   map[string]interface{}{"type": "object", "description": "Variable definitions"},
				"includes":    map[string]interface{}{"type": "array", "description": "Include conditions"},
				"excludes":    map[string]interface{}{"type": "array", "description": "Exclude conditions"},
			},
			Required: []string{"uid", "version"},
		},
	}, h.handleFilterSegmentUpdate)

	s.AddTool(mcp.Tool{
		Name:        "dt_filter_segment_delete",
		Description: `Delete a filter segment.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"uid":     map[string]interface{}{"type": "string", "description": "Filter segment UID"},
				"version": map[string]interface{}{"type": "string", "description": "Current version"},
			},
			Required: []string{"uid", "version"},
		},
	}, h.handleFilterSegmentDelete)

	s.AddTool(mcp.Tool{
		Name:        "dt_filter_segments_entity_model",
		Description: `Get filter segments entity model (available fields and attributes).`,
		InputSchema: mcp.ToolInputSchema{
			Type:       "object",
			Properties: map[string]interface{}{},
		},
	}, h.handleFilterSegmentsEntityModel)

	s.AddTool(mcp.Tool{
		Name:        "dt_filter_segments_lean",
		Description: `Get filter segments in lean format (minimal data).`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"uids": map[string]interface{}{"type": "array", "description": "List of UIDs to fetch", "items": map[string]interface{}{"type": "string"}},
			},
		},
	}, h.handleFilterSegmentsLean)

	// ==================== Fieldsets ====================
	s.AddTool(mcp.Tool{
		Name:        "dt_fieldsets_list",
		Description: `List custom field definitions.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"table":     map[string]interface{}{"type": "string", "description": "Table name"},
				"page_size": map[string]interface{}{"type": "integer", "description": "Page size"},
			},
		},
	}, h.handleFieldsetsList)

	s.AddTool(mcp.Tool{
		Name:        "dt_fieldset_get",
		Description: `Get a fieldset by UID.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"uid": map[string]interface{}{"type": "string", "description": "Fieldset UID"},
			},
			Required: []string{"uid"},
		},
	}, h.handleFieldsetGet)

	s.AddTool(mcp.Tool{
		Name:        "dt_fieldset_create",
		Description: `Create a new fieldset.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"name":        map[string]interface{}{"type": "string", "description": "Fieldset name"},
				"table":       map[string]interface{}{"type": "string", "description": "Target table"},
				"description": map[string]interface{}{"type": "string", "description": "Description"},
				"fields":      map[string]interface{}{"type": "array", "description": "Field definitions"},
			},
			Required: []string{"name", "table", "fields"},
		},
	}, h.handleFieldsetCreate)

	s.AddTool(mcp.Tool{
		Name:        "dt_fieldset_update",
		Description: `Update a fieldset.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"uid":         map[string]interface{}{"type": "string", "description": "Fieldset UID"},
				"version":     map[string]interface{}{"type": "string", "description": "Current version"},
				"name":        map[string]interface{}{"type": "string", "description": "New name"},
				"description": map[string]interface{}{"type": "string", "description": "New description"},
				"fields":      map[string]interface{}{"type": "array", "description": "Updated field definitions"},
			},
			Required: []string{"uid", "version"},
		},
	}, h.handleFieldsetUpdate)

	s.AddTool(mcp.Tool{
		Name:        "dt_fieldset_delete",
		Description: `Delete a fieldset.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"uid":     map[string]interface{}{"type": "string", "description": "Fieldset UID"},
				"version": map[string]interface{}{"type": "string", "description": "Current version"},
			},
			Required: []string{"uid", "version"},
		},
	}, h.handleFieldsetDelete)

	// ==================== Record Deletion (GDPR) ====================
	s.AddTool(mcp.Tool{
		Name: "dt_record_deletion_execute",
		Description: `Execute a record deletion request (GDPR compliance).

Deletes records matching the specified criteria from Grail storage.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"table":  map[string]interface{}{"type": "string", "description": "Table name"},
				"filter": map[string]interface{}{"type": "string", "description": "DQL filter for records to delete"},
			},
			Required: []string{"table", "filter"},
		},
	}, h.handleRecordDeletionExecute)

	s.AddTool(mcp.Tool{
		Name:        "dt_record_deletion_status",
		Description: `Get status of a record deletion request.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"request_token": map[string]interface{}{"type": "string", "description": "Request token from execute"},
			},
			Required: []string{"request_token"},
		},
	}, h.handleRecordDeletionStatus)

	s.AddTool(mcp.Tool{
		Name:        "dt_record_deletion_cancel",
		Description: `Cancel a record deletion request.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"request_token": map[string]interface{}{"type": "string", "description": "Request token from execute"},
			},
			Required: []string{"request_token"},
		},
	}, h.handleRecordDeletionCancel)

	// ==================== Resource Store (Lookup Tables) ====================
	s.AddTool(mcp.Tool{
		Name: "dt_lookup_table_upload",
		Description: `Upload a lookup table for data enrichment.

Lookup tables can be used in DQL queries for data enrichment.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"name":    map[string]interface{}{"type": "string", "description": "Lookup table name"},
				"content": map[string]interface{}{"type": "string", "description": "CSV content"},
			},
			Required: []string{"name", "content"},
		},
	}, h.handleLookupTableUpload)

	s.AddTool(mcp.Tool{
		Name:        "dt_lookup_table_test_pattern",
		Description: `Test a lookup table pattern/query.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"name":    map[string]interface{}{"type": "string", "description": "Lookup table name"},
				"pattern": map[string]interface{}{"type": "string", "description": "Pattern to test"},
			},
			Required: []string{"name", "pattern"},
		},
	}, h.handleLookupTableTestPattern)

	s.AddTool(mcp.Tool{
		Name:        "dt_resource_files_delete",
		Description: `Delete resource files.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"files": map[string]interface{}{"type": "array", "description": "List of file paths to delete", "items": map[string]interface{}{"type": "string"}},
			},
			Required: []string{"files"},
		},
	}, h.handleResourceFilesDelete)
}

// ==================== Handler Implementations ====================

func (h *Handlers) handleBucketUpdate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	name := getStringParam(args, "name")
	version := getIntParam(args, "version", 0)
	if name == "" || version == 0 {
		return toolError(fmt.Errorf("name and version are required")), nil
	}

	body := map[string]interface{}{}
	if retention := getIntParam(args, "retention_days", 0); retention > 0 {
		body["retentionDays"] = retention
	}
	if displayName := getStringParam(args, "display_name"); displayName != "" {
		body["displayName"] = displayName
	}

	params := map[string]string{"optimistic-locking-version": fmt.Sprintf("%d", version)}
	resp, err := h.Client.Put(ctx, "/platform/storage/management/v1/bucket-definitions/"+name, body, client.WithQueryParams(params))
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

func (h *Handlers) handleBucketTruncate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name := getStringParam(req.Params.Arguments, "name")
	if name == "" {
		return toolError(fmt.Errorf("name is required")), nil
	}

	resp, err := h.Client.Post(ctx, "/platform/storage/management/v1/bucket-definitions/"+name+":truncate", nil)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(fmt.Sprintf("Bucket %s truncated", name)), nil
}

func (h *Handlers) handleFilterSegmentsList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	params := map[string]string{}
	if filter := getStringParam(args, "filter"); filter != "" {
		params["filter"] = filter
	}
	if pageSize := getIntParam(args, "page_size", 0); pageSize > 0 {
		params["page-size"] = fmt.Sprintf("%d", pageSize)
	}

	resp, err := h.Client.Get(ctx, "/platform/storage/filter-segments/v1/filter-segments", client.WithQueryParams(params))
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

func (h *Handlers) handleFilterSegmentGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	uid := getStringParam(args, "uid")
	if uid == "" {
		return toolError(fmt.Errorf("uid is required")), nil
	}

	params := map[string]string{}
	if getBoolParam(args, "include_filter") {
		params["include-filter"] = "true"
	}

	resp, err := h.Client.Get(ctx, "/platform/storage/filter-segments/v1/filter-segments/"+uid, client.WithQueryParams(params))
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

func (h *Handlers) handleFilterSegmentCreate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	name := getStringParam(args, "name")
	if name == "" {
		return toolError(fmt.Errorf("name is required")), nil
	}

	body := map[string]interface{}{
		"name": name,
	}
	if desc := getStringParam(args, "description"); desc != "" {
		body["description"] = desc
	}
	if isPublic, ok := args["is_public"].(bool); ok {
		body["isPublic"] = isPublic
	}
	if variables, ok := args["variables"].(map[string]interface{}); ok {
		body["variables"] = variables
	}
	if includes, ok := args["includes"].([]interface{}); ok {
		body["includes"] = includes
	}
	if excludes, ok := args["excludes"].([]interface{}); ok {
		body["excludes"] = excludes
	}

	resp, err := h.Client.Post(ctx, "/platform/storage/filter-segments/v1/filter-segments", body)
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

func (h *Handlers) handleFilterSegmentUpdate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	uid := getStringParam(args, "uid")
	version := getStringParam(args, "version")
	if uid == "" || version == "" {
		return toolError(fmt.Errorf("uid and version are required")), nil
	}

	body := map[string]interface{}{}
	if name := getStringParam(args, "name"); name != "" {
		body["name"] = name
	}
	if desc := getStringParam(args, "description"); desc != "" {
		body["description"] = desc
	}
	if isPublic, ok := args["is_public"].(bool); ok {
		body["isPublic"] = isPublic
	}
	if variables, ok := args["variables"].(map[string]interface{}); ok {
		body["variables"] = variables
	}
	if includes, ok := args["includes"].([]interface{}); ok {
		body["includes"] = includes
	}
	if excludes, ok := args["excludes"].([]interface{}); ok {
		body["excludes"] = excludes
	}

	params := map[string]string{"optimistic-locking-version": version}
	resp, err := h.Client.Put(ctx, "/platform/storage/filter-segments/v1/filter-segments/"+uid, body, client.WithQueryParams(params))
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

func (h *Handlers) handleFilterSegmentDelete(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	uid := getStringParam(args, "uid")
	version := getStringParam(args, "version")
	if uid == "" || version == "" {
		return toolError(fmt.Errorf("uid and version are required")), nil
	}

	params := map[string]string{"optimistic-locking-version": version}
	resp, err := h.Client.Delete(ctx, "/platform/storage/filter-segments/v1/filter-segments/"+uid, client.WithQueryParams(params))
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(fmt.Sprintf("Filter segment %s deleted", uid)), nil
}

func (h *Handlers) handleFilterSegmentsEntityModel(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	resp, err := h.Client.Get(ctx, "/platform/storage/filter-segments/v1/filter-segments-entity-model")
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

func (h *Handlers) handleFilterSegmentsLean(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	body := map[string]interface{}{}
	if uids, ok := args["uids"].([]interface{}); ok {
		body["uids"] = uids
	}

	resp, err := h.Client.Post(ctx, "/platform/storage/filter-segments/v1/filter-segments:lean", body)
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

func (h *Handlers) handleFieldsetsList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	params := map[string]string{}
	if table := getStringParam(args, "table"); table != "" {
		params["table"] = table
	}
	if pageSize := getIntParam(args, "page_size", 0); pageSize > 0 {
		params["page-size"] = fmt.Sprintf("%d", pageSize)
	}

	resp, err := h.Client.Get(ctx, "/platform/storage/fieldsets/v1/fieldsets", client.WithQueryParams(params))
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

func (h *Handlers) handleFieldsetGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	uid := getStringParam(req.Params.Arguments, "uid")
	if uid == "" {
		return toolError(fmt.Errorf("uid is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/storage/fieldsets/v1/fieldsets/"+uid)
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

func (h *Handlers) handleFieldsetCreate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	name := getStringParam(args, "name")
	table := getStringParam(args, "table")
	fields, ok := args["fields"].([]interface{})
	if name == "" || table == "" || !ok {
		return toolError(fmt.Errorf("name, table, and fields are required")), nil
	}

	body := map[string]interface{}{
		"name":   name,
		"table":  table,
		"fields": fields,
	}
	if desc := getStringParam(args, "description"); desc != "" {
		body["description"] = desc
	}

	resp, err := h.Client.Post(ctx, "/platform/storage/fieldsets/v1/fieldsets", body)
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

func (h *Handlers) handleFieldsetUpdate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	uid := getStringParam(args, "uid")
	version := getStringParam(args, "version")
	if uid == "" || version == "" {
		return toolError(fmt.Errorf("uid and version are required")), nil
	}

	body := map[string]interface{}{}
	if name := getStringParam(args, "name"); name != "" {
		body["name"] = name
	}
	if desc := getStringParam(args, "description"); desc != "" {
		body["description"] = desc
	}
	if fields, ok := args["fields"].([]interface{}); ok {
		body["fields"] = fields
	}

	params := map[string]string{"optimistic-locking-version": version}
	resp, err := h.Client.Put(ctx, "/platform/storage/fieldsets/v1/fieldsets/"+uid, body, client.WithQueryParams(params))
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

func (h *Handlers) handleFieldsetDelete(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	uid := getStringParam(args, "uid")
	version := getStringParam(args, "version")
	if uid == "" || version == "" {
		return toolError(fmt.Errorf("uid and version are required")), nil
	}

	params := map[string]string{"optimistic-locking-version": version}
	resp, err := h.Client.Delete(ctx, "/platform/storage/fieldsets/v1/fieldsets/"+uid, client.WithQueryParams(params))
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(fmt.Sprintf("Fieldset %s deleted", uid)), nil
}

func (h *Handlers) handleRecordDeletionExecute(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	table := getStringParam(args, "table")
	filter := getStringParam(args, "filter")
	if table == "" || filter == "" {
		return toolError(fmt.Errorf("table and filter are required")), nil
	}

	body := map[string]interface{}{
		"table":  table,
		"filter": filter,
	}

	resp, err := h.Client.Post(ctx, "/platform/storage/record-deletion/v1/delete:execute", body)
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

func (h *Handlers) handleRecordDeletionStatus(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	token := getStringParam(req.Params.Arguments, "request_token")
	if token == "" {
		return toolError(fmt.Errorf("request_token is required")), nil
	}

	params := map[string]string{"request-token": token}
	resp, err := h.Client.Get(ctx, "/platform/storage/record-deletion/v1/delete:status", client.WithQueryParams(params))
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

func (h *Handlers) handleRecordDeletionCancel(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	token := getStringParam(req.Params.Arguments, "request_token")
	if token == "" {
		return toolError(fmt.Errorf("request_token is required")), nil
	}

	params := map[string]string{"request-token": token}
	resp, err := h.Client.Post(ctx, "/platform/storage/record-deletion/v1/delete:cancel", nil, client.WithQueryParams(params))
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult("Record deletion cancelled"), nil
}

func (h *Handlers) handleLookupTableUpload(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	name := getStringParam(args, "name")
	content := getStringParam(args, "content")
	if name == "" || content == "" {
		return toolError(fmt.Errorf("name and content are required")), nil
	}

	body := map[string]interface{}{
		"name":    name,
		"content": content,
	}

	resp, err := h.Client.Post(ctx, "/platform/storage/resource-store/v1/files/tabular/lookup:upload", body)
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

func (h *Handlers) handleLookupTableTestPattern(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	name := getStringParam(args, "name")
	pattern := getStringParam(args, "pattern")
	if name == "" || pattern == "" {
		return toolError(fmt.Errorf("name and pattern are required")), nil
	}

	body := map[string]interface{}{
		"name":    name,
		"pattern": pattern,
	}

	resp, err := h.Client.Post(ctx, "/platform/storage/resource-store/v1/files/tabular/lookup:test-pattern", body)
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

func (h *Handlers) handleResourceFilesDelete(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	files, ok := req.Params.Arguments["files"].([]interface{})
	if !ok || len(files) == 0 {
		return toolError(fmt.Errorf("files array is required")), nil
	}

	body := map[string]interface{}{
		"files": files,
	}

	resp, err := h.Client.Post(ctx, "/platform/storage/resource-store/v1/files:delete", body)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(fmt.Sprintf("Deleted %d files", len(files))), nil
}
