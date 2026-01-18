// Package client provides an HTTP client for Dynatrace Platform APIs.
package client

import (
	"bytes"
	"mime/multipart"
	"net/textproto"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/dynatrace/dynatrace-platform-mcp-server/internal/config"
)

// Client is an HTTP client for Dynatrace Platform APIs.
type Client struct {
	cfg        *config.Config
	httpClient *http.Client
}

// Response wraps an HTTP response with helper methods.
type Response struct {
	StatusCode int
	Body       []byte // Public field for direct access
}

// IsSuccess returns true if the response status code is 2xx.
func (r *Response) IsSuccess() bool {
	return r.StatusCode >= 200 && r.StatusCode < 300
}

// JSON unmarshals the response body into the provided interface.
func (r *Response) JSON(v interface{}) error {
	return json.Unmarshal(r.Body, v)
}

// String returns the response body as a string.
func (r *Response) String() string {
	return string(r.Body)
}

// FormatError formats an error response for display.
func FormatError(r *Response) string {
	return fmt.Sprintf("API error %d: %s", r.StatusCode, string(r.Body))
}

// New creates a new Dynatrace API client.
func New(cfg *config.Config) *Client {
	return &Client{
		cfg: cfg,
		httpClient: &http.Client{
			Timeout: 120 * time.Second,
		},
	}
}

// RequestOption configures an HTTP request.
type RequestOption func(*http.Request)

// WithQueryParams adds query parameters to the request.
func WithQueryParams(params map[string]string) RequestOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		for k, v := range params {
			if v != "" {
				q.Set(k, v)
			}
		}
		req.URL.RawQuery = q.Encode()
	}
}

// WithQueryParamsMulti adds query parameters with multiple values.
func WithQueryParamsMulti(params map[string][]string) RequestOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		for k, values := range params {
			for _, v := range values {
				if v != "" {
					q.Add(k, v)
				}
			}
		}
		req.URL.RawQuery = q.Encode()
	}
}

// WithHeader adds a header to the request.
func WithHeader(key, value string) RequestOption {
	return func(req *http.Request) {
		req.Header.Set(key, value)
	}
}

// Get performs a GET request.
func (c *Client) Get(ctx context.Context, path string, opts ...RequestOption) (*Response, error) {
	return c.doRequest(ctx, http.MethodGet, path, nil, opts...)
}

// Post performs a POST request.
func (c *Client) Post(ctx context.Context, path string, body interface{}, opts ...RequestOption) (*Response, error) {
	return c.doRequest(ctx, http.MethodPost, path, body, opts...)
}

// Put performs a PUT request.
func (c *Client) Put(ctx context.Context, path string, body interface{}, opts ...RequestOption) (*Response, error) {
	return c.doRequest(ctx, http.MethodPut, path, body, opts...)
}

// Patch performs a PATCH request.
func (c *Client) Patch(ctx context.Context, path string, body interface{}, opts ...RequestOption) (*Response, error) {
	return c.doRequest(ctx, http.MethodPatch, path, body, opts...)
}

// Delete performs a DELETE request.
func (c *Client) Delete(ctx context.Context, path string, opts ...RequestOption) (*Response, error) {
	return c.doRequest(ctx, http.MethodDelete, path, nil, opts...)
}

// GetConfig returns the client configuration.
func (c *Client) GetConfig() *config.Config {
	return c.cfg
}

func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}, opts ...RequestOption) (*Response, error) {
	// Use apps URL for platform APIs
	baseURL := c.cfg.GetAppsURL()
	fullURL := baseURL + path

	var bodyReader io.Reader
	if body != nil {
		switch v := body.(type) {
		case string:
			bodyReader = strings.NewReader(v)
		case []byte:
			bodyReader = bytes.NewReader(v)
		default:
			jsonBody, err := json.Marshal(body)
			if err != nil {
				return nil, fmt.Errorf("marshal request body: %w", err)
			}
			bodyReader = bytes.NewReader(jsonBody)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+c.cfg.PlatformToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Apply options
	for _, opt := range opts {
		opt(req)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	return &Response{
		StatusCode: resp.StatusCode,
		Body:       respBody,
	}, nil
}


// PostMultipart performs a POST with multipart/form-data (required for Documents API)
func (c *Client) PostMultipart(ctx context.Context, path string, fields map[string]string) (*Response, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	for key, val := range fields {
		if key == "content" {
			part, _ := writer.CreatePart(textproto.MIMEHeader{
				"Content-Disposition": []string{`form-data; name="content"; filename="content.json"`},
				"Content-Type":        []string{"application/json"},
			})
			part.Write([]byte(val))
		} else {
			writer.WriteField(key, val)
		}
	}
	writer.Close()
	baseURL := c.cfg.GetAppsURL()
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+path, &buf)
	req.Header.Set("Authorization", "Bearer "+c.cfg.PlatformToken)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	return &Response{StatusCode: resp.StatusCode, Body: respBody}, nil
}
