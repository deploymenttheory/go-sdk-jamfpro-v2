package mocks

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
)

//go:embed validate_get.json
var mockGetResponse []byte

//go:embed validate_update.json
var mockUpdateResponse []byte

//go:embed validate_history.json
var mockHistoryResponse []byte

//go:embed validate_error_not_found.json
var mockErrorNotFoundResponse []byte

//go:embed validate_history_invalid.json
var mockHistoryInvalidResponse []byte

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
		return nil, fmt.Errorf("GSXConnectionMock: no response registered for %s %s", method, path)
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
	m.register("GET", "/api/v1/gsx-connection", 200, mockGetResponse)
}

func (m *GSXConnectionMock) RegisterUpdateGSXConnectionMock() {
	m.register("PATCH", "/api/v1/gsx-connection", 200, mockUpdateResponse)
}

func (m *GSXConnectionMock) RegisterGetHistoryMock() {
	m.register("GET", "/api/v1/gsx-connection/history", 200, mockHistoryResponse)
}

func (m *GSXConnectionMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/api/v1/gsx-connection", 404, mockErrorNotFoundResponse)
}

// RegisterGetHistoryInvalidMock registers a response with invalid JSON to test mergePage unmarshal error.
func (m *GSXConnectionMock) RegisterGetHistoryInvalidMock() {
	m.register("GET", "/api/v1/gsx-connection/history", 200, mockHistoryInvalidResponse)
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
		_ = mergePage(resp.Body)
	}
	return resp, nil
}

func (m *GSXConnectionMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *GSXConnectionMock) InvalidateToken() error                    { return nil }
func (m *GSXConnectionMock) KeepAliveToken() error                     { return nil }
func (m *GSXConnectionMock) GetLogger() *zap.Logger                    { return m.logger }
