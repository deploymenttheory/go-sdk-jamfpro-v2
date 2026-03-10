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
	"go.uber.org/zap"
	"resty.dev/v3"

	mockhelpers "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// MobileDevicesMock is a test double implementing client.Client for Classic API mobile devices.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Classic API mocks use XML for serialization to match the Classic API wire format.
type MobileDevicesMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewMobileDevicesMock returns an empty mock ready for response registration.
func NewMobileDevicesMock() *MobileDevicesMock {
	return &MobileDevicesMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *MobileDevicesMock) RegisterMocks() {
	m.RegisterListMobileDevicesMock()
	m.RegisterGetMobileDeviceByIDMock()
	m.RegisterGetMobileDeviceByNameMock()
	m.RegisterGetMobileDeviceByIDAndDataSubsetMock()
	m.RegisterGetMobileDeviceByNameAndDataSubsetMock()
	m.RegisterCreateMobileDeviceMock()
	m.RegisterUpdateMobileDeviceByIDMock()
	m.RegisterUpdateMobileDeviceByNameMock()
	m.RegisterDeleteMobileDeviceByIDMock()
	m.RegisterDeleteMobileDeviceByNameMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *MobileDevicesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
}

// ---- Success responders ----

// RegisterListMobileDevicesMock registers GET /JSSResource/mobiledevices → 200.
func (m *MobileDevicesMock) RegisterListMobileDevicesMock() {
	m.register("GET", "/JSSResource/mobiledevices", 200, "validate_list_mobile_devices.xml")
}

// RegisterGetMobileDeviceByIDMock registers GET /JSSResource/mobiledevices/id/1 → 200.
func (m *MobileDevicesMock) RegisterGetMobileDeviceByIDMock() {
	m.register("GET", "/JSSResource/mobiledevices/id/1", 200, "validate_get_mobile_device.xml")
}

// RegisterGetMobileDeviceByNameMock registers GET /JSSResource/mobiledevices/name/iPhone-01 → 200.
func (m *MobileDevicesMock) RegisterGetMobileDeviceByNameMock() {
	m.register("GET", "/JSSResource/mobiledevices/name/iPhone-01", 200, "validate_get_mobile_device.xml")
}

// RegisterGetMobileDeviceByIDAndDataSubsetMock registers GET /JSSResource/mobiledevices/id/1/subset/General → 200.
func (m *MobileDevicesMock) RegisterGetMobileDeviceByIDAndDataSubsetMock() {
	m.register("GET", "/JSSResource/mobiledevices/id/1/subset/General", 200, "validate_get_mobile_device.xml")
}

// RegisterGetMobileDeviceByNameAndDataSubsetMock registers GET /JSSResource/mobiledevices/name/iPhone-01/subset/General → 200.
func (m *MobileDevicesMock) RegisterGetMobileDeviceByNameAndDataSubsetMock() {
	m.register("GET", "/JSSResource/mobiledevices/name/iPhone-01/subset/General", 200, "validate_get_mobile_device.xml")
}

// RegisterCreateMobileDeviceMock registers POST /JSSResource/mobiledevices/id/0 → 201.
func (m *MobileDevicesMock) RegisterCreateMobileDeviceMock() {
	m.register("POST", "/JSSResource/mobiledevices/id/0", 201, "validate_create_mobile_device.xml")
}

// RegisterUpdateMobileDeviceByIDMock registers PUT /JSSResource/mobiledevices/id/1 → 200.
func (m *MobileDevicesMock) RegisterUpdateMobileDeviceByIDMock() {
	m.register("PUT", "/JSSResource/mobiledevices/id/1", 200, "validate_update_mobile_device.xml")
}

// RegisterUpdateMobileDeviceByNameMock registers PUT /JSSResource/mobiledevices/name/iPhone-01 → 200.
func (m *MobileDevicesMock) RegisterUpdateMobileDeviceByNameMock() {
	m.register("PUT", "/JSSResource/mobiledevices/name/iPhone-01", 200, "validate_update_mobile_device.xml")
}

// RegisterDeleteMobileDeviceByIDMock registers DELETE /JSSResource/mobiledevices/id/1 → 200.
func (m *MobileDevicesMock) RegisterDeleteMobileDeviceByIDMock() {
	m.register("DELETE", "/JSSResource/mobiledevices/id/1", 200, "")
}

// RegisterDeleteMobileDeviceByNameMock registers DELETE /JSSResource/mobiledevices/name/iPhone-01 → 200.
func (m *MobileDevicesMock) RegisterDeleteMobileDeviceByNameMock() {
	m.register("DELETE", "/JSSResource/mobiledevices/name/iPhone-01", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/mobiledevices/id/999 → 404.
func (m *MobileDevicesMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/JSSResource/mobiledevices/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

// ---- client.Client implementation ----

func (m *MobileDevicesMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *MobileDevicesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDevicesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDevicesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDevicesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDevicesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *MobileDevicesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *MobileDevicesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MobileDevicesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MobileDevicesMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *MobileDevicesMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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
func (m *MobileDevicesMock) NewRequest(ctx context.Context) *client.RequestBuilder {
	return client.NewMockRequestBuilder(ctx, func(method, path string, result any) (*resty.Response, error) {
		return m.dispatch(method, path, result)
	})
}

func (m *MobileDevicesMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *MobileDevicesMock) InvalidateToken() error                { return nil }
func (m *MobileDevicesMock) KeepAliveToken() error                 { return nil }
func (m *MobileDevicesMock) GetLogger() *zap.Logger                { return m.logger }

// ---- Internal helpers ----

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *MobileDevicesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("MobileDevicesMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// registerError stores an error response with externalized XML body.
func (m *MobileDevicesMock) registerError(method, path string, statusCode int, fixture, errMsg string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("MobileDevicesMock: failed to load error fixture %q: %v", fixture, err))
	}
	m.responses[method+":"+path] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
		errMsg:     errMsg,
	}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *MobileDevicesMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		headers := http.Header{"Content-Type": {constants.ApplicationXML}}
		return mockhelpers.NewMockResponse(http.StatusNotFound, headers, []byte(`<error>no mock registered</error>`)), fmt.Errorf("MobileDevicesMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {constants.ApplicationXML}}
	resp := mockhelpers.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := xml.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("MobileDevicesMock: unmarshal into result: %w", err)
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
