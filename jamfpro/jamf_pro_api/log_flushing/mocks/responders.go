package mocks

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"resty.dev/v3"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"go.uber.org/zap"
)

//go:embed validate_settings.json
var validateSettingsJSON []byte

//go:embed validate_tasks_list.json
var validateTasksListJSON []byte

//go:embed validate_task_get.json
var validateTaskGetJSON []byte

//go:embed validate_queue_task.json
var validateQueueTaskJSON []byte

var errNoMockRegistered = fmt.Errorf("no mock registered")

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

func (m *LogFlushingMock) register(method, path string, statusCode int, rawBody []byte) {
	key := method + " " + path
	m.responses[key] = registeredResponse{
		statusCode: statusCode,
		rawBody:    rawBody,
	}
}

func (m *LogFlushingMock) registerError(method, path string, errMsg string) {
	key := method + " " + path
	m.responses[key] = registeredResponse{
		errMsg: errMsg,
	}
}

func (m *LogFlushingMock) RegisterGetSettingsMock() {
	m.register("GET", "/api/v1/log-flushing", 200, validateSettingsJSON)
}

func (m *LogFlushingMock) RegisterGetSettingsErrorMock() {
	m.registerError("GET", "/api/v1/log-flushing", "api error")
}

func (m *LogFlushingMock) RegisterListTasksMock() {
	m.register("GET", "/api/v1/log-flushing/task", 200, validateTasksListJSON)
}

func (m *LogFlushingMock) RegisterListTasksErrorMock() {
	m.registerError("GET", "/api/v1/log-flushing/task", "api error")
}

func (m *LogFlushingMock) RegisterGetTaskByIDMock() {
	m.register("GET", "/api/v1/log-flushing/task/1", 200, validateTaskGetJSON)
}

func (m *LogFlushingMock) RegisterGetTaskByIDErrorMock(id string) {
	m.registerError("GET", "/api/v1/log-flushing/task/"+id, "api error")
}

func (m *LogFlushingMock) RegisterQueueTaskMock() {
	m.register("POST", "/api/v1/log-flushing/task", 201, validateQueueTaskJSON)
}

func (m *LogFlushingMock) RegisterQueueTaskErrorMock() {
	m.registerError("POST", "/api/v1/log-flushing/task", "api error")
}

func (m *LogFlushingMock) RegisterDeleteTaskMock() {
	m.register("DELETE", "/api/v1/log-flushing/task/1", 204, nil)
}

func (m *LogFlushingMock) RegisterDeleteTaskErrorMock(id string) {
	m.registerError("DELETE", "/api/v1/log-flushing/task/"+id, "api error")
}

func (m *LogFlushingMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	key := "GET " + path
	r, ok := m.responses[key]
	if !ok {
		return nil, errNoMockRegistered
	}
	if r.errMsg != "" {
		return nil, fmt.Errorf("%s", r.errMsg)
	}
	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return shared.NewMockResponse(r.statusCode, http.Header{}, r.rawBody), nil
}

func (m *LogFlushingMock) Post(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	key := "POST " + path
	r, ok := m.responses[key]
	if !ok {
		return nil, errNoMockRegistered
	}
	if r.errMsg != "" {
		return nil, fmt.Errorf("%s", r.errMsg)
	}
	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return shared.NewMockResponse(r.statusCode, http.Header{}, r.rawBody), nil
}

func (m *LogFlushingMock) PostWithQuery(ctx context.Context, path string, rsqlQuery map[string]string, body any, headers map[string]string, result any) (*resty.Response, error) {
	return m.Post(ctx, path, body, headers, result)
}

func (m *LogFlushingMock) PostForm(ctx context.Context, path string, formData map[string]string, headers map[string]string, result any) (*resty.Response, error) {
	return m.Post(ctx, path, formData, headers, result)
}

func (m *LogFlushingMock) PostMultipart(ctx context.Context, path string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.Post(ctx, path, nil, headers, result)
}

func (m *LogFlushingMock) Put(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	key := "PUT " + path
	r, ok := m.responses[key]
	if !ok {
		return nil, errNoMockRegistered
	}
	if r.errMsg != "" {
		return nil, fmt.Errorf("%s", r.errMsg)
	}
	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return shared.NewMockResponse(r.statusCode, http.Header{}, r.rawBody), nil
}

func (m *LogFlushingMock) Patch(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	key := "PATCH " + path
	r, ok := m.responses[key]
	if !ok {
		return nil, errNoMockRegistered
	}
	if r.errMsg != "" {
		return nil, fmt.Errorf("%s", r.errMsg)
	}
	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return shared.NewMockResponse(r.statusCode, http.Header{}, r.rawBody), nil
}

func (m *LogFlushingMock) Delete(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*resty.Response, error) {
	key := "DELETE " + path
	r, ok := m.responses[key]
	if !ok {
		return nil, errNoMockRegistered
	}
	if r.errMsg != "" {
		return nil, fmt.Errorf("%s", r.errMsg)
	}
	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return shared.NewMockResponse(r.statusCode, http.Header{}, r.rawBody), nil
}

func (m *LogFlushingMock) DeleteWithBody(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	return m.Delete(ctx, path, nil, headers, result)
}

func (m *LogFlushingMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string) (*resty.Response, []byte, error) {
	m.LastRSQLQuery = rsqlQuery
	key := "GET " + path
	r, ok := m.responses[key]
	if !ok {
		return nil, nil, errNoMockRegistered
	}
	if r.errMsg != "" {
		return nil, nil, fmt.Errorf("%s", r.errMsg)
	}
	return shared.NewMockResponse(r.statusCode, http.Header{}, r.rawBody), r.rawBody, nil
}

func (m *LogFlushingMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, mergePage func(pageData []byte) error) (*resty.Response, error) {
	return nil, fmt.Errorf("GetPaginated not implemented in LogFlushingMock")
}

func (m *LogFlushingMock) RSQLBuilder() client.RSQLFilterBuilder {
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
