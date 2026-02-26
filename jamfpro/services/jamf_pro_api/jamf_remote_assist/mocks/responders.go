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

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// JamfRemoteAssistMock implements interfaces.HTTPClient for tests.
type JamfRemoteAssistMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewJamfRemoteAssistMock() *JamfRemoteAssistMock {
	return &JamfRemoteAssistMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *JamfRemoteAssistMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("JamfRemoteAssistMock: load %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *JamfRemoteAssistMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("JamfRemoteAssistMock: load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *JamfRemoteAssistMock) RegisterMocks() {
	m.register("GET", "/api/v1/jamf-remote-assist/session", 200, "validate_list.json")
	m.register("GET", "/api/v1/jamf-remote-assist/session/session-abc", 200, "validate_session.json")
	m.register("GET", "/api/v2/jamf-remote-assist/session", 200, "validate_list_v2.json")
	m.register("GET", "/api/v2/jamf-remote-assist/session/session-abc", 200, "validate_session.json")
	m.register("POST", "/api/v2/jamf-remote-assist/session/export", 200, "")
}

// RegisterListSessionsV1ErrorMock registers an error response for ListSessionsV1.
func (m *JamfRemoteAssistMock) RegisterListSessionsV1ErrorMock() {
	m.registerError("GET", "/api/v1/jamf-remote-assist/session", 500, "error_not_found.json")
}

// RegisterListSessionsV2ErrorMock registers an error response for ListSessionsV2.
func (m *JamfRemoteAssistMock) RegisterListSessionsV2ErrorMock() {
	m.registerError("GET", "/api/v2/jamf-remote-assist/session", 500, "error_not_found.json")
}

// RegisterListSessionsV2InvalidMock registers invalid JSON for ListSessionsV2 (triggers mergePage error).
func (m *JamfRemoteAssistMock) RegisterListSessionsV2InvalidMock() {
	m.register("GET", "/api/v2/jamf-remote-assist/session", 200, "validate_list_v2_invalid.json")
}

// RegisterGetSessionByIDV2ErrorMock registers an error response for GetSessionByIDV2.
func (m *JamfRemoteAssistMock) RegisterGetSessionByIDV2ErrorMock() {
	m.registerError("GET", "/api/v2/jamf-remote-assist/session/nonexistent", 404, "error_not_found.json")
}

// RegisterExportSessionsV2ErrorMock registers an error response for ExportSessionsV2.
func (m *JamfRemoteAssistMock) RegisterExportSessionsV2ErrorMock() {
	m.registerError("POST", "/api/v2/jamf-remote-assist/session/export", 500, "error_not_found.json")
}

func (m *JamfRemoteAssistMock) Get(ctx context.Context, path string, q map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *JamfRemoteAssistMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *JamfRemoteAssistMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *JamfRemoteAssistMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *JamfRemoteAssistMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *JamfRemoteAssistMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *JamfRemoteAssistMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *JamfRemoteAssistMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *JamfRemoteAssistMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *JamfRemoteAssistMock) GetBytes(ctx context.Context, path string, q map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *JamfRemoteAssistMock) GetPaginated(ctx context.Context, path string, q map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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
func (m *JamfRemoteAssistMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *JamfRemoteAssistMock) InvalidateToken() error                    { return nil }
func (m *JamfRemoteAssistMock) KeepAliveToken() error                     { return nil }
func (m *JamfRemoteAssistMock) GetLogger() *zap.Logger                    { return m.logger }

func (m *JamfRemoteAssistMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("JamfRemoteAssistMock: no response registered for %s %s", method, path)
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
	_, callerPath, _, _ := runtime.Caller(0)
	dir := filepath.Dir(callerPath)
	return os.ReadFile(filepath.Join(dir, filename))
}
