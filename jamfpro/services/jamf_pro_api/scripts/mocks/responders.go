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

// ScriptsMock is a test double implementing interfaces.HTTPClient.
// Responses are keyed by "METHOD:path" and loaded from JSON fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
type ScriptsMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string // captures the rsqlQuery from the most recent Get call
}

// NewScriptsMock returns an empty mock ready for response registration.
func NewScriptsMock() *ScriptsMock {
	return &ScriptsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *ScriptsMock) RegisterMocks() {
	m.RegisterListScriptsMock()
	m.RegisterGetScriptMock()
	m.RegisterDownloadScriptMock()
	m.RegisterCreateScriptMock()
	m.RegisterUpdateScriptMock()
	m.RegisterDeleteScriptMock()
	m.RegisterGetScriptHistoryMock()
	m.RegisterAddScriptHistoryNotesMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *ScriptsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListScriptsMock registers GET /api/v1/scripts → 200.
func (m *ScriptsMock) RegisterListScriptsMock() {
	m.register("GET", "/api/v1/scripts", 200, "validate_list_scripts.json")
}

// RegisterListScriptsRSQLMock registers GET /api/v1/scripts → 200 with a
// single-result fixture that simulates a server-side RSQL-filtered response.
func (m *ScriptsMock) RegisterListScriptsRSQLMock() {
	m.register("GET", "/api/v1/scripts", 200, "validate_list_scripts_rsql.json")
}

// RegisterGetScriptMock registers GET /api/v1/scripts/1 → 200.
func (m *ScriptsMock) RegisterGetScriptMock() {
	m.register("GET", "/api/v1/scripts/1", 200, "validate_get_script.json")
}

// RegisterDownloadScriptMock registers GET /api/v1/scripts/1/download → 200.
func (m *ScriptsMock) RegisterDownloadScriptMock() {
	m.register("GET", "/api/v1/scripts/1/download", 200, "validate_download_script.txt")
}

// RegisterCreateScriptMock registers POST /api/v1/scripts → 201.
func (m *ScriptsMock) RegisterCreateScriptMock() {
	m.register("POST", "/api/v1/scripts", 201, "validate_create_script.json")
}

// RegisterUpdateScriptMock registers PUT /api/v1/scripts/1 → 200.
func (m *ScriptsMock) RegisterUpdateScriptMock() {
	m.register("PUT", "/api/v1/scripts/1", 200, "validate_update_script.json")
}

// RegisterDeleteScriptMock registers DELETE /api/v1/scripts/1 → 204.
func (m *ScriptsMock) RegisterDeleteScriptMock() {
	m.register("DELETE", "/api/v1/scripts/1", 204, "")
}

// RegisterGetScriptHistoryMock registers GET /api/v1/scripts/1/history → 200.
func (m *ScriptsMock) RegisterGetScriptHistoryMock() {
	m.register("GET", "/api/v1/scripts/1/history", 200, "validate_get_history.json")
}

// RegisterAddScriptHistoryNotesMock registers POST /api/v1/scripts/1/history → 201.
func (m *ScriptsMock) RegisterAddScriptHistoryNotesMock() {
	m.register("POST", "/api/v1/scripts/1/history", 201, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /api/v1/scripts/999 → 404.
func (m *ScriptsMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/api/v1/scripts/999", 404, "error_not_found.json")
}

// RegisterConflictErrorMock registers POST /api/v1/scripts → 409
// when the caller wishes to simulate a duplicate-name conflict.
func (m *ScriptsMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/api/v1/scripts", 409, "error_conflict.json")
}

// ---- interfaces.HTTPClient implementation ----

func (m *ScriptsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *ScriptsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ScriptsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ScriptsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ScriptsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ScriptsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *ScriptsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *ScriptsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *ScriptsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *ScriptsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *ScriptsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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

func (m *ScriptsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *ScriptsMock) InvalidateToken() error                    { return nil }
func (m *ScriptsMock) KeepAliveToken() error                     { return nil }
func (m *ScriptsMock) GetLogger() *zap.Logger                    { return m.logger }

// ---- Internal helpers ----

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 204 No Content responses).
func (m *ScriptsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("ScriptsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// registerError stores an error response. The fixture body is returned alongside the error.
func (m *ScriptsMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("ScriptsMock: failed to load error fixture %q: %v", fixture, err))
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
func (m *ScriptsMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/json"}},
			Body:       []byte(`{"code":"NOT-FOUND","message":"no mock registered"}`),
		}, fmt.Errorf("ScriptsMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("ScriptsMock: unmarshal into result: %w", err)
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
