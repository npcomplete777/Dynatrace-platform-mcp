// Package tools provides MCP tool handlers for Dynatrace Platform APIs.
package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/dynatrace/dynatrace-platform-mcp-server/internal/client"
	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

const (
	sloBasePath         = "/platform/slo/v1/slos"
	sloTemplatePath     = "/platform/slo/v1/objective-templates"
	maxRetryAttempts    = 2
	evalPollInterval    = 1 * time.Second
	evalMaxPollTime     = 60 * time.Second
	evalDefaultTimeout  = 1000 // milliseconds for initial request
)

// SLOResponse represents the structure of an SLO response from Dynatrace API
type SLOResponse struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Version     string                 `json:"version"`
	CustomSli   map[string]interface{} `json:"customSli,omitempty"`
	SliRef      map[string]interface{} `json:"sliReference,omitempty"`
	Criteria    []interface{}          `json:"criteria"`
	Tags        []string               `json:"tags,omitempty"`
}

// SLOEvaluationStartResponse handles both sync and async responses
type SLOEvaluationStartResponse struct {
	// Async response - contains token for polling
	EvaluationToken string `json:"evaluationToken,omitempty"`
	// Sync response - contains immediate results
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
	EvaluationResults []SLOEvaluationResult  `json:"evaluationResults,omitempty"`
}

// SLOEvaluationResult represents a single evaluation result
type SLOEvaluationResult struct {
	Status      string  `json:"status"`
	Value       float64 `json:"value"`
	ErrorBudget float64 `json:"errorBudget,omitempty"`
	Target      float64 `json:"target,omitempty"`
	Warning     float64 `json:"warning,omitempty"`
}

// SLOPollResponse represents the poll endpoint response
type SLOPollResponse struct {
	Status   string                 `json:"status"`
	Progress int                    `json:"progress,omitempty"`
	Result   *SLOEvaluationResult   `json:"result,omitempty"`
	Results  map[string]interface{} `json:"results,omitempty"`
}

