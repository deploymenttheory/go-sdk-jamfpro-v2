package mocks

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"slices"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"go.uber.org/zap"
	"resty.dev/v3"
)

// ContentType defines the serialization format for mock responses.
type ContentType string

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// GenericMock is a reusable test double implementing client.Client.
// It can be configured for JSON (Jamf Pro API) or XML (Classic API) responses.
type GenericMock struct {
	name          string
	responses     map[string]registeredResponse
	logger        *zap.Logger
	contentType   ContentType
	LastRSQLQuery map[string]string
	fixtureDir    string // Directory containing fixture files
}

// GenericMockConfig configures a GenericMock instance.
type GenericMockConfig struct {
	Name        string      // Mock name for error messages (e.g., "BuildingsMock")
	ContentType ContentType // JSON or XML
	FixtureDir  string      // Optional: custom fixture directory (defaults to "mocks" in caller's directory)
}

// NewGenericMock creates a new generic mock with the specified configuration.
func NewGenericMock(config GenericMockConfig) *GenericMock {
	if config.Name == "" {
		config.Name = "GenericMock"
	}
	if config.ContentType == "" {
		config.ContentType = constants.ApplicationJSON
	}
	if config.FixtureDir == "" {
		// Default to "mocks" directory relative to the caller's package
		// Walk up the call stack to find the first non-mocks package
		for i := 1; i < 10; i++ {
			_, filename, _, ok := runtime.Caller(i)
			if !ok {
				break
			}
			dir := filepath.Dir(filename)
			// Skip if we're still in the mocks package
			if filepath.Base(dir) == "mocks" {
				continue
			}
			// Found the service package, use its mocks subdirectory
			config.FixtureDir = filepath.Join(dir, "mocks")
			break
		}
	}

	return &GenericMock{
		name:        config.Name,
		responses:   make(map[string]registeredResponse),
		logger:      zap.NewNop(),
		contentType: config.ContentType,
		fixtureDir:  config.FixtureDir,
	}
}

// NewJSONMock creates a mock configured for JSON responses (Jamf Pro API).
func NewJSONMock(name string) *GenericMock {
	return NewGenericMock(GenericMockConfig{
		Name:        name,
		ContentType: constants.ApplicationJSON,
	})
}

// NewXMLMock creates a mock configured for XML responses (Classic API).
func NewXMLMock(name string) *GenericMock {
	return NewGenericMock(GenericMockConfig{
		Name:        name,
		ContentType: constants.ApplicationXML,
	})
}

