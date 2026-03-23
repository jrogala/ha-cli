// Package ops contains business logic for Home Assistant operations.
// Functions return Go structs and errors — zero I/O, zero formatting.
package ops

import "github.com/jrogala/ha-cli/client"

// ServerInfo holds the Home Assistant server configuration.
type ServerInfo struct {
	LocationName string `json:"location_name"`
	Version      string `json:"version"`
	TimeZone     string `json:"time_zone"`
}

// GetConfig returns the Home Assistant server configuration.
func GetConfig(c *client.Client) (*ServerInfo, error) {
	cfg, err := c.GetConfig()
	if err != nil {
		return nil, err
	}
	return &ServerInfo{
		LocationName: cfg.LocationName,
		Version:      cfg.Version,
		TimeZone:     cfg.TimeZone,
	}, nil
}
