package tools

import (
	"context"
	"fmt"

	"github.com/dynatrace/dynatrace-platform-mcp-server/internal/client"
	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

// RegisterDocumentsExtendedTools registers additional document tools not in the base set.
func RegisterDocumentsExtendedTools(s *mcpserver.MCPServer, h *Handlers) {
	// ==================== Document Metadata & Content ====================
	s.AddTool(mcp.Tool{
		Name:        "dt_document_metadata_get",
		Description: `Get document metadata without content.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{"type": "string", "description": "Document ID"},
			},
			Required: []string{"id"},
		},
	}, h.handleDocumentMetadataGet)

	s.AddTool(mcp.Tool{
		Name:        "dt_document_content_get",
		Description: `Get document content only.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{"type": "string", "description": "Document ID"},
			},
			Required: []string{"id"},
		},
	}, h.handleDocumentContentGet)

	// ==================== Document Snapshots ====================
	s.AddTool(mcp.Tool{
		Name:        "dt_document_snapshots_list",
		Description: `List snapshots (versions) for a document.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{"type": "string", "description": "Document ID"},
			},
			Required: []string{"id"},
		},
	}, h.handleDocumentSnapshotsList)

	s.AddTool(mcp.Tool{
		Name:        "dt_document_snapshot_get",
		Description: `Get a specific snapshot version of a document.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id":               map[string]interface{}{"type": "string", "description": "Document ID"},
				"snapshot_version": map[string]interface{}{"type": "string", "description": "Snapshot version"},
			},
			Required: []string{"id", "snapshot_version"},
		},
	}, h.handleDocumentSnapshotGet)

	s.AddTool(mcp.Tool{
		Name:        "dt_document_snapshot_restore",
		Description: `Restore a document to a specific snapshot version.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id":               map[string]interface{}{"type": "string", "description": "Document ID"},
				"snapshot_version": map[string]interface{}{"type": "string", "description": "Snapshot version to restore"},
			},
			Required: []string{"id", "snapshot_version"},
		},
	}, h.handleDocumentSnapshotRestore)

	// ==================== Document Locking ====================
	s.AddTool(mcp.Tool{
		Name:        "dt_document_lock_inspect",
		Description: `Inspect lock status of a document.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{"type": "string", "description": "Document ID"},
			},
			Required: []string{"id"},
		},
	}, h.handleDocumentLockInspect)

	s.AddTool(mcp.Tool{
		Name:        "dt_document_lock_acquire",
		Description: `Acquire a lock on a document.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id":      map[string]interface{}{"type": "string", "description": "Document ID"},
				"version": map[string]interface{}{"type": "string", "description": "Document version"},
			},
			Required: []string{"id", "version"},
		},
	}, h.handleDocumentLockAcquire)

	s.AddTool(mcp.Tool{
		Name:        "dt_document_lock_release",
		Description: `Release a lock on a document.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{"type": "string", "description": "Document ID"},
			},
			Required: []string{"id"},
		},
	}, h.handleDocumentLockRelease)

	// ==================== Document Ownership ====================
	s.AddTool(mcp.Tool{
		Name:        "dt_document_transfer_owner",
		Description: `Transfer ownership of a document to another user.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id":        map[string]interface{}{"type": "string", "description": "Document ID"},
				"new_owner": map[string]interface{}{"type": "string", "description": "New owner user ID"},
			},
			Required: []string{"id", "new_owner"},
		},
	}, h.handleDocumentTransferOwner)

	// ==================== Environment Shares ====================
	s.AddTool(mcp.Tool{
		Name:        "dt_environment_shares_list",
		Description: `List environment-wide document shares.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"filter":    map[string]interface{}{"type": "string", "description": "Filter query"},
				"page_size": map[string]interface{}{"type": "integer", "description": "Page size"},
			},
		},
	}, h.handleEnvironmentSharesList)

	s.AddTool(mcp.Tool{
		Name:        "dt_environment_share_create",
		Description: `Create an environment-wide share for a document.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"document_id": map[string]interface{}{"type": "string", "description": "Document ID to share"},
				"access":      map[string]interface{}{"type": "string", "description": "Access level (read, write)"},
				"claim_type":  map[string]interface{}{"type": "string", "description": "Claim type"},
			},
			Required: []string{"document_id", "access"},
		},
	}, h.handleEnvironmentShareCreate)

	s.AddTool(mcp.Tool{
		Name:        "dt_environment_share_get",
		Description: `Get an environment share by ID.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{"type": "string", "description": "Share ID"},
			},
			Required: []string{"id"},
		},
	}, h.handleEnvironmentShareGet)

	s.AddTool(mcp.Tool{
		Name:        "dt_environment_share_update",
		Description: `Update an environment share.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id":      map[string]interface{}{"type": "string", "description": "Share ID"},
				"access":  map[string]interface{}{"type": "string", "description": "Access level"},
				"enabled": map[string]interface{}{"type": "boolean", "description": "Enable/disable share"},
			},
			Required: []string{"id"},
		},
	}, h.handleEnvironmentShareUpdate)

	s.AddTool(mcp.Tool{
		Name:        "dt_environment_share_delete",
		Description: `Delete an environment share.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{"type": "string", "description": "Share ID"},
			},
			Required: []string{"id"},
		},
	}, h.handleEnvironmentShareDelete)

	s.AddTool(mcp.Tool{
		Name:        "dt_environment_share_claim",
		Description: `Claim a shared document.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{"type": "string", "description": "Share ID"},
			},
			Required: []string{"id"},
		},
	}, h.handleEnvironmentShareClaim)

	s.AddTool(mcp.Tool{
		Name:        "dt_environment_share_claimers_list",
		Description: `List users who have claimed a shared document.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{"type": "string", "description": "Share ID"},
			},
			Required: []string{"id"},
		},
	}, h.handleEnvironmentShareClaimersList)

	// ==================== Direct Shares ====================
	s.AddTool(mcp.Tool{
		Name:        "dt_direct_shares_list",
		Description: `List direct document shares.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"filter":    map[string]interface{}{"type": "string", "description": "Filter query"},
				"page_size": map[string]interface{}{"type": "integer", "description": "Page size"},
			},
		},
	}, h.handleDirectSharesList)

	s.AddTool(mcp.Tool{
		Name:        "dt_direct_share_create",
		Description: `Create a direct share for a document.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"document_id": map[string]interface{}{"type": "string", "description": "Document ID to share"},
				"recipients":  map[string]interface{}{"type": "array", "description": "List of recipient user IDs"},
				"access":      map[string]interface{}{"type": "string", "description": "Access level (read, write)"},
			},
			Required: []string{"document_id", "recipients", "access"},
		},
	}, h.handleDirectShareCreate)

	s.AddTool(mcp.Tool{
		Name:        "dt_direct_share_get",
		Description: `Get a direct share by ID.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{"type": "string", "description": "Share ID"},
			},
			Required: []string{"id"},
		},
	}, h.handleDirectShareGet)

	s.AddTool(mcp.Tool{
		Name:        "dt_direct_share_update",
		Description: `Update a direct share.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id":     map[string]interface{}{"type": "string", "description": "Share ID"},
				"access": map[string]interface{}{"type": "string", "description": "Access level"},
			},
			Required: []string{"id"},
		},
	}, h.handleDirectShareUpdate)

	s.AddTool(mcp.Tool{
		Name:        "dt_direct_share_delete",
		Description: `Delete a direct share.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{"type": "string", "description": "Share ID"},
			},
			Required: []string{"id"},
		},
	}, h.handleDirectShareDelete)

	s.AddTool(mcp.Tool{
		Name:        "dt_direct_share_recipients_list",
		Description: `List recipients of a direct share.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{"type": "string", "description": "Share ID"},
			},
			Required: []string{"id"},
		},
	}, h.handleDirectShareRecipientsList)

	s.AddTool(mcp.Tool{
		Name:        "dt_direct_share_recipients_add",
		Description: `Add recipients to a direct share.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id":         map[string]interface{}{"type": "string", "description": "Share ID"},
				"recipients": map[string]interface{}{"type": "array", "description": "Recipients to add"},
			},
			Required: []string{"id", "recipients"},
		},
	}, h.handleDirectShareRecipientsAdd)

	s.AddTool(mcp.Tool{
		Name:        "dt_direct_share_recipients_remove",
		Description: `Remove recipients from a direct share.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id":         map[string]interface{}{"type": "string", "description": "Share ID"},
				"recipients": map[string]interface{}{"type": "array", "description": "Recipients to remove"},
			},
			Required: []string{"id", "recipients"},
		},
	}, h.handleDirectShareRecipientsRemove)

	// ==================== Trash ====================
	s.AddTool(mcp.Tool{
		Name:        "dt_trash_documents_list",
		Description: `List documents in trash.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"filter":    map[string]interface{}{"type": "string", "description": "Filter query"},
				"page_size": map[string]interface{}{"type": "integer", "description": "Page size"},
			},
		},
	}, h.handleTrashDocumentsList)

	s.AddTool(mcp.Tool{
		Name:        "dt_trash_document_get",
		Description: `Get a document from trash.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{"type": "string", "description": "Document ID"},
			},
			Required: []string{"id"},
		},
	}, h.handleTrashDocumentGet)

	s.AddTool(mcp.Tool{
		Name:        "dt_trash_document_delete",
		Description: `Permanently delete a document from trash.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{"type": "string", "description": "Document ID"},
			},
			Required: []string{"id"},
		},
	}, h.handleTrashDocumentDelete)

	s.AddTool(mcp.Tool{
		Name:        "dt_trash_document_restore",
		Description: `Restore a document from trash.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"id": map[string]interface{}{"type": "string", "description": "Document ID"},
			},
			Required: []string{"id"},
		},
	}, h.handleTrashDocumentRestore)

	// ==================== Bulk Operations ====================
	s.AddTool(mcp.Tool{
		Name:        "dt_documents_bulk_delete",
		Description: `Bulk delete documents.`,
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"ids": map[string]interface{}{"type": "array", "description": "List of document IDs to delete", "items": map[string]interface{}{"type": "string"}},
			},
			Required: []string{"ids"},
		},
	}, h.handleDocumentsBulkDelete)
}

