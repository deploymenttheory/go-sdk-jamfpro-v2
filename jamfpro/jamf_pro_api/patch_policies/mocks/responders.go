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

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"go.uber.org/zap"
	"resty.dev/v3"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
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

func (m *PatchPoliciesMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	if err != nil {
		panic("PatchPoliciesMock: failed to load error fixture: " + err.Error())
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
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

func (m *PatchPoliciesMock) RegisterListSummaryMock() {
	m.register("GET", "/api/v2/patch-policies", 200, "validate_list_summary.json")
}

func (m *PatchPoliciesMock) RegisterListSummaryEmptyMock() {
	m.register("GET", "/api/v2/patch-policies", 200, "validate_list_summary_empty.json")
}

func (m *PatchPoliciesMock) RegisterListErrorMock() {
	m.registerError("GET", "/api/v2/patch-policies/policy-details", 500, "error_not_found.json")
}

func (m *PatchPoliciesMock) RegisterListInvalidMock() {
	m.register("GET", "/api/v2/patch-policies/policy-details", 200, "validate_list_invalid.json")
}

func (m *PatchPoliciesMock) RegisterListSummaryErrorMock() {
	m.registerError("GET", "/api/v2/patch-policies", 500, "error_not_found.json")
}

func (m *PatchPoliciesMock) RegisterGetDashboardStatusErrorMock(id string) {
	m.registerError("GET", "/api/v2/patch-policies/"+id+"/dashboard", 404, "error_not_found.json")
}

func (m *PatchPoliciesMock) RegisterAddToDashboardErrorMock(id string) {
	m.registerError("POST", "/api/v2/patch-policies/"+id+"/dashboard", 500, "error_not_found.json")
}

func (m *PatchPoliciesMock) RegisterRemoveFromDashboardErrorMock(id string) {
	m.registerError("DELETE", "/api/v2/patch-policies/"+id+"/dashboard", 500, "error_not_found.json")
}

func (m *PatchPoliciesMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("PatchPoliciesMock: no response registered for %s %s", method, path)
	}
	headers := http.Header{"Content-Type": {"application/json"}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)
	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}
	if result != nil && len(r.rawBody) > 0 {
		_ = json.Unmarshal(r.rawBody, result)
	}
	return resp, nil
}

func mustMocksDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func (m *PatchPoliciesMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *PatchPoliciesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *PatchPoliciesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *PatchPoliciesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *PatchPoliciesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ transport.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *PatchPoliciesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *PatchPoliciesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *PatchPoliciesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *PatchPoliciesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *PatchPoliciesMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}
func (m *PatchPoliciesMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	bodyBytes := resp.Bytes()
	if mergePage != nil && len(bodyBytes) > 0 {
		var page struct {
			Results json.RawMessage `json:"results"`
		}
		if err := json.Unmarshal(bodyBytes, &page); err != nil {
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
		if err := mergePage(page.Results); err != nil {
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
	}
	return resp, nil
}
func (m *PatchPoliciesMock) RSQLBuilder() transport.RSQLFilterBuilder { return nil }
func (m *PatchPoliciesMock) InvalidateToken() error                     { return nil }
func (m *PatchPoliciesMock) KeepAliveToken() error                      { return nil }
func (m *PatchPoliciesMock) GetLogger() *zap.Logger                     { return m.logger }
