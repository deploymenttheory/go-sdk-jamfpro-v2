package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

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

// StaticMobileDeviceGroupsMock is a test double implementing transport.HTTPClient.
type StaticMobileDeviceGroupsMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewStaticMobileDeviceGroupsMock returns an empty mock ready for response registration.
func NewStaticMobileDeviceGroupsMock() *StaticMobileDeviceGroupsMock {
	return &StaticMobileDeviceGroupsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *StaticMobileDeviceGroupsMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetMock()
	m.RegisterCreateMock()
	m.RegisterUpdateMock()
	m.RegisterDeleteMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *StaticMobileDeviceGroupsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
}

func (m *StaticMobileDeviceGroupsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("StaticMobileDeviceGroupsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *StaticMobileDeviceGroupsMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("StaticMobileDeviceGroupsMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *StaticMobileDeviceGroupsMock) RegisterListMock() {
	m.register("GET", "/api/v2/mobile-device-groups/static-groups", 200, "validate_list.json")
}

func (m *StaticMobileDeviceGroupsMock) RegisterGetMock() {
	m.register("GET", "/api/v2/mobile-device-groups/static-groups/10", 200, "validate_get.json")
}

func (m *StaticMobileDeviceGroupsMock) RegisterCreateMock() {
	m.register("POST", "/api/v2/mobile-device-groups/static-groups", 201, "validate_create.json")
}

func (m *StaticMobileDeviceGroupsMock) RegisterUpdateMock() {
	m.register("PATCH", "/api/v2/mobile-device-groups/static-groups/10", 200, "validate_update.json")
}

func (m *StaticMobileDeviceGroupsMock) RegisterDeleteMock() {
	m.register("DELETE", "/api/v2/mobile-device-groups/static-groups/10", 204, "")
}

func (m *StaticMobileDeviceGroupsMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/api/v2/mobile-device-groups/static-groups/999", 404, "error_not_found.json")
}

func (m *StaticMobileDeviceGroupsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *StaticMobileDeviceGroupsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *StaticMobileDeviceGroupsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *StaticMobileDeviceGroupsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *StaticMobileDeviceGroupsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ transport.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *StaticMobileDeviceGroupsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *StaticMobileDeviceGroupsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *StaticMobileDeviceGroupsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *StaticMobileDeviceGroupsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *StaticMobileDeviceGroupsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *StaticMobileDeviceGroupsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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

func (m *StaticMobileDeviceGroupsMock) RSQLBuilder() transport.RSQLFilterBuilder { return nil }
func (m *StaticMobileDeviceGroupsMock) InvalidateToken() error                    { return nil }
func (m *StaticMobileDeviceGroupsMock) KeepAliveToken() error                     { return nil }
func (m *StaticMobileDeviceGroupsMock) GetLogger() *zap.Logger                    { return m.logger }

func (m *StaticMobileDeviceGroupsMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("StaticMobileDeviceGroupsMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {"application/json"}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("StaticMobileDeviceGroupsMock: unmarshal into result: %w", err)
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
