package tests

import (
	"fmt"
	"strings"

	"github.com/jrogala/ha-cli/pkg/ops"
)

var testEntities = []map[string]any{
	{
		"entity_id":    "light.kitchen",
		"state":        "on",
		"attributes":   map[string]any{"friendly_name": "Kitchen Light", "brightness": 255},
		"last_changed": "2024-01-01T00:00:00Z",
		"last_updated": "2024-01-01T00:00:00Z",
	},
	{
		"entity_id":    "light.bedroom",
		"state":        "off",
		"attributes":   map[string]any{"friendly_name": "Bedroom Light"},
		"last_changed": "2024-01-01T00:00:00Z",
		"last_updated": "2024-01-01T00:00:00Z",
	},
	{
		"entity_id":    "switch.garden",
		"state":        "on",
		"attributes":   map[string]any{"friendly_name": "Garden Switch"},
		"last_changed": "2024-01-01T00:00:00Z",
		"last_updated": "2024-01-01T00:00:00Z",
	},
	{
		"entity_id":    "sensor.temperature",
		"state":        "21.5",
		"attributes":   map[string]any{"friendly_name": "Temperature Sensor", "unit_of_measurement": "°C"},
		"last_changed": "2024-01-01T00:00:00Z",
		"last_updated": "2024-01-01T00:00:00Z",
	},
}

func (sc *scenarioCtx) entitiesExistInSystem() error {
	sc.mock.On("GET", "/api/states", 200, testEntities)
	return nil
}

func (sc *scenarioCtx) entityExists(entityID string) error {
	for _, e := range testEntities {
		if e["entity_id"] == entityID {
			sc.mock.On("GET", "/api/states/"+entityID, 200, e)
			return nil
		}
	}
	// Register the entity even if not in testEntities, with a default
	sc.mock.On("GET", "/api/states/"+entityID, 200, map[string]any{
		"entity_id":    entityID,
		"state":        "on",
		"attributes":   map[string]any{"friendly_name": entityID},
		"last_changed": "2024-01-01T00:00:00Z",
		"last_updated": "2024-01-01T00:00:00Z",
	})
	return nil
}

func (sc *scenarioCtx) iListAllEntities() error {
	entries, err := ops.ListEntities(sc.client, ops.ListOptions{})
	sc.lastErr = err
	sc.entityList = entries
	return nil
}

func (sc *scenarioCtx) iListEntitiesWithDomain(domain string) error {
	entries, err := ops.ListEntities(sc.client, ops.ListOptions{Domain: domain})
	sc.lastErr = err
	sc.entityList = entries
	return nil
}

func (sc *scenarioCtx) iListEntitiesWithSearch(search string) error {
	entries, err := ops.ListEntities(sc.client, ops.ListOptions{Search: search})
	sc.lastErr = err
	sc.entityList = entries
	return nil
}

func (sc *scenarioCtx) iShouldReceiveEntityList() error {
	entries, ok := sc.entityList.([]ops.EntityEntry)
	if !ok || entries == nil {
		return fmt.Errorf("no entity list available")
	}
	if len(entries) == 0 {
		return fmt.Errorf("entity list is empty")
	}
	return nil
}

func (sc *scenarioCtx) eachEntityShouldHaveIDAndState() error {
	entries, ok := sc.entityList.([]ops.EntityEntry)
	if !ok {
		return fmt.Errorf("no entity list available")
	}
	for _, e := range entries {
		if e.EntityID == "" {
			return fmt.Errorf("entity has empty ID")
		}
		if e.State == "" {
			return fmt.Errorf("entity %s has empty state", e.EntityID)
		}
	}
	return nil
}

func (sc *scenarioCtx) allEntitiesShouldBelongToDomain(domain string) error {
	entries, ok := sc.entityList.([]ops.EntityEntry)
	if !ok {
		return fmt.Errorf("no entity list available")
	}
	for _, e := range entries {
		if !strings.HasPrefix(e.EntityID, domain+".") {
			return fmt.Errorf("entity %s does not belong to domain %s", e.EntityID, domain)
		}
	}
	return nil
}

func (sc *scenarioCtx) allEntitiesShouldMatchSearch(search string) error {
	entries, ok := sc.entityList.([]ops.EntityEntry)
	if !ok {
		return fmt.Errorf("no entity list available")
	}
	for _, e := range entries {
		combined := strings.ToLower(e.EntityID + e.Name)
		if !strings.Contains(combined, strings.ToLower(search)) {
			return fmt.Errorf("entity %s does not match search %q", e.EntityID, search)
		}
	}
	return nil
}

// --- State ---

func (sc *scenarioCtx) iGetStateOf(entityID string) error {
	detail, err := ops.GetState(sc.client, entityID)
	sc.lastErr = err
	sc.entityDetail = detail
	return nil
}

func (sc *scenarioCtx) iShouldGetEntityDetails() error {
	if sc.entityDetail == nil {
		return fmt.Errorf("no entity detail available")
	}
	return nil
}

func (sc *scenarioCtx) entityIDShouldBe(expected string) error {
	detail, ok := sc.entityDetail.(*ops.EntityDetail)
	if !ok || detail == nil {
		return fmt.Errorf("no entity detail available")
	}
	if detail.EntityID != expected {
		return fmt.Errorf("expected entity ID %q, got %q", expected, detail.EntityID)
	}
	return nil
}

func (sc *scenarioCtx) entityShouldHaveStateValue() error {
	detail, ok := sc.entityDetail.(*ops.EntityDetail)
	if !ok || detail == nil {
		return fmt.Errorf("no entity detail available")
	}
	if detail.State == "" {
		return fmt.Errorf("entity state is empty")
	}
	return nil
}
