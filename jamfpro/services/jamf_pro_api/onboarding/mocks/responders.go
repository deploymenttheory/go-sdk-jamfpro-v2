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

type OnboardingMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewOnboardingMock() *OnboardingMock {
	return &OnboardingMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *OnboardingMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		body, _ = os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *OnboardingMock) RegisterGetMock() {
	m.register("GET", "/api/v1/onboarding", 200, "validate_get.json")
}

func (m *OnboardingMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{StatusCode: 404, Headers: http.Header{}, Body: nil}, fmt.Errorf("OnboardingMock: no response for %s %s", method, path)
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

func (m *OnboardingMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *OnboardingMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *OnboardingMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *OnboardingMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *OnboardingMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *OnboardingMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *OnboardingMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *OnboardingMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *OnboardingMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *OnboardingMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *OnboardingMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil && len(resp.Body) > 0 {
		_ = mergePage(resp.Body)
	}
	return resp, nil
}
func (m *OnboardingMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *OnboardingMock) InvalidateToken() error                    { return nil }
func (m *OnboardingMock) KeepAliveToken() error                      { return nil }
func (m *OnboardingMock) GetLogger() *zap.Logger                      { return m.logger }
