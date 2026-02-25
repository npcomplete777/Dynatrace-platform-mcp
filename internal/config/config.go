// Package config handles configuration loading for the Dynatrace Platform MCP server.
package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// ToolConfig controls whether an individual MCP tool is enabled.
type ToolConfig struct {
	Enabled *bool `yaml:"enabled"`
}

// yamlConfig is the raw YAML file structure.
type yamlConfig struct {
	Tools map[string]ToolConfig `yaml:"tools"`
}

// Config holds the server configuration.
type Config struct {
	BaseURL       string
	PlatformToken string
	Debug         bool
	Tools         map[string]ToolConfig
}

// Load loads configuration from environment variables and an optional YAML config file.
// The config file path is read from DYNATRACE_CONFIG_FILE, defaulting to config.yaml.
func Load() (*Config, error) {
	cfg := &Config{
		BaseURL:       os.Getenv("DT_BASE_URL"),
		PlatformToken: os.Getenv("DT_PLATFORM_TOKEN"),
		Debug:         os.Getenv("DT_DEBUG") == "true",
		Tools:         make(map[string]ToolConfig),
	}

	// Fall back to alternative env var names
	if cfg.BaseURL == "" {
		cfg.BaseURL = os.Getenv("DYNATRACE_BASE_URL")
	}
	if cfg.PlatformToken == "" {
		cfg.PlatformToken = os.Getenv("DYNATRACE_PLATFORM_TOKEN")
	}

	// Load optional YAML config file
	configFile := os.Getenv("DYNATRACE_CONFIG_FILE")
	if configFile == "" {
		configFile = "config.yaml"
	}
	if err := cfg.loadYAML(configFile); err != nil {
		return nil, err
	}

	return cfg, nil
}

// loadYAML reads tool enable/disable settings from the given YAML file.
// Missing file is silently ignored — all tools default to enabled.
func (c *Config) loadYAML(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // no config file is fine
		}
		return fmt.Errorf("reading config file %q: %w", path, err)
	}
	var y yamlConfig
	if err := yaml.Unmarshal(data, &y); err != nil {
		return fmt.Errorf("parsing config file %q: %w", path, err)
	}
	if y.Tools != nil {
		c.Tools = y.Tools
	}
	return nil
}

// IsEnabled reports whether the named tool should be registered.
// Tools absent from the config file default to enabled.
func (c *Config) IsEnabled(name string) bool {
	tc, ok := c.Tools[name]
	if !ok {
		return true
	}
	if tc.Enabled == nil {
		return true
	}
	return *tc.Enabled
}

// Validate validates the configuration.
func (c *Config) Validate() error {
	if c.BaseURL == "" {
		return errors.New("DT_BASE_URL or DYNATRACE_BASE_URL environment variable is required")
	}
	if c.PlatformToken == "" {
		return errors.New("DT_PLATFORM_TOKEN or DYNATRACE_PLATFORM_TOKEN environment variable is required")
	}

	// Normalize base URL
	c.BaseURL = strings.TrimSuffix(c.BaseURL, "/")

	// Validate URL format
	if !strings.HasPrefix(c.BaseURL, "https://") {
		return fmt.Errorf("base URL must start with https://: %s", c.BaseURL)
	}

	return nil
}

// GetAppsURL returns the base URL for apps.dynatrace.com endpoints.
func (c *Config) GetAppsURL() string {
	// If already using apps subdomain, return as-is
	if strings.Contains(c.BaseURL, ".apps.dynatrace.com") {
		return c.BaseURL
	}
	// Convert live.dynatrace.com to apps.dynatrace.com
	return strings.Replace(c.BaseURL, ".live.dynatrace.com", ".apps.dynatrace.com", 1)
}

// GetAccountURN returns the account URN for IAM API calls.
func (c *Config) GetAccountURN() string {
	return os.Getenv("DT_ACCOUNT_URN")
}
