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

// UserGroupsMock is a test double implementing interfaces.HTTPClient for Classic API user groups.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Classic API mocks use XML for serialization to match the Classic API wire format.
type UserGroupsMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewUserGroupsMock returns an empty mock ready for response registration.
func NewUserGroupsMock() *UserGroupsMock {
	return &UserGroupsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *UserGroupsMock) RegisterMocks() {
	m.RegisterListUserGroupsMock()
	m.RegisterGetUserGroupByIDMock()
	m.RegisterGetUserGroupByNameMock()
	m.RegisterCreateUserGroupMock()
	m.RegisterUpdateUserGroupByIDMock()
	m.RegisterUpdateUserGroupByNameMock()
	m.RegisterDeleteUserGroupByIDMock()
	m.RegisterDeleteUserGroupByNameMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *UserGroupsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListUserGroupsMock registers GET /JSSResource/usergroups → 200.
func (m *UserGroupsMock) RegisterListUserGroupsMock() {
	m.register("GET", "/JSSResource/usergroups", 200, "validate_list_user_groups.xml")
}

// RegisterGetUserGroupByIDMock registers GET /JSSResource/usergroups/id/1 → 200.
func (m *UserGroupsMock) RegisterGetUserGroupByIDMock() {
	m.register("GET", "/JSSResource/usergroups/id/1", 200, "validate_get_user_group.xml")
}

// RegisterGetUserGroupByNameMock registers GET /JSSResource/usergroups/name/All Users → 200.
func (m *UserGroupsMock) RegisterGetUserGroupByNameMock() {
	m.register("GET", "/JSSResource/usergroups/name/All Users", 200, "validate_get_user_group.xml")
}

// RegisterCreateUserGroupMock registers POST /JSSResource/usergroups/id/0 → 201.
func (m *UserGroupsMock) RegisterCreateUserGroupMock() {
	m.register("POST", "/JSSResource/usergroups/id/0", 201, "validate_create_user_group.xml")
}

// RegisterUpdateUserGroupByIDMock registers PUT /JSSResource/usergroups/id/1 → 200.
func (m *UserGroupsMock) RegisterUpdateUserGroupByIDMock() {
	m.register("PUT", "/JSSResource/usergroups/id/1", 200, "validate_update_user_group.xml")
}

// RegisterUpdateUserGroupByNameMock registers PUT /JSSResource/usergroups/name/All Users → 200.
func (m *UserGroupsMock) RegisterUpdateUserGroupByNameMock() {
	m.register("PUT", "/JSSResource/usergroups/name/All Users", 200, "validate_update_user_group.xml")
}

// RegisterDeleteUserGroupByIDMock registers DELETE /JSSResource/usergroups/id/1 → 200.
func (m *UserGroupsMock) RegisterDeleteUserGroupByIDMock() {
	m.register("DELETE", "/JSSResource/usergroups/id/1", 200, "")
}

// RegisterDeleteUserGroupByNameMock registers DELETE /JSSResource/usergroups/name/All Users → 200.
func (m *UserGroupsMock) RegisterDeleteUserGroupByNameMock() {
	m.register("DELETE", "/JSSResource/usergroups/name/All Users", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/usergroups/id/999 → 404.
func (m *UserGroupsMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/JSSResource/usergroups/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

// RegisterConflictErrorMock registers POST /JSSResource/usergroups/id/0 → 409.
func (m *UserGroupsMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/JSSResource/usergroups/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A user group with that name already exists")
}

// ---- interfaces.HTTPClient implementation ----

func (m *UserGroupsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *UserGroupsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *UserGroupsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *UserGroupsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *UserGroupsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *UserGroupsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *UserGroupsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *UserGroupsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *UserGroupsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *UserGroupsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *UserGroupsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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

func (m *UserGroupsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *UserGroupsMock) InvalidateToken() error                     { return nil }
func (m *UserGroupsMock) KeepAliveToken() error                      { return nil }
func (m *UserGroupsMock) GetLogger() *zap.Logger                     { return m.logger }

// ---- Internal helpers ----

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *UserGroupsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("UserGroupsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// registerError stores an error response with externalized XML body.
func (m *UserGroupsMock) registerError(method, path string, statusCode int, fixture, errMsg string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("UserGroupsMock: failed to load error fixture %q: %v", fixture, err))
	}
	m.responses[method+":"+path] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
		errMsg:     errMsg,
	}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *UserGroupsMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/xml"}},
			Body:       []byte(`<error>no mock registered</error>`),
		}, fmt.Errorf("UserGroupsMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("UserGroupsMock: unmarshal into result: %w", err)
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
