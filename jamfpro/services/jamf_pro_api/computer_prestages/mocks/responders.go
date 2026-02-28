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

type ComputerPrestagesMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewComputerPrestagesMock() *ComputerPrestagesMock {
	return &ComputerPrestagesMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *ComputerPrestagesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		body, _ = os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *ComputerPrestagesMock) RegisterListMock() {
	m.register("GET", "/api/v3/computer-prestages", 200, "validate_list.json")
}

func (m *ComputerPrestagesMock) RegisterGetByIDMock(id string) {
	m.register("GET", "/api/v3/computer-prestages/"+id, 200, "validate_get.json")
}

func (m *ComputerPrestagesMock) RegisterCreateMock() {
	m.register("POST", "/api/v3/computer-prestages", 200, "validate_create.json")
}

func (m *ComputerPrestagesMock) RegisterUpdateByIDMock(id string) {
	m.register("PUT", "/api/v3/computer-prestages/"+id, 200, "validate_get.json")
}

func (m *ComputerPrestagesMock) RegisterDeleteByIDMock(id string) {
	m.register("DELETE", "/api/v3/computer-prestages/"+id, 200, "")
}

func (m *ComputerPrestagesMock) RegisterGetDeviceScopeMock(id string) {
	m.register("GET", "/api/v2/computer-prestages/"+id+"/scope", 200, "validate_scope.json")
}

func (m *ComputerPrestagesMock) RegisterReplaceDeviceScopeMock(id string) {
	m.register("PUT", "/api/v2/computer-prestages/"+id+"/scope", 200, "validate_scope.json")
}

func (m *ComputerPrestagesMock) RegisterGetAllDeviceScopeMock() {
	m.register("GET", "/api/v2/computer-prestages/scope", 200, "validate_all_scope.json")
}

func (m *ComputerPrestagesMock) RegisterAddDeviceScopeMock(id string) {
	m.register("POST", "/api/v2/computer-prestages/"+id+"/scope", 200, "validate_scope.json")
}

func (m *ComputerPrestagesMock) RegisterRemoveDeviceScopeMock(id string) {
	m.register("POST", "/api/v2/computer-prestages/"+id+"/scope/delete-multiple", 200, "validate_scope.json")
}

func (m *ComputerPrestagesMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("ComputerPrestagesMock: no response for %s %s", method, path)
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

func (m *ComputerPrestagesMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *ComputerPrestagesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ComputerPrestagesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ComputerPrestagesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ComputerPrestagesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ComputerPrestagesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *ComputerPrestagesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *ComputerPrestagesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *ComputerPrestagesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *ComputerPrestagesMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *ComputerPrestagesMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil && len(resp.Body) > 0 {
		var wrapper struct {
			Results json.RawMessage `json:"results"`
		}
		if err := json.Unmarshal(resp.Body, &wrapper); err != nil {
			return resp, fmt.Errorf("failed to extract results: %w", err)
		}
		if err := mergePage(wrapper.Results); err != nil {
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
	}
	return resp, nil
}
func (m *ComputerPrestagesMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *ComputerPrestagesMock) InvalidateToken() error                     { return nil }
func (m *ComputerPrestagesMock) KeepAliveToken() error                      { return nil }
func (m *ComputerPrestagesMock) GetLogger() *zap.Logger                     { return m.logger }
