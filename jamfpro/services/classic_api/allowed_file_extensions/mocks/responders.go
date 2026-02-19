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

// AllowedFileExtensionsMock is a test double implementing interfaces.HTTPClient for Classic API allowed file extensions.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Unlike Jamf Pro API mocks which use JSON, Classic API mocks use XML for
// serialization to match the Classic API wire format.
type AllowedFileExtensionsMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string // captures the rsqlQuery from the most recent Get call
}

// NewAllowedFileExtensionsMock returns an empty mock ready for response registration.
func NewAllowedFileExtensionsMock() *AllowedFileExtensionsMock {
	return &AllowedFileExtensionsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *AllowedFileExtensionsMock) RegisterMocks() {
	m.RegisterListAllowedFileExtensionsMock()
	m.RegisterGetAllowedFileExtensionByIDMock()
	m.RegisterGetAllowedFileExtensionByExtensionMock()
	m.RegisterCreateAllowedFileExtensionMock()
	m.RegisterDeleteAllowedFileExtensionByIDMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *AllowedFileExtensionsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListAllowedFileExtensionsMock registers GET /JSSResource/allowedfileextensions → 200.
func (m *AllowedFileExtensionsMock) RegisterListAllowedFileExtensionsMock() {
	m.register("GET", "/JSSResource/allowedfileextensions", 200, "validate_list_allowed_file_extensions.xml")
}

// RegisterGetAllowedFileExtensionByIDMock registers GET /JSSResource/allowedfileextensions/id/1 → 200.
func (m *AllowedFileExtensionsMock) RegisterGetAllowedFileExtensionByIDMock() {
	m.register("GET", "/JSSResource/allowedfileextensions/id/1", 200, "validate_get_allowed_file_extension.xml")
}

// RegisterGetAllowedFileExtensionByExtensionMock registers GET /JSSResource/allowedfileextensions/extension/dmg → 200.
func (m *AllowedFileExtensionsMock) RegisterGetAllowedFileExtensionByExtensionMock() {
	m.register("GET", "/JSSResource/allowedfileextensions/extension/dmg", 200, "validate_get_allowed_file_extension.xml")
}

// RegisterCreateAllowedFileExtensionMock registers POST /JSSResource/allowedfileextensions/id/0 → 201.
func (m *AllowedFileExtensionsMock) RegisterCreateAllowedFileExtensionMock() {
	m.register("POST", "/JSSResource/allowedfileextensions/id/0", 201, "validate_create_allowed_file_extension.xml")
}

// RegisterDeleteAllowedFileExtensionByIDMock registers DELETE /JSSResource/allowedfileextensions/id/1 → 200.
func (m *AllowedFileExtensionsMock) RegisterDeleteAllowedFileExtensionByIDMock() {
	m.register("DELETE", "/JSSResource/allowedfileextensions/id/1", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/allowedfileextensions/id/999 → 404.
func (m *AllowedFileExtensionsMock) RegisterNotFoundErrorMock() {
	body := []byte("<br>An error has occurred.<br>Resource not found<br><br>")
	m.responses["GET:/JSSResource/allowedfileextensions/id/999"] = registeredResponse{
		statusCode: 404,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (404): Resource not found",
	}
}

// RegisterConflictErrorMock registers POST /JSSResource/allowedfileextensions/id/0 → 409.
func (m *AllowedFileExtensionsMock) RegisterConflictErrorMock() {
	body := []byte("<br>An error has occurred.<br>An allowed file extension with that name already exists.<br><br>")
	m.responses["POST:/JSSResource/allowedfileextensions/id/0"] = registeredResponse{
		statusCode: 409,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (409): An allowed file extension with that name already exists",
	}
}

// ---- interfaces.HTTPClient implementation ----

func (m *AllowedFileExtensionsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *AllowedFileExtensionsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *AllowedFileExtensionsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *AllowedFileExtensionsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *AllowedFileExtensionsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *AllowedFileExtensionsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *AllowedFileExtensionsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *AllowedFileExtensionsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *AllowedFileExtensionsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *AllowedFileExtensionsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *AllowedFileExtensionsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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

func (m *AllowedFileExtensionsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *AllowedFileExtensionsMock) InvalidateToken() error                    { return nil }
func (m *AllowedFileExtensionsMock) KeepAliveToken() error                     { return nil }
func (m *AllowedFileExtensionsMock) GetLogger() *zap.Logger                    { return m.logger }

// ---- Internal helpers ----

func (m *AllowedFileExtensionsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("AllowedFileExtensionsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *AllowedFileExtensionsMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/xml"}},
			Body:       []byte(`<error>no mock registered</error>`),
		}, fmt.Errorf("AllowedFileExtensionsMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("AllowedFileExtensionsMock: unmarshal into result: %w", err)
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
