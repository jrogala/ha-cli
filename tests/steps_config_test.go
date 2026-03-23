package tests

import (
	"fmt"

	"github.com/jrogala/ha-cli/pkg/ops"
)

func (sc *scenarioCtx) iRequestServerConfig() error {
	sc.mock.On("GET", "/api/config", 200, map[string]any{
		"location_name": "Home",
		"version":       "2024.1.0",
		"time_zone":     "Europe/Paris",
		"components":    []string{"light", "switch"},
		"latitude":      48.8566,
		"longitude":     2.3522,
		"unit_system":   map[string]string{"temperature": "°C"},
	})

	info, err := ops.GetConfig(sc.client)
	sc.lastErr = err
	sc.configInfo = info
	return nil
}

func (sc *scenarioCtx) iShouldGetLocationName() error {
	info, ok := sc.configInfo.(*ops.ServerInfo)
	if !ok || info == nil {
		return fmt.Errorf("no config info available")
	}
	if info.LocationName == "" {
		return fmt.Errorf("location name is empty")
	}
	return nil
}

func (sc *scenarioCtx) iShouldGetVersion() error {
	info, ok := sc.configInfo.(*ops.ServerInfo)
	if !ok || info == nil {
		return fmt.Errorf("no config info available")
	}
	if info.Version == "" {
		return fmt.Errorf("version is empty")
	}
	return nil
}

func (sc *scenarioCtx) iShouldGetTimezone() error {
	info, ok := sc.configInfo.(*ops.ServerInfo)
	if !ok || info == nil {
		return fmt.Errorf("no config info available")
	}
	if info.TimeZone == "" {
		return fmt.Errorf("timezone is empty")
	}
	return nil
}
