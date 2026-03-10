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
	"go.uber.org/zap"
	"resty.dev/v3"

	mockhelpers "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// UserAccountSettingsMock is a test double implementing client.Client.
type UserAccountSettingsMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewUserAccountSettingsMock returns an empty mock ready for response registration.
func NewUserAccountSettingsMock() *UserAccountSettingsMock {
	return &UserAccountSettingsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func mustMocksDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func (m *UserAccountSettingsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := os.ReadFile(filepath.Join(mustMocksDir(), fixture))
		if err != nil {
			panic(fmt.Sprintf("UserAccountSettingsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *UserAccountSettingsMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	if err != nil {
		panic(fmt.Sprintf("UserAccountSettingsMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *UserAccountSettingsMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("UserAccountSettingsMock: no response registered for %s %s", method, path)
	}
	headers := http.Header{"Content-Type": {"application/json"}}
	resp := mockhelpers.NewMockResponse(r.statusCode, headers, r.rawBody)
	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}
	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("UserAccountSettingsMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

const (
	settingsPath = "/api/v1/user/preferences/settings"
	prefsPath    = "/api/v1/user/preferences"
)

func (m *UserAccountSettingsMock) RegisterGetSettingsV1Mock(keyID string) {
	m.register("GET", settingsPath+"/"+keyID, 200, "validate_get_settings.json")
}

func (m *UserAccountSettingsMock) RegisterGetV1Mock(keyID string) {
	m.register("GET", prefsPath+"/"+keyID, 200, "validate_get.json")
}

func (m *UserAccountSettingsMock) RegisterGetV1PlainTextMock(keyID string) {
	m.responses["GET:"+prefsPath+"/"+keyID] = registeredResponse{statusCode: 200, rawBody: []byte("plaintext-value")}
}

func (m *UserAccountSettingsMock) RegisterPutV1Mock(keyID string) {
	m.register("PUT", prefsPath+"/"+keyID, 200, "")
}

func (m *UserAccountSettingsMock) RegisterDeleteV1Mock(keyID string) {
	m.register("DELETE", prefsPath+"/"+keyID, 204, "")
}

func (m *UserAccountSettingsMock) RegisterNotFoundErrorMock(keyID string) {
	m.registerError("GET", settingsPath+"/"+keyID, 404, "error_not_found.json")
	m.registerError("GET", prefsPath+"/"+keyID, 404, "error_not_found.json")
	m.registerError("PUT", prefsPath+"/"+keyID, 404, "error_not_found.json")
	m.registerError("DELETE", prefsPath+"/"+keyID, 404, "error_not_found.json")
}

func (m *UserAccountSettingsMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}

func (m *UserAccountSettingsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *UserAccountSettingsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *UserAccountSettingsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *UserAccountSettingsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *UserAccountSettingsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *UserAccountSettingsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *UserAccountSettingsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *UserAccountSettingsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *UserAccountSettingsMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *UserAccountSettingsMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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
func (m *UserAccountSettingsMock) NewRequest(ctx context.Context) *client.RequestBuilder {
	return client.NewMockRequestBuilder(ctx, func(method, path string, result any) (*resty.Response, error) {
		return m.dispatch(method, path, result)
	})
}

func (m *UserAccountSettingsMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *UserAccountSettingsMock) InvalidateToken() error                { return nil }
func (m *UserAccountSettingsMock) KeepAliveToken() error                 { return nil }
func (m *UserAccountSettingsMock) GetLogger() *zap.Logger                { return m.logger }
