package tests

import (
	"context"
	"fmt"
	"strings"

	"github.com/cucumber/godog"
)

func initializeScenario(ctx *godog.ScenarioContext) {
	sc := newScenarioCtx(globalMock)

	ctx.Before(func(ctx context.Context, sc2 *godog.Scenario) (context.Context, error) {
		sc.mock.Reset()
		sc.lastErr = nil
		sc.configInfo = nil
		sc.entityList = nil
		sc.entityDetail = nil
		sc.serviceList = nil
		sc.controlResult = nil
		sc.callResult = nil
		sc.client = sc.newClient()
		return ctx, nil
	})

	// --- Background ---
	ctx.Step(`^a running Home Assistant instance$`, sc.aRunningHAInstance)
	ctx.Step(`^an authenticated user$`, sc.anAuthenticatedUser)

	// --- Config ---
	ctx.Step(`^I request the server config$`, sc.iRequestServerConfig)
	ctx.Step(`^I should get a location name$`, sc.iShouldGetLocationName)
	ctx.Step(`^I should get a version$`, sc.iShouldGetVersion)
	ctx.Step(`^I should get a timezone$`, sc.iShouldGetTimezone)

	// --- Entities ---
	ctx.Step(`^entities exist in the system$`, sc.entitiesExistInSystem)
	ctx.Step(`^entity "([^"]*)" exists$`, sc.entityExists)
	ctx.Step(`^I list all entities$`, sc.iListAllEntities)
	ctx.Step(`^I list entities with domain "([^"]*)"$`, sc.iListEntitiesWithDomain)
	ctx.Step(`^I list entities with search "([^"]*)"$`, sc.iListEntitiesWithSearch)
	ctx.Step(`^I should receive a list of entities$`, sc.iShouldReceiveEntityList)
	ctx.Step(`^each entity should have an ID and state$`, sc.eachEntityShouldHaveIDAndState)
	ctx.Step(`^all entities should belong to domain "([^"]*)"$`, sc.allEntitiesShouldBelongToDomain)
	ctx.Step(`^all entities should match search "([^"]*)"$`, sc.allEntitiesShouldMatchSearch)

	// --- State ---
	ctx.Step(`^I get the state of "([^"]*)"$`, sc.iGetStateOf)
	ctx.Step(`^I should get entity details$`, sc.iShouldGetEntityDetails)
	ctx.Step(`^the entity ID should be "([^"]*)"$`, sc.entityIDShouldBe)
	ctx.Step(`^the entity should have a state value$`, sc.entityShouldHaveStateValue)

	// --- Control ---
	ctx.Step(`^I turn on "([^"]*)"$`, sc.iTurnOn)
	ctx.Step(`^I turn off "([^"]*)"$`, sc.iTurnOff)
	ctx.Step(`^I toggle "([^"]*)"$`, sc.iToggle)
	ctx.Step(`^the action should succeed with action "([^"]*)"$`, sc.actionShouldSucceedWithAction)

	// --- Services ---
	ctx.Step(`^I list all services$`, sc.iListAllServices)
	ctx.Step(`^I list services with domain "([^"]*)"$`, sc.iListServicesWithDomain)
	ctx.Step(`^I should receive a list of services$`, sc.iShouldReceiveServiceList)
	ctx.Step(`^each service should have a domain and name$`, sc.eachServiceShouldHaveDomainAndName)
	ctx.Step(`^all services should belong to domain "([^"]*)"$`, sc.allServicesShouldBelongToDomain)

	// --- Call ---
	ctx.Step(`^I call service "([^"]*)" "([^"]*)" with data:$`, sc.iCallServiceWithData)
	ctx.Step(`^the service call should succeed$`, sc.serviceCallShouldSucceed)

	// --- Common ---
	ctx.Step(`^it should fail with "([^"]*)"$`, sc.itShouldFailWith)
}

// --- Background steps ---

func (sc *scenarioCtx) aRunningHAInstance() error {
	if sc.mock == nil {
		return fmt.Errorf("mock server is not running")
	}
	return nil
}

func (sc *scenarioCtx) anAuthenticatedUser() error {
	if sc.client == nil {
		return fmt.Errorf("no authenticated client")
	}
	return nil
}

func (sc *scenarioCtx) itShouldFailWith(expected string) error {
	if sc.lastErr == nil {
		return fmt.Errorf("expected an error containing %q but got none", expected)
	}
	if !strings.Contains(sc.lastErr.Error(), expected) {
		return fmt.Errorf("expected error containing %q, got: %s", expected, sc.lastErr.Error())
	}
	return nil
}
