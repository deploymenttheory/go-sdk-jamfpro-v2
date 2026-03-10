package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"go.uber.org/zap"
	"resty.dev/v3"

	mockhelpers "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
}

// ClassicLdapMock is a test double implementing client.Client.
type ClassicLdapMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewClassicLdapMock returns an empty mock ready for response registration.
func NewClassicLdapMock() *ClassicLdapMock {
	return &ClassicLdapMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *ClassicLdapMock) register(method, path string, statusCode int, fixture string) {
	body, _ := os.ReadFile(filepath.Join(mustGetwd(), "mocks", fixture))
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// RegisterGetMappingsByIDMock registers a successful GET response for the given id.
func (m *ClassicLdapMock) RegisterGetMappingsByIDMock(id string) {
	m.register("GET", "/api/v1/classic-ldap/"+id, 200, "validate_get.json")
}

func (m *ClassicLdapMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("ClassicLdapMock: no response for %s %s", method, path)
	}
	resp := mockhelpers.NewMockResponse(r.statusCode, http.Header{"Content-Type": {"application/json"}}, r.rawBody)
	if result != nil && len(r.rawBody) > 0 {
		_ = json.Unmarshal(r.rawBody, result)
	}
	return resp, nil
}

func mustGetwd() string {
	dir, _ := os.Getwd()
	return dir
}

func (m *ClassicLdapMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *ClassicLdapMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ClassicLdapMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ClassicLdapMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ClassicLdapMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ClassicLdapMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *ClassicLdapMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *ClassicLdapMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *ClassicLdapMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *ClassicLdapMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}
func (m *ClassicLdapMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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
func (m *ClassicLdapMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *ClassicLdapMock) InvalidateToken() error                { return nil }
func (m *ClassicLdapMock) KeepAliveToken() error                 { return nil }
func (m *ClassicLdapMock) GetLogger() *zap.Logger                { return m.logger }
