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

// OIDCMock is a test double implementing interfaces.HTTPClient.
type OIDCMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewOIDCMock returns an empty mock ready for response registration.
func NewOIDCMock() *OIDCMock {
	return &OIDCMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *OIDCMock) RegisterMocks() {
	m.RegisterGetDirectIdPLoginURLMock()
	m.RegisterGetPublicKeyMock()
	m.RegisterGenerateCertificateMock()
	m.RegisterGetRedirectURLMock()
}

func (m *OIDCMock) register(method, path string, statusCode int, rawJSON string) {
	key := method + " " + path
	m.responses[key] = registeredResponse{
		statusCode: statusCode,
		rawBody:    []byte(rawJSON),
	}
}

// RegisterGetDirectIdPLoginURLMock registers a successful response for GetDirectIdPLoginURLV1.
func (m *OIDCMock) RegisterGetDirectIdPLoginURLMock() {
	m.register("GET", "/api/v1/oidc/direct-idp-login-url", 200, `{
		"url": "https://idp.example.com/authorize?client_id=jamfpro"
	}`)
}

// RegisterGetPublicKeyMock registers a successful response for GetPublicKeyV1.
func (m *OIDCMock) RegisterGetPublicKeyMock() {
	m.register("GET", "/api/v1/oidc/public-key", 200, `{
		"keys": [{
			"kty": "RSA",
			"e": "AQAB",
			"use": "sig",
			"kid": "test-key-id",
			"alg": "RS256",
			"iat": 1609459200,
			"n": "xGOr-H7A..."
		}]
	}`)
}

// RegisterGenerateCertificateMock registers a successful response for GenerateCertificateV1.
func (m *OIDCMock) RegisterGenerateCertificateMock() {
	m.register("POST", "/api/v1/oidc/generate-certificate", 204, ``)
}

// RegisterGetRedirectURLMock registers a successful response for GetRedirectURLV1.
func (m *OIDCMock) RegisterGetRedirectURLMock() {
	m.register("POST", "/api/v1/oidc/dispatch", 200, `{
		"redirectUrl": "https://idp.example.com/login?redirect=https://jamf.example.com"
	}`)
}

// Get implements interfaces.HTTPClient.
func (m *OIDCMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*interfaces.Response, error) {
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
func (m *OIDCMock) Post(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
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
func (m *OIDCMock) PostWithQuery(ctx context.Context, path string, rsqlQuery map[string]string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
	return m.Post(ctx, path, body, headers, result)
}

// PostForm implements interfaces.HTTPClient.
func (m *OIDCMock) PostForm(ctx context.Context, path string, formData map[string]string, headers map[string]string, result any) (*interfaces.Response, error) {
	return m.Post(ctx, path, formData, headers, result)
}

// PostMultipart implements interfaces.HTTPClient.
func (m *OIDCMock) PostMultipart(ctx context.Context, path string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.Post(ctx, path, nil, headers, result)
}

// Put implements interfaces.HTTPClient.
func (m *OIDCMock) Put(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
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
func (m *OIDCMock) Patch(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
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
func (m *OIDCMock) Delete(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*interfaces.Response, error) {
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
func (m *OIDCMock) DeleteWithBody(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
	return m.Delete(ctx, path, nil, headers, result)
}

// GetBytes implements interfaces.HTTPClient.
func (m *OIDCMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string) (*interfaces.Response, []byte, error) {
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
func (m *OIDCMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, mergePage func(pageData []byte) error) (*interfaces.Response, error) {
	return nil, fmt.Errorf("GetPaginated not implemented in OIDCMock")
}

// RSQLBuilder implements interfaces.HTTPClient.
func (m *OIDCMock) RSQLBuilder() interfaces.RSQLFilterBuilder {
	return nil
}

// InvalidateToken implements interfaces.HTTPClient.
func (m *OIDCMock) InvalidateToken() error {
	return nil
}

// KeepAliveToken implements interfaces.HTTPClient.
func (m *OIDCMock) KeepAliveToken() error {
	return nil
}

// GetLogger implements interfaces.HTTPClient.
func (m *OIDCMock) GetLogger() *zap.Logger {
	return m.logger
}
