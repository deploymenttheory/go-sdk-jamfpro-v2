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

// DepartmentsMock is a test double implementing interfaces.HTTPClient.
type DepartmentsMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewDepartmentsMock returns an empty mock ready for response registration.
func NewDepartmentsMock() *DepartmentsMock {
	return &DepartmentsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *DepartmentsMock) RegisterMocks() {
	m.RegisterListDepartmentsMock()
	m.RegisterGetDepartmentMock()
	m.RegisterCreateDepartmentMock()
	m.RegisterUpdateDepartmentMock()
	m.RegisterDeleteDepartmentMock()
	m.RegisterGetDepartmentHistoryMock()
	m.RegisterAddDepartmentHistoryNotesMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *DepartmentsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *DepartmentsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("DepartmentsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *DepartmentsMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("DepartmentsMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *DepartmentsMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/json"}},
			Body:       []byte(`{"code":"NOT-FOUND","message":"no mock registered"}`),
		}, fmt.Errorf("DepartmentsMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("DepartmentsMock: unmarshal into result: %w", err)
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

func (m *DepartmentsMock) RegisterListDepartmentsMock() {
	m.register("GET", "/api/v1/departments", 200, "validate_list_departments.json")
}

func (m *DepartmentsMock) RegisterListDepartmentsRSQLMock() {
	m.register("GET", "/api/v1/departments", 200, "validate_list_departments_rsql.json")
}

func (m *DepartmentsMock) RegisterGetDepartmentMock() {
	m.register("GET", "/api/v1/departments/1", 200, "validate_get_department.json")
}

func (m *DepartmentsMock) RegisterCreateDepartmentMock() {
	m.register("POST", "/api/v1/departments", 201, "validate_create_department.json")
}

func (m *DepartmentsMock) RegisterUpdateDepartmentMock() {
	m.register("PUT", "/api/v1/departments/1", 200, "validate_update_department.json")
}

func (m *DepartmentsMock) RegisterDeleteDepartmentMock() {
	m.register("DELETE", "/api/v1/departments/1", 204, "")
}

func (m *DepartmentsMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/api/v1/departments/999", 404, "error_not_found.json")
}

func (m *DepartmentsMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/api/v1/departments", 409, "error_conflict.json")
}

func (m *DepartmentsMock) RegisterGetDepartmentHistoryMock() {
	m.register("GET", "/api/v1/departments/1/history", 200, "validate_get_history.json")
}

func (m *DepartmentsMock) RegisterAddDepartmentHistoryNotesMock() {
	m.register("POST", "/api/v1/departments/1/history", 201, "")
}

func (m *DepartmentsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *DepartmentsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *DepartmentsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *DepartmentsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *DepartmentsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *DepartmentsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *DepartmentsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *DepartmentsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *DepartmentsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *DepartmentsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *DepartmentsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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

func (m *DepartmentsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *DepartmentsMock) InvalidateToken() error                    { return nil }
func (m *DepartmentsMock) KeepAliveToken() error                     { return nil }
func (m *DepartmentsMock) GetLogger() *zap.Logger                    { return m.logger }