// Register registers a mock response for the given method and path.
// If fixture is empty, no body is set.
func (m *GenericMock) Register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := m.loadFixture(fixture)
		if err != nil {
			panic(fmt.Sprintf("%s: failed to load fixture %q: %v", m.name, fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// RegisterError registers a mock error response.
// For JSON responses, it attempts to parse error details from the fixture.
func (m *GenericMock) RegisterError(method, path string, statusCode int, fixture string, errMsg string) {
	var body []byte
	if fixture != "" {
		var err error
		body, err = m.loadFixture(fixture)
		if err != nil {
			panic(fmt.Sprintf("%s: failed to load error fixture %q: %v", m.name, fixture, err))
		}

		// If no custom error message provided, try to extract from JSON
		if errMsg == "" && m.contentType == constants.ApplicationJSON {
			var parsed struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			}
			if json.Unmarshal(body, &parsed) == nil {
				errMsg = fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
			}
		}
	}

	if errMsg == "" {
		errMsg = fmt.Sprintf("%s: error response %d", m.name, statusCode)
	}

	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

// RegisterRawBody registers a mock response with raw body bytes.
// Useful for testing error paths with malformed responses.
func (m *GenericMock) RegisterRawBody(method, path string, statusCode int, body []byte) {
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// loadFixture loads a fixture file from the configured fixture directory.
// If not found, falls back to centralized test_fixtures directory for common errors.
func (m *GenericMock) loadFixture(filename string) ([]byte, error) {
	// Try local fixture directory first
	path := filepath.Join(m.fixtureDir, filename)
	data, err := os.ReadFile(path)
	if err == nil {
		return data, nil
	}

	// Fall back to centralized test_fixtures for common error files
	if isCommonErrorFixture(filename) {
		centralPath := filepath.Join(filepath.Dir(m.fixtureDir), "..", "..", "mocks", "test_fixtures", filename)
		data, err = os.ReadFile(centralPath)
		if err == nil {
			return data, nil
		}
	}

	return nil, fmt.Errorf("read fixture %s: %w", filename, err)
}

// isCommonErrorFixture checks if a filename is a common error fixture.
func isCommonErrorFixture(filename string) bool {
	commonErrors := []string{
		"error_not_found.json", "error_not_found.xml",
		"error_conflict.json", "error_conflict.xml",
		"error_internal.json", "error_internal.xml",
		"error_bad_request.json",
	}
	return slices.Contains(commonErrors, filename)
}

// dispatch is the core routing logic for all HTTP methods.
func (m *GenericMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("%s: no response registered for %s %s", m.name, method, path)
	}

	headers := http.Header{"Content-Type": {string(m.contentType)}}
	resp := NewMockResponse(r.statusCode, headers, r.rawBody)

	// If errMsg is set, return response with error (consistent with error handling pattern)
	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	// Unmarshal response body into result if provided
	if result != nil && len(r.rawBody) > 0 {
		// Special case: if result is *[]byte, copy raw bytes directly
		if byteSlicePtr, ok := result.(*[]byte); ok {
			*byteSlicePtr = r.rawBody
		} else {
			var err error
			if m.contentType == constants.ApplicationJSON {
				err = json.Unmarshal(r.rawBody, result)
			} else {
				err = xml.Unmarshal(r.rawBody, result)
			}
			if err != nil {
				return resp, fmt.Errorf("%s: unmarshal into result: %w", m.name, err)
			}
		}
	}

	return resp, nil
}

// -----------------------------------------------------------------------------
// client.Client Interface Implementation
// -----------------------------------------------------------------------------

func (m *GenericMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *GenericMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *GenericMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *GenericMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *GenericMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *GenericMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *GenericMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *GenericMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *GenericMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *GenericMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	m.LastRSQLQuery = rsqlQuery
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *GenericMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil {
		body := resp.Bytes()
		if err := mergePage(body); err != nil {
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
	}
	return resp, nil
}

func (m *GenericMock) NewRequest(ctx context.Context) *client.RequestBuilder {
	return client.NewMockRequestBuilderWithQueryCapture(ctx, func(method, path string, result any) (*resty.Response, error) {
		return m.dispatch(method, path, result)
	}, &m.LastRSQLQuery)
}

func (m *GenericMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *GenericMock) InvalidateToken() error                { return nil }
func (m *GenericMock) KeepAliveToken() error                 { return nil }
func (m *GenericMock) GetLogger() *zap.Logger                { return m.logger }

// Convenience methods for registering common error responses

// RegisterNotFoundError registers a 404 Not Found error for the given method and path.
func (m *GenericMock) RegisterNotFoundError(method, path string) {
	ext := ".json"
	if m.contentType == constants.ApplicationXML {
		ext = ".xml"
	}
	m.RegisterError(method, path, http.StatusNotFound, "error_not_found"+ext, "")
}

// RegisterConflictError registers a 409 Conflict error for the given method and path.
func (m *GenericMock) RegisterConflictError(method, path string) {
	ext := ".json"
	if m.contentType == constants.ApplicationXML {
		ext = ".xml"
	}
	m.RegisterError(method, path, http.StatusConflict, "error_conflict"+ext, "")
}

// RegisterInternalError registers a 500 Internal Server Error for the given method and path.
func (m *GenericMock) RegisterInternalError(method, path string) {
	ext := ".json"
	if m.contentType == constants.ApplicationXML {
		ext = ".xml"
	}
	m.RegisterError(method, path, http.StatusInternalServerError, "error_internal"+ext, "")
}

// RegisterBadRequestError registers a 400 Bad Request error for the given method and path.
func (m *GenericMock) RegisterBadRequestError(method, path string) {
	if m.contentType == constants.ApplicationXML {
		// XML API doesn't typically use 400 errors
		m.RegisterInternalError(method, path)
		return
	}
	m.RegisterError(method, path, http.StatusBadRequest, "error_bad_request.json", "")
}
