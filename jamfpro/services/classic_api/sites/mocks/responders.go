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
	"go.uber.org/zap"
)

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	// errMsg causes the mock to return an error (simulating API or transport failures).
	errMsg string
}

// SitesMock is a test double implementing interfaces.HTTPClient for Classic API sites.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Unlike Jamf Pro API mocks which use JSON, Classic API mocks use XML for
// serialization to match the Classic API wire format.
type SitesMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string // captures the rsqlQuery from the most recent Get call
}

// NewSitesMock returns an empty mock ready for response registration.
func NewSitesMock() *SitesMock {
	return &SitesMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *SitesMock) RegisterMocks() {
	m.RegisterListSitesMock()
	m.RegisterGetSiteByIDMock()
	m.RegisterGetSiteByNameMock()
	m.RegisterCreateSiteMock()
	m.RegisterUpdateSiteByIDMock()
	m.RegisterUpdateSiteByNameMock()
	m.RegisterDeleteSiteByIDMock()
	m.RegisterDeleteSiteByNameMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *SitesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListSitesMock registers GET /JSSResource/sites → 200.
func (m *SitesMock) RegisterListSitesMock() {
	m.register("GET", "/JSSResource/sites", 200, "validate_list_sites.xml")
}

// RegisterGetSiteByIDMock registers GET /JSSResource/sites/id/1 → 200.
func (m *SitesMock) RegisterGetSiteByIDMock() {
	m.register("GET", "/JSSResource/sites/id/1", 200, "validate_get_site.xml")
}

// RegisterGetSiteByNameMock registers GET /JSSResource/sites/name/Main%20Campus → 200.
func (m *SitesMock) RegisterGetSiteByNameMock() {
	m.register("GET", "/JSSResource/sites/name/Main Campus", 200, "validate_get_site.xml")
}

// RegisterCreateSiteMock registers POST /JSSResource/sites/id/0 → 201.
func (m *SitesMock) RegisterCreateSiteMock() {
	m.register("POST", "/JSSResource/sites/id/0", 201, "validate_create_site.xml")
}

// RegisterUpdateSiteByIDMock registers PUT /JSSResource/sites/id/1 → 200.
func (m *SitesMock) RegisterUpdateSiteByIDMock() {
	m.register("PUT", "/JSSResource/sites/id/1", 200, "validate_update_site.xml")
}

// RegisterUpdateSiteByNameMock registers PUT /JSSResource/sites/name/Main Campus → 200.
func (m *SitesMock) RegisterUpdateSiteByNameMock() {
	m.register("PUT", "/JSSResource/sites/name/Main Campus", 200, "validate_update_site.xml")
}

// RegisterDeleteSiteByIDMock registers DELETE /JSSResource/sites/id/1 → 200.
func (m *SitesMock) RegisterDeleteSiteByIDMock() {
	m.register("DELETE", "/JSSResource/sites/id/1", 200, "")
}

// RegisterDeleteSiteByNameMock registers DELETE /JSSResource/sites/name/Main Campus → 200.
func (m *SitesMock) RegisterDeleteSiteByNameMock() {
	m.register("DELETE", "/JSSResource/sites/name/Main Campus", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/sites/id/999 → 404.
func (m *SitesMock) RegisterNotFoundErrorMock() {
	body := []byte("<br>An error has occurred.<br>Resource not found<br><br>")
	m.responses["GET:/JSSResource/sites/id/999"] = registeredResponse{
		statusCode: 404,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (404): Resource not found",
	}
}

// RegisterConflictErrorMock registers POST /JSSResource/sites/id/0 → 409
// when the caller wishes to simulate a duplicate-name conflict.
func (m *SitesMock) RegisterConflictErrorMock() {
	body := []byte("<br>An error has occurred.<br>A site with that name already exists.<br><br>")
	m.responses["POST:/JSSResource/sites/id/0"] = registeredResponse{
		statusCode: 409,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (409): A site with that name already exists",
	}
}

// ---- interfaces.HTTPClient implementation ----

func (m *SitesMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *SitesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *SitesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *SitesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *SitesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *SitesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *SitesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *SitesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *SitesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *SitesMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *SitesMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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

func (m *SitesMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *SitesMock) InvalidateToken() error                    { return nil }
func (m *SitesMock) KeepAliveToken() error                     { return nil }
func (m *SitesMock) GetLogger() *zap.Logger                    { return m.logger }

// ---- Internal helpers ----

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *SitesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("SitesMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *SitesMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/xml"}},
			Body:       []byte(`<error>no mock registered</error>`),
		}, fmt.Errorf("SitesMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("SitesMock: unmarshal into result: %w", err)
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
