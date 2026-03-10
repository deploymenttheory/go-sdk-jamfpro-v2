package mocks

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"resty.dev/v3"

	mockhelpers "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"go.uber.org/zap"
)

//go:embed validate_list.json
var validateListJSON []byte

//go:embed validate_get.json
var validateGetJSON []byte

//go:embed validate_list_empty.json
var validateListEmptyJSON []byte

//go:embed validate_list_invalid.json
var validateListInvalidJSON []byte

type registeredResponse struct {
	method   string
	path     string
	response []byte
	status   int
}

type GroupsMock struct {
	responses []registeredResponse
}

func NewGroupsMock() *GroupsMock {
	return &GroupsMock{
		responses: make([]registeredResponse, 0),
	}
}

func (m *GroupsMock) dispatch(method, path string) ([]byte, int, bool) {
	var bestMatch *registeredResponse
	longestMatch := 0

	for i := range m.responses {
		r := &m.responses[i]
		if r.method == method && strings.HasPrefix(path, r.path) {
			if len(r.path) > longestMatch {
				longestMatch = len(r.path)
				bestMatch = r
			}
		}
	}

	if bestMatch != nil {
		return bestMatch.response, bestMatch.status, true
	}
	return nil, 0, false
}

func (m *GroupsMock) Get(ctx context.Context, endpoint string, params map[string]string, headers map[string]string, result any) (*resty.Response, error) {
	data, status, ok := m.dispatch("GET", endpoint)
	if !ok {
		return nil, fmt.Errorf("no mock registered for GET %s", endpoint)
	}

	if result != nil && data != nil {
		if err := json.Unmarshal(data, result); err != nil {
			return mockhelpers.NewMockResponse(http.StatusInternalServerError, http.Header{}, nil), err
		}
	}

	return mockhelpers.NewMockResponse(status, http.Header{}, nil), nil
}

func (m *GroupsMock) Post(ctx context.Context, endpoint string, body any, headers map[string]string, result any) (*resty.Response, error) {
	data, status, ok := m.dispatch("POST", endpoint)
	if !ok {
		return nil, fmt.Errorf("no mock registered for POST %s", endpoint)
	}

	if result != nil && data != nil {
		if err := json.Unmarshal(data, result); err != nil {
			return mockhelpers.NewMockResponse(http.StatusInternalServerError, http.Header{}, nil), err
		}
	}

	return mockhelpers.NewMockResponse(status, http.Header{}, nil), nil
}

func (m *GroupsMock) Put(ctx context.Context, endpoint string, body any, headers map[string]string, result any) (*resty.Response, error) {
	data, status, ok := m.dispatch("PUT", endpoint)
	if !ok {
		return nil, fmt.Errorf("no mock registered for PUT %s", endpoint)
	}

	if result != nil && data != nil {
		if err := json.Unmarshal(data, result); err != nil {
			return mockhelpers.NewMockResponse(http.StatusInternalServerError, http.Header{}, nil), err
		}
	}

	return mockhelpers.NewMockResponse(status, http.Header{}, nil), nil
}

func (m *GroupsMock) Patch(ctx context.Context, endpoint string, body any, headers map[string]string, result any) (*resty.Response, error) {
	data, status, ok := m.dispatch("PATCH", endpoint)
	if !ok {
		return nil, fmt.Errorf("no mock registered for PATCH %s", endpoint)
	}

	if result != nil && data != nil {
		if err := json.Unmarshal(data, result); err != nil {
			return mockhelpers.NewMockResponse(http.StatusInternalServerError, http.Header{}, nil), err
		}
	}

	return mockhelpers.NewMockResponse(status, http.Header{}, nil), nil
}

func (m *GroupsMock) Delete(ctx context.Context, endpoint string, params map[string]string, headers map[string]string, result any) (*resty.Response, error) {
	_, status, ok := m.dispatch("DELETE", endpoint)
	if !ok {
		return nil, fmt.Errorf("no mock registered for DELETE %s", endpoint)
	}

	return mockhelpers.NewMockResponse(status, http.Header{}, nil), nil
}

func (m *GroupsMock) DeleteWithBody(ctx context.Context, endpoint string, body any, headers map[string]string, result any) (*resty.Response, error) {
	return mockhelpers.NewMockResponse(http.StatusNotImplemented, http.Header{}, nil), fmt.Errorf("DeleteWithBody not implemented in mock")
}

