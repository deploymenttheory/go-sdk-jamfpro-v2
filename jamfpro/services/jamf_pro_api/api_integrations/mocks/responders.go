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

type ApiIntegrationsMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewApiIntegrationsMock() *ApiIntegrationsMock {
	return &ApiIntegrationsMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *ApiIntegrationsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		body, _ = os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *ApiIntegrationsMock) RegisterListMock() {
	m.register("GET", "/api/v1/api-integrations", 200, "validate_list.json")
}

func (m *ApiIntegrationsMock) RegisterGetByIDMock(id string) {
	m.register("GET", "/api/v1/api-integrations/"+id, 200, "validate_get.json")
}

func (m *ApiIntegrationsMock) RegisterCreateMock() {
	m.register("POST", "/api/v1/api-integrations", 200, "validate_get.json")
}

func (m *ApiIntegrationsMock) RegisterUpdateByIDMock(id string) {
	m.register("PUT", "/api/v1/api-integrations/"+id, 200, "validate_get.json")
}

func (m *ApiIntegrationsMock) RegisterDeleteByIDMock(id string) {
	m.register("DELETE", "/api/v1/api-integrations/"+id, 200, "")
}

func (m *ApiIntegrationsMock) RegisterRefreshClientCredentialsMock(id string) {
	m.register("POST", "/api/v1/api-integrations/"+id+"/client-credentials", 200, "validate_client_credentials.json")
}

func (m *ApiIntegrationsMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{StatusCode: 404, Headers: http.Header{}, Body: nil}, fmt.Errorf("ApiIntegrationsMock: no response for %s %s", method, path)
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

func (m *ApiIntegrationsMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *ApiIntegrationsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ApiIntegrationsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ApiIntegrationsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ApiIntegrationsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ApiIntegrationsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *ApiIntegrationsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *ApiIntegrationsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *ApiIntegrationsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *ApiIntegrationsMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *ApiIntegrationsMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil && len(resp.Body) > 0 {
		_ = mergePage(resp.Body)
	}
	return resp, nil
}
func (m *ApiIntegrationsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *ApiIntegrationsMock) InvalidateToken() error                     { return nil }
func (m *ApiIntegrationsMock) KeepAliveToken() error                      { return nil }
func (m *ApiIntegrationsMock) GetLogger() *zap.Logger                     { return m.logger }
