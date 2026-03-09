package mocks

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"go.uber.org/zap"
	"resty.dev/v3"
)

// errNoMockRegistered is returned when no mock is registered for the request.
var errNoMockRegistered = fmt.Errorf("no mock registered")

//go:embed validate_direct_idp_login_url.json
var validateDirectIdPLoginURLJSON []byte

//go:embed validate_public_key.json
var validatePublicKeyJSON []byte

//go:embed validate_public_features.json
var validatePublicFeaturesJSON []byte

//go:embed validate_redirect_url.json
var validateRedirectURLJSON []byte

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// OIDCMock is a test double implementing transport.HTTPClient.
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
	m.RegisterGetPublicFeaturesMock()
	m.RegisterGenerateCertificateMock()
	m.RegisterGetRedirectURLMock()
}

func (m *OIDCMock) register(method, path string, statusCode int, rawBody []byte) {
	key := method + " " + path
	m.responses[key] = registeredResponse{
		statusCode: statusCode,
		rawBody:    rawBody,
	}
}

// dispatch returns the registered response for the given method and path.
// When no mock is registered, it returns (nil, 0, false).
func (m *OIDCMock) dispatch(method, path string) ([]byte, int, bool) {
	key := method + " " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, 0, false
	}
	if resp.errMsg != "" {
		return nil, 0, false
	}
	return resp.rawBody, resp.statusCode, true
}

// RegisterGetDirectIdPLoginURLMock registers a successful response for GetDirectIdPLoginURLV1.
func (m *OIDCMock) RegisterGetDirectIdPLoginURLMock() {
	m.register("GET", "/api/v1/oidc/direct-idp-login-url", 200, validateDirectIdPLoginURLJSON)
}

// RegisterGetPublicKeyMock registers a successful response for GetPublicKeyV1.
func (m *OIDCMock) RegisterGetPublicKeyMock() {
	m.register("GET", "/api/v1/oidc/public-key", 200, validatePublicKeyJSON)
}

// RegisterGetPublicFeaturesMock registers a successful response for GetPublicFeaturesV1.
func (m *OIDCMock) RegisterGetPublicFeaturesMock() {
	m.register("GET", "/api/v1/oidc/public-features", 200, validatePublicFeaturesJSON)
}

// RegisterGenerateCertificateMock registers a successful response for GenerateCertificateV1.
func (m *OIDCMock) RegisterGenerateCertificateMock() {
	m.register("POST", "/api/v1/oidc/generate-certificate", 204, []byte{})
}

// RegisterGetRedirectURLMock registers a successful response for GetRedirectURLV1.
func (m *OIDCMock) RegisterGetRedirectURLMock() {
	m.register("POST", "/api/v2/oidc/dispatch", 200, validateRedirectURLJSON)
}

// Get implements transport.HTTPClient.
func (m *OIDCMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*resty.Response, error) {
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

// Post implements transport.HTTPClient.
func (m *OIDCMock) Post(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
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

// PostWithQuery implements transport.HTTPClient.
func (m *OIDCMock) PostWithQuery(ctx context.Context, path string, rsqlQuery map[string]string, body any, headers map[string]string, result any) (*resty.Response, error) {
	return m.Post(ctx, path, body, headers, result)
}

// PostForm implements transport.HTTPClient.
func (m *OIDCMock) PostForm(ctx context.Context, path string, formData map[string]string, headers map[string]string, result any) (*resty.Response, error) {
	return m.Post(ctx, path, formData, headers, result)
}

// PostMultipart implements transport.HTTPClient.
func (m *OIDCMock) PostMultipart(ctx context.Context, path string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback transport.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.Post(ctx, path, nil, headers, result)
}

// Put implements transport.HTTPClient.
func (m *OIDCMock) Put(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	rawBody, statusCode, found := m.dispatch("PUT", path)
	if !found {
		return nil, errNoMockRegistered
	}
	if result != nil && len(rawBody) > 0 {
		if err := json.Unmarshal(rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return shared.NewMockResponse(statusCode, http.Header{}, rawBody), nil
}

// Patch implements transport.HTTPClient.
func (m *OIDCMock) Patch(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	rawBody, statusCode, found := m.dispatch("PATCH", path)
	if !found {
		return nil, errNoMockRegistered
	}
	if result != nil && len(rawBody) > 0 {
		if err := json.Unmarshal(rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return shared.NewMockResponse(statusCode, http.Header{}, rawBody), nil
}

// Delete implements transport.HTTPClient.
func (m *OIDCMock) Delete(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*resty.Response, error) {
	rawBody, statusCode, found := m.dispatch("DELETE", path)
	if !found {
		return nil, errNoMockRegistered
	}
	if result != nil && len(rawBody) > 0 {
		if err := json.Unmarshal(rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return shared.NewMockResponse(statusCode, http.Header{}, rawBody), nil
}

// DeleteWithBody implements transport.HTTPClient.
func (m *OIDCMock) DeleteWithBody(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	return m.Delete(ctx, path, nil, headers, result)
}

// GetBytes implements transport.HTTPClient.
func (m *OIDCMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string) (*resty.Response, []byte, error) {
	m.LastRSQLQuery = rsqlQuery
	rawBody, statusCode, found := m.dispatch("GET", path)
	if !found {
		return nil, nil, errNoMockRegistered
	}
	return shared.NewMockResponse(statusCode, http.Header{}, rawBody), rawBody, nil
}

// GetPaginated implements transport.HTTPClient.
func (m *OIDCMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, mergePage func(pageData []byte) error) (*resty.Response, error) {
	return nil, fmt.Errorf("GetPaginated not implemented in OIDCMock")
}

// RSQLBuilder implements transport.HTTPClient.
func (m *OIDCMock) RSQLBuilder() transport.RSQLFilterBuilder {
	return nil
}

// InvalidateToken implements transport.HTTPClient.
func (m *OIDCMock) InvalidateToken() error {
	return nil
}

// KeepAliveToken implements transport.HTTPClient.
func (m *OIDCMock) KeepAliveToken() error {
	return nil
}

// GetLogger implements transport.HTTPClient.
func (m *OIDCMock) GetLogger() *zap.Logger {
	return m.logger
}
