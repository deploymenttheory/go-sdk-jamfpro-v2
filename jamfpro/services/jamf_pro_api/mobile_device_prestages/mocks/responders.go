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

type MobileDevicePrestagesMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewMobileDevicePrestagesMock() *MobileDevicePrestagesMock {
	return &MobileDevicePrestagesMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *MobileDevicePrestagesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		body, _ = os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *MobileDevicePrestagesMock) RegisterListMock() {
	m.register("GET", "/api/v3/mobile-device-prestages", 200, "validate_list.json")
}

func (m *MobileDevicePrestagesMock) RegisterGetByIDMock(id string) {
	m.register("GET", "/api/v3/mobile-device-prestages/"+id, 200, "validate_get.json")
}

func (m *MobileDevicePrestagesMock) RegisterCreateMock() {
	m.register("POST", "/api/v3/mobile-device-prestages", 200, "validate_create.json")
}

func (m *MobileDevicePrestagesMock) RegisterUpdateByIDMock(id string) {
	m.register("PUT", "/api/v3/mobile-device-prestages/"+id, 200, "validate_get.json")
}

func (m *MobileDevicePrestagesMock) RegisterDeleteByIDMock(id string) {
	m.register("DELETE", "/api/v3/mobile-device-prestages/"+id, 200, "")
}

func (m *MobileDevicePrestagesMock) RegisterGetScopeByIDMock(id string) {
	m.register("GET", "/api/v2/mobile-device-prestages/"+id+"/scope", 200, "validate_scope.json")
}

func (m *MobileDevicePrestagesMock) RegisterEmptyListMock() {
	m.register("GET", "/api/v3/mobile-device-prestages", 200, "validate_empty_list.json")
}

func (m *MobileDevicePrestagesMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{StatusCode: 404, Headers: http.Header{}, Body: nil}, fmt.Errorf("MobileDevicePrestagesMock: no response for %s %s", method, path)
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

func (m *MobileDevicePrestagesMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *MobileDevicePrestagesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *MobileDevicePrestagesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *MobileDevicePrestagesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *MobileDevicePrestagesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *MobileDevicePrestagesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *MobileDevicePrestagesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *MobileDevicePrestagesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *MobileDevicePrestagesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *MobileDevicePrestagesMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *MobileDevicePrestagesMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil && len(resp.Body) > 0 {
		// Extract the results array from the response for pagination
		var wrapper struct {
			Results json.RawMessage `json:"results"`
		}
		if err := json.Unmarshal(resp.Body, &wrapper); err == nil {
			_ = mergePage(wrapper.Results)
		}
	}
	return resp, nil
}
func (m *MobileDevicePrestagesMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *MobileDevicePrestagesMock) InvalidateToken() error                     { return nil }
func (m *MobileDevicePrestagesMock) KeepAliveToken() error                      { return nil }
func (m *MobileDevicePrestagesMock) GetLogger() *zap.Logger                     { return m.logger }
