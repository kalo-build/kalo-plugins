package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Client is a typed HTTP client for the API.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient creates a new API client.
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: http.DefaultClient,
	}
}

// apiError represents an API error response.
type apiError struct {
	StatusCode int
	Body       string
}

func (e *apiError) Error() string {
	return fmt.Sprintf("API error %d: %s", e.StatusCode, e.Body)
}

// MetaRoot calls GET /meta/root.
func (c *Client) MetaRoot(ctx context.Context) (ApiOverview, error) {
	reqURL := c.BaseURL + "/meta/root"
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		var zero ApiOverview
		return zero, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		var zero ApiOverview
		return zero, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		var zero ApiOverview
		return zero, &apiError{StatusCode: resp.StatusCode, Body: string(respBody)}
	}

	var result ApiOverview
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		var zero ApiOverview
		return zero, fmt.Errorf("failed to decode response: %w", err)
	}
	return result, nil
}

// ActionsCacheList calls GET /orgs/{org-name}/cache-usage.
func (c *Client) ActionsCacheList(ctx context.Context, orgName string, p2fa string) (Status, error) {
	path := fmt.Sprintf("/orgs/%v/cache-usage", orgName)
	reqURL := c.BaseURL + path
	query := url.Values{}
	if p2fa != "" {
		query.Set("2fa", p2fa)
	}
	if len(query) > 0 {
		reqURL += "?" + query.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		var zero Status
		return zero, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		var zero Status
		return zero, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		var zero Status
		return zero, &apiError{StatusCode: resp.StatusCode, Body: string(respBody)}
	}

	var result Status
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		var zero Status
		return zero, fmt.Errorf("failed to decode response: %w", err)
	}
	return result, nil
}

// ListPets calls GET /pets.
func (c *Client) ListPets(ctx context.Context, limit string, new_ string) (Pets, error) {
	reqURL := c.BaseURL + "/pets"
	query := url.Values{}
	if limit != "" {
		query.Set("limit", limit)
	}
	if new_ != "" {
		query.Set("new", new_)
	}
	if len(query) > 0 {
		reqURL += "?" + query.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		var zero Pets
		return zero, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		var zero Pets
		return zero, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		var zero Pets
		return zero, &apiError{StatusCode: resp.StatusCode, Body: string(respBody)}
	}

	var result Pets
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		var zero Pets
		return zero, fmt.Errorf("failed to decode response: %w", err)
	}
	return result, nil
}

// GetStatus calls GET /status.
func (c *Client) GetStatus(ctx context.Context) (Status, error) {
	reqURL := c.BaseURL + "/status"
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		var zero Status
		return zero, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		var zero Status
		return zero, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		var zero Status
		return zero, &apiError{StatusCode: resp.StatusCode, Body: string(respBody)}
	}

	var result Status
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		var zero Status
		return zero, fmt.Errorf("failed to decode response: %w", err)
	}
	return result, nil
}

// CreateStatus calls POST /status.
func (c *Client) CreateStatus(ctx context.Context, body ApiOverview) (Status, error) {
	reqURL := c.BaseURL + "/status"
	var bodyReader io.Reader
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		var zero Status
		return zero, fmt.Errorf("failed to marshal request body: %w", err)
	}
	bodyReader = bytes.NewReader(bodyBytes)
	req, err := http.NewRequestWithContext(ctx, "POST", reqURL, bodyReader)
	if err != nil {
		var zero Status
		return zero, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		var zero Status
		return zero, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		var zero Status
		return zero, &apiError{StatusCode: resp.StatusCode, Body: string(respBody)}
	}

	var result Status
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		var zero Status
		return zero, fmt.Errorf("failed to decode response: %w", err)
	}
	return result, nil
}
