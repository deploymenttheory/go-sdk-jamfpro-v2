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

// MobileDeviceConfigurationProfilesMock is a test double implementing interfaces.HTTPClient for Classic API mobile device configuration profiles.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Classic API mocks use XML for serialization to match the Classic API wire format.
type MobileDeviceConfigurationProfilesMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewMobileDeviceConfigurationProfilesMock returns an empty mock ready for response registration.
func NewMobileDeviceConfigurationProfilesMock() *MobileDeviceConfigurationProfilesMock {
	return &MobileDeviceConfigurationProfilesMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *MobileDeviceConfigurationProfilesMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetByIDMock()
	m.RegisterGetByNameMock()
	m.RegisterGetByIDWithSubsetMock()
	m.RegisterGetByNameWithSubsetMock()
	m.RegisterCreateMock()
	m.RegisterUpdateByIDMock()
	m.RegisterUpdateByNameMock()
	m.RegisterDeleteByIDMock()
	m.RegisterDeleteByNameMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *MobileDeviceConfigurationProfilesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListMock registers GET /JSSResource/mobiledeviceconfigurationprofiles → 200.
func (m *MobileDeviceConfigurationProfilesMock) RegisterListMock() {
	m.register("GET", "/JSSResource/mobiledeviceconfigurationprofiles", 200, "validate_list_mobile_device_configuration_profiles.xml")
}

// RegisterGetByIDMock registers GET /JSSResource/mobiledeviceconfigurationprofiles/id/1 → 200.
func (m *MobileDeviceConfigurationProfilesMock) RegisterGetByIDMock() {
	m.register("GET", "/JSSResource/mobiledeviceconfigurationprofiles/id/1", 200, "validate_get_mobile_device_configuration_profile.xml")
}

// RegisterGetByNameMock registers GET /JSSResource/mobiledeviceconfigurationprofiles/name/Wi-Fi Profile → 200.
func (m *MobileDeviceConfigurationProfilesMock) RegisterGetByNameMock() {
	m.register("GET", "/JSSResource/mobiledeviceconfigurationprofiles/name/Wi-Fi Profile", 200, "validate_get_mobile_device_configuration_profile.xml")
}

// RegisterGetByIDWithSubsetMock registers GET /JSSResource/mobiledeviceconfigurationprofiles/id/1/subset/General → 200.
func (m *MobileDeviceConfigurationProfilesMock) RegisterGetByIDWithSubsetMock() {
	m.register("GET", "/JSSResource/mobiledeviceconfigurationprofiles/id/1/subset/General", 200, "validate_get_mobile_device_configuration_profile.xml")
}

// RegisterGetByNameWithSubsetMock registers GET /JSSResource/mobiledeviceconfigurationprofiles/name/Wi-Fi Profile/subset/General → 200.
func (m *MobileDeviceConfigurationProfilesMock) RegisterGetByNameWithSubsetMock() {
	m.register("GET", "/JSSResource/mobiledeviceconfigurationprofiles/name/Wi-Fi Profile/subset/General", 200, "validate_get_mobile_device_configuration_profile.xml")
}

// RegisterCreateMock registers POST /JSSResource/mobiledeviceconfigurationprofiles/id/0 → 201.
func (m *MobileDeviceConfigurationProfilesMock) RegisterCreateMock() {
	m.register("POST", "/JSSResource/mobiledeviceconfigurationprofiles/id/0", 201, "validate_create_mobile_device_configuration_profile.xml")
}

// RegisterUpdateByIDMock registers PUT /JSSResource/mobiledeviceconfigurationprofiles/id/1 → 200.
func (m *MobileDeviceConfigurationProfilesMock) RegisterUpdateByIDMock() {
	m.register("PUT", "/JSSResource/mobiledeviceconfigurationprofiles/id/1", 200, "validate_update_mobile_device_configuration_profile.xml")
}

// RegisterUpdateByNameMock registers PUT /JSSResource/mobiledeviceconfigurationprofiles/name/Wi-Fi Profile → 200.
func (m *MobileDeviceConfigurationProfilesMock) RegisterUpdateByNameMock() {
	m.register("PUT", "/JSSResource/mobiledeviceconfigurationprofiles/name/Wi-Fi Profile", 200, "validate_update_mobile_device_configuration_profile.xml")
}

// RegisterDeleteByIDMock registers DELETE /JSSResource/mobiledeviceconfigurationprofiles/id/1 → 200.
func (m *MobileDeviceConfigurationProfilesMock) RegisterDeleteByIDMock() {
	m.register("DELETE", "/JSSResource/mobiledeviceconfigurationprofiles/id/1", 200, "")
}

// RegisterDeleteByNameMock registers DELETE /JSSResource/mobiledeviceconfigurationprofiles/name/Wi-Fi Profile → 200.
func (m *MobileDeviceConfigurationProfilesMock) RegisterDeleteByNameMock() {
	m.register("DELETE", "/JSSResource/mobiledeviceconfigurationprofiles/name/Wi-Fi Profile", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/mobiledeviceconfigurationprofiles/id/999 → 404.
func (m *MobileDeviceConfigurationProfilesMock) RegisterNotFoundErrorMock() {
	body := []byte("<br>An error has occurred.<br>Resource not found<br><br>")
	m.responses["GET:/JSSResource/mobiledeviceconfigurationprofiles/id/999"] = registeredResponse{
		statusCode: 404,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (404): Resource not found",
	}
}

// RegisterConflictErrorMock registers POST /JSSResource/mobiledeviceconfigurationprofiles/id/0 → 409.
func (m *MobileDeviceConfigurationProfilesMock) RegisterConflictErrorMock() {
	body := []byte("<br>An error has occurred.<br>A mobile device configuration profile with that name already exists.<br><br>")
	m.responses["POST:/JSSResource/mobiledeviceconfigurationprofiles/id/0"] = registeredResponse{
		statusCode: 409,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (409): A mobile device configuration profile with that name already exists",
	}
}

// ---- interfaces.HTTPClient implementation ----

func (m *MobileDeviceConfigurationProfilesMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *MobileDeviceConfigurationProfilesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceConfigurationProfilesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceConfigurationProfilesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceConfigurationProfilesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceConfigurationProfilesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *MobileDeviceConfigurationProfilesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *MobileDeviceConfigurationProfilesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MobileDeviceConfigurationProfilesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MobileDeviceConfigurationProfilesMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *MobileDeviceConfigurationProfilesMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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

func (m *MobileDeviceConfigurationProfilesMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *MobileDeviceConfigurationProfilesMock) InvalidateToken() error                     { return nil }
func (m *MobileDeviceConfigurationProfilesMock) KeepAliveToken() error                      { return nil }
func (m *MobileDeviceConfigurationProfilesMock) GetLogger() *zap.Logger                     { return m.logger }

// ---- Internal helpers ----

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *MobileDeviceConfigurationProfilesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("MobileDeviceConfigurationProfilesMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *MobileDeviceConfigurationProfilesMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/xml"}},
			Body:       []byte(`<error>no mock registered</error>`),
		}, fmt.Errorf("MobileDeviceConfigurationProfilesMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("MobileDeviceConfigurationProfilesMock: unmarshal into result: %w", err)
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
