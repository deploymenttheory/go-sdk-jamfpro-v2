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

type ComputerInventoryCollectionSettingsMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewComputerInventoryCollectionSettingsMock() *ComputerInventoryCollectionSettingsMock {
	return &ComputerInventoryCollectionSettingsMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *ComputerInventoryCollectionSettingsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		body, _ = os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *ComputerInventoryCollectionSettingsMock) RegisterGetMock() {
	m.register("GET", "/api/v2/computer-inventory-collection-settings", 200, "validate_get.json")
}

func (m *ComputerInventoryCollectionSettingsMock) RegisterUpdateMock() {
	m.register("PATCH", "/api/v2/computer-inventory-collection-settings", 204, "")
}

func (m *ComputerInventoryCollectionSettingsMock) RegisterCreateCustomPathMock() {
	m.register("POST", "/api/v2/computer-inventory-collection-settings/custom-path", 201, "validate_create_path.json")
}

func (m *ComputerInventoryCollectionSettingsMock) RegisterDeleteCustomPathMock(id string) {
	m.register("DELETE", "/api/v2/computer-inventory-collection-settings/custom-path/"+id, 204, "")
}

func (m *ComputerInventoryCollectionSettingsMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{StatusCode: 404, Headers: http.Header{}, Body: nil}, fmt.Errorf("ComputerInventoryCollectionSettingsMock: no response for %s %s", method, path)
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

func (m *ComputerInventoryCollectionSettingsMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *ComputerInventoryCollectionSettingsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ComputerInventoryCollectionSettingsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ComputerInventoryCollectionSettingsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ComputerInventoryCollectionSettingsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ComputerInventoryCollectionSettingsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *ComputerInventoryCollectionSettingsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *ComputerInventoryCollectionSettingsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *ComputerInventoryCollectionSettingsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *ComputerInventoryCollectionSettingsMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *ComputerInventoryCollectionSettingsMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil && len(resp.Body) > 0 {
		_ = mergePage(resp.Body)
	}
	return resp, nil
}
func (m *ComputerInventoryCollectionSettingsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *ComputerInventoryCollectionSettingsMock) InvalidateToken() error                     { return nil }
func (m *ComputerInventoryCollectionSettingsMock) KeepAliveToken() error                      { return nil }
func (m *ComputerInventoryCollectionSettingsMock) GetLogger() *zap.Logger                     { return m.logger }