// RegisterSLOTools registers all SLO-related MCP tools with optimistic locking support.
func RegisterSLOToolsV2(s *mcpserver.MCPServer, h *Handlers, isEnabled func(string) bool) {
	// List SLOs
	if isEnabled("dt_slos_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_slos_list",
			Description: "List Service Level Objectives (SLOs) with pagination support. Returns SLO definitions including their IDs and versions for subsequent operations.",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"page_size": map[string]interface{}{
						"type":        "integer",
						"description": "Number of SLOs per page (max 100, default 100)",
					},
				},
			},
		}, h.handleSLOsListV2)
	}

	// Get SLO
	if isEnabled("dt_slo_get") {
		s.AddTool(mcp.Tool{
			Name:        "dt_slo_get",
			Description: "Get details of a specific SLO by ID. Returns the full SLO definition including the current version hash required for updates/deletes.",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{
						"type":        "string",
						"description": "SLO ID",
					},
				},
				Required: []string{"id"},
			},
		}, h.handleSLOGetV2)
	}

	// Create SLO
	if isEnabled("dt_slo_create") {
		s.AddTool(mcp.Tool{
			Name:        "dt_slo_create",
			Description: "Create a new SLO. Use either customSli (for DQL-based indicators) OR sliReference (for built-in templates) - not both. Returns the created SLO with its ID and version hash.",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"name": map[string]interface{}{
						"type":        "string",
						"description": "SLO name (required)",
					},
					"description": map[string]interface{}{
						"type":        "string",
						"description": "SLO description",
					},
					"customSli": map[string]interface{}{
						"type":        "object",
						"description": "Custom SLI using DQL query. Mutually exclusive with sliReference.",
						"properties": map[string]interface{}{
							"indicator": map[string]interface{}{
								"type":        "string",
								"description": "DQL query that outputs a 'sli' field with value 0-100. Example: 'timeseries sli=avg(dt.host.cpu.idle)'",
							},
						},
					},
					"sliReference": map[string]interface{}{
						"type":        "object",
						"description": "SLI template reference. Mutually exclusive with customSli.",
						"properties": map[string]interface{}{
							"templateId": map[string]interface{}{
								"type":        "string",
								"description": "SLI template ID from dt_slo_templates_list",
							},
							"variables": map[string]interface{}{
								"type":        "array",
								"description": "Template variables (name/value pairs)",
								"items": map[string]interface{}{
									"type": "object",
									"properties": map[string]interface{}{
										"name":  map[string]interface{}{"type": "string"},
										"value": map[string]interface{}{"type": "string"},
									},
								},
							},
						},
					},
					"criteria": map[string]interface{}{
						"type":        "array",
						"description": "SLO criteria (required)",
						"items": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"timeframeFrom": map[string]interface{}{
									"type":        "string",
									"description": "Start of timeframe (e.g., 'now-7d')",
								},
								"timeframeTo": map[string]interface{}{
									"type":        "string",
									"description": "End of timeframe (e.g., 'now')",
								},
								"target": map[string]interface{}{
									"type":        "number",
									"description": "Target percentage (0-100)",
								},
								"warning": map[string]interface{}{
									"type":        "number",
									"description": "Warning threshold percentage (0-100)",
								},
							},
						},
					},
					"tags": map[string]interface{}{
						"type":        "array",
						"description": "Tags for the SLO (e.g., 'Stage:DEV')",
						"items":       map[string]interface{}{"type": "string"},
					},
				},
				Required: []string{"name", "criteria"},
			},
		}, h.handleSLOCreateV2)
	}

	// Update SLO - atomic fetch-mutate pattern
	if isEnabled("dt_slo_update") {
		s.AddTool(mcp.Tool{
			Name:        "dt_slo_update",
			Description: "Update an existing SLO. This tool implements atomic fetch-mutate: it fetches the current version, applies changes, and handles 409 conflicts automatically with retry.",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{
						"type":        "string",
						"description": "SLO ID (required)",
					},
					"name": map[string]interface{}{
						"type":        "string",
						"description": "Updated SLO name",
					},
					"description": map[string]interface{}{
						"type":        "string",
						"description": "Updated SLO description",
					},
					"customSli": map[string]interface{}{
						"type":        "object",
						"description": "Updated custom SLI (mutually exclusive with sliReference)",
						"properties": map[string]interface{}{
							"indicator": map[string]interface{}{
								"type":        "string",
								"description": "DQL query that outputs a 'sli' field with value 0-100",
							},
						},
					},
					"sliReference": map[string]interface{}{
						"type":        "object",
						"description": "Updated SLI template reference (mutually exclusive with customSli)",
					},
					"criteria": map[string]interface{}{
						"type":        "array",
						"description": "Updated SLO criteria",
					},
					"tags": map[string]interface{}{
						"type":        "array",
						"description": "Updated tags",
						"items":       map[string]interface{}{"type": "string"},
					},
				},
				Required: []string{"id"},
			},
		}, h.handleSLOUpdateV2)
	}

	// Delete SLO
	if isEnabled("dt_slo_delete") {
		s.AddTool(mcp.Tool{
			Name:        "dt_slo_delete",
			Description: "Delete an SLO by ID. This tool fetches the current version before deletion to satisfy optimistic locking requirements.",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{
						"type":        "string",
						"description": "SLO ID (required)",
					},
				},
				Required: []string{"id"},
			},
		}, h.handleSLODeleteV2)
	}

	// List SLO Templates
	if isEnabled("dt_slo_templates_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_slo_templates_list",
			Description: "List built-in SLO objective templates. Templates provide pre-defined DQL indicators for common use cases like Service Availability, Service Performance, Host CPU, and Kubernetes Efficiency. Use template IDs with dt_slo_create's sliReference parameter instead of writing custom DQL.",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"page_size": map[string]interface{}{
						"type":        "integer",
						"description": "Number of templates per page (max 400)",
					},
				},
			},
		}, h.handleSLOTemplatesListV2)
	}

	// Get SLO Template
	if isEnabled("dt_slo_template_get") {
		s.AddTool(mcp.Tool{
			Name:        "dt_slo_template_get",
			Description: "Get details of a specific SLO objective template by ID, including required variables. Use this to discover what variables (e.g., services, responseTimeInMilliSeconds) are needed before creating an SLO with sliReference.",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{
						"type":        "string",
						"description": "Template ID (Base64 encoded, from dt_slo_templates_list)",
					},
				},
				Required: []string{"id"},
			},
		}, h.handleSLOTemplateGetV2)
	}

	// ==================== SLO Evaluation Tools ====================

	// Evaluate SLO (high-level, handles async automatically)
	if isEnabled("dt_slo_evaluate") {
		s.AddTool(mcp.Tool{
			Name:        "dt_slo_evaluate",
			Description: "Evaluate an SLO and get its current status. This tool handles the async evaluation pattern automatically: it starts the evaluation, polls for completion if needed, and returns the final result. Use this for a simple one-call evaluation.",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{
						"type":        "string",
						"description": "SLO ID to evaluate (required)",
					},
					"timeframe_from": map[string]interface{}{
						"type":        "string",
						"description": "Start of custom timeframe (e.g., 'now-30d'). If not provided, uses SLO's default criteria timeframe.",
					},
					"timeframe_to": map[string]interface{}{
						"type":        "string",
						"description": "End of custom timeframe (e.g., 'now'). If not provided, uses SLO's default criteria timeframe.",
					},
				},
				Required: []string{"id"},
			},
		}, h.handleSLOEvaluate)
	}

	// Start Evaluation (low-level)
	if isEnabled("dt_slo_evaluation_start") {
		s.AddTool(mcp.Tool{
			Name:        "dt_slo_evaluation_start",
			Description: "Start an SLO evaluation. May return results immediately (sync) or an evaluation token for polling (async). For most use cases, prefer dt_slo_evaluate which handles both scenarios automatically.",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{
						"type":        "string",
						"description": "SLO ID to evaluate (required)",
					},
					"timeframe_from": map[string]interface{}{
						"type":        "string",
						"description": "Start of custom timeframe (e.g., 'now-30d')",
					},
					"timeframe_to": map[string]interface{}{
						"type":        "string",
						"description": "End of custom timeframe (e.g., 'now')",
					},
					"timeout_ms": map[string]interface{}{
						"type":        "integer",
						"description": "Request timeout in milliseconds. Lower values increase chance of async response. Default: 1000",
					},
				},
				Required: []string{"id"},
			},
		}, h.handleSLOEvaluationStart)
	}

	// Poll Evaluation (low-level)
	if isEnabled("dt_slo_evaluation_poll") {
		s.AddTool(mcp.Tool{
			Name:        "dt_slo_evaluation_poll",
			Description: "Poll the status of an async SLO evaluation. Returns IN_PROGRESS with progress percentage, or COMPLETED with results.",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"evaluation_token": map[string]interface{}{
						"type":        "string",
						"description": "Evaluation token from dt_slo_evaluation_start (required)",
					},
				},
				Required: []string{"evaluation_token"},
			},
		}, h.handleSLOEvaluationPoll)
	}

	// Cancel Evaluation (low-level)
	if isEnabled("dt_slo_evaluation_cancel") {
		s.AddTool(mcp.Tool{
			Name:        "dt_slo_evaluation_cancel",
			Description: "Cancel a running SLO evaluation. Use this to abort long-running evaluations.",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"evaluation_token": map[string]interface{}{
						"type":        "string",
						"description": "Evaluation token from dt_slo_evaluation_start (required)",
					},
				},
				Required: []string{"evaluation_token"},
			},
		}, h.handleSLOEvaluationCancel)
	}
}

