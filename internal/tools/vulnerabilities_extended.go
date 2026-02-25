package tools

import (
	"context"
	"fmt"

	"github.com/dynatrace/dynatrace-platform-mcp-server/internal/client"
	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

// RegisterVulnerabilitiesExtendedTools registers additional vulnerability/security tools.
func RegisterVulnerabilitiesExtendedTools(s *mcpserver.MCPServer, h *Handlers, isEnabled func(string) bool) {
	// ==================== Vulnerabilities Segmentation ====================
	if isEnabled("dt_vulnerabilities_segment") {
		s.AddTool(mcp.Tool{
			Name:        "dt_vulnerabilities_segment",
			Description: `Get vulnerabilities with segmentation (grouping/aggregation).`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"filter":       map[string]interface{}{"type": "string", "description": "Filter query"},
					"segment_by":   map[string]interface{}{"type": "array", "description": "Fields to segment by"},
					"aggregations": map[string]interface{}{"type": "array", "description": "Aggregation functions"},
				},
			},
		}, h.handleVulnerabilitiesSegment)
	}

	if isEnabled("dt_vulnerability_segment") {
		s.AddTool(mcp.Tool{
			Name:        "dt_vulnerability_segment",
			Description: `Get a specific vulnerability with segmentation.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":           map[string]interface{}{"type": "string", "description": "Vulnerability ID"},
					"segment_by":   map[string]interface{}{"type": "array", "description": "Fields to segment by"},
					"aggregations": map[string]interface{}{"type": "array", "description": "Aggregation functions"},
				},
				Required: []string{"id"},
			},
		}, h.handleVulnerabilitySegment)
	}

	// ==================== Affected Entities ====================
	if isEnabled("dt_vulnerability_affected_entities_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_vulnerability_affected_entities_list",
			Description: `List entities affected by a vulnerability.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":        map[string]interface{}{"type": "string", "description": "Vulnerability ID"},
					"filter":    map[string]interface{}{"type": "string", "description": "Filter query"},
					"page_size": map[string]interface{}{"type": "integer", "description": "Page size"},
				},
				Required: []string{"id"},
			},
		}, h.handleVulnerabilityAffectedEntitiesList)
	}

	if isEnabled("dt_vulnerability_affected_entities_segment") {
		s.AddTool(mcp.Tool{
			Name:        "dt_vulnerability_affected_entities_segment",
			Description: `Get affected entities with segmentation.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":           map[string]interface{}{"type": "string", "description": "Vulnerability ID"},
					"segment_by":   map[string]interface{}{"type": "array", "description": "Fields to segment by"},
					"aggregations": map[string]interface{}{"type": "array", "description": "Aggregations"},
				},
				Required: []string{"id"},
			},
		}, h.handleVulnerabilityAffectedEntitiesSegment)
	}

	if isEnabled("dt_vulnerability_affected_entities_muting") {
		s.AddTool(mcp.Tool{
			Name:        "dt_vulnerability_affected_entities_muting",
			Description: `Mute/unmute affected entities for a vulnerability.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":       map[string]interface{}{"type": "string", "description": "Vulnerability ID"},
					"entities": map[string]interface{}{"type": "array", "description": "Entity IDs to mute/unmute"},
					"muted":    map[string]interface{}{"type": "boolean", "description": "Mute (true) or unmute (false)"},
					"reason":   map[string]interface{}{"type": "string", "description": "Reason for muting"},
				},
				Required: []string{"id", "entities", "muted"},
			},
		}, h.handleVulnerabilityAffectedEntitiesMuting)
	}

	if isEnabled("dt_vulnerability_affected_entities_set_tracking_links") {
		s.AddTool(mcp.Tool{
			Name:        "dt_vulnerability_affected_entities_set_tracking_links",
			Description: `Set tracking links for affected entities.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":             map[string]interface{}{"type": "string", "description": "Vulnerability ID"},
					"entities":       map[string]interface{}{"type": "array", "description": "Entity IDs"},
					"tracking_links": map[string]interface{}{"type": "array", "description": "Tracking links to set"},
				},
				Required: []string{"id", "entities", "tracking_links"},
			},
		}, h.handleVulnerabilitySetTrackingLinks)
	}

	if isEnabled("dt_vulnerability_affected_entities_delete_tracking_links") {
		s.AddTool(mcp.Tool{
			Name:        "dt_vulnerability_affected_entities_delete_tracking_links",
			Description: `Delete tracking links from affected entities.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":             map[string]interface{}{"type": "string", "description": "Vulnerability ID"},
					"entities":       map[string]interface{}{"type": "array", "description": "Entity IDs"},
					"tracking_links": map[string]interface{}{"type": "array", "description": "Tracking link IDs to delete"},
				},
				Required: []string{"id", "entities", "tracking_links"},
			},
		}, h.handleVulnerabilityDeleteTrackingLinks)
	}

	// ==================== Davis Assessment ====================
	if isEnabled("dt_vulnerability_davis_assessment") {
		s.AddTool(mcp.Tool{
			Name:        "dt_vulnerability_davis_assessment",
			Description: `Get Davis AI security assessment for a vulnerability.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id": map[string]interface{}{"type": "string", "description": "Vulnerability ID"},
				},
				Required: []string{"id"},
			},
		}, h.handleVulnerabilityDavisAssessment)
	}

	if isEnabled("dt_vulnerability_davis_assessment_segment") {
		s.AddTool(mcp.Tool{
			Name:        "dt_vulnerability_davis_assessment_segment",
			Description: `Get Davis assessment with segmentation.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"id":           map[string]interface{}{"type": "string", "description": "Vulnerability ID"},
					"segment_by":   map[string]interface{}{"type": "array", "description": "Fields to segment by"},
					"aggregations": map[string]interface{}{"type": "array", "description": "Aggregations"},
				},
				Required: []string{"id"},
			},
		}, h.handleVulnerabilityDavisAssessmentSegment)
	}

	// ==================== Davis Security Recommendations ====================
	if isEnabled("dt_davis_security_recommendations_list") {
		s.AddTool(mcp.Tool{
			Name:        "dt_davis_security_recommendations_list",
			Description: `List Davis AI security recommendations.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"filter":    map[string]interface{}{"type": "string", "description": "Filter query"},
					"page_size": map[string]interface{}{"type": "integer", "description": "Page size"},
				},
			},
		}, h.handleDavisSecurityRecommendationsList)
	}

	if isEnabled("dt_davis_security_recommendations_segment") {
		s.AddTool(mcp.Tool{
			Name:        "dt_davis_security_recommendations_segment",
			Description: `Get security recommendations with segmentation.`,
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]interface{}{
					"filter":       map[string]interface{}{"type": "string", "description": "Filter query"},
					"segment_by":   map[string]interface{}{"type": "array", "description": "Fields to segment by"},
					"aggregations": map[string]interface{}{"type": "array", "description": "Aggregations"},
				},
			},
		}, h.handleDavisSecurityRecommendationsSegment)
	}
}

// ==================== Handler Implementations ====================

func (h *Handlers) handleVulnerabilitiesSegment(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	body := map[string]interface{}{}
	if filter := getStringParam(args, "filter"); filter != "" {
		body["filter"] = filter
	}
	if segmentBy, ok := args["segment_by"].([]interface{}); ok {
		body["segmentBy"] = segmentBy
	}
	if aggregations, ok := args["aggregations"].([]interface{}); ok {
		body["aggregations"] = aggregations
	}

	resp, err := h.Client.Post(ctx, "/platform/app-security/v1/vulnerabilities:segment", body)
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

func (h *Handlers) handleVulnerabilitySegment(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	body := map[string]interface{}{}
	if segmentBy, ok := args["segment_by"].([]interface{}); ok {
		body["segmentBy"] = segmentBy
	}
	if aggregations, ok := args["aggregations"].([]interface{}); ok {
		body["aggregations"] = aggregations
	}

	resp, err := h.Client.Post(ctx, "/platform/app-security/v1/vulnerabilities/"+id+":segment", body)
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

func (h *Handlers) handleVulnerabilityAffectedEntitiesList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	params := map[string]string{}
	if filter := getStringParam(args, "filter"); filter != "" {
		params["filter"] = filter
	}
	if pageSize := getIntParam(args, "page_size", 0); pageSize > 0 {
		params["page-size"] = fmt.Sprintf("%d", pageSize)
	}

	resp, err := h.Client.Get(ctx, "/platform/app-security/v1/vulnerabilities/"+id+"/affected-entities", client.WithQueryParams(params))
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

func (h *Handlers) handleVulnerabilityAffectedEntitiesSegment(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	body := map[string]interface{}{}
	if segmentBy, ok := args["segment_by"].([]interface{}); ok {
		body["segmentBy"] = segmentBy
	}
	if aggregations, ok := args["aggregations"].([]interface{}); ok {
		body["aggregations"] = aggregations
	}

	resp, err := h.Client.Post(ctx, "/platform/app-security/v1/vulnerabilities/"+id+"/affected-entities:segment", body)
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

func (h *Handlers) handleVulnerabilityAffectedEntitiesMuting(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	entities, ok := args["entities"].([]interface{})
	if !ok {
		return toolError(fmt.Errorf("entities are required")), nil
	}

	muted, ok := args["muted"].(bool)
	if !ok {
		return toolError(fmt.Errorf("muted is required")), nil
	}

	body := map[string]interface{}{
		"entities": entities,
		"muted":    muted,
	}
	if reason := getStringParam(args, "reason"); reason != "" {
		body["reason"] = reason
	}

	resp, err := h.Client.Post(ctx, "/platform/app-security/v1/vulnerabilities/"+id+"/affected-entities:muting", body)
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

func (h *Handlers) handleVulnerabilitySetTrackingLinks(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	entities, ok := args["entities"].([]interface{})
	if !ok {
		return toolError(fmt.Errorf("entities are required")), nil
	}

	trackingLinks, ok := args["tracking_links"].([]interface{})
	if !ok {
		return toolError(fmt.Errorf("tracking_links are required")), nil
	}

	body := map[string]interface{}{
		"entities":      entities,
		"trackingLinks": trackingLinks,
	}

	resp, err := h.Client.Post(ctx, "/platform/app-security/v1/vulnerabilities/"+id+"/affected-entities:set-tracking-links", body)
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

func (h *Handlers) handleVulnerabilityDeleteTrackingLinks(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	entities, ok := args["entities"].([]interface{})
	if !ok {
		return toolError(fmt.Errorf("entities are required")), nil
	}

	trackingLinks, ok := args["tracking_links"].([]interface{})
	if !ok {
		return toolError(fmt.Errorf("tracking_links are required")), nil
	}

	body := map[string]interface{}{
		"entities":      entities,
		"trackingLinks": trackingLinks,
	}

	resp, err := h.Client.Post(ctx, "/platform/app-security/v1/vulnerabilities/"+id+"/affected-entities:delete-tracking-links", body)
	if err != nil {
		return toolError(err), nil
	}
	if !resp.IsSuccess() {
		return toolError(fmt.Errorf(client.FormatError(resp))), nil
	}
	return toolResult("Tracking links deleted"), nil
}

func (h *Handlers) handleVulnerabilityDavisAssessment(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id := getStringParam(req.Params.Arguments, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	resp, err := h.Client.Get(ctx, "/platform/app-security/v1/vulnerabilities/"+id+"/davis-assessment")
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

func (h *Handlers) handleVulnerabilityDavisAssessmentSegment(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	id := getStringParam(args, "id")
	if id == "" {
		return toolError(fmt.Errorf("id is required")), nil
	}

	body := map[string]interface{}{}
	if segmentBy, ok := args["segment_by"].([]interface{}); ok {
		body["segmentBy"] = segmentBy
	}
	if aggregations, ok := args["aggregations"].([]interface{}); ok {
		body["aggregations"] = aggregations
	}

	resp, err := h.Client.Post(ctx, "/platform/app-security/v1/vulnerabilities/"+id+"/davis-assessment:segment", body)
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

func (h *Handlers) handleDavisSecurityRecommendationsList(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	params := map[string]string{}
	if filter := getStringParam(args, "filter"); filter != "" {
		params["filter"] = filter
	}
	if pageSize := getIntParam(args, "page_size", 0); pageSize > 0 {
		params["page-size"] = fmt.Sprintf("%d", pageSize)
	}

	resp, err := h.Client.Get(ctx, "/platform/app-security/v1/davis-security-recommendations", client.WithQueryParams(params))
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

func (h *Handlers) handleDavisSecurityRecommendationsSegment(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.Params.Arguments
	body := map[string]interface{}{}
	if filter := getStringParam(args, "filter"); filter != "" {
		body["filter"] = filter
	}
	if segmentBy, ok := args["segment_by"].([]interface{}); ok {
		body["segmentBy"] = segmentBy
	}
	if aggregations, ok := args["aggregations"].([]interface{}); ok {
		body["aggregations"] = aggregations
	}

	resp, err := h.Client.Post(ctx, "/platform/app-security/v1/davis-security-recommendations:segment", body)
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
