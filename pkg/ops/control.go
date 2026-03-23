package ops

import "github.com/jrogala/ha-cli/client"

// ControlResult holds the result of a turn on/off/toggle operation.
type ControlResult struct {
	EntityID string `json:"entity_id"`
	Action   string `json:"action"`
}

// TurnOn turns on an entity.
func TurnOn(c *client.Client, entityID string) (*ControlResult, error) {
	_, err := c.TurnOn(entityID)
	if err != nil {
		return nil, err
	}
	return &ControlResult{EntityID: entityID, Action: "on"}, nil
}

// TurnOff turns off an entity.
func TurnOff(c *client.Client, entityID string) (*ControlResult, error) {
	_, err := c.TurnOff(entityID)
	if err != nil {
		return nil, err
	}
	return &ControlResult{EntityID: entityID, Action: "off"}, nil
}

// Toggle toggles an entity.
func Toggle(c *client.Client, entityID string) (*ControlResult, error) {
	_, err := c.Toggle(entityID)
	if err != nil {
		return nil, err
	}
	return &ControlResult{EntityID: entityID, Action: "toggle"}, nil
}