// fetchSLO retrieves an SLO and returns the parsed response
func (h *Handlers) fetchSLO(ctx context.Context, id string) (*SLOResponse, error) {
	resp, err := h.Client.Get(ctx, sloBasePath+"/"+id)
	if err != nil {
		return nil, fmt.Errorf("fetch SLO: %w", err)
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("fetch SLO failed: %s", client.FormatError(resp))
	}

	var slo SLOResponse
	if err := json.Unmarshal(resp.Body, &slo); err != nil {
		return nil, fmt.Errorf("parse SLO response: %w", err)
	}
	return &slo, nil
}

// scrubPayload ensures mutual exclusivity between customSli and sliReference
func scrubPayload(payload map[string]interface{}, current *SLOResponse) map[string]interface{} {
	cleaned := make(map[string]interface{})

	// Copy all fields
	for k, v := range payload {
		cleaned[k] = v
	}

	// Ensure mutual exclusivity: if customSli is set, remove sliReference and vice versa
	_, hasCustomSli := cleaned["customSli"]
	_, hasSliRef := cleaned["sliReference"]

	if hasCustomSli && hasSliRef {
		// Prefer customSli if both are provided
		delete(cleaned, "sliReference")
	}

	// If neither is set but current has one, use that
	if !hasCustomSli && !hasSliRef && current != nil {
		if current.CustomSli != nil && len(current.CustomSli) > 0 {
			cleaned["customSli"] = current.CustomSli
		} else if current.SliRef != nil && len(current.SliRef) > 0 {
			cleaned["sliReference"] = current.SliRef
		}
	}

	// Never send null/empty sliReference if customSli is present
	if hasCustomSli {
		if ref, exists := cleaned["sliReference"]; exists {
			if ref == nil {
				delete(cleaned, "sliReference")
			}
		}
	}

	// Never send null/empty customSli if sliReference is present
	if hasSliRef {
		if csi, exists := cleaned["customSli"]; exists {
			if csi == nil {
				delete(cleaned, "customSli")
			}
		}
	}

	return cleaned
}

