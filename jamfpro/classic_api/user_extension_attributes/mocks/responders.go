package mocks

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
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

// UserExtensionAttributesMock is a test double implementing transport.HTTPClient for Classic API user extension attributes.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Classic API mocks use XML for serialization to match the Classic API wire format.
type UserExtensionAttributesMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewUserExtensionAttributesMock returns an empty mock ready for response registration.
func NewUserExtensionAttributesMock() *UserExtensionAttributesMock {
	return &UserExtensionAttributesMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *UserExtensionAttributesMock) RegisterMocks() {
	m.RegisterListUserExtensionAttributesMock()
	m.RegisterGetUserExtensionAttributeByIDMock()
	m.RegisterGetUserExtensionAttributeByNameMock()
	m.RegisterCreateUserExtensionAttributeMock()
	m.RegisterUpdateUserExtensionAttributeByIDMock()
	m.RegisterUpdateUserExtensionAttributeByNameMock()
	m.RegisterDeleteUserExtensionAttributeByIDMock()
	m.RegisterDeleteUserExtensionAttributeByNameMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *UserExtensionAttributesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListUserExtensionAttributesMock registers GET /JSSResource/userextensionattributes → 200.
func (m *UserExtensionAttributesMock) RegisterListUserExtensionAttributesMock() {
	m.register("GET", "/JSSResource/userextensionattributes", 200, "validate_list_user_extension_attributes.xml")
}

// RegisterGetUserExtensionAttributeByIDMock registers GET /JSSResource/userextensionattributes/id/1 → 200.
func (m *UserExtensionAttributesMock) RegisterGetUserExtensionAttributeByIDMock() {
	m.register("GET", "/JSSResource/userextensionattributes/id/1", 200, "validate_get_user_extension_attribute.xml")
}

// RegisterGetUserExtensionAttributeByNameMock registers GET /JSSResource/userextensionattributes/name/Department → 200.
func (m *UserExtensionAttributesMock) RegisterGetUserExtensionAttributeByNameMock() {
	m.register("GET", "/JSSResource/userextensionattributes/name/Department", 200, "validate_get_user_extension_attribute.xml")
}

// RegisterCreateUserExtensionAttributeMock registers POST /JSSResource/userextensionattributes/id/0 → 201.
func (m *UserExtensionAttributesMock) RegisterCreateUserExtensionAttributeMock() {
	m.register("POST", "/JSSResource/userextensionattributes/id/0", 201, "validate_create_user_extension_attribute.xml")
}

// RegisterUpdateUserExtensionAttributeByIDMock registers PUT /JSSResource/userextensionattributes/id/1 → 200.
func (m *UserExtensionAttributesMock) RegisterUpdateUserExtensionAttributeByIDMock() {
	m.register("PUT", "/JSSResource/userextensionattributes/id/1", 200, "validate_update_user_extension_attribute.xml")
}

// RegisterUpdateUserExtensionAttributeByNameMock registers PUT /JSSResource/userextensionattributes/name/Department → 200.
func (m *UserExtensionAttributesMock) RegisterUpdateUserExtensionAttributeByNameMock() {
	m.register("PUT", "/JSSResource/userextensionattributes/name/Department", 200, "validate_update_user_extension_attribute.xml")
}

// RegisterDeleteUserExtensionAttributeByIDMock registers DELETE /JSSResource/userextensionattributes/id/1 → 200.
func (m *UserExtensionAttributesMock) RegisterDeleteUserExtensionAttributeByIDMock() {
	m.register("DELETE", "/JSSResource/userextensionattributes/id/1", 200, "")
}

// RegisterDeleteUserExtensionAttributeByNameMock registers DELETE /JSSResource/userextensionattributes/name/Department → 200.
func (m *UserExtensionAttributesMock) RegisterDeleteUserExtensionAttributeByNameMock() {
	m.register("DELETE", "/JSSResource/userextensionattributes/name/Department", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/userextensionattributes/id/999 → 404.
func (m *UserExtensionAttributesMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/JSSResource/userextensionattributes/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

// RegisterConflictErrorMock registers POST /JSSResource/userextensionattributes/id/0 → 409.
func (m *UserExtensionAttributesMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/JSSResource/userextensionattributes/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A user extension attribute with that name already exists")
}

// ---- transport.HTTPClient implementation ----

func (m *UserExtensionAttributesMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *UserExtensionAttributesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *UserExtensionAttributesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *UserExtensionAttributesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *UserExtensionAttributesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ transport.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *UserExtensionAttributesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *UserExtensionAttributesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *UserExtensionAttributesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *UserExtensionAttributesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *UserExtensionAttributesMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *UserExtensionAttributesMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil {
		body := resp.Bytes()
		if err := mergePage(body); err != nil {
			return resp, err
		}
	}
	return resp, nil
}

func (m *UserExtensionAttributesMock) RSQLBuilder() transport.RSQLFilterBuilder { return nil }
func (m *UserExtensionAttributesMock) InvalidateToken() error                    { return nil }
func (m *UserExtensionAttributesMock) KeepAliveToken() error                     { return nil }
func (m *UserExtensionAttributesMock) GetLogger() *zap.Logger                    { return m.logger }

// ---- Internal helpers ----

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *UserExtensionAttributesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("UserExtensionAttributesMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// registerError stores an error response with externalized XML body.
func (m *UserExtensionAttributesMock) registerError(method, path string, statusCode int, fixture, errMsg string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("UserExtensionAttributesMock: failed to load error fixture %q: %v", fixture, err))
	}
	m.responses[method+":"+path] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
		errMsg:     errMsg,
	}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *UserExtensionAttributesMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		headers := http.Header{"Content-Type": {constants.ApplicationXML}}
		return shared.NewMockResponse(http.StatusNotFound, headers, []byte(`<error>no mock registered</error>`)), fmt.Errorf("UserExtensionAttributesMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {constants.ApplicationXML}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := xml.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("UserExtensionAttributesMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

// loadMockResponse reads an XML fixture file from the mocks/ directory
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
