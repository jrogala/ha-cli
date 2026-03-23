package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Client communicates with the Home Assistant REST API.
type Client struct {
	baseURL string
	token   string
}

// New creates a new Home Assistant client.
func New(baseURL, token string) *Client {
	return &Client{
		baseURL: strings.TrimRight(baseURL, "/"),
		token:   token,
	}
}

func (c *Client) do(method, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, c.baseURL+path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		data, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("ha error %d: %s", resp.StatusCode, string(data))
	}
	return resp, nil
}

func (c *Client) doJSON(method, path string, body io.Reader, result any) error {
	resp, err := c.do(method, path, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}
	return nil
}

// --- Config ---

func (c *Client) GetConfig() (*Config, error) {
	var cfg Config
	if err := c.doJSON("GET", "/api/config", nil, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// --- States ---

func (c *Client) GetStates() ([]Entity, error) {
	var entities []Entity
	if err := c.doJSON("GET", "/api/states", nil, &entities); err != nil {
		return nil, err
	}
	return entities, nil
}

func (c *Client) GetState(entityID string) (*Entity, error) {
	var entity Entity
	if err := c.doJSON("GET", "/api/states/"+entityID, nil, &entity); err != nil {
		return nil, err
	}
	return &entity, nil
}

// --- Services ---

func (c *Client) GetServices() (map[string]map[string]any, error) {
	var services []struct {
		Domain   string                    `json:"domain"`
		Services map[string]any            `json:"services"`
	}
	if err := c.doJSON("GET", "/api/services", nil, &services); err != nil {
		return nil, err
	}
	result := make(map[string]map[string]any)
	for _, s := range services {
		result[s.Domain] = s.Services
	}
	return result, nil
}

func (c *Client) CallService(domain, service string, data map[string]any) ([]Entity, error) {
	body, _ := json.Marshal(data)
	var entities []Entity
	if err := c.doJSON("POST", "/api/services/"+domain+"/"+service, strings.NewReader(string(body)), &entities); err != nil {
		return nil, err
	}
	return entities, nil
}

// --- Convenience ---

func (c *Client) TurnOn(entityID string) ([]Entity, error) {
	domain := entityDomain(entityID)
	return c.CallService(domain, "turn_on", map[string]any{"entity_id": entityID})
}

func (c *Client) TurnOff(entityID string) ([]Entity, error) {
	domain := entityDomain(entityID)
	return c.CallService(domain, "turn_off", map[string]any{"entity_id": entityID})
}

func (c *Client) Toggle(entityID string) ([]Entity, error) {
	domain := entityDomain(entityID)
	return c.CallService(domain, "toggle", map[string]any{"entity_id": entityID})
}

func entityDomain(entityID string) string {
	if i := strings.Index(entityID, "."); i >= 0 {
		return entityID[:i]
	}
	return "homeassistant"
}
