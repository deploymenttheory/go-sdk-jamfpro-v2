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

// MobileDeviceEnrollmentProfilesMock is a test double implementing transport.HTTPClient for Classic API mobile device enrollment profiles.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Classic API mocks use XML for serialization to match the Classic API wire format.
type MobileDeviceEnrollmentProfilesMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewMobileDeviceEnrollmentProfilesMock returns an empty mock ready for response registration.
func NewMobileDeviceEnrollmentProfilesMock() *MobileDeviceEnrollmentProfilesMock {
	return &MobileDeviceEnrollmentProfilesMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *MobileDeviceEnrollmentProfilesMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetByIDMock()
	m.RegisterGetByNameMock()
	m.RegisterGetByInvitationMock()
	m.RegisterGetByIDWithSubsetMock()
	m.RegisterGetByNameWithSubsetMock()
	m.RegisterCreateMock()
	m.RegisterUpdateByIDMock()
	m.RegisterUpdateByNameMock()
	m.RegisterUpdateByInvitationMock()
	m.RegisterDeleteByIDMock()
	m.RegisterDeleteByNameMock()
	m.RegisterDeleteByInvitationMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *MobileDeviceEnrollmentProfilesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
}

// ---- Success responders ----

// RegisterListMock registers GET /JSSResource/mobiledeviceenrollmentprofiles → 200.
func (m *MobileDeviceEnrollmentProfilesMock) RegisterListMock() {
	m.register("GET", "/JSSResource/mobiledeviceenrollmentprofiles", 200, "validate_list_mobile_device_enrollment_profiles.xml")
}

// RegisterGetByIDMock registers GET /JSSResource/mobiledeviceenrollmentprofiles/id/1 → 200.
func (m *MobileDeviceEnrollmentProfilesMock) RegisterGetByIDMock() {
	m.register("GET", "/JSSResource/mobiledeviceenrollmentprofiles/id/1", 200, "validate_get_mobile_device_enrollment_profile.xml")
}

// RegisterGetByNameMock registers GET /JSSResource/mobiledeviceenrollmentprofiles/name/Test Enrollment Profile → 200.
func (m *MobileDeviceEnrollmentProfilesMock) RegisterGetByNameMock() {
	m.register("GET", "/JSSResource/mobiledeviceenrollmentprofiles/name/Test Enrollment Profile", 200, "validate_get_mobile_device_enrollment_profile.xml")
}

// RegisterGetByInvitationMock registers GET /JSSResource/mobiledeviceenrollmentprofiles/invitation/1234567890.123456 → 200.
func (m *MobileDeviceEnrollmentProfilesMock) RegisterGetByInvitationMock() {
	m.register("GET", "/JSSResource/mobiledeviceenrollmentprofiles/invitation/1234567890.123456", 200, "validate_get_mobile_device_enrollment_profile.xml")
}

// RegisterGetByIDWithSubsetMock registers GET /JSSResource/mobiledeviceenrollmentprofiles/id/1/subset/General → 200.
func (m *MobileDeviceEnrollmentProfilesMock) RegisterGetByIDWithSubsetMock() {
	m.register("GET", "/JSSResource/mobiledeviceenrollmentprofiles/id/1/subset/General", 200, "validate_get_mobile_device_enrollment_profile.xml")
}

// RegisterGetByNameWithSubsetMock registers GET /JSSResource/mobiledeviceenrollmentprofiles/name/Test Enrollment Profile/subset/General → 200.
func (m *MobileDeviceEnrollmentProfilesMock) RegisterGetByNameWithSubsetMock() {
	m.register("GET", "/JSSResource/mobiledeviceenrollmentprofiles/name/Test Enrollment Profile/subset/General", 200, "validate_get_mobile_device_enrollment_profile.xml")
}

// RegisterCreateMock registers POST /JSSResource/mobiledeviceenrollmentprofiles/id/0 → 201.
func (m *MobileDeviceEnrollmentProfilesMock) RegisterCreateMock() {
	m.register("POST", "/JSSResource/mobiledeviceenrollmentprofiles/id/0", 201, "validate_create_mobile_device_enrollment_profile.xml")
}

// RegisterUpdateByIDMock registers PUT /JSSResource/mobiledeviceenrollmentprofiles/id/1 → 200.
func (m *MobileDeviceEnrollmentProfilesMock) RegisterUpdateByIDMock() {
	m.register("PUT", "/JSSResource/mobiledeviceenrollmentprofiles/id/1", 200, "validate_update_mobile_device_enrollment_profile.xml")
}

