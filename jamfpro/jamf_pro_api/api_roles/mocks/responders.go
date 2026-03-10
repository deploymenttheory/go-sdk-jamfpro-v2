package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"resty.dev/v3"

	mockhelpers "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"go.uber.org/zap"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// APIRolesMock is a test double implementing client.Client.
type APIRolesMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewAPIRolesMock returns an empty mock ready for response registration.
func NewAPIRolesMock() *APIRolesMock {
	return &APIRolesMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func (m *APIRolesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("APIRolesMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *APIRolesMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("APIRolesMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)}
}

// RegisterMocks registers all standard success responses.
func (m *APIRolesMock) RegisterMocks() {
	m.register("GET", "/api/v1/api-roles", 200, "validate_list.json")
	m.register("GET", "/api/v1/api-roles/1", 200, "validate_get.json")
	m.register("POST", "/api/v1/api-roles", 200, "validate_create.json")
	m.register("PUT", "/api/v1/api-roles/1", 200, "validate_get.json")
	m.register("DELETE", "/api/v1/api-roles/1", 204, "")
}

// RegisterNotFoundErrorMock registers GET .../999 → 404.
func (m *APIRolesMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/api/v1/api-roles/999", 404, "error_not_found.json")
}

func (m *APIRolesMock) Get(ctx context.Context, path string, q map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *APIRolesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *APIRolesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *APIRolesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *APIRolesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *APIRolesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *APIRolesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *APIRolesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *APIRolesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *APIRolesMock) GetBytes(ctx context.Context, path string, q map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}
func (m *APIRolesMock) GetPaginated(ctx context.Context, path string, q map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	bodyBytes := resp.Bytes()
	if mergePage != nil && len(bodyBytes) > 0 {
		var wrapper struct {
			Results json.RawMessage `json:"results"`
		}
		if err := json.Unmarshal(bodyBytes, &wrapper); err != nil {
			return resp, fmt.Errorf("failed to extract results field: %w", err)
		}
		if err := mergePage(wrapper.Results); err != nil {
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
	}
	return resp, nil
}
func (m *APIRolesMock) NewRequest(ctx context.Context) *client.RequestBuilder {
	return client.NewMockRequestBuilder(ctx, func(method, path string, result any) (*resty.Response, error) {
		return m.dispatch(method, path, result)
	})
}
func (m *APIRolesMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *APIRolesMock) InvalidateToken() error                { return nil }
func (m *APIRolesMock) KeepAliveToken() error                 { return nil }
func (m *APIRolesMock) GetLogger() *zap.Logger                { return m.logger }

func (m *APIRolesMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return mockhelpers.NewMockResponse(http.StatusNotFound, http.Header{}, nil), fmt.Errorf("APIRolesMock: no response for %s %s", method, path)
	}
	headers := http.Header{"Content-Type": {"application/json"}}
	resp := mockhelpers.NewMockResponse(r.statusCode, headers, r.rawBody)
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
