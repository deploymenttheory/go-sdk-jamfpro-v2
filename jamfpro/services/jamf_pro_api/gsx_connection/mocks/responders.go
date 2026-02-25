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

// GSXConnectionMock is a test double implementing interfaces.HTTPClient.
type GSXConnectionMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewGSXConnectionMock returns an empty mock ready for response registration.
func NewGSXConnectionMock() *GSXConnectionMock {
	return &GSXConnectionMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *GSXConnectionMock) RegisterMocks() {
	m.RegisterGetGSXConnectionMock()
	m.RegisterUpdateGSXConnectionMock()
	m.RegisterGetHistoryMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *GSXConnectionMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
}

func (m *GSXConnectionMock) register(method, path string, statusCode int, body []byte) {
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *GSXConnectionMock) registerError(method, path string, statusCode int, body []byte) {
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *GSXConnectionMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/json"}},
			Body:       []byte(`{"code":"NOT-FOUND","message":"no mock registered"}`),
		}, fmt.Errorf("GSXConnectionMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("GSXConnectionMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

func (m *GSXConnectionMock) RegisterGetGSXConnectionMock() {
	body := []byte(`{
		"enabled": true,
		"username": "test@example.com",
		"serviceAccountNo": "12345",
		"shipToNo": "67890",
		"gsxKeystore": {
			"name": "certificate.p12",
			"expirationEpoch": 1691954900000,
			"errorMessage": ""
		}
	}`)
	m.register("GET", "/api/v1/gsx-connection", 200, body)
}

func (m *GSXConnectionMock) RegisterUpdateGSXConnectionMock() {
	body := []byte(`{
		"enabled": false,
		"username": "updated@example.com",
		"serviceAccountNo": "54321",
		"shipToNo": "09876",
		"gsxKeystore": {
			"name": "certificate.p12",
			"expirationEpoch": 1691954900000,
			"errorMessage": ""
		}
	}`)
	m.register("PATCH", "/api/v1/gsx-connection", 200, body)
}

func (m *GSXConnectionMock) RegisterGetHistoryMock() {
	body := []byte(`{
		"totalCount": 2,
		"results": [
			{
				"id": "1",
				"username": "admin",
				"date": "2024-01-15T10:30:00Z",
				"note": "GSX connection enabled",
				"details": "Enabled GSX connection for service account"
			},
			{
				"id": "2",
				"username": "admin",
				"date": "2024-01-16T14:20:00Z",
				"note": "Updated service account",
				"details": "Changed service account number"
			}
		]
	}`)
	m.register("GET", "/api/v1/gsx-connection/history", 200, body)
}

func (m *GSXConnectionMock) RegisterNotFoundErrorMock() {
	body := []byte(`{
		"code": "NOT-FOUND",
		"message": "GSX connection not found"
	}`)
	m.registerError("GET", "/api/v1/gsx-connection", 404, body)
}

func (m *GSXConnectionMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *GSXConnectionMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *GSXConnectionMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *GSXConnectionMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *GSXConnectionMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *GSXConnectionMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *GSXConnectionMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *GSXConnectionMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *GSXConnectionMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *GSXConnectionMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *GSXConnectionMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
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

func (m *GSXConnectionMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *GSXConnectionMock) InvalidateToken() error                    { return nil }
func (m *GSXConnectionMock) KeepAliveToken() error                     { return nil }
func (m *GSXConnectionMock) GetLogger() *zap.Logger                    { return m.logger }
