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

// ComputerExtensionAttributesMock is a test double implementing interfaces.HTTPClient.
type ComputerExtensionAttributesMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewComputerExtensionAttributesMock returns an empty mock ready for response registration.
func NewComputerExtensionAttributesMock() *ComputerExtensionAttributesMock {
	return &ComputerExtensionAttributesMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *ComputerExtensionAttributesMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetMock()
	m.RegisterCreateMock()
	m.RegisterUpdateMock()
	m.RegisterDeleteMock()
	m.RegisterDeleteMultipleMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *ComputerExtensionAttributesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
}

func (m *ComputerExtensionAttributesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("ComputerExtensionAttributesMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *ComputerExtensionAttributesMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("ComputerExtensionAttributesMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *ComputerExtensionAttributesMock) RegisterListMock() {
	m.register("GET", "/api/v1/computer-extension-attributes", 200, "validate_list.json")
}

func (m *ComputerExtensionAttributesMock) RegisterGetMock() {
	m.register("GET", "/api/v1/computer-extension-attributes/1", 200, "validate_get.json")
}

func (m *ComputerExtensionAttributesMock) RegisterCreateMock() {
	m.register("POST", "/api/v1/computer-extension-attributes", 201, "validate_create.json")
}

func (m *ComputerExtensionAttributesMock) RegisterUpdateMock() {
	m.register("PUT", "/api/v1/computer-extension-attributes/1", 200, "validate_update.json")
}

func (m *ComputerExtensionAttributesMock) RegisterDeleteMock() {
	m.register("DELETE", "/api/v1/computer-extension-attributes/1", 204, "")
}

func (m *ComputerExtensionAttributesMock) RegisterDeleteMultipleMock() {
	m.register("POST", "/api/v1/computer-extension-attributes/delete-multiple", 204, "")
}

func (m *ComputerExtensionAttributesMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/api/v1/computer-extension-attributes/999", 404, "error_not_found.json")
}

func (m *ComputerExtensionAttributesMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *ComputerExtensionAttributesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputerExtensionAttributesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputerExtensionAttributesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputerExtensionAttributesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputerExtensionAttributesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *ComputerExtensionAttributesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *ComputerExtensionAttributesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *ComputerExtensionAttributesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *ComputerExtensionAttributesMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *ComputerExtensionAttributesMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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

func (m *ComputerExtensionAttributesMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *ComputerExtensionAttributesMock) InvalidateToken() error                    { return nil }
func (m *ComputerExtensionAttributesMock) KeepAliveToken() error                      { return nil }
func (m *ComputerExtensionAttributesMock) GetLogger() *zap.Logger                      { return m.logger }

func (m *ComputerExtensionAttributesMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/json"}},
			Body:       []byte(`{"code":"NOT-FOUND","message":"no mock registered"}`),
		}, fmt.Errorf("ComputerExtensionAttributesMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("ComputerExtensionAttributesMock: unmarshal into result: %w", err)
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