// handleSLOsListV2 lists SLOs with pagination
func (h *Handlers) handleSLOsListV2(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	params := make(map[string]string)
	if pageSize := getIntParam(req.Params.Arguments, "page_size", 0); pageSize > 0 {
		params["page-size"] = fmt.Sprintf("%d", pageSize)
	}

	resp, err := h.Client.Get(ctx, sloBasePath, client.WithQueryParams(params))
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

// handleSLOGetV2 retrieves a specific SLO
func (h *Handlers) handleSLOGetV2(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, sloBasePath+"/"+id)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

// handleSLOCreateV2 creates a new SLO with proper payload validation
func (h *Handlers) handleSLOCreateV2(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	payload := make(map[string]interface{})

	// Extract and validate required fields
	name := getStringParam(req.Params.Arguments, "name")
	if name == "" {
		return toolError(fmt.Errorf("name is required")), nil
	}
	payload["name"] = name

	criteria := getSliceParam(req.Params.Arguments, "criteria")
	if criteria == nil || len(criteria) == 0 {
		return toolError(fmt.Errorf("criteria is required and must not be empty")), nil
	}
	payload["criteria"] = criteria

	// Optional description
	if desc := getStringParam(req.Params.Arguments, "description"); desc != "" {
		payload["description"] = desc
	}

	// Handle SLI - must have either customSli OR sliReference (not both)
	customSli := getMapParam(req.Params.Arguments, "customSli")
	sliRef := getMapParam(req.Params.Arguments, "sliReference")

	if customSli != nil && sliRef != nil {
		return toolError(fmt.Errorf("cannot specify both customSli and sliReference - they are mutually exclusive")), nil
	}

	if customSli != nil {
		payload["customSli"] = customSli
	} else if sliRef != nil {
		payload["sliReference"] = sliRef
	} else {
		return toolError(fmt.Errorf("either customSli or sliReference is required")), nil
	}

	// Optional tags
	if tags := getStringSliceParam(req.Params.Arguments, "tags"); tags != nil {
		payload["tags"] = tags
	} else {
		payload["tags"] = []string{}
	}

	// Create the SLO
	resp, err := h.Client.Post(ctx, sloBasePath, payload)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}

	// Return the created SLO (includes id and version for subsequent operations)
	return toolResult(resp.Body), nil
}

