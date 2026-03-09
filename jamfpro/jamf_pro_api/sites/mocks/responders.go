package mocks

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"resty.dev/v3"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"go.uber.org/zap"
)

//go:embed *.json
var fixtureFS embed.FS

type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// SitesMock is a test double implementing transport.HTTPClient for sites operations.
type SitesMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewSitesMock returns an empty mock ready for response registration.
func NewSitesMock() *SitesMock {
	return &SitesMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses.
func (m *SitesMock) RegisterMocks() {
	m.RegisterListV1Mock()
	m.RegisterGetObjectsByIDV1Mock()
}

func (m *SitesMock) register(method, path string, statusCode int, fixtureFile string) {
	key := method + " " + path
	var body []byte
	if fixtureFile != "" {
		data, err := fixtureFS.ReadFile(fixtureFile)
		if err != nil {
			panic(fmt.Sprintf("SitesMock: failed to load fixture %q: %v", fixtureFile, err))
		}
		body = data
	}
	m.responses[key] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
	}
}

// RegisterListV1Mock registers a successful response for ListV1.
func (m *SitesMock) RegisterListV1Mock() {
	m.register("GET", "/api/v1/sites", 200, "validate_list.json")
}

// RegisterGetObjectsByIDV1Mock registers a successful response for GetObjectsByIDV1.
func (m *SitesMock) RegisterGetObjectsByIDV1Mock() {
	m.register("GET", "/api/v1/sites/1/objects", 200, "validate_objects.json")
}

// Get implements transport.HTTPClient.
func (m *SitesMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	key := "GET " + path
	resp, ok := m.responses[key]
	if !ok {
		return shared.NewMockResponse(404, http.Header{}, nil), fmt.Errorf("no mock registered for GET %s", path)
	}
	if resp.errMsg != "" {
		return nil, fmt.Errorf("%s", resp.errMsg)
	}
	if result != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return shared.NewMockResponse(resp.statusCode, http.Header{}, resp.rawBody), nil
}

// GetPaginated implements transport.HTTPClient for paginated endpoints.
func (m *SitesMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, mergePage func(pageData []byte) error) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	key := "GET " + path
	resp, ok := m.responses[key]
	if !ok {
		return shared.NewMockResponse(404, http.Header{}, nil), fmt.Errorf("no mock registered for GET %s", path)
	}
	if resp.errMsg != "" {
		return nil, fmt.Errorf("%s", resp.errMsg)
	}
	if mergePage != nil && len(resp.rawBody) > 0 {
		var page struct {
			Results json.RawMessage `json:"results"`
		}
		if err := json.Unmarshal(resp.rawBody, &page); err != nil {
			return nil, fmt.Errorf("merge page failed: %w", err)
		}
		if err := mergePage(page.Results); err != nil {
			return nil, fmt.Errorf("merge page failed: %w", err)
		}
	}
	return shared.NewMockResponse(resp.statusCode, http.Header{}, resp.rawBody), nil
}

// Post implements transport.HTTPClient.
func (m *SitesMock) Post(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	return shared.NewMockResponse(http.StatusMethodNotAllowed, http.Header{}, nil), nil
}

// PostWithQuery implements transport.HTTPClient.
func (m *SitesMock) PostWithQuery(ctx context.Context, path string, rsqlQuery map[string]string, body any, headers map[string]string, result any) (*resty.Response, error) {
	return shared.NewMockResponse(http.StatusMethodNotAllowed, http.Header{}, nil), nil
}

// PostForm implements transport.HTTPClient.
func (m *SitesMock) PostForm(ctx context.Context, path string, formData map[string]string, headers map[string]string, result any) (*resty.Response, error) {
	return shared.NewMockResponse(http.StatusMethodNotAllowed, http.Header{}, nil), nil
}

// PostMultipart implements transport.HTTPClient.
func (m *SitesMock) PostMultipart(ctx context.Context, path string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback transport.MultipartProgressCallback, result any) (*resty.Response, error) {
	return shared.NewMockResponse(http.StatusMethodNotAllowed, http.Header{}, nil), nil
}

// Put implements transport.HTTPClient.
func (m *SitesMock) Put(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	return shared.NewMockResponse(http.StatusMethodNotAllowed, http.Header{}, nil), nil
}

// Patch implements transport.HTTPClient.
func (m *SitesMock) Patch(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	return shared.NewMockResponse(http.StatusMethodNotAllowed, http.Header{}, nil), nil
}

// Delete implements transport.HTTPClient.
func (m *SitesMock) Delete(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*resty.Response, error) {
	return shared.NewMockResponse(http.StatusMethodNotAllowed, http.Header{}, nil), nil
}

// DeleteWithBody implements transport.HTTPClient.
func (m *SitesMock) DeleteWithBody(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	return shared.NewMockResponse(http.StatusMethodNotAllowed, http.Header{}, nil), nil
}

// GetBytes implements transport.HTTPClient.
func (m *SitesMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string) (*resty.Response, []byte, error) {
	return shared.NewMockResponse(http.StatusMethodNotAllowed, http.Header{}, nil), nil, nil
}

// RSQLBuilder implements transport.HTTPClient.
func (m *SitesMock) RSQLBuilder() transport.RSQLFilterBuilder {
	return nil
}

// InvalidateToken implements transport.HTTPClient.
func (m *SitesMock) InvalidateToken() error {
	return nil
}

// KeepAliveToken implements transport.HTTPClient.
func (m *SitesMock) KeepAliveToken() error {
	return nil
}

// GetLogger implements transport.HTTPClient.
func (m *SitesMock) GetLogger() *zap.Logger {
	return m.logger
}
