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

const endpointIssueTomcatSslCertificate = "/api/settings/issueTomcatSslCertificate"

type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

type TomcatSettingsMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewTomcatSettingsMock() *TomcatSettingsMock {
	return &TomcatSettingsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func mustMocksDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func (m *TomcatSettingsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := os.ReadFile(filepath.Join(mustMocksDir(), fixture))
		if err != nil {
			panic(fmt.Sprintf("TomcatSettingsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *TomcatSettingsMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	if err != nil {
		panic(fmt.Sprintf("TomcatSettingsMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *TomcatSettingsMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("TomcatSettingsMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {"application/json"}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("TomcatSettingsMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

func (m *TomcatSettingsMock) RegisterIssueTomcatSslCertificateMock() {
	m.register("POST", endpointIssueTomcatSslCertificate, 200, "")
}

func (m *TomcatSettingsMock) RegisterIssueTomcatSslCertificateErrorMock() {
	m.registerError("POST", endpointIssueTomcatSslCertificate, 500, "error_not_found.json")
}

func (m *TomcatSettingsMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}

func (m *TomcatSettingsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *TomcatSettingsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *TomcatSettingsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *TomcatSettingsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *TomcatSettingsMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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

func (m *TomcatSettingsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *TomcatSettingsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *TomcatSettingsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *TomcatSettingsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *TomcatSettingsMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *TomcatSettingsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *TomcatSettingsMock) InvalidateToken() error                    { return nil }
func (m *TomcatSettingsMock) KeepAliveToken() error                     { return nil }
func (m *TomcatSettingsMock) GetLogger() *zap.Logger                    { return m.logger }

func (m *TomcatSettingsMock) SetLogger(logger *zap.Logger) {
	m.logger = logger
}

func (m *TomcatSettingsMock) DownloadFile(ctx context.Context, url string) (io.ReadCloser, *http.Response, error) {
	return nil, nil, fmt.Errorf("DownloadFile not implemented in TomcatSettingsMock")
}
