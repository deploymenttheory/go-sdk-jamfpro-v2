package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
}

// ServiceDiscoveryEnrollmentMock is a test double implementing interfaces.HTTPClient.
type ServiceDiscoveryEnrollmentMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewServiceDiscoveryEnrollmentMock returns an empty mock ready for response registration.
func NewServiceDiscoveryEnrollmentMock() *ServiceDiscoveryEnrollmentMock {
	return &ServiceDiscoveryEnrollmentMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *ServiceDiscoveryEnrollmentMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		body, _ = os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// RegisterGetWellKnownSettingsMock registers a successful GET well-known-settings response.
func (m *ServiceDiscoveryEnrollmentMock) RegisterGetWellKnownSettingsMock() {
	m.register("GET", "/api/v1/service-discovery-enrollment/well-known-settings", 200, "validate_get.json")
}

// RegisterUpdateWellKnownSettingsMock registers a successful PUT well-known-settings response (204 No Content).
func (m *ServiceDiscoveryEnrollmentMock) RegisterUpdateWellKnownSettingsMock() {
	m.register("PUT", "/api/v1/service-discovery-enrollment/well-known-settings", 204, "validate_update_204.json")
}

func (m *ServiceDiscoveryEnrollmentMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{StatusCode: 404, Headers: http.Header{}, Body: nil}, fmt.Errorf("ServiceDiscoveryEnrollmentMock: no response for %s %s", method, path)
	}
	resp := &interfaces.Response{StatusCode: r.statusCode, Status: fmt.Sprintf("%d", r.statusCode), Headers: http.Header{"Content-Type": {"application/json"}}, Body: r.rawBody}
	if result != nil && len(r.rawBody) > 0 {
		_ = json.Unmarshal(r.rawBody, result)
	}
	return resp, nil
}

func mustMocksDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func (m *ServiceDiscoveryEnrollmentMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *ServiceDiscoveryEnrollmentMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ServiceDiscoveryEnrollmentMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ServiceDiscoveryEnrollmentMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ServiceDiscoveryEnrollmentMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ServiceDiscoveryEnrollmentMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *ServiceDiscoveryEnrollmentMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *ServiceDiscoveryEnrollmentMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *ServiceDiscoveryEnrollmentMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *ServiceDiscoveryEnrollmentMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *ServiceDiscoveryEnrollmentMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil && len(resp.Body) > 0 {
		_ = mergePage(resp.Body)
	}
	return resp, nil
}
func (m *ServiceDiscoveryEnrollmentMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *ServiceDiscoveryEnrollmentMock) InvalidateToken() error                    { return nil }
func (m *ServiceDiscoveryEnrollmentMock) KeepAliveToken() error                      { return nil }
func (m *ServiceDiscoveryEnrollmentMock) GetLogger() *zap.Logger                      { return m.logger }
