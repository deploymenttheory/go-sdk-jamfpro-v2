package mocks

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
)

//go:embed *.json
var fixtureFS embed.FS

type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// SLASAMock is a test double implementing interfaces.HTTPClient for SLASA operations.
// Responses are keyed by "METHOD path". Use RegisterGetStatusAcceptedMock, RegisterGetStatusNotAcceptedMock,
// and RegisterAcceptMock to set up responses.
type SLASAMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewSLASAMock returns an empty mock ready for response registration.
func NewSLASAMock() *SLASAMock {
	return &SLASAMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func (m *SLASAMock) register(method, path string, statusCode int, fixtureFile string) {
	key := method + " " + path
	var body []byte
	if fixtureFile != "" {
		data, err := fixtureFS.ReadFile(fixtureFile)
		if err != nil {
			panic(fmt.Sprintf("SLASAMock: failed to load fixture %q: %v", fixtureFile, err))
		}
		body = data
	}
	m.responses[key] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
	}
}

// RegisterGetStatusAcceptedMock registers GET /api/v1/slasa → 200 with ACCEPTED status.
func (m *SLASAMock) RegisterGetStatusAcceptedMock() {
	m.register("GET", "/api/v1/slasa", 200, "validate_get_status_accepted.json")
}

// RegisterGetStatusNotAcceptedMock registers GET /api/v1/slasa → 200 with NOT_ACCEPTED status.
func (m *SLASAMock) RegisterGetStatusNotAcceptedMock() {
	m.register("GET", "/api/v1/slasa", 200, "validate_get_status_not_accepted.json")
}

// RegisterAcceptMock registers POST /api/v1/slasa → 200 (no body).
func (m *SLASAMock) RegisterAcceptMock() {
	m.register("POST", "/api/v1/slasa", 200, "")
}

// RegisterGetStatusErrorMock registers GET with a 500 error.
func (m *SLASAMock) RegisterGetStatusErrorMock() {
	m.responses["GET /api/v1/slasa"] = registeredResponse{
		statusCode: 500,
		rawBody:    []byte(`{"error":"internal server error"}`),
		errMsg:     "request failed: 500 Internal Server Error",
	}
}

func (m *SLASAMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*interfaces.Response, error) {
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

func (m *SLASAMock) Post(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
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

func (m *SLASAMock) PostWithQuery(ctx context.Context, path string, rsqlQuery map[string]string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
	return m.Post(ctx, path, body, headers, result)
}

func (m *SLASAMock) PostForm(ctx context.Context, path string, formData map[string]string, headers map[string]string, result any) (*interfaces.Response, error) {
	return m.Post(ctx, path, formData, headers, result)
}

func (m *SLASAMock) PostMultipart(ctx context.Context, path string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.Post(ctx, path, nil, headers, result)
}

func (m *SLASAMock) Put(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
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

func (m *SLASAMock) Patch(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
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

func (m *SLASAMock) Delete(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*interfaces.Response, error) {
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

func (m *SLASAMock) DeleteWithBody(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
	return m.Delete(ctx, path, nil, headers, result)
}

func (m *SLASAMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string) (*interfaces.Response, []byte, error) {
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

func (m *SLASAMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, mergePage func(pageData []byte) error) (*interfaces.Response, error) {
	return nil, fmt.Errorf("GetPaginated not implemented in SLASAMock")
}

func (m *SLASAMock) RSQLBuilder() interfaces.RSQLFilterBuilder {
	return nil
}

func (m *SLASAMock) InvalidateToken() error {
	return nil
}

func (m *SLASAMock) KeepAliveToken() error {
	return nil
}

func (m *SLASAMock) GetLogger() *zap.Logger {
	return m.logger
}