// handleSLOUpdateV2 implements atomic fetch-mutate with conflict retry
func (h *Handlers) handleSLOUpdateV2(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	// Atomic fetch-mutate loop with retry
	for attempt := 0; attempt < maxRetryAttempts; attempt++ {
		// Step 1: Fetch current SLO to get version
		current, err := h.fetchSLO(ctx, id)
		if err != nil {
			return toolError(err), nil
		}

		// Step 2: Build updated payload, merging with current values
		payload := make(map[string]interface{})

		// Start with current values
		payload["name"] = current.Name
		if current.Description != "" {
			payload["description"] = current.Description
		}
		if current.CustomSli != nil {
			payload["customSli"] = current.CustomSli
		}
		if current.SliRef != nil {
			payload["sliReference"] = current.SliRef
		}
		payload["criteria"] = current.Criteria
		payload["tags"] = current.Tags

		// Override with provided values
		if name := getStringParam(req.Params.Arguments, "name"); name != "" {
			payload["name"] = name
		}
		if desc := getStringParam(req.Params.Arguments, "description"); desc != "" {
			payload["description"] = desc
		}
		if customSli := getMapParam(req.Params.Arguments, "customSli"); customSli != nil {
			payload["customSli"] = customSli
			delete(payload, "sliReference") // Mutually exclusive
		}
		if sliRef := getMapParam(req.Params.Arguments, "sliReference"); sliRef != nil {
			payload["sliReference"] = sliRef
			delete(payload, "customSli") // Mutually exclusive
		}
		if criteria := getSliceParam(req.Params.Arguments, "criteria"); criteria != nil {
			payload["criteria"] = criteria
		}
		if tags := getStringSliceParam(req.Params.Arguments, "tags"); tags != nil {
			payload["tags"] = tags
		}

		// Scrub payload for API compatibility
		payload = scrubPayload(payload, current)

		// Step 3: PUT with optimistic-locking-version
		resp, err := h.Client.Put(
			ctx,
			sloBasePath+"/"+id,
			payload,
			client.WithQueryParams(map[string]string{
				"optimistic-locking-version": current.Version,
			}),
		)
		if err != nil {
			return toolError(err), nil
		}

		// Success - PUT returns empty body, so fetch the updated SLO
		if resp.IsSuccess() {
			updated, err := h.fetchSLO(ctx, id)
			if err != nil {
				// Update succeeded but couldn't fetch result, return success message
				return textResult(fmt.Sprintf("SLO %s updated successfully (new version available)", id)), nil
			}
			// Return the updated SLO with new version
			result, _ := json.MarshalIndent(updated, "", "  ")
			return toolResult(result), nil
		}

		// Handle 409 Conflict - version mismatch, retry with fresh version
		if resp.StatusCode == 409 && attempt < maxRetryAttempts-1 {
			continue // Retry loop will fetch fresh version
		}

		// Other errors or final retry attempt failed
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}

	return toolError(fmt.Errorf("update failed after %d attempts due to concurrent modifications", maxRetryAttempts)), nil
}

// handleSLODeleteV2 deletes an SLO with optimistic locking
func (h *Handlers) handleSLODeleteV2(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	// Fetch current SLO to get version for optimistic locking
	for attempt := 0; attempt < maxRetryAttempts; attempt++ {
		current, err := h.fetchSLO(ctx, id)
		if err != nil {
			// If 404, SLO doesn't exist
			return toolError(err), nil
		}

		// DELETE with optimistic-locking-version
		resp, err := h.Client.Delete(
			ctx,
			sloBasePath+"/"+id,
			client.WithQueryParams(map[string]string{
				"optimistic-locking-version": current.Version,
			}),
		)
		if err != nil {
			return toolError(err), nil
		}

		// Success (204 No Content)
		if resp.IsSuccess() {
			return textResult(fmt.Sprintf("SLO %s deleted successfully", id)), nil
		}

		// Handle 409 Conflict - retry with fresh version
		if resp.StatusCode == 409 && attempt < maxRetryAttempts-1 {
			continue
		}

		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}

	return toolError(fmt.Errorf("delete failed after %d attempts due to concurrent modifications", maxRetryAttempts)), nil
}

