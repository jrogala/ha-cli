package client

// Entity represents a Home Assistant entity state.
type Entity struct {
	EntityID    string                 `json:"entity_id"`
	State       string                 `json:"state"`
	Attributes  map[string]any         `json:"attributes"`
	LastChanged string                 `json:"last_changed"`
	LastUpdated string                 `json:"last_updated"`
}

// Config represents the HA server config.
type Config struct {
	LocationName string   `json:"location_name"`
	Version      string   `json:"version"`
	Components   []string `json:"components"`
	Latitude     float64  `json:"latitude"`
	Longitude    float64  `json:"longitude"`
	TimeZone     string   `json:"time_zone"`
	UnitSystem   struct {
		Temperature string `json:"temperature"`
	} `json:"unit_system"`
}

// Service represents a HA service.
type Service struct {
	Domain   string
	Service  string
	Name     string
	Desc     string
}

// Area represents a HA area.
type Area struct {
	AreaID string `json:"area_id"`
	Name   string `json:"name"`
}

// Device represents a HA device.
type Device struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	AreaID       string `json:"area_id"`
}

// Automation represents a HA automation.
type Automation struct {
	EntityID string `json:"entity_id"`
	State    string `json:"state"`
	Alias    string
}

// FriendlyName returns the friendly name from attributes, or entity_id.
func (e *Entity) FriendlyName() string {
	if name, ok := e.Attributes["friendly_name"].(string); ok {
		return name
	}
	return e.EntityID
}

// Domain returns the domain part of the entity_id (e.g. "light" from "light.kitchen").
func (e *Entity) Domain() string {
	for i, c := range e.EntityID {
		if c == '.' {
			return e.EntityID[:i]
		}
	}
	return e.EntityID
}
