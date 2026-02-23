package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
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

func (m *JCDSMock) register(method, path string, statusCode int, rawJSON string) {
	key := method + " " + path
	m.responses[key] = registeredResponse{
		statusCode: statusCode,
		rawBody:    []byte(rawJSON),
	}
}

func (m *JCDSMock) RegisterGetPackagesMock() {
	m.register("GET", "/api/v1/jcds/files", 200, `[{
		"fileName": "test-package.pkg",
		"length": 1024000,
		"md5": "abc123",
		"region": "us-east-1",
		"sha3": "def456"
	}]`)
}

func (m *JCDSMock) RegisterGetPackageURIByNameMock() {
	m.register("GET", "/api/v1/jcds/files/test-package.pkg", 200, `{
		"uri": "s3://jamf-bucket/path/test-package.pkg"
	}`)
}

func (m *JCDSMock) RegisterRenewCredentialsMock() {
	m.register("POST", "/api/v1/jcds/renew-credentials", 200, `{
		"accessKeyID": "test-access-key",
		"secretAccessKey": "test-secret-key",
		"sessionToken": "test-session-token",
		"region": "us-east-1",
		"bucketName": "jamf-bucket",
		"path": "path/",
		"uuid": "test-uuid"
	}`)
}

func (m *JCDSMock) RegisterRefreshInventoryMock() {
	m.register("POST", "/api/v1/jcds/refresh-inventory", 204, ``)
}

// Get implements interfaces.HTTPClient.
func (m *JCDSMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*interfaces.Response, error) {
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
	return &interfaces.Response{
		StatusCode: resp.statusCode,
		Headers:    http.Header{},
		Body:       resp.rawBody,
	}, nil
}

// Post implements interfaces.HTTPClient.
func (m *JCDSMock) Post(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
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
	return &interfaces.Response{
		StatusCode: resp.statusCode,
		Headers:    http.Header{},
		Body:       resp.rawBody,
	}, nil
}

// PostWithQuery implements interfaces.HTTPClient.
func (m *JCDSMock) PostWithQuery(ctx context.Context, path string, rsqlQuery map[string]string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
	return m.Post(ctx, path, body, headers, result)
}

// PostForm implements interfaces.HTTPClient.
func (m *JCDSMock) PostForm(ctx context.Context, path string, formData map[string]string, headers map[string]string, result any) (*interfaces.Response, error) {
	return m.Post(ctx, path, formData, headers, result)
}

// PostMultipart implements interfaces.HTTPClient.
func (m *JCDSMock) PostMultipart(ctx context.Context, path string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.Post(ctx, path, nil, headers, result)
}

// Put implements interfaces.HTTPClient.
func (m *JCDSMock) Put(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
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
	return &interfaces.Response{
		StatusCode: resp.statusCode,
		Headers:    http.Header{},
		Body:       resp.rawBody,
	}, nil
}

// Patch implements interfaces.HTTPClient.
func (m *JCDSMock) Patch(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
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
	return &interfaces.Response{
		StatusCode: resp.statusCode,
		Headers:    http.Header{},
		Body:       resp.rawBody,
	}, nil
}

// Delete implements interfaces.HTTPClient.
func (m *JCDSMock) Delete(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*interfaces.Response, error) {
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
	return &interfaces.Response{
		StatusCode: resp.statusCode,
		Headers:    http.Header{},
		Body:       resp.rawBody,
	}, nil
}

// DeleteWithBody implements interfaces.HTTPClient.
func (m *JCDSMock) DeleteWithBody(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
	return m.Delete(ctx, path, nil, headers, result)
}

// GetBytes implements interfaces.HTTPClient.
func (m *JCDSMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string) (*interfaces.Response, []byte, error) {
	m.LastRSQLQuery = rsqlQuery
	key := "GET " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, nil, fmt.Errorf("no mock registered for GET %s", path)
	}
	if resp.errMsg != "" {
		return nil, nil, fmt.Errorf("%s", resp.errMsg)
	}
	return &interfaces.Response{
		StatusCode: resp.statusCode,
		Headers:    http.Header{},
		Body:       resp.rawBody,
	}, resp.rawBody, nil
}

// GetPaginated implements interfaces.HTTPClient.
func (m *JCDSMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, mergePage func(pageData []byte) error) (*interfaces.Response, error) {
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
