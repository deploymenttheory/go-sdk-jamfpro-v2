package mocks

import (
	"context"
	_ "embed"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
)

//go:embed validate_get.json
var mockGetResponse []byte

type registeredResponse struct {
	method   string
	path     string
	response []byte
	status   int
}

type DSSDeclarationsMock struct {
	responses []registeredResponse
}

func NewDSSDeclarationsMock() *DSSDeclarationsMock {
	return &DSSDeclarationsMock{
		responses: make([]registeredResponse, 0),
	}
}

func (m *DSSDeclarationsMock) dispatch(method, path string) ([]byte, int, bool) {
	for _, r := range m.responses {
		if r.method == method && strings.HasPrefix(path, r.path) {
			return r.response, r.status, true
		}
	}
	return nil, 0, false
}

func (m *DSSDeclarationsMock) Get(ctx context.Context, endpoint string, queryParams map[string]string, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	body, status, found := m.dispatch("GET", endpoint)
	if !found {
		return &interfaces.Response{StatusCode: http.StatusNotFound}, nil
	}

	if out != nil {
		if err := json.Unmarshal(body, out); err != nil {
			return nil, err
		}
	}

	return &interfaces.Response{StatusCode: status}, nil
}

func (m *DSSDeclarationsMock) Post(ctx context.Context, endpoint string, body interface{}, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	return &interfaces.Response{StatusCode: http.StatusMethodNotAllowed}, nil
}

func (m *DSSDeclarationsMock) PostWithQuery(ctx context.Context, endpoint string, queryParams map[string]string, body interface{}, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	return &interfaces.Response{StatusCode: http.StatusMethodNotAllowed}, nil
}

func (m *DSSDeclarationsMock) PostForm(ctx context.Context, endpoint string, formData map[string]string, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	return &interfaces.Response{StatusCode: http.StatusMethodNotAllowed}, nil
}

func (m *DSSDeclarationsMock) PostMultipart(ctx context.Context, endpoint string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback interfaces.MultipartProgressCallback, out any) (*interfaces.Response, error) {
	return &interfaces.Response{StatusCode: http.StatusMethodNotAllowed}, nil
}

func (m *DSSDeclarationsMock) Put(ctx context.Context, endpoint string, body interface{}, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	return &interfaces.Response{StatusCode: http.StatusMethodNotAllowed}, nil
}

func (m *DSSDeclarationsMock) Patch(ctx context.Context, endpoint string, body interface{}, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	return &interfaces.Response{StatusCode: http.StatusMethodNotAllowed}, nil
}

func (m *DSSDeclarationsMock) Delete(ctx context.Context, endpoint string, queryParams map[string]string, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	return &interfaces.Response{StatusCode: http.StatusMethodNotAllowed}, nil
}

func (m *DSSDeclarationsMock) DeleteWithBody(ctx context.Context, endpoint string, body interface{}, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	return &interfaces.Response{StatusCode: http.StatusMethodNotAllowed}, nil
}

func (m *DSSDeclarationsMock) GetBytes(ctx context.Context, endpoint string, queryParams map[string]string, headers map[string]string) (*interfaces.Response, []byte, error) {
	body, status, found := m.dispatch("GET", endpoint)
	if !found {
		return &interfaces.Response{StatusCode: http.StatusNotFound}, nil, nil
	}

	return &interfaces.Response{StatusCode: status}, body, nil
}

func (m *DSSDeclarationsMock) GetPaginated(ctx context.Context, endpoint string, rsqlQuery map[string]string, headers map[string]string, mergePage func(page []byte) error) (*interfaces.Response, error) {
	return &interfaces.Response{StatusCode: http.StatusMethodNotAllowed}, nil
}

func (m *DSSDeclarationsMock) GetLogger() *zap.Logger {
	return nil
}

func (m *DSSDeclarationsMock) RSQLBuilder() interfaces.RSQLFilterBuilder {
	return nil
}

func (m *DSSDeclarationsMock) InvalidateToken() error {
	return nil
}

func (m *DSSDeclarationsMock) KeepAliveToken() error {
	return nil
}

func RegisterGetByUUIDMock(mock *DSSDeclarationsMock) {
	mock.responses = append(mock.responses, registeredResponse{
		method:   "GET",
		path:     "/api/v1/dss-declarations/",
		response: mockGetResponse,
		status:   http.StatusOK,
	})
}
