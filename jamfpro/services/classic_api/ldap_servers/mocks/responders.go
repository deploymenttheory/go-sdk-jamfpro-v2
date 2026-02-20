package mocks

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"go.uber.org/zap"
)

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	// errMsg causes the mock to return an error (simulating API or transport failures).
	errMsg string
}

// LDAPServersMock is a test double implementing interfaces.HTTPClient for Classic API LDAP servers.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Unlike Jamf Pro API mocks which use JSON, Classic API mocks use XML for
// serialization to match the Classic API wire format.
type LDAPServersMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string // captures the rsqlQuery from the most recent Get call
}

// NewLDAPServersMock returns an empty mock ready for response registration.
func NewLDAPServersMock() *LDAPServersMock {
	return &LDAPServersMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *LDAPServersMock) RegisterMocks() {
	m.RegisterListLDAPServersMock()
	m.RegisterGetLDAPServerByIDMock()
	m.RegisterGetLDAPServerByNameMock()
	m.RegisterCreateLDAPServerMock()
	m.RegisterUpdateLDAPServerByIDMock()
	m.RegisterUpdateLDAPServerByNameMock()
	m.RegisterDeleteLDAPServerByIDMock()
	m.RegisterDeleteLDAPServerByNameMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *LDAPServersMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListLDAPServersMock registers GET /JSSResource/ldapservers → 200.
func (m *LDAPServersMock) RegisterListLDAPServersMock() {
	m.register("GET", "/JSSResource/ldapservers", 200, "validate_list_ldap_servers.xml")
}

// RegisterGetLDAPServerByIDMock registers GET /JSSResource/ldapservers/id/1 → 200.
func (m *LDAPServersMock) RegisterGetLDAPServerByIDMock() {
	m.register("GET", "/JSSResource/ldapservers/id/1", 200, "validate_get_ldap_server.xml")
}

// RegisterGetLDAPServerByNameMock registers GET /JSSResource/ldapservers/name/Test LDAP Server → 200.
func (m *LDAPServersMock) RegisterGetLDAPServerByNameMock() {
	m.register("GET", "/JSSResource/ldapservers/name/Test LDAP Server", 200, "validate_get_ldap_server.xml")
}

// RegisterCreateLDAPServerMock registers POST /JSSResource/ldapservers/id/0 → 201.
func (m *LDAPServersMock) RegisterCreateLDAPServerMock() {
	m.register("POST", "/JSSResource/ldapservers/id/0", 201, "validate_create_ldap_server.xml")
}

// RegisterUpdateLDAPServerByIDMock registers PUT /JSSResource/ldapservers/id/1 → 200.
func (m *LDAPServersMock) RegisterUpdateLDAPServerByIDMock() {
	m.register("PUT", "/JSSResource/ldapservers/id/1", 200, "validate_update_ldap_server.xml")
}

// RegisterUpdateLDAPServerByNameMock registers PUT /JSSResource/ldapservers/name/Test LDAP Server → 200.
func (m *LDAPServersMock) RegisterUpdateLDAPServerByNameMock() {
	m.register("PUT", "/JSSResource/ldapservers/name/Test LDAP Server", 200, "validate_update_ldap_server.xml")
}

// RegisterDeleteLDAPServerByIDMock registers DELETE /JSSResource/ldapservers/id/1 → 200.
func (m *LDAPServersMock) RegisterDeleteLDAPServerByIDMock() {
	m.register("DELETE", "/JSSResource/ldapservers/id/1", 200, "")
}

// RegisterDeleteLDAPServerByNameMock registers DELETE /JSSResource/ldapservers/name/Test LDAP Server → 200.
func (m *LDAPServersMock) RegisterDeleteLDAPServerByNameMock() {
	m.register("DELETE", "/JSSResource/ldapservers/name/Test LDAP Server", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/ldapservers/id/999 → 404.
func (m *LDAPServersMock) RegisterNotFoundErrorMock() {
	body := []byte("<br>An error has occurred.<br>Resource not found<br><br>")
	m.responses["GET:/JSSResource/ldapservers/id/999"] = registeredResponse{
		statusCode: 404,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (404): Resource not found",
	}
}

// RegisterConflictErrorMock registers POST /JSSResource/ldapservers/id/0 → 409
// when the caller wishes to simulate a duplicate-name conflict.
func (m *LDAPServersMock) RegisterConflictErrorMock() {
	body := []byte("<br>An error has occurred.<br>An LDAP server with that name already exists.<br><br>")
	m.responses["POST:/JSSResource/ldapservers/id/0"] = registeredResponse{
		statusCode: 409,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (409): An LDAP server with that name already exists",
	}
}

// ---- interfaces.HTTPClient implementation ----

func (m *LDAPServersMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *LDAPServersMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *LDAPServersMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *LDAPServersMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *LDAPServersMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *LDAPServersMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *LDAPServersMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *LDAPServersMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *LDAPServersMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *LDAPServersMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *LDAPServersMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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

func (m *LDAPServersMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *LDAPServersMock) InvalidateToken() error                     { return nil }
func (m *LDAPServersMock) KeepAliveToken() error                      { return nil }
func (m *LDAPServersMock) GetLogger() *zap.Logger                     { return m.logger }

// ---- Internal helpers ----

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *LDAPServersMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("LDAPServersMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *LDAPServersMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {mime.ApplicationXML}},
			Body:       []byte(`<error>no mock registered</error>`),
		}, fmt.Errorf("LDAPServersMock: no response registered for %s %s", method, path)
	}

	resp := &interfaces.Response{
		StatusCode: r.statusCode,
		Status:     fmt.Sprintf("%d", r.statusCode),
		Headers:    http.Header{"Content-Type": {mime.ApplicationXML}},
		Body:       r.rawBody,
	}

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := xml.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("LDAPServersMock: unmarshal into result: %w", err)
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
