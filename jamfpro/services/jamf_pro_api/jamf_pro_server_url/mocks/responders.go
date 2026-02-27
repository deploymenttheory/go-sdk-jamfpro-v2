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

type JamfProServerURLMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewJamfProServerURLMock() *JamfProServerURLMock {
	return &JamfProServerURLMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func (m *JamfProServerURLMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		body, _ = os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *JamfProServerURLMock) RegisterGetMock() {
	m.register("GET", "/api/v1/jamf-pro-server-url", 200, "validate_get.json")
}

func (m *JamfProServerURLMock) RegisterUpdateMock() {
	m.register("PUT", "/api/v1/jamf-pro-server-url", 200, "validate_update.json")
}

func (m *JamfProServerURLMock) RegisterGetHistoryMock() {
	m.register("GET", "/api/v1/jamf-pro-server-url/history", 200, "validate_history.json")
}

func (m *JamfProServerURLMock) RegisterCreateHistoryNoteMock() {
	m.register("POST", "/api/v1/jamf-pro-server-url/history", 201, "validate_history_note.json")
}

var errNoMockRegistered = fmt.Errorf("JamfProServerURLMock: no response registered")

func (m *JamfProServerURLMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, errNoMockRegistered
	}
	resp := &interfaces.Response{
		StatusCode: r.statusCode,
		Status:     fmt.Sprintf("%d", r.statusCode),
		Headers:    http.Header{"Content-Type": {"application/json"}},
		Body:       r.rawBody,
	}
	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("JamfProServerURLMock: unmarshal: %w", err)
		}
	}
	return resp, nil
}

func mustMocksDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func (m *JamfProServerURLMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}

func (m *JamfProServerURLMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *JamfProServerURLMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *JamfProServerURLMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *JamfProServerURLMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *JamfProServerURLMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *JamfProServerURLMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *JamfProServerURLMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *JamfProServerURLMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *JamfProServerURLMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *JamfProServerURLMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil {
		_ = mergePage(resp.Body)
	}
	return resp, nil
}

func (m *JamfProServerURLMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *JamfProServerURLMock) InvalidateToken() error                     { return nil }
func (m *JamfProServerURLMock) KeepAliveToken() error                      { return nil }
func (m *JamfProServerURLMock) GetLogger() *zap.Logger                     { return m.logger }
