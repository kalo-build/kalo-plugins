package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

// ListCompany calls GET /api/companies.
func (c *Client) ListCompany(ctx context.Context) ([]Company, error) {
	reqURL := c.BaseURL + "/api/companies"
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		var zero []Company
		return zero, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		var zero []Company
		return zero, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		var zero []Company
		return zero, &apiError{StatusCode: resp.StatusCode, Body: string(respBody)}
	}

	var result []Company
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		var zero []Company
		return zero, fmt.Errorf("failed to decode response: %w", err)
	}
	return result, nil
}

// CreateCompany calls POST /api/companies.
func (c *Client) CreateCompany(ctx context.Context, body CompanyCreate) (Company, error) {
	reqURL := c.BaseURL + "/api/companies"
	var bodyReader io.Reader
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		var zero Company
		return zero, fmt.Errorf("failed to marshal request body: %w", err)
	}
	bodyReader = bytes.NewReader(bodyBytes)
	req, err := http.NewRequestWithContext(ctx, "POST", reqURL, bodyReader)
	if err != nil {
		var zero Company
		return zero, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		var zero Company
		return zero, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		var zero Company
		return zero, &apiError{StatusCode: resp.StatusCode, Body: string(respBody)}
	}

	var result Company
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		var zero Company
		return zero, fmt.Errorf("failed to decode response: %w", err)
	}
	return result, nil
}

// GetCompany calls GET /api/companies/{id}.
func (c *Client) GetCompany(ctx context.Context, id string) (Company, error) {
	path := fmt.Sprintf("/api/companies/%v", id)
	reqURL := c.BaseURL + path
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		var zero Company
		return zero, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		var zero Company
		return zero, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		var zero Company
		return zero, &apiError{StatusCode: resp.StatusCode, Body: string(respBody)}
	}

	var result Company
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		var zero Company
		return zero, fmt.Errorf("failed to decode response: %w", err)
	}
	return result, nil
}

// UpdateCompany calls PUT /api/companies/{id}.
func (c *Client) UpdateCompany(ctx context.Context, id string, body CompanyUpdate) (Company, error) {
	path := fmt.Sprintf("/api/companies/%v", id)
	reqURL := c.BaseURL + path
	var bodyReader io.Reader
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		var zero Company
		return zero, fmt.Errorf("failed to marshal request body: %w", err)
	}
	bodyReader = bytes.NewReader(bodyBytes)
	req, err := http.NewRequestWithContext(ctx, "PUT", reqURL, bodyReader)
	if err != nil {
		var zero Company
		return zero, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		var zero Company
		return zero, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		var zero Company
		return zero, &apiError{StatusCode: resp.StatusCode, Body: string(respBody)}
	}

	var result Company
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		var zero Company
		return zero, fmt.Errorf("failed to decode response: %w", err)
	}
	return result, nil
}

// DeleteCompany calls DELETE /api/companies/{id}.
func (c *Client) DeleteCompany(ctx context.Context, id string) error {
	path := fmt.Sprintf("/api/companies/%v", id)
	reqURL := c.BaseURL + path
	req, err := http.NewRequestWithContext(ctx, "DELETE", reqURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		return &apiError{StatusCode: resp.StatusCode, Body: string(respBody)}
	}
	return nil
}
