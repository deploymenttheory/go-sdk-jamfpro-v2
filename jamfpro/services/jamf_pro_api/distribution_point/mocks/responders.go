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

// DistributionPointMock implements interfaces.HTTPClient.
type DistributionPointMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewDistributionPointMock() *DistributionPointMock {
	return &DistributionPointMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func (m *DistributionPointMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("DistributionPointMock: load %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *DistributionPointMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("DistributionPointMock: load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)}
}

func (m *DistributionPointMock) RegisterMocks() {
	m.register("GET", "/api/v1/distribution-points", 200, "validate_list.json")
	m.register("POST", "/api/v1/distribution-points", 201, "validate_create.json")
	m.register("POST", "/api/v1/distribution-points/delete-multiple", 204, "")
	m.register("GET", "/api/v1/distribution-points/1", 200, "validate_get.json")
	m.register("PUT", "/api/v1/distribution-points/1", 200, "validate_get.json")
	m.register("DELETE", "/api/v1/distribution-points/1", 204, "")
	m.register("PATCH", "/api/v1/distribution-points/1", 200, "validate_get.json")
	m.register("GET", "/api/v1/distribution-points/1/history", 200, "validate_history.json")
	m.register("POST", "/api/v1/distribution-points/1/history", 201, "validate_history_note.json")
}

func (m *DistributionPointMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/api/v1/distribution-points/999", 404, "error_not_found.json")
}

// RegisterListInvalidMock registers GET list with invalid JSON for testing mergePage error paths.
func (m *DistributionPointMock) RegisterListInvalidMock() {
	m.register("GET", "/api/v1/distribution-points", 200, "validate_list_invalid.json")
}

// RegisterHistoryInvalidMock registers GET history with invalid JSON for testing mergePage error paths.
func (m *DistributionPointMock) RegisterHistoryInvalidMock() {
	m.register("GET", "/api/v1/distribution-points/1/history", 200, "validate_history_invalid.json")
}

func (m *DistributionPointMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("DistributionPointMock: no response for %s %s", method, path)
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

func (m *DistributionPointMock) Get(ctx context.Context, path string, q map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *DistributionPointMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *DistributionPointMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *DistributionPointMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *DistributionPointMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *DistributionPointMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *DistributionPointMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *DistributionPointMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *DistributionPointMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *DistributionPointMock) GetBytes(ctx context.Context, path string, q map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *DistributionPointMock) GetPaginated(ctx context.Context, path string, q map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil && len(resp.Body) > 0 {
		// Parse the paginated response structure to extract the results field
		var pageResp struct {
			Results json.RawMessage `json:"results"`
		}
		if err := json.Unmarshal(resp.Body, &pageResp); err != nil {
			return resp, fmt.Errorf("failed to unmarshal paginated response: %w", err)
		}
		if err := mergePage(pageResp.Results); err != nil {
			return resp, err
		}
	}
	return resp, nil
}
func (m *DistributionPointMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *DistributionPointMock) InvalidateToken() error                    { return nil }
func (m *DistributionPointMock) KeepAliveToken() error                     { return nil }
func (m *DistributionPointMock) GetLogger() *zap.Logger                    { return m.logger }
