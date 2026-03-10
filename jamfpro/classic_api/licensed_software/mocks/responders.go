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

// LicensedSoftwareMock is a test double implementing client.Client for Classic API licensed software.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Classic API mocks use XML for serialization to match the Classic API wire format.
type LicensedSoftwareMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewLicensedSoftwareMock returns an empty mock ready for response registration.
func NewLicensedSoftwareMock() *LicensedSoftwareMock {
	return &LicensedSoftwareMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *LicensedSoftwareMock) RegisterMocks() {
	m.RegisterListLicensedSoftwareMock()
	m.RegisterGetLicensedSoftwareByIDMock()
	m.RegisterGetLicensedSoftwareByNameMock()
	m.RegisterCreateLicensedSoftwareMock()
	m.RegisterUpdateLicensedSoftwareByIDMock()
	m.RegisterUpdateLicensedSoftwareByNameMock()
	m.RegisterDeleteLicensedSoftwareByIDMock()
	m.RegisterDeleteLicensedSoftwareByNameMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *LicensedSoftwareMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListLicensedSoftwareMock registers GET /JSSResource/licensedsoftware → 200.
func (m *LicensedSoftwareMock) RegisterListLicensedSoftwareMock() {
	m.register("GET", "/JSSResource/licensedsoftware", 200, "validate_list_licensed_software.xml")
}

// RegisterGetLicensedSoftwareByIDMock registers GET /JSSResource/licensedsoftware/id/1 → 200.
func (m *LicensedSoftwareMock) RegisterGetLicensedSoftwareByIDMock() {
	m.register("GET", "/JSSResource/licensedsoftware/id/1", 200, "validate_get_licensed_software.xml")
}

// RegisterGetLicensedSoftwareByNameMock registers GET /JSSResource/licensedsoftware/name/Sample Licensed Software → 200.
func (m *LicensedSoftwareMock) RegisterGetLicensedSoftwareByNameMock() {
	m.register("GET", "/JSSResource/licensedsoftware/name/Sample Licensed Software", 200, "validate_get_licensed_software.xml")
}

// RegisterCreateLicensedSoftwareMock registers POST /JSSResource/licensedsoftware/id/0 → 201.
func (m *LicensedSoftwareMock) RegisterCreateLicensedSoftwareMock() {
	m.register("POST", "/JSSResource/licensedsoftware/id/0", 201, "validate_create_licensed_software.xml")
}

// RegisterUpdateLicensedSoftwareByIDMock registers PUT /JSSResource/licensedsoftware/id/1 → 200.
func (m *LicensedSoftwareMock) RegisterUpdateLicensedSoftwareByIDMock() {
	m.register("PUT", "/JSSResource/licensedsoftware/id/1", 200, "validate_update_licensed_software.xml")
}

// RegisterUpdateLicensedSoftwareByNameMock registers PUT /JSSResource/licensedsoftware/name/Sample Licensed Software → 200.
func (m *LicensedSoftwareMock) RegisterUpdateLicensedSoftwareByNameMock() {
	m.register("PUT", "/JSSResource/licensedsoftware/name/Sample Licensed Software", 200, "validate_update_licensed_software.xml")
}

// RegisterDeleteLicensedSoftwareByIDMock registers DELETE /JSSResource/licensedsoftware/id/1 → 200.
func (m *LicensedSoftwareMock) RegisterDeleteLicensedSoftwareByIDMock() {
	m.register("DELETE", "/JSSResource/licensedsoftware/id/1", 200, "")
}

// RegisterDeleteLicensedSoftwareByNameMock registers DELETE /JSSResource/licensedsoftware/name/Sample Licensed Software → 200.
func (m *LicensedSoftwareMock) RegisterDeleteLicensedSoftwareByNameMock() {
	m.register("DELETE", "/JSSResource/licensedsoftware/name/Sample Licensed Software", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/licensedsoftware/id/999 → 404.
func (m *LicensedSoftwareMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/JSSResource/licensedsoftware/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

// RegisterConflictErrorMock registers POST /JSSResource/licensedsoftware/id/0 → 409.
func (m *LicensedSoftwareMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/JSSResource/licensedsoftware/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): Licensed software with that name already exists")
}

// ---- client.Client implementation ----

func (m *LicensedSoftwareMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	_ = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *LicensedSoftwareMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *LicensedSoftwareMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *LicensedSoftwareMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *LicensedSoftwareMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *LicensedSoftwareMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *LicensedSoftwareMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *LicensedSoftwareMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *LicensedSoftwareMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *LicensedSoftwareMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	_ = rsqlQuery
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *LicensedSoftwareMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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

func (m *LicensedSoftwareMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *LicensedSoftwareMock) InvalidateToken() error                { return nil }
func (m *LicensedSoftwareMock) KeepAliveToken() error                 { return nil }
func (m *LicensedSoftwareMock) GetLogger() *zap.Logger                { return m.logger }

// ---- Internal helpers ----

// registerError stores an error response with externalized XML body.
func (m *LicensedSoftwareMock) registerError(method, path string, statusCode int, fixture, errMsg string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("LicensedSoftwareMock: failed to load error fixture %q: %v", fixture, err))
	}
	m.responses[method+":"+path] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
		errMsg:     errMsg,
	}
}

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *LicensedSoftwareMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("LicensedSoftwareMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *LicensedSoftwareMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		headers := http.Header{"Content-Type": {constants.ApplicationXML}}
		return shared.NewMockResponse(http.StatusNotFound, headers, []byte(`<error>no mock registered</error>`)), fmt.Errorf("LicensedSoftwareMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {constants.ApplicationXML}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := xml.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("LicensedSoftwareMock: unmarshal into result: %w", err)
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
