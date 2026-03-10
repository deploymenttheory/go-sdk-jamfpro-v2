package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"resty.dev/v3"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"go.uber.org/zap"
)

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// AccountsMock is a test double implementing client.Client.
type AccountsMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewAccountsMock returns an empty mock ready for response registration.
func NewAccountsMock() *AccountsMock {
	return &AccountsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *AccountsMock) RegisterMocks() {
	m.RegisterListAccountsMock()
	m.RegisterGetAccountMock()
	m.RegisterCreateAccountMock()
	m.RegisterDeleteAccountMock()
}

func (m *AccountsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("AccountsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// RegisterListAccountsMock registers the list accounts response.
func (m *AccountsMock) RegisterListAccountsMock() {
	m.register("GET", "/api/v1/accounts", 200, "validate_list.json")
}

// RegisterGetAccountMock registers the get account by ID response.
func (m *AccountsMock) RegisterGetAccountMock() {
	m.register("GET", "/api/v1/accounts/1", 200, "validate_get.json")
}

// RegisterCreateAccountMock registers the create account response.
func (m *AccountsMock) RegisterCreateAccountMock() {
	m.register("POST", "/api/v1/accounts", 201, "validate_create.json")
}

// RegisterDeleteAccountMock registers the delete account response.
func (m *AccountsMock) RegisterDeleteAccountMock() {
	m.register("DELETE", "/api/v1/accounts/1", 204, "")
}

// loadMockResponse loads a JSON fixture from the mocks directory.
func loadMockResponse(filename string) ([]byte, error) {
	wd, _ := os.Getwd()
	return os.ReadFile(filepath.Join(wd, "mocks", filename))
}

// Get implements client.Client.Get.
func (m *AccountsMock) Get(ctx context.Context, path string, query map[string]string, headers map[string]string, out any) (*resty.Response, error) {
	m.LastRSQLQuery = query
	key := "GET:" + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for GET %s", path)
	}

	if resp.errMsg != "" {
		return shared.NewMockResponse(resp.statusCode, http.Header{}, nil), fmt.Errorf("%s", resp.errMsg)
	}

	if out != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, out); err != nil {
			return nil, fmt.Errorf("failed to unmarshal mock response: %w", err)
		}
	}

	return shared.NewMockResponse(resp.statusCode, http.Header{}, nil), nil
}

// GetPaginated implements client.Client.GetPaginated.
func (m *AccountsMock) GetPaginated(ctx context.Context, path string, query map[string]string, headers map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	m.LastRSQLQuery = query
	key := "GET:" + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for GET %s", path)
	}

	if resp.errMsg != "" {
		return shared.NewMockResponse(resp.statusCode, http.Header{}, nil), fmt.Errorf("%s", resp.errMsg)
	}

	if mergePage != nil && len(resp.rawBody) > 0 {
		var page struct {
			Results json.RawMessage `json:"results"`
		}
		if err := json.Unmarshal(resp.rawBody, &page); err != nil {
			return nil, fmt.Errorf("failed to unmarshal paginated response: %w", err)
		}
		if err := mergePage(page.Results); err != nil {
			return nil, fmt.Errorf("mergePage failed: %w", err)
		}
	}

	return shared.NewMockResponse(resp.statusCode, http.Header{}, nil), nil
}

// Post implements client.Client.Post.
func (m *AccountsMock) Post(ctx context.Context, path string, body any, headers map[string]string, out any) (*resty.Response, error) {
	key := "POST:" + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for POST %s", path)
	}

	if resp.errMsg != "" {
		return shared.NewMockResponse(resp.statusCode, http.Header{}, nil), fmt.Errorf("%s", resp.errMsg)
	}

	if out != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, out); err != nil {
			return nil, fmt.Errorf("failed to unmarshal mock response: %w", err)
		}
	}

	return shared.NewMockResponse(resp.statusCode, http.Header{}, nil), nil
}

// Delete implements client.Client.Delete.
func (m *AccountsMock) Delete(ctx context.Context, path string, query map[string]string, headers map[string]string, out any) (*resty.Response, error) {
	key := "DELETE:" + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for DELETE %s", path)
	}

	if resp.errMsg != "" {
		return shared.NewMockResponse(resp.statusCode, http.Header{}, nil), fmt.Errorf("%s", resp.errMsg)
	}

	return shared.NewMockResponse(resp.statusCode, http.Header{}, nil), nil
}

// Put implements client.Client.Put (unused in accounts).
func (m *AccountsMock) Put(ctx context.Context, path string, body any, headers map[string]string, out any) (*resty.Response, error) {
	return nil, fmt.Errorf("Put not implemented in AccountsMock")
}

// Patch implements client.Client.Patch (unused in accounts).
func (m *AccountsMock) Patch(ctx context.Context, path string, body any, headers map[string]string, out any) (*resty.Response, error) {
	return nil, fmt.Errorf("Patch not implemented in AccountsMock")
}

// DownloadFile implements client.Client.DownloadFile (unused in accounts).
func (m *AccountsMock) DownloadFile(ctx context.Context, url string) (io.ReadCloser, *http.Response, error) {
	return nil, nil, fmt.Errorf("DownloadFile not implemented in AccountsMock")
}

// SetLogger implements client.Client.SetLogger.
func (m *AccountsMock) SetLogger(logger *zap.Logger) {
	m.logger = logger
}

// GetLogger implements client.Client.GetLogger.
func (m *AccountsMock) GetLogger() *zap.Logger {
	return m.logger
}

// DeleteWithBody implements client.Client.DeleteWithBody.
func (m *AccountsMock) DeleteWithBody(ctx context.Context, path string, body any, headers map[string]string, out any) (*resty.Response, error) {
	key := "DELETE:" + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for DELETE %s", path)
	}

	if resp.errMsg != "" {
		return shared.NewMockResponse(resp.statusCode, http.Header{}, nil), fmt.Errorf("%s", resp.errMsg)
	}

	return shared.NewMockResponse(resp.statusCode, http.Header{}, nil), nil
}

// PostWithQuery implements client.Client.PostWithQuery.
func (m *AccountsMock) PostWithQuery(ctx context.Context, path string, rsqlQuery map[string]string, body any, headers map[string]string, out any) (*resty.Response, error) {
	return nil, fmt.Errorf("PostWithQuery not implemented in AccountsMock")
}

// PostForm implements client.Client.PostForm.
func (m *AccountsMock) PostForm(ctx context.Context, path string, formData map[string]string, headers map[string]string, out any) (*resty.Response, error) {
	return nil, fmt.Errorf("PostForm not implemented in AccountsMock")
}

// PostMultipart implements client.Client.PostMultipart.
func (m *AccountsMock) PostMultipart(ctx context.Context, path string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback client.MultipartProgressCallback, out any) (*resty.Response, error) {
	return nil, fmt.Errorf("PostMultipart not implemented in AccountsMock")
}

// GetBytes implements client.Client.GetBytes.
func (m *AccountsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string) (*resty.Response, []byte, error) {
	return nil, nil, fmt.Errorf("GetBytes not implemented in AccountsMock")
}

// RSQLBuilder implements client.Client.RSQLBuilder.
func (m *AccountsMock) RSQLBuilder() client.RSQLFilterBuilder {
	return nil
}

// InvalidateToken implements client.Client.InvalidateToken.
func (m *AccountsMock) InvalidateToken() error {
	return nil
}

// KeepAliveToken implements client.Client.KeepAliveToken.
func (m *AccountsMock) KeepAliveToken() error {
	return nil
}
