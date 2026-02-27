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

type AdvancedUserContentSearchesMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewAdvancedUserContentSearchesMock() *AdvancedUserContentSearchesMock {
	return &AdvancedUserContentSearchesMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *AdvancedUserContentSearchesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("AdvancedUserContentSearchesMock: load %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *AdvancedUserContentSearchesMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("AdvancedUserContentSearchesMock: load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)}
}

func (m *AdvancedUserContentSearchesMock) RegisterMocks() {
	m.register("GET", "/api/v1/advanced-user-content-searches", 200, "validate_list.json")
	m.register("GET", "/api/v1/advanced-user-content-searches/1", 200, "validate_get.json")
	m.register("POST", "/api/v1/advanced-user-content-searches", 201, "validate_create.json")
	m.register("PUT", "/api/v1/advanced-user-content-searches/1", 200, "validate_get.json")
	m.register("DELETE", "/api/v1/advanced-user-content-searches/1", 204, "")
}

func (m *AdvancedUserContentSearchesMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/api/v1/advanced-user-content-searches/999", 404, "error_not_found.json")
}

func (m *AdvancedUserContentSearchesMock) Get(ctx context.Context, path string, q map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *AdvancedUserContentSearchesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *AdvancedUserContentSearchesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *AdvancedUserContentSearchesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *AdvancedUserContentSearchesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *AdvancedUserContentSearchesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *AdvancedUserContentSearchesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *AdvancedUserContentSearchesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *AdvancedUserContentSearchesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *AdvancedUserContentSearchesMock) GetBytes(ctx context.Context, path string, q map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *AdvancedUserContentSearchesMock) GetPaginated(ctx context.Context, path string, q map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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
func (m *AdvancedUserContentSearchesMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *AdvancedUserContentSearchesMock) InvalidateToken() error                    { return nil }
func (m *AdvancedUserContentSearchesMock) KeepAliveToken() error                     { return nil }
func (m *AdvancedUserContentSearchesMock) GetLogger() *zap.Logger                     { return m.logger }

func (m *AdvancedUserContentSearchesMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{StatusCode: 404, Headers: http.Header{}, Body: nil}, fmt.Errorf("AdvancedUserContentSearchesMock: no response for %s %s", method, path)
	}
	resp := &interfaces.Response{StatusCode: r.statusCode, Status: fmt.Sprintf("%d", r.statusCode), Headers: http.Header{"Content-Type": {"application/json"}}, Body: r.rawBody}
	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}
	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, err
		}
	}
	return resp, nil
}

func loadMockResponse(filename string) ([]byte, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return os.ReadFile(filepath.Join(dir, "mocks", filename))
}
