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

// MobileDeviceProvisioningProfilesMock is a test double implementing transport.HTTPClient for Classic API mobile device provisioning profiles.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Classic API mocks use XML for serialization to match the Classic API wire format.
type MobileDeviceProvisioningProfilesMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewMobileDeviceProvisioningProfilesMock returns an empty mock ready for response registration.
func NewMobileDeviceProvisioningProfilesMock() *MobileDeviceProvisioningProfilesMock {
	return &MobileDeviceProvisioningProfilesMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *MobileDeviceProvisioningProfilesMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetByIDMock()
	m.RegisterGetByNameMock()
	m.RegisterGetByUUIDMock()
	m.RegisterCreateByIDMock()
	m.RegisterCreateByNameMock()
	m.RegisterCreateByUUIDMock()
	m.RegisterUpdateByIDMock()
	m.RegisterUpdateByNameMock()
	m.RegisterUpdateByUUIDMock()
	m.RegisterDeleteByIDMock()
	m.RegisterDeleteByNameMock()
	m.RegisterDeleteByUUIDMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *MobileDeviceProvisioningProfilesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListMock registers GET /JSSResource/mobiledeviceprovisioningprofiles → 200.
func (m *MobileDeviceProvisioningProfilesMock) RegisterListMock() {
	m.register("GET", "/JSSResource/mobiledeviceprovisioningprofiles", 200, "validate_list_mobile_device_provisioning_profiles.xml")
}

// RegisterGetByIDMock registers GET /JSSResource/mobiledeviceprovisioningprofiles/id/1 → 200.
func (m *MobileDeviceProvisioningProfilesMock) RegisterGetByIDMock() {
	m.register("GET", "/JSSResource/mobiledeviceprovisioningprofiles/id/1", 200, "validate_get_mobile_device_provisioning_profile.xml")
}

// RegisterGetByNameMock registers GET /JSSResource/mobiledeviceprovisioningprofiles/name/Test Provisioning Profile → 200.
func (m *MobileDeviceProvisioningProfilesMock) RegisterGetByNameMock() {
	m.register("GET", "/JSSResource/mobiledeviceprovisioningprofiles/name/Test Provisioning Profile", 200, "validate_get_mobile_device_provisioning_profile.xml")
}

// RegisterGetByUUIDMock registers GET /JSSResource/mobiledeviceprovisioningprofiles/uuid/550e8400-e29b-41d4-a716-446655440000 → 200.
func (m *MobileDeviceProvisioningProfilesMock) RegisterGetByUUIDMock() {
	m.register("GET", "/JSSResource/mobiledeviceprovisioningprofiles/uuid/550e8400-e29b-41d4-a716-446655440000", 200, "validate_get_mobile_device_provisioning_profile.xml")
}

// RegisterCreateByIDMock registers POST /JSSResource/mobiledeviceprovisioningprofiles/id/0 → 201.
func (m *MobileDeviceProvisioningProfilesMock) RegisterCreateByIDMock() {
	m.register("POST", "/JSSResource/mobiledeviceprovisioningprofiles/id/0", 201, "validate_create_mobile_device_provisioning_profile.xml")
}

// RegisterCreateByNameMock registers POST /JSSResource/mobiledeviceprovisioningprofiles/name/New Profile → 201.
func (m *MobileDeviceProvisioningProfilesMock) RegisterCreateByNameMock() {
	m.register("POST", "/JSSResource/mobiledeviceprovisioningprofiles/name/New Profile", 201, "validate_create_mobile_device_provisioning_profile.xml")
}

// RegisterCreateByUUIDMock registers POST /JSSResource/mobiledeviceprovisioningprofiles/uuid/550e8400-e29b-41d4-a716-446655440001 → 201.
func (m *MobileDeviceProvisioningProfilesMock) RegisterCreateByUUIDMock() {
	m.register("POST", "/JSSResource/mobiledeviceprovisioningprofiles/uuid/550e8400-e29b-41d4-a716-446655440001", 201, "validate_create_mobile_device_provisioning_profile.xml")
}

// RegisterUpdateByIDMock registers PUT /JSSResource/mobiledeviceprovisioningprofiles/id/1 → 200.
func (m *MobileDeviceProvisioningProfilesMock) RegisterUpdateByIDMock() {
	m.register("PUT", "/JSSResource/mobiledeviceprovisioningprofiles/id/1", 200, "validate_update_mobile_device_provisioning_profile.xml")
}

// RegisterUpdateByNameMock registers PUT /JSSResource/mobiledeviceprovisioningprofiles/name/Test Provisioning Profile → 200.
func (m *MobileDeviceProvisioningProfilesMock) RegisterUpdateByNameMock() {
	m.register("PUT", "/JSSResource/mobiledeviceprovisioningprofiles/name/Test Provisioning Profile", 200, "validate_update_mobile_device_provisioning_profile.xml")
}

// RegisterUpdateByUUIDMock registers PUT /JSSResource/mobiledeviceprovisioningprofiles/uuid/550e8400-e29b-41d4-a716-446655440000 → 200.
func (m *MobileDeviceProvisioningProfilesMock) RegisterUpdateByUUIDMock() {
	m.register("PUT", "/JSSResource/mobiledeviceprovisioningprofiles/uuid/550e8400-e29b-41d4-a716-446655440000", 200, "validate_update_mobile_device_provisioning_profile.xml")
}

// RegisterDeleteByIDMock registers DELETE /JSSResource/mobiledeviceprovisioningprofiles/id/1 → 200.
func (m *MobileDeviceProvisioningProfilesMock) RegisterDeleteByIDMock() {
	m.register("DELETE", "/JSSResource/mobiledeviceprovisioningprofiles/id/1", 200, "")
}

// RegisterDeleteByNameMock registers DELETE /JSSResource/mobiledeviceprovisioningprofiles/name/Test Provisioning Profile → 200.
func (m *MobileDeviceProvisioningProfilesMock) RegisterDeleteByNameMock() {
	m.register("DELETE", "/JSSResource/mobiledeviceprovisioningprofiles/name/Test Provisioning Profile", 200, "")
}

// RegisterDeleteByUUIDMock registers DELETE /JSSResource/mobiledeviceprovisioningprofiles/uuid/550e8400-e29b-41d4-a716-446655440000 → 200.
func (m *MobileDeviceProvisioningProfilesMock) RegisterDeleteByUUIDMock() {
	m.register("DELETE", "/JSSResource/mobiledeviceprovisioningprofiles/uuid/550e8400-e29b-41d4-a716-446655440000", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/mobiledeviceprovisioningprofiles/id/999 → 404.
func (m *MobileDeviceProvisioningProfilesMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/JSSResource/mobiledeviceprovisioningprofiles/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

// RegisterConflictErrorMock registers POST /JSSResource/mobiledeviceprovisioningprofiles/id/0 → 409.
func (m *MobileDeviceProvisioningProfilesMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/JSSResource/mobiledeviceprovisioningprofiles/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A mobile device provisioning profile with that name already exists")
}

// ---- transport.HTTPClient implementation ----

func (m *MobileDeviceProvisioningProfilesMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *MobileDeviceProvisioningProfilesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceProvisioningProfilesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceProvisioningProfilesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceProvisioningProfilesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ transport.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceProvisioningProfilesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *MobileDeviceProvisioningProfilesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *MobileDeviceProvisioningProfilesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MobileDeviceProvisioningProfilesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MobileDeviceProvisioningProfilesMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *MobileDeviceProvisioningProfilesMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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

func (m *MobileDeviceProvisioningProfilesMock) RSQLBuilder() transport.RSQLFilterBuilder { return nil }
func (m *MobileDeviceProvisioningProfilesMock) InvalidateToken() error                    { return nil }
func (m *MobileDeviceProvisioningProfilesMock) KeepAliveToken() error                     { return nil }
func (m *MobileDeviceProvisioningProfilesMock) GetLogger() *zap.Logger                    { return m.logger }

// ---- Internal helpers ----

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *MobileDeviceProvisioningProfilesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("MobileDeviceProvisioningProfilesMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// registerError stores an error response with externalized XML body.
func (m *MobileDeviceProvisioningProfilesMock) registerError(method, path string, statusCode int, fixture, errMsg string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("MobileDeviceProvisioningProfilesMock: failed to load error fixture %q: %v", fixture, err))
	}
	m.responses[method+":"+path] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
		errMsg:     errMsg,
	}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *MobileDeviceProvisioningProfilesMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		headers := http.Header{"Content-Type": {constants.ApplicationXML}}
		return shared.NewMockResponse(http.StatusNotFound, headers, []byte(`<error>no mock registered</error>`)), fmt.Errorf("MobileDeviceProvisioningProfilesMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {constants.ApplicationXML}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := xml.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("MobileDeviceProvisioningProfilesMock: unmarshal into result: %w", err)
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
