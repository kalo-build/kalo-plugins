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

// ListNotes calls GET /notes.
func (c *Client) ListNotes(ctx context.Context) ([]Note, error) {
	reqURL := c.BaseURL + "/notes"
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		var zero []Note
		return zero, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		var zero []Note
		return zero, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		var zero []Note
		return zero, &apiError{StatusCode: resp.StatusCode, Body: string(respBody)}
	}

	var result []Note
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		var zero []Note
		return zero, fmt.Errorf("failed to decode response: %w", err)
	}
	return result, nil
}

// CreateNote calls POST /notes.
func (c *Client) CreateNote(ctx context.Context, body NewNote) (Note, error) {
	reqURL := c.BaseURL + "/notes"
	var bodyReader io.Reader
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		var zero Note
		return zero, fmt.Errorf("failed to marshal request body: %w", err)
	}
	bodyReader = bytes.NewReader(bodyBytes)
	req, err := http.NewRequestWithContext(ctx, "POST", reqURL, bodyReader)
	if err != nil {
		var zero Note
		return zero, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		var zero Note
		return zero, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		var zero Note
		return zero, &apiError{StatusCode: resp.StatusCode, Body: string(respBody)}
	}

	var result Note
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		var zero Note
		return zero, fmt.Errorf("failed to decode response: %w", err)
	}
	return result, nil
}

// GetLatestNote calls GET /notes/latest.
func (c *Client) GetLatestNote(ctx context.Context) (Note, error) {
	reqURL := c.BaseURL + "/notes/latest"
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		var zero Note
		return zero, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		var zero Note
		return zero, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		var zero Note
		return zero, &apiError{StatusCode: resp.StatusCode, Body: string(respBody)}
	}

	var result struct {
		Data Note `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		var zero Note
		return zero, fmt.Errorf("failed to decode response: %w", err)
	}
	return result.Data, nil
}

// ListNotesPage calls GET /notes/page.
func (c *Client) ListNotesPage(ctx context.Context) ([]Note, error) {
	reqURL := c.BaseURL + "/notes/page"
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		var zero []Note
		return zero, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		var zero []Note
		return zero, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		var zero []Note
		return zero, &apiError{StatusCode: resp.StatusCode, Body: string(respBody)}
	}

	var result struct {
		Data []Note `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		var zero []Note
		return zero, fmt.Errorf("failed to decode response: %w", err)
	}
	return result.Data, nil
}

// GetStats calls GET /notes/stats.
func (c *Client) GetStats(ctx context.Context) error {
	reqURL := c.BaseURL + "/notes/stats"
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
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

// ListSummaries calls GET /notes/summaries.
func (c *Client) ListSummaries(ctx context.Context) error {
	reqURL := c.BaseURL + "/notes/summaries"
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
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

// ListTags calls GET /notes/tags.
func (c *Client) ListTags(ctx context.Context) error {
	reqURL := c.BaseURL + "/notes/tags"
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
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

// GetNote calls GET /notes/{noteId}.
func (c *Client) GetNote(ctx context.Context, noteId string) (Note, error) {
	path := fmt.Sprintf("/notes/%v", noteId)
	reqURL := c.BaseURL + path
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		var zero Note
		return zero, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		var zero Note
		return zero, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		var zero Note
		return zero, &apiError{StatusCode: resp.StatusCode, Body: string(respBody)}
	}

	var result Note
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		var zero Note
		return zero, fmt.Errorf("failed to decode response: %w", err)
	}
	return result, nil
}
