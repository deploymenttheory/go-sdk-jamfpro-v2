package mocks

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
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

// ComputerGroupsMock is a test double implementing interfaces.HTTPClient for Classic API computer groups.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Classic API mocks use XML for serialization to match the Classic API wire format.
type ComputerGroupsMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewComputerGroupsMock returns an empty mock ready for response registration.
func NewComputerGroupsMock() *ComputerGroupsMock {
	return &ComputerGroupsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *ComputerGroupsMock) RegisterMocks() {
	m.RegisterListComputerGroupsMock()
	m.RegisterGetComputerGroupByIDMock()
	m.RegisterGetComputerGroupByNameMock()
	m.RegisterCreateComputerGroupMock()
	m.RegisterUpdateComputerGroupByIDMock()
	m.RegisterUpdateComputerGroupByNameMock()
	m.RegisterDeleteComputerGroupByIDMock()
	m.RegisterDeleteComputerGroupByNameMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *ComputerGroupsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListComputerGroupsMock registers GET /JSSResource/computergroups → 200.
func (m *ComputerGroupsMock) RegisterListComputerGroupsMock() {
	m.register("GET", "/JSSResource/computergroups", 200, "validate_list_computer_groups.xml")
}

// RegisterGetComputerGroupByIDMock registers GET /JSSResource/computergroups/id/1 → 200.
func (m *ComputerGroupsMock) RegisterGetComputerGroupByIDMock() {
	m.register("GET", "/JSSResource/computergroups/id/1", 200, "validate_get_computer_group.xml")
}

// RegisterGetComputerGroupByNameMock registers GET /JSSResource/computergroups/name/All Managed Clients → 200.
func (m *ComputerGroupsMock) RegisterGetComputerGroupByNameMock() {
	m.register("GET", "/JSSResource/computergroups/name/All Managed Clients", 200, "validate_get_computer_group.xml")
}

// RegisterCreateComputerGroupMock registers POST /JSSResource/computergroups/id/0 → 201.
func (m *ComputerGroupsMock) RegisterCreateComputerGroupMock() {
	m.register("POST", "/JSSResource/computergroups/id/0", 201, "validate_create_computer_group.xml")
}

// RegisterUpdateComputerGroupByIDMock registers PUT /JSSResource/computergroups/id/1 → 200.
func (m *ComputerGroupsMock) RegisterUpdateComputerGroupByIDMock() {
	m.register("PUT", "/JSSResource/computergroups/id/1", 200, "validate_update_computer_group.xml")
}

// RegisterUpdateComputerGroupByNameMock registers PUT /JSSResource/computergroups/name/All Managed Clients → 200.
func (m *ComputerGroupsMock) RegisterUpdateComputerGroupByNameMock() {
	m.register("PUT", "/JSSResource/computergroups/name/All Managed Clients", 200, "validate_update_computer_group.xml")
}

// RegisterDeleteComputerGroupByIDMock registers DELETE /JSSResource/computergroups/id/1 → 200.
func (m *ComputerGroupsMock) RegisterDeleteComputerGroupByIDMock() {
	m.register("DELETE", "/JSSResource/computergroups/id/1", 200, "")
}

// RegisterDeleteComputerGroupByNameMock registers DELETE /JSSResource/computergroups/name/All Managed Clients → 200.
func (m *ComputerGroupsMock) RegisterDeleteComputerGroupByNameMock() {
	m.register("DELETE", "/JSSResource/computergroups/name/All Managed Clients", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/computergroups/id/999 → 404.
func (m *ComputerGroupsMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/JSSResource/computergroups/id/999", 404, "Computer group not found")
}

// RegisterConflictErrorMock registers POST /JSSResource/computergroups/id/0 → 409.
func (m *ComputerGroupsMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/JSSResource/computergroups/id/0", 409, "Computer group name already exists")
}

// ---- Implementation ----

func (m *ComputerGroupsMock) register(method, path string, statusCode int, filename string) {
	key := method + ":" + path
	var body []byte
	if filename != "" {
		body = m.readFixture(filename)
	}
	m.responses[key] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
	}
}

func (m *ComputerGroupsMock) registerError(method, path string, statusCode int, errMsg string) {
	key := method + ":" + path
	m.responses[key] = registeredResponse{
		statusCode: statusCode,
		errMsg:     errMsg,
	}
}

func (m *ComputerGroupsMock) readFixture(filename string) []byte {
	cwd, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("cannot determine working directory: %v", err))
	}

	path := filepath.Join(cwd, "mocks", filename)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		path = filepath.Join(cwd, filename)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("failed to read fixture %q: %v", filename, err))
	}
	return data
}

// Get implements interfaces.HTTPClient.Get for the mock.
func (m *ComputerGroupsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, out any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery

	key := "GET:" + path
	reg, found := m.responses[key]
	if !found {
		return nil, fmt.Errorf("no mock registered for GET %s", path)
	}
	if reg.errMsg != "" {
		return &interfaces.Response{StatusCode: reg.statusCode}, fmt.Errorf(reg.errMsg)
	}
	if out != nil && len(reg.rawBody) > 0 {
		if err := xml.Unmarshal(reg.rawBody, out); err != nil {
			return nil, fmt.Errorf("unmarshal error for GET %s: %w", path, err)
		}
	}
	return &interfaces.Response{StatusCode: reg.statusCode}, nil
}

