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

// CommandFlushMock is a test double implementing interfaces.HTTPClient for command flush operations.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files (if any) in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Classic API mocks use XML for serialization to match the Classic API wire format.
type CommandFlushMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewCommandFlushMock returns an empty mock ready for response registration.
func NewCommandFlushMock() *CommandFlushMock {
	return &CommandFlushMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *CommandFlushMock) RegisterMocks() {
	// Computers - all status combinations
	m.RegisterFlushComputersPendingMock()
	m.RegisterFlushComputersFailedMock()
	m.RegisterFlushComputersPendingAndFailedMock()

	// Computer groups - all status combinations
	m.RegisterFlushComputerGroupsPendingMock()
	m.RegisterFlushComputerGroupsFailedMock()
	m.RegisterFlushComputerGroupsPendingAndFailedMock()

	// Mobile devices - all status combinations
	m.RegisterFlushMobileDevicesPendingMock()
	m.RegisterFlushMobileDevicesFailedMock()
	m.RegisterFlushMobileDevicesPendingAndFailedMock()

	// Mobile device groups - all status combinations
	m.RegisterFlushMobileDeviceGroupsPendingMock()
	m.RegisterFlushMobileDeviceGroupsFailedMock()
	m.RegisterFlushMobileDeviceGroupsPendingAndFailedMock()

	// XML batch operation
	m.RegisterFlushWithXMLMock()
}

// ---- Success responders - Computers ----

// RegisterFlushComputersPendingMock registers DELETE /JSSResource/commandflush/computers/id/123/status/Pending → 204.
func (m *CommandFlushMock) RegisterFlushComputersPendingMock() {
	m.register("DELETE", "/JSSResource/commandflush/computers/id/123/status/Pending", 204, "")
}

// RegisterFlushComputersFailedMock registers DELETE /JSSResource/commandflush/computers/id/123/status/Failed → 204.
func (m *CommandFlushMock) RegisterFlushComputersFailedMock() {
	m.register("DELETE", "/JSSResource/commandflush/computers/id/123/status/Failed", 204, "")
}

// RegisterFlushComputersPendingAndFailedMock registers DELETE /JSSResource/commandflush/computers/id/123/status/Pending%2BFailed → 204.
func (m *CommandFlushMock) RegisterFlushComputersPendingAndFailedMock() {
	m.register("DELETE", "/JSSResource/commandflush/computers/id/123/status/Pending%2BFailed", 204, "")
}

// ---- Success responders - Computer Groups ----

// RegisterFlushComputerGroupsPendingMock registers DELETE /JSSResource/commandflush/computergroups/id/456/status/Pending → 204.
func (m *CommandFlushMock) RegisterFlushComputerGroupsPendingMock() {
	m.register("DELETE", "/JSSResource/commandflush/computergroups/id/456/status/Pending", 204, "")
}

// RegisterFlushComputerGroupsFailedMock registers DELETE /JSSResource/commandflush/computergroups/id/456/status/Failed → 204.
func (m *CommandFlushMock) RegisterFlushComputerGroupsFailedMock() {
	m.register("DELETE", "/JSSResource/commandflush/computergroups/id/456/status/Failed", 204, "")
}

// RegisterFlushComputerGroupsPendingAndFailedMock registers DELETE /JSSResource/commandflush/computergroups/id/456/status/Pending%2BFailed → 204.
func (m *CommandFlushMock) RegisterFlushComputerGroupsPendingAndFailedMock() {
	m.register("DELETE", "/JSSResource/commandflush/computergroups/id/456/status/Pending%2BFailed", 204, "")
}

// ---- Success responders - Mobile Devices ----

// RegisterFlushMobileDevicesPendingMock registers DELETE /JSSResource/commandflush/mobiledevices/id/789/status/Pending → 204.
func (m *CommandFlushMock) RegisterFlushMobileDevicesPendingMock() {
	m.register("DELETE", "/JSSResource/commandflush/mobiledevices/id/789/status/Pending", 204, "")
}

// RegisterFlushMobileDevicesFailedMock registers DELETE /JSSResource/commandflush/mobiledevices/id/789/status/Failed → 204.
func (m *CommandFlushMock) RegisterFlushMobileDevicesFailedMock() {
	m.register("DELETE", "/JSSResource/commandflush/mobiledevices/id/789/status/Failed", 204, "")
}

// RegisterFlushMobileDevicesPendingAndFailedMock registers DELETE /JSSResource/commandflush/mobiledevices/id/789/status/Pending%2BFailed → 204.
func (m *CommandFlushMock) RegisterFlushMobileDevicesPendingAndFailedMock() {
	m.register("DELETE", "/JSSResource/commandflush/mobiledevices/id/789/status/Pending%2BFailed", 204, "")
}

// ---- Success responders - Mobile Device Groups ----

// RegisterFlushMobileDeviceGroupsPendingMock registers DELETE /JSSResource/commandflush/mobiledevicegroups/id/101112/status/Pending → 204.
func (m *CommandFlushMock) RegisterFlushMobileDeviceGroupsPendingMock() {
	m.register("DELETE", "/JSSResource/commandflush/mobiledevicegroups/id/101112/status/Pending", 204, "")
}

// RegisterFlushMobileDeviceGroupsFailedMock registers DELETE /JSSResource/commandflush/mobiledevicegroups/id/101112/status/Failed → 204.
func (m *CommandFlushMock) RegisterFlushMobileDeviceGroupsFailedMock() {
	m.register("DELETE", "/JSSResource/commandflush/mobiledevicegroups/id/101112/status/Failed", 204, "")
}

// RegisterFlushMobileDeviceGroupsPendingAndFailedMock registers DELETE /JSSResource/commandflush/mobiledevicegroups/id/101112/status/Pending%2BFailed → 204.
func (m *CommandFlushMock) RegisterFlushMobileDeviceGroupsPendingAndFailedMock() {
	m.register("DELETE", "/JSSResource/commandflush/mobiledevicegroups/id/101112/status/Pending%2BFailed", 204, "")
}

// ---- Success responders - XML Batch ----

// RegisterFlushWithXMLMock registers DELETE /JSSResource/commandflush → 204.
func (m *CommandFlushMock) RegisterFlushWithXMLMock() {
	m.register("DELETE", "/JSSResource/commandflush", 204, "")
}

// ---- interfaces.HTTPClient implementation ----

func (m *CommandFlushMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}

func (m *CommandFlushMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *CommandFlushMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *CommandFlushMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *CommandFlushMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *CommandFlushMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *CommandFlushMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *CommandFlushMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *CommandFlushMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *CommandFlushMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *CommandFlushMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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

func (m *CommandFlushMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *CommandFlushMock) InvalidateToken() error                     { return nil }
func (m *CommandFlushMock) KeepAliveToken() error                      { return nil }
func (m *CommandFlushMock) GetLogger() *zap.Logger                     { return m.logger }

// ---- Internal helpers ----

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *CommandFlushMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("CommandFlushMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *CommandFlushMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/xml"}},
			Body:       []byte(`<error>no mock registered</error>`),
		}, fmt.Errorf("CommandFlushMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("CommandFlushMock: unmarshal into result: %w", err)
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
