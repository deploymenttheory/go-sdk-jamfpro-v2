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

// MacOSConfigurationProfilesMock is a test double implementing interfaces.HTTPClient for Classic API macOS configuration profiles.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Classic API mocks use XML for serialization to match the Classic API wire format.
type MacOSConfigurationProfilesMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewMacOSConfigurationProfilesMock returns an empty mock ready for response registration.
func NewMacOSConfigurationProfilesMock() *MacOSConfigurationProfilesMock {
	return &MacOSConfigurationProfilesMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *MacOSConfigurationProfilesMock) RegisterMocks() {
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
func (m *MacOSConfigurationProfilesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListMock registers GET /JSSResource/osxconfigurationprofiles → 200.
func (m *MacOSConfigurationProfilesMock) RegisterListMock() {
	m.register("GET", "/JSSResource/osxconfigurationprofiles", 200, "validate_list_osx_configuration_profiles.xml")
}

// RegisterGetByIDMock registers GET /JSSResource/osxconfigurationprofiles/id/1 → 200.
func (m *MacOSConfigurationProfilesMock) RegisterGetByIDMock() {
	m.register("GET", "/JSSResource/osxconfigurationprofiles/id/1", 200, "validate_get_osx_configuration_profile.xml")
}

// RegisterGetByNameMock registers GET /JSSResource/osxconfigurationprofiles/name/Wi-Fi Profile → 200.
func (m *MacOSConfigurationProfilesMock) RegisterGetByNameMock() {
	m.register("GET", "/JSSResource/osxconfigurationprofiles/name/Wi-Fi Profile", 200, "validate_get_osx_configuration_profile.xml")
}

// RegisterCreateMock registers POST /JSSResource/osxconfigurationprofiles/id/0 → 201.
func (m *MacOSConfigurationProfilesMock) RegisterCreateMock() {
	m.register("POST", "/JSSResource/osxconfigurationprofiles/id/0", 201, "validate_create_osx_configuration_profile.xml")
}

// RegisterUpdateByIDMock registers PUT /JSSResource/osxconfigurationprofiles/id/1 → 200.
func (m *MacOSConfigurationProfilesMock) RegisterUpdateByIDMock() {
	m.register("PUT", "/JSSResource/osxconfigurationprofiles/id/1", 200, "validate_update_osx_configuration_profile.xml")
}

// RegisterUpdateByNameMock registers PUT /JSSResource/osxconfigurationprofiles/name/Wi-Fi Profile → 200.
func (m *MacOSConfigurationProfilesMock) RegisterUpdateByNameMock() {
	m.register("PUT", "/JSSResource/osxconfigurationprofiles/name/Wi-Fi Profile", 200, "validate_update_osx_configuration_profile.xml")
}

// RegisterDeleteByIDMock registers DELETE /JSSResource/osxconfigurationprofiles/id/1 → 200.
func (m *MacOSConfigurationProfilesMock) RegisterDeleteByIDMock() {
	m.register("DELETE", "/JSSResource/osxconfigurationprofiles/id/1", 200, "")
}

// RegisterDeleteByNameMock registers DELETE /JSSResource/osxconfigurationprofiles/name/Wi-Fi Profile → 200.
func (m *MacOSConfigurationProfilesMock) RegisterDeleteByNameMock() {
	m.register("DELETE", "/JSSResource/osxconfigurationprofiles/name/Wi-Fi Profile", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/osxconfigurationprofiles/id/999 → 404.
func (m *MacOSConfigurationProfilesMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/JSSResource/osxconfigurationprofiles/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

// RegisterConflictErrorMock registers POST /JSSResource/osxconfigurationprofiles/id/0 → 409.
func (m *MacOSConfigurationProfilesMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/JSSResource/osxconfigurationprofiles/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A configuration profile with that name already exists")
}

// ---- interfaces.HTTPClient implementation ----

func (m *MacOSConfigurationProfilesMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}

func (m *MacOSConfigurationProfilesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MacOSConfigurationProfilesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MacOSConfigurationProfilesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MacOSConfigurationProfilesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MacOSConfigurationProfilesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *MacOSConfigurationProfilesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *MacOSConfigurationProfilesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MacOSConfigurationProfilesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MacOSConfigurationProfilesMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *MacOSConfigurationProfilesMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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

func (m *MacOSConfigurationProfilesMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *MacOSConfigurationProfilesMock) InvalidateToken() error                    { return nil }
func (m *MacOSConfigurationProfilesMock) KeepAliveToken() error                     { return nil }
func (m *MacOSConfigurationProfilesMock) GetLogger() *zap.Logger                    { return m.logger }

// ---- Internal helpers ----

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *MacOSConfigurationProfilesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("MacOSConfigurationProfilesMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// registerError stores an error response with externalized XML body.
func (m *MacOSConfigurationProfilesMock) registerError(method, path string, statusCode int, fixture, errMsg string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("MacOSConfigurationProfilesMock: failed to load error fixture %q: %v", fixture, err))
	}
	m.responses[method+":"+path] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
		errMsg:     errMsg,
	}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *MacOSConfigurationProfilesMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		headers := http.Header{"Content-Type": {mime.ApplicationXML}}
		return shared.NewMockResponse(http.StatusNotFound, headers, []byte(`<error>no mock registered</error>`)), fmt.Errorf("MacOSConfigurationProfilesMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {mime.ApplicationXML}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := xml.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("MacOSConfigurationProfilesMock: unmarshal into result: %w", err)
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
