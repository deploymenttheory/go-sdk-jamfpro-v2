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

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// ImpactAlertNotificationSettingsMock is a test double implementing transport.HTTPClient.
type ImpactAlertNotificationSettingsMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewImpactAlertNotificationSettingsMock returns an empty mock ready for response registration.
func NewImpactAlertNotificationSettingsMock() *ImpactAlertNotificationSettingsMock {
	return &ImpactAlertNotificationSettingsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *ImpactAlertNotificationSettingsMock) RegisterMocks() {
	m.RegisterGetMock()
	m.RegisterUpdateMock()
}

func (m *ImpactAlertNotificationSettingsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("ImpactAlertNotificationSettingsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *ImpactAlertNotificationSettingsMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("ImpactAlertNotificationSettingsMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func loadMockResponse(filename string) ([]byte, error) {
	_, callerPath, _, _ := runtime.Caller(0)
	dir := filepath.Dir(callerPath)
	data, err := os.ReadFile(filepath.Join(dir, filename))
	if err != nil {
		return nil, fmt.Errorf("read fixture %s: %w", filename, err)
	}
	return data, nil
}

func (m *ImpactAlertNotificationSettingsMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("ImpactAlertNotificationSettingsMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {"application/json"}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("ImpactAlertNotificationSettingsMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

func (m *ImpactAlertNotificationSettingsMock) RegisterGetMock() {
	m.register("GET", "/api/v1/impact-alert-notification-settings", 200, "validate_get.json")
}

func (m *ImpactAlertNotificationSettingsMock) RegisterUpdateMock() {
	// Update returns 204 No Content
	m.register("PUT", "/api/v1/impact-alert-notification-settings", 204, "")
}

func (m *ImpactAlertNotificationSettingsMock) RegisterGetErrorMock() {
	m.registerError("GET", "/api/v1/impact-alert-notification-settings", 404, "error_not_found.json")
}

func (m *ImpactAlertNotificationSettingsMock) RegisterUpdateErrorMock() {
	m.registerError("PUT", "/api/v1/impact-alert-notification-settings", 400, "error_update.json")
}

func (m *ImpactAlertNotificationSettingsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}

func (m *ImpactAlertNotificationSettingsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ImpactAlertNotificationSettingsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ImpactAlertNotificationSettingsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ImpactAlertNotificationSettingsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ transport.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ImpactAlertNotificationSettingsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *ImpactAlertNotificationSettingsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *ImpactAlertNotificationSettingsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *ImpactAlertNotificationSettingsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *ImpactAlertNotificationSettingsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *ImpactAlertNotificationSettingsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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

func (m *ImpactAlertNotificationSettingsMock) RSQLBuilder() transport.RSQLFilterBuilder { return nil }
func (m *ImpactAlertNotificationSettingsMock) InvalidateToken() error                    { return nil }
func (m *ImpactAlertNotificationSettingsMock) KeepAliveToken() error                     { return nil }
func (m *ImpactAlertNotificationSettingsMock) GetLogger() *zap.Logger                    { return m.logger }
