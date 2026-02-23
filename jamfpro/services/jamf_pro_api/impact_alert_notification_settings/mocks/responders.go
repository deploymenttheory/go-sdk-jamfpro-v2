package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
)

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// ImpactAlertNotificationSettingsMock is a test double implementing interfaces.HTTPClient.
type ImpactAlertNotificationSettingsMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewImpactAlertNotificationSettingsMock returns an empty mock ready for response registration.
func NewImpactAlertNotificationSettingsMock() *ImpactAlertNotificationSettingsMock {
	return &ImpactAlertNotificationSettingsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *ImpactAlertNotificationSettingsMock) RegisterMocks() {
	m.RegisterGetMock()
	m.RegisterUpdateMock()
}

func (m *ImpactAlertNotificationSettingsMock) register(method, path string, statusCode int, body []byte) {
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *ImpactAlertNotificationSettingsMock) registerError(method, path string, statusCode int, body []byte) {
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *ImpactAlertNotificationSettingsMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/json"}},
			Body:       []byte(`{"code":"NOT-FOUND","message":"no mock registered"}`),
		}, fmt.Errorf("ImpactAlertNotificationSettingsMock: no response registered for %s %s", method, path)
	}

	resp := &interfaces.Response{
		StatusCode: r.statusCode,
		Status:     fmt.Sprintf("%d", r.statusCode),
		Headers:    http.Header{"Content-Type": {"application/json"}},
		Body:       r.rawBody,
	}

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("ImpactAlertNotificationSettingsMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

func (m *ImpactAlertNotificationSettingsMock) RegisterGetMock() {
	body := []byte(`{
		"scopeableObjectsAlertEnabled": true,
		"scopeableObjectsConfirmationCodeEnabled": false,
		"deployableObjectsAlertEnabled": true,
		"deployableObjectsConfirmationCodeEnabled": false
	}`)
	m.register("GET", "/api/v1/impact-alert-notification-settings", 200, body)
}

func (m *ImpactAlertNotificationSettingsMock) RegisterUpdateMock() {
	// Update returns 204 No Content
	m.register("PUT", "/api/v1/impact-alert-notification-settings", 204, []byte{})
}

func (m *ImpactAlertNotificationSettingsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}

func (m *ImpactAlertNotificationSettingsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ImpactAlertNotificationSettingsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ImpactAlertNotificationSettingsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ImpactAlertNotificationSettingsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ImpactAlertNotificationSettingsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *ImpactAlertNotificationSettingsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *ImpactAlertNotificationSettingsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *ImpactAlertNotificationSettingsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *ImpactAlertNotificationSettingsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *ImpactAlertNotificationSettingsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil {
		if err := mergePage(resp.Body); err != nil {
			return resp, err
		}
	}
	return resp, nil
}

func (m *ImpactAlertNotificationSettingsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *ImpactAlertNotificationSettingsMock) InvalidateToken() error                    { return nil }
func (m *ImpactAlertNotificationSettingsMock) KeepAliveToken() error                     { return nil }
func (m *ImpactAlertNotificationSettingsMock) GetLogger() *zap.Logger                    { return m.logger }
