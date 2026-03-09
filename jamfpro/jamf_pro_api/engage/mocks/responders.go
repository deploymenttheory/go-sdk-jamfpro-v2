package mocks

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"resty.dev/v3"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"go.uber.org/zap"
)

// errNoMockRegistered is returned when no mock is registered for the request.
var errNoMockRegistered = fmt.Errorf("no mock registered")

//go:embed validate_get.json
var validateGetJSON []byte

//go:embed validate_update.json
var validateUpdateJSON []byte

//go:embed validate_history.json
var validateHistoryJSON []byte

//go:embed validate_add_history_notes.json
var validateAddHistoryNotesJSON []byte

//go:embed validate_get_invalid.json
var validateGetInvalidJSON []byte

type registeredResponse struct {
	method   string
	path     string
	response []byte
	status   int
}

type EngageMock struct {
	responses []registeredResponse
}

func NewEngageMock() *EngageMock {
	return &EngageMock{
		responses: make([]registeredResponse, 0),
	}
}

func (m *EngageMock) dispatch(method, path string) ([]byte, int, bool) {
	for _, r := range m.responses {
		if r.method == method && strings.HasPrefix(path, r.path) {
			return r.response, r.status, true
		}
	}
	return nil, 0, false
}

func (m *EngageMock) Get(ctx context.Context, endpoint string, queryParams map[string]string, headers map[string]string, out any) (*resty.Response, error) {
	body, status, found := m.dispatch("GET", endpoint)
	if !found {
		return nil, errNoMockRegistered
	}

	if out != nil {
		if strPtr, ok := out.(*string); ok {
			*strPtr = string(body)
		} else if bytesPtr, ok := out.(*[]byte); ok {
			*bytesPtr = body
		} else if err := json.Unmarshal(body, out); err != nil {
			return shared.NewMockResponse(status, http.Header{}, nil), err
		}
	}

	return shared.NewMockResponse(status, http.Header{}, nil), nil
}

func (m *EngageMock) Post(ctx context.Context, endpoint string, body any, headers map[string]string, out any) (*resty.Response, error) {
	respBody, status, found := m.dispatch("POST", endpoint)
	if !found {
		return nil, errNoMockRegistered
	}

	if out != nil {
		if strPtr, ok := out.(*string); ok {
			*strPtr = string(respBody)
		} else if bytesPtr, ok := out.(*[]byte); ok {
			*bytesPtr = respBody
		} else if err := json.Unmarshal(respBody, out); err != nil {
			return shared.NewMockResponse(status, http.Header{}, nil), err
		}
	}

	return shared.NewMockResponse(status, http.Header{}, nil), nil
}

func (m *EngageMock) PostWithQuery(ctx context.Context, endpoint string, queryParams map[string]string, body any, headers map[string]string, out any) (*resty.Response, error) {
	return shared.NewMockResponse(http.StatusMethodNotAllowed, http.Header{}, nil), nil
}

func (m *EngageMock) PostForm(ctx context.Context, endpoint string, formData map[string]string, headers map[string]string, out any) (*resty.Response, error) {
	return shared.NewMockResponse(http.StatusMethodNotAllowed, http.Header{}, nil), nil
}

func (m *EngageMock) PostMultipart(ctx context.Context, endpoint string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback transport.MultipartProgressCallback, out any) (*resty.Response, error) {
	return shared.NewMockResponse(http.StatusMethodNotAllowed, http.Header{}, nil), nil
}

func (m *EngageMock) Put(ctx context.Context, endpoint string, body any, headers map[string]string, out any) (*resty.Response, error) {
	respBody, status, found := m.dispatch("PUT", endpoint)
	if !found {
		return nil, errNoMockRegistered
	}

	if out != nil {
		if strPtr, ok := out.(*string); ok {
			*strPtr = string(respBody)
		} else if bytesPtr, ok := out.(*[]byte); ok {
			*bytesPtr = respBody
		} else if err := json.Unmarshal(respBody, out); err != nil {
			return shared.NewMockResponse(status, http.Header{}, nil), err
		}
	}

	return shared.NewMockResponse(status, http.Header{}, nil), nil
}

func (m *EngageMock) Patch(ctx context.Context, endpoint string, body any, headers map[string]string, out any) (*resty.Response, error) {
	return shared.NewMockResponse(http.StatusMethodNotAllowed, http.Header{}, nil), nil
}

func (m *EngageMock) Delete(ctx context.Context, endpoint string, queryParams map[string]string, headers map[string]string, out any) (*resty.Response, error) {
	return shared.NewMockResponse(http.StatusMethodNotAllowed, http.Header{}, nil), nil
}

func (m *EngageMock) DeleteWithBody(ctx context.Context, endpoint string, body any, headers map[string]string, out any) (*resty.Response, error) {
	return shared.NewMockResponse(http.StatusMethodNotAllowed, http.Header{}, nil), nil
}

func (m *EngageMock) GetBytes(ctx context.Context, endpoint string, queryParams map[string]string, headers map[string]string) (*resty.Response, []byte, error) {
	body, status, found := m.dispatch("GET", endpoint)
	if !found {
		return nil, nil, errNoMockRegistered
	}

	return shared.NewMockResponse(status, http.Header{}, nil), body, nil
}

func (m *EngageMock) GetPaginated(ctx context.Context, endpoint string, rsqlQuery map[string]string, headers map[string]string, mergePage func(page []byte) error) (*resty.Response, error) {
	return shared.NewMockResponse(http.StatusMethodNotAllowed, http.Header{}, nil), nil
}

func (m *EngageMock) GetLogger() *zap.Logger {
	return nil
}

func (m *EngageMock) RSQLBuilder() transport.RSQLFilterBuilder {
	return nil
}

func (m *EngageMock) InvalidateToken() error {
	return nil
}

func (m *EngageMock) KeepAliveToken() error {
	return nil
}

func (m *EngageMock) RegisterGetMock() {
	m.responses = append(m.responses, registeredResponse{
		method:   "GET",
		path:     "/api/v2/engage",
		response: validateGetJSON,
		status:   http.StatusOK,
	})
}

// RegisterGetInvalidJSONMock registers a mock that returns invalid JSON for GET /api/v2/engage.
func (m *EngageMock) RegisterGetInvalidJSONMock() {
	m.responses = append(m.responses, registeredResponse{
		method:   "GET",
		path:     "/api/v2/engage",
		response: validateGetInvalidJSON,
		status:   http.StatusOK,
	})
}

func (m *EngageMock) RegisterUpdateMock() {
	m.responses = append(m.responses, registeredResponse{
		method:   "PUT",
		path:     "/api/v2/engage",
		response: validateUpdateJSON,
		status:   http.StatusOK,
	})
}

func (m *EngageMock) RegisterGetHistoryMock() {
	m.responses = append(m.responses, registeredResponse{
		method:   "GET",
		path:     "/api/v2/engage/history",
		response: validateHistoryJSON,
		status:   http.StatusOK,
	})
}

func (m *EngageMock) RegisterAddHistoryNotesMock() {
	m.responses = append(m.responses, registeredResponse{
		method:   "POST",
		path:     "/api/v2/engage/history",
		response: validateAddHistoryNotesJSON,
		status:   http.StatusCreated,
	})
}
