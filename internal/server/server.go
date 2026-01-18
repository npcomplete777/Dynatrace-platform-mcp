// Package server provides the main MCP server implementation.
package server

import (
	"github.com/dynatrace/dynatrace-platform-mcp-server/internal/client"
	"github.com/dynatrace/dynatrace-platform-mcp-server/internal/tools"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

func RegisterTools(s *mcpserver.MCPServer, c *client.Client) {
	h := &tools.Handlers{Client: c}

	// ==================== Base Tools (existing 65 tools) ====================
	tools.RegisterDQLTools(s, h)
	tools.RegisterAutomationTools(s, h)
	tools.RegisterDocumentTools(s, h)
	tools.RegisterDavisTools(s, h)
	tools.RegisterSLOToolsV2(s, h) // Refactored with optimistic locking support
	tools.RegisterOpenPipelineTools(s, h)
	tools.RegisterNotificationTools(s, h)
	tools.RegisterStorageTools(s, h)
	tools.RegisterVulnerabilityTools(s, h)
	tools.RegisterHubTools(s, h)
	tools.RegisterIAMTools(s, h)
	tools.RegisterPlatformTools(s, h)
	tools.RegisterStateTools(s, h)
	tools.RegisterAppEngineTools(s, h)
	tools.RegisterEmailTools(s, h)

	// ==================== Extended Tools (151 new tools for full API coverage) ====================
	tools.RegisterAutomationExtendedTools(s, h)      // +47: execution logs/tasks, business calendars CRUD, scheduling rules CRUD, workflow history
	tools.RegisterDocumentsExtendedTools(s, h)       // +29: snapshots, locking, shares, trash management
	tools.RegisterDavisExtendedTools(s, h)           // +10: nl2dql, dql2nl, document search, analyzer schemas, feedback
	tools.RegisterStorageExtendedTools(s, h)         // +21: filter segments, fieldsets, record deletion, lookup tables
	tools.RegisterHubExtendedTools(s, h)             // +5:  releases, extension/technology details, categories
	tools.RegisterVulnerabilitiesExtendedTools(s, h) // +12: segmentation, affected entities, davis assessment
	tools.RegisterOpenPipelineExtendedTools(s, h)    // +9:  config update, DQL/matcher autocomplete/verify, processor preview
	tools.RegisterNotificationsExtendedTools(s, h)   // +15: event/resource/self notifications CRUD
	tools.RegisterIAMExtendedTools(s, h)             // +13: user/group CRUD, service users, license settings
}
