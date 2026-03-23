package tests

import (
	"fmt"

	"github.com/cucumber/godog"
	"github.com/jrogala/ha-cli/pkg/ops"
)

var testServices = []map[string]any{
	{
		"domain": "light",
		"services": map[string]any{
			"turn_on":  map[string]any{},
			"turn_off": map[string]any{},
			"toggle":   map[string]any{},
		},
	},
	{
		"domain": "switch",
		"services": map[string]any{
			"turn_on":  map[string]any{},
			"turn_off": map[string]any{},
		},
	},
}

func (sc *scenarioCtx) iListAllServices() error {
	sc.mock.On("GET", "/api/services", 200, testServices)

	entries, err := ops.ListServices(sc.client, "")
	sc.lastErr = err
	sc.serviceList = entries
	return nil
}

func (sc *scenarioCtx) iListServicesWithDomain(domain string) error {
	sc.mock.On("GET", "/api/services", 200, testServices)

	entries, err := ops.ListServices(sc.client, domain)
	sc.lastErr = err
	sc.serviceList = entries
	return nil
}

func (sc *scenarioCtx) iShouldReceiveServiceList() error {
	entries, ok := sc.serviceList.([]ops.ServiceEntry)
	if !ok || entries == nil {
		return fmt.Errorf("no service list available")
	}
	if len(entries) == 0 {
		return fmt.Errorf("service list is empty")
	}
	return nil
}

func (sc *scenarioCtx) eachServiceShouldHaveDomainAndName() error {
	entries, ok := sc.serviceList.([]ops.ServiceEntry)
	if !ok {
		return fmt.Errorf("no service list available")
	}
	for _, e := range entries {
		if e.Domain == "" {
			return fmt.Errorf("service has empty domain")
		}
		if e.Service == "" {
			return fmt.Errorf("service has empty name")
		}
	}
	return nil
}

func (sc *scenarioCtx) allServicesShouldBelongToDomain(domain string) error {
	entries, ok := sc.serviceList.([]ops.ServiceEntry)
	if !ok {
		return fmt.Errorf("no service list available")
	}
	for _, e := range entries {
		if e.Domain != domain {
			return fmt.Errorf("service %s.%s does not belong to domain %s", e.Domain, e.Service, domain)
		}
	}
	return nil
}

func (sc *scenarioCtx) iCallServiceWithData(domain, service string, table *godog.Table) error {
	sc.mock.On("POST", "/api/services/"+domain+"/"+service, 200, []map[string]any{})

	// Parse table data
	data := make(map[string]any)
	if len(table.Rows) > 1 {
		headers := table.Rows[0]
		values := table.Rows[1]
		for i, cell := range headers.Cells {
			if i < len(values.Cells) {
				data[cell.Value] = values.Cells[i].Value
			}
		}
	}

	result, err := ops.CallService(sc.client, domain, service, data)
	sc.lastErr = err
	sc.callResult = result
	return nil
}

func (sc *scenarioCtx) serviceCallShouldSucceed() error {
	if sc.lastErr != nil {
		return fmt.Errorf("expected success, got error: %v", sc.lastErr)
	}
	result, ok := sc.callResult.(*ops.CallResult)
	if !ok || result == nil {
		return fmt.Errorf("no call result available")
	}
	if !result.Success {
		return fmt.Errorf("service call was not successful")
	}
	return nil
}
