package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
}

// JamfProSystemInitializationMock is a test double implementing interfaces.HTTPClient.
type JamfProSystemInitializationMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewJamfProSystemInitializationMock returns an empty mock ready for response registration.
func NewJamfProSystemInitializationMock() *JamfProSystemInitializationMock {
	return &JamfProSystemInitializationMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func (m *JamfProSystemInitializationMock) register(method, path string, statusCode int, fixture string) {
	body, _ := os.ReadFile(filepath.Join(mustGetwd(), "mocks", fixture))
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// RegisterInitializeMock registers a successful POST /api/v1/system/initialize response.
func (m *JamfProSystemInitializationMock) RegisterInitializeMock() {
	m.register("POST", "/api/v1/system/initialize", 200, "validate_initialize.json")
}

// RegisterInitializeDatabaseConnectionMock registers a successful POST /api/v1/system/initialize-database-connection response.
func (m *JamfProSystemInitializationMock) RegisterInitializeDatabaseConnectionMock() {
	m.register("POST", "/api/v1/system/initialize-database-connection", 200, "validate_initialize_database_connection.json")
}

// RegisterPlatformInitializeMock registers a successful POST /api/v1/system/platform-initialize response.
func (m *JamfProSystemInitializationMock) RegisterPlatformInitializeMock() {
	m.register("POST", "/api/v1/system/platform-initialize", 201, "validate_platform_initialize.json")
}

func (m *JamfProSystemInitializationMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{StatusCode: 404, Headers: http.Header{}, Body: nil}, fmt.Errorf("JamfProSystemInitializationMock: no response for %s %s", method, path)
	}
	resp := &interfaces.Response{
		StatusCode: r.statusCode,
		Status:     fmt.Sprintf("%d", r.statusCode),
		Headers:    http.Header{"Content-Type": {"application/json"}},
		Body:       r.rawBody,
	}
	if result != nil && len(r.rawBody) > 0 {
		_ = json.Unmarshal(r.rawBody, result)
	}
	return resp, nil
}

func mustGetwd() string {
	dir, _ := os.Getwd()
	return dir
}

func (m *JamfProSystemInitializationMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *JamfProSystemInitializationMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *JamfProSystemInitializationMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *JamfProSystemInitializationMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *JamfProSystemInitializationMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *JamfProSystemInitializationMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *JamfProSystemInitializationMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *JamfProSystemInitializationMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *JamfProSystemInitializationMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *JamfProSystemInitializationMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *JamfProSystemInitializationMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil && len(resp.Body) > 0 {
		if err := mergePage(resp.Body); err != nil {
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
	}
	return resp, nil
}
func (m *JamfProSystemInitializationMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *JamfProSystemInitializationMock) InvalidateToken() error                    { return nil }
func (m *JamfProSystemInitializationMock) KeepAliveToken() error                      { return nil }
func (m *JamfProSystemInitializationMock) GetLogger() *zap.Logger                      { return m.logger }
