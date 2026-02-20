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

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
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

func (m *CloudLdapMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{StatusCode: 404, Headers: http.Header{}, Body: nil}, fmt.Errorf("CloudLdapMock: no response for %s %s", method, path)
	}
	resp := &interfaces.Response{StatusCode: r.statusCode, Status: fmt.Sprintf("%d", r.statusCode), Headers: http.Header{"Content-Type": {"application/json"}}, Body: r.rawBody}
	if result != nil && len(r.rawBody) > 0 {
		_ = json.Unmarshal(r.rawBody, result)
	}
	return resp, nil
}

func mustMocksDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func (m *CloudLdapMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *CloudLdapMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CloudLdapMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CloudLdapMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CloudLdapMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CloudLdapMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *CloudLdapMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *CloudLdapMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *CloudLdapMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *CloudLdapMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *CloudLdapMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil && len(resp.Body) > 0 {
		_ = mergePage(resp.Body)
	}
	return resp, nil
}
func (m *CloudLdapMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *CloudLdapMock) InvalidateToken() error                     { return nil }
func (m *CloudLdapMock) KeepAliveToken() error                      { return nil }
func (m *CloudLdapMock) GetLogger() *zap.Logger                     { return m.logger }