func (m *GroupsMock) GetBytes(ctx context.Context, endpoint string, params map[string]string, headers map[string]string) (*resty.Response, []byte, error) {
	data, status, ok := m.dispatch("GET", endpoint)
	if !ok {
		return nil, nil, fmt.Errorf("no mock registered for GET %s", endpoint)
	}

	return mockhelpers.NewMockResponse(status, http.Header{}, nil), data, nil
}

func (m *GroupsMock) GetPaginated(ctx context.Context, endpoint string, params map[string]string, headers map[string]string, mergePage func(page []byte) error) (*resty.Response, error) {
	data, status, ok := m.dispatch("GET", endpoint)
	if !ok {
		return nil, fmt.Errorf("no mock registered for GET %s", endpoint)
	}
	if mergePage != nil && data != nil && len(data) > 0 {
		var page struct {
			Results json.RawMessage `json:"results"`
		}
		if err := json.Unmarshal(data, &page); err != nil {
			return mockhelpers.NewMockResponse(http.StatusInternalServerError, http.Header{}, nil), err
		}
		if err := mergePage(page.Results); err != nil {
			return mockhelpers.NewMockResponse(http.StatusInternalServerError, http.Header{}, nil), err
		}
	}
	return mockhelpers.NewMockResponse(status, http.Header{}, nil), nil
}

func (m *GroupsMock) PostMultipart(ctx context.Context, endpoint string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback client.MultipartProgressCallback, out any) (*resty.Response, error) {
	return mockhelpers.NewMockResponse(http.StatusNotImplemented, http.Header{}, nil), fmt.Errorf("PostMultipart not implemented in mock")
}

func (m *GroupsMock) PostWithQuery(ctx context.Context, endpoint string, queryParams map[string]string, body any, headers map[string]string, out any) (*resty.Response, error) {
	return mockhelpers.NewMockResponse(http.StatusNotImplemented, http.Header{}, nil), fmt.Errorf("PostWithQuery not implemented in mock")
}

func (m *GroupsMock) PostForm(ctx context.Context, endpoint string, formData map[string]string, headers map[string]string, out any) (*resty.Response, error) {
	return mockhelpers.NewMockResponse(http.StatusNotImplemented, http.Header{}, nil), fmt.Errorf("PostForm not implemented in mock")
}

func (m *GroupsMock) GetLogger() *zap.Logger {
	return nil
}

func (m *GroupsMock) RSQLBuilder() client.RSQLFilterBuilder {
	return nil
}

func (m *GroupsMock) InvalidateToken() error {
	return nil
}

func (m *GroupsMock) KeepAliveToken() error {
	return nil
}

func (m *GroupsMock) RegisterListMock() {
	m.responses = append(m.responses, registeredResponse{
		method:   "GET",
		path:     "/api/v1/groups",
		response: validateListJSON,
		status:   http.StatusOK,
	})
}

func (m *GroupsMock) RegisterGetByIDMock() {
	m.responses = append(m.responses, registeredResponse{
		method:   "GET",
		path:     "/api/v1/groups/",
		response: validateGetJSON,
		status:   http.StatusOK,
	})
}

func (m *GroupsMock) RegisterUpdateMock() {
	m.responses = append(m.responses, registeredResponse{
		method:   "PATCH",
		path:     "/api/v1/groups/",
		response: validateGetJSON,
		status:   http.StatusOK,
	})
}

func (m *GroupsMock) RegisterDeleteMock() {
	m.responses = append(m.responses, registeredResponse{
		method:   "DELETE",
		path:     "/api/v1/groups/",
		response: nil,
		status:   http.StatusNoContent,
	})
}

func (m *GroupsMock) RegisterEmptyListMock() {
	m.responses = append(m.responses, registeredResponse{
		method:   "GET",
		path:     "/api/v1/groups",
		response: validateListEmptyJSON,
		status:   http.StatusOK,
	})
}

// RegisterListInvalidJSONMock returns invalid JSON to trigger mergePage unmarshal/decode errors.
func (m *GroupsMock) RegisterListInvalidJSONMock() {
	m.responses = append(m.responses, registeredResponse{
		method:   "GET",
		path:     "/api/v1/groups",
		response: validateListInvalidJSON,
		status:   http.StatusOK,
	})
}
