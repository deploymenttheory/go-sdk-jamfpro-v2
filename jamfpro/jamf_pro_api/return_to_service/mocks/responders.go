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

type ReturnToServiceMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewReturnToServiceMock() *ReturnToServiceMock {
	return &ReturnToServiceMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *ReturnToServiceMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		body, _ = os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *ReturnToServiceMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	if err != nil {
		panic(fmt.Sprintf("ReturnToServiceMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *ReturnToServiceMock) RegisterListMock() {
	m.register("GET", "/api/v1/return-to-service", 200, "validate_list.json")
}

func (m *ReturnToServiceMock) RegisterGetByIDMock() {
	m.register("GET", "/api/v1/return-to-service/1", 200, "validate_get.json")
}

func (m *ReturnToServiceMock) RegisterCreateMock() {
	m.register("POST", "/api/v1/return-to-service", 201, "validate_create.json")
}

func (m *ReturnToServiceMock) RegisterUpdateMock() {
	m.register("PUT", "/api/v1/return-to-service/1", 200, "validate_update.json")
}

func (m *ReturnToServiceMock) RegisterDeleteMock() {
	m.register("DELETE", "/api/v1/return-to-service/1", 204, "")
}

func (m *ReturnToServiceMock) RegisterListErrorMock() {
	m.registerError("GET", "/api/v1/return-to-service", 500, "error_api.json")
}

func (m *ReturnToServiceMock) RegisterGetByIDErrorMock() {
	m.registerError("GET", "/api/v1/return-to-service/1", 500, "error_api.json")
}

func (m *ReturnToServiceMock) RegisterCreateErrorMock() {
	m.registerError("POST", "/api/v1/return-to-service", 500, "error_api.json")
}

func (m *ReturnToServiceMock) RegisterUpdateErrorMock() {
	m.registerError("PUT", "/api/v1/return-to-service/1", 500, "error_api.json")
}

func (m *ReturnToServiceMock) RegisterDeleteErrorMock() {
	m.registerError("DELETE", "/api/v1/return-to-service/1", 500, "error_api.json")
}

func (m *ReturnToServiceMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("ReturnToServiceMock: no response for %s %s", method, path)
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

func mustMocksDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func (m *ReturnToServiceMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *ReturnToServiceMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ReturnToServiceMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ReturnToServiceMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ReturnToServiceMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ReturnToServiceMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *ReturnToServiceMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *ReturnToServiceMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *ReturnToServiceMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *ReturnToServiceMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}
func (m *ReturnToServiceMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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
func (m *ReturnToServiceMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *ReturnToServiceMock) InvalidateToken() error                    { return nil }
func (m *ReturnToServiceMock) KeepAliveToken() error                     { return nil }
func (m *ReturnToServiceMock) GetLogger() *zap.Logger                    { return m.logger }
