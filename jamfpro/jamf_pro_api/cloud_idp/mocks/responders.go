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

type CloudIdpMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewCloudIdpMock() *CloudIdpMock {
	return &CloudIdpMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *CloudIdpMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		body, _ = os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *CloudIdpMock) RegisterListMock() {
	m.register("GET", "/api/v1/cloud-idp", 200, "validate_list.json")
}

func (m *CloudIdpMock) RegisterGetByIDMock(id string) {
	m.register("GET", "/api/v1/cloud-idp/"+id, 200, "validate_get.json")
}

func (m *CloudIdpMock) RegisterExportMock() {
	m.register("POST", "/api/v1/cloud-idp/export", 200, "validate_list.json")
}

func (m *CloudIdpMock) RegisterGetHistoryByIDMock(id string) {
	m.register("GET", "/api/v1/cloud-idp/"+id+"/history", 200, "validate_history.json")
}

func (m *CloudIdpMock) RegisterAddHistoryNoteMock(id string) {
	m.register("POST", "/api/v1/cloud-idp/"+id+"/history", 201, "")
}

func (m *CloudIdpMock) RegisterTestGroupSearchMock(id string) {
	m.register("POST", "/api/v1/cloud-idp/"+id+"/test-group", 200, "validate_test_group.json")
}

func (m *CloudIdpMock) RegisterTestUserSearchMock(id string) {
	m.register("POST", "/api/v1/cloud-idp/"+id+"/test-user", 200, "validate_test_user.json")
}

func (m *CloudIdpMock) RegisterTestUserMembershipMock(id string) {
	m.register("POST", "/api/v1/cloud-idp/"+id+"/test-user-membership", 200, "validate_test_membership.json")
}

func (m *CloudIdpMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("CloudIdpMock: no response for %s %s", method, path)
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

func (m *CloudIdpMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *CloudIdpMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CloudIdpMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CloudIdpMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CloudIdpMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CloudIdpMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *CloudIdpMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *CloudIdpMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *CloudIdpMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *CloudIdpMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}
func (m *CloudIdpMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	bodyBytes := resp.Bytes()
	if mergePage != nil && len(bodyBytes) > 0 {
		var pageResponse map[string]any
		if err := json.Unmarshal(bodyBytes, &pageResponse); err != nil {
			return resp, fmt.Errorf("failed to unmarshal response: %w", err)
		}
		if results, ok := pageResponse["results"]; ok {
			resultsJSON, err := json.Marshal(results)
			if err != nil {
				return resp, fmt.Errorf("failed to marshal results: %w", err)
			}
			if err := mergePage(resultsJSON); err != nil {
				return resp, fmt.Errorf("mergePage failed: %w", err)
			}
		}
	}
	return resp, nil
}
func (m *CloudIdpMock) NewRequest(ctx context.Context) *client.RequestBuilder {
	return client.NewMockRequestBuilder(ctx, func(method, path string, result any) (*resty.Response, error) {
		return m.dispatch(method, path, result)
	})
}
func (m *CloudIdpMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *CloudIdpMock) InvalidateToken() error                    { return nil }
func (m *CloudIdpMock) KeepAliveToken() error                     { return nil }
func (m *CloudIdpMock) GetLogger() *zap.Logger                    { return m.logger }
