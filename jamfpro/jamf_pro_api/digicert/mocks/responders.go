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
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"go.uber.org/zap"
	"resty.dev/v3"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

type DigicertMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewDigicertMock() *DigicertMock {
	return &DigicertMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *DigicertMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		body, _ = os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *DigicertMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	if err != nil {
		panic(fmt.Sprintf("DigicertMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func mustMocksDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func (m *DigicertMock) RegisterCreateMock() {
	m.register("POST", "/api/v1/pki/digicert/trust-lifecycle-manager", 201, "validate_create.json")
}

func (m *DigicertMock) RegisterGetByIDMock(id string) {
	m.register("GET", "/api/v1/pki/digicert/trust-lifecycle-manager/"+id, 200, "validate_get.json")
}

func (m *DigicertMock) RegisterUpdateByIDMock(id string) {
	m.register("PATCH", "/api/v1/pki/digicert/trust-lifecycle-manager/"+id, 200, "")
}

func (m *DigicertMock) RegisterDeleteByIDMock(id string) {
	m.register("DELETE", "/api/v1/pki/digicert/trust-lifecycle-manager/"+id, 204, "")
}

func (m *DigicertMock) RegisterValidateClientCertificateMock() {
	m.register("POST", "/api/v1/pki/digicert/trust-lifecycle-manager/validate-client-certificate", 200, "")
}

func (m *DigicertMock) RegisterGetConnectionStatusMock(id string) {
	m.register("GET", "/api/v1/pki/digicert/trust-lifecycle-manager/"+id+"/connection-status", 200, "validate_connection_status.json")
}

func (m *DigicertMock) RegisterGetDependenciesMock(id string) {
	m.register("GET", "/api/v1/pki/digicert/trust-lifecycle-manager/"+id+"/dependencies", 200, "validate_dependencies.json")
}

func (m *DigicertMock) RegisterNotFoundErrorMock(id string) {
	m.registerError("GET", "/api/v1/pki/digicert/trust-lifecycle-manager/"+id, 404, "error_not_found.json")
}

func (m *DigicertMock) RegisterConnectionStatusNotFoundErrorMock(id string) {
	m.registerError("GET", "/api/v1/pki/digicert/trust-lifecycle-manager/"+id+"/connection-status", 404, "error_not_found.json")
}

func (m *DigicertMock) RegisterDependenciesNotFoundErrorMock(id string) {
	m.registerError("GET", "/api/v1/pki/digicert/trust-lifecycle-manager/"+id+"/dependencies", 404, "error_not_found.json")
}

func (m *DigicertMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("DigicertMock: no response for %s %s", method, path)
	}
	headers := http.Header{"Content-Type": {"application/json"}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)
	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}
	if result != nil && len(r.rawBody) > 0 {
		_ = json.Unmarshal(r.rawBody, result)
	}
	return resp, nil
}

func (m *DigicertMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}

func (m *DigicertMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *DigicertMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *DigicertMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *DigicertMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *DigicertMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *DigicertMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *DigicertMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *DigicertMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *DigicertMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *DigicertMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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

func (m *DigicertMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *DigicertMock) InvalidateToken() error                    { return nil }
func (m *DigicertMock) KeepAliveToken() error                     { return nil }
func (m *DigicertMock) GetLogger() *zap.Logger                    { return m.logger }
