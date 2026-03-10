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

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"go.uber.org/zap"
	"resty.dev/v3"

	mockhelpers "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
}

type DeclarativeDeviceManagementMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewDeclarativeDeviceManagementMock() *DeclarativeDeviceManagementMock {
	return &DeclarativeDeviceManagementMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *DeclarativeDeviceManagementMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		body, _ = os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *DeclarativeDeviceManagementMock) RegisterForceSyncMock(clientManagementID string) {
	m.register("POST", "/api/v1/ddm/"+clientManagementID+"/sync", 204, "")
}

func (m *DeclarativeDeviceManagementMock) RegisterGetStatusItemsMock(clientManagementID string) {
	m.register("GET", "/api/v1/ddm/"+clientManagementID+"/status-items", 200, "validate_status_items.json")
}

func (m *DeclarativeDeviceManagementMock) RegisterGetStatusItemByKeyMock(clientManagementID, key string) {
	m.register("GET", "/api/v1/ddm/"+clientManagementID+"/status-items/"+key, 200, "validate_status_item.json")
}

func (m *DeclarativeDeviceManagementMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("DeclarativeDeviceManagementMock: no response for %s %s", method, path)
	}
	headers := http.Header{"Content-Type": {"application/json"}}
	resp := mockhelpers.NewMockResponse(r.statusCode, headers, r.rawBody)
	if result != nil && len(r.rawBody) > 0 {
		_ = json.Unmarshal(r.rawBody, result)
	}
	return resp, nil
}

func mustMocksDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func (m *DeclarativeDeviceManagementMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *DeclarativeDeviceManagementMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *DeclarativeDeviceManagementMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *DeclarativeDeviceManagementMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *DeclarativeDeviceManagementMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *DeclarativeDeviceManagementMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *DeclarativeDeviceManagementMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *DeclarativeDeviceManagementMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *DeclarativeDeviceManagementMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *DeclarativeDeviceManagementMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}
func (m *DeclarativeDeviceManagementMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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
func (m *DeclarativeDeviceManagementMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *DeclarativeDeviceManagementMock) InvalidateToken() error                    { return nil }
func (m *DeclarativeDeviceManagementMock) KeepAliveToken() error                     { return nil }
func (m *DeclarativeDeviceManagementMock) GetLogger() *zap.Logger                    { return m.logger }
