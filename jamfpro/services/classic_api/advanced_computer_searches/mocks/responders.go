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

// AdvancedComputerSearchesMock is a test double implementing interfaces.HTTPClient for Classic API advanced computer searches.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Unlike Jamf Pro API mocks which use JSON, Classic API mocks use XML for
// serialization to match the Classic API wire format.
type AdvancedComputerSearchesMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string // captures the rsqlQuery from the most recent Get call
}

// NewAdvancedComputerSearchesMock returns an empty mock ready for response registration.
func NewAdvancedComputerSearchesMock() *AdvancedComputerSearchesMock {
	return &AdvancedComputerSearchesMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *AdvancedComputerSearchesMock) RegisterMocks() {
	m.RegisterListAdvancedComputerSearchesMock()
	m.RegisterGetAdvancedComputerSearchByIDMock()
	m.RegisterGetAdvancedComputerSearchByNameMock()
	m.RegisterCreateAdvancedComputerSearchMock()
	m.RegisterUpdateAdvancedComputerSearchByIDMock()
	m.RegisterUpdateAdvancedComputerSearchByNameMock()
	m.RegisterDeleteAdvancedComputerSearchByIDMock()
	m.RegisterDeleteAdvancedComputerSearchByNameMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *AdvancedComputerSearchesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListAdvancedComputerSearchesMock registers GET /JSSResource/advancedcomputersearches → 200.
func (m *AdvancedComputerSearchesMock) RegisterListAdvancedComputerSearchesMock() {
	m.register("GET", "/JSSResource/advancedcomputersearches", 200, "validate_list_advanced_computer_searches.xml")
}

// RegisterGetAdvancedComputerSearchByIDMock registers GET /JSSResource/advancedcomputersearches/id/1 → 200.
func (m *AdvancedComputerSearchesMock) RegisterGetAdvancedComputerSearchByIDMock() {
	m.register("GET", "/JSSResource/advancedcomputersearches/id/1", 200, "validate_get_advanced_computer_search.xml")
}

// RegisterGetAdvancedComputerSearchByNameMock registers GET /JSSResource/advancedcomputersearches/name/Test Search → 200.
func (m *AdvancedComputerSearchesMock) RegisterGetAdvancedComputerSearchByNameMock() {
	m.register("GET", "/JSSResource/advancedcomputersearches/name/Test Search", 200, "validate_get_advanced_computer_search.xml")
}

// RegisterCreateAdvancedComputerSearchMock registers POST /JSSResource/advancedcomputersearches/id/0 → 201.
func (m *AdvancedComputerSearchesMock) RegisterCreateAdvancedComputerSearchMock() {
	m.register("POST", "/JSSResource/advancedcomputersearches/id/0", 201, "validate_create_advanced_computer_search.xml")
}

// RegisterUpdateAdvancedComputerSearchByIDMock registers PUT /JSSResource/advancedcomputersearches/id/1 → 200.
func (m *AdvancedComputerSearchesMock) RegisterUpdateAdvancedComputerSearchByIDMock() {
	m.register("PUT", "/JSSResource/advancedcomputersearches/id/1", 200, "validate_update_advanced_computer_search.xml")
}

// RegisterUpdateAdvancedComputerSearchByNameMock registers PUT /JSSResource/advancedcomputersearches/name/Test Search → 200.
func (m *AdvancedComputerSearchesMock) RegisterUpdateAdvancedComputerSearchByNameMock() {
	m.register("PUT", "/JSSResource/advancedcomputersearches/name/Test Search", 200, "validate_update_advanced_computer_search.xml")
}

// RegisterDeleteAdvancedComputerSearchByIDMock registers DELETE /JSSResource/advancedcomputersearches/id/1 → 200.
func (m *AdvancedComputerSearchesMock) RegisterDeleteAdvancedComputerSearchByIDMock() {
	m.register("DELETE", "/JSSResource/advancedcomputersearches/id/1", 200, "")
}

// RegisterDeleteAdvancedComputerSearchByNameMock registers DELETE /JSSResource/advancedcomputersearches/name/Test Search → 200.
func (m *AdvancedComputerSearchesMock) RegisterDeleteAdvancedComputerSearchByNameMock() {
	m.register("DELETE", "/JSSResource/advancedcomputersearches/name/Test Search", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/advancedcomputersearches/id/999 → 404.
func (m *AdvancedComputerSearchesMock) RegisterNotFoundErrorMock() {
	body := []byte("<br>An error has occurred.<br>Resource not found<br><br>")
	m.responses["GET:/JSSResource/advancedcomputersearches/id/999"] = registeredResponse{
		statusCode: 404,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (404): Resource not found",
	}
}

// RegisterConflictErrorMock registers POST /JSSResource/advancedcomputersearches/id/0 → 409
// when the caller wishes to simulate a duplicate-name conflict.
func (m *AdvancedComputerSearchesMock) RegisterConflictErrorMock() {
	body := []byte("<br>An error has occurred.<br>An advanced computer search with that name already exists.<br><br>")
	m.responses["POST:/JSSResource/advancedcomputersearches/id/0"] = registeredResponse{
		statusCode: 409,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (409): An advanced computer search with that name already exists",
	}
}

// ---- interfaces.HTTPClient implementation ----

func (m *AdvancedComputerSearchesMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *AdvancedComputerSearchesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *AdvancedComputerSearchesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *AdvancedComputerSearchesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *AdvancedComputerSearchesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *AdvancedComputerSearchesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *AdvancedComputerSearchesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *AdvancedComputerSearchesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *AdvancedComputerSearchesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *AdvancedComputerSearchesMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *AdvancedComputerSearchesMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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

func (m *AdvancedComputerSearchesMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *AdvancedComputerSearchesMock) InvalidateToken() error                     { return nil }
func (m *AdvancedComputerSearchesMock) KeepAliveToken() error                      { return nil }
func (m *AdvancedComputerSearchesMock) GetLogger() *zap.Logger                     { return m.logger }

// ---- Internal helpers ----

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *AdvancedComputerSearchesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("AdvancedComputerSearchesMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *AdvancedComputerSearchesMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {mime.ApplicationXML}},
			Body:       []byte(`<error>no mock registered</error>`),
		}, fmt.Errorf("AdvancedComputerSearchesMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("AdvancedComputerSearchesMock: unmarshal into result: %w", err)
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
