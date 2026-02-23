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

// NotificationsMock is a test double implementing interfaces.HTTPClient.
type NotificationsMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewNotificationsMock returns an empty mock ready for response registration.
func NewNotificationsMock() *NotificationsMock {
	return &NotificationsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func (m *NotificationsMock) register(method, path string, statusCode int, rawJSON string) {
	key := method + " " + path
	m.responses[key] = registeredResponse{
		statusCode: statusCode,
		rawBody:    []byte(rawJSON),
	}
}

// RegisterGetNotificationsMock registers a successful response for GetForUserAndSiteV1.
func (m *NotificationsMock) RegisterGetNotificationsMock() {
	m.register("GET", "/api/v1/notifications", 200, `[{
		"type": "SYSTEM_ALERT",
		"id": "notification-1",
		"params": {
			"message": "Test notification",
			"severity": "info"
		}
	}]`)
}

// Get implements interfaces.HTTPClient.
func (m *NotificationsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*interfaces.Response, error) {
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
	return &interfaces.Response{
		StatusCode: resp.statusCode,
		Headers:    http.Header{},
		Body:       resp.rawBody,
	}, nil
}

// Post implements interfaces.HTTPClient.
func (m *NotificationsMock) Post(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
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
	return &interfaces.Response{
		StatusCode: resp.statusCode,
		Headers:    http.Header{},
		Body:       resp.rawBody,
	}, nil
}

// PostWithQuery implements interfaces.HTTPClient.
func (m *NotificationsMock) PostWithQuery(ctx context.Context, path string, rsqlQuery map[string]string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
	return m.Post(ctx, path, body, headers, result)
}

// PostForm implements interfaces.HTTPClient.
func (m *NotificationsMock) PostForm(ctx context.Context, path string, formData map[string]string, headers map[string]string, result any) (*interfaces.Response, error) {
	return m.Post(ctx, path, formData, headers, result)
}

// PostMultipart implements interfaces.HTTPClient.
func (m *NotificationsMock) PostMultipart(ctx context.Context, path string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.Post(ctx, path, nil, headers, result)
}

// Put implements interfaces.HTTPClient.
func (m *NotificationsMock) Put(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
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
	return &interfaces.Response{
		StatusCode: resp.statusCode,
		Headers:    http.Header{},
		Body:       resp.rawBody,
	}, nil
}

// Patch implements interfaces.HTTPClient.
func (m *NotificationsMock) Patch(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
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
	return &interfaces.Response{
		StatusCode: resp.statusCode,
		Headers:    http.Header{},
		Body:       resp.rawBody,
	}, nil
}

// Delete implements interfaces.HTTPClient.
func (m *NotificationsMock) Delete(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*interfaces.Response, error) {
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
	return &interfaces.Response{
		StatusCode: resp.statusCode,
		Headers:    http.Header{},
		Body:       resp.rawBody,
	}, nil
}

// DeleteWithBody implements interfaces.HTTPClient.
func (m *NotificationsMock) DeleteWithBody(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
	return m.Delete(ctx, path, nil, headers, result)
}

// GetBytes implements interfaces.HTTPClient.
func (m *NotificationsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string) (*interfaces.Response, []byte, error) {
	m.LastRSQLQuery = rsqlQuery
	key := "GET " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, nil, fmt.Errorf("no mock registered for GET %s", path)
	}
	if resp.errMsg != "" {
		return nil, nil, fmt.Errorf("%s", resp.errMsg)
	}
	return &interfaces.Response{
		StatusCode: resp.statusCode,
		Headers:    http.Header{},
		Body:       resp.rawBody,
	}, resp.rawBody, nil
}

// GetPaginated implements interfaces.HTTPClient.
func (m *NotificationsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, mergePage func(pageData []byte) error) (*interfaces.Response, error) {
	return nil, fmt.Errorf("GetPaginated not implemented in NotificationsMock")
}

// RSQLBuilder implements interfaces.HTTPClient.
func (m *NotificationsMock) RSQLBuilder() interfaces.RSQLFilterBuilder {
	return nil
}

// InvalidateToken implements interfaces.HTTPClient.
func (m *NotificationsMock) InvalidateToken() error {
	return nil
}

// KeepAliveToken implements interfaces.HTTPClient.
func (m *NotificationsMock) KeepAliveToken() error {
	return nil
}

// GetLogger implements interfaces.HTTPClient.
func (m *NotificationsMock) GetLogger() *zap.Logger {
	return m.logger
}
