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
	// errMsg causes the mock to return an error (simulating API or transport failures).
	errMsg string
}

// IBeaconsMock is a test double implementing interfaces.HTTPClient for Classic API iBeacons.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Unlike Jamf Pro API mocks which use JSON, Classic API mocks use XML for
// serialization to match the Classic API wire format.
type IBeaconsMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string // captures the rsqlQuery from the most recent Get call
}

// NewIBeaconsMock returns an empty mock ready for response registration.
func NewIBeaconsMock() *IBeaconsMock {
	return &IBeaconsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *IBeaconsMock) RegisterMocks() {
	m.RegisterListIBeaconsMock()
	m.RegisterGetIBeaconByIDMock()
	m.RegisterGetIBeaconByNameMock()
	m.RegisterCreateIBeaconMock()
	m.RegisterUpdateIBeaconByIDMock()
	m.RegisterUpdateIBeaconByNameMock()
	m.RegisterDeleteIBeaconByIDMock()
	m.RegisterDeleteIBeaconByNameMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *IBeaconsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListIBeaconsMock registers GET /JSSResource/ibeacons → 200.
func (m *IBeaconsMock) RegisterListIBeaconsMock() {
	m.register("GET", "/JSSResource/ibeacons", 200, "validate_list_ibeacons.xml")
}

// RegisterGetIBeaconByIDMock registers GET /JSSResource/ibeacons/id/1 → 200.
func (m *IBeaconsMock) RegisterGetIBeaconByIDMock() {
	m.register("GET", "/JSSResource/ibeacons/id/1", 200, "validate_get_ibeacon.xml")
}

// RegisterGetIBeaconByNameMock registers GET /JSSResource/ibeacons/name/Lobby Beacon → 200.
func (m *IBeaconsMock) RegisterGetIBeaconByNameMock() {
	m.register("GET", "/JSSResource/ibeacons/name/Lobby Beacon", 200, "validate_get_ibeacon.xml")
}

// RegisterCreateIBeaconMock registers POST /JSSResource/ibeacons/id/0 → 201.
func (m *IBeaconsMock) RegisterCreateIBeaconMock() {
	m.register("POST", "/JSSResource/ibeacons/id/0", 201, "validate_create_ibeacon.xml")
}

// RegisterUpdateIBeaconByIDMock registers PUT /JSSResource/ibeacons/id/1 → 200.
func (m *IBeaconsMock) RegisterUpdateIBeaconByIDMock() {
	m.register("PUT", "/JSSResource/ibeacons/id/1", 200, "validate_update_ibeacon.xml")
}

// RegisterUpdateIBeaconByNameMock registers PUT /JSSResource/ibeacons/name/Lobby Beacon → 200.
func (m *IBeaconsMock) RegisterUpdateIBeaconByNameMock() {
	m.register("PUT", "/JSSResource/ibeacons/name/Lobby Beacon", 200, "validate_update_ibeacon.xml")
}

// RegisterDeleteIBeaconByIDMock registers DELETE /JSSResource/ibeacons/id/1 → 200.
func (m *IBeaconsMock) RegisterDeleteIBeaconByIDMock() {
	m.register("DELETE", "/JSSResource/ibeacons/id/1", 200, "")
}

// RegisterDeleteIBeaconByNameMock registers DELETE /JSSResource/ibeacons/name/Lobby Beacon → 200.
func (m *IBeaconsMock) RegisterDeleteIBeaconByNameMock() {
	m.register("DELETE", "/JSSResource/ibeacons/name/Lobby Beacon", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/ibeacons/id/999 → 404.
func (m *IBeaconsMock) RegisterNotFoundErrorMock() {
	body := []byte("<br>An error has occurred.<br>Resource not found<br><br>")
	m.responses["GET:/JSSResource/ibeacons/id/999"] = registeredResponse{
		statusCode: 404,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (404): Resource not found",
	}
}

// RegisterConflictErrorMock registers POST /JSSResource/ibeacons/id/0 → 409.
func (m *IBeaconsMock) RegisterConflictErrorMock() {
	body := []byte("<br>An error has occurred.<br>An iBeacon with that name already exists.<br><br>")
	m.responses["POST:/JSSResource/ibeacons/id/0"] = registeredResponse{
		statusCode: 409,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (409): An iBeacon with that name already exists",
	}
}

// ---- interfaces.HTTPClient implementation ----

func (m *IBeaconsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *IBeaconsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *IBeaconsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *IBeaconsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *IBeaconsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *IBeaconsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *IBeaconsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *IBeaconsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *IBeaconsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *IBeaconsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *IBeaconsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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

func (m *IBeaconsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *IBeaconsMock) InvalidateToken() error                    { return nil }
func (m *IBeaconsMock) KeepAliveToken() error                     { return nil }
func (m *IBeaconsMock) GetLogger() *zap.Logger                    { return m.logger }

// ---- Internal helpers ----

func (m *IBeaconsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("IBeaconsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *IBeaconsMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/xml"}},
			Body:       []byte(`<error>no mock registered</error>`),
		}, fmt.Errorf("IBeaconsMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("IBeaconsMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

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
