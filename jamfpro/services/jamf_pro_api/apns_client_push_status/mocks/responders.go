package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"resty.dev/v3"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
)

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// APNSClientPushStatusMock is a test double implementing interfaces.HTTPClient.
type APNSClientPushStatusMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewAPNSClientPushStatusMock returns an empty mock ready for response registration.
func NewAPNSClientPushStatusMock() *APNSClientPushStatusMock {
	return &APNSClientPushStatusMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *APNSClientPushStatusMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterEnableAllClientsMock()
	m.RegisterGetEnableAllClientsStatusMock()
	m.RegisterEnableClientMock()
}

func (m *APNSClientPushStatusMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("APNSClientPushStatusMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// RegisterListMock registers the list APNS client push status response.
func (m *APNSClientPushStatusMock) RegisterListMock() {
	m.register("GET", "/api/v1/apns-client-push-status", 200, "validate_list.json")
}

// RegisterEnableAllClientsMock registers the enable-all-clients POST response (no content).
func (m *APNSClientPushStatusMock) RegisterEnableAllClientsMock() {
	m.register("POST", "/api/v1/apns-client-push-status/enable-all-clients", 202, "")
}

// RegisterGetEnableAllClientsStatusMock registers the enable-all-clients status GET response.
func (m *APNSClientPushStatusMock) RegisterGetEnableAllClientsStatusMock() {
	m.register("GET", "/api/v1/apns-client-push-status/enable-all-clients/status", 200, "validate_get_enable_all_clients_status.json")
}

// RegisterEnableClientMock registers the enable-client POST response (204 No Content).
func (m *APNSClientPushStatusMock) RegisterEnableClientMock() {
	m.register("POST", "/api/v1/apns-client-push-status/enable-client", 204, "")
}

// loadMockResponse loads a JSON fixture from the mocks directory.
func loadMockResponse(filename string) ([]byte, error) {
	wd, _ := os.Getwd()
	return os.ReadFile(filepath.Join(wd, "mocks", filename))
}

// Get implements interfaces.HTTPClient.Get.
func (m *APNSClientPushStatusMock) Get(ctx context.Context, path string, query map[string]string, headers map[string]string, out any) (*resty.Response, error) {
	m.LastRSQLQuery = query
	key := "GET:" + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for GET %s", path)
	}

	if resp.errMsg != "" {
		return shared.NewMockResponse(resp.statusCode, http.Header{}, nil), fmt.Errorf("%s", resp.errMsg)
	}

	if out != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, out); err != nil {
			return nil, fmt.Errorf("failed to unmarshal mock response: %w", err)
		}
	}

	return shared.NewMockResponse(resp.statusCode, http.Header{}, nil), nil
}

// GetPaginated implements interfaces.HTTPClient.GetPaginated.
func (m *APNSClientPushStatusMock) GetPaginated(ctx context.Context, path string, query map[string]string, headers map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	m.LastRSQLQuery = query
	key := "GET:" + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for GET %s", path)
	}

	if resp.errMsg != "" {
		return shared.NewMockResponse(resp.statusCode, http.Header{}, nil), fmt.Errorf("%s", resp.errMsg)
	}

	if mergePage != nil && len(resp.rawBody) > 0 {
		// Extract "results" field from the response, just like the real GetPaginated does
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

	return shared.NewMockResponse(resp.statusCode, http.Header{}, nil), nil
}

// Post implements interfaces.HTTPClient.Post.
func (m *APNSClientPushStatusMock) Post(ctx context.Context, path string, body any, headers map[string]string, out any) (*resty.Response, error) {
	key := "POST:" + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for POST %s", path)
	}

	if resp.errMsg != "" {
		return shared.NewMockResponse(resp.statusCode, http.Header{}, nil), fmt.Errorf("%s", resp.errMsg)
	}

	if out != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, out); err != nil {
			return nil, fmt.Errorf("failed to unmarshal mock response: %w", err)
		}
	}

	return shared.NewMockResponse(resp.statusCode, http.Header{}, nil), nil
}

