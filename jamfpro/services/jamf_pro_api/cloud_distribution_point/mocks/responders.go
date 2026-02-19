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

type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// CloudDistributionPointMock implements interfaces.HTTPClient.
type CloudDistributionPointMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewCloudDistributionPointMock() *CloudDistributionPointMock {
	return &CloudDistributionPointMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func (m *CloudDistributionPointMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(err)
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *CloudDistributionPointMock) RegisterMocks() {
	m.register("GET", "/api/v1/cloud-distribution-point", 200, "validate_get.json")
	m.register("POST", "/api/v1/cloud-distribution-point", 200, "validate_get.json")
	m.register("PATCH", "/api/v1/cloud-distribution-point", 200, "validate_get.json")
	m.register("DELETE", "/api/v1/cloud-distribution-point", 204, "")
	m.register("GET", "/api/v1/cloud-distribution-point/upload-capability", 200, "validate_upload_capability.json")
	m.register("GET", "/api/v1/cloud-distribution-point/test-connection", 200, "validate_test_connection.json")
}

func (m *CloudDistributionPointMock) Get(ctx context.Context, path string, q map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *CloudDistributionPointMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CloudDistributionPointMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CloudDistributionPointMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CloudDistributionPointMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CloudDistributionPointMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *CloudDistributionPointMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *CloudDistributionPointMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *CloudDistributionPointMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *CloudDistributionPointMock) GetBytes(ctx context.Context, path string, q map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *CloudDistributionPointMock) GetPaginated(ctx context.Context, path string, q map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	return m.dispatch("GET", path, nil)
}
func (m *CloudDistributionPointMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *CloudDistributionPointMock) InvalidateToken() error                    { return nil }
func (m *CloudDistributionPointMock) KeepAliveToken() error                     { return nil }
func (m *CloudDistributionPointMock) GetLogger() *zap.Logger                    { return m.logger }

func (m *CloudDistributionPointMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{StatusCode: 404, Body: nil}, fmt.Errorf("no mock for %s %s", method, path)
	}
	resp := &interfaces.Response{StatusCode: r.statusCode, Headers: http.Header{"Content-Type": {"application/json"}}, Body: r.rawBody}
	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}
	if result != nil && len(r.rawBody) > 0 {
		_ = json.Unmarshal(r.rawBody, result)
	}
	return resp, nil
}

func loadMockResponse(filename string) ([]byte, error) {
	dir, _ := os.Getwd()
	return os.ReadFile(filepath.Join(dir, "mocks", filename))
}
