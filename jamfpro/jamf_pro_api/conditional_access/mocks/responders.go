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

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"go.uber.org/zap"
	"resty.dev/v3"

	mockhelpers "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
}

type ConditionalAccessMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewConditionalAccessMock() *ConditionalAccessMock {
	return &ConditionalAccessMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *ConditionalAccessMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		body, _ = os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *ConditionalAccessMock) RegisterGetDeviceComplianceFeatureToggleMock() {
	m.register("GET", "/api/v1/conditional-access/device-compliance/feature-toggle", 200, "validate_get.json")
}

func (m *ConditionalAccessMock) RegisterGetDeviceComplianceInformationComputerMock() {
	m.register("GET", "/api/v1/conditional-access/device-compliance-information/computer/1", 200, "validate_compliance_computer.json")
}

func (m *ConditionalAccessMock) RegisterGetDeviceComplianceInformationMobileMock() {
	m.register("GET", "/api/v1/conditional-access/device-compliance-information/mobile/1", 200, "validate_compliance_mobile.json")
}

func (m *ConditionalAccessMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return mockhelpers.NewMockResponse(404, http.Header{}, nil), fmt.Errorf("ConditionalAccessMock: no response for %s %s", method, path)
	}
	resp := mockhelpers.NewMockResponse(r.statusCode, http.Header{"Content-Type": {"application/json"}}, r.rawBody)
	if result != nil && len(r.rawBody) > 0 {
		_ = json.Unmarshal(r.rawBody, result)
	}
	return resp, nil
}

func mustMocksDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func (m *ConditionalAccessMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *ConditionalAccessMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ConditionalAccessMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ConditionalAccessMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ConditionalAccessMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ConditionalAccessMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *ConditionalAccessMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *ConditionalAccessMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *ConditionalAccessMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *ConditionalAccessMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}
func (m *ConditionalAccessMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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
func (m *ConditionalAccessMock) NewRequest(ctx context.Context) *client.RequestBuilder {
	return client.NewMockRequestBuilder(ctx, func(method, path string, result any) (*resty.Response, error) {
		return m.dispatch(method, path, result)
	})
}
func (m *ConditionalAccessMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *ConditionalAccessMock) InvalidateToken() error                     { return nil }
func (m *ConditionalAccessMock) KeepAliveToken() error                      { return nil }
func (m *ConditionalAccessMock) GetLogger() *zap.Logger                     { return m.logger }
