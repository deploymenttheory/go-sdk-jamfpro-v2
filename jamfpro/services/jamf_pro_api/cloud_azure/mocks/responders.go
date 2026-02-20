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

type CloudAzureMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewCloudAzureMock() *CloudAzureMock {
	return &CloudAzureMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *CloudAzureMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		body, _ = os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *CloudAzureMock) RegisterGetDefaultServerConfigurationMock() {
	m.register("GET", "/api/v1/cloud-azure/defaults/server-configuration", 200, "validate_default_server.json")
}

func (m *CloudAzureMock) RegisterGetByIDMock(id string) {
	m.register("GET", "/api/v1/cloud-azure/"+id, 200, "validate_get.json")
}

func (m *CloudAzureMock) RegisterListMock() {
	m.register("GET", "/api/v1/cloud-azure", 200, "validate_list.json")
}

func (m *CloudAzureMock) RegisterCreateMock() {
	m.register("POST", "/api/v1/cloud-azure", 201, "validate_create.json")
}

func (m *CloudAzureMock) RegisterUpdateByIDMock(id string) {
	m.register("PUT", "/api/v1/cloud-azure/"+id, 200, "validate_get.json")
}

func (m *CloudAzureMock) RegisterDeleteByIDMock(id string) {
	m.register("DELETE", "/api/v1/cloud-azure/"+id, 204, "")
}

func (m *CloudAzureMock) RegisterGetDefaultMappingsMock() {
	m.register("GET", "/api/v1/cloud-azure/defaults/mappings", 200, "validate_default_mappings.json")
}

func (m *CloudAzureMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{StatusCode: 404, Headers: http.Header{}, Body: nil}, fmt.Errorf("CloudAzureMock: no response for %s %s", method, path)
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

func (m *CloudAzureMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *CloudAzureMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CloudAzureMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CloudAzureMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CloudAzureMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CloudAzureMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *CloudAzureMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *CloudAzureMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *CloudAzureMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *CloudAzureMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *CloudAzureMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil && len(resp.Body) > 0 {
		_ = mergePage(resp.Body)
	}
	return resp, nil
}
func (m *CloudAzureMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *CloudAzureMock) InvalidateToken() error                     { return nil }
func (m *CloudAzureMock) KeepAliveToken() error                      { return nil }
func (m *CloudAzureMock) GetLogger() *zap.Logger                     { return m.logger }