// Post implements interfaces.HTTPClient.Post for the mock.
func (m *ComputerGroupsMock) Post(ctx context.Context, path string, body any, headers map[string]string, out any) (*interfaces.Response, error) {
	key := "POST:" + path
	reg, found := m.responses[key]
	if !found {
		return nil, fmt.Errorf("no mock registered for POST %s", path)
	}
	if reg.errMsg != "" {
		return &interfaces.Response{StatusCode: reg.statusCode}, fmt.Errorf(reg.errMsg)
	}
	if out != nil && len(reg.rawBody) > 0 {
		if err := xml.Unmarshal(reg.rawBody, out); err != nil {
			return nil, fmt.Errorf("unmarshal error for POST %s: %w", path, err)
		}
	}
	return &interfaces.Response{StatusCode: reg.statusCode}, nil
}

// Put implements interfaces.HTTPClient.Put for the mock.
func (m *ComputerGroupsMock) Put(ctx context.Context, path string, body any, headers map[string]string, out any) (*interfaces.Response, error) {
	key := "PUT:" + path
	reg, found := m.responses[key]
	if !found {
		return nil, fmt.Errorf("no mock registered for PUT %s", path)
	}
	if reg.errMsg != "" {
		return &interfaces.Response{StatusCode: reg.statusCode}, fmt.Errorf(reg.errMsg)
	}
	if out != nil && len(reg.rawBody) > 0 {
		if err := xml.Unmarshal(reg.rawBody, out); err != nil {
			return nil, fmt.Errorf("unmarshal error for PUT %s: %w", path, err)
		}
	}
	return &interfaces.Response{StatusCode: reg.statusCode}, nil
}

// Delete implements interfaces.HTTPClient.Delete for the mock.
func (m *ComputerGroupsMock) Delete(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, out any) (*interfaces.Response, error) {
	key := "DELETE:" + path
	reg, found := m.responses[key]
	if !found {
		return nil, fmt.Errorf("no mock registered for DELETE %s", path)
	}
	if reg.errMsg != "" {
		return &interfaces.Response{StatusCode: reg.statusCode}, fmt.Errorf(reg.errMsg)
	}
	return &interfaces.Response{StatusCode: reg.statusCode}, nil
}

// PostWithQuery is a placeholder for interfaces.HTTPClient compliance (not used in Classic API).
func (m *ComputerGroupsMock) PostWithQuery(ctx context.Context, path string, rsqlQuery map[string]string, body any, headers map[string]string, out any) (*interfaces.Response, error) {
	return nil, fmt.Errorf("PostWithQuery not implemented for Classic API mock")
}

// PostForm is a placeholder for interfaces.HTTPClient compliance (not used in Classic API).
func (m *ComputerGroupsMock) PostForm(ctx context.Context, path string, formData map[string]string, headers map[string]string, out any) (*interfaces.Response, error) {
	return nil, fmt.Errorf("PostForm not implemented for Classic API mock")
}

// PostMultipart is a placeholder for interfaces.HTTPClient compliance (not used in Classic API).
func (m *ComputerGroupsMock) PostMultipart(ctx context.Context, path string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback interfaces.MultipartProgressCallback, out any) (*interfaces.Response, error) {
	return nil, fmt.Errorf("PostMultipart not implemented for Classic API mock")
}

// Patch is a placeholder for interfaces.HTTPClient compliance (not used in Classic API).
func (m *ComputerGroupsMock) Patch(ctx context.Context, path string, body any, headers map[string]string, out any) (*interfaces.Response, error) {
	return nil, fmt.Errorf("Patch not implemented for Classic API mock")
}

// DeleteWithBody is a placeholder for interfaces.HTTPClient compliance (not used in Classic API).
func (m *ComputerGroupsMock) DeleteWithBody(ctx context.Context, path string, body any, headers map[string]string, out any) (*interfaces.Response, error) {
	return nil, fmt.Errorf("DeleteWithBody not implemented for Classic API mock")
}

// GetBytes is a placeholder for interfaces.HTTPClient compliance (not used in Classic API).
func (m *ComputerGroupsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string) (*interfaces.Response, []byte, error) {
	return nil, nil, fmt.Errorf("GetBytes not implemented for Classic API mock")
}

// GetPaginated is a placeholder for interfaces.HTTPClient compliance (not used in Classic API).
func (m *ComputerGroupsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, mergePage func(pageData []byte) error) (*interfaces.Response, error) {
	return nil, fmt.Errorf("GetPaginated not implemented for Classic API mock")
}

// RSQLBuilder is a placeholder for interfaces.HTTPClient compliance (not used in Classic API).
func (m *ComputerGroupsMock) RSQLBuilder() interfaces.RSQLFilterBuilder {
	return nil
}

// InvalidateToken is a placeholder for interfaces.HTTPClient compliance (not used in Classic API).
func (m *ComputerGroupsMock) InvalidateToken() error {
	return fmt.Errorf("InvalidateToken not implemented for mock")
}

// KeepAliveToken is a placeholder for interfaces.HTTPClient compliance (not used in Classic API).
func (m *ComputerGroupsMock) KeepAliveToken() error {
	return fmt.Errorf("KeepAliveToken not implemented for mock")
}

// GetLogger is a placeholder for interfaces.HTTPClient compliance.
func (m *ComputerGroupsMock) GetLogger() *zap.Logger {
	return m.logger
}
