# Dynatrace Platform MCP Server

A [Model Context Protocol (MCP)](https://modelcontextprotocol.io) server that exposes the Dynatrace Platform (Grail) APIs as tools for AI assistants like Claude. Covers the full surface area of the Dynatrace Platform API — DQL queries, automation workflows, documents, SLOs, Davis AI, OpenPipeline, storage, security vulnerabilities, IAM, and more.

**~160 tools across 12 API domains.**

---

## Requirements

- Go 1.23+
- A Dynatrace environment with Platform/Grail enabled
- A Platform API token (`dt0s16.xxx`) with the scopes you need

---

## Build

```bash
go build -o bin/dynatrace-platform-mcp ./cmd/server
```

---

## Configuration

### Environment variables (required)

| Variable | Alt name | Description |
|---|---|---|
| `DT_BASE_URL` | `DYNATRACE_BASE_URL` | Your environment URL, e.g. `https://abc12345.apps.dynatrace.com` |
| `DT_PLATFORM_TOKEN` | `DYNATRACE_PLATFORM_TOKEN` | Platform API token (`dt0s16.xxx`) |
| `DT_ACCOUNT_URN` | — | Account URN for IAM tools (optional) |
| `DT_DEBUG` | — | Set to `true` for verbose logging |

The server accepts both `.live.dynatrace.com` and `.apps.dynatrace.com` URLs and normalises them automatically.

### Tool configuration (optional)

Tools are individually enabled or disabled via a YAML file. By default every tool is enabled. The config file path is resolved in this order:

1. `DYNATRACE_CONFIG_FILE` environment variable
2. `config.yaml` in the working directory
3. No file → all tools enabled

**Format:**

```yaml
tools:
  tool_name:
    enabled: false   # omit or set true to enable
```

See [Recommended Profiles](#recommended-configuration-profiles) below for ready-to-use configurations.

---

## Running with Claude Desktop

Add to `~/Library/Application Support/Claude/claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "dynatrace-platform": {
      "command": "/path/to/bin/dynatrace-platform-mcp",
      "env": {
        "DT_BASE_URL": "https://abc12345.apps.dynatrace.com",
        "DT_PLATFORM_TOKEN": "dt0s16.YOUR_TOKEN",
        "DT_ACCOUNT_URN": "urn:dtaccount:your-account-id",
        "DYNATRACE_CONFIG_FILE": "/path/to/config.yaml"
      }
    }
  }
}
```

## Running with Claude Code (CLI)

Add to `~/.mcp.json`:

```json
{
  "mcpServers": {
    "dynatrace-platform": {
      "command": "/path/to/bin/dynatrace-platform-mcp",
      "env": {
        "DT_BASE_URL": "https://abc12345.apps.dynatrace.com",
        "DT_PLATFORM_TOKEN": "dt0s16.YOUR_TOKEN",
        "DYNATRACE_CONFIG_FILE": "/path/to/config.yaml"
      }
    }
  }
}
```

---

## Tool Domains

### DQL (3 tools)
Query any Dynatrace Grail data — logs, metrics, traces, events, entities, business events.

| Tool | Description |
|---|---|
| `dt_dql_query` | Execute a DQL query against Grail |
| `dt_dql_autocomplete` | Get autocomplete suggestions for a partial query |
| `dt_dql_parse` | Parse and validate a DQL query without executing it |

### Automation / Workflows (47 tools)
Full lifecycle management for Dynatrace Workflows — create, run, inspect executions, manage business calendars and scheduling rules.

Key tools: `dt_workflows_list`, `dt_workflow_create`, `dt_workflow_run`, `dt_workflow_executions`, `dt_execution_task_log`, `dt_business_calendar_create`, `dt_scheduling_rule_create`

### Documents (34 tools)
Manage Dynatrace documents (notebooks, dashboards) including snapshots, locking, sharing, and trash.

Key tools: `dt_documents_list`, `dt_document_create`, `dt_document_snapshots_list`, `dt_document_lock_acquire`, `dt_environment_share_create`, `dt_direct_share_create`

### Davis AI / Copilot (12 tools)
Natural language to DQL translation, DQL to natural language, document semantic search, and analyzer schemas.

Key tools: `dt_copilot_nl2dql`, `dt_copilot_dql2nl`, `dt_copilot_document_search`, `dt_davis_analyze`, `dt_davis_analyzers_list`

### SLOs (11 tools)
Create, evaluate, and manage Service Level Objectives with optimistic locking support.

Key tools: `dt_slos_list`, `dt_slo_create`, `dt_slo_evaluate`, `dt_slo_evaluation_start`, `dt_slo_evaluation_poll`

### OpenPipeline (11 tools)
Manage ingestion pipelines, verify DQL/matcher expressions, preview processor results.

Key tools: `dt_openpipeline_configs_list`, `dt_openpipeline_configuration_update`, `dt_openpipeline_dql_verify`, `dt_openpipeline_processor_preview`

### Notifications (16 tools)
Event notifications, resource notifications, and personal (self) notification subscriptions — full CRUD.

Key tools: `dt_event_notifications_list`, `dt_event_notification_create`, `dt_resource_notification_create`, `dt_self_notification_create`

### Storage / Grail Buckets (22 tools)
Bucket management, filter segments, fieldsets, record deletion jobs, and lookup table uploads.

Key tools: `dt_buckets_list`, `dt_bucket_update`, `dt_filter_segment_create`, `dt_fieldset_create`, `dt_record_deletion_execute`, `dt_lookup_table_upload`

### Hub (6 tools)
Browse the Dynatrace Hub — extensions, technologies, app releases, categories.

Key tools: `dt_hub_items_list`, `dt_hub_extension_get`, `dt_hub_technology_get`

### Security Vulnerabilities (12 tools)
List, segment, and manage vulnerabilities; control Davis security recommendations and affected entity tracking.

Key tools: `dt_vulnerabilities_list`, `dt_vulnerability_affected_entities_list`, `dt_vulnerability_davis_assessment`, `dt_davis_security_recommendations_list`

### IAM (13 tools)
User and group CRUD, service users, policy listing, and environment license settings.

Key tools: `dt_iam_groups_list`, `dt_iam_user_create`, `dt_iam_group_create`, `dt_environment_license_settings`

### Platform / Misc (4 tools)
Environment info, app state, app listing, and email sending.

Key tools: `dt_environment_info`, `dt_apps_list`, `dt_email_send`

---

## Recommended Configuration Profiles

Copy the relevant block into your `config.yaml` (or a separate file pointed to by `DYNATRACE_CONFIG_FILE`).

### Read-Only

Safe for shared environments, demos, or any context where you want to prevent writes. All list/get/query/evaluate tools remain active; every create, update, delete, and mutation tool is disabled.

```yaml
# config-readonly.yaml
# All write operations disabled — safe for shared or production read access.
tools:
  # Automation writes
  dt_workflow_create:
    enabled: false
  dt_workflow_update:
    enabled: false
  dt_workflow_delete:
    enabled: false
  dt_workflow_run:
    enabled: false
  dt_workflow_duplicate:
    enabled: false
  dt_workflow_history_restore:
    enabled: false
  dt_workflow_reset_throttles:
    enabled: false
  dt_execution_task_cancel:
    enabled: false
  dt_business_calendar_create:
    enabled: false
  dt_business_calendar_update:
    enabled: false
  dt_business_calendar_delete:
    enabled: false
  dt_business_calendar_duplicate:
    enabled: false
  dt_scheduling_rule_create:
    enabled: false
  dt_scheduling_rule_update:
    enabled: false
  dt_scheduling_rule_delete:
    enabled: false
  dt_scheduling_rule_duplicate:
    enabled: false
  # Document writes
  dt_document_create:
    enabled: false
  dt_document_update:
    enabled: false
  dt_document_delete:
    enabled: false
  dt_document_snapshot_restore:
    enabled: false
  dt_document_lock_acquire:
    enabled: false
  dt_document_lock_release:
    enabled: false
  dt_document_transfer_owner:
    enabled: false
  dt_documents_bulk_delete:
    enabled: false
  dt_environment_share_create:
    enabled: false
  dt_environment_share_update:
    enabled: false
  dt_environment_share_delete:
    enabled: false
  dt_environment_share_claim:
    enabled: false
  dt_direct_share_create:
    enabled: false
  dt_direct_share_update:
    enabled: false
  dt_direct_share_delete:
    enabled: false
  dt_direct_share_recipients_add:
    enabled: false
  dt_direct_share_recipients_remove:
    enabled: false
  dt_trash_document_delete:
    enabled: false
  dt_trash_document_restore:
    enabled: false
  # SLO writes
  dt_slo_create:
    enabled: false
  dt_slo_update:
    enabled: false
  dt_slo_delete:
    enabled: false
  dt_slo_evaluation_start:
    enabled: false
  dt_slo_evaluation_cancel:
    enabled: false
  # OpenPipeline writes
  dt_openpipeline_configuration_update:
    enabled: false
  # Notification writes
  dt_event_notification_create:
    enabled: false
  dt_event_notification_update:
    enabled: false
  dt_event_notification_delete:
    enabled: false
  dt_resource_notification_create:
    enabled: false
  dt_resource_notification_update:
    enabled: false
  dt_resource_notification_delete:
    enabled: false
  dt_self_notification_create:
    enabled: false
  dt_self_notification_update:
    enabled: false
  dt_self_notification_delete:
    enabled: false
  # Storage writes
  dt_bucket_update:
    enabled: false
  dt_bucket_truncate:
    enabled: false
  dt_filter_segment_create:
    enabled: false
  dt_filter_segment_update:
    enabled: false
  dt_filter_segment_delete:
    enabled: false
  dt_fieldset_create:
    enabled: false
  dt_fieldset_update:
    enabled: false
  dt_fieldset_delete:
    enabled: false
  dt_record_deletion_execute:
    enabled: false
  dt_record_deletion_cancel:
    enabled: false
  dt_lookup_table_upload:
    enabled: false
  dt_resource_files_delete:
    enabled: false
  # IAM writes
  dt_iam_user_create:
    enabled: false
  dt_iam_user_update:
    enabled: false
  dt_iam_user_delete:
    enabled: false
  dt_iam_group_create:
    enabled: false
  dt_iam_group_update:
    enabled: false
  dt_iam_group_delete:
    enabled: false
  dt_all_user_app_states_delete:
    enabled: false
  # Vulnerability writes
  dt_vulnerability_affected_entities_muting:
    enabled: false
  dt_vulnerability_affected_entities_set_tracking_links:
    enabled: false
  dt_vulnerability_affected_entities_delete_tracking_links:
    enabled: false
  # Email
  dt_email_send:
    enabled: false
```

---

### SRE / On-Call

Focused on observability and reliability — DQL, SLOs, vulnerabilities, and Davis AI. No automation workflow management, no IAM, no document or notification management.

```yaml
# config-sre.yaml
# SLO monitoring, DQL queries, Davis AI, and vulnerability awareness.
tools:
  # ---- Disable: Automation ----
  dt_workflow_create:
    enabled: false
  dt_workflow_update:
    enabled: false
  dt_workflow_delete:
    enabled: false
  dt_workflow_run:
    enabled: false
  dt_workflow_duplicate:
    enabled: false
  dt_workflow_history_restore:
    enabled: false
  dt_workflow_reset_throttles:
    enabled: false
  dt_execution_task_cancel:
    enabled: false
  dt_business_calendar_create:
    enabled: false
  dt_business_calendar_update:
    enabled: false
  dt_business_calendar_delete:
    enabled: false
  dt_business_calendar_duplicate:
    enabled: false
  dt_scheduling_rule_create:
    enabled: false
  dt_scheduling_rule_update:
    enabled: false
  dt_scheduling_rule_delete:
    enabled: false
  dt_scheduling_rule_duplicate:
    enabled: false
  # ---- Disable: Documents ----
  dt_document_create:
    enabled: false
  dt_document_update:
    enabled: false
  dt_document_delete:
    enabled: false
  dt_document_snapshot_restore:
    enabled: false
  dt_document_lock_acquire:
    enabled: false
  dt_document_lock_release:
    enabled: false
  dt_document_transfer_owner:
    enabled: false
  dt_documents_bulk_delete:
    enabled: false
  dt_environment_share_create:
    enabled: false
  dt_environment_share_update:
    enabled: false
  dt_environment_share_delete:
    enabled: false
  dt_environment_share_claim:
    enabled: false
  dt_direct_share_create:
    enabled: false
  dt_direct_share_update:
    enabled: false
  dt_direct_share_delete:
    enabled: false
  dt_direct_share_recipients_add:
    enabled: false
  dt_direct_share_recipients_remove:
    enabled: false
  dt_trash_document_delete:
    enabled: false
  dt_trash_document_restore:
    enabled: false
  # ---- Disable: OpenPipeline ----
  dt_openpipeline_configuration_update:
    enabled: false
  # ---- Disable: Notifications ----
  dt_event_notification_create:
    enabled: false
  dt_event_notification_update:
    enabled: false
  dt_event_notification_delete:
    enabled: false
  dt_resource_notification_create:
    enabled: false
  dt_resource_notification_update:
    enabled: false
  dt_resource_notification_delete:
    enabled: false
  dt_self_notification_create:
    enabled: false
  dt_self_notification_update:
    enabled: false
  dt_self_notification_delete:
    enabled: false
  # ---- Disable: Storage mutations ----
  dt_bucket_update:
    enabled: false
  dt_bucket_truncate:
    enabled: false
  dt_filter_segment_create:
    enabled: false
  dt_filter_segment_update:
    enabled: false
  dt_filter_segment_delete:
    enabled: false
  dt_fieldset_create:
    enabled: false
  dt_fieldset_update:
    enabled: false
  dt_fieldset_delete:
    enabled: false
  dt_record_deletion_execute:
    enabled: false
  dt_record_deletion_cancel:
    enabled: false
  dt_lookup_table_upload:
    enabled: false
  dt_resource_files_delete:
    enabled: false
  # ---- Disable: IAM ----
  dt_iam_user_create:
    enabled: false
  dt_iam_user_update:
    enabled: false
  dt_iam_user_delete:
    enabled: false
  dt_iam_group_create:
    enabled: false
  dt_iam_group_update:
    enabled: false
  dt_iam_group_delete:
    enabled: false
  dt_all_user_app_states_delete:
    enabled: false
  # ---- Disable: Email ----
  dt_email_send:
    enabled: false
```

---

### Developer / Automation Engineer

Workflow automation, DQL, documents, and Davis AI for development work. IAM, storage mutations, and email disabled.

```yaml
# config-developer.yaml
# Workflow automation, documents, DQL, Davis AI. No IAM or storage mutations.
tools:
  # ---- Disable: Destructive storage ops ----
  dt_bucket_truncate:
    enabled: false
  dt_record_deletion_execute:
    enabled: false
  dt_record_deletion_cancel:
    enabled: false
  dt_resource_files_delete:
    enabled: false
  dt_filter_segment_delete:
    enabled: false
  dt_fieldset_delete:
    enabled: false
  # ---- Disable: IAM ----
  dt_iam_user_create:
    enabled: false
  dt_iam_user_update:
    enabled: false
  dt_iam_user_delete:
    enabled: false
  dt_iam_group_create:
    enabled: false
  dt_iam_group_update:
    enabled: false
  dt_iam_group_delete:
    enabled: false
  dt_all_user_app_states_delete:
    enabled: false
  # ---- Disable: Vulnerability mutations ----
  dt_vulnerability_affected_entities_muting:
    enabled: false
  dt_vulnerability_affected_entities_set_tracking_links:
    enabled: false
  dt_vulnerability_affected_entities_delete_tracking_links:
    enabled: false
  # ---- Disable: OpenPipeline config changes ----
  dt_openpipeline_configuration_update:
    enabled: false
  # ---- Disable: Email ----
  dt_email_send:
    enabled: false
```

---

### Admin

All tools enabled. This is the default when no config file is present. Use explicitly to document intent.

```yaml
# config-admin.yaml
# All tools enabled. Full platform access — use in trusted environments only.
tools: {}
```

---

### DQL Only

Minimal footprint — only the three DQL query tools. Useful for embedding in pipelines or giving an AI assistant query-only access.

```yaml
# config-dql-only.yaml
# Only DQL tools active. Everything else disabled.
tools:
  dt_dql_autocomplete:
    enabled: true
  dt_dql_parse:
    enabled: true
  dt_dql_query:
    enabled: true
  # Automation
  dt_workflows_list:
    enabled: false
  dt_workflow_get:
    enabled: false
  dt_workflow_create:
    enabled: false
  dt_workflow_update:
    enabled: false
  dt_workflow_delete:
    enabled: false
  dt_workflow_run:
    enabled: false
  dt_workflow_executions:
    enabled: false
  dt_action_executions:
    enabled: false
  dt_action_execution_log:
    enabled: false
  dt_execution_log:
    enabled: false
  dt_execution_all_event_logs:
    enabled: false
  dt_execution_actions_list:
    enabled: false
  dt_execution_tasks_list:
    enabled: false
  dt_execution_task_get:
    enabled: false
  dt_execution_task_log:
    enabled: false
  dt_execution_task_result:
    enabled: false
  dt_execution_task_input:
    enabled: false
  dt_execution_task_cancel:
    enabled: false
  dt_execution_transitions_list:
    enabled: false
  dt_workflow_duplicate:
    enabled: false
  dt_workflow_export:
    enabled: false
  dt_workflow_history_list:
    enabled: false
  dt_workflow_history_get:
    enabled: false
  dt_workflow_history_restore:
    enabled: false
  dt_workflow_tasks_list:
    enabled: false
  dt_workflow_reset_throttles:
    enabled: false
  dt_business_calendar_get:
    enabled: false
  dt_business_calendar_create:
    enabled: false
  dt_business_calendar_update:
    enabled: false
  dt_business_calendar_delete:
    enabled: false
  dt_business_calendar_duplicate:
    enabled: false
  dt_business_calendar_history_list:
    enabled: false
  dt_scheduling_rule_get:
    enabled: false
  dt_scheduling_rule_create:
    enabled: false
  dt_scheduling_rule_update:
    enabled: false
  dt_scheduling_rule_delete:
    enabled: false
  dt_scheduling_rule_duplicate:
    enabled: false
  dt_scheduling_rule_preview:
    enabled: false
  dt_holiday_calendars_list:
    enabled: false
  dt_holiday_calendar_get:
    enabled: false
  dt_timezones_list:
    enabled: false
  dt_schedule_preview:
    enabled: false
  dt_event_trigger_filter_preview:
    enabled: false
  dt_automation_settings_get:
    enabled: false
  dt_automation_service_users_list:
    enabled: false
  dt_automation_user_settings_get:
    enabled: false
  dt_automation_user_permissions_get:
    enabled: false
  dt_automation_version_get:
    enabled: false
  # Documents
  dt_documents_list:
    enabled: false
  dt_document_get:
    enabled: false
  dt_document_create:
    enabled: false
  dt_document_update:
    enabled: false
  dt_document_delete:
    enabled: false
  dt_document_metadata_get:
    enabled: false
  dt_document_content_get:
    enabled: false
  dt_document_snapshots_list:
    enabled: false
  dt_document_snapshot_get:
    enabled: false
  dt_document_snapshot_restore:
    enabled: false
  dt_document_lock_inspect:
    enabled: false
  dt_document_lock_acquire:
    enabled: false
  dt_document_lock_release:
    enabled: false
  dt_document_transfer_owner:
    enabled: false
  dt_environment_shares_list:
    enabled: false
  dt_environment_share_create:
    enabled: false
  dt_environment_share_get:
    enabled: false
  dt_environment_share_update:
    enabled: false
  dt_environment_share_delete:
    enabled: false
  dt_environment_share_claim:
    enabled: false
  dt_environment_share_claimers_list:
    enabled: false
  dt_direct_shares_list:
    enabled: false
  dt_direct_share_create:
    enabled: false
  dt_direct_share_get:
    enabled: false
  dt_direct_share_update:
    enabled: false
  dt_direct_share_delete:
    enabled: false
  dt_direct_share_recipients_list:
    enabled: false
  dt_direct_share_recipients_add:
    enabled: false
  dt_direct_share_recipients_remove:
    enabled: false
  dt_trash_documents_list:
    enabled: false
  dt_trash_document_get:
    enabled: false
  dt_trash_document_delete:
    enabled: false
  dt_trash_document_restore:
    enabled: false
  dt_documents_bulk_delete:
    enabled: false
  # Davis AI
  dt_davis_analyzers_list:
    enabled: false
  dt_davis_analyze:
    enabled: false
  dt_davis_analyzer_documentation:
    enabled: false
  dt_davis_analyzer_input_schema:
    enabled: false
  dt_davis_analyzer_result_schema:
    enabled: false
  dt_davis_analyzer_validate:
    enabled: false
  dt_copilot_nl2dql:
    enabled: false
  dt_copilot_dql2nl:
    enabled: false
  dt_copilot_document_search:
    enabled: false
  dt_copilot_conversation_feedback:
    enabled: false
  dt_copilot_nl2dql_feedback:
    enabled: false
  dt_copilot_dql2nl_feedback:
    enabled: false
  # SLOs
  dt_slos_list:
    enabled: false
  dt_slo_get:
    enabled: false
  dt_slo_create:
    enabled: false
  dt_slo_update:
    enabled: false
  dt_slo_delete:
    enabled: false
  dt_slo_templates_list:
    enabled: false
  dt_slo_template_get:
    enabled: false
  dt_slo_evaluate:
    enabled: false
  dt_slo_evaluation_start:
    enabled: false
  dt_slo_evaluation_poll:
    enabled: false
  dt_slo_evaluation_cancel:
    enabled: false
  # OpenPipeline
  dt_openpipeline_configs_list:
    enabled: false
  dt_openpipeline_config_get:
    enabled: false
  dt_openpipeline_configuration_update:
    enabled: false
  dt_openpipeline_dql_autocomplete:
    enabled: false
  dt_openpipeline_dql_verify:
    enabled: false
  dt_openpipeline_matcher_autocomplete:
    enabled: false
  dt_openpipeline_matcher_verify:
    enabled: false
  dt_openpipeline_lql_to_dql:
    enabled: false
  dt_openpipeline_processor_preview:
    enabled: false
  dt_openpipeline_technologies_list:
    enabled: false
  dt_openpipeline_technology_processors_list:
    enabled: false
  # Notifications
  dt_notifications_list:
    enabled: false
  dt_event_notifications_list:
    enabled: false
  dt_event_notification_create:
    enabled: false
  dt_event_notification_get:
    enabled: false
  dt_event_notification_update:
    enabled: false
  dt_event_notification_delete:
    enabled: false
  dt_resource_notifications_list:
    enabled: false
  dt_resource_notification_create:
    enabled: false
  dt_resource_notification_get_by_resource:
    enabled: false
  dt_resource_notification_get:
    enabled: false
  dt_resource_notification_update:
    enabled: false
  dt_resource_notification_delete:
    enabled: false
  dt_self_notifications_list:
    enabled: false
  dt_self_notification_create:
    enabled: false
  dt_self_notification_get:
    enabled: false
  dt_self_notification_update:
    enabled: false
  dt_self_notification_delete:
    enabled: false
  # Storage
  dt_buckets_list:
    enabled: false
  dt_bucket_get:
    enabled: false
  dt_bucket_update:
    enabled: false
  dt_bucket_truncate:
    enabled: false
  dt_filter_segments_list:
    enabled: false
  dt_filter_segment_get:
    enabled: false
  dt_filter_segment_create:
    enabled: false
  dt_filter_segment_update:
    enabled: false
  dt_filter_segment_delete:
    enabled: false
  dt_filter_segments_entity_model:
    enabled: false
  dt_filter_segments_lean:
    enabled: false
  dt_fieldsets_list:
    enabled: false
  dt_fieldset_get:
    enabled: false
  dt_fieldset_create:
    enabled: false
  dt_fieldset_update:
    enabled: false
  dt_fieldset_delete:
    enabled: false
  dt_record_deletion_execute:
    enabled: false
  dt_record_deletion_status:
    enabled: false
  dt_record_deletion_cancel:
    enabled: false
  dt_lookup_table_upload:
    enabled: false
  dt_lookup_table_test_pattern:
    enabled: false
  dt_resource_files_delete:
    enabled: false
  # Hub
  dt_hub_items_list:
    enabled: false
  dt_hub_app_releases_list:
    enabled: false
  dt_hub_extension_get:
    enabled: false
  dt_hub_extension_releases_list:
    enabled: false
  dt_hub_technology_get:
    enabled: false
  dt_hub_categories_list:
    enabled: false
  # Vulnerabilities
  dt_vulnerabilities_list:
    enabled: false
  dt_vulnerabilities_segment:
    enabled: false
  dt_vulnerability_segment:
    enabled: false
  dt_vulnerability_affected_entities_list:
    enabled: false
  dt_vulnerability_affected_entities_segment:
    enabled: false
  dt_vulnerability_affected_entities_muting:
    enabled: false
  dt_vulnerability_affected_entities_set_tracking_links:
    enabled: false
  dt_vulnerability_affected_entities_delete_tracking_links:
    enabled: false
  dt_vulnerability_davis_assessment:
    enabled: false
  dt_vulnerability_davis_assessment_segment:
    enabled: false
  dt_davis_security_recommendations_list:
    enabled: false
  dt_davis_security_recommendations_segment:
    enabled: false
  # IAM
  dt_iam_groups_list:
    enabled: false
  dt_iam_policies_list:
    enabled: false
  dt_iam_user_get:
    enabled: false
  dt_iam_user_create:
    enabled: false
  dt_iam_user_update:
    enabled: false
  dt_iam_user_delete:
    enabled: false
  dt_iam_service_users_list:
    enabled: false
  dt_iam_group_get:
    enabled: false
  dt_iam_group_create:
    enabled: false
  dt_iam_group_update:
    enabled: false
  dt_iam_group_delete:
    enabled: false
  dt_environment_license_settings:
    enabled: false
  dt_all_user_app_states_delete:
    enabled: false
  # Platform
  dt_environment_info:
    enabled: false
  dt_app_state_get:
    enabled: false
  dt_apps_list:
    enabled: false
  dt_email_send:
    enabled: false
```

---

## API Token Scopes

The minimum token scopes depend on which tools you enable. Below are the scopes for each domain:

| Domain | Required scopes |
|---|---|
| DQL | `storage:query:read` |
| Automation / Workflows | `automation:workflows:read`, `automation:workflows:write` |
| Documents | `document:documents:read`, `document:documents:write` |
| Davis AI / Copilot | `davis:analyzers:read`, `davis:analyzers:execute` |
| SLOs | `slo:slos:read`, `slo:slos:write` |
| OpenPipeline | `openpipeline:configurations:read`, `openpipeline:configurations:write` |
| Notifications | `app-engine:apps:run` (Notifications API) |
| Storage / Buckets | `storage:buckets:read`, `storage:buckets:write`, `storage:filter-segments:read`, `storage:filter-segments:write` |
| Hub | `hub:catalog:read` |
| Vulnerabilities | `securityPosture:findings:read`, `securityPosture:findings:write` |
| IAM | `iam:users:read`, `iam:users:write`, `iam:groups:read`, `iam:groups:write` |
| Email | `app-engine:apps:run` |

---

## Project Structure

```
.
├── cmd/server/main.go              # Entry point
├── config.yaml                     # Tool enable/disable configuration
├── internal/
│   ├── client/client.go            # HTTP client (Bearer auth, multipart support)
│   ├── config/config.go            # Config loading + IsEnabled()
│   ├── server/server.go            # Tool registration
│   └── tools/
│       ├── handlers.go             # Shared handler utilities
│       ├── dql.go                  # DQL tools
│       ├── base_tools.go           # Automation, Documents, Davis, Storage, Hub, IAM, Platform
│       ├── slo_tools.go            # SLO tools
│       ├── automation_extended.go  # Extended automation tools
│       ├── documents_extended.go   # Extended document tools
│       ├── davis_extended.go       # Davis AI / Copilot tools
│       ├── storage_extended.go     # Extended storage tools
│       ├── hub_extended.go         # Extended Hub tools
│       ├── vulnerabilities_extended.go
│       ├── openpipeline_extended.go
│       ├── notifications_extended.go
│       └── iam_extended.go
└── go.mod
```
