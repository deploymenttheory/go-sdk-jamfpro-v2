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

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"go.uber.org/zap"
	"resty.dev/v3"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// SsoSettingsMock is a test double implementing client.Client.
type SsoSettingsMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewSsoSettingsMock returns an empty mock ready for response registration.
func NewSsoSettingsMock() *SsoSettingsMock {
	return &SsoSettingsMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *SsoSettingsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := os.ReadFile(filepath.Join(mustMocksDir(), fixture))
		if err != nil {
			panic(fmt.Sprintf("SsoSettingsMock: load %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *SsoSettingsMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	if err != nil {
		panic(fmt.Sprintf("SsoSettingsMock: load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

// RegisterGetMock registers a successful GET /api/v3/sso response.
func (m *SsoSettingsMock) RegisterGetMock() {
	m.register("GET", "/api/v3/sso", 200, "validate_get.json")
}

// RegisterUpdateMock registers a successful PUT /api/v3/sso response.
func (m *SsoSettingsMock) RegisterUpdateMock() {
	m.register("PUT", "/api/v3/sso", 200, "validate_get.json")
}

// RegisterGetDependenciesMock registers a successful GET /api/v3/sso/dependencies response.
func (m *SsoSettingsMock) RegisterGetDependenciesMock() {
	m.register("GET", "/api/v3/sso/dependencies", 200, "validate_dependencies.json")
}

// RegisterDisableMock registers a successful POST /api/v3/sso/disable response.
func (m *SsoSettingsMock) RegisterDisableMock() {
	m.register("POST", "/api/v3/sso/disable", 204, "")
}

// RegisterGetHistoryMock registers a successful GET /api/v3/sso/history response.
func (m *SsoSettingsMock) RegisterGetHistoryMock() {
	m.register("GET", "/api/v3/sso/history", 200, "validate_history.json")
}

// RegisterAddHistoryNoteMock registers a successful POST /api/v3/sso/history response.
func (m *SsoSettingsMock) RegisterAddHistoryNoteMock() {
	m.register("POST", "/api/v3/sso/history", 201, "validate_add_history_note.json")
}

// RegisterDownloadMetadataMock registers a successful GET /api/v3/sso/metadata/download response.
func (m *SsoSettingsMock) RegisterDownloadMetadataMock() {
	m.register("GET", "/api/v3/sso/metadata/download", 200, "")
}

// RegisterGetErrorMock registers an error response for GetV3.
func (m *SsoSettingsMock) RegisterGetErrorMock() {
	m.registerError("GET", "/api/v3/sso", 500, "error_not_found.json")
}

// RegisterUpdateErrorMock registers an error response for UpdateV3.
func (m *SsoSettingsMock) RegisterUpdateErrorMock() {
	m.registerError("PUT", "/api/v3/sso", 500, "error_not_found.json")
}

// RegisterGetDependenciesErrorMock registers an error response for GetEnrollmentCustomizationDependenciesV3.
func (m *SsoSettingsMock) RegisterGetDependenciesErrorMock() {
	m.registerError("GET", "/api/v3/sso/dependencies", 500, "error_not_found.json")
}

// RegisterDisableErrorMock registers an error response for DisableV3.
func (m *SsoSettingsMock) RegisterDisableErrorMock() {
	m.registerError("POST", "/api/v3/sso/disable", 500, "error_not_found.json")
}

// RegisterGetHistoryErrorMock registers an error response for GetHistoryV3.
func (m *SsoSettingsMock) RegisterGetHistoryErrorMock() {
	m.registerError("GET", "/api/v3/sso/history", 500, "error_not_found.json")
}

// RegisterAddHistoryNoteErrorMock registers an error response for AddHistoryNoteV3.
func (m *SsoSettingsMock) RegisterAddHistoryNoteErrorMock() {
	m.registerError("POST", "/api/v3/sso/history", 500, "error_not_found.json")
}

// RegisterDownloadMetadataErrorMock registers an error response for DownloadMetadataV3.
func (m *SsoSettingsMock) RegisterDownloadMetadataErrorMock() {
	m.registerError("GET", "/api/v3/sso/metadata/download", 500, "error_not_found.json")
}

func (m *SsoSettingsMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("SsoSettingsMock: no response for %s %s", method, path)
	}
	headers := http.Header{"Content-Type": {"application/json"}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)
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

func mustMocksDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func (m *SsoSettingsMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *SsoSettingsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *SsoSettingsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *SsoSettingsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *SsoSettingsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *SsoSettingsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *SsoSettingsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *SsoSettingsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *SsoSettingsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *SsoSettingsMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	return resp, resp.Bytes(), nil
}
func (m *SsoSettingsMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	bodyBytes := resp.Bytes()
	if mergePage != nil && len(bodyBytes) > 0 {
		var page struct {
			Results json.RawMessage `json:"results"`
		}
		if err := json.Unmarshal(bodyBytes, &page); err != nil {
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
		if err := mergePage(page.Results); err != nil {
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
	}
	return resp, nil
}
func (m *SsoSettingsMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *SsoSettingsMock) InvalidateToken() error                { return nil }
func (m *SsoSettingsMock) KeepAliveToken() error                 { return nil }
func (m *SsoSettingsMock) GetLogger() *zap.Logger                { return m.logger }
