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

// DockItemsMock is a test double implementing interfaces.HTTPClient.
type DockItemsMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewDockItemsMock returns an empty mock ready for response registration.
func NewDockItemsMock() *DockItemsMock {
	return &DockItemsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *DockItemsMock) RegisterMocks() {
	m.RegisterGetDockItemMock()
	m.RegisterCreateDockItemMock()
	m.RegisterUpdateDockItemMock()
	m.RegisterDeleteDockItemMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *DockItemsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *DockItemsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("DockItemsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *DockItemsMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("DockItemsMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *DockItemsMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/json"}},
			Body:       []byte(`{"code":"NOT-FOUND","message":"no mock registered"}`),
		}, fmt.Errorf("DockItemsMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("DockItemsMock: unmarshal into result: %w", err)
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

func (m *DockItemsMock) RegisterGetDockItemMock() {
	m.register("GET", "/api/v1/dock-items/1", 200, "validate_get_dock_item.json")
}

func (m *DockItemsMock) RegisterCreateDockItemMock() {
	m.register("POST", "/api/v1/dock-items", 201, "validate_create_dock_item.json")
}

func (m *DockItemsMock) RegisterUpdateDockItemMock() {
	m.register("PUT", "/api/v1/dock-items/1", 200, "validate_update_dock_item.json")
}

func (m *DockItemsMock) RegisterDeleteDockItemMock() {
	m.register("DELETE", "/api/v1/dock-items/1", 204, "")
}

func (m *DockItemsMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/api/v1/dock-items/999", 404, "error_not_found.json")
}

func (m *DockItemsMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/api/v1/dock-items", 409, "error_conflict.json")
}

func (m *DockItemsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}

func (m *DockItemsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *DockItemsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *DockItemsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *DockItemsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *DockItemsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *DockItemsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *DockItemsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *DockItemsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *DockItemsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *DockItemsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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

func (m *DockItemsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *DockItemsMock) InvalidateToken() error                    { return nil }
func (m *DockItemsMock) KeepAliveToken() error                      { return nil }
func (m *DockItemsMock) GetLogger() *zap.Logger                     { return m.logger }
