package mocks

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"go.uber.org/zap"
	"resty.dev/v3"
)

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	// errMsg causes the mock to return an error (simulating API or transport failures).
	errMsg string
}

// RestrictedSoftwareMock is a test double implementing client.Client for Classic API restricted software.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Unlike Jamf Pro API mocks which use JSON, Classic API mocks use XML for
// serialization to match the Classic API wire format.
type RestrictedSoftwareMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string // captures the rsqlQuery from the most recent Get call
}

// NewRestrictedSoftwareMock returns an empty mock ready for response registration.
func NewRestrictedSoftwareMock() *RestrictedSoftwareMock {
	return &RestrictedSoftwareMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *RestrictedSoftwareMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetByIDMock()
	m.RegisterGetByNameMock()
	m.RegisterCreateMock()
	m.RegisterUpdateByIDMock()
	m.RegisterUpdateByNameMock()
	m.RegisterDeleteByIDMock()
	m.RegisterDeleteByNameMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *RestrictedSoftwareMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListMock registers GET /JSSResource/restrictedsoftware → 200.
func (m *RestrictedSoftwareMock) RegisterListMock() {
	m.register("GET", "/JSSResource/restrictedsoftware", 200, "validate_list_restricted_software.xml")
}

// RegisterGetByIDMock registers GET /JSSResource/restrictedsoftware/id/1 → 200.
func (m *RestrictedSoftwareMock) RegisterGetByIDMock() {
	m.register("GET", "/JSSResource/restrictedsoftware/id/1", 200, "validate_get_restricted_software.xml")
}

// RegisterGetByNameMock registers GET /JSSResource/restrictedsoftware/name/Calculator → 200.
func (m *RestrictedSoftwareMock) RegisterGetByNameMock() {
	m.register("GET", "/JSSResource/restrictedsoftware/name/Calculator", 200, "validate_get_restricted_software.xml")
}

// RegisterCreateMock registers POST /JSSResource/restrictedsoftware/id/0 → 201.
func (m *RestrictedSoftwareMock) RegisterCreateMock() {
	m.register("POST", "/JSSResource/restrictedsoftware/id/0", 201, "validate_create_restricted_software.xml")
}

// RegisterUpdateByIDMock registers PUT /JSSResource/restrictedsoftware/id/1 → 200.
func (m *RestrictedSoftwareMock) RegisterUpdateByIDMock() {
	m.register("PUT", "/JSSResource/restrictedsoftware/id/1", 200, "validate_update_restricted_software.xml")
}

// RegisterUpdateByNameMock registers PUT /JSSResource/restrictedsoftware/name/Calculator → 200.
func (m *RestrictedSoftwareMock) RegisterUpdateByNameMock() {
	m.register("PUT", "/JSSResource/restrictedsoftware/name/Calculator", 200, "validate_update_restricted_software.xml")
}

// RegisterDeleteByIDMock registers DELETE /JSSResource/restrictedsoftware/id/1 → 200.
func (m *RestrictedSoftwareMock) RegisterDeleteByIDMock() {
	m.register("DELETE", "/JSSResource/restrictedsoftware/id/1", 200, "")
}

// RegisterDeleteByNameMock registers DELETE /JSSResource/restrictedsoftware/name/Calculator → 200.
func (m *RestrictedSoftwareMock) RegisterDeleteByNameMock() {
	m.register("DELETE", "/JSSResource/restrictedsoftware/name/Calculator", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/restrictedsoftware/id/999 → 404.
func (m *RestrictedSoftwareMock) RegisterNotFoundErrorMock() {
	body := []byte("<br>An error has occurred.<br>Resource not found<br><br>")
	m.responses["GET:/JSSResource/restrictedsoftware/id/999"] = registeredResponse{
		statusCode: 404,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (404): Resource not found",
	}
}

// RegisterConflictErrorMock registers POST /JSSResource/restrictedsoftware/id/0 → 409
// when the caller wishes to simulate a duplicate-name conflict.
func (m *RestrictedSoftwareMock) RegisterConflictErrorMock() {
	body := []byte("<br>An error has occurred.<br>A restricted software with that name already exists.<br><br>")
	m.responses["POST:/JSSResource/restrictedsoftware/id/0"] = registeredResponse{
		statusCode: 409,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (409): A restricted software with that name already exists",
	}
}

// ---- client.Client implementation ----

func (m *RestrictedSoftwareMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *RestrictedSoftwareMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *RestrictedSoftwareMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *RestrictedSoftwareMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *RestrictedSoftwareMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *RestrictedSoftwareMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *RestrictedSoftwareMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *RestrictedSoftwareMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *RestrictedSoftwareMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *RestrictedSoftwareMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *RestrictedSoftwareMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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

func (m *RestrictedSoftwareMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *RestrictedSoftwareMock) InvalidateToken() error                { return nil }
func (m *RestrictedSoftwareMock) KeepAliveToken() error                 { return nil }
func (m *RestrictedSoftwareMock) GetLogger() *zap.Logger                { return m.logger }

// ---- Internal helpers ----

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *RestrictedSoftwareMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("RestrictedSoftwareMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *RestrictedSoftwareMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		headers := http.Header{"Content-Type": {constants.ApplicationXML}}
		return shared.NewMockResponse(http.StatusNotFound, headers, []byte(`<error>no mock registered</error>`)), fmt.Errorf("RestrictedSoftwareMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {constants.ApplicationXML}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := xml.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("RestrictedSoftwareMock: unmarshal into result: %w", err)
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
