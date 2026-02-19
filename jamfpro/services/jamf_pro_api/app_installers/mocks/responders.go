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

// AppInstallersMock is a test double implementing interfaces.HTTPClient.
type AppInstallersMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewAppInstallersMock returns an empty mock ready for response registration.
func NewAppInstallersMock() *AppInstallersMock {
	return &AppInstallersMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func (m *AppInstallersMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("AppInstallersMock: load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *AppInstallersMock) RegisterMocks() {
	m.register("GET", "/api/v1/app-installers/titles", 200, "validate_list_titles.json")
	m.register("GET", "/api/v1/app-installers/titles/1", 200, "validate_get_title.json")
	m.register("GET", "/api/v1/app-installers/deployments", 200, "validate_list_deployments.json")
	m.register("GET", "/api/v1/app-installers/deployments/1", 200, "validate_get_deployment.json")
	m.register("POST", "/api/v1/app-installers/deployments", 201, "validate_create_deployment.json")
	m.register("PUT", "/api/v1/app-installers/deployments/1", 200, "validate_get_deployment.json")
	m.register("DELETE", "/api/v1/app-installers/deployments/1", 204, "")
}

func (m *AppInstallersMock) Get(ctx context.Context, path string, q map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *AppInstallersMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *AppInstallersMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *AppInstallersMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *AppInstallersMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *AppInstallersMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *AppInstallersMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *AppInstallersMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *AppInstallersMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *AppInstallersMock) GetBytes(ctx context.Context, path string, q map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *AppInstallersMock) GetPaginated(ctx context.Context, path string, q map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil && len(resp.Body) > 0 {
		_ = mergePage(resp.Body)
	}
	return resp, nil
}
func (m *AppInstallersMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *AppInstallersMock) InvalidateToken() error                    { return nil }
func (m *AppInstallersMock) KeepAliveToken() error                     { return nil }
func (m *AppInstallersMock) GetLogger() *zap.Logger                     { return m.logger }

func (m *AppInstallersMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{StatusCode: 404, Status: "404", Headers: http.Header{}, Body: nil},
			fmt.Errorf("AppInstallersMock: no response for %s %s", method, path)
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
	dir, _ := os.Getwd()
	return os.ReadFile(filepath.Join(dir, "mocks", filename))
}
