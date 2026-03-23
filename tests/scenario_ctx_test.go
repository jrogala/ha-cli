package tests

import (
	"github.com/jrogala/ha-cli/client"
)

// scenarioCtx holds per-scenario state.
type scenarioCtx struct {
	mock   *MockServer
	client *client.Client

	// results from the latest "When" step
	lastErr       error
	configInfo    any
	entityList    any
	entityDetail  any
	serviceList   any
	controlResult any
	callResult    any
}

func newScenarioCtx(mock *MockServer) *scenarioCtx {
	return &scenarioCtx{
		mock: mock,
	}
}

// newClient creates a client pointing at the mock server.
func (sc *scenarioCtx) newClient() *client.Client {
	return client.New(sc.mock.URL(), "test-token")
}
