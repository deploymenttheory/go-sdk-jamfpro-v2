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
	errMsg     string
}

// ComputersMock is a test double implementing interfaces.HTTPClient for Classic API computers.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Classic API mocks use XML for serialization to match the Classic API wire format.
type ComputersMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewComputersMock returns an empty mock ready for response registration.
func NewComputersMock() *ComputersMock {
	return &ComputersMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *ComputersMock) RegisterMocks() {
	m.RegisterListComputersMock()
	m.RegisterGetComputerByIDMock()
	m.RegisterGetComputerByNameMock()
	m.RegisterCreateComputerMock()
	m.RegisterUpdateComputerByIDMock()
	m.RegisterUpdateComputerByNameMock()
	m.RegisterDeleteComputerByIDMock()
	m.RegisterDeleteComputerByNameMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *ComputersMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
}

// ---- Success responders ----

// RegisterListComputersMock registers GET /JSSResource/computers → 200.
func (m *ComputersMock) RegisterListComputersMock() {
	m.register("GET", "/JSSResource/computers", 200, "validate_list_computers.xml")
}

// RegisterGetComputerByIDMock registers GET /JSSResource/computers/id/1 → 200.
func (m *ComputersMock) RegisterGetComputerByIDMock() {
	m.register("GET", "/JSSResource/computers/id/1", 200, "validate_get_computer.xml")
}

// RegisterGetComputerByNameMock registers GET /JSSResource/computers/name/MacBook-Pro-01 → 200.
func (m *ComputersMock) RegisterGetComputerByNameMock() {
	m.register("GET", "/JSSResource/computers/name/MacBook-Pro-01", 200, "validate_get_computer.xml")
}

// RegisterCreateComputerMock registers POST /JSSResource/computers → 201.
func (m *ComputersMock) RegisterCreateComputerMock() {
	m.register("POST", "/JSSResource/computers", 201, "validate_create_computer.xml")
}

// RegisterUpdateComputerByIDMock registers PUT /JSSResource/computers/id/1 → 200.
func (m *ComputersMock) RegisterUpdateComputerByIDMock() {
	m.register("PUT", "/JSSResource/computers/id/1", 200, "validate_update_computer.xml")
}

// RegisterUpdateComputerByNameMock registers PUT /JSSResource/computers/name/MacBook-Pro-01 → 200.
func (m *ComputersMock) RegisterUpdateComputerByNameMock() {
	m.register("PUT", "/JSSResource/computers/name/MacBook-Pro-01", 200, "validate_update_computer.xml")
}

// RegisterDeleteComputerByIDMock registers DELETE /JSSResource/computers/id/1 → 200.
func (m *ComputersMock) RegisterDeleteComputerByIDMock() {
	m.register("DELETE", "/JSSResource/computers/id/1", 200, "")
}

// RegisterDeleteComputerByNameMock registers DELETE /JSSResource/computers/name/MacBook-Pro-01 → 200.
func (m *ComputersMock) RegisterDeleteComputerByNameMock() {
	m.register("DELETE", "/JSSResource/computers/name/MacBook-Pro-01", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/computers/id/999 → 404.
func (m *ComputersMock) RegisterNotFoundErrorMock() {
	body := []byte("<br>An error has occurred.<br>Resource not found<br><br>")
	m.responses["GET:/JSSResource/computers/id/999"] = registeredResponse{
		statusCode: 404,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (404): Resource not found",
	}
}

// ---- interfaces.HTTPClient implementation ----

func (m *ComputersMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *ComputersMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputersMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputersMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputersMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputersMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *ComputersMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *ComputersMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *ComputersMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *ComputersMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *ComputersMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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

func (m *ComputersMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *ComputersMock) InvalidateToken() error                     { return nil }
func (m *ComputersMock) KeepAliveToken() error                      { return nil }
func (m *ComputersMock) GetLogger() *zap.Logger                     { return m.logger }

// ---- Internal helpers ----

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *ComputersMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("ComputersMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *ComputersMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/xml"}},
			Body:       []byte(`<error>no mock registered</error>`),
		}, fmt.Errorf("ComputersMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("ComputersMock: unmarshal into result: %w", err)
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
