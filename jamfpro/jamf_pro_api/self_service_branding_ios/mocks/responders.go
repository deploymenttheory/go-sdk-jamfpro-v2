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

// SelfServiceBrandingMobileMock is a test double implementing interfaces.HTTPClient.
type SelfServiceBrandingMobileMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewSelfServiceBrandingMobileMock returns an empty mock ready for response registration.
func NewSelfServiceBrandingMobileMock() *SelfServiceBrandingMobileMock {
	return &SelfServiceBrandingMobileMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *SelfServiceBrandingMobileMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetByIDMock()
	m.RegisterCreateMock()
	m.RegisterUpdateMock()
	m.RegisterDeleteMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *SelfServiceBrandingMobileMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *SelfServiceBrandingMobileMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("SelfServiceBrandingMobileMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *SelfServiceBrandingMobileMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("SelfServiceBrandingMobileMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *SelfServiceBrandingMobileMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("SelfServiceBrandingMobileMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {"application/json"}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("SelfServiceBrandingMobileMock: unmarshal into result: %w", err)
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

func (m *SelfServiceBrandingMobileMock) RegisterListMock() {
	m.register("GET", "/api/v1/self-service/branding/ios", 200, "validate_list.json")
}

func (m *SelfServiceBrandingMobileMock) RegisterGetByIDMock() {
	m.register("GET", "/api/v1/self-service/branding/ios/1", 200, "validate_get.json")
}

func (m *SelfServiceBrandingMobileMock) RegisterCreateMock() {
	m.register("POST", "/api/v1/self-service/branding/ios", 201, "validate_create.json")
}

func (m *SelfServiceBrandingMobileMock) RegisterUpdateMock() {
	m.register("PUT", "/api/v1/self-service/branding/ios/1", 200, "validate_update.json")
}

func (m *SelfServiceBrandingMobileMock) RegisterDeleteMock() {
	m.register("DELETE", "/api/v1/self-service/branding/ios/1", 204, "")
}

func (m *SelfServiceBrandingMobileMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/api/v1/self-service/branding/ios/999", 404, "error_not_found.json")
}

func (m *SelfServiceBrandingMobileMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/api/v1/self-service/branding/ios", 409, "error_conflict.json")
}

func (m *SelfServiceBrandingMobileMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}

func (m *SelfServiceBrandingMobileMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *SelfServiceBrandingMobileMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *SelfServiceBrandingMobileMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *SelfServiceBrandingMobileMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *SelfServiceBrandingMobileMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *SelfServiceBrandingMobileMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *SelfServiceBrandingMobileMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *SelfServiceBrandingMobileMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *SelfServiceBrandingMobileMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *SelfServiceBrandingMobileMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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

func (m *SelfServiceBrandingMobileMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *SelfServiceBrandingMobileMock) InvalidateToken() error                    { return nil }
func (m *SelfServiceBrandingMobileMock) KeepAliveToken() error                     { return nil }
func (m *SelfServiceBrandingMobileMock) GetLogger() *zap.Logger                    { return m.logger }
