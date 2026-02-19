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

// ComputerGroupsMock is a test double implementing interfaces.HTTPClient.
type ComputerGroupsMock struct {
	responses     map[string]registeredResponse
	logger         *zap.Logger
	LastRSQLQuery map[string]string
}

// NewComputerGroupsMock returns an empty mock ready for response registration.
func NewComputerGroupsMock() *ComputerGroupsMock {
	return &ComputerGroupsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *ComputerGroupsMock) RegisterMocks() {
	m.RegisterListSmartGroupsMock()
	m.RegisterGetSmartGroupMock()
	m.RegisterCreateSmartGroupMock()
	m.RegisterUpdateSmartGroupMock()
	m.RegisterDeleteSmartGroupMock()
	m.RegisterListStaticGroupsMock()
	m.RegisterGetStaticGroupMock()
	m.RegisterCreateStaticGroupMock()
	m.RegisterUpdateStaticGroupMock()
	m.RegisterDeleteStaticGroupMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *ComputerGroupsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *ComputerGroupsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("ComputerGroupsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *ComputerGroupsMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("ComputerGroupsMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *ComputerGroupsMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/json"}},
			Body:       []byte(`{"code":"NOT-FOUND","message":"no mock registered"}`),
		}, fmt.Errorf("ComputerGroupsMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("ComputerGroupsMock: unmarshal into result: %w", err)
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

// Smart groups
func (m *ComputerGroupsMock) RegisterListSmartGroupsMock() {
	m.register("GET", "/api/v2/computer-groups/smart-groups", 200, "validate_list_smart_groups.json")
}

func (m *ComputerGroupsMock) RegisterGetSmartGroupMock() {
	m.register("GET", "/api/v2/computer-groups/smart-groups/1", 200, "validate_get_smart_group.json")
}

func (m *ComputerGroupsMock) RegisterCreateSmartGroupMock() {
	m.register("POST", "/api/v2/computer-groups/smart-groups", 201, "validate_create_smart_group.json")
}

func (m *ComputerGroupsMock) RegisterUpdateSmartGroupMock() {
	m.register("PUT", "/api/v2/computer-groups/smart-groups/1", 200, "validate_update_smart_group.json")
}

func (m *ComputerGroupsMock) RegisterDeleteSmartGroupMock() {
	m.register("DELETE", "/api/v2/computer-groups/smart-groups/1", 204, "")
}

// Static groups
func (m *ComputerGroupsMock) RegisterListStaticGroupsMock() {
	m.register("GET", "/api/v2/computer-groups/static-groups", 200, "validate_list_static_groups.json")
}

func (m *ComputerGroupsMock) RegisterGetStaticGroupMock() {
	m.register("GET", "/api/v2/computer-groups/static-groups/10", 200, "validate_get_static_group.json")
}

func (m *ComputerGroupsMock) RegisterCreateStaticGroupMock() {
	m.register("POST", "/api/v2/computer-groups/static-groups", 201, "validate_create_static_group.json")
}

func (m *ComputerGroupsMock) RegisterUpdateStaticGroupMock() {
	m.register("PATCH", "/api/v2/computer-groups/static-groups/10", 200, "validate_update_static_group.json")
}

func (m *ComputerGroupsMock) RegisterDeleteStaticGroupMock() {
	m.register("DELETE", "/api/v2/computer-groups/static-groups/10", 204, "")
}

// Errors
func (m *ComputerGroupsMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/api/v2/computer-groups/smart-groups/999", 404, "error_not_found.json")
}

func (m *ComputerGroupsMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/api/v2/computer-groups/smart-groups", 409, "error_conflict.json")
}

func (m *ComputerGroupsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *ComputerGroupsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputerGroupsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputerGroupsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputerGroupsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputerGroupsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *ComputerGroupsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *ComputerGroupsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *ComputerGroupsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *ComputerGroupsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *ComputerGroupsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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

func (m *ComputerGroupsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *ComputerGroupsMock) InvalidateToken() error                    { return nil }
func (m *ComputerGroupsMock) KeepAliveToken() error                     { return nil }
func (m *ComputerGroupsMock) GetLogger() *zap.Logger                    { return m.logger }
