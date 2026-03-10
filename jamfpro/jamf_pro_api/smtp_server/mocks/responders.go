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

type SMTPServerMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewSMTPServerMock() *SMTPServerMock {
	return &SMTPServerMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *SMTPServerMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := os.ReadFile(filepath.Join(mustMocksDir(), fixture))
		if err != nil {
			panic(fmt.Sprintf("SMTPServerMock: load %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *SMTPServerMock) RegisterMocks() {
	m.register("GET", "/api/v2/smtp-server", 200, "validate_get.json")
	m.register("PUT", "/api/v2/smtp-server", 200, "validate_get.json")
	m.register("GET", "/api/v1/smtp-server/history", 200, "validate_history.json")
	m.register("POST", "/api/v1/smtp-server/history", 201, "validate_add_history_note.json")
	m.register("POST", "/api/v1/smtp-server/test", 202, "")
}

func (m *SMTPServerMock) registerError(method, path string, statusCode int, fixture string, errMsg string) {
	var body []byte
	if fixture != "" {
		data, _ := os.ReadFile(filepath.Join(mustMocksDir(), fixture))
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *SMTPServerMock) RegisterGetErrorMock() {
	m.registerError("GET", "/api/v2/smtp-server", 500, "validate_get.json", "Jamf Pro API error (500): server error")
}

func (m *SMTPServerMock) RegisterPutErrorMock() {
	m.registerError("PUT", "/api/v2/smtp-server", 500, "validate_get.json", "Jamf Pro API error (500): server error")
}

func (m *SMTPServerMock) RegisterGetHistoryErrorMock() {
	m.registerError("GET", "/api/v1/smtp-server/history", 500, "validate_history.json", "Jamf Pro API error (500): server error")
}

func (m *SMTPServerMock) RegisterGetHistoryInvalidJSONMock() {
	m.responses["GET:/api/v1/smtp-server/history"] = registeredResponse{statusCode: 200, rawBody: []byte(`{invalid json`)}
}

func (m *SMTPServerMock) RegisterGetHistoryInvalidItemMock() {
	body, _ := os.ReadFile(filepath.Join(mustMocksDir(), "validate_history_invalid_item.json"))
	m.responses["GET:/api/v1/smtp-server/history"] = registeredResponse{statusCode: 200, rawBody: body}
}

func (m *SMTPServerMock) RegisterAddHistoryNoteErrorMock() {
	m.registerError("POST", "/api/v1/smtp-server/history", 500, "validate_add_history_note.json", "Jamf Pro API error (500): server error")
}

func (m *SMTPServerMock) RegisterTestErrorMock() {
	m.responses["POST:/api/v1/smtp-server/test"] = registeredResponse{statusCode: 500, rawBody: nil, errMsg: "Jamf Pro API error (500): server error"}
}

func (m *SMTPServerMock) Get(ctx context.Context, path string, q map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *SMTPServerMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *SMTPServerMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *SMTPServerMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *SMTPServerMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *SMTPServerMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *SMTPServerMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *SMTPServerMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *SMTPServerMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *SMTPServerMock) GetBytes(ctx context.Context, path string, q map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}
func (m *SMTPServerMock) GetPaginated(ctx context.Context, path string, q map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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
func (m *SMTPServerMock) NewRequest(ctx context.Context) *client.RequestBuilder {
	return client.NewMockRequestBuilder(ctx, func(method, path string, result any) (*resty.Response, error) {
		return m.dispatch(method, path, result)
	})
}
func (m *SMTPServerMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *SMTPServerMock) InvalidateToken() error                    { return nil }
func (m *SMTPServerMock) KeepAliveToken() error                     { return nil }
func (m *SMTPServerMock) GetLogger() *zap.Logger                    { return m.logger }

func (m *SMTPServerMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("SMTPServerMock: no response for %s %s", method, path)
	}
	headers := http.Header{"Content-Type": {"application/json"}}
	resp := mockhelpers.NewMockResponse(r.statusCode, headers, r.rawBody)
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