// Delete implements interfaces.HTTPClient.Delete.
func (m *APNSClientPushStatusMock) Delete(ctx context.Context, path string, query map[string]string, headers map[string]string, out any) (*resty.Response, error) {
	return nil, fmt.Errorf("Delete not implemented in APNSClientPushStatusMock")
}

// Put implements interfaces.HTTPClient.Put.
func (m *APNSClientPushStatusMock) Put(ctx context.Context, path string, body any, headers map[string]string, out any) (*resty.Response, error) {
	return nil, fmt.Errorf("Put not implemented in APNSClientPushStatusMock")
}

// Patch implements interfaces.HTTPClient.Patch.
func (m *APNSClientPushStatusMock) Patch(ctx context.Context, path string, body any, headers map[string]string, out any) (*resty.Response, error) {
	return nil, fmt.Errorf("Patch not implemented in APNSClientPushStatusMock")
}

// DownloadFile implements interfaces.HTTPClient.DownloadFile.
func (m *APNSClientPushStatusMock) DownloadFile(ctx context.Context, url string) (io.ReadCloser, *http.Response, error) {
	return nil, nil, fmt.Errorf("DownloadFile not implemented in APNSClientPushStatusMock")
}

// SetLogger implements interfaces.HTTPClient.SetLogger.
func (m *APNSClientPushStatusMock) SetLogger(logger *zap.Logger) {
	m.logger = logger
}

// GetLogger implements interfaces.HTTPClient.GetLogger.
func (m *APNSClientPushStatusMock) GetLogger() *zap.Logger {
	return m.logger
}

// DeleteWithBody implements interfaces.HTTPClient.DeleteWithBody.
func (m *APNSClientPushStatusMock) DeleteWithBody(ctx context.Context, path string, body any, headers map[string]string, out any) (*resty.Response, error) {
	return nil, fmt.Errorf("DeleteWithBody not implemented in APNSClientPushStatusMock")
}

// PostWithQuery implements interfaces.HTTPClient.PostWithQuery.
func (m *APNSClientPushStatusMock) PostWithQuery(ctx context.Context, path string, rsqlQuery map[string]string, body any, headers map[string]string, out any) (*resty.Response, error) {
	return nil, fmt.Errorf("PostWithQuery not implemented in APNSClientPushStatusMock")
}

// PostForm implements interfaces.HTTPClient.PostForm.
func (m *APNSClientPushStatusMock) PostForm(ctx context.Context, path string, formData map[string]string, headers map[string]string, out any) (*resty.Response, error) {
	return nil, fmt.Errorf("PostForm not implemented in APNSClientPushStatusMock")
}

// PostMultipart implements interfaces.HTTPClient.PostMultipart.
func (m *APNSClientPushStatusMock) PostMultipart(ctx context.Context, path string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback interfaces.MultipartProgressCallback, out any) (*resty.Response, error) {
	return nil, fmt.Errorf("PostMultipart not implemented in APNSClientPushStatusMock")
}

// GetBytes implements interfaces.HTTPClient.GetBytes.
func (m *APNSClientPushStatusMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string) (*resty.Response, []byte, error) {
	return nil, nil, fmt.Errorf("GetBytes not implemented in APNSClientPushStatusMock")
}

// RSQLBuilder implements interfaces.HTTPClient.RSQLBuilder.
func (m *APNSClientPushStatusMock) RSQLBuilder() interfaces.RSQLFilterBuilder {
	return nil
}

// InvalidateToken implements interfaces.HTTPClient.InvalidateToken.
func (m *APNSClientPushStatusMock) InvalidateToken() error {
	return nil
}

// KeepAliveToken implements interfaces.HTTPClient.KeepAliveToken.
func (m *APNSClientPushStatusMock) KeepAliveToken() error {
	return nil
}