// handleSLOTemplatesListV2 lists SLO templates
func (h *Handlers) handleSLOTemplatesListV2(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	params := make(map[string]string)
	if pageSize := getIntParam(req.Params.Arguments, "page_size", 0); pageSize > 0 {
		params["page-size"] = fmt.Sprintf("%d", pageSize)
	}

	resp, err := h.Client.Get(ctx, sloTemplatePath, client.WithQueryParams(params))
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

// handleSLOTemplateGetV2 retrieves a specific SLO template
func (h *Handlers) handleSLOTemplateGetV2(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	// URL-encode the template ID (CRITICAL: IDs are Base64 and contain == or /)
	encodedID := url.PathEscape(id)

	resp, err := h.Client.Get(ctx, sloTemplatePath+"/"+encodedID)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult(resp.Body), nil
}

// ==================== SLO Evaluation Handlers ====================

// handleSLOEvaluate provides a high-level evaluation that handles async automatically
func (h *Handlers) handleSLOEvaluate(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	// Build request payload
	payload := map[string]interface{}{
		"id":                          id,
		"requestTimeoutMilliseconds":  evalDefaultTimeout,
	}

	// Add custom timeframe if provided
	timeframeFrom := getStringParam(req.Params.Arguments, "timeframe_from")
	timeframeTo := getStringParam(req.Params.Arguments, "timeframe_to")
	if timeframeFrom != "" || timeframeTo != "" {
		customTimeframe := make(map[string]string)
		if timeframeFrom != "" {
			customTimeframe["timeframeFrom"] = timeframeFrom
		}
		if timeframeTo != "" {
			customTimeframe["timeframeTo"] = timeframeTo
		}
		payload["customTimeframe"] = customTimeframe
	}

	// Step 1: Start evaluation
	resp, err := h.Client.Post(ctx, sloBasePath+"/evaluation:start", payload)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}

	// Parse response to determine if sync or async
	var startResp SLOEvaluationStartResponse
	if err := json.Unmarshal(resp.Body, &startResp); err != nil {
		return toolError(fmt.Errorf("parse evaluation response: %w", err)), nil
	}

	// Scenario A: Synchronous - results available immediately
	if len(startResp.EvaluationResults) > 0 {
		return toolResult(resp.Body), nil
	}

	// Scenario B: Asynchronous - poll for results
	if startResp.EvaluationToken == "" {
		return toolError(fmt.Errorf("unexpected response: no results and no evaluation token")), nil
	}

	// URL-encode the token (CRITICAL: token contains Base64 == characters)
	encodedToken := url.QueryEscape(startResp.EvaluationToken)

	// Poll loop
	startTime := time.Now()
	for {
		// Check timeout
		if time.Since(startTime) > evalMaxPollTime {
			// Try to cancel and return timeout error
			h.cancelEvaluation(ctx, encodedToken)
			return toolError(fmt.Errorf("evaluation timed out after %v", evalMaxPollTime)), nil
		}

		// Wait before polling
		time.Sleep(evalPollInterval)

		// Poll for status
		pollResp, err := h.Client.Get(
			ctx,
			sloBasePath+"/evaluation:poll",
			client.WithQueryParams(map[string]string{
				"evaluation-token": encodedToken,
			}),
		)
		if err != nil {
			return toolError(err), nil
		}
		if !pollResp.IsSuccess() {
			return toolError(fmt.Errorf(client.FormatError(pollResp))), nil
		}

		// Parse poll response
		var poll SLOPollResponse
		if err := json.Unmarshal(pollResp.Body, &poll); err != nil {
			return toolError(fmt.Errorf("parse poll response: %w", err)), nil
		}

		switch poll.Status {
		case "COMPLETED":
			// Return the full poll response with results
			return toolResult(pollResp.Body), nil
		case "IN_PROGRESS":
			// Continue polling
			continue
		default:
			return toolError(fmt.Errorf("evaluation failed with status: %s", poll.Status)), nil
		}
	}
}

