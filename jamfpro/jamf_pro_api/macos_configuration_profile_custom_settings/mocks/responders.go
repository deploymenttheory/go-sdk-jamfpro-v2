package mocks

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"resty.dev/v3"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"go.uber.org/zap"
)

//go:embed *.json
var fixtureFS embed.FS

type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// MacOSConfigProfileCustomSettingsMock is a test double implementing transport.HTTPClient
// for macOS configuration profile custom settings operations.
//
// Responses are keyed by "METHOD path". Use RegisterGetSchemaListMock,
// RegisterGetByPayloadUUIDMock, and RegisterCreateMock to set up responses.
type MacOSConfigProfileCustomSettingsMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewMacOSConfigProfileCustomSettingsMock returns an empty mock ready for response registration.
func NewMacOSConfigProfileCustomSettingsMock() *MacOSConfigProfileCustomSettingsMock {
	return &MacOSConfigProfileCustomSettingsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func (m *MacOSConfigProfileCustomSettingsMock) register(method, path string, statusCode int, fixtureFile string) {
	key := method + " " + path
	var body []byte
	if fixtureFile != "" {
		data, err := fixtureFS.ReadFile(fixtureFile)
		if err != nil {
			panic(fmt.Sprintf("MacOSConfigProfileCustomSettingsMock: failed to load fixture %q: %v", fixtureFile, err))
		}
		body = data
	}
	m.responses[key] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
	}
}

// RegisterGetSchemaListMock registers GET /api/config-profiles/macos/custom-settings/v1/schema-list → 200.
func (m *MacOSConfigProfileCustomSettingsMock) RegisterGetSchemaListMock() {
	m.register("GET", "/api/config-profiles/macos/custom-settings/v1/schema-list", 200, "validate_list_schema.json")
}

// RegisterGetByPayloadUUIDMock registers GET /api/config-profiles/macos/{id} → 200.
// Use the same id in tests as passed here (e.g. "test-uuid-12345").
func (m *MacOSConfigProfileCustomSettingsMock) RegisterGetByPayloadUUIDMock(id string) {
	path := "/api/config-profiles/macos/" + id
	m.register("GET", path, 200, "validate_get.json")
}

// RegisterCreateMock registers POST /api/config-profiles/macos → 200.
func (m *MacOSConfigProfileCustomSettingsMock) RegisterCreateMock() {
	m.register("POST", "/api/config-profiles/macos", 200, "validate_create.json")
}

// RegisterGetSchemaListErrorMock registers GET schema-list with a 500 error.
func (m *MacOSConfigProfileCustomSettingsMock) RegisterGetSchemaListErrorMock() {
	m.responses["GET /api/config-profiles/macos/custom-settings/v1/schema-list"] = registeredResponse{
		statusCode: 500,
		rawBody:    []byte(`{"error":"internal server error"}`),
		errMsg:     "request failed: 500 Internal Server Error",
	}
}

// RegisterGetByPayloadUUIDErrorMock registers GET by ID with a 404 error.
func (m *MacOSConfigProfileCustomSettingsMock) RegisterGetByPayloadUUIDErrorMock(id string) {
	path := "/api/config-profiles/macos/" + id
	m.responses["GET "+path] = registeredResponse{
		statusCode: 404,
		rawBody:    []byte(`{"error":"not found"}`),
		errMsg:     "request failed: 404 Not Found",
	}
}

// RegisterCreateErrorMock registers POST with a 500 error.
func (m *MacOSConfigProfileCustomSettingsMock) RegisterCreateErrorMock() {
	m.responses["POST /api/config-profiles/macos"] = registeredResponse{
		statusCode: 500,
		rawBody:    []byte(`{"error":"internal server error"}`),
		errMsg:     "request failed: 500 Internal Server Error",
	}
}

func (m *MacOSConfigProfileCustomSettingsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*resty.Response, error) {
	key := "GET " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for GET %s", path)
	}
	if resp.errMsg != "" {
		return nil, fmt.Errorf("%s", resp.errMsg)
	}
	if result != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return shared.NewMockResponse(resp.statusCode, http.Header{}, resp.rawBody), nil
}

func (m *MacOSConfigProfileCustomSettingsMock) Post(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	key := "POST " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for POST %s", path)
	}
	if resp.errMsg != "" {
		return nil, fmt.Errorf("%s", resp.errMsg)
	}
	if result != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return shared.NewMockResponse(resp.statusCode, http.Header{}, resp.rawBody), nil
}

func (m *MacOSConfigProfileCustomSettingsMock) PostWithQuery(ctx context.Context, path string, rsqlQuery map[string]string, body any, headers map[string]string, result any) (*resty.Response, error) {
	return m.Post(ctx, path, body, headers, result)
}

func (m *MacOSConfigProfileCustomSettingsMock) PostForm(ctx context.Context, path string, formData map[string]string, headers map[string]string, result any) (*resty.Response, error) {
	return m.Post(ctx, path, formData, headers, result)
}

func (m *MacOSConfigProfileCustomSettingsMock) PostMultipart(ctx context.Context, path string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback transport.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.Post(ctx, path, nil, headers, result)
}

func (m *MacOSConfigProfileCustomSettingsMock) Put(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	key := "PUT " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for PUT %s", path)
	}
	if resp.errMsg != "" {
		return nil, fmt.Errorf("%s", resp.errMsg)
	}
	if result != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return shared.NewMockResponse(resp.statusCode, http.Header{}, resp.rawBody), nil
}

func (m *MacOSConfigProfileCustomSettingsMock) Patch(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	key := "PATCH " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for PATCH %s", path)
	}
	if resp.errMsg != "" {
		return nil, fmt.Errorf("%s", resp.errMsg)
	}
	if result != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return shared.NewMockResponse(resp.statusCode, http.Header{}, resp.rawBody), nil
}

func (m *MacOSConfigProfileCustomSettingsMock) Delete(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*resty.Response, error) {
	key := "DELETE " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for DELETE %s", path)
	}
	if resp.errMsg != "" {
		return nil, fmt.Errorf("%s", resp.errMsg)
	}
	if result != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return shared.NewMockResponse(resp.statusCode, http.Header{}, resp.rawBody), nil
}

func (m *MacOSConfigProfileCustomSettingsMock) DeleteWithBody(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	return m.Delete(ctx, path, nil, headers, result)
}

func (m *MacOSConfigProfileCustomSettingsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string) (*resty.Response, []byte, error) {
	key := "GET " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, nil, fmt.Errorf("no mock registered for GET %s", path)
	}
	if resp.errMsg != "" {
		return nil, nil, fmt.Errorf("%s", resp.errMsg)
	}
	return shared.NewMockResponse(resp.statusCode, http.Header{}, resp.rawBody), resp.rawBody, nil
}

func (m *MacOSConfigProfileCustomSettingsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, mergePage func(pageData []byte) error) (*resty.Response, error) {
	return nil, fmt.Errorf("GetPaginated not implemented in MacOSConfigProfileCustomSettingsMock")
}

func (m *MacOSConfigProfileCustomSettingsMock) RSQLBuilder() transport.RSQLFilterBuilder {
	return nil
}

func (m *MacOSConfigProfileCustomSettingsMock) InvalidateToken() error {
	return nil
}

func (m *MacOSConfigProfileCustomSettingsMock) KeepAliveToken() error {
	return nil
}

func (m *MacOSConfigProfileCustomSettingsMock) GetLogger() *zap.Logger {
	return m.logger
}
