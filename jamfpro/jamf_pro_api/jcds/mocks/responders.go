package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
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

// JCDSMock is a test double implementing interfaces.HTTPClient.
type JCDSMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewJCDSMock returns an empty mock ready for response registration.
func NewJCDSMock() *JCDSMock {
	return &JCDSMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *JCDSMock) RegisterMocks() {
	m.RegisterGetPackagesMock()
	m.RegisterGetPackageURIByNameMock()
	m.RegisterRenewCredentialsMock()
	m.RegisterRefreshInventoryMock()
}

func (m *JCDSMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("JCDSMock: load %q: %v", fixture, err))
		}
		body = data
	}
	key := method + " " + path
	m.responses[key] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
	}
}

func (m *JCDSMock) RegisterGetPackagesMock() {
	m.register("GET", "/api/v1/jcds/files", 200, "validate_get_packages.json")
}

func (m *JCDSMock) RegisterGetPackageURIByNameMock() {
	m.register("GET", "/api/v1/jcds/files/test-package.pkg", 200, "validate_get_package_uri.json")
}

func (m *JCDSMock) RegisterRenewCredentialsMock() {
	m.register("POST", "/api/v1/jcds/renew-credentials", 200, "validate_renew_credentials.json")
}

func (m *JCDSMock) RegisterRefreshInventoryMock() {
	m.register("POST", "/api/v1/jcds/refresh-inventory", 204, "")
}

// RegisterUploadCredentialsMock registers POST /api/v1/jcds/files for CreatePackageV1/DeletePackageV1.
func (m *JCDSMock) RegisterUploadCredentialsMock() {
	m.register("POST", "/api/v1/jcds/files", 200, "validate_upload_credentials.json")
}

// RegisterErrorMock registers a mock that returns an error for testing error paths.
func (m *JCDSMock) RegisterErrorMock(method, path string, errMsg string) {
	key := method + " " + path
	m.responses[key] = registeredResponse{
		statusCode: 500,
		rawBody:    nil,
		errMsg:     errMsg,
	}
}

// RegisterIncompleteCredentialsMock registers POST /api/v1/jcds/files with incomplete credentials.
func (m *JCDSMock) RegisterIncompleteCredentialsMock() {
	m.register("POST", "/api/v1/jcds/files", 200, "validate_upload_credentials_incomplete.json")
}

func loadMockResponse(filename string) ([]byte, error) {
	_, callerPath, _, _ := runtime.Caller(0)
	dir := filepath.Dir(callerPath)
	return os.ReadFile(filepath.Join(dir, filename))
}

// Get implements interfaces.HTTPClient.
func (m *JCDSMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	key := "GET " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for GET %s", path)
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

// Post implements interfaces.HTTPClient.
func (m *JCDSMock) Post(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	key := "POST " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for POST %s", path)
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

// PostWithQuery implements interfaces.HTTPClient.
func (m *JCDSMock) PostWithQuery(ctx context.Context, path string, rsqlQuery map[string]string, body any, headers map[string]string, result any) (*resty.Response, error) {
	return m.Post(ctx, path, body, headers, result)
}

// PostForm implements interfaces.HTTPClient.
func (m *JCDSMock) PostForm(ctx context.Context, path string, formData map[string]string, headers map[string]string, result any) (*resty.Response, error) {
	return m.Post(ctx, path, formData, headers, result)
}

// PostMultipart implements interfaces.HTTPClient.
func (m *JCDSMock) PostMultipart(ctx context.Context, path string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback interfaces.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.Post(ctx, path, nil, headers, result)
}

// Put implements interfaces.HTTPClient.
func (m *JCDSMock) Put(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	key := "PUT " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for PUT %s", path)
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

// Patch implements interfaces.HTTPClient.
func (m *JCDSMock) Patch(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	key := "PATCH " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for PATCH %s", path)
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

// Delete implements interfaces.HTTPClient.
func (m *JCDSMock) Delete(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*resty.Response, error) {
	key := "DELETE " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for DELETE %s", path)
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

// DeleteWithBody implements interfaces.HTTPClient.
func (m *JCDSMock) DeleteWithBody(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	return m.Delete(ctx, path, nil, headers, result)
}

// GetBytes implements interfaces.HTTPClient.
func (m *JCDSMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string) (*resty.Response, []byte, error) {
	m.LastRSQLQuery = rsqlQuery
	key := "GET " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, nil, fmt.Errorf("no mock registered for GET %s", path)
	}
	if resp.errMsg != "" {
		return nil, nil, fmt.Errorf("%s", resp.errMsg)
	}
	return shared.NewMockResponse(resp.statusCode, http.Header{}, resp.rawBody), resp.rawBody, nil
}

// GetPaginated implements interfaces.HTTPClient.
func (m *JCDSMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, mergePage func(pageData []byte) error) (*resty.Response, error) {
	return nil, fmt.Errorf("GetPaginated not implemented in JCDSMock")
}

// RSQLBuilder implements interfaces.HTTPClient.
func (m *JCDSMock) RSQLBuilder() interfaces.RSQLFilterBuilder {
	return nil
}

// InvalidateToken implements interfaces.HTTPClient.
func (m *JCDSMock) InvalidateToken() error {
	return nil
}

// KeepAliveToken implements interfaces.HTTPClient.
func (m *JCDSMock) KeepAliveToken() error {
	return nil
}

// GetLogger implements interfaces.HTTPClient.
func (m *JCDSMock) GetLogger() *zap.Logger {
	return m.logger
}
