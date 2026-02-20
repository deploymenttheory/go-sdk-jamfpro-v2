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

type AdcsSettingsMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewAdcsSettingsMock() *AdcsSettingsMock {
	return &AdcsSettingsMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *AdcsSettingsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		body, _ = os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *AdcsSettingsMock) RegisterCreateMock() {
	m.register("POST", "/api/v1/pki/adcs-settings", 201, "validate_create.json")
}

func (m *AdcsSettingsMock) RegisterGetByIDMock(id string) {
	m.register("GET", "/api/v1/pki/adcs-settings/"+id, 200, "validate_get.json")
}

func (m *AdcsSettingsMock) RegisterUpdateByIDMock(id string) {
	m.register("PATCH", "/api/v1/pki/adcs-settings/"+id, 204, "")
}

func (m *AdcsSettingsMock) RegisterDeleteByIDMock(id string) {
	m.register("DELETE", "/api/v1/pki/adcs-settings/"+id, 204, "")
}

func (m *AdcsSettingsMock) RegisterValidateServerCertificateMock() {
	m.register("POST", "/api/v1/pki/adcs-settings/validate-certificate", 204, "")
}

func (m *AdcsSettingsMock) RegisterValidateClientCertificateMock() {
	m.register("POST", "/api/v1/pki/adcs-settings/validate-client-certificate", 204, "")
}

func (m *AdcsSettingsMock) RegisterGetDependenciesByIDMock(id string) {
	m.register("GET", "/api/v1/pki/adcs-settings/"+id+"/dependencies", 200, "validate_dependencies.json")
}

func (m *AdcsSettingsMock) RegisterGetHistoryByIDMock(id string) {
	m.register("GET", "/api/v1/pki/adcs-settings/"+id+"/history", 200, "validate_history.json")
}

func (m *AdcsSettingsMock) RegisterAddHistoryNoteMock(id string) {
	m.register("POST", "/api/v1/pki/adcs-settings/"+id+"/history", 201, "")
}

func (m *AdcsSettingsMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{StatusCode: 404, Headers: http.Header{}, Body: nil}, fmt.Errorf("AdcsSettingsMock: no response for %s %s", method, path)
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

func (m *AdcsSettingsMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *AdcsSettingsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *AdcsSettingsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *AdcsSettingsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *AdcsSettingsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *AdcsSettingsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *AdcsSettingsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *AdcsSettingsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *AdcsSettingsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *AdcsSettingsMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *AdcsSettingsMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil && len(resp.Body) > 0 {
		_ = mergePage(resp.Body)
	}
	return resp, nil
}
func (m *AdcsSettingsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *AdcsSettingsMock) InvalidateToken() error                     { return nil }
func (m *AdcsSettingsMock) KeepAliveToken() error                      { return nil }
func (m *AdcsSettingsMock) GetLogger() *zap.Logger                     { return m.logger }
