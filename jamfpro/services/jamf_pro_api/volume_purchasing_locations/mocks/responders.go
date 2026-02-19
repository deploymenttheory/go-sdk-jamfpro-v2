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

// VolumePurchasingLocationsMock is a test double implementing interfaces.HTTPClient.
type VolumePurchasingLocationsMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewVolumePurchasingLocationsMock returns an empty mock ready for response registration.
func NewVolumePurchasingLocationsMock() *VolumePurchasingLocationsMock {
	return &VolumePurchasingLocationsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *VolumePurchasingLocationsMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetMock()
	m.RegisterCreateMock()
	m.RegisterUpdateMock()
	m.RegisterDeleteMock()
	m.RegisterReclaimMock()
	m.RegisterGetContentMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *VolumePurchasingLocationsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
}

func (m *VolumePurchasingLocationsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("VolumePurchasingLocationsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *VolumePurchasingLocationsMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("VolumePurchasingLocationsMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *VolumePurchasingLocationsMock) RegisterListMock() {
	m.register("GET", "/api/v1/volume-purchasing-locations", 200, "validate_list.json")
}

func (m *VolumePurchasingLocationsMock) RegisterGetMock() {
	m.register("GET", "/api/v1/volume-purchasing-locations/1", 200, "validate_get.json")
}

func (m *VolumePurchasingLocationsMock) RegisterCreateMock() {
	m.register("POST", "/api/v1/volume-purchasing-locations", 201, "validate_create.json")
}

func (m *VolumePurchasingLocationsMock) RegisterUpdateMock() {
	m.register("PATCH", "/api/v1/volume-purchasing-locations/1", 200, "validate_update.json")
}

func (m *VolumePurchasingLocationsMock) RegisterDeleteMock() {
	m.register("DELETE", "/api/v1/volume-purchasing-locations/1", 204, "")
}

func (m *VolumePurchasingLocationsMock) RegisterReclaimMock() {
	m.register("POST", "/api/v1/volume-purchasing-locations/1/reclaim", 202, "")
}

func (m *VolumePurchasingLocationsMock) RegisterGetContentMock() {
	m.register("GET", "/api/v1/volume-purchasing-locations/1/content", 200, "validate_content.json")
}

func (m *VolumePurchasingLocationsMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/api/v1/volume-purchasing-locations/999", 404, "error_not_found.json")
}

func (m *VolumePurchasingLocationsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *VolumePurchasingLocationsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *VolumePurchasingLocationsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *VolumePurchasingLocationsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *VolumePurchasingLocationsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *VolumePurchasingLocationsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *VolumePurchasingLocationsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *VolumePurchasingLocationsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *VolumePurchasingLocationsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *VolumePurchasingLocationsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *VolumePurchasingLocationsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil && len(resp.Body) > 0 {
		if err := mergePage(resp.Body); err != nil {
			return resp, err
		}
	}
	return resp, nil
}

func (m *VolumePurchasingLocationsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *VolumePurchasingLocationsMock) InvalidateToken() error                      { return nil }
func (m *VolumePurchasingLocationsMock) KeepAliveToken() error                        { return nil }
func (m *VolumePurchasingLocationsMock) GetLogger() *zap.Logger                        { return m.logger }

func (m *VolumePurchasingLocationsMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/json"}},
			Body:       []byte(`{"code":"NOT-FOUND","message":"no mock registered"}`),
		}, fmt.Errorf("VolumePurchasingLocationsMock: no response registered for %s %s", method, path)
	}

	resp := &interfaces.Response{
		StatusCode: r.statusCode,
		Status:     fmt.Sprintf("%d", r.statusCode),
		Headers:    http.Header{"Content-Type": {"application/json"}},
		Body:       r.rawBody,
	}

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("VolumePurchasingLocationsMock: unmarshal into result: %w", err)
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
