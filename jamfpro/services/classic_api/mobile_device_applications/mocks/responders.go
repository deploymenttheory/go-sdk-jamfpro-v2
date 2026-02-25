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

// MobileDeviceApplicationsMock is a test double implementing interfaces.HTTPClient for Classic API mobile device applications.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Classic API mocks use XML for serialization to match the Classic API wire format.
type MobileDeviceApplicationsMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewMobileDeviceApplicationsMock returns an empty mock ready for response registration.
func NewMobileDeviceApplicationsMock() *MobileDeviceApplicationsMock {
	return &MobileDeviceApplicationsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *MobileDeviceApplicationsMock) RegisterMocks() {
	m.RegisterListMobileDeviceApplicationsMock()
	m.RegisterGetMobileDeviceApplicationByIDMock()
	m.RegisterGetMobileDeviceApplicationByNameMock()
	m.RegisterGetMobileDeviceApplicationByBundleIDMock()
	m.RegisterGetMobileDeviceApplicationByBundleIDAndVersionMock()
	m.RegisterGetMobileDeviceApplicationByIDAndSubsetMock()
	m.RegisterGetMobileDeviceApplicationByNameAndSubsetMock()
	m.RegisterCreateMobileDeviceApplicationMock()
	m.RegisterUpdateMobileDeviceApplicationByIDMock()
	m.RegisterUpdateMobileDeviceApplicationByNameMock()
	m.RegisterUpdateMobileDeviceApplicationByBundleIDMock()
	m.RegisterUpdateMobileDeviceApplicationByIDAndVersionMock()
	m.RegisterDeleteMobileDeviceApplicationByIDMock()
	m.RegisterDeleteMobileDeviceApplicationByNameMock()
	m.RegisterDeleteMobileDeviceApplicationByBundleIDMock()
	m.RegisterDeleteMobileDeviceApplicationByBundleIDAndVersionMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *MobileDeviceApplicationsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListMobileDeviceApplicationsMock registers GET /JSSResource/mobiledeviceapplications → 200.
func (m *MobileDeviceApplicationsMock) RegisterListMobileDeviceApplicationsMock() {
	m.register("GET", "/JSSResource/mobiledeviceapplications", 200, "validate_list_mobile_device_applications.xml")
}

// RegisterGetMobileDeviceApplicationByIDMock registers GET /JSSResource/mobiledeviceapplications/id/1 → 200.
func (m *MobileDeviceApplicationsMock) RegisterGetMobileDeviceApplicationByIDMock() {
	m.register("GET", "/JSSResource/mobiledeviceapplications/id/1", 200, "validate_get_mobile_device_application.xml")
}

// RegisterGetMobileDeviceApplicationByNameMock registers GET /JSSResource/mobiledeviceapplications/name/Sample iOS App → 200.
func (m *MobileDeviceApplicationsMock) RegisterGetMobileDeviceApplicationByNameMock() {
	m.register("GET", "/JSSResource/mobiledeviceapplications/name/Sample iOS App", 200, "validate_get_mobile_device_application.xml")
}

// RegisterGetMobileDeviceApplicationByBundleIDMock registers GET /JSSResource/mobiledeviceapplications/bundleid/com.example.app → 200.
func (m *MobileDeviceApplicationsMock) RegisterGetMobileDeviceApplicationByBundleIDMock() {
	m.register("GET", "/JSSResource/mobiledeviceapplications/bundleid/com.example.app", 200, "validate_get_mobile_device_application.xml")
}

// RegisterGetMobileDeviceApplicationByBundleIDAndVersionMock registers GET .../bundleid/com.example.app/version/1.0 → 200.
func (m *MobileDeviceApplicationsMock) RegisterGetMobileDeviceApplicationByBundleIDAndVersionMock() {
	m.register("GET", "/JSSResource/mobiledeviceapplications/bundleid/com.example.app/version/1.0", 200, "validate_get_mobile_device_application.xml")
}

// RegisterGetMobileDeviceApplicationByIDAndSubsetMock registers GET .../id/1/subset/General → 200.
func (m *MobileDeviceApplicationsMock) RegisterGetMobileDeviceApplicationByIDAndSubsetMock() {
	m.register("GET", "/JSSResource/mobiledeviceapplications/id/1/subset/General", 200, "validate_get_mobile_device_application.xml")
}

// RegisterGetMobileDeviceApplicationByNameAndSubsetMock registers GET .../name/Sample iOS App/subset/General → 200.
func (m *MobileDeviceApplicationsMock) RegisterGetMobileDeviceApplicationByNameAndSubsetMock() {
	m.register("GET", "/JSSResource/mobiledeviceapplications/name/Sample iOS App/subset/General", 200, "validate_get_mobile_device_application.xml")
}

// RegisterCreateMobileDeviceApplicationMock registers POST /JSSResource/mobiledeviceapplications/id/0 → 201.
func (m *MobileDeviceApplicationsMock) RegisterCreateMobileDeviceApplicationMock() {
	m.register("POST", "/JSSResource/mobiledeviceapplications/id/0", 201, "validate_create_mobile_device_application.xml")
}

// RegisterUpdateMobileDeviceApplicationByIDMock registers PUT /JSSResource/mobiledeviceapplications/id/1 → 200.
func (m *MobileDeviceApplicationsMock) RegisterUpdateMobileDeviceApplicationByIDMock() {
	m.register("PUT", "/JSSResource/mobiledeviceapplications/id/1", 200, "validate_update_mobile_device_application.xml")
}

// RegisterUpdateMobileDeviceApplicationByNameMock registers PUT .../name/Sample iOS App → 200.
func (m *MobileDeviceApplicationsMock) RegisterUpdateMobileDeviceApplicationByNameMock() {
	m.register("PUT", "/JSSResource/mobiledeviceapplications/name/Sample iOS App", 200, "validate_update_mobile_device_application.xml")
}

// RegisterUpdateMobileDeviceApplicationByBundleIDMock registers PUT .../bundleid/com.example.app → 200.
func (m *MobileDeviceApplicationsMock) RegisterUpdateMobileDeviceApplicationByBundleIDMock() {
	m.register("PUT", "/JSSResource/mobiledeviceapplications/bundleid/com.example.app", 200, "validate_update_mobile_device_application.xml")
}

// RegisterUpdateMobileDeviceApplicationByIDAndVersionMock registers PUT .../id/1/version/1.0 → 200.
func (m *MobileDeviceApplicationsMock) RegisterUpdateMobileDeviceApplicationByIDAndVersionMock() {
	m.register("PUT", "/JSSResource/mobiledeviceapplications/id/1/version/1.0", 200, "validate_update_mobile_device_application.xml")
}

// RegisterDeleteMobileDeviceApplicationByIDMock registers DELETE /JSSResource/mobiledeviceapplications/id/1 → 200.
func (m *MobileDeviceApplicationsMock) RegisterDeleteMobileDeviceApplicationByIDMock() {
	m.register("DELETE", "/JSSResource/mobiledeviceapplications/id/1", 200, "")
}

// RegisterDeleteMobileDeviceApplicationByNameMock registers DELETE .../name/Sample iOS App → 200.
func (m *MobileDeviceApplicationsMock) RegisterDeleteMobileDeviceApplicationByNameMock() {
	m.register("DELETE", "/JSSResource/mobiledeviceapplications/name/Sample iOS App", 200, "")
}

// RegisterDeleteMobileDeviceApplicationByBundleIDMock registers DELETE .../bundleid/com.example.app → 200.
func (m *MobileDeviceApplicationsMock) RegisterDeleteMobileDeviceApplicationByBundleIDMock() {
	m.register("DELETE", "/JSSResource/mobiledeviceapplications/bundleid/com.example.app", 200, "")
}

// RegisterDeleteMobileDeviceApplicationByBundleIDAndVersionMock registers DELETE .../bundleid/com.example.app/version/1.0 → 200.
func (m *MobileDeviceApplicationsMock) RegisterDeleteMobileDeviceApplicationByBundleIDAndVersionMock() {
	m.register("DELETE", "/JSSResource/mobiledeviceapplications/bundleid/com.example.app/version/1.0", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/mobiledeviceapplications/id/999 → 404.
func (m *MobileDeviceApplicationsMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/JSSResource/mobiledeviceapplications/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

// RegisterConflictErrorMock registers POST /JSSResource/mobiledeviceapplications/id/0 → 409.
func (m *MobileDeviceApplicationsMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/JSSResource/mobiledeviceapplications/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A mobile device application with that name already exists")
}

// ---- interfaces.HTTPClient implementation ----

func (m *MobileDeviceApplicationsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	_ = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *MobileDeviceApplicationsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceApplicationsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceApplicationsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceApplicationsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceApplicationsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *MobileDeviceApplicationsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *MobileDeviceApplicationsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MobileDeviceApplicationsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MobileDeviceApplicationsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	_ = rsqlQuery
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *MobileDeviceApplicationsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	_ = rsqlQuery
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

func (m *MobileDeviceApplicationsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *MobileDeviceApplicationsMock) InvalidateToken() error                     { return nil }
func (m *MobileDeviceApplicationsMock) KeepAliveToken() error                       { return nil }
func (m *MobileDeviceApplicationsMock) GetLogger() *zap.Logger                      { return m.logger }

// ---- Internal helpers ----

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *MobileDeviceApplicationsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("MobileDeviceApplicationsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// registerError stores an error response with externalized XML body.
func (m *MobileDeviceApplicationsMock) registerError(method, path string, statusCode int, fixture, errMsg string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("MobileDeviceApplicationsMock: failed to load error fixture %q: %v", fixture, err))
	}
	m.responses[method+":"+path] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
		errMsg:     errMsg,
	}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *MobileDeviceApplicationsMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/xml"}},
			Body:       []byte(`<error>no mock registered</error>`),
		}, fmt.Errorf("MobileDeviceApplicationsMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("MobileDeviceApplicationsMock: unmarshal into result: %w", err)
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
