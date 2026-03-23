package ops

import (
	"strings"

	"github.com/jrogala/ha-cli/client"
)

// EntityEntry represents an entity in listing results.
type EntityEntry struct {
	EntityID string `json:"entity_id"`
	State    string `json:"state"`
	Name     string `json:"name"`
}

// EntityDetail holds detailed information about a single entity.
type EntityDetail struct {
	EntityID    string         `json:"entity_id"`
	Name        string         `json:"name"`
	State       string         `json:"state"`
	LastUpdated string         `json:"last_updated"`
	Attributes  map[string]any `json:"attributes"`
}

// ListOptions configures entity listing.
type ListOptions struct {
	Domain string
	Search string
}

// ListEntities returns entities, optionally filtered by domain and search term.
func ListEntities(c *client.Client, opts ListOptions) ([]EntityEntry, error) {
	entities, err := c.GetStates()
	if err != nil {
		return nil, err
	}

	var entries []EntityEntry
	for _, e := range entities {
		if opts.Domain != "" && e.Domain() != opts.Domain {
			continue
		}
		if opts.Search != "" && !strings.Contains(
			strings.ToLower(e.EntityID+e.FriendlyName()),
			strings.ToLower(opts.Search),
		) {
			continue
		}
		entries = append(entries, EntityEntry{
			EntityID: e.EntityID,
			State:    e.State,
			Name:     e.FriendlyName(),
		})
	}
	return entries, nil
}

// GetState returns detailed information about a single entity.
func GetState(c *client.Client, entityID string) (*EntityDetail, error) {
	entity, err := c.GetState(entityID)
	if err != nil {
		return nil, err
	}
	return &EntityDetail{
		EntityID:    entity.EntityID,
		Name:        entity.FriendlyName(),
		State:       entity.State,
		LastUpdated: entity.LastUpdated,
		Attributes:  entity.Attributes,
	}, nil
}
