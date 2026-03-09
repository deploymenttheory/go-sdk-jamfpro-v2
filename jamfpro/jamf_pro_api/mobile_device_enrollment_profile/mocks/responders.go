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
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"go.uber.org/zap"
	"resty.dev/v3"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

type MobileDeviceEnrollmentProfileMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewMobileDeviceEnrollmentProfileMock() *MobileDeviceEnrollmentProfileMock {
	return &MobileDeviceEnrollmentProfileMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func mustMocksDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func (m *MobileDeviceEnrollmentProfileMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := os.ReadFile(filepath.Join(mustMocksDir(), fixture))
		if err != nil {
			panic(fmt.Sprintf("MobileDeviceEnrollmentProfileMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *MobileDeviceEnrollmentProfileMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	if err != nil {
		panic(fmt.Sprintf("MobileDeviceEnrollmentProfileMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *MobileDeviceEnrollmentProfileMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("MobileDeviceEnrollmentProfileMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {"application/json"}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("MobileDeviceEnrollmentProfileMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

func (m *MobileDeviceEnrollmentProfileMock) RegisterGetDownloadProfileMock(id string) {
	path := "/api/v1/mobile-device-enrollment-profile/" + id + "/download-profile"
	m.register("GET", path, 200, "validate_download_profile.bin")
}

func (m *MobileDeviceEnrollmentProfileMock) RegisterNotFoundErrorMock(id string) {
	path := "/api/v1/mobile-device-enrollment-profile/" + id + "/download-profile"
	m.registerError("GET", path, 404, "error_not_found.json")
}

func (m *MobileDeviceEnrollmentProfileMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}

func (m *MobileDeviceEnrollmentProfileMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceEnrollmentProfileMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceEnrollmentProfileMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceEnrollmentProfileMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MobileDeviceEnrollmentProfileMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *MobileDeviceEnrollmentProfileMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *MobileDeviceEnrollmentProfileMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MobileDeviceEnrollmentProfileMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MobileDeviceEnrollmentProfileMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *MobileDeviceEnrollmentProfileMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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

func (m *MobileDeviceEnrollmentProfileMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *MobileDeviceEnrollmentProfileMock) InvalidateToken() error                    { return nil }
func (m *MobileDeviceEnrollmentProfileMock) KeepAliveToken() error                     { return nil }
func (m *MobileDeviceEnrollmentProfileMock) GetLogger() *zap.Logger                    { return m.logger }
