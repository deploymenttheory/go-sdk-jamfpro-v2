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

type PatchPoliciesMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewPatchPoliciesMock() *PatchPoliciesMock {
	return &PatchPoliciesMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *PatchPoliciesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		body, _ = os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *PatchPoliciesMock) RegisterListMock() {
	m.register("GET", "/api/v2/patch-policies/policy-details", 200, "validate_list.json")
}

func (m *PatchPoliciesMock) RegisterGetDashboardStatusMock(id string, onDashboard bool) {
	fixture := "validate_dashboard_status_false.json"
	if onDashboard {
		fixture = "validate_dashboard_status.json"
	}
	m.register("GET", "/api/v2/patch-policies/"+id+"/dashboard", 200, fixture)
}

func (m *PatchPoliciesMock) RegisterAddToDashboardMock(id string) {
	m.register("POST", "/api/v2/patch-policies/"+id+"/dashboard", 200, "")
}

func (m *PatchPoliciesMock) RegisterRemoveFromDashboardMock(id string) {
	m.register("DELETE", "/api/v2/patch-policies/"+id+"/dashboard", 200, "")
}

func (m *PatchPoliciesMock) RegisterEmptyListMock() {
	m.register("GET", "/api/v2/patch-policies/policy-details", 200, "validate_empty_list.json")
}

func (m *PatchPoliciesMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{StatusCode: 404, Headers: http.Header{}, Body: nil}, fmt.Errorf("PatchPoliciesMock: no response for %s %s", method, path)
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

func (m *PatchPoliciesMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *PatchPoliciesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *PatchPoliciesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *PatchPoliciesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *PatchPoliciesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *PatchPoliciesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *PatchPoliciesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *PatchPoliciesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *PatchPoliciesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *PatchPoliciesMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *PatchPoliciesMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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
func (m *PatchPoliciesMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *PatchPoliciesMock) InvalidateToken() error                     { return nil }
func (m *PatchPoliciesMock) KeepAliveToken() error                      { return nil }
func (m *PatchPoliciesMock) GetLogger() *zap.Logger                     { return m.logger }
