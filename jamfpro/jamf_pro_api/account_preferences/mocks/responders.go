package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"resty.dev/v3"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
}

// AccountPreferencesMock is a test double implementing interfaces.HTTPClient.
type AccountPreferencesMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewAccountPreferencesMock returns an empty mock ready for response registration.
func NewAccountPreferencesMock() *AccountPreferencesMock {
	return &AccountPreferencesMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *AccountPreferencesMock) register(method, path string, statusCode int, fixture string) {
	body, _ := os.ReadFile(filepath.Join(mustGetwd(), "mocks", fixture))
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// RegisterGetAccountPreferencesMock registers a successful GET /api/v3/account-preferences response.
func (m *AccountPreferencesMock) RegisterGetAccountPreferencesMock() {
	m.register("GET", "/api/v3/account-preferences", 200, "validate_get.json")
}

// RegisterUpdateAccountPreferencesMock registers a successful PATCH /api/v3/account-preferences response.
func (m *AccountPreferencesMock) RegisterUpdateAccountPreferencesMock() {
	m.register("PATCH", "/api/v3/account-preferences", 200, "validate_get.json")
}

func (m *AccountPreferencesMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return shared.NewMockResponse(http.StatusNotFound, http.Header{}, nil), fmt.Errorf("AccountPreferencesMock: no response for %s %s", method, path)
	}
	headers := http.Header{"Content-Type": {"application/json"}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)
	if result != nil && len(r.rawBody) > 0 {
		_ = json.Unmarshal(r.rawBody, result)
	}
	return resp, nil
}

func mustGetwd() string {
	dir, _ := os.Getwd()
	return dir
}

func (m *AccountPreferencesMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *AccountPreferencesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *AccountPreferencesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *AccountPreferencesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *AccountPreferencesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *AccountPreferencesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *AccountPreferencesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *AccountPreferencesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *AccountPreferencesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *AccountPreferencesMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}
func (m *AccountPreferencesMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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
func (m *AccountPreferencesMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *AccountPreferencesMock) InvalidateToken() error                    { return nil }
func (m *AccountPreferencesMock) KeepAliveToken() error                     { return nil }
func (m *AccountPreferencesMock) GetLogger() *zap.Logger                    { return m.logger }
