// Package server provides the main MCP server implementation.
package server

import (
	"github.com/dynatrace/dynatrace-platform-mcp-server/internal/client"
	"github.com/dynatrace/dynatrace-platform-mcp-server/internal/config"
	"github.com/dynatrace/dynatrace-platform-mcp-server/internal/tools"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

func RegisterTools(s *mcpserver.MCPServer, c *client.Client, cfg *config.Config) {
	h := &tools.Handlers{Client: c}
	isEnabled := func(name string) bool { return cfg.IsEnabled(name) }

	// ==================== Base Tools (existing 65 tools) ====================
	tools.RegisterDQLTools(s, h, isEnabled)
	tools.RegisterAutomationTools(s, h, isEnabled)
	tools.RegisterDocumentTools(s, h, isEnabled)
	tools.RegisterDavisTools(s, h, isEnabled)
	tools.RegisterSLOToolsV2(s, h, isEnabled) // Refactored with optimistic locking support
	tools.RegisterOpenPipelineTools(s, h, isEnabled)
	tools.RegisterNotificationTools(s, h, isEnabled)
	tools.RegisterStorageTools(s, h, isEnabled)
	tools.RegisterVulnerabilityTools(s, h, isEnabled)
	tools.RegisterHubTools(s, h, isEnabled)
	tools.RegisterIAMTools(s, h, isEnabled)
	tools.RegisterPlatformTools(s, h, isEnabled)
	tools.RegisterStateTools(s, h, isEnabled)
	tools.RegisterAppEngineTools(s, h, isEnabled)
	tools.RegisterEmailTools(s, h, isEnabled)

	// ==================== Extended Tools (151 new tools for full API coverage) ====================
	tools.RegisterAutomationExtendedTools(s, h, isEnabled)      // +47: execution logs/tasks, business calendars CRUD, scheduling rules CRUD, workflow history
	tools.RegisterDocumentsExtendedTools(s, h, isEnabled)       // +29: snapshots, locking, shares, trash management
	tools.RegisterDavisExtendedTools(s, h, isEnabled)           // +10: nl2dql, dql2nl, document search, analyzer schemas, feedback
	tools.RegisterStorageExtendedTools(s, h, isEnabled)         // +21: filter segments, fieldsets, record deletion, lookup tables
	tools.RegisterHubExtendedTools(s, h, isEnabled)             // +5:  releases, extension/technology details, categories
	tools.RegisterVulnerabilitiesExtendedTools(s, h, isEnabled) // +12: segmentation, affected entities, davis assessment
	tools.RegisterOpenPipelineExtendedTools(s, h, isEnabled)    // +9:  config update, DQL/matcher autocomplete/verify, processor preview
	tools.RegisterNotificationsExtendedTools(s, h, isEnabled)   // +15: event/resource/self notifications CRUD
	tools.RegisterIAMExtendedTools(s, h, isEnabled)             // +13: user/group CRUD, service users, license settings
}
