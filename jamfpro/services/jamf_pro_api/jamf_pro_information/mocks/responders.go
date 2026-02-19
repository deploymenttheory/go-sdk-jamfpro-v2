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

type JamfProInformationMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewJamfProInformationMock() *JamfProInformationMock {
	return &JamfProInformationMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *JamfProInformationMock) register(method, path string, statusCode int, fixture string) {
	body, _ := os.ReadFile(filepath.Join(mustGetwd(), "mocks", fixture))
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *JamfProInformationMock) RegisterMocks() {
	m.register("GET", "/api/v2/jamf-pro-information", 200, "validate_get.json")
}

func (m *JamfProInformationMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *JamfProInformationMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *JamfProInformationMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *JamfProInformationMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *JamfProInformationMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *JamfProInformationMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *JamfProInformationMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *JamfProInformationMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *JamfProInformationMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *JamfProInformationMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *JamfProInformationMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil && len(resp.Body) > 0 {
		_ = mergePage(resp.Body)
	}
	return resp, nil
}
func (m *JamfProInformationMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *JamfProInformationMock) InvalidateToken() error                    { return nil }
func (m *JamfProInformationMock) KeepAliveToken() error                     { return nil }
func (m *JamfProInformationMock) GetLogger() *zap.Logger                     { return m.logger }

func (m *JamfProInformationMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{StatusCode: 404, Headers: http.Header{}, Body: nil}, fmt.Errorf("no mock for %s %s", method, path)
	}
	resp := &interfaces.Response{StatusCode: r.statusCode, Status: fmt.Sprintf("%d", r.statusCode), Headers: http.Header{"Content-Type": {"application/json"}}, Body: r.rawBody}
	if result != nil && len(r.rawBody) > 0 {
		_ = json.Unmarshal(r.rawBody, result)
	}
	return resp, nil
}

func mustGetwd() string {
	dir, _ := os.Getwd()
	return dir
}
