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

// MobileDeviceExtensionAttributesMock is a test double implementing interfaces.HTTPClient.
type MobileDeviceExtensionAttributesMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewMobileDeviceExtensionAttributesMock returns an empty mock ready for response registration.
func NewMobileDeviceExtensionAttributesMock() *MobileDeviceExtensionAttributesMock {
	return &MobileDeviceExtensionAttributesMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *MobileDeviceExtensionAttributesMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetMock()
	m.RegisterCreateMock()
	m.RegisterUpdateMock()
	m.RegisterDeleteMock()
	m.RegisterDeleteMultipleMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *MobileDeviceExtensionAttributesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
}

func (m *MobileDeviceExtensionAttributesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("MobileDeviceExtensionAttributesMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *MobileDeviceExtensionAttributesMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("MobileDeviceExtensionAttributesMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *MobileDeviceExtensionAttributesMock) RegisterListMock() {
	m.register("GET", "/api/v1/mobile-device-extension-attributes", 200, "validate_list.json")
}

func (m *MobileDeviceExtensionAttributesMock) RegisterGetMock() {
	m.register("GET", "/api/v1/mobile-device-extension-attributes/1", 200, "validate_get.json")
}

func (m *MobileDeviceExtensionAttributesMock) RegisterCreateMock() {
	m.register("POST", "/api/v1/mobile-device-extension-attributes", 201, "validate_create.json")
}

func (m *MobileDeviceExtensionAttributesMock) RegisterUpdateMock() {
	m.register("PUT", "/api/v1/mobile-device-extension-attributes/1", 200, "validate_update.json")
}

func (m *MobileDeviceExtensionAttributesMock) RegisterDeleteMock() {
	m.register("DELETE", "/api/v1/mobile-device-extension-attributes/1", 204, "")
}

func (m *MobileDeviceExtensionAttributesMock) RegisterDeleteMultipleMock() {
	m.register("POST", "/api/v1/mobile-device-extension-attributes/delete-multiple", 204, "")
}

func (m *MobileDeviceExtensionAttributesMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/api/v1/mobile-device-extension-attributes/999", 404, "error_not_found.json")
}

func (m *MobileDeviceExtensionAttributesMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *MobileDeviceExtensionAttributesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceExtensionAttributesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceExtensionAttributesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceExtensionAttributesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceExtensionAttributesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *MobileDeviceExtensionAttributesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *MobileDeviceExtensionAttributesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MobileDeviceExtensionAttributesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MobileDeviceExtensionAttributesMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *MobileDeviceExtensionAttributesMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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

func (m *MobileDeviceExtensionAttributesMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *MobileDeviceExtensionAttributesMock) InvalidateToken() error                    { return nil }
func (m *MobileDeviceExtensionAttributesMock) KeepAliveToken() error                    { return nil }
func (m *MobileDeviceExtensionAttributesMock) GetLogger() *zap.Logger                    { return m.logger }

func (m *MobileDeviceExtensionAttributesMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/json"}},
			Body:       []byte(`{"code":"NOT-FOUND","message":"no mock registered"}`),
		}, fmt.Errorf("MobileDeviceExtensionAttributesMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("MobileDeviceExtensionAttributesMock: unmarshal into result: %w", err)
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
