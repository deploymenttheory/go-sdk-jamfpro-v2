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

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"resty.dev/v3"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
}

type BrandingMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewBrandingMock() *BrandingMock {
	return &BrandingMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *BrandingMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		_, file, _, _ := runtime.Caller(0)
		mocksDir := filepath.Dir(file)
		body, _ = os.ReadFile(filepath.Join(mocksDir, fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *BrandingMock) RegisterMocks() {
	m.register("GET", "/api/v1/branding-images/download/test-id", 200, "validate_download.bin")
}

func (m *BrandingMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *BrandingMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *BrandingMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *BrandingMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *BrandingMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *BrandingMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *BrandingMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *BrandingMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *BrandingMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

// GetBytes implements interfaces.HTTPClient.GetBytes.
func (m *BrandingMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}
func (m *BrandingMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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
func (m *BrandingMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *BrandingMock) InvalidateToken() error                    { return nil }
func (m *BrandingMock) KeepAliveToken() error                     { return nil }
func (m *BrandingMock) GetLogger() *zap.Logger                    { return m.logger }

func (m *BrandingMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return shared.NewMockResponse(http.StatusNotFound, http.Header{}, nil), fmt.Errorf("no mock for %s %s", method, path)
	}
	headers := http.Header{"Content-Type": {"image/png"}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)
	if result != nil && len(r.rawBody) > 0 {
		_ = json.Unmarshal(r.rawBody, result)
	}
	return resp, nil
}
