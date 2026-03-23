package tests

import (
	"fmt"

	"github.com/jrogala/ha-cli/pkg/ops"
)

func (sc *scenarioCtx) iTurnOn(entityID string) error {
	// Mock the service call endpoint
	domain := entityDomain(entityID)
	sc.mock.On("POST", "/api/services/"+domain+"/turn_on", 200, []map[string]any{})

	result, err := ops.TurnOn(sc.client, entityID)
	sc.lastErr = err
	sc.controlResult = result
	return nil
}

func (sc *scenarioCtx) iTurnOff(entityID string) error {
	domain := entityDomain(entityID)
	sc.mock.On("POST", "/api/services/"+domain+"/turn_off", 200, []map[string]any{})

	result, err := ops.TurnOff(sc.client, entityID)
	sc.lastErr = err
	sc.controlResult = result
	return nil
}

func (sc *scenarioCtx) iToggle(entityID string) error {
	domain := entityDomain(entityID)
	sc.mock.On("POST", "/api/services/"+domain+"/toggle", 200, []map[string]any{})

	result, err := ops.Toggle(sc.client, entityID)
	sc.lastErr = err
	sc.controlResult = result
	return nil
}

func (sc *scenarioCtx) actionShouldSucceedWithAction(action string) error {
	if sc.lastErr != nil {
		return fmt.Errorf("expected success, got error: %v", sc.lastErr)
	}
	result, ok := sc.controlResult.(*ops.ControlResult)
	if !ok || result == nil {
		return fmt.Errorf("no control result available")
	}
	if result.Action != action {
		return fmt.Errorf("expected action %q, got %q", action, result.Action)
	}
	return nil
}

func entityDomain(entityID string) string {
	for i, c := range entityID {
		if c == '.' {
			return entityID[:i]
		}
	}
	return "homeassistant"
}
