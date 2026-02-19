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

// NotificationsMock implements interfaces.HTTPClient.
type NotificationsMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewNotificationsMock() *NotificationsMock {
	return &NotificationsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func (m *NotificationsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		data, err := os.ReadFile(filepath.Join(dir, "mocks", fixture))
		if err != nil {
			panic(err)
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *NotificationsMock) RegisterMocks() {
	m.register("GET", "/api/v1/notifications", 200, "validate_list.json")
}

func (m *NotificationsMock) Get(ctx context.Context, path string, q map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *NotificationsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *NotificationsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *NotificationsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *NotificationsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *NotificationsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *NotificationsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *NotificationsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *NotificationsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *NotificationsMock) GetBytes(ctx context.Context, path string, q map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *NotificationsMock) GetPaginated(ctx context.Context, path string, q map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	return m.dispatch("GET", path, nil)
}
func (m *NotificationsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *NotificationsMock) InvalidateToken() error                    { return nil }
func (m *NotificationsMock) KeepAliveToken() error                     { return nil }
func (m *NotificationsMock) GetLogger() *zap.Logger                    { return m.logger }

func (m *NotificationsMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{StatusCode: 404}, fmt.Errorf("no mock for %s %s", method, path)
	}
	resp := &interfaces.Response{StatusCode: r.statusCode, Headers: http.Header{"Content-Type": {"application/json"}}, Body: r.rawBody}
	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}
	if result != nil && len(r.rawBody) > 0 {
		_ = json.Unmarshal(r.rawBody, result)
	}
	return resp, nil
}
