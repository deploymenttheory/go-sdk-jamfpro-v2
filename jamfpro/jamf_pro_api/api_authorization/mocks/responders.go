package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"resty.dev/v3"

	mockhelpers "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"go.uber.org/zap"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
}

// ApiAuthorizationMock is a test double implementing client.Client.
type ApiAuthorizationMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewApiAuthorizationMock returns an empty mock ready for response registration.
func NewApiAuthorizationMock() *ApiAuthorizationMock {
	return &ApiAuthorizationMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *ApiAuthorizationMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := os.ReadFile(filepath.Join(mustGetwd(), "mocks", fixture))
		if err != nil {
			panic(fmt.Sprintf("ApiAuthorizationMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// RegisterGetV1Mock registers a successful GET response for /api/v1/auth.
func (m *ApiAuthorizationMock) RegisterGetV1Mock() {
	m.register("GET", "/api/v1/auth", 200, "validate_get.json")
}

func (m *ApiAuthorizationMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return mockhelpers.NewMockResponse(http.StatusNotFound, http.Header{}, nil), fmt.Errorf("ApiAuthorizationMock: no response for %s %s", method, path)
	}
	headers := http.Header{"Content-Type": {"application/json"}}
	resp := mockhelpers.NewMockResponse(r.statusCode, headers, r.rawBody)
	if result != nil && len(r.rawBody) > 0 {
		_ = json.Unmarshal(r.rawBody, result)
	}
	return resp, nil
}

func mustGetwd() string {
	dir, _ := os.Getwd()
	return dir
}

func (m *ApiAuthorizationMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *ApiAuthorizationMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ApiAuthorizationMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ApiAuthorizationMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ApiAuthorizationMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ApiAuthorizationMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *ApiAuthorizationMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *ApiAuthorizationMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *ApiAuthorizationMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *ApiAuthorizationMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}
func (m *ApiAuthorizationMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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
func (m *ApiAuthorizationMock) NewRequest(ctx context.Context) *client.RequestBuilder {
	return client.NewMockRequestBuilder(ctx, func(method, path string, result any) (*resty.Response, error) {
		return m.dispatch(method, path, result)
	})
}
func (m *ApiAuthorizationMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *ApiAuthorizationMock) InvalidateToken() error                { return nil }
func (m *ApiAuthorizationMock) KeepAliveToken() error                 { return nil }
func (m *ApiAuthorizationMock) GetLogger() *zap.Logger                { return m.logger }
