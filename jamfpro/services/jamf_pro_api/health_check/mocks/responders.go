package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"resty.dev/v3"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// HealthCheckMock implements interfaces.HTTPClient for tests.
type HealthCheckMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewHealthCheckMock() *HealthCheckMock {
	return &HealthCheckMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *HealthCheckMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("HealthCheckMock: load %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *HealthCheckMock) RegisterMocks() {
	m.register("GET", "/api/v1/health-check", 200, "validate_get.json")
	m.register("GET", "/api/v1/health-status", 200, "validate_health_status.json")
}

func (m *HealthCheckMock) RegisterErrorMock() {
	m.responses["GET:/api/v1/health-check"] = registeredResponse{
		statusCode: 503,
		rawBody:    []byte(`{"code":"SERVICE_UNAVAILABLE","message":"service unavailable"}`),
		errMsg:     "Jamf Pro API error (503): service unavailable",
	}
}

func (m *HealthCheckMock) Get(ctx context.Context, path string, q map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *HealthCheckMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *HealthCheckMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *HealthCheckMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *HealthCheckMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *HealthCheckMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *HealthCheckMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *HealthCheckMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *HealthCheckMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *HealthCheckMock) GetBytes(ctx context.Context, path string, q map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}
func (m *HealthCheckMock) GetPaginated(ctx context.Context, path string, q map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	bodyBytes := resp.Bytes()
	if mergePage != nil && len(bodyBytes) > 0 {
		if err := mergePage(bodyBytes); err != nil {
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
	}
	return resp, nil
}
func (m *HealthCheckMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *HealthCheckMock) InvalidateToken() error                    { return nil }
func (m *HealthCheckMock) KeepAliveToken() error                     { return nil }
func (m *HealthCheckMock) GetLogger() *zap.Logger                    { return m.logger }

func (m *HealthCheckMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return shared.NewMockResponse(404, http.Header{}, nil), fmt.Errorf("HealthCheckMock: no response for %s %s", method, path)
	}
	headers := http.Header{"Content-Type": {"application/json"}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)
	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}
	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, err
		}
	}
	return resp, nil
}

func loadMockResponse(filename string) ([]byte, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return os.ReadFile(filepath.Join(dir, "mocks", filename))
}
