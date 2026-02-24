package mocks

import (
	"context"
	"encoding/json"
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

// AccountsMock is a test double implementing interfaces.HTTPClient.
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
	m.register("GET", "/api/v1/accounts", 200, "list_success")
}

// RegisterGetAccountMock registers the get account by ID response.
func (m *AccountsMock) RegisterGetAccountMock() {
	m.register("GET", "/api/v1/accounts/1", 200, "get_by_id_success")
}

// RegisterCreateAccountMock registers the create account response.
func (m *AccountsMock) RegisterCreateAccountMock() {
	m.register("POST", "/api/v1/accounts", 201, "create_success")
}

// RegisterDeleteAccountMock registers the delete account response.
func (m *AccountsMock) RegisterDeleteAccountMock() {
	m.register("DELETE", "/api/v1/accounts/1", 204, "")
}

// loadMockResponse loads a JSON fixture from the mocks/mock_responses.json file.
func loadMockResponse(key string) ([]byte, error) {
	wd, _ := os.Getwd()
	jsonPath := filepath.Join(wd, "mocks", "mock_responses.json")

	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read mock_responses.json: %w", err)
	}

	var allResponses map[string]json.RawMessage
	if err := json.Unmarshal(data, &allResponses); err != nil {
		return nil, fmt.Errorf("failed to unmarshal mock_responses.json: %w", err)
	}

	response, ok := allResponses[key]
	if !ok {
		return nil, fmt.Errorf("mock response key %q not found", key)
	}

	return response, nil
}

// Get implements interfaces.HTTPClient.Get.
func (m *AccountsMock) Get(ctx context.Context, path string, query map[string]string, headers map[string]string, out any) (*interfaces.Response, error) {
	m.LastRSQLQuery = query
	key := "GET:" + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for GET %s", path)
	}

	if resp.errMsg != "" {
		return &interfaces.Response{StatusCode: resp.statusCode}, fmt.Errorf(resp.errMsg)
	}

	if out != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, out); err != nil {
			return nil, fmt.Errorf("failed to unmarshal mock response: %w", err)
		}
	}

	return &interfaces.Response{StatusCode: resp.statusCode}, nil
}

// GetPaginated implements interfaces.HTTPClient.GetPaginated.
func (m *AccountsMock) GetPaginated(ctx context.Context, path string, query map[string]string, headers map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	m.LastRSQLQuery = query
	key := "GET:" + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for GET %s", path)
	}

	if resp.errMsg != "" {
		return &interfaces.Response{StatusCode: resp.statusCode}, fmt.Errorf(resp.errMsg)
	}

	if len(resp.rawBody) > 0 {
		if err := mergePage(resp.rawBody); err != nil {
			return nil, fmt.Errorf("mergePage failed: %w", err)
		}
	}

	return &interfaces.Response{StatusCode: resp.statusCode}, nil
}

// Post implements interfaces.HTTPClient.Post.
func (m *AccountsMock) Post(ctx context.Context, path string, body any, headers map[string]string, out any) (*interfaces.Response, error) {
	key := "POST:" + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for POST %s", path)
	}

	if resp.errMsg != "" {
		return &interfaces.Response{StatusCode: resp.statusCode}, fmt.Errorf(resp.errMsg)
	}

	if out != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, out); err != nil {
			return nil, fmt.Errorf("failed to unmarshal mock response: %w", err)
		}
	}

	return &interfaces.Response{StatusCode: resp.statusCode}, nil
}

// Delete implements interfaces.HTTPClient.Delete.
func (m *AccountsMock) Delete(ctx context.Context, path string, query map[string]string, headers map[string]string, out any) (*interfaces.Response, error) {
	key := "DELETE:" + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for DELETE %s", path)
	}

	if resp.errMsg != "" {
		return &interfaces.Response{StatusCode: resp.statusCode}, fmt.Errorf(resp.errMsg)
	}

	return &interfaces.Response{StatusCode: resp.statusCode}, nil
}

// Put implements interfaces.HTTPClient.Put (unused in accounts).
func (m *AccountsMock) Put(ctx context.Context, path string, body any, headers map[string]string, out any) (*interfaces.Response, error) {
	return nil, fmt.Errorf("Put not implemented in AccountsMock")
}

// Patch implements interfaces.HTTPClient.Patch (unused in accounts).
func (m *AccountsMock) Patch(ctx context.Context, path string, body any, headers map[string]string, out any) (*interfaces.Response, error) {
	return nil, fmt.Errorf("Patch not implemented in AccountsMock")
}

// DownloadFile implements interfaces.HTTPClient.DownloadFile (unused in accounts).
func (m *AccountsMock) DownloadFile(ctx context.Context, url string) (io.ReadCloser, *http.Response, error) {
	return nil, nil, fmt.Errorf("DownloadFile not implemented in AccountsMock")
}

// SetLogger implements interfaces.HTTPClient.SetLogger.
func (m *AccountsMock) SetLogger(logger *zap.Logger) {
	m.logger = logger
}

// GetLogger implements interfaces.HTTPClient.GetLogger.
func (m *AccountsMock) GetLogger() *zap.Logger {
	return m.logger
}

// DeleteWithBody implements interfaces.HTTPClient.DeleteWithBody.
func (m *AccountsMock) DeleteWithBody(ctx context.Context, path string, body any, headers map[string]string, out any) (*interfaces.Response, error) {
	key := "DELETE:" + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for DELETE %s", path)
	}

	if resp.errMsg != "" {
		return &interfaces.Response{StatusCode: resp.statusCode}, fmt.Errorf("%s", resp.errMsg)
	}

	return &interfaces.Response{StatusCode: resp.statusCode}, nil
}

// PostWithQuery implements interfaces.HTTPClient.PostWithQuery.
func (m *AccountsMock) PostWithQuery(ctx context.Context, path string, rsqlQuery map[string]string, body any, headers map[string]string, out any) (*interfaces.Response, error) {
	return nil, fmt.Errorf("PostWithQuery not implemented in AccountsMock")
}

// PostForm implements interfaces.HTTPClient.PostForm.
func (m *AccountsMock) PostForm(ctx context.Context, path string, formData map[string]string, headers map[string]string, out any) (*interfaces.Response, error) {
	return nil, fmt.Errorf("PostForm not implemented in AccountsMock")
}

// PostMultipart implements interfaces.HTTPClient.PostMultipart.
func (m *AccountsMock) PostMultipart(ctx context.Context, path string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback interfaces.MultipartProgressCallback, out any) (*interfaces.Response, error) {
	return nil, fmt.Errorf("PostMultipart not implemented in AccountsMock")
}

// GetBytes implements interfaces.HTTPClient.GetBytes.
func (m *AccountsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string) (*interfaces.Response, []byte, error) {
	return nil, nil, fmt.Errorf("GetBytes not implemented in AccountsMock")
}

// RSQLBuilder implements interfaces.HTTPClient.RSQLBuilder.
func (m *AccountsMock) RSQLBuilder() interfaces.RSQLFilterBuilder {
	return nil
}

// InvalidateToken implements interfaces.HTTPClient.InvalidateToken.
func (m *AccountsMock) InvalidateToken() error {
	return nil
}

// KeepAliveToken implements interfaces.HTTPClient.KeepAliveToken.
func (m *AccountsMock) KeepAliveToken() error {
	return nil
}
