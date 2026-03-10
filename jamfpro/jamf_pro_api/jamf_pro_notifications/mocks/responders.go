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
	"strings"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"go.uber.org/zap"
	"resty.dev/v3"

	mockhelpers "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// NotificationsMock is a test double implementing client.Client.
type NotificationsMock struct {
	responses       map[string]registeredResponse
	prefixResponses []struct {
		method string
		prefix string
		resp   registeredResponse
	}
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewNotificationsMock returns an empty mock ready for response registration.
func NewNotificationsMock() *NotificationsMock {
	return &NotificationsMock{
		responses:       make(map[string]registeredResponse),
		prefixResponses: nil,
		logger:          zap.NewNop(),
	}
}

func (m *NotificationsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("NotificationsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
	}
}

func (m *NotificationsMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("NotificationsMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *NotificationsMock) registerPrefix(method, pathPrefix string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("NotificationsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.prefixResponses = append(m.prefixResponses, struct {
		method string
		prefix string
		resp   registeredResponse
	}{
		method: method,
		prefix: pathPrefix,
		resp:   registeredResponse{statusCode: statusCode, rawBody: body},
	})
}

func (m *NotificationsMock) dispatch(method, path string, result any) (*resty.Response, error) {
	key := method + ":" + path
	if r, ok := m.responses[key]; ok {
		return m.buildResponse(r, result)
	}
	for _, pr := range m.prefixResponses {
		if pr.method == method && strings.HasPrefix(path, pr.prefix) {
			return m.buildResponse(pr.resp, result)
		}
	}
	return nil, fmt.Errorf("NotificationsMock: no response registered for %s %s", method, path)
}

func (m *NotificationsMock) buildResponse(r registeredResponse, result any) (*resty.Response, error) {
	headers := http.Header{"Content-Type": {"application/json"}}
	resp := mockhelpers.NewMockResponse(r.statusCode, headers, r.rawBody)
	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}
	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return resp, nil
}

// loadMockResponse reads a JSON fixture from the mocks directory adjacent to this file.
func loadMockResponse(filename string) ([]byte, error) {
	_, callerPath, _, ok := runtime.Caller(1)
	if !ok {
		return nil, fmt.Errorf("runtime.Caller failed")
	}
	dir := filepath.Dir(callerPath)
	data, err := os.ReadFile(filepath.Join(dir, filename))
	if err != nil {
		return nil, fmt.Errorf("read fixture %s: %w", filename, err)
	}
	return data, nil
}

// RegisterGetNotificationsMock registers a successful response for GetForUserAndSiteV1.
func (m *NotificationsMock) RegisterGetNotificationsMock() {
	m.register("GET", "/api/v1/notifications", 200, "validate_list.json")
}

// RegisterGetNotificationsEmptyMock registers an empty list response.
func (m *NotificationsMock) RegisterGetNotificationsEmptyMock() {
	m.register("GET", "/api/v1/notifications", 200, "validate_list_empty.json")
}

// RegisterGetNotificationsErrorMock registers an API error for GetForUserAndSiteV1.
func (m *NotificationsMock) RegisterGetNotificationsErrorMock() {
	m.registerError("GET", "/api/v1/notifications", 500, "error_api.json")
}

// RegisterDeleteNotificationMock registers a successful 204 response for DeleteByTypeAndIDV1.
// Uses path prefix matching so any /api/v1/notifications/{type}/{id} is matched.
func (m *NotificationsMock) RegisterDeleteNotificationMock() {
	m.registerPrefix("DELETE", "/api/v1/notifications/", 204, "")
}

// Get implements client.Client.
func (m *NotificationsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

// Post implements client.Client.
func (m *NotificationsMock) Post(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

// PostWithQuery implements client.Client.
func (m *NotificationsMock) PostWithQuery(ctx context.Context, path string, rsqlQuery map[string]string, body any, headers map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

// PostForm implements client.Client.
func (m *NotificationsMock) PostForm(ctx context.Context, path string, formData map[string]string, headers map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

// PostMultipart implements client.Client.
func (m *NotificationsMock) PostMultipart(ctx context.Context, path string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

// Put implements client.Client.
func (m *NotificationsMock) Put(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

// Patch implements client.Client.
func (m *NotificationsMock) Patch(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

// Delete implements client.Client.
func (m *NotificationsMock) Delete(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

// DeleteWithBody implements client.Client.
func (m *NotificationsMock) DeleteWithBody(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

// GetBytes implements client.Client.
func (m *NotificationsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string) (*resty.Response, []byte, error) {
	m.LastRSQLQuery = rsqlQuery
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

// GetPaginated implements client.Client.
func (m *NotificationsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, mergePage func(pageData []byte) error) (*resty.Response, error) {
	return nil, fmt.Errorf("GetPaginated not implemented in NotificationsMock")
}

// RSQLBuilder implements client.Client.
func (m *NotificationsMock) RSQLBuilder() client.RSQLFilterBuilder {
	return nil
}

// InvalidateToken implements client.Client.
func (m *NotificationsMock) InvalidateToken() error {
	return nil
}

// KeepAliveToken implements client.Client.
func (m *NotificationsMock) KeepAliveToken() error {
	return nil
}

// GetLogger implements client.Client.
func (m *NotificationsMock) GetLogger() *zap.Logger {
	return m.logger
}
