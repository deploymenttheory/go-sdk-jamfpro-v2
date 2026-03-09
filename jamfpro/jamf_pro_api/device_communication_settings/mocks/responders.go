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

type DeviceCommunicationSettingsMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewDeviceCommunicationSettingsMock() *DeviceCommunicationSettingsMock {
	return &DeviceCommunicationSettingsMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *DeviceCommunicationSettingsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		body, _ = os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *DeviceCommunicationSettingsMock) RegisterGetMock() {
	m.register("GET", "/api/v1/device-communication-settings", 200, "validate_get.json")
}

func (m *DeviceCommunicationSettingsMock) RegisterPutMock() {
	m.register("PUT", "/api/v1/device-communication-settings", 200, "validate_get.json")
}

func (m *DeviceCommunicationSettingsMock) RegisterGetHistoryMock() {
	m.register("GET", "/api/v1/device-communication-settings/history", 200, "validate_history.json")
}

func (m *DeviceCommunicationSettingsMock) RegisterAddHistoryNotesMock() {
	m.register("POST", "/api/v1/device-communication-settings/history", 201, "validate_add_history_note.json")
}

func (m *DeviceCommunicationSettingsMock) RegisterGetErrorMock() {
	body, _ := os.ReadFile(filepath.Join(mustMocksDir(), "validate_get.json"))
	m.responses["GET:/api/v1/device-communication-settings"] = registeredResponse{statusCode: 500, rawBody: body, errMsg: "Jamf Pro API error (500): server error"}
}

func (m *DeviceCommunicationSettingsMock) RegisterPutErrorMock() {
	body, _ := os.ReadFile(filepath.Join(mustMocksDir(), "validate_get.json"))
	m.responses["PUT:/api/v1/device-communication-settings"] = registeredResponse{statusCode: 500, rawBody: body, errMsg: "Jamf Pro API error (500): server error"}
}

func (m *DeviceCommunicationSettingsMock) RegisterGetHistoryErrorMock() {
	body, _ := os.ReadFile(filepath.Join(mustMocksDir(), "validate_history.json"))
	m.responses["GET:/api/v1/device-communication-settings/history"] = registeredResponse{statusCode: 500, rawBody: body, errMsg: "Jamf Pro API error (500): server error"}
}

func (m *DeviceCommunicationSettingsMock) RegisterGetHistoryInvalidJSONMock() {
	m.responses["GET:/api/v1/device-communication-settings/history"] = registeredResponse{statusCode: 200, rawBody: []byte(`{invalid json`)}
}

func (m *DeviceCommunicationSettingsMock) RegisterGetHistoryInvalidItemMock() {
	body, _ := os.ReadFile(filepath.Join(mustMocksDir(), "validate_history_invalid_item.json"))
	m.responses["GET:/api/v1/device-communication-settings/history"] = registeredResponse{statusCode: 200, rawBody: body}
}

func (m *DeviceCommunicationSettingsMock) RegisterAddHistoryNotesErrorMock() {
	body, _ := os.ReadFile(filepath.Join(mustMocksDir(), "validate_add_history_note.json"))
	m.responses["POST:/api/v1/device-communication-settings/history"] = registeredResponse{statusCode: 500, rawBody: body, errMsg: "Jamf Pro API error (500): server error"}
}

func (m *DeviceCommunicationSettingsMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("DeviceCommunicationSettingsMock: no response for %s %s", method, path)
	}
	headers := http.Header{"Content-Type": {"application/json"}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)
	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}
	if result != nil && len(r.rawBody) > 0 {
		_ = json.Unmarshal(r.rawBody, result)
	}
	return resp, nil
}

func mustMocksDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func (m *DeviceCommunicationSettingsMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *DeviceCommunicationSettingsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *DeviceCommunicationSettingsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *DeviceCommunicationSettingsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *DeviceCommunicationSettingsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *DeviceCommunicationSettingsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *DeviceCommunicationSettingsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *DeviceCommunicationSettingsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *DeviceCommunicationSettingsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *DeviceCommunicationSettingsMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}
func (m *DeviceCommunicationSettingsMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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
func (m *DeviceCommunicationSettingsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *DeviceCommunicationSettingsMock) InvalidateToken() error                    { return nil }
func (m *DeviceCommunicationSettingsMock) KeepAliveToken() error                     { return nil }
func (m *DeviceCommunicationSettingsMock) GetLogger() *zap.Logger                    { return m.logger }
