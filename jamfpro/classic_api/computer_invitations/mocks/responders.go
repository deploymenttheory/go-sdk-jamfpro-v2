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

// ComputerInvitationsMock is a test double implementing transport.HTTPClient for Classic API computer invitations.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Classic API mocks use XML for serialization to match the Classic API wire format.
type ComputerInvitationsMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewComputerInvitationsMock returns an empty mock ready for response registration.
func NewComputerInvitationsMock() *ComputerInvitationsMock {
	return &ComputerInvitationsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *ComputerInvitationsMock) RegisterMocks() {
	m.RegisterListComputerInvitationsMock()
	m.RegisterGetComputerInvitationByIDMock()
	m.RegisterGetComputerInvitationByInvitationIDMock()
	m.RegisterCreateComputerInvitationMock()
	m.RegisterDeleteComputerInvitationByIDMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *ComputerInvitationsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
}

// ---- Success responders ----

// RegisterListComputerInvitationsMock registers GET /JSSResource/computerinvitations → 200.
func (m *ComputerInvitationsMock) RegisterListComputerInvitationsMock() {
	m.register("GET", "/JSSResource/computerinvitations", 200, "validate_list_computer_invitations.xml")
}

// RegisterGetComputerInvitationByIDMock registers GET /JSSResource/computerinvitations/id/1 → 200.
func (m *ComputerInvitationsMock) RegisterGetComputerInvitationByIDMock() {
	m.register("GET", "/JSSResource/computerinvitations/id/1", 200, "validate_get_computer_invitation.xml")
}

// RegisterGetComputerInvitationByInvitationIDMock registers GET /JSSResource/computerinvitations/invitation/1234567890 → 200.
func (m *ComputerInvitationsMock) RegisterGetComputerInvitationByInvitationIDMock() {
	m.register("GET", "/JSSResource/computerinvitations/invitation/1234567890", 200, "validate_get_computer_invitation.xml")
}

// RegisterCreateComputerInvitationMock registers POST /JSSResource/computerinvitations/id/0 → 201.
func (m *ComputerInvitationsMock) RegisterCreateComputerInvitationMock() {
	m.register("POST", "/JSSResource/computerinvitations/id/0", 201, "validate_create_computer_invitation.xml")
}

// RegisterDeleteComputerInvitationByIDMock registers DELETE /JSSResource/computerinvitations/id/1 → 200.
func (m *ComputerInvitationsMock) RegisterDeleteComputerInvitationByIDMock() {
	m.register("DELETE", "/JSSResource/computerinvitations/id/1", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/computerinvitations/id/999 → 404.
func (m *ComputerInvitationsMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/JSSResource/computerinvitations/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

// ---- transport.HTTPClient implementation ----

func (m *ComputerInvitationsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *ComputerInvitationsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputerInvitationsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputerInvitationsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputerInvitationsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ transport.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputerInvitationsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *ComputerInvitationsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *ComputerInvitationsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *ComputerInvitationsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *ComputerInvitationsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *ComputerInvitationsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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

func (m *ComputerInvitationsMock) RSQLBuilder() transport.RSQLFilterBuilder { return nil }
func (m *ComputerInvitationsMock) InvalidateToken() error                    { return nil }
func (m *ComputerInvitationsMock) KeepAliveToken() error                     { return nil }
func (m *ComputerInvitationsMock) GetLogger() *zap.Logger                    { return m.logger }

// ---- Internal helpers ----

// registerError stores an error response with externalized XML body.
func (m *ComputerInvitationsMock) registerError(method, path string, statusCode int, fixture, errMsg string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("ComputerInvitationsMock: failed to load error fixture %q: %v", fixture, err))
	}
	m.responses[method+":"+path] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
		errMsg:     errMsg,
	}
}

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *ComputerInvitationsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("ComputerInvitationsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *ComputerInvitationsMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		headers := http.Header{"Content-Type": {mime.ApplicationXML}}
		return shared.NewMockResponse(http.StatusNotFound, headers, []byte(`<error>no mock registered</error>`)), fmt.Errorf("ComputerInvitationsMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {mime.ApplicationXML}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := xml.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("ComputerInvitationsMock: unmarshal into result: %w", err)
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
