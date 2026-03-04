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

// FileShareDistributionPointsMock is a test double implementing interfaces.HTTPClient for Classic API file share distribution points.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Classic API mocks use XML for serialization to match the Classic API wire format.
type FileShareDistributionPointsMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewFileShareDistributionPointsMock returns an empty mock ready for response registration.
func NewFileShareDistributionPointsMock() *FileShareDistributionPointsMock {
	return &FileShareDistributionPointsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *FileShareDistributionPointsMock) RegisterMocks() {
	m.RegisterListFileShareDistributionPointsMock()
	m.RegisterGetFileShareDistributionPointByIDMock()
	m.RegisterGetFileShareDistributionPointByNameMock()
	m.RegisterCreateFileShareDistributionPointMock()
	m.RegisterUpdateFileShareDistributionPointByIDMock()
	m.RegisterUpdateFileShareDistributionPointByNameMock()
	m.RegisterDeleteFileShareDistributionPointByIDMock()
	m.RegisterDeleteFileShareDistributionPointByNameMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *FileShareDistributionPointsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListFileShareDistributionPointsMock registers GET /JSSResource/distributionpoints → 200.
func (m *FileShareDistributionPointsMock) RegisterListFileShareDistributionPointsMock() {
	m.register("GET", "/JSSResource/distributionpoints", 200, "validate_list_file_share_distribution_points.xml")
}

// RegisterGetFileShareDistributionPointByIDMock registers GET /JSSResource/distributionpoints/id/1 → 200.
func (m *FileShareDistributionPointsMock) RegisterGetFileShareDistributionPointByIDMock() {
	m.register("GET", "/JSSResource/distributionpoints/id/1", 200, "validate_get_file_share_distribution_point.xml")
}

// RegisterGetFileShareDistributionPointByNameMock registers GET /JSSResource/distributionpoints/name/Main File Share DP → 200.
func (m *FileShareDistributionPointsMock) RegisterGetFileShareDistributionPointByNameMock() {
	m.register("GET", "/JSSResource/distributionpoints/name/Main File Share DP", 200, "validate_get_file_share_distribution_point.xml")
}

// RegisterCreateFileShareDistributionPointMock registers POST /JSSResource/distributionpoints/id/0 → 201.
func (m *FileShareDistributionPointsMock) RegisterCreateFileShareDistributionPointMock() {
	m.register("POST", "/JSSResource/distributionpoints/id/0", 201, "validate_create_file_share_distribution_point.xml")
}

// RegisterUpdateFileShareDistributionPointByIDMock registers PUT /JSSResource/distributionpoints/id/1 → 200.
func (m *FileShareDistributionPointsMock) RegisterUpdateFileShareDistributionPointByIDMock() {
	m.register("PUT", "/JSSResource/distributionpoints/id/1", 200, "validate_update_file_share_distribution_point.xml")
}

// RegisterUpdateFileShareDistributionPointByNameMock registers PUT /JSSResource/distributionpoints/name/Main File Share DP → 200.
func (m *FileShareDistributionPointsMock) RegisterUpdateFileShareDistributionPointByNameMock() {
	m.register("PUT", "/JSSResource/distributionpoints/name/Main File Share DP", 200, "validate_update_file_share_distribution_point.xml")
}

// RegisterDeleteFileShareDistributionPointByIDMock registers DELETE /JSSResource/distributionpoints/id/1 → 200.
func (m *FileShareDistributionPointsMock) RegisterDeleteFileShareDistributionPointByIDMock() {
	m.register("DELETE", "/JSSResource/distributionpoints/id/1", 200, "")
}

// RegisterDeleteFileShareDistributionPointByNameMock registers DELETE /JSSResource/distributionpoints/name/Main File Share DP → 200.
func (m *FileShareDistributionPointsMock) RegisterDeleteFileShareDistributionPointByNameMock() {
	m.register("DELETE", "/JSSResource/distributionpoints/name/Main File Share DP", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/distributionpoints/id/999 → 404.
func (m *FileShareDistributionPointsMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/JSSResource/distributionpoints/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

// RegisterConflictErrorMock registers POST /JSSResource/distributionpoints/id/0 → 409.
func (m *FileShareDistributionPointsMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/JSSResource/distributionpoints/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A distribution point with that name already exists")
}

// ---- interfaces.HTTPClient implementation ----

func (m *FileShareDistributionPointsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *FileShareDistributionPointsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *FileShareDistributionPointsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *FileShareDistributionPointsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *FileShareDistributionPointsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *FileShareDistributionPointsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *FileShareDistributionPointsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *FileShareDistributionPointsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *FileShareDistributionPointsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *FileShareDistributionPointsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *FileShareDistributionPointsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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

func (m *FileShareDistributionPointsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *FileShareDistributionPointsMock) InvalidateToken() error                    { return nil }
func (m *FileShareDistributionPointsMock) KeepAliveToken() error                     { return nil }
func (m *FileShareDistributionPointsMock) GetLogger() *zap.Logger                    { return m.logger }

// ---- Internal helpers ----

// registerError stores an error response with externalized XML body.
func (m *FileShareDistributionPointsMock) registerError(method, path string, statusCode int, fixture, errMsg string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("FileShareDistributionPointsMock: failed to load error fixture %q: %v", fixture, err))
	}
	m.responses[method+":"+path] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
		errMsg:     errMsg,
	}
}

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *FileShareDistributionPointsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("FileShareDistributionPointsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *FileShareDistributionPointsMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		headers := http.Header{"Content-Type": {mime.ApplicationXML}}
		return shared.NewMockResponse(http.StatusNotFound, headers, []byte(`<error>no mock registered</error>`)), fmt.Errorf("FileShareDistributionPointsMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {mime.ApplicationXML}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := xml.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("FileShareDistributionPointsMock: unmarshal into result: %w", err)
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
