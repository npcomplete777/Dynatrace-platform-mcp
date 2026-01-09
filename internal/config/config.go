// Package config handles configuration loading for the Dynatrace Platform MCP server.
package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Config holds the server configuration.
type Config struct {
	BaseURL       string
	PlatformToken string
	Debug         bool
}

// Load loads configuration from environment variables.
func Load() (*Config, error) {
	cfg := &Config{
		BaseURL:       os.Getenv("DT_BASE_URL"),
		PlatformToken: os.Getenv("DT_PLATFORM_TOKEN"),
		Debug:         os.Getenv("DT_DEBUG") == "true",
	}

	// Fall back to alternative env var names
	if cfg.BaseURL == "" {
		cfg.BaseURL = os.Getenv("DYNATRACE_BASE_URL")
	}
	if cfg.PlatformToken == "" {
		cfg.PlatformToken = os.Getenv("DYNATRACE_PLATFORM_TOKEN")
	}

	return cfg, nil
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
