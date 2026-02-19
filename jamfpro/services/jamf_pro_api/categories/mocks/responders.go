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
	// errMsg causes the mock to return an error (simulating API or transport failures).
	errMsg string
}

// CategoriesMock is a test double implementing interfaces.HTTPClient.
// Responses are keyed by "METHOD:path" and loaded from JSON fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
type CategoriesMock struct {
	responses    map[string]registeredResponse
	logger       *zap.Logger
	LastRSQLQuery map[string]string // captures the rsqlQuery from the most recent Get call
}

// NewCategoriesMock returns an empty mock ready for response registration.
func NewCategoriesMock() *CategoriesMock {
	return &CategoriesMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *CategoriesMock) RegisterMocks() {
	m.RegisterListCategoriesMock()
	m.RegisterGetCategoryMock()
	m.RegisterCreateCategoryMock()
	m.RegisterUpdateCategoryMock()
	m.RegisterDeleteCategoryMock()
	m.RegisterDeleteCategoriesBulkMock()
	m.RegisterGetCategoryHistoryMock()
	m.RegisterAddCategoryHistoryNotesMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *CategoriesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListCategoriesMock registers GET /api/v1/categories → 200.
func (m *CategoriesMock) RegisterListCategoriesMock() {
	m.register("GET", "/api/v1/categories", 200, "validate_list_categories.json")
}

// RegisterListCategoriesRSQLMock registers GET /api/v1/categories → 200 with a
// single-result fixture that simulates a server-side RSQL-filtered response.
func (m *CategoriesMock) RegisterListCategoriesRSQLMock() {
	m.register("GET", "/api/v1/categories", 200, "validate_list_categories_rsql.json")
}

// RegisterGetCategoryMock registers GET /api/v1/categories/1 → 200.
func (m *CategoriesMock) RegisterGetCategoryMock() {
	m.register("GET", "/api/v1/categories/1", 200, "validate_get_category.json")
}

// RegisterCreateCategoryMock registers POST /api/v1/categories → 201.
func (m *CategoriesMock) RegisterCreateCategoryMock() {
	m.register("POST", "/api/v1/categories", 201, "validate_create_category.json")
}

// RegisterUpdateCategoryMock registers PUT /api/v1/categories/1 → 200.
func (m *CategoriesMock) RegisterUpdateCategoryMock() {
	m.register("PUT", "/api/v1/categories/1", 200, "validate_update_category.json")
}

// RegisterDeleteCategoryMock registers DELETE /api/v1/categories/1 → 204.
func (m *CategoriesMock) RegisterDeleteCategoryMock() {
	m.register("DELETE", "/api/v1/categories/1", 204, "")
}

// RegisterDeleteCategoriesBulkMock registers POST /api/v1/categories/delete-multiple → 204.
func (m *CategoriesMock) RegisterDeleteCategoriesBulkMock() {
	m.register("POST", "/api/v1/categories/delete-multiple", 204, "")
}

// RegisterGetCategoryHistoryMock registers GET /api/v1/categories/1/history → 200.
func (m *CategoriesMock) RegisterGetCategoryHistoryMock() {
	m.register("GET", "/api/v1/categories/1/history", 200, "validate_get_history.json")
}

// RegisterAddCategoryHistoryNotesMock registers POST /api/v1/categories/1/history → 201.
func (m *CategoriesMock) RegisterAddCategoryHistoryNotesMock() {
	m.register("POST", "/api/v1/categories/1/history", 201, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /api/v1/categories/999 → 404.
func (m *CategoriesMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/api/v1/categories/999", 404, "error_not_found.json")
}

// RegisterConflictErrorMock registers POST /api/v1/categories → 409
// when the caller wishes to simulate a duplicate-name conflict.
func (m *CategoriesMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/api/v1/categories", 409, "error_conflict.json")
}

// ---- interfaces.HTTPClient implementation ----

func (m *CategoriesMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *CategoriesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *CategoriesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *CategoriesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *CategoriesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *CategoriesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *CategoriesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *CategoriesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *CategoriesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *CategoriesMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *CategoriesMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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

func (m *CategoriesMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *CategoriesMock) InvalidateToken() error                    { return nil }
func (m *CategoriesMock) KeepAliveToken() error                     { return nil }
func (m *CategoriesMock) GetLogger() *zap.Logger                    { return m.logger }

// ---- Internal helpers ----

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 204 No Content responses).
func (m *CategoriesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("CategoriesMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// registerError stores an error response. The fixture body is returned alongside the error.
func (m *CategoriesMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("CategoriesMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

// dispatch looks up the registered response and either unmarshals the body
// into result or returns an error depending on the registration type.
func (m *CategoriesMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/json"}},
			Body:       []byte(`{"code":"NOT-FOUND","message":"no mock registered"}`),
		}, fmt.Errorf("CategoriesMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("CategoriesMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

// loadMockResponse reads a JSON fixture file from the mocks/ directory
// adjacent to the test being run.
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
