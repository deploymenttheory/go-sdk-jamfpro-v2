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

// VPPAssignmentsMock is a test double implementing interfaces.HTTPClient for Classic API VPP assignments.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Classic API mocks use XML for serialization to match the Classic API wire format.
type VPPAssignmentsMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewVPPAssignmentsMock returns an empty mock ready for response registration.
func NewVPPAssignmentsMock() *VPPAssignmentsMock {
	return &VPPAssignmentsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *VPPAssignmentsMock) RegisterMocks() {
	m.RegisterListVPPAssignmentsMock()
	m.RegisterGetVPPAssignmentByIDMock()
	m.RegisterCreateVPPAssignmentMock()
	m.RegisterUpdateVPPAssignmentByIDMock()
	m.RegisterDeleteVPPAssignmentByIDMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *VPPAssignmentsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListVPPAssignmentsMock registers GET /JSSResource/vppassignments → 200.
func (m *VPPAssignmentsMock) RegisterListVPPAssignmentsMock() {
	m.register("GET", "/JSSResource/vppassignments", 200, "validate_list_vpp_assignments.xml")
}

// RegisterGetVPPAssignmentByIDMock registers GET /JSSResource/vppassignments/id/1 → 200.
func (m *VPPAssignmentsMock) RegisterGetVPPAssignmentByIDMock() {
	m.register("GET", "/JSSResource/vppassignments/id/1", 200, "validate_get_vpp_assignment.xml")
}

// RegisterCreateVPPAssignmentMock registers POST /JSSResource/vppassignments/id/0 → 201.
func (m *VPPAssignmentsMock) RegisterCreateVPPAssignmentMock() {
	m.register("POST", "/JSSResource/vppassignments/id/0", 201, "validate_create_vpp_assignment.xml")
}

// RegisterUpdateVPPAssignmentByIDMock registers PUT /JSSResource/vppassignments/id/1 → 200.
func (m *VPPAssignmentsMock) RegisterUpdateVPPAssignmentByIDMock() {
	m.register("PUT", "/JSSResource/vppassignments/id/1", 200, "validate_update_vpp_assignment.xml")
}

// RegisterDeleteVPPAssignmentByIDMock registers DELETE /JSSResource/vppassignments/id/1 → 200.
func (m *VPPAssignmentsMock) RegisterDeleteVPPAssignmentByIDMock() {
	m.register("DELETE", "/JSSResource/vppassignments/id/1", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/vppassignments/id/999 → 404.
func (m *VPPAssignmentsMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/JSSResource/vppassignments/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

// RegisterConflictErrorMock registers POST /JSSResource/vppassignments/id/0 → 409.
func (m *VPPAssignmentsMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/JSSResource/vppassignments/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A VPP assignment with that name already exists")
}

// ---- interfaces.HTTPClient implementation ----

func (m *VPPAssignmentsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}

func (m *VPPAssignmentsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *VPPAssignmentsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *VPPAssignmentsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *VPPAssignmentsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *VPPAssignmentsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *VPPAssignmentsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *VPPAssignmentsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *VPPAssignmentsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *VPPAssignmentsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *VPPAssignmentsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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

func (m *VPPAssignmentsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *VPPAssignmentsMock) InvalidateToken() error                     { return nil }
func (m *VPPAssignmentsMock) KeepAliveToken() error                      { return nil }
func (m *VPPAssignmentsMock) GetLogger() *zap.Logger                      { return m.logger }

// ---- Internal helpers ----

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *VPPAssignmentsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("VPPAssignmentsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// registerError stores an error response with externalized XML body.
func (m *VPPAssignmentsMock) registerError(method, path string, statusCode int, fixture, errMsg string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("VPPAssignmentsMock: failed to load error fixture %q: %v", fixture, err))
	}
	m.responses[method+":"+path] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
		errMsg:     errMsg,
	}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *VPPAssignmentsMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/xml"}},
			Body:       []byte(`<error>no mock registered</error>`),
		}, fmt.Errorf("VPPAssignmentsMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("VPPAssignmentsMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

// loadMockResponse reads an XML fixture file from the mocks/ directory
// (same directory as this responders.go file).
func loadMockResponse(filename string) ([]byte, error) {
	_, currentFile, _, _ := runtime.Caller(0)
	dir := filepath.Dir(currentFile)
	data, err := os.ReadFile(filepath.Join(dir, filename))
	if err != nil {
		return nil, fmt.Errorf("read fixture %s: %w", filename, err)
	}
	return data, nil
}
