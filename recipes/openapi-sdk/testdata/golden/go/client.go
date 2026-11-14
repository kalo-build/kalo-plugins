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

// ListPets calls GET /pets.
func (c *Client) ListPets(ctx context.Context, limit string) (Pets, error) {
	reqURL := c.BaseURL + "/pets"
	query := url.Values{}
	if limit != "" {
		query.Set("limit", limit)
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

// CreatePets calls POST /pets.
func (c *Client) CreatePets(ctx context.Context, body Pet) error {
	reqURL := c.BaseURL + "/pets"
	var bodyReader io.Reader
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}
	bodyReader = bytes.NewReader(bodyBytes)
	req, err := http.NewRequestWithContext(ctx, "POST", reqURL, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
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

// ShowPetById calls GET /pets/{petId}.
func (c *Client) ShowPetById(ctx context.Context, petId string) (Pet, error) {
	path := fmt.Sprintf("/pets/%v", petId)
	reqURL := c.BaseURL + path
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		var zero Pet
		return zero, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		var zero Pet
		return zero, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		var zero Pet
		return zero, &apiError{StatusCode: resp.StatusCode, Body: string(respBody)}
	}

	var result Pet
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		var zero Pet
		return zero, fmt.Errorf("failed to decode response: %w", err)
	}
	return result, nil
}
