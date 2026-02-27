package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
)

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// ActivationCodeMock is a test double implementing interfaces.HTTPClient.
type ActivationCodeMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewActivationCodeMock returns an empty mock ready for response registration.
func NewActivationCodeMock() *ActivationCodeMock {
	return &ActivationCodeMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *ActivationCodeMock) RegisterMocks() {
	m.RegisterGetHistoryMock()
	m.RegisterUpdateActivationCodeMock()
	m.RegisterUpdateOrganizationNameMock()
	m.RegisterAddHistoryNoteMock()
	m.RegisterExportHistoryMock()
}

func (m *ActivationCodeMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("ActivationCodeMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// RegisterGetHistoryMock registers the get activation code history response.
func (m *ActivationCodeMock) RegisterGetHistoryMock() {
	m.register("GET", "/api/v1/activation-code/history", 200, "validate_history.json")
}

// RegisterUpdateActivationCodeMock registers the update activation code response.
func (m *ActivationCodeMock) RegisterUpdateActivationCodeMock() {
	m.register("PUT", "/api/v1/activation-code", 202, "")
}

// RegisterUpdateOrganizationNameMock registers the update organization name response.
func (m *ActivationCodeMock) RegisterUpdateOrganizationNameMock() {
	m.register("PATCH", "/api/v1/activation-code/organization-name", 202, "")
}

// RegisterAddHistoryNoteMock registers the add history note response.
func (m *ActivationCodeMock) RegisterAddHistoryNoteMock() {
	m.register("POST", "/api/v1/activation-code/history", 201, "validate_add_history_note.json")
}

// RegisterExportHistoryMock registers the export history response.
func (m *ActivationCodeMock) RegisterExportHistoryMock() {
	csvData := "id,username,date,note,details\n1,admin,2019-02-04 21:09:31,Buildings update,Some details\n"
	m.responses["POST:"+"/api/v1/activation-code/history/export"] = registeredResponse{
		statusCode: 200,
		rawBody:    []byte(csvData),
	}
}

// loadMockResponse loads a JSON fixture from the mocks directory.
func loadMockResponse(filename string) ([]byte, error) {
	wd, _ := os.Getwd()
	return os.ReadFile(filepath.Join(wd, "mocks", filename))
}

// Get implements interfaces.HTTPClient.Get.
func (m *ActivationCodeMock) Get(ctx context.Context, path string, query map[string]string, headers map[string]string, out any) (*interfaces.Response, error) {
	m.LastRSQLQuery = query
	key := "GET:" + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for GET %s", path)
	}

	if resp.errMsg != "" {
		return &interfaces.Response{StatusCode: resp.statusCode}, fmt.Errorf("%s", resp.errMsg)
	}

	if out != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, out); err != nil {
			return nil, fmt.Errorf("failed to unmarshal mock response: %w", err)
		}
	}

	return &interfaces.Response{StatusCode: resp.statusCode}, nil
}

// GetPaginated implements interfaces.HTTPClient.GetPaginated.
func (m *ActivationCodeMock) GetPaginated(ctx context.Context, path string, query map[string]string, headers map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	m.LastRSQLQuery = query
	key := "GET:" + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for GET %s", path)
	}

	if resp.errMsg != "" {
		return &interfaces.Response{StatusCode: resp.statusCode}, fmt.Errorf("%s", resp.errMsg)
	}

	if len(resp.rawBody) > 0 {
		// Parse the paginated response structure to extract the results field
		var pageResp struct {
			Results json.RawMessage `json:"results"`
		}
		if err := json.Unmarshal(resp.rawBody, &pageResp); err != nil {
			return nil, fmt.Errorf("failed to unmarshal paginated response: %w", err)
		}
		
		if err := mergePage(pageResp.Results); err != nil {
			return nil, fmt.Errorf("mergePage failed: %w", err)
		}
	}

	return &interfaces.Response{StatusCode: resp.statusCode}, nil
}

// Post implements interfaces.HTTPClient.Post.
func (m *ActivationCodeMock) Post(ctx context.Context, path string, body any, headers map[string]string, out any) (*interfaces.Response, error) {
	key := "POST:" + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for POST %s", path)
	}

	if resp.errMsg != "" {
		return &interfaces.Response{StatusCode: resp.statusCode}, fmt.Errorf("%s", resp.errMsg)
	}

	if out != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, out); err != nil {
			return nil, fmt.Errorf("failed to unmarshal mock response: %w", err)
		}
	}

	return &interfaces.Response{StatusCode: resp.statusCode}, nil
}

