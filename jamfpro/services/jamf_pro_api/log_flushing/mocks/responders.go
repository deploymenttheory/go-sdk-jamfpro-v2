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

type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

type LogFlushingMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

func NewLogFlushingMock() *LogFlushingMock {
	return &LogFlushingMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func (m *LogFlushingMock) register(method, path string, statusCode int, rawJSON string) {
	key := method + " " + path
	m.responses[key] = registeredResponse{
		statusCode: statusCode,
		rawBody:    []byte(rawJSON),
	}
}

func (m *LogFlushingMock) RegisterGetSettingsMock() {
	m.register("GET", "/api/v1/log-flushing", 200, `{
		"retentionPolicies": [{
			"displayName": "Jamf Pro Server Logs",
			"qualifier": "JAMFSoftwareServer",
			"retentionPeriod": 30,
			"retentionPeriodUnit": "Days"
		}],
		"hourOfDay": 2
	}`)
}

func (m *LogFlushingMock) RegisterListTasksMock() {
	m.register("GET", "/api/v1/log-flushing/task", 200, `[{
		"id": "1",
		"qualifier": "JAMFSoftwareServer",
		"retentionPeriod": 30,
		"retentionPeriodUnit": "Days",
		"state": "COMPLETED"
	}]`)
}

func (m *LogFlushingMock) RegisterGetTaskByIDMock() {
	m.register("GET", "/api/v1/log-flushing/task/1", 200, `{
		"id": "1",
		"qualifier": "JAMFSoftwareServer",
		"retentionPeriod": 30,
		"retentionPeriodUnit": "Days",
		"state": "COMPLETED"
	}`)
}

func (m *LogFlushingMock) RegisterQueueTaskMock() {
	m.register("POST", "/api/v1/log-flushing/task", 201, `{
		"id": "2",
		"href": "/api/v1/log-flushing/task/2"
	}`)
}

func (m *LogFlushingMock) RegisterDeleteTaskMock() {
	m.register("DELETE", "/api/v1/log-flushing/task/1", 204, ``)
}

func (m *LogFlushingMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	key := "GET " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for GET %s", path)
	}
	if resp.errMsg != "" {
		return nil, fmt.Errorf("%s", resp.errMsg)
	}
	if result != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return &interfaces.Response{StatusCode: resp.statusCode, Headers: http.Header{}, Body: resp.rawBody}, nil
}

func (m *LogFlushingMock) Post(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
	key := "POST " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for POST %s", path)
	}
	if resp.errMsg != "" {
		return nil, fmt.Errorf("%s", resp.errMsg)
	}
	if result != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return &interfaces.Response{StatusCode: resp.statusCode, Headers: http.Header{}, Body: resp.rawBody}, nil
}

func (m *LogFlushingMock) PostWithQuery(ctx context.Context, path string, rsqlQuery map[string]string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
	return m.Post(ctx, path, body, headers, result)
}

func (m *LogFlushingMock) PostForm(ctx context.Context, path string, formData map[string]string, headers map[string]string, result any) (*interfaces.Response, error) {
	return m.Post(ctx, path, formData, headers, result)
}

func (m *LogFlushingMock) PostMultipart(ctx context.Context, path string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.Post(ctx, path, nil, headers, result)
}

func (m *LogFlushingMock) Put(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
	key := "PUT " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for PUT %s", path)
	}
	if resp.errMsg != "" {
		return nil, fmt.Errorf("%s", resp.errMsg)
	}
	if result != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return &interfaces.Response{StatusCode: resp.statusCode, Headers: http.Header{}, Body: resp.rawBody}, nil
}

func (m *LogFlushingMock) Patch(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
	key := "PATCH " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for PATCH %s", path)
	}
	if resp.errMsg != "" {
		return nil, fmt.Errorf("%s", resp.errMsg)
	}
	if result != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return &interfaces.Response{StatusCode: resp.statusCode, Headers: http.Header{}, Body: resp.rawBody}, nil
}

func (m *LogFlushingMock) Delete(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*interfaces.Response, error) {
	key := "DELETE " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for DELETE %s", path)
	}
	if resp.errMsg != "" {
		return nil, fmt.Errorf("%s", resp.errMsg)
	}
	if result != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return &interfaces.Response{StatusCode: resp.statusCode, Headers: http.Header{}, Body: resp.rawBody}, nil
}

func (m *LogFlushingMock) DeleteWithBody(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
	return m.Delete(ctx, path, nil, headers, result)
}

func (m *LogFlushingMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string) (*interfaces.Response, []byte, error) {
	m.LastRSQLQuery = rsqlQuery
	key := "GET " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, nil, fmt.Errorf("no mock registered for GET %s", path)
	}
	if resp.errMsg != "" {
		return nil, nil, fmt.Errorf("%s", resp.errMsg)
	}
	return &interfaces.Response{StatusCode: resp.statusCode, Headers: http.Header{}, Body: resp.rawBody}, resp.rawBody, nil
}

func (m *LogFlushingMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, mergePage func(pageData []byte) error) (*interfaces.Response, error) {
	return nil, fmt.Errorf("GetPaginated not implemented in LogFlushingMock")
}

func (m *LogFlushingMock) RSQLBuilder() interfaces.RSQLFilterBuilder {
	return nil
}

func (m *LogFlushingMock) InvalidateToken() error {
	return nil
}

func (m *LogFlushingMock) KeepAliveToken() error {
	return nil
}

func (m *LogFlushingMock) GetLogger() *zap.Logger {
	return m.logger
}
