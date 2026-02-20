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

// BYOProfilesMock is a test double implementing interfaces.HTTPClient for Classic API BYO profiles.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Unlike Jamf Pro API mocks which use JSON, Classic API mocks use XML for
// serialization to match the Classic API wire format.
type BYOProfilesMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string // captures the rsqlQuery from the most recent Get call
}

// NewBYOProfilesMock returns an empty mock ready for response registration.
func NewBYOProfilesMock() *BYOProfilesMock {
	return &BYOProfilesMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *BYOProfilesMock) RegisterMocks() {
	m.RegisterListBYOProfilesMock()
	m.RegisterGetBYOProfileByIDMock()
	m.RegisterGetBYOProfileByNameMock()
	m.RegisterCreateBYOProfileMock()
	m.RegisterUpdateBYOProfileByIDMock()
	m.RegisterUpdateBYOProfileByNameMock()
	m.RegisterDeleteBYOProfileByIDMock()
	m.RegisterDeleteBYOProfileByNameMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *BYOProfilesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListBYOProfilesMock registers GET /JSSResource/byoprofiles → 200.
func (m *BYOProfilesMock) RegisterListBYOProfilesMock() {
	m.register("GET", "/JSSResource/byoprofiles", 200, "validate_list_byoprofiles.xml")
}

// RegisterGetBYOProfileByIDMock registers GET /JSSResource/byoprofiles/id/1 → 200.
func (m *BYOProfilesMock) RegisterGetBYOProfileByIDMock() {
	m.register("GET", "/JSSResource/byoprofiles/id/1", 200, "validate_get_byoprofile.xml")
}

// RegisterGetBYOProfileByNameMock registers GET /JSSResource/byoprofiles/name/Test BYO Profile 1 → 200.
func (m *BYOProfilesMock) RegisterGetBYOProfileByNameMock() {
	m.register("GET", "/JSSResource/byoprofiles/name/Test BYO Profile 1", 200, "validate_get_byoprofile.xml")
}

// RegisterCreateBYOProfileMock registers POST /JSSResource/byoprofiles/id/0 → 201.
func (m *BYOProfilesMock) RegisterCreateBYOProfileMock() {
	m.register("POST", "/JSSResource/byoprofiles/id/0", 201, "validate_create_byoprofile.xml")
}

// RegisterUpdateBYOProfileByIDMock registers PUT /JSSResource/byoprofiles/id/1 → 200.
func (m *BYOProfilesMock) RegisterUpdateBYOProfileByIDMock() {
	m.register("PUT", "/JSSResource/byoprofiles/id/1", 200, "validate_update_byoprofile.xml")
}

// RegisterUpdateBYOProfileByNameMock registers PUT /JSSResource/byoprofiles/name/Test BYO Profile 1 → 200.
func (m *BYOProfilesMock) RegisterUpdateBYOProfileByNameMock() {
	m.register("PUT", "/JSSResource/byoprofiles/name/Test BYO Profile 1", 200, "validate_update_byoprofile.xml")
}

// RegisterDeleteBYOProfileByIDMock registers DELETE /JSSResource/byoprofiles/id/1 → 200.
func (m *BYOProfilesMock) RegisterDeleteBYOProfileByIDMock() {
	m.register("DELETE", "/JSSResource/byoprofiles/id/1", 200, "")
}

// RegisterDeleteBYOProfileByNameMock registers DELETE /JSSResource/byoprofiles/name/Test BYO Profile 1 → 200.
func (m *BYOProfilesMock) RegisterDeleteBYOProfileByNameMock() {
	m.register("DELETE", "/JSSResource/byoprofiles/name/Test BYO Profile 1", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/byoprofiles/id/999 → 404.
func (m *BYOProfilesMock) RegisterNotFoundErrorMock() {
	body := []byte("<br>An error has occurred.<br>Resource not found<br><br>")
	m.responses["GET:/JSSResource/byoprofiles/id/999"] = registeredResponse{
		statusCode: 404,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (404): Resource not found",
	}
}

// RegisterConflictErrorMock registers POST /JSSResource/byoprofiles/id/0 → 409
// when the caller wishes to simulate a duplicate-name conflict.
func (m *BYOProfilesMock) RegisterConflictErrorMock() {
	body := []byte("<br>An error has occurred.<br>A BYO profile with that name already exists.<br><br>")
	m.responses["POST:/JSSResource/byoprofiles/id/0"] = registeredResponse{
		statusCode: 409,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (409): A BYO profile with that name already exists",
	}
}

// ---- interfaces.HTTPClient implementation ----

func (m *BYOProfilesMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *BYOProfilesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *BYOProfilesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *BYOProfilesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *BYOProfilesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *BYOProfilesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *BYOProfilesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *BYOProfilesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *BYOProfilesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *BYOProfilesMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *BYOProfilesMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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

func (m *BYOProfilesMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *BYOProfilesMock) InvalidateToken() error                     { return nil }
func (m *BYOProfilesMock) KeepAliveToken() error                      { return nil }
func (m *BYOProfilesMock) GetLogger() *zap.Logger                     { return m.logger }

// ---- Internal helpers ----

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *BYOProfilesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("BYOProfilesMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *BYOProfilesMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {mime.ApplicationXML}},
			Body:       []byte(`<error>no mock registered</error>`),
		}, fmt.Errorf("BYOProfilesMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("BYOProfilesMock: unmarshal into result: %w", err)
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