// Delete implements interfaces.HTTPClient.Delete.
func (m *ActivationCodeMock) Delete(ctx context.Context, path string, query map[string]string, headers map[string]string, out any) (*interfaces.Response, error) {
	return nil, fmt.Errorf("Delete not implemented in ActivationCodeMock")
}

// Put implements interfaces.HTTPClient.Put.
func (m *ActivationCodeMock) Put(ctx context.Context, path string, body any, headers map[string]string, out any) (*interfaces.Response, error) {
	key := "PUT:" + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for PUT %s", path)
	}

	if resp.errMsg != "" {
		return &interfaces.Response{StatusCode: resp.statusCode}, fmt.Errorf("%s", resp.errMsg)
	}

	return &interfaces.Response{StatusCode: resp.statusCode}, nil
}

// Patch implements interfaces.HTTPClient.Patch.
func (m *ActivationCodeMock) Patch(ctx context.Context, path string, body any, headers map[string]string, out any) (*interfaces.Response, error) {
	key := "PATCH:" + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for PATCH %s", path)
	}

	if resp.errMsg != "" {
		return &interfaces.Response{StatusCode: resp.statusCode}, fmt.Errorf("%s", resp.errMsg)
	}

	return &interfaces.Response{StatusCode: resp.statusCode}, nil
}

// DownloadFile implements interfaces.HTTPClient.DownloadFile.
func (m *ActivationCodeMock) DownloadFile(ctx context.Context, url string) (io.ReadCloser, *http.Response, error) {
	return nil, nil, fmt.Errorf("DownloadFile not implemented in ActivationCodeMock")
}

// SetLogger implements interfaces.HTTPClient.SetLogger.
func (m *ActivationCodeMock) SetLogger(logger *zap.Logger) {
	m.logger = logger
}

// GetLogger implements interfaces.HTTPClient.GetLogger.
func (m *ActivationCodeMock) GetLogger() *zap.Logger {
	return m.logger
}

// DeleteWithBody implements interfaces.HTTPClient.DeleteWithBody.
func (m *ActivationCodeMock) DeleteWithBody(ctx context.Context, path string, body any, headers map[string]string, out any) (*interfaces.Response, error) {
	return nil, fmt.Errorf("DeleteWithBody not implemented in ActivationCodeMock")
}

// PostWithQuery implements interfaces.HTTPClient.PostWithQuery.
func (m *ActivationCodeMock) PostWithQuery(ctx context.Context, path string, rsqlQuery map[string]string, body any, headers map[string]string, out any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	key := "POST:" + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for POST %s", path)
	}

	if resp.errMsg != "" {
		return &interfaces.Response{StatusCode: resp.statusCode, Body: []byte(resp.errMsg)}, fmt.Errorf("%s", resp.errMsg)
	}

	return &interfaces.Response{StatusCode: resp.statusCode, Body: resp.rawBody}, nil
}

// PostForm implements interfaces.HTTPClient.PostForm.
func (m *ActivationCodeMock) PostForm(ctx context.Context, path string, formData map[string]string, headers map[string]string, out any) (*interfaces.Response, error) {
	return nil, fmt.Errorf("PostForm not implemented in ActivationCodeMock")
}

// PostMultipart implements interfaces.HTTPClient.PostMultipart.
func (m *ActivationCodeMock) PostMultipart(ctx context.Context, path string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback interfaces.MultipartProgressCallback, out any) (*interfaces.Response, error) {
	return nil, fmt.Errorf("PostMultipart not implemented in ActivationCodeMock")
}

// GetBytes implements interfaces.HTTPClient.GetBytes.
func (m *ActivationCodeMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string) (*interfaces.Response, []byte, error) {
	return nil, nil, fmt.Errorf("GetBytes not implemented in ActivationCodeMock")
}

// RSQLBuilder implements interfaces.HTTPClient.RSQLBuilder.
func (m *ActivationCodeMock) RSQLBuilder() interfaces.RSQLFilterBuilder {
	return nil
}

// InvalidateToken implements interfaces.HTTPClient.InvalidateToken.
func (m *ActivationCodeMock) InvalidateToken() error {
	return nil
}

// KeepAliveToken implements interfaces.HTTPClient.KeepAliveToken.
func (m *ActivationCodeMock) KeepAliveToken() error {
	return nil
}
