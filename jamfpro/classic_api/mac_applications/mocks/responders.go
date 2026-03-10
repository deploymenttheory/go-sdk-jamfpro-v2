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

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"go.uber.org/zap"
	"resty.dev/v3"

	mockhelpers "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// MacApplicationsMock is a test double implementing client.Client for Classic API Mac applications.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Classic API mocks use XML for serialization to match the Classic API wire format.
type MacApplicationsMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewMacApplicationsMock returns an empty mock ready for response registration.
func NewMacApplicationsMock() *MacApplicationsMock {
	return &MacApplicationsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *MacApplicationsMock) RegisterMocks() {
	m.RegisterListMacApplicationsMock()
	m.RegisterGetMacApplicationByIDMock()
	m.RegisterGetMacApplicationByNameMock()
	m.RegisterGetMacApplicationByIDAndSubsetMock()
	m.RegisterGetMacApplicationByNameAndSubsetMock()
	m.RegisterCreateMacApplicationMock()
	m.RegisterUpdateMacApplicationByIDMock()
	m.RegisterUpdateMacApplicationByNameMock()
	m.RegisterDeleteMacApplicationByIDMock()
	m.RegisterDeleteMacApplicationByNameMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *MacApplicationsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListMacApplicationsMock registers GET /JSSResource/macapplications → 200.
func (m *MacApplicationsMock) RegisterListMacApplicationsMock() {
	m.register("GET", "/JSSResource/macapplications", 200, "validate_list_mac_applications.xml")
}

// RegisterGetMacApplicationByIDMock registers GET /JSSResource/macapplications/id/1 → 200.
func (m *MacApplicationsMock) RegisterGetMacApplicationByIDMock() {
	m.register("GET", "/JSSResource/macapplications/id/1", 200, "validate_get_mac_application.xml")
}

// RegisterGetMacApplicationByNameMock registers GET /JSSResource/macapplications/name/Sample Mac App → 200.
func (m *MacApplicationsMock) RegisterGetMacApplicationByNameMock() {
	m.register("GET", "/JSSResource/macapplications/name/Sample Mac App", 200, "validate_get_mac_application.xml")
}

// RegisterGetMacApplicationByIDAndSubsetMock registers GET /JSSResource/macapplications/id/1/subset/General → 200.
func (m *MacApplicationsMock) RegisterGetMacApplicationByIDAndSubsetMock() {
	m.register("GET", "/JSSResource/macapplications/id/1/subset/General", 200, "validate_get_mac_application.xml")
}

// RegisterGetMacApplicationByNameAndSubsetMock registers GET /JSSResource/macapplications/name/Sample Mac App/subset/General → 200.
func (m *MacApplicationsMock) RegisterGetMacApplicationByNameAndSubsetMock() {
	m.register("GET", "/JSSResource/macapplications/name/Sample Mac App/subset/General", 200, "validate_get_mac_application.xml")
}

// RegisterCreateMacApplicationMock registers POST /JSSResource/macapplications/id/0 → 201.
func (m *MacApplicationsMock) RegisterCreateMacApplicationMock() {
	m.register("POST", "/JSSResource/macapplications/id/0", 201, "validate_create_mac_application.xml")
}

// RegisterUpdateMacApplicationByIDMock registers PUT /JSSResource/macapplications/id/1 → 200.
func (m *MacApplicationsMock) RegisterUpdateMacApplicationByIDMock() {
	m.register("PUT", "/JSSResource/macapplications/id/1", 200, "validate_update_mac_application.xml")
}

// RegisterUpdateMacApplicationByNameMock registers PUT /JSSResource/macapplications/name/Sample Mac App → 200.
func (m *MacApplicationsMock) RegisterUpdateMacApplicationByNameMock() {
	m.register("PUT", "/JSSResource/macapplications/name/Sample Mac App", 200, "validate_update_mac_application.xml")
}

// RegisterDeleteMacApplicationByIDMock registers DELETE /JSSResource/macapplications/id/1 → 200.
func (m *MacApplicationsMock) RegisterDeleteMacApplicationByIDMock() {
	m.register("DELETE", "/JSSResource/macapplications/id/1", 200, "")
}

// RegisterDeleteMacApplicationByNameMock registers DELETE /JSSResource/macapplications/name/Sample Mac App → 200.
func (m *MacApplicationsMock) RegisterDeleteMacApplicationByNameMock() {
	m.register("DELETE", "/JSSResource/macapplications/name/Sample Mac App", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/macapplications/id/999 → 404.
func (m *MacApplicationsMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/JSSResource/macapplications/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

// RegisterConflictErrorMock registers POST /JSSResource/macapplications/id/0 → 409.
func (m *MacApplicationsMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/JSSResource/macapplications/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A Mac application with that name already exists")
}

// ---- client.Client implementation ----

func (m *MacApplicationsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	_ = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *MacApplicationsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MacApplicationsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MacApplicationsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MacApplicationsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MacApplicationsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *MacApplicationsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *MacApplicationsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MacApplicationsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MacApplicationsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	_ = rsqlQuery
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *MacApplicationsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	_ = rsqlQuery
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

func (m *MacApplicationsMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *MacApplicationsMock) InvalidateToken() error                { return nil }
func (m *MacApplicationsMock) KeepAliveToken() error                 { return nil }
func (m *MacApplicationsMock) GetLogger() *zap.Logger                { return m.logger }

// ---- Internal helpers ----

// registerError stores an error response with externalized XML body.
func (m *MacApplicationsMock) registerError(method, path string, statusCode int, fixture, errMsg string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("MacApplicationsMock: failed to load error fixture %q: %v", fixture, err))
	}
	m.responses[method+":"+path] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
		errMsg:     errMsg,
	}
}

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *MacApplicationsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("MacApplicationsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *MacApplicationsMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		headers := http.Header{"Content-Type": {constants.ApplicationXML}}
		return mockhelpers.NewMockResponse(http.StatusNotFound, headers, []byte(`<error>no mock registered</error>`)), fmt.Errorf("MacApplicationsMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {constants.ApplicationXML}}
	resp := mockhelpers.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := xml.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("MacApplicationsMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

// loadMockResponse reads an XML fixture file from the mocks package directory.
func loadMockResponse(filename string) ([]byte, error) {
	_, callerFile, _, ok := runtime.Caller(1)
	if !ok {
		return nil, fmt.Errorf("could not get caller path")
	}
	dir := filepath.Dir(callerFile)
	data, err := os.ReadFile(filepath.Join(dir, filename))
	if err != nil {
		return nil, fmt.Errorf("read fixture %s: %w", filename, err)
	}
	return data, nil
}
