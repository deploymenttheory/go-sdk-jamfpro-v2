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
	errMsg     string
}

// LdapMock is a test double implementing interfaces.HTTPClient.
type LdapMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewLdapMock returns an empty mock ready for response registration.
func NewLdapMock() *LdapMock {
	return &LdapMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *LdapMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		body, _ = os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// RegisterGetLdapGroupsMock registers a successful GET /api/v1/ldap/groups response.
func (m *LdapMock) RegisterGetLdapGroupsMock() {
	m.register("GET", "/api/v1/ldap/groups", 200, "validate_list_groups.json")
}

// RegisterGetLdapServersMock registers a successful GET /api/v1/ldap/servers response.
func (m *LdapMock) RegisterGetLdapServersMock() {
	m.register("GET", "/api/v1/ldap/servers", 200, "validate_list_servers.json")
}

// RegisterGetLdapServersOnlyMock registers a successful GET /api/v1/ldap/ldap-servers response.
func (m *LdapMock) RegisterGetLdapServersOnlyMock() {
	m.register("GET", "/api/v1/ldap/ldap-servers", 200, "validate_list_servers_only.json")
}

// RegisterGetLdapGroupsErrorMock registers GET /api/v1/ldap/groups with an error for error-path testing.
func (m *LdapMock) RegisterGetLdapGroupsErrorMock() {
	body, _ := os.ReadFile(filepath.Join(mustMocksDir(), "validate_list_groups.json"))
	m.responses["GET:/api/v1/ldap/groups"] = registeredResponse{statusCode: 500, rawBody: body, errMsg: "Jamf Pro API error (500): server error"}
}

// RegisterGetLdapServersErrorMock registers GET /api/v1/ldap/servers with an error for error-path testing.
func (m *LdapMock) RegisterGetLdapServersErrorMock() {
	body, _ := os.ReadFile(filepath.Join(mustMocksDir(), "validate_list_servers.json"))
	m.responses["GET:/api/v1/ldap/servers"] = registeredResponse{statusCode: 404, rawBody: body, errMsg: "Jamf Pro API error (404): not found"}
}

// RegisterGetLdapServersOnlyErrorMock registers GET /api/v1/ldap/ldap-servers with an error for error-path testing.
func (m *LdapMock) RegisterGetLdapServersOnlyErrorMock() {
	body, _ := os.ReadFile(filepath.Join(mustMocksDir(), "validate_list_servers_only.json"))
	m.responses["GET:/api/v1/ldap/ldap-servers"] = registeredResponse{statusCode: 500, rawBody: body, errMsg: "Jamf Pro API error (500): server error"}
}

func (m *LdapMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("LdapMock: no response for %s %s", method, path)
	}
	resp := &interfaces.Response{StatusCode: r.statusCode, Status: fmt.Sprintf("%d", r.statusCode), Headers: http.Header{"Content-Type": {"application/json"}}, Body: r.rawBody}
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

func (m *LdapMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *LdapMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *LdapMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *LdapMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *LdapMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *LdapMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *LdapMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *LdapMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *LdapMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *LdapMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *LdapMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil && len(resp.Body) > 0 {
		if err := mergePage(resp.Body); err != nil {
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
	}
	return resp, nil
}
func (m *LdapMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *LdapMock) InvalidateToken() error                    { return nil }
func (m *LdapMock) KeepAliveToken() error                     { return nil }
func (m *LdapMock) GetLogger() *zap.Logger                     { return m.logger }
