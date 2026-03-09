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

// MobileDeviceGroupsMock is a test double implementing transport.HTTPClient for Classic API mobile device groups.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Classic API mocks use XML for serialization to match the Classic API wire format.
type MobileDeviceGroupsMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewMobileDeviceGroupsMock returns an empty mock ready for response registration.
func NewMobileDeviceGroupsMock() *MobileDeviceGroupsMock {
	return &MobileDeviceGroupsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *MobileDeviceGroupsMock) RegisterMocks() {
	m.RegisterListMobileDeviceGroupsMock()
	m.RegisterGetMobileDeviceGroupByIDMock()
	m.RegisterGetMobileDeviceGroupByNameMock()
	m.RegisterCreateMobileDeviceGroupMock()
	m.RegisterUpdateMobileDeviceGroupByIDMock()
	m.RegisterUpdateMobileDeviceGroupByNameMock()
	m.RegisterDeleteMobileDeviceGroupByIDMock()
	m.RegisterDeleteMobileDeviceGroupByNameMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *MobileDeviceGroupsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListMobileDeviceGroupsMock registers GET /JSSResource/mobiledevicegroups → 200.
func (m *MobileDeviceGroupsMock) RegisterListMobileDeviceGroupsMock() {
	m.register("GET", "/JSSResource/mobiledevicegroups", 200, "validate_list_mobile_device_groups.xml")
}

// RegisterGetMobileDeviceGroupByIDMock registers GET /JSSResource/mobiledevicegroups/id/1 → 200.
func (m *MobileDeviceGroupsMock) RegisterGetMobileDeviceGroupByIDMock() {
	m.register("GET", "/JSSResource/mobiledevicegroups/id/1", 200, "validate_get_mobile_device_group.xml")
}

// RegisterGetMobileDeviceGroupByNameMock registers GET /JSSResource/mobiledevicegroups/name/All Mobile Devices → 200.
func (m *MobileDeviceGroupsMock) RegisterGetMobileDeviceGroupByNameMock() {
	m.register("GET", "/JSSResource/mobiledevicegroups/name/All Mobile Devices", 200, "validate_get_mobile_device_group.xml")
}

// RegisterCreateMobileDeviceGroupMock registers POST /JSSResource/mobiledevicegroups/id/0 → 201.
func (m *MobileDeviceGroupsMock) RegisterCreateMobileDeviceGroupMock() {
	m.register("POST", "/JSSResource/mobiledevicegroups/id/0", 201, "validate_create_mobile_device_group.xml")
}

// RegisterUpdateMobileDeviceGroupByIDMock registers PUT /JSSResource/mobiledevicegroups/id/1 → 200.
func (m *MobileDeviceGroupsMock) RegisterUpdateMobileDeviceGroupByIDMock() {
	m.register("PUT", "/JSSResource/mobiledevicegroups/id/1", 200, "validate_update_mobile_device_group.xml")
}

// RegisterUpdateMobileDeviceGroupByNameMock registers PUT /JSSResource/mobiledevicegroups/name/All Mobile Devices → 200.
func (m *MobileDeviceGroupsMock) RegisterUpdateMobileDeviceGroupByNameMock() {
	m.register("PUT", "/JSSResource/mobiledevicegroups/name/All Mobile Devices", 200, "validate_update_mobile_device_group.xml")
}

// RegisterDeleteMobileDeviceGroupByIDMock registers DELETE /JSSResource/mobiledevicegroups/id/1 → 200.
func (m *MobileDeviceGroupsMock) RegisterDeleteMobileDeviceGroupByIDMock() {
	m.register("DELETE", "/JSSResource/mobiledevicegroups/id/1", 200, "")
}

// RegisterDeleteMobileDeviceGroupByNameMock registers DELETE /JSSResource/mobiledevicegroups/name/All Mobile Devices → 200.
func (m *MobileDeviceGroupsMock) RegisterDeleteMobileDeviceGroupByNameMock() {
	m.register("DELETE", "/JSSResource/mobiledevicegroups/name/All Mobile Devices", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/mobiledevicegroups/id/999 → 404.
func (m *MobileDeviceGroupsMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/JSSResource/mobiledevicegroups/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Mobile device group not found")
}

// RegisterConflictErrorMock registers POST /JSSResource/mobiledevicegroups/id/0 → 409.
func (m *MobileDeviceGroupsMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/JSSResource/mobiledevicegroups/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A mobile device group with that name already exists")
}

// ---- transport.HTTPClient implementation ----

func (m *MobileDeviceGroupsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *MobileDeviceGroupsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceGroupsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceGroupsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceGroupsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ transport.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceGroupsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *MobileDeviceGroupsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *MobileDeviceGroupsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MobileDeviceGroupsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MobileDeviceGroupsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *MobileDeviceGroupsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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

func (m *MobileDeviceGroupsMock) RSQLBuilder() transport.RSQLFilterBuilder { return nil }
func (m *MobileDeviceGroupsMock) InvalidateToken() error                    { return nil }
func (m *MobileDeviceGroupsMock) KeepAliveToken() error                     { return nil }
func (m *MobileDeviceGroupsMock) GetLogger() *zap.Logger                    { return m.logger }

// ---- Internal helpers ----

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *MobileDeviceGroupsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("MobileDeviceGroupsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// registerError stores an error response with externalized XML body.
func (m *MobileDeviceGroupsMock) registerError(method, path string, statusCode int, fixture, errMsg string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("MobileDeviceGroupsMock: failed to load error fixture %q: %v", fixture, err))
	}
	m.responses[method+":"+path] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
		errMsg:     errMsg,
	}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *MobileDeviceGroupsMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		headers := http.Header{"Content-Type": {constants.ApplicationXML}}
		return shared.NewMockResponse(http.StatusNotFound, headers, []byte(`<error>no mock registered</error>`)), fmt.Errorf("MobileDeviceGroupsMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {constants.ApplicationXML}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := xml.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("MobileDeviceGroupsMock: unmarshal into result: %w", err)
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
