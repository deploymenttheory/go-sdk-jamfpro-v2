package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"go.uber.org/zap"
	"resty.dev/v3"

	mockhelpers "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// StaticComputerGroupsMock is a test double implementing client.Client.
type StaticComputerGroupsMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewStaticComputerGroupsMock returns an empty mock ready for response registration.
func NewStaticComputerGroupsMock() *StaticComputerGroupsMock {
	return &StaticComputerGroupsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *StaticComputerGroupsMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetByIDMock()
	m.RegisterCreateMock()
	m.RegisterUpdateMock()
	m.RegisterDeleteMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *StaticComputerGroupsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *StaticComputerGroupsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("StaticComputerGroupsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+normalizePath(path)] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *StaticComputerGroupsMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("StaticComputerGroupsMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+normalizePath(path)] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

// normalizePath strips query string for consistent mock key lookup.
func normalizePath(path string) string {
	if idx := strings.Index(path, "?"); idx >= 0 {
		return path[:idx]
	}
	return path
}

func (m *StaticComputerGroupsMock) dispatch(method, path string, result any) (*resty.Response, error) {
	key := method + ":" + normalizePath(path)
	r, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("StaticComputerGroupsMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {"application/json"}}
	resp := mockhelpers.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("StaticComputerGroupsMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

func loadMockResponse(filename string) ([]byte, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("get working directory: %w", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "mocks", filename))
	if err != nil {
		return nil, fmt.Errorf("read fixture %s: %w", filename, err)
	}
	return data, nil
}

func (m *StaticComputerGroupsMock) RegisterListMock() {
	m.register("GET", "/api/v2/computer-groups/static-groups", 200, "validate_list.json")
}

func (m *StaticComputerGroupsMock) RegisterGetByIDMock() {
	m.register("GET", "/api/v2/computer-groups/static-groups/1", 200, "validate_get.json")
}

func (m *StaticComputerGroupsMock) RegisterCreateMock() {
	m.register("POST", "/api/v2/computer-groups/static-groups", 201, "validate_create.json")
}

func (m *StaticComputerGroupsMock) RegisterUpdateMock() {
	m.register("PUT", "/api/v2/computer-groups/static-groups/1", 200, "validate_update.json")
}

func (m *StaticComputerGroupsMock) RegisterDeleteMock() {
	m.register("DELETE", "/api/v2/computer-groups/static-groups/1", 204, "")
}

func (m *StaticComputerGroupsMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/api/v2/computer-groups/static-groups/999", 404, "error_not_found.json")
}

func (m *StaticComputerGroupsMock) RegisterUpdateNotFoundErrorMock() {
	m.registerError("PUT", "/api/v2/computer-groups/static-groups/999", 404, "error_not_found.json")
}

func (m *StaticComputerGroupsMock) RegisterDeleteNotFoundErrorMock() {
	m.registerError("DELETE", "/api/v2/computer-groups/static-groups/999", 404, "error_not_found.json")
}

func (m *StaticComputerGroupsMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/api/v2/computer-groups/static-groups", 409, "error_conflict.json")
}

func (m *StaticComputerGroupsMock) RegisterListEmptyMock() {
	m.register("GET", "/api/v2/computer-groups/static-groups", 200, "validate_list_empty.json")
}

func (m *StaticComputerGroupsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}

func (m *StaticComputerGroupsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *StaticComputerGroupsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *StaticComputerGroupsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *StaticComputerGroupsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *StaticComputerGroupsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *StaticComputerGroupsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *StaticComputerGroupsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *StaticComputerGroupsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *StaticComputerGroupsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *StaticComputerGroupsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	bodyBytes := resp.Bytes()
	if mergePage != nil && len(bodyBytes) > 0 {
		var pageResp struct {
			Results json.RawMessage `json:"results"`
		}
		if err := json.Unmarshal(bodyBytes, &pageResp); err != nil {
			return resp, fmt.Errorf("failed to extract results field: %w", err)
		}
		if err := mergePage(pageResp.Results); err != nil {
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
	}
	return resp, nil
}

func (m *StaticComputerGroupsMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *StaticComputerGroupsMock) InvalidateToken() error                { return nil }
func (m *StaticComputerGroupsMock) KeepAliveToken() error                 { return nil }
func (m *StaticComputerGroupsMock) GetLogger() *zap.Logger                { return m.logger }
