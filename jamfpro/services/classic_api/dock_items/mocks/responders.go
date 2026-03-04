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
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"go.uber.org/zap"
	"resty.dev/v3"
)

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// DockItemsMock is a test double implementing interfaces.HTTPClient for Classic API dock items.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Classic API mocks use XML for serialization to match the Classic API wire format.
type DockItemsMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewDockItemsMock returns an empty mock ready for response registration.
func NewDockItemsMock() *DockItemsMock {
	return &DockItemsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *DockItemsMock) RegisterMocks() {
	m.RegisterListDockItemsMock()
	m.RegisterGetDockItemByIDMock()
	m.RegisterGetDockItemByNameMock()
	m.RegisterCreateDockItemMock()
	m.RegisterUpdateDockItemByIDMock()
	m.RegisterUpdateDockItemByNameMock()
	m.RegisterDeleteDockItemByIDMock()
	m.RegisterDeleteDockItemByNameMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *DockItemsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListDockItemsMock registers GET /JSSResource/dockitems → 200.
func (m *DockItemsMock) RegisterListDockItemsMock() {
	m.register("GET", "/JSSResource/dockitems", 200, "validate_list_dock_items.xml")
}

// RegisterGetDockItemByIDMock registers GET /JSSResource/dockitems/id/1 → 200.
func (m *DockItemsMock) RegisterGetDockItemByIDMock() {
	m.register("GET", "/JSSResource/dockitems/id/1", 200, "validate_get_dock_item.xml")
}

// RegisterGetDockItemByNameMock registers GET /JSSResource/dockitems/name/Safari → 200.
func (m *DockItemsMock) RegisterGetDockItemByNameMock() {
	m.register("GET", "/JSSResource/dockitems/name/Safari", 200, "validate_get_dock_item.xml")
}

// RegisterCreateDockItemMock registers POST /JSSResource/dockitems/id/0 → 201.
func (m *DockItemsMock) RegisterCreateDockItemMock() {
	m.register("POST", "/JSSResource/dockitems/id/0", 201, "validate_create_dock_item.xml")
}

// RegisterUpdateDockItemByIDMock registers PUT /JSSResource/dockitems/id/1 → 200.
func (m *DockItemsMock) RegisterUpdateDockItemByIDMock() {
	m.register("PUT", "/JSSResource/dockitems/id/1", 200, "validate_update_dock_item.xml")
}

// RegisterUpdateDockItemByNameMock registers PUT /JSSResource/dockitems/name/Safari → 200.
func (m *DockItemsMock) RegisterUpdateDockItemByNameMock() {
	m.register("PUT", "/JSSResource/dockitems/name/Safari", 200, "validate_update_dock_item.xml")
}

// RegisterDeleteDockItemByIDMock registers DELETE /JSSResource/dockitems/id/1 → 200.
func (m *DockItemsMock) RegisterDeleteDockItemByIDMock() {
	m.register("DELETE", "/JSSResource/dockitems/id/1", 200, "")
}

// RegisterDeleteDockItemByNameMock registers DELETE /JSSResource/dockitems/name/Safari → 200.
func (m *DockItemsMock) RegisterDeleteDockItemByNameMock() {
	m.register("DELETE", "/JSSResource/dockitems/name/Safari", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/dockitems/id/999 → 404.
func (m *DockItemsMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/JSSResource/dockitems/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

// RegisterConflictErrorMock registers POST /JSSResource/dockitems/id/0 → 409.
func (m *DockItemsMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/JSSResource/dockitems/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A dock item with that name already exists")
}

// ---- interfaces.HTTPClient implementation ----

func (m *DockItemsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *DockItemsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *DockItemsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *DockItemsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *DockItemsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *DockItemsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *DockItemsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *DockItemsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *DockItemsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *DockItemsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *DockItemsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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

func (m *DockItemsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *DockItemsMock) InvalidateToken() error                    { return nil }
func (m *DockItemsMock) KeepAliveToken() error                     { return nil }
func (m *DockItemsMock) GetLogger() *zap.Logger                    { return m.logger }

// ---- Internal helpers ----

// registerError stores an error response with externalized XML body.
func (m *DockItemsMock) registerError(method, path string, statusCode int, fixture, errMsg string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("DockItemsMock: failed to load error fixture %q: %v", fixture, err))
	}
	m.responses[method+":"+path] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
		errMsg:     errMsg,
	}
}

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *DockItemsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("DockItemsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *DockItemsMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		headers := http.Header{"Content-Type": {mime.ApplicationXML}}
		return shared.NewMockResponse(http.StatusNotFound, headers, []byte(`<error>no mock registered</error>`)), fmt.Errorf("DockItemsMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {mime.ApplicationXML}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := xml.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("DockItemsMock: unmarshal into result: %w", err)
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
