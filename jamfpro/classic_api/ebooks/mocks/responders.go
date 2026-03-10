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

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
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

// EbooksMock is a test double implementing client.Client for Classic API ebooks.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Classic API mocks use XML for serialization to match the Classic API wire format.
type EbooksMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewEbooksMock returns an empty mock ready for response registration.
func NewEbooksMock() *EbooksMock {
	return &EbooksMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *EbooksMock) RegisterMocks() {
	m.RegisterListEbooksMock()
	m.RegisterGetEbookByIDMock()
	m.RegisterGetEbookByNameMock()
	m.RegisterGetEbookByNameAndSubsetMock()
	m.RegisterCreateEbookMock()
	m.RegisterUpdateEbookByIDMock()
	m.RegisterUpdateEbookByNameMock()
	m.RegisterDeleteEbookByIDMock()
	m.RegisterDeleteEbookByNameMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *EbooksMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListEbooksMock registers GET /JSSResource/ebooks → 200.
func (m *EbooksMock) RegisterListEbooksMock() {
	m.register("GET", "/JSSResource/ebooks", 200, "validate_list_ebooks.xml")
}

// RegisterGetEbookByIDMock registers GET /JSSResource/ebooks/id/1 → 200.
func (m *EbooksMock) RegisterGetEbookByIDMock() {
	m.register("GET", "/JSSResource/ebooks/id/1", 200, "validate_get_ebook.xml")
}

// RegisterGetEbookByNameMock registers GET /JSSResource/ebooks/name/Sample Ebook → 200.
func (m *EbooksMock) RegisterGetEbookByNameMock() {
	m.register("GET", "/JSSResource/ebooks/name/Sample Ebook", 200, "validate_get_ebook.xml")
}

// RegisterGetEbookByNameAndSubsetMock registers GET /JSSResource/ebooks/name/Sample Ebook/subset/General → 200.
func (m *EbooksMock) RegisterGetEbookByNameAndSubsetMock() {
	m.register("GET", "/JSSResource/ebooks/name/Sample Ebook/subset/General", 200, "validate_get_ebook.xml")
}

// RegisterCreateEbookMock registers POST /JSSResource/ebooks/id/0 → 201.
func (m *EbooksMock) RegisterCreateEbookMock() {
	m.register("POST", "/JSSResource/ebooks/id/0", 201, "validate_create_ebook.xml")
}

// RegisterUpdateEbookByIDMock registers PUT /JSSResource/ebooks/id/1 → 200.
func (m *EbooksMock) RegisterUpdateEbookByIDMock() {
	m.register("PUT", "/JSSResource/ebooks/id/1", 200, "validate_update_ebook.xml")
}

// RegisterUpdateEbookByNameMock registers PUT /JSSResource/ebooks/name/Sample Ebook → 200.
func (m *EbooksMock) RegisterUpdateEbookByNameMock() {
	m.register("PUT", "/JSSResource/ebooks/name/Sample Ebook", 200, "validate_update_ebook.xml")
}

// RegisterDeleteEbookByIDMock registers DELETE /JSSResource/ebooks/id/1 → 200.
func (m *EbooksMock) RegisterDeleteEbookByIDMock() {
	m.register("DELETE", "/JSSResource/ebooks/id/1", 200, "")
}

// RegisterDeleteEbookByNameMock registers DELETE /JSSResource/ebooks/name/Sample Ebook → 200.
func (m *EbooksMock) RegisterDeleteEbookByNameMock() {
	m.register("DELETE", "/JSSResource/ebooks/name/Sample Ebook", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/ebooks/id/999 → 404.
func (m *EbooksMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/JSSResource/ebooks/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

// RegisterConflictErrorMock registers POST /JSSResource/ebooks/id/0 → 409.
func (m *EbooksMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/JSSResource/ebooks/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): An ebook with that name already exists")
}

// ---- client.Client implementation ----

func (m *EbooksMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	_ = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *EbooksMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *EbooksMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *EbooksMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *EbooksMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *EbooksMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *EbooksMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *EbooksMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *EbooksMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *EbooksMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	_ = rsqlQuery
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *EbooksMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	_ = rsqlQuery
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

func (m *EbooksMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *EbooksMock) InvalidateToken() error                { return nil }
func (m *EbooksMock) KeepAliveToken() error                 { return nil }
func (m *EbooksMock) GetLogger() *zap.Logger                { return m.logger }

// ---- Internal helpers ----

// registerError stores an error response with externalized XML body.
func (m *EbooksMock) registerError(method, path string, statusCode int, fixture, errMsg string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("EbooksMock: failed to load error fixture %q: %v", fixture, err))
	}
	m.responses[method+":"+path] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
		errMsg:     errMsg,
	}
}

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *EbooksMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("EbooksMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *EbooksMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		headers := http.Header{"Content-Type": {constants.ApplicationXML}}
		return shared.NewMockResponse(http.StatusNotFound, headers, []byte(`<error>no mock registered</error>`)), fmt.Errorf("EbooksMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {constants.ApplicationXML}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := xml.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("EbooksMock: unmarshal into result: %w", err)
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
