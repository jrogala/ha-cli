package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
)

// routeKey uniquely identifies a mock route.
type routeKey struct {
	Method string
	Path   string
}

// mockRoute holds the response for a mocked endpoint.
type mockRoute struct {
	Status int
	Body   any
}

// MockServer is a test HTTP server that returns pre-configured responses.
type MockServer struct {
	Server *httptest.Server
	mu     sync.Mutex
	routes map[routeKey]mockRoute
}

// NewMockServer creates and starts a new mock HTTP server.
func NewMockServer() *MockServer {
	ms := &MockServer{
		routes: make(map[routeKey]mockRoute),
	}
	ms.Server = httptest.NewServer(http.HandlerFunc(ms.handler))
	return ms
}

// On registers a response for a given method and path.
func (ms *MockServer) On(method, path string, status int, body any) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.routes[routeKey{Method: method, Path: path}] = mockRoute{Status: status, Body: body}
}

// Reset clears all registered routes.
func (ms *MockServer) Reset() {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.routes = make(map[routeKey]mockRoute)
}

// Close shuts down the mock server.
func (ms *MockServer) Close() {
	ms.Server.Close()
}

// URL returns the base URL of the mock server.
func (ms *MockServer) URL() string {
	return ms.Server.URL
}

func (ms *MockServer) handler(w http.ResponseWriter, r *http.Request) {
	ms.mu.Lock()
	route, ok := ms.routes[routeKey{Method: r.Method, Path: r.URL.Path}]
	ms.mu.Unlock()

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "not found"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(route.Status)
	if route.Body != nil {
		json.NewEncoder(w).Encode(route.Body)
	}
}
