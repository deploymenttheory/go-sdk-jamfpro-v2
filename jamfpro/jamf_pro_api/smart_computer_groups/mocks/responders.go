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
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"go.uber.org/zap"
	"resty.dev/v3"
)

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// SmartComputerGroupsMock is a test double implementing client.Client.
type SmartComputerGroupsMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewSmartComputerGroupsMock returns an empty mock ready for response registration.
func NewSmartComputerGroupsMock() *SmartComputerGroupsMock {
	return &SmartComputerGroupsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *SmartComputerGroupsMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetByIDMock()
	m.RegisterGetMembershipMock()
	m.RegisterCreateMock()
	m.RegisterUpdateMock()
	m.RegisterDeleteMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *SmartComputerGroupsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *SmartComputerGroupsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("SmartComputerGroupsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+normalizePath(path)] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *SmartComputerGroupsMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("SmartComputerGroupsMock: failed to load error fixture %q: %v", fixture, err))
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

func (m *SmartComputerGroupsMock) dispatch(method, path string, result any) (*resty.Response, error) {
	key := method + ":" + normalizePath(path)
	r, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("SmartComputerGroupsMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {"application/json"}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("SmartComputerGroupsMock: unmarshal into result: %w", err)
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

func (m *SmartComputerGroupsMock) RegisterListMock() {
	m.register("GET", "/api/v2/computer-groups/smart-groups", 200, "validate_list.json")
}

// RegisterListEmptyMock registers a list response with no results (for GetByName not found).
func (m *SmartComputerGroupsMock) RegisterListEmptyMock() {
	m.register("GET", "/api/v2/computer-groups/smart-groups", 200, "validate_list_empty.json")
}

func (m *SmartComputerGroupsMock) RegisterGetByIDMock() {
	m.register("GET", "/api/v2/computer-groups/smart-groups/1", 200, "validate_get_by_id.json")
}

func (m *SmartComputerGroupsMock) RegisterGetMembershipMock() {
	m.register("GET", "/api/v2/computer-groups/smart-group-membership/1", 200, "validate_membership.json")
}

func (m *SmartComputerGroupsMock) RegisterCreateMock() {
	m.register("POST", "/api/v2/computer-groups/smart-groups", 201, "validate_create.json")
}

func (m *SmartComputerGroupsMock) RegisterUpdateMock() {
	m.register("PUT", "/api/v2/computer-groups/smart-groups/1", 200, "validate_update.json")
}

func (m *SmartComputerGroupsMock) RegisterDeleteMock() {
	m.register("DELETE", "/api/v2/computer-groups/smart-groups/1", 204, "")
}

func (m *SmartComputerGroupsMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/api/v2/computer-groups/smart-groups/999", 404, "error_not_found.json")
	m.registerError("GET", "/api/v2/computer-groups/smart-group-membership/999", 404, "error_not_found.json")
}

func (m *SmartComputerGroupsMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/api/v2/computer-groups/smart-groups", 409, "error_conflict.json")
}

func (m *SmartComputerGroupsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}

func (m *SmartComputerGroupsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *SmartComputerGroupsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *SmartComputerGroupsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *SmartComputerGroupsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *SmartComputerGroupsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *SmartComputerGroupsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *SmartComputerGroupsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *SmartComputerGroupsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *SmartComputerGroupsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *SmartComputerGroupsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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

func (m *SmartComputerGroupsMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *SmartComputerGroupsMock) InvalidateToken() error                { return nil }
func (m *SmartComputerGroupsMock) KeepAliveToken() error                 { return nil }
func (m *SmartComputerGroupsMock) GetLogger() *zap.Logger                { return m.logger }
