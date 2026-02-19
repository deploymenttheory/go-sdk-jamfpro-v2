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

// DirectoryBindingsMock is a test double implementing interfaces.HTTPClient for Classic API directory bindings.
type DirectoryBindingsMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewDirectoryBindingsMock returns an empty mock ready for response registration.
func NewDirectoryBindingsMock() *DirectoryBindingsMock {
	return &DirectoryBindingsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *DirectoryBindingsMock) RegisterMocks() {
	m.RegisterListDirectoryBindingsMock()
	m.RegisterGetDirectoryBindingByIDMock()
	m.RegisterGetDirectoryBindingByNameMock()
	m.RegisterCreateDirectoryBindingMock()
	m.RegisterUpdateDirectoryBindingByIDMock()
	m.RegisterUpdateDirectoryBindingByNameMock()
	m.RegisterDeleteDirectoryBindingByIDMock()
	m.RegisterDeleteDirectoryBindingByNameMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *DirectoryBindingsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *DirectoryBindingsMock) RegisterListDirectoryBindingsMock() {
	m.register("GET", "/JSSResource/directorybindings", 200, "validate_list_directory_bindings.xml")
}

func (m *DirectoryBindingsMock) RegisterGetDirectoryBindingByIDMock() {
	m.register("GET", "/JSSResource/directorybindings/id/1", 200, "validate_get_directory_binding.xml")
}

func (m *DirectoryBindingsMock) RegisterGetDirectoryBindingByNameMock() {
	m.register("GET", "/JSSResource/directorybindings/name/AD Binding", 200, "validate_get_directory_binding.xml")
}

func (m *DirectoryBindingsMock) RegisterCreateDirectoryBindingMock() {
	m.register("POST", "/JSSResource/directorybindings/id/0", 201, "validate_create_directory_binding.xml")
}

func (m *DirectoryBindingsMock) RegisterUpdateDirectoryBindingByIDMock() {
	m.register("PUT", "/JSSResource/directorybindings/id/1", 200, "validate_update_directory_binding.xml")
}

func (m *DirectoryBindingsMock) RegisterUpdateDirectoryBindingByNameMock() {
	m.register("PUT", "/JSSResource/directorybindings/name/AD Binding", 200, "validate_update_directory_binding.xml")
}

func (m *DirectoryBindingsMock) RegisterDeleteDirectoryBindingByIDMock() {
	m.register("DELETE", "/JSSResource/directorybindings/id/1", 200, "")
}

func (m *DirectoryBindingsMock) RegisterDeleteDirectoryBindingByNameMock() {
	m.register("DELETE", "/JSSResource/directorybindings/name/AD Binding", 200, "")
}

func (m *DirectoryBindingsMock) RegisterNotFoundErrorMock() {
	body := []byte("<br>An error has occurred.<br>Resource not found<br><br>")
	m.responses["GET:/JSSResource/directorybindings/id/999"] = registeredResponse{
		statusCode: 404,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (404): Resource not found",
	}
}

func (m *DirectoryBindingsMock) RegisterConflictErrorMock() {
	body := []byte("<br>An error has occurred.<br>A directory binding with that name already exists.<br><br>")
	m.responses["POST:/JSSResource/directorybindings/id/0"] = registeredResponse{
		statusCode: 409,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (409): A directory binding with that name already exists",
	}
}

func (m *DirectoryBindingsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}
func (m *DirectoryBindingsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *DirectoryBindingsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *DirectoryBindingsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *DirectoryBindingsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *DirectoryBindingsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *DirectoryBindingsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *DirectoryBindingsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *DirectoryBindingsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *DirectoryBindingsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *DirectoryBindingsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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
func (m *DirectoryBindingsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *DirectoryBindingsMock) InvalidateToken() error                    { return nil }
func (m *DirectoryBindingsMock) KeepAliveToken() error                     { return nil }
func (m *DirectoryBindingsMock) GetLogger() *zap.Logger                    { return m.logger }

func (m *DirectoryBindingsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("DirectoryBindingsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *DirectoryBindingsMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/xml"}},
			Body:       []byte(`<error>no mock registered</error>`),
		}, fmt.Errorf("DirectoryBindingsMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("DirectoryBindingsMock: unmarshal into result: %w", err)
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
