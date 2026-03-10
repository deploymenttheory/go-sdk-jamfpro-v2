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

type VenafiMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewVenafiMock() *VenafiMock {
	return &VenafiMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *VenafiMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		body, _ = os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *VenafiMock) RegisterCreateMock() {
	m.register("POST", "/api/v1/pki/venafi", 201, "validate_create.json")
}

func (m *VenafiMock) RegisterGetByIDMock(id string) {
	m.register("GET", "/api/v1/pki/venafi/"+id, 200, "validate_get.json")
}

func (m *VenafiMock) RegisterUpdateByIDMock(id string) {
	m.register("PATCH", "/api/v1/pki/venafi/"+id, 200, "validate_get.json")
}

func (m *VenafiMock) RegisterDeleteByIDMock(id string) {
	m.register("DELETE", "/api/v1/pki/venafi/"+id, 204, "")
}

func (m *VenafiMock) RegisterGetConnectionStatusMock(id string) {
	m.register("GET", "/api/v1/pki/venafi/"+id+"/connection-status", 200, "validate_connection_status.json")
}

func (m *VenafiMock) RegisterGetDependentProfilesMock(id string) {
	m.register("GET", "/api/v1/pki/venafi/"+id+"/dependent-profiles", 200, "validate_dependent_profiles.json")
}

func (m *VenafiMock) RegisterGetHistoryMock(id string) {
	m.register("GET", "/api/v1/pki/venafi/"+id+"/history", 200, "validate_history.json")
}

func (m *VenafiMock) RegisterAddHistoryNoteMock(id string) {
	m.register("POST", "/api/v1/pki/venafi/"+id+"/history", 201, "validate_history_note.json")
}

func (m *VenafiMock) RegisterGetJamfPublicKeyMock(id string) {
	m.register("GET", "/api/v1/pki/venafi/"+id+"/jamf-public-key", 200, "validate_jamf_public_key.pem")
}

func (m *VenafiMock) RegisterGetProxyTrustStoreMock(id string) {
	m.register("GET", "/api/v1/pki/venafi/"+id+"/proxy-trust-store", 200, "validate_jamf_public_key.pem")
}

func (m *VenafiMock) RegisterRegenerateJamfPublicKeyMock(id string) {
	m.register("POST", "/api/v1/pki/venafi/"+id+"/jamf-public-key/regenerate", 200, "")
}

func (m *VenafiMock) RegisterUploadProxyTrustStoreMock(id string) {
	m.register("POST", "/api/v1/pki/venafi/"+id+"/proxy-trust-store", 200, "")
}

func (m *VenafiMock) RegisterDeleteProxyTrustStoreMock(id string) {
	m.register("DELETE", "/api/v1/pki/venafi/"+id+"/proxy-trust-store", 204, "")
}

func (m *VenafiMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("VenafiMock: no response for %s %s", method, path)
	}
	headers := http.Header{"Content-Type": {"application/json"}}
	resp := mockhelpers.NewMockResponse(r.statusCode, headers, r.rawBody)
	if result != nil && len(r.rawBody) > 0 {
		_ = json.Unmarshal(r.rawBody, result)
	}
	return resp, nil
}

func mustMocksDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func (m *VenafiMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *VenafiMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *VenafiMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *VenafiMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *VenafiMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *VenafiMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *VenafiMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *VenafiMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *VenafiMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *VenafiMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}
func (m *VenafiMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	bodyBytes := resp.Bytes()
	if mergePage != nil && len(bodyBytes) > 0 {
		var page struct {
			Results json.RawMessage `json:"results"`
		}
		if err := json.Unmarshal(bodyBytes, &page); err != nil {
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
		if err := mergePage(page.Results); err != nil {
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
	}
	return resp, nil
}
func (m *VenafiMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *VenafiMock) InvalidateToken() error                    { return nil }
func (m *VenafiMock) KeepAliveToken() error                     { return nil }
func (m *VenafiMock) GetLogger() *zap.Logger                    { return m.logger }