// handleSLOEvaluationStart starts an evaluation (low-level)
func (h *Handlers) handleSLOEvaluationStart(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	// Build request payload
	payload := map[string]interface{}{
		"id": id,
	}

	// Timeout
	timeout := getIntParam(req.Params.Arguments, "timeout_ms", evalDefaultTimeout)
	payload["requestTimeoutMilliseconds"] = timeout

	// Custom timeframe
	timeframeFrom := getStringParam(req.Params.Arguments, "timeframe_from")
	timeframeTo := getStringParam(req.Params.Arguments, "timeframe_to")
	if timeframeFrom != "" || timeframeTo != "" {
		customTimeframe := make(map[string]string)
		if timeframeFrom != "" {
			customTimeframe["timeframeFrom"] = timeframeFrom
		}
		if timeframeTo != "" {
			customTimeframe["timeframeTo"] = timeframeTo
		}
		payload["customTimeframe"] = customTimeframe
	}

	resp, err := h.Client.Post(ctx, sloBasePath+"/evaluation:start", payload)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}

	return toolResult(resp.Body), nil
}

// handleSLOEvaluationPoll polls evaluation status (low-level)
func (h *Handlers) handleSLOEvaluationPoll(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	token := getStringParam(req.Params.Arguments, "evaluation_token")
	if token == "" {
		return toolError(fmt.Errorf("evaluation_token is required")), nil
	}

	// URL-encode the token (CRITICAL: token contains Base64 == characters)
	encodedToken := url.QueryEscape(token)

	resp, err := h.Client.Get(
		ctx,
		sloBasePath+"/evaluation:poll",
		client.WithQueryParams(map[string]string{
			"evaluation-token": encodedToken,
		}),
	)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}

	return toolResult(resp.Body), nil
}

// handleSLOEvaluationCancel cancels a running evaluation (low-level)
func (h *Handlers) handleSLOEvaluationCancel(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	token := getStringParam(req.Params.Arguments, "evaluation_token")
	if token == "" {
		return toolError(fmt.Errorf("evaluation_token is required")), nil
	}

	// URL-encode the token
	encodedToken := url.QueryEscape(token)

	resp, err := h.Client.Post(
		ctx,
		sloBasePath+"/evaluation:cancel",
		nil, // Empty body
		client.WithQueryParams(map[string]string{
			"evaluation-token": encodedToken,
		}),
	)
	if err != nil {
		return toolError(err), nil
	}

	// 204 No Content = success
	if resp.StatusCode == 204 {
		return textResult("Evaluation cancelled successfully"), nil
	}

	// 410 Gone = job already finished (treat as success)
	if resp.StatusCode == 410 {
		return textResult("Evaluation already completed before cancellation"), nil
	}

	// 404 = token invalid
	if resp.StatusCode == 404 {
		return toolError(fmt.Errorf("evaluation token not found or invalid")), nil
	}

	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}

	return textResult("Evaluation cancelled"), nil
}

// cancelEvaluation is a helper to cancel an evaluation (used internally for timeout cleanup)
func (h *Handlers) cancelEvaluation(ctx context.Context, encodedToken string) {
	// Fire-and-forget cancel request
	_, _ = h.Client.Post(
		ctx,
		sloBasePath+"/evaluation:cancel",
		nil,
		client.WithQueryParams(map[string]string{
			"evaluation-token": encodedToken,
		}),
	)
}
