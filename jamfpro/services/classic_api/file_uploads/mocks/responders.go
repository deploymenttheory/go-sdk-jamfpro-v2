package mocks

import (
	"context"
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

// FileUploadsMock is a test double implementing interfaces.HTTPClient for Classic API file uploads.
// Responses are keyed by "METHOD:path". File uploads use PostMultipart and dispatch to the same
// lookup; register with RegisterCreateAttachmentMock for the path your test uses.
//
// Classic API file uploads typically return 200/201 with no body on success.
type FileUploadsMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewFileUploadsMock returns an empty mock ready for response registration.
func NewFileUploadsMock() *FileUploadsMock {
	return &FileUploadsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *FileUploadsMock) RegisterMocks() {
	m.RegisterCreateAttachmentMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *FileUploadsMock) RegisterErrorMocks() {
	m.RegisterInvalidResourceErrorMock()
	m.RegisterPeripheralsNameErrorMock()
}

// ---- Success responders ----

// RegisterCreateAttachmentMock registers POST /JSSResource/fileuploads/policies/id/1 → 200.
// Use this path in tests: resource="policies", idType=ResourceIDTypeID, identifier="1".
func (m *FileUploadsMock) RegisterCreateAttachmentMock() {
	m.register("POST", "/JSSResource/fileuploads/policies/id/1", 200, "")
}

// RegisterCreateAttachmentMockForPath registers a success response for an arbitrary path.
// Use when testing with a custom resource/idType/identifier combination.
func (m *FileUploadsMock) RegisterCreateAttachmentMockForPath(path string) {
	m.register("POST", path, 200, "")
}

// ---- Error responders ----

// RegisterInvalidResourceErrorMock registers POST with invalid resource → 400.
func (m *FileUploadsMock) RegisterInvalidResourceErrorMock() {
	path := "/JSSResource/fileuploads/invalidresource/id/1"
	m.responses["POST:"+path] = registeredResponse{
		statusCode: 400,
		rawBody:    []byte("<br>An error has occurred.<br>Invalid resource type<br><br>"),
		errMsg:     "Jamf Pro Classic API error (400): Invalid resource type",
	}
}

// RegisterPeripheralsNameErrorMock registers an error for peripherals with name type.
func (m *FileUploadsMock) RegisterPeripheralsNameErrorMock() {
	path := "/JSSResource/fileuploads/peripherals/name/somename"
	m.responses["POST:"+path] = registeredResponse{
		statusCode: 400,
		rawBody:    []byte("<br>An error has occurred.<br>Peripherals only support ID type<br><br>"),
		errMsg:     "Jamf Pro Classic API error (400): Peripherals only support ID type",
	}
}

// ---- interfaces.HTTPClient implementation ----

func (m *FileUploadsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}

func (m *FileUploadsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *FileUploadsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *FileUploadsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

// PostMultipart implements interfaces.HTTPClient.
// File uploads dispatch by path; the mock ignores file content for unit tests.
func (m *FileUploadsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *FileUploadsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *FileUploadsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *FileUploadsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *FileUploadsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *FileUploadsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *FileUploadsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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

func (m *FileUploadsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *FileUploadsMock) InvalidateToken() error                     { return nil }
func (m *FileUploadsMock) KeepAliveToken() error                      { return nil }
func (m *FileUploadsMock) GetLogger() *zap.Logger                     { return m.logger }

// ---- Internal helpers ----

func (m *FileUploadsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("FileUploadsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *FileUploadsMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/xml"}},
			Body:       []byte(`<error>no mock registered</error>`),
		}, fmt.Errorf("FileUploadsMock: no response registered for %s %s", method, path)
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
		// File uploads typically return empty body; no unmarshaling needed
		_ = result
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
