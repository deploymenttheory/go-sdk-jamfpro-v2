package mocks

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
)

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// UsersMock is a test double implementing interfaces.HTTPClient for Classic API users.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Classic API mocks use XML for serialization to match the Classic API wire format.
type UsersMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
	LastRSQLQuery map[string]string
}

// NewUsersMock returns an empty mock ready for response registration.
func NewUsersMock() *UsersMock {
	return &UsersMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *UsersMock) RegisterMocks() {
	m.RegisterListUsersMock()
	m.RegisterGetUserByIDMock()
	m.RegisterGetUserByNameMock()
	m.RegisterGetUserByEmailMock()
	m.RegisterCreateUserMock()
	m.RegisterUpdateUserByIDMock()
	m.RegisterUpdateUserByNameMock()
	m.RegisterUpdateUserByEmailMock()
	m.RegisterDeleteUserByIDMock()
	m.RegisterDeleteUserByNameMock()
	m.RegisterDeleteUserByEmailMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *UsersMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListUsersMock registers GET /JSSResource/users → 200.
func (m *UsersMock) RegisterListUsersMock() {
	m.register("GET", "/JSSResource/users", 200, "validate_list_users.xml")
}

// RegisterGetUserByIDMock registers GET /JSSResource/users/id/1 → 200.
func (m *UsersMock) RegisterGetUserByIDMock() {
	m.register("GET", "/JSSResource/users/id/1", 200, "validate_get_user.xml")
}

// RegisterGetUserByNameMock registers GET /JSSResource/users/name/admin → 200.
func (m *UsersMock) RegisterGetUserByNameMock() {
	m.register("GET", "/JSSResource/users/name/admin", 200, "validate_get_user.xml")
}

// RegisterGetUserByEmailMock registers GET /JSSResource/users/email/{admin@example.com|admin%40example.com} → 200.
func (m *UsersMock) RegisterGetUserByEmailMock() {
	m.register("GET", "/JSSResource/users/email/admin@example.com", 200, "validate_get_user.xml")
	m.register("GET", "/JSSResource/users/email/admin%40example.com", 200, "validate_get_user.xml")
}

// RegisterCreateUserMock registers POST /JSSResource/users/id/0 → 201.
func (m *UsersMock) RegisterCreateUserMock() {
	m.register("POST", "/JSSResource/users/id/0", 201, "validate_create_user.xml")
}

// RegisterUpdateUserByIDMock registers PUT /JSSResource/users/id/1 → 200.
func (m *UsersMock) RegisterUpdateUserByIDMock() {
	m.register("PUT", "/JSSResource/users/id/1", 200, "validate_update_user.xml")
}

// RegisterUpdateUserByNameMock registers PUT /JSSResource/users/name/admin → 200.
func (m *UsersMock) RegisterUpdateUserByNameMock() {
	m.register("PUT", "/JSSResource/users/name/admin", 200, "validate_update_user.xml")
}

// RegisterUpdateUserByEmailMock registers PUT /JSSResource/users/email/{admin@example.com|admin%40example.com} → 200.
func (m *UsersMock) RegisterUpdateUserByEmailMock() {
	m.register("PUT", "/JSSResource/users/email/admin@example.com", 200, "validate_update_user.xml")
	m.register("PUT", "/JSSResource/users/email/admin%40example.com", 200, "validate_update_user.xml")
}

// RegisterDeleteUserByIDMock registers DELETE /JSSResource/users/id/1 → 200.
func (m *UsersMock) RegisterDeleteUserByIDMock() {
	m.register("DELETE", "/JSSResource/users/id/1", 200, "")
}

// RegisterDeleteUserByNameMock registers DELETE /JSSResource/users/name/admin → 200.
func (m *UsersMock) RegisterDeleteUserByNameMock() {
	m.register("DELETE", "/JSSResource/users/name/admin", 200, "")
}

// RegisterDeleteUserByEmailMock registers DELETE /JSSResource/users/email/{admin@example.com|admin%40example.com} → 200.
func (m *UsersMock) RegisterDeleteUserByEmailMock() {
	m.register("DELETE", "/JSSResource/users/email/admin@example.com", 200, "")
	m.register("DELETE", "/JSSResource/users/email/admin%40example.com", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/users/id/999 → 404.
func (m *UsersMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/JSSResource/users/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

// RegisterConflictErrorMock registers POST /JSSResource/users/id/0 → 409.
func (m *UsersMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/JSSResource/users/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A user with that name already exists")
}

// ---- interfaces.HTTPClient implementation ----

func (m *UsersMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *UsersMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *UsersMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *UsersMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *UsersMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *UsersMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *UsersMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *UsersMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *UsersMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *UsersMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *UsersMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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

func (m *UsersMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *UsersMock) InvalidateToken() error                     { return nil }
func (m *UsersMock) KeepAliveToken() error                      { return nil }
func (m *UsersMock) GetLogger() *zap.Logger                     { return m.logger }

// ---- Internal helpers ----

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *UsersMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("UsersMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// registerError stores an error response with externalized XML body.
func (m *UsersMock) registerError(method, path string, statusCode int, fixture, errMsg string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("UsersMock: failed to load error fixture %q: %v", fixture, err))
	}
	m.responses[method+":"+path] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
		errMsg:     errMsg,
	}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *UsersMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/xml"}},
			Body:       []byte(`<error>no mock registered</error>`),
		}, fmt.Errorf("UsersMock: no response registered for %s %s", method, path)
	}

	resp := &interfaces.Response{
		StatusCode: r.statusCode,
		Status:     fmt.Sprintf("%d", r.statusCode),
		Headers:    http.Header{"Content-Type": {"application/xml"}},
		Body:       r.rawBody,
	}

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := xml.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("UsersMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

// loadMockResponse reads an XML fixture file from the mocks/ directory
// adjacent to this file, so it works regardless of the test working directory.
func loadMockResponse(filename string) ([]byte, error) {
	_, callerPath, _, ok := runtime.Caller(1)
	if !ok {
		return nil, fmt.Errorf("runtime.Caller failed")
	}
	dir := filepath.Dir(callerPath)
	data, err := os.ReadFile(filepath.Join(dir, filename))
	if err != nil {
		return nil, fmt.Errorf("read fixture %s: %w", filename, err)
	}
	return data, nil
}
