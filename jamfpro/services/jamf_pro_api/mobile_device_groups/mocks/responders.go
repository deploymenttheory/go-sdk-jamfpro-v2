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

type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// MobileDeviceGroupsMock is a test double implementing interfaces.HTTPClient.
type MobileDeviceGroupsMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewMobileDeviceGroupsMock returns an empty mock ready for response registration.
func NewMobileDeviceGroupsMock() *MobileDeviceGroupsMock {
	return &MobileDeviceGroupsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *MobileDeviceGroupsMock) RegisterMocks() {
	m.RegisterListSmartMock()
	m.RegisterGetSmartMock()
	m.RegisterCreateSmartMock()
	m.RegisterUpdateSmartMock()
	m.RegisterDeleteSmartMock()
	m.RegisterListStaticMock()
	m.RegisterGetStaticMock()
	m.RegisterCreateStaticMock()
	m.RegisterUpdateStaticMock()
	m.RegisterDeleteStaticMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *MobileDeviceGroupsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
}

func (m *MobileDeviceGroupsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("MobileDeviceGroupsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *MobileDeviceGroupsMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("MobileDeviceGroupsMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *MobileDeviceGroupsMock) RegisterListSmartMock() {
	m.register("GET", "/api/v1/mobile-device-groups/smart-groups", 200, "validate_list_smart_groups.json")
}

func (m *MobileDeviceGroupsMock) RegisterGetSmartMock() {
	m.register("GET", "/api/v1/mobile-device-groups/smart-groups/1", 200, "validate_get_smart_group.json")
}

func (m *MobileDeviceGroupsMock) RegisterCreateSmartMock() {
	m.register("POST", "/api/v1/mobile-device-groups/smart-groups", 201, "validate_create_smart_group.json")
}

func (m *MobileDeviceGroupsMock) RegisterUpdateSmartMock() {
	m.register("PUT", "/api/v1/mobile-device-groups/smart-groups/1", 200, "validate_update_smart_group.json")
}

func (m *MobileDeviceGroupsMock) RegisterDeleteSmartMock() {
	m.register("DELETE", "/api/v1/mobile-device-groups/smart-groups/1", 204, "")
}

func (m *MobileDeviceGroupsMock) RegisterListStaticMock() {
	m.register("GET", "/api/v1/mobile-device-groups/static-groups", 200, "validate_list_static_groups.json")
}

func (m *MobileDeviceGroupsMock) RegisterGetStaticMock() {
	m.register("GET", "/api/v1/mobile-device-groups/static-groups/10", 200, "validate_get_static_group.json")
}

func (m *MobileDeviceGroupsMock) RegisterCreateStaticMock() {
	m.register("POST", "/api/v1/mobile-device-groups/static-groups", 201, "validate_create_static_group.json")
}

func (m *MobileDeviceGroupsMock) RegisterUpdateStaticMock() {
	m.register("PATCH", "/api/v1/mobile-device-groups/static-groups/10", 200, "validate_update_static_group.json")
}

func (m *MobileDeviceGroupsMock) RegisterDeleteStaticMock() {
	m.register("DELETE", "/api/v1/mobile-device-groups/static-groups/10", 204, "")
}

func (m *MobileDeviceGroupsMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/api/v1/mobile-device-groups/smart-groups/999", 404, "error_not_found.json")
}

func (m *MobileDeviceGroupsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *MobileDeviceGroupsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceGroupsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceGroupsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceGroupsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceGroupsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *MobileDeviceGroupsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *MobileDeviceGroupsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MobileDeviceGroupsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MobileDeviceGroupsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *MobileDeviceGroupsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil && len(resp.Body) > 0 {
		if err := mergePage(resp.Body); err != nil {
			return resp, err
		}
	}
	return resp, nil
}

func (m *MobileDeviceGroupsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *MobileDeviceGroupsMock) InvalidateToken() error                      { return nil }
func (m *MobileDeviceGroupsMock) KeepAliveToken() error                     { return nil }
func (m *MobileDeviceGroupsMock) GetLogger() *zap.Logger                    { return m.logger }

func (m *MobileDeviceGroupsMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/json"}},
			Body:       []byte(`{"code":"NOT-FOUND","message":"no mock registered"}`),
		}, fmt.Errorf("MobileDeviceGroupsMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("MobileDeviceGroupsMock: unmarshal into result: %w", err)
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
