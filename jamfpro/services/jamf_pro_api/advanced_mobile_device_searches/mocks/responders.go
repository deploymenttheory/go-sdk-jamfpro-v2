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

type AdvancedMobileDeviceSearchesMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewAdvancedMobileDeviceSearchesMock() *AdvancedMobileDeviceSearchesMock {
	return &AdvancedMobileDeviceSearchesMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *AdvancedMobileDeviceSearchesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("AdvancedMobileDeviceSearchesMock: load %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *AdvancedMobileDeviceSearchesMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("AdvancedMobileDeviceSearchesMock: load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)}
}

func (m *AdvancedMobileDeviceSearchesMock) RegisterMocks() {
	m.register("GET", "/api/v1/advanced-mobile-device-searches", 200, "validate_list.json")
	m.register("GET", "/api/v1/advanced-mobile-device-searches/1", 200, "validate_get.json")
	m.register("POST", "/api/v1/advanced-mobile-device-searches", 201, "validate_create.json")
	m.register("PUT", "/api/v1/advanced-mobile-device-searches/1", 200, "validate_get.json")
	m.register("DELETE", "/api/v1/advanced-mobile-device-searches/1", 204, "")
	m.register("GET", "/api/v1/advanced-mobile-device-searches/choices?criteria=Device%20Name&site=-1&contains=", 200, "validate_choices.json")
	m.register("GET", "/api/v1/advanced-mobile-device-searches/choices?criteria=Device+Name&site=-1&contains=", 200, "validate_choices.json")
}

func (m *AdvancedMobileDeviceSearchesMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/api/v1/advanced-mobile-device-searches/999", 404, "error_not_found.json")
}

func (m *AdvancedMobileDeviceSearchesMock) Get(ctx context.Context, path string, q map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *AdvancedMobileDeviceSearchesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *AdvancedMobileDeviceSearchesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *AdvancedMobileDeviceSearchesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *AdvancedMobileDeviceSearchesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *AdvancedMobileDeviceSearchesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *AdvancedMobileDeviceSearchesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *AdvancedMobileDeviceSearchesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *AdvancedMobileDeviceSearchesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *AdvancedMobileDeviceSearchesMock) GetBytes(ctx context.Context, path string, q map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *AdvancedMobileDeviceSearchesMock) GetPaginated(ctx context.Context, path string, q map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil && len(resp.Body) > 0 {
		_ = mergePage(resp.Body)
	}
	return resp, nil
}
func (m *AdvancedMobileDeviceSearchesMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *AdvancedMobileDeviceSearchesMock) InvalidateToken() error                    { return nil }
func (m *AdvancedMobileDeviceSearchesMock) KeepAliveToken() error                     { return nil }
func (m *AdvancedMobileDeviceSearchesMock) GetLogger() *zap.Logger                     { return m.logger }

func (m *AdvancedMobileDeviceSearchesMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{StatusCode: 404, Headers: http.Header{}, Body: nil}, fmt.Errorf("AdvancedMobileDeviceSearchesMock: no response for %s %s", method, path)
	}
	resp := &interfaces.Response{StatusCode: r.statusCode, Status: fmt.Sprintf("%d", r.statusCode), Headers: http.Header{"Content-Type": {"application/json"}}, Body: r.rawBody}
	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}
	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, err
		}
	}
	return resp, nil
}

func loadMockResponse(filename string) ([]byte, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return os.ReadFile(filepath.Join(dir, "mocks", filename))
}