// RegisterUpdateByNameMock registers PUT /JSSResource/mobiledeviceenrollmentprofiles/name/Test Enrollment Profile → 200.
func (m *MobileDeviceEnrollmentProfilesMock) RegisterUpdateByNameMock() {
	m.register("PUT", "/JSSResource/mobiledeviceenrollmentprofiles/name/Test Enrollment Profile", 200, "validate_update_mobile_device_enrollment_profile.xml")
}

// RegisterUpdateByInvitationMock registers PUT /JSSResource/mobiledeviceenrollmentprofiles/invitation/1234567890.123456 → 200.
func (m *MobileDeviceEnrollmentProfilesMock) RegisterUpdateByInvitationMock() {
	m.register("PUT", "/JSSResource/mobiledeviceenrollmentprofiles/invitation/1234567890.123456", 200, "validate_update_mobile_device_enrollment_profile.xml")
}

// RegisterDeleteByIDMock registers DELETE /JSSResource/mobiledeviceenrollmentprofiles/id/1 → 200.
func (m *MobileDeviceEnrollmentProfilesMock) RegisterDeleteByIDMock() {
	m.register("DELETE", "/JSSResource/mobiledeviceenrollmentprofiles/id/1", 200, "")
}

// RegisterDeleteByNameMock registers DELETE /JSSResource/mobiledeviceenrollmentprofiles/name/Test Enrollment Profile → 200.
func (m *MobileDeviceEnrollmentProfilesMock) RegisterDeleteByNameMock() {
	m.register("DELETE", "/JSSResource/mobiledeviceenrollmentprofiles/name/Test Enrollment Profile", 200, "")
}

// RegisterDeleteByInvitationMock registers DELETE /JSSResource/mobiledeviceenrollmentprofiles/invitation/1234567890.123456 → 200.
func (m *MobileDeviceEnrollmentProfilesMock) RegisterDeleteByInvitationMock() {
	m.register("DELETE", "/JSSResource/mobiledeviceenrollmentprofiles/invitation/1234567890.123456", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/mobiledeviceenrollmentprofiles/id/999 → 404.
func (m *MobileDeviceEnrollmentProfilesMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/JSSResource/mobiledeviceenrollmentprofiles/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

// ---- transport.HTTPClient implementation ----

func (m *MobileDeviceEnrollmentProfilesMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}

func (m *MobileDeviceEnrollmentProfilesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceEnrollmentProfilesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceEnrollmentProfilesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceEnrollmentProfilesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ transport.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceEnrollmentProfilesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *MobileDeviceEnrollmentProfilesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *MobileDeviceEnrollmentProfilesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MobileDeviceEnrollmentProfilesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MobileDeviceEnrollmentProfilesMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *MobileDeviceEnrollmentProfilesMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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

func (m *MobileDeviceEnrollmentProfilesMock) RSQLBuilder() transport.RSQLFilterBuilder { return nil }
func (m *MobileDeviceEnrollmentProfilesMock) InvalidateToken() error                    { return nil }
func (m *MobileDeviceEnrollmentProfilesMock) KeepAliveToken() error                     { return nil }
func (m *MobileDeviceEnrollmentProfilesMock) GetLogger() *zap.Logger                    { return m.logger }

// ---- Internal helpers ----

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *MobileDeviceEnrollmentProfilesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("MobileDeviceEnrollmentProfilesMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// registerError stores an error response with externalized XML body.
func (m *MobileDeviceEnrollmentProfilesMock) registerError(method, path string, statusCode int, fixture, errMsg string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("MobileDeviceEnrollmentProfilesMock: failed to load error fixture %q: %v", fixture, err))
	}
	m.responses[method+":"+path] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
		errMsg:     errMsg,
	}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *MobileDeviceEnrollmentProfilesMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		headers := http.Header{"Content-Type": {constants.ApplicationXML}}
		return shared.NewMockResponse(http.StatusNotFound, headers, []byte(`<error>no mock registered</error>`)), fmt.Errorf("MobileDeviceEnrollmentProfilesMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {constants.ApplicationXML}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := xml.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("MobileDeviceEnrollmentProfilesMock: unmarshal into result: %w", err)
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