// ==================== Handler Implementations ====================

func (h *Handlers) handleDocumentMetadataGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/document/v1/documents/"+id+"/metadata")
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

func (h *Handlers) handleDocumentContentGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/document/v1/documents/"+id+"/content")
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

func (h *Handlers) handleDocumentSnapshotsList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/document/v1/documents/"+id+"/snapshots")
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

func (h *Handlers) handleDocumentSnapshotGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	snapshotVersion := getStringParam(args, "snapshot_version")
	if id == "" || snapshotVersion == "" {
		return toolError(fmt.Errorf("id and snapshot_version are required")), nil
	}

	resp, err := h.Client.Get(ctx, fmt.Sprintf("/platform/document/v1/documents/%s/snapshots/%s", id, snapshotVersion))
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

func (h *Handlers) handleDocumentSnapshotRestore(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	snapshotVersion := getStringParam(args, "snapshot_version")
	if id == "" || snapshotVersion == "" {
		return toolError(fmt.Errorf("id and snapshot_version are required")), nil
	}

	resp, err := h.Client.Post(ctx, fmt.Sprintf("/platform/document/v1/documents/%s/snapshots/%s:restore", id, snapshotVersion), nil)
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

func (h *Handlers) handleDocumentLockInspect(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/document/v1/documents/"+id+":inspect-lock")
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

func (h *Handlers) handleDocumentLockAcquire(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	version := getStringParam(args, "version")
	if id == "" || version == "" {
		return toolError(fmt.Errorf("id and version are required")), nil
	}

	params := map[string]string{"optimistic-locking-version": version}
	resp, err := h.Client.Post(ctx, "/platform/document/v1/documents/"+id+":acquire-lock", nil, client.WithQueryParams(params))
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

func (h *Handlers) handleDocumentLockRelease(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Post(ctx, "/platform/document/v1/documents/"+id+":release-lock", nil)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(fmt.Sprintf("Lock released for document %s", id)), nil
}

func (h *Handlers) handleDocumentTransferOwner(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	newOwner := getStringParam(args, "new_owner")
	if id == "" || newOwner == "" {
		return toolError(fmt.Errorf("id and new_owner are required")), nil
	}

	body := map[string]interface{}{
		"newOwner": newOwner,
	}

	resp, err := h.Client.Post(ctx, "/platform/document/v1/documents/"+id+":transfer-owner", body)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(fmt.Sprintf("Ownership transferred to %s", newOwner)), nil
}

func (h *Handlers) handleEnvironmentSharesList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	params := map[string]string{}
	if filter := getStringParam(args, "filter"); filter != "" {
		params["filter"] = filter
	}
	if pageSize := getIntParam(args, "page_size", 0); pageSize > 0 {
		params["page-size"] = fmt.Sprintf("%d", pageSize)
	}

	resp, err := h.Client.Get(ctx, "/platform/document/v1/environment-shares", client.WithQueryParams(params))
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

func (h *Handlers) handleEnvironmentShareCreate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	documentID := getStringParam(args, "document_id")
	access := getStringParam(args, "access")
	if documentID == "" || access == "" {
		return toolError(fmt.Errorf("document_id and access are required")), nil
	}

	body := map[string]interface{}{
		"documentId": documentID,
		"access":     access,
	}
	if claimType := getStringParam(args, "claim_type"); claimType != "" {
		body["claimType"] = claimType
	}

	resp, err := h.Client.Post(ctx, "/platform/document/v1/environment-shares", body)
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

func (h *Handlers) handleEnvironmentShareGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/document/v1/environment-shares/"+id)
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

func (h *Handlers) handleEnvironmentShareUpdate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	body := map[string]interface{}{}
	if access := getStringParam(args, "access"); access != "" {
		body["access"] = access
	}
	if enabled, ok := args["enabled"].(bool); ok {
		body["enabled"] = enabled
	}

	resp, err := h.Client.Put(ctx, "/platform/document/v1/environment-shares/"+id, body)
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

func (h *Handlers) handleEnvironmentShareDelete(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Delete(ctx, "/platform/document/v1/environment-shares/"+id)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(fmt.Sprintf("Environment share %s deleted", id)), nil
}

func (h *Handlers) handleEnvironmentShareClaim(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Post(ctx, "/platform/document/v1/environment-shares/"+id+"/claim", nil)
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

func (h *Handlers) handleEnvironmentShareClaimersList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/document/v1/environment-shares/"+id+"/claimers")
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

func (h *Handlers) handleDirectSharesList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	params := map[string]string{}
	if filter := getStringParam(args, "filter"); filter != "" {
		params["filter"] = filter
	}
	if pageSize := getIntParam(args, "page_size", 0); pageSize > 0 {
		params["page-size"] = fmt.Sprintf("%d", pageSize)
	}

	resp, err := h.Client.Get(ctx, "/platform/document/v1/direct-shares", client.WithQueryParams(params))
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

func (h *Handlers) handleDirectShareCreate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	documentID := getStringParam(args, "document_id")
	access := getStringParam(args, "access")
	recipients, ok := args["recipients"].([]interface{})
	if documentID == "" || access == "" || !ok {
		return toolError(fmt.Errorf("document_id, recipients, and access are required")), nil
	}

	body := map[string]interface{}{
		"documentId": documentID,
		"access":     access,
		"recipients": recipients,
	}

	resp, err := h.Client.Post(ctx, "/platform/document/v1/direct-shares", body)
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

func (h *Handlers) handleDirectShareGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/document/v1/direct-shares/"+id)
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

func (h *Handlers) handleDirectShareUpdate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	body := map[string]interface{}{}
	if access := getStringParam(args, "access"); access != "" {
		body["access"] = access
	}

	resp, err := h.Client.Put(ctx, "/platform/document/v1/direct-shares/"+id, body)
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

func (h *Handlers) handleDirectShareDelete(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Delete(ctx, "/platform/document/v1/direct-shares/"+id)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(fmt.Sprintf("Direct share %s deleted", id)), nil
}

func (h *Handlers) handleDirectShareRecipientsList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/document/v1/direct-shares/"+id+"/recipients")
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

func (h *Handlers) handleDirectShareRecipientsAdd(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	recipients, ok := args["recipients"].([]interface{})
	if id == "" || !ok {
		return toolError(fmt.Errorf("id and recipients are required")), nil
	}

	body := map[string]interface{}{
		"recipients": recipients,
	}

	resp, err := h.Client.Post(ctx, "/platform/document/v1/direct-shares/"+id+"/recipients/add", body)
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

func (h *Handlers) handleDirectShareRecipientsRemove(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	recipients, ok := args["recipients"].([]interface{})
	if id == "" || !ok {
		return toolError(fmt.Errorf("id and recipients are required")), nil
	}

	body := map[string]interface{}{
		"recipients": recipients,
	}

	resp, err := h.Client.Post(ctx, "/platform/document/v1/direct-shares/"+id+"/recipients/remove", body)
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

func (h *Handlers) handleTrashDocumentsList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	params := map[string]string{}
	if filter := getStringParam(args, "filter"); filter != "" {
		params["filter"] = filter
	}
	if pageSize := getIntParam(args, "page_size", 0); pageSize > 0 {
		params["page-size"] = fmt.Sprintf("%d", pageSize)
	}

	resp, err := h.Client.Get(ctx, "/platform/document/v1/trash/documents", client.WithQueryParams(params))
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

func (h *Handlers) handleTrashDocumentGet(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/document/v1/trash/documents/"+id)
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

func (h *Handlers) handleTrashDocumentDelete(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Delete(ctx, "/platform/document/v1/trash/documents/"+id)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(fmt.Sprintf("Document %s permanently deleted", id)), nil
}

func (h *Handlers) handleTrashDocumentRestore(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Post(ctx, "/platform/document/v1/trash/documents/"+id+"/restore", nil)
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

func (h *Handlers) handleDocumentsBulkDelete(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ids, ok := req.Params.Arguments["ids"].([]interface{})
	if !ok || len(ids) == 0 {
		return toolError(fmt.Errorf("ids array is required")), nil
	}

	body := map[string]interface{}{
		"ids": ids,
	}

	resp, err := h.Client.Post(ctx, "/platform/document/v1/documents:delete", body)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(fmt.Sprintf("Deleted %d documents", len(ids))), nil
}
