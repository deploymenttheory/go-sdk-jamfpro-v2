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
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
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

// ComputerGroupsMock is a test double implementing transport.HTTPClient for Classic API computer groups.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Classic API mocks use XML for serialization to match the Classic API wire format.
type ComputerGroupsMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewComputerGroupsMock returns an empty mock ready for response registration.
func NewComputerGroupsMock() *ComputerGroupsMock {
	return &ComputerGroupsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *ComputerGroupsMock) RegisterMocks() {
	m.RegisterListComputerGroupsMock()
	m.RegisterGetComputerGroupByIDMock()
	m.RegisterGetComputerGroupByNameMock()
	m.RegisterCreateComputerGroupMock()
	m.RegisterUpdateComputerGroupByIDMock()
	m.RegisterUpdateComputerGroupByNameMock()
	m.RegisterDeleteComputerGroupByIDMock()
	m.RegisterDeleteComputerGroupByNameMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *ComputerGroupsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListComputerGroupsMock registers GET /JSSResource/computergroups → 200.
func (m *ComputerGroupsMock) RegisterListComputerGroupsMock() {
	m.register("GET", "/JSSResource/computergroups", 200, "validate_list_computer_groups.xml")
}

// RegisterGetComputerGroupByIDMock registers GET /JSSResource/computergroups/id/1 → 200.
func (m *ComputerGroupsMock) RegisterGetComputerGroupByIDMock() {
	m.register("GET", "/JSSResource/computergroups/id/1", 200, "validate_get_computer_group.xml")
}

// RegisterGetComputerGroupByNameMock registers GET /JSSResource/computergroups/name/All Managed Clients → 200.
func (m *ComputerGroupsMock) RegisterGetComputerGroupByNameMock() {
	m.register("GET", "/JSSResource/computergroups/name/All Managed Clients", 200, "validate_get_computer_group.xml")
}

// RegisterCreateComputerGroupMock registers POST /JSSResource/computergroups/id/0 → 201.
func (m *ComputerGroupsMock) RegisterCreateComputerGroupMock() {
	m.register("POST", "/JSSResource/computergroups/id/0", 201, "validate_create_computer_group.xml")
}

// RegisterUpdateComputerGroupByIDMock registers PUT /JSSResource/computergroups/id/1 → 200.
func (m *ComputerGroupsMock) RegisterUpdateComputerGroupByIDMock() {
	m.register("PUT", "/JSSResource/computergroups/id/1", 200, "validate_update_computer_group.xml")
}

// RegisterUpdateComputerGroupByNameMock registers PUT /JSSResource/computergroups/name/All Managed Clients → 200.
func (m *ComputerGroupsMock) RegisterUpdateComputerGroupByNameMock() {
	m.register("PUT", "/JSSResource/computergroups/name/All Managed Clients", 200, "validate_update_computer_group.xml")
}

// RegisterDeleteComputerGroupByIDMock registers DELETE /JSSResource/computergroups/id/1 → 200.
func (m *ComputerGroupsMock) RegisterDeleteComputerGroupByIDMock() {
	m.register("DELETE", "/JSSResource/computergroups/id/1", 200, "")
}

// RegisterDeleteComputerGroupByNameMock registers DELETE /JSSResource/computergroups/name/All Managed Clients → 200.
func (m *ComputerGroupsMock) RegisterDeleteComputerGroupByNameMock() {
	m.register("DELETE", "/JSSResource/computergroups/name/All Managed Clients", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/computergroups/id/999 → 404.
func (m *ComputerGroupsMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/JSSResource/computergroups/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

// RegisterConflictErrorMock registers POST /JSSResource/computergroups/id/0 → 409.
func (m *ComputerGroupsMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/JSSResource/computergroups/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A computer group with that name already exists")
}

// ---- transport.HTTPClient implementation ----

func (m *ComputerGroupsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *ComputerGroupsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputerGroupsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputerGroupsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputerGroupsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ transport.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputerGroupsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *ComputerGroupsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *ComputerGroupsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *ComputerGroupsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *ComputerGroupsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *ComputerGroupsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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

func (m *ComputerGroupsMock) RSQLBuilder() transport.RSQLFilterBuilder { return nil }
func (m *ComputerGroupsMock) InvalidateToken() error                    { return nil }
func (m *ComputerGroupsMock) KeepAliveToken() error                     { return nil }
func (m *ComputerGroupsMock) GetLogger() *zap.Logger                    { return m.logger }

// ---- Internal helpers ----

// registerError stores an error response with externalized XML body.
func (m *ComputerGroupsMock) registerError(method, path string, statusCode int, fixture, errMsg string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("ComputerGroupsMock: failed to load error fixture %q: %v", fixture, err))
	}
	m.responses[method+":"+path] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
		errMsg:     errMsg,
	}
}

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *ComputerGroupsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("ComputerGroupsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *ComputerGroupsMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		headers := http.Header{"Content-Type": {mime.ApplicationXML}}
		return shared.NewMockResponse(http.StatusNotFound, headers, []byte(`<error>no mock registered</error>`)), fmt.Errorf("ComputerGroupsMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {mime.ApplicationXML}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := xml.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("ComputerGroupsMock: unmarshal into result: %w", err)
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
