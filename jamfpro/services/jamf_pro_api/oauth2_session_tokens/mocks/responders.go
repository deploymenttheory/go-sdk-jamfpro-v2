package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"resty.dev/v3"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// OAuth2SessionTokensMock implements interfaces.HTTPClient for tests.
type OAuth2SessionTokensMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewOAuth2SessionTokensMock() *OAuth2SessionTokensMock {
	return &OAuth2SessionTokensMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *OAuth2SessionTokensMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("OAuth2SessionTokensMock: load %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *OAuth2SessionTokensMock) RegisterMocks() {
	m.register("GET", "/api/v1/oauth2/session-tokens", 200, "validate_get.json")
}

func (m *OAuth2SessionTokensMock) RegisterGetErrorMock() {
	m.responses["GET:/api/v1/oauth2/session-tokens"] = registeredResponse{
		statusCode: 401,
		rawBody:    []byte(`{"code":"UNAUTHORIZED","message":"unauthorized"}`),
		errMsg:     "Jamf Pro API error (401): unauthorized",
	}
}

func (m *OAuth2SessionTokensMock) Get(ctx context.Context, path string, q map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *OAuth2SessionTokensMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *OAuth2SessionTokensMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *OAuth2SessionTokensMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *OAuth2SessionTokensMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *OAuth2SessionTokensMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *OAuth2SessionTokensMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *OAuth2SessionTokensMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *OAuth2SessionTokensMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *OAuth2SessionTokensMock) GetBytes(ctx context.Context, path string, q map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}
func (m *OAuth2SessionTokensMock) GetPaginated(ctx context.Context, path string, q map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	bodyBytes := resp.Bytes()
	if mergePage != nil && len(bodyBytes) > 0 {
		if err := mergePage(bodyBytes); err != nil {
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
	}
	return resp, nil
}
func (m *OAuth2SessionTokensMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *OAuth2SessionTokensMock) InvalidateToken() error                    { return nil }
func (m *OAuth2SessionTokensMock) KeepAliveToken() error                     { return nil }
func (m *OAuth2SessionTokensMock) GetLogger() *zap.Logger                    { return m.logger }

func (m *OAuth2SessionTokensMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return shared.NewMockResponse(404, http.Header{}, nil), fmt.Errorf("OAuth2SessionTokensMock: no response for %s %s", method, path)
	}
	headers := http.Header{"Content-Type": {"application/json"}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)
	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}
	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, err
		}
	}
	return resp, nil
}

func loadMockResponse(filename string) ([]byte, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return os.ReadFile(filepath.Join(dir, "mocks", filename))
}
