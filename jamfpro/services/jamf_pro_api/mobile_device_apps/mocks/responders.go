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
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"go.uber.org/zap"
	"resty.dev/v3"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

type MobileDeviceAppsMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewMobileDeviceAppsMock() *MobileDeviceAppsMock {
	return &MobileDeviceAppsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func mustMocksDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func (m *MobileDeviceAppsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := os.ReadFile(filepath.Join(mustMocksDir(), fixture))
		if err != nil {
			panic(fmt.Sprintf("MobileDeviceAppsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *MobileDeviceAppsMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	if err != nil {
		panic(fmt.Sprintf("MobileDeviceAppsMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *MobileDeviceAppsMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("MobileDeviceAppsMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {"application/json"}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("MobileDeviceAppsMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

const reinstallAppConfigPath = "/api/v1/mobile-device-apps/reinstall-app-config"

func (m *MobileDeviceAppsMock) RegisterReinstallAppConfigMock() {
	m.register("POST", reinstallAppConfigPath, 204, "")
}

func (m *MobileDeviceAppsMock) RegisterNotFoundErrorMock() {
	m.registerError("POST", reinstallAppConfigPath, 404, "error_not_found.json")
}

func (m *MobileDeviceAppsMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}

func (m *MobileDeviceAppsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceAppsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *MobileDeviceAppsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *MobileDeviceAppsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MobileDeviceAppsMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	bodyBytes := resp.Bytes()
	if mergePage != nil && resp != nil && len(bodyBytes) > 0 {
		if err := mergePage(bodyBytes); err != nil {
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
	}
	return resp, nil
}

func (m *MobileDeviceAppsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceAppsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceAppsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceAppsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MobileDeviceAppsMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *MobileDeviceAppsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *MobileDeviceAppsMock) InvalidateToken() error                    { return nil }
func (m *MobileDeviceAppsMock) KeepAliveToken() error                     { return nil }
func (m *MobileDeviceAppsMock) GetLogger() *zap.Logger                    { return m.logger }

func (m *MobileDeviceAppsMock) SetLogger(logger *zap.Logger) {
	m.logger = logger
}

func (m *MobileDeviceAppsMock) DownloadFile(ctx context.Context, url string) (io.ReadCloser, *http.Response, error) {
	return nil, nil, fmt.Errorf("DownloadFile not implemented in MobileDeviceAppsMock")
}
