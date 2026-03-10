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
	errMsg     string
}

type CloudLdapMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewCloudLdapMock() *CloudLdapMock {
	return &CloudLdapMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *CloudLdapMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		body, _ = os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *CloudLdapMock) RegisterGetDefaultMappingsMock(providerName string) {
	m.register("GET", "/api/v2/cloud-ldaps/defaults/"+providerName+"/mappings", 200, "validate_default_mappings.json")
}

func (m *CloudLdapMock) RegisterGetDefaultServerConfigurationMock(providerName string) {
	m.register("GET", "/api/v2/cloud-ldaps/defaults/"+providerName+"/server-configuration", 200, "validate_default_server.json")
}

func (m *CloudLdapMock) RegisterCreateMock() {
	m.register("POST", "/api/v2/cloud-ldaps", 201, "validate_create.json")
}

func (m *CloudLdapMock) RegisterGetByIDMock(id string) {
	m.register("GET", "/api/v2/cloud-ldaps/"+id, 200, "validate_get.json")
}

func (m *CloudLdapMock) RegisterUpdateByIDMock(id string) {
	m.register("PUT", "/api/v2/cloud-ldaps/"+id, 200, "validate_get.json")
}

func (m *CloudLdapMock) RegisterDeleteByIDMock(id string) {
	m.register("DELETE", "/api/v2/cloud-ldaps/"+id, 204, "")
}

func (m *CloudLdapMock) RegisterGetBindConnectionPoolStatsMock(id string) {
	m.register("GET", "/api/v2/cloud-ldaps/"+id+"/connection/bind", 200, "validate_connection_pool.json")
}

func (m *CloudLdapMock) RegisterGetSearchConnectionPoolStatsMock(id string) {
	m.register("GET", "/api/v2/cloud-ldaps/"+id+"/connection/search", 200, "validate_connection_pool.json")
}

func (m *CloudLdapMock) RegisterTestConnectionMock(id string) {
	m.register("GET", "/api/v2/cloud-ldaps/"+id+"/connection/status", 200, "validate_connection_status.json")
}

func (m *CloudLdapMock) RegisterGetMappingsByIDMock(id string) {
	m.register("GET", "/api/v2/cloud-ldaps/"+id+"/mappings", 200, "validate_default_mappings.json")
}

func (m *CloudLdapMock) RegisterUpdateMappingsByIDMock(id string) {
	m.register("PUT", "/api/v2/cloud-ldaps/"+id+"/mappings", 200, "validate_default_mappings.json")
}

// RegisterGetByIDErrorMock registers GET /api/v2/cloud-ldaps/{id} with an error response for API error path testing.
func (m *CloudLdapMock) RegisterGetByIDErrorMock(id string) {
	body, _ := os.ReadFile(filepath.Join(mustMocksDir(), "validate_get.json"))
	m.responses["GET:/api/v2/cloud-ldaps/"+id] = registeredResponse{statusCode: 404, rawBody: body, errMsg: "Jamf Pro API error (404): not found"}
}

// RegisterCreateErrorMock registers POST /api/v2/cloud-ldaps with an error response for API error path testing.
func (m *CloudLdapMock) RegisterCreateErrorMock() {
	body, _ := os.ReadFile(filepath.Join(mustMocksDir(), "validate_create.json"))
	m.responses["POST:/api/v2/cloud-ldaps"] = registeredResponse{statusCode: 500, rawBody: body, errMsg: "Jamf Pro API error (500): server error"}
}

func (m *CloudLdapMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("CloudLdapMock: no response for %s %s", method, path)
	}
	resp := mockhelpers.NewMockResponse(r.statusCode, http.Header{"Content-Type": {"application/json"}}, r.rawBody)
	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}
	if result != nil && len(r.rawBody) > 0 {
		_ = json.Unmarshal(r.rawBody, result)
	}
	return resp, nil
}

func mustMocksDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func (m *CloudLdapMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *CloudLdapMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CloudLdapMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CloudLdapMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CloudLdapMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CloudLdapMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *CloudLdapMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *CloudLdapMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *CloudLdapMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *CloudLdapMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}
func (m *CloudLdapMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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
func (m *CloudLdapMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *CloudLdapMock) InvalidateToken() error                    { return nil }
func (m *CloudLdapMock) KeepAliveToken() error                     { return nil }
func (m *CloudLdapMock) GetLogger() *zap.Logger                    { return m.logger }
