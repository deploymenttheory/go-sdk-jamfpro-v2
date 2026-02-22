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

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// LocalAdminPasswordMock is a test double implementing interfaces.HTTPClient.
type LocalAdminPasswordMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewLocalAdminPasswordMock returns an empty mock ready for response registration.
func NewLocalAdminPasswordMock() *LocalAdminPasswordMock {
	return &LocalAdminPasswordMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *LocalAdminPasswordMock) RegisterMocks() {
	m.RegisterGetPendingRotationsMock()
	m.RegisterGetSettingsMock()
	m.RegisterUpdateSettingsMock()
	m.RegisterGetPasswordHistoryMock()
	m.RegisterGetCurrentPasswordMock()
	m.RegisterGetFullHistoryMock()
	m.RegisterGetCapableAccountsMock()
	m.RegisterSetPasswordMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *LocalAdminPasswordMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
}

func (m *LocalAdminPasswordMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("LocalAdminPasswordMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *LocalAdminPasswordMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("LocalAdminPasswordMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *LocalAdminPasswordMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/json"}},
			Body:       []byte(`{"code":"NOT-FOUND","message":"no mock registered"}`),
		}, fmt.Errorf("LocalAdminPasswordMock: no response registered for %s %s", method, path)
	}

	resp := &interfaces.Response{
		StatusCode: r.statusCode,
		Status:     fmt.Sprintf("%d", r.statusCode),
		Headers:    http.Header{"Content-Type": {"application/json"}},
		Body:       r.rawBody,
	}

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("LocalAdminPasswordMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

func loadMockResponse(filename string) ([]byte, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("get working directory: %w", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "mocks", filename))
	if err != nil {
		return nil, fmt.Errorf("read fixture %s: %w", filename, err)
	}
	return data, nil
}

func (m *LocalAdminPasswordMock) RegisterGetPendingRotationsMock() {
	m.register("GET", "/api/v2/local-admin-password/pending-rotations", 200, "validate_pending_rotations.json")
}

func (m *LocalAdminPasswordMock) RegisterGetSettingsMock() {
	m.register("GET", "/api/v2/local-admin-password/settings", 200, "validate_get_settings.json")
}

func (m *LocalAdminPasswordMock) RegisterUpdateSettingsMock() {
	m.register("PUT", "/api/v2/local-admin-password/settings", 200, "")
}

func (m *LocalAdminPasswordMock) RegisterGetPasswordHistoryMock() {
	m.register("GET", "/api/v2/local-admin-password/device-001/account/admin/audit", 200, "validate_password_history.json")
}

func (m *LocalAdminPasswordMock) RegisterGetCurrentPasswordMock() {
	m.register("GET", "/api/v2/local-admin-password/device-001/account/admin/password", 200, "validate_current_password.json")
}

func (m *LocalAdminPasswordMock) RegisterGetFullHistoryMock() {
	m.register("GET", "/api/v2/local-admin-password/device-001/history", 200, "validate_full_history.json")
}

func (m *LocalAdminPasswordMock) RegisterGetCapableAccountsMock() {
	m.register("GET", "/api/v2/local-admin-password/device-001/accounts", 200, "validate_capable_accounts.json")
}

func (m *LocalAdminPasswordMock) RegisterSetPasswordMock() {
	m.register("PUT", "/api/v2/local-admin-password/device-001/set-password", 200, "validate_set_password.json")
}

func (m *LocalAdminPasswordMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/api/v2/local-admin-password/device-999/accounts", 404, "error_not_found.json")
}

func (m *LocalAdminPasswordMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}

func (m *LocalAdminPasswordMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *LocalAdminPasswordMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *LocalAdminPasswordMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *LocalAdminPasswordMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *LocalAdminPasswordMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *LocalAdminPasswordMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *LocalAdminPasswordMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *LocalAdminPasswordMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *LocalAdminPasswordMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *LocalAdminPasswordMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil {
		if err := mergePage(resp.Body); err != nil {
			return resp, err
		}
	}
	return resp, nil
}

func (m *LocalAdminPasswordMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *LocalAdminPasswordMock) InvalidateToken() error                    { return nil }
func (m *LocalAdminPasswordMock) KeepAliveToken() error                     { return nil }
func (m *LocalAdminPasswordMock) GetLogger() *zap.Logger                    { return m.logger }
