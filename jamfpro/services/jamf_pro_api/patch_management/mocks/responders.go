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

// PatchManagementMock is a test double implementing interfaces.HTTPClient for Patch Management operations.
// Responses are keyed by "METHOD path". Use RegisterAcceptDisclaimerMock to set up responses.
type PatchManagementMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewPatchManagementMock returns an empty mock ready for response registration.
func NewPatchManagementMock() *PatchManagementMock {
	return &PatchManagementMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func (m *PatchManagementMock) register(method, path string, statusCode int, fixtureFile string) {
	key := method + " " + path
	var body []byte
	if fixtureFile != "" {
		data, err := fixtureFS.ReadFile(fixtureFile)
		if err != nil {
			panic(fmt.Sprintf("PatchManagementMock: failed to load fixture %q: %v", fixtureFile, err))
		}
		body = data
	}
	m.responses[key] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
	}
}

// RegisterAcceptDisclaimerMock registers POST /api/v2/patch-management-accept-disclaimer → 200 (no body).
func (m *PatchManagementMock) RegisterAcceptDisclaimerMock() {
	m.register("POST", "/api/v2/patch-management-accept-disclaimer", 200, "")
}

// RegisterAcceptDisclaimerErrorMock registers POST with a 500 error.
func (m *PatchManagementMock) RegisterAcceptDisclaimerErrorMock() {
	m.responses["POST /api/v2/patch-management-accept-disclaimer"] = registeredResponse{
		statusCode: 500,
		rawBody:    []byte(`{"error":"internal server error"}`),
		errMsg:     "request failed: 500 Internal Server Error",
	}
}

func (m *PatchManagementMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*interfaces.Response, error) {
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

func (m *PatchManagementMock) Post(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
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

func (m *PatchManagementMock) PostWithQuery(ctx context.Context, path string, rsqlQuery map[string]string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
	return m.Post(ctx, path, body, headers, result)
}

func (m *PatchManagementMock) PostForm(ctx context.Context, path string, formData map[string]string, headers map[string]string, result any) (*interfaces.Response, error) {
	return m.Post(ctx, path, formData, headers, result)
}

func (m *PatchManagementMock) PostMultipart(ctx context.Context, path string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.Post(ctx, path, nil, headers, result)
}

func (m *PatchManagementMock) Put(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
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

func (m *PatchManagementMock) Patch(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
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

func (m *PatchManagementMock) Delete(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*interfaces.Response, error) {
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

func (m *PatchManagementMock) DeleteWithBody(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
	return m.Delete(ctx, path, nil, headers, result)
}

func (m *PatchManagementMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string) (*interfaces.Response, []byte, error) {
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

func (m *PatchManagementMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, mergePage func(pageData []byte) error) (*interfaces.Response, error) {
	return nil, fmt.Errorf("GetPaginated not implemented in PatchManagementMock")
}

func (m *PatchManagementMock) RSQLBuilder() interfaces.RSQLFilterBuilder {
	return nil
}

func (m *PatchManagementMock) InvalidateToken() error {
	return nil
}

func (m *PatchManagementMock) KeepAliveToken() error {
	return nil
}

func (m *PatchManagementMock) GetLogger() *zap.Logger {
	return m.logger
}
